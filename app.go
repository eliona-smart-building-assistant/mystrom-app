//  This file is part of the eliona project.
//  Copyright © 2022 LEICOM iTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"context"
	"fmt"
	"mystrom/apiserver"
	"mystrom/apiservices"
	"mystrom/appdb"
	"mystrom/broker"
	"mystrom/conf"
	"mystrom/eliona"
	"mystrom/model"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/eliona-smart-building-assistant/go-eliona/app"
	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-eliona/dashboard"
	"github.com/eliona-smart-building-assistant/go-eliona/frontend"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/db"
	utilshttp "github.com/eliona-smart-building-assistant/go-utils/http"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func initialize() {
	ctx := context.Background()

	conn := db.NewInitConnectionWithContextAndApplicationName(ctx, app.AppName())
	defer conn.Close(ctx)

	app.Init(conn, app.AppName(),
		app.ExecSqlFile("conf/init.sql"),
		asset.InitAssetTypeFiles("resources/asset-types/*.json"),
		dashboard.InitWidgetTypeFiles("resources/widget-types/*.json"),
	)

	app.Patch(conn, app.AppName(), "010100",
		asset.InitAssetTypeFiles("resources/asset-types/*.json"),
	)
}

var once sync.Once

func collectData() {
	configs, err := conf.GetConfigs(context.Background())
	if err != nil {
		log.Fatal("conf", "Couldn't read configs from DB: %v", err)
		return
	}
	if len(configs) == 0 {
		once.Do(func() {
			log.Info("conf", "No configs in DB. Please configure the app in Eliona.")
		})
		return
	}

	for _, config := range configs {
		if !conf.IsConfigEnabled(config) {
			if conf.IsConfigActive(config) {
				conf.SetConfigActiveState(context.Background(), config, false)
			}
			continue
		}

		if !conf.IsConfigActive(config) {
			conf.SetConfigActiveState(context.Background(), config, true)
			log.Info("conf", "Collecting initialized with Configuration %d:\n"+
				"Enable: %t\n"+
				"Refresh Interval: %d\n"+
				"Request Timeout: %d\n"+
				"Project IDs: %v\n",
				*config.Id,
				*config.Enable,
				config.RefreshInterval,
				*config.RequestTimeout,
				*config.ProjectIDs)
		}

		common.RunOnceWithParam(func(config apiserver.Configuration) {
			log.Info("main", "Collecting %d started.", *config.Id)
			if err := collectResources(config); err != nil {
				return // Error is handled in the method itself.
			}
			pollTicker := time.NewTicker(time.Second * time.Duration(config.DataPollInterval))
			defer pollTicker.Stop()

			done := time.After(time.Second * time.Duration(config.RefreshInterval))
			for {
				select {
				case <-pollTicker.C:
					pollData(config)
				case <-done:
					log.Info("main", "Collecting %d finished.", *config.Id)
					return
				}
			}
		}, config, *config.Id)
	}
}

func collectResources(config apiserver.Configuration) error {
	var _ asset.FunctionalNode = (*model.Switch)(nil)
	root, err := broker.GetDevices(config)
	if err != nil {
		log.Error("broker", "getting root: %v", err)
		return err
	}
	if err := eliona.CreateAssets(config, &root); err != nil {
		log.Error("eliona", "creating assets: %v", err)
		return err
	}
	if err := eliona.UpsertSwitchData(config, root.GetDevices()); err != nil {
		log.Error("eliona", "inserting data into Eliona: %v", err)
		return err
	}
	return nil
}

func pollData(config apiserver.Configuration) {
	devices, err := broker.GetData(config)
	if err != nil {
		log.Error("broker", "getting data: %v", err)
		return
	}
	if err := eliona.UpsertSwitchData(config, devices); err != nil {
		log.Error("eliona", "inserting data into Eliona: %v", err)
		return
	}
}

// listenForOutputChanges listens to output attribute changes from Eliona.
func listenForOutputChanges() {
	for { // We want to restart listening in case something breaks.
		outputs, err := eliona.ListenForOutputChanges()
		if err != nil {
			log.Error("eliona", "listening for output changes: %v", err)
			return
		}
		for output := range outputs {
			if cr := output.ClientReference.Get(); cr != nil && *cr == eliona.ClientReference {
				// Just an echoed value this app sent.
				continue
			}
			asset, err := conf.GetAssetById(output.AssetId)
			if err != nil {
				log.Error("conf", "getting asset by assetID %v: %v", output.AssetId, err)
				continue
			}
			config, err := conf.GetConfigForAsset(asset)
			if err != nil {
				log.Error("conf", "getting configuration for asset id %v: %v", asset.AssetID.Int32, err)
				continue
			}
			if err := outputData(asset, config, output.Data); err != nil {
				log.Error("conf", "outputting data (%v) for config %v, assetId %v and device id %v: %v", output.Data, config.Id, asset.AssetID.Int32, asset.ProviderID, err)
				continue
			}
			pollData(config)
		}
		time.Sleep(time.Second * 5) // Give the server a little break.
	}
}

// outputData implements passing output data to broker.
func outputData(asset appdb.Asset, config apiserver.Configuration, data map[string]interface{}) error {
	val, ok := data["relay"]
	if !ok {
		return fmt.Errorf("data does not contain \"relay\": %v", data)
	}
	var value int64
	var err error
	switch v := val.(type) {
	case float64:
		value = int64(v)
	case string:
		if value, err = strconv.ParseInt(v, 10, 64); err != nil {
			return fmt.Errorf("output: parsing %v: %v", v, err)
		}
	default:
		return fmt.Errorf("output: got value of unknown type: (%T) %v", val, val)
	}

	return broker.PostData(config, asset.ProviderID, value)
}

// listenApi starts the API server and listen for requests
func listenApi() {
	err := http.ListenAndServe(":"+common.Getenv("API_SERVER_PORT", "3000"),
		frontend.NewEnvironmentHandler(
			utilshttp.NewCORSEnabledHandler(
				apiserver.NewRouter(
					apiserver.NewConfigurationAPIController(apiservices.NewConfigurationApiService()),
					apiserver.NewVersionAPIController(apiservices.NewVersionApiService()),
					apiserver.NewCustomizationAPIController(apiservices.NewCustomizationApiService()),
				))))
	log.Fatal("main", "API server: %v", err)
}
