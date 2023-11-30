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

package broker

import (
	"fmt"
	"mystrom/apiserver"
	nethttp "net/http"
	"net/url"
	"time"

	"github.com/eliona-smart-building-assistant/go-eliona/utils"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/http"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

type Switch struct {
	ID   string `eliona:"id" subtype:"info"`
	Name string `eliona:"name,filterable" subtype:"info"`

	Power float32 `eliona:"power" subtype:"input"`
	Temp  float32 `eliona:"temperature" subtype:"input"`

	Relay int `eliona:"relay" subtype:"output"`

	Room Room
}

func (s *Switch) AssetType() string {
	return "mystrom_switch"
}

func (s *Switch) Id() string {
	return s.AssetType() + "_" + s.ID
}

type Room struct {
	ID   string
	Name string
}

type devicesResponse struct {
	Devices []struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		Power          float32 `json:"power"`
		WifiSwitchTemp float32 `json:"wifiSwitchTemp"`
		State          string  `json:"state"`
		Type           string  `json:"type"`
		Room           struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"room"`
	} `json:"devices"`
	Status string `json:"status"`
}

func GetDevices(config apiserver.Configuration) ([]Switch, error) {
	// API v1 is called here for the rooms list. Be careful not to overuse it, though. No frequent
	// polling should be done to api v1.
	r, err := http.NewRequestWithApiKey("https://mystrom.ch/api/devices", "Auth-Token", config.ApiKey)
	if err != nil {
		return nil, fmt.Errorf("creating request for devices: %v", err)
	}
	resp, statusCode, err := http.ReadWithStatusCode[devicesResponse](r, time.Duration(*config.RequestTimeout)*time.Second, true)
	if err != nil {
		return nil, fmt.Errorf("querying API for devices: %v", err)
	}
	if statusCode != nethttp.StatusOK {
		return nil, fmt.Errorf("querying API for devices: got status %v", statusCode)
	}
	if resp.Status != "ok" {
		return nil, fmt.Errorf("API reports non-ok status: %v", resp.Status)
	}

	var devices []Switch
	for _, d := range resp.Devices {
		if d.Type != "ws2" && d.Type != "wse" {
			// We suport only WS2 and WSE smart plugs.
			continue
		}
		relayState := 0
		if d.State == "on" {
			relayState = 1
		}
		s := Switch{
			ID:    d.ID,
			Name:  d.Name,
			Power: d.Power,
			Temp:  d.WifiSwitchTemp,
			Relay: relayState,
			Room: struct {
				ID   string
				Name string
			}(d.Room),
		}
		if adheres, err := s.AdheresToFilter(config.AssetFilter); err != nil {
			return nil, fmt.Errorf("checking if adheres to filter: %v", err)
		} else if !adheres {
			continue
		}
		devices = append(devices, s)
	}

	return devices, nil
}

type devicesResponseV2 struct {
	Devices []struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Power       float32 `json:"power"`
		State       string  `json:"state"`
		Temperature float32 `json:"temperature"`
		Type        string  `json:"type"`
	} `json:"devices"`
}

func GetData(config apiserver.Configuration) ([]Switch, error) {
	// API v2 should be the preferred choice when communicating with myStrom. But ideally the
	// fetching of data should be done using webhooks.
	r, err := http.NewRequestWithApiKey("https://mystrom.ch/api/v2/devices", "Auth-Token", config.ApiKey)
	if err != nil {
		return nil, fmt.Errorf("creating request for devices: %v", err)
	}
	resp, statusCode, err := http.ReadWithStatusCode[devicesResponseV2](r, time.Duration(*config.RequestTimeout)*time.Second, true)
	if err != nil {
		return nil, fmt.Errorf("querying API for devices: %v", err)
	}
	if statusCode != nethttp.StatusOK {
		return nil, fmt.Errorf("querying API for devices: got status %v", statusCode)
	}

	var switches []Switch
	for _, device := range resp.Devices {
		if device.Type != "WS2" && device.Type != "WSE" {
			// We suport only WS2 and WSE smart plugs.
			continue
		}
		relayState := 0
		if device.State == "ON" {
			relayState = 1
		}
		switches = append(switches, Switch{
			ID:    device.ID,
			Name:  device.Name,
			Power: device.Power,
			Temp:  device.Temperature,
			Relay: relayState,
		})
	}

	return switches, nil
}

func PostData(config apiserver.Configuration, deviceID string, value int64) error {
	u, err := url.Parse(fmt.Sprintf("https://mystrom.ch/api/v2/device/%s", deviceID))
	if err != nil {
		return fmt.Errorf("shouldn't happen: parsing URL: %v", err)
	}
	q := u.Query()
	action := "on"
	if value == 0 {
		action = "off"
	}
	q.Set("action", action)
	u.RawQuery = q.Encode()
	var body interface{}
	r, err := http.NewPostRequestWithApiKey(u.String(), body, "Auth-Token", config.ApiKey)
	if err != nil {
		return fmt.Errorf("creating request for devices: %v", err)
	}
	resp, statusCode, err := http.DoWithStatusCode(r, time.Duration(*config.RequestTimeout)*time.Second, true)
	if err != nil {
		return fmt.Errorf("querying API for devices: %v", err)
	}
	if statusCode != nethttp.StatusOK {
		return fmt.Errorf("querying API for devices: got status %v and response %s", statusCode, resp)
	}
	log.Debug("broker", "posted action %v to device %v", action, deviceID)
	return nil
}

func (s *Switch) AdheresToFilter(filter [][]apiserver.FilterRule) (bool, error) {
	f := apiFilterToCommonFilter(filter)
	fp, err := utils.StructToMap(s)
	if err != nil {
		return false, fmt.Errorf("converting strict to map: %v", err)
	}
	adheres, err := common.Filter(f, fp)
	if err != nil {
		return false, err
	}
	return adheres, nil
}

func apiFilterToCommonFilter(input [][]apiserver.FilterRule) [][]common.FilterRule {
	result := make([][]common.FilterRule, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = make([]common.FilterRule, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			result[i][j] = common.FilterRule{
				Parameter: input[i][j].Parameter,
				Regex:     input[i][j].Regex,
			}
		}
	}
	return result
}
