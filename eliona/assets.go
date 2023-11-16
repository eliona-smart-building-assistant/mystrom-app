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

package eliona

import (
	"context"
	"fmt"
	"mystrom/apiserver"
	"mystrom/broker"
	"mystrom/conf"

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-eliona/client"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

type Asset interface {
	AssetType() string
	Id() string
}

func createRoomAssetIfNecessary(config apiserver.Configuration, projectId string, room broker.Room) (int32, error) {
	rootAssetID, err := upsertRootAsset(config, projectId)
	if err != nil {
		return 0, fmt.Errorf("upserting root asset: %v", err)
	}

	assetType := "mystrom_room"
	_, roomId, err := upsertAsset(assetData{
		config:                  config,
		projectId:               projectId,
		parentFunctionalAssetId: &rootAssetID,
		parentLocationalAssetId: &rootAssetID,
		identifier:              fmt.Sprintf("%s_%s", assetType, room.ID),
		assetType:               assetType,
		name:                    room.Name,
		description:             fmt.Sprintf("%s (%v)", room.Name, room.ID),
	})
	if err != nil {
		return 0, fmt.Errorf("upserting room %s: %v", room.ID, err)
	}
	return roomId, nil
}

func CreateAssetsIfNecessary(config apiserver.Configuration, devices []broker.Switch) error {
	for _, projectId := range conf.ProjIds(config) {
		assetsCreated := 0
		rootAssetID, err := upsertRootAsset(config, projectId)
		if err != nil {
			return fmt.Errorf("upserting root asset: %v", err)
		}
		for _, device := range devices {
			locParentId := rootAssetID
			if device.Room.ID != "" {
				locParentId, err = createRoomAssetIfNecessary(config, projectId, device.Room)
				if err != nil {
					return fmt.Errorf("upserting room asset: %v", err)
				}
			}
			assetType := "mystrom_switch"
			ad := assetData{
				config:                  config,
				projectId:               projectId,
				parentFunctionalAssetId: &rootAssetID,
				parentLocationalAssetId: &locParentId,
				identifier:              device.Id(),
				assetType:               assetType,
				name:                    fmt.Sprintf("%s | %s", device.Room.Name, device.Name),
				description:             fmt.Sprintf("%s (%v)", device.Name, device.Id()),
			}

			created, _, err := upsertAsset(ad)
			if err != nil {
				return fmt.Errorf("upserting device %s: %v", device.Id(), err)
			}
			if created {
				assetsCreated++
			}
		}
		if assetsCreated > 0 {
			if err := notifyUsers(projectId, assetsCreated); err != nil {
				return fmt.Errorf("notifying users about CAC: %v", err)
			}
		}
	}
	return nil
}

func upsertRootAsset(config apiserver.Configuration, projectId string) (int32, error) {
	_, rootAssetID, err := upsertAsset(assetData{
		config:                  config,
		projectId:               projectId,
		parentLocationalAssetId: nil,
		identifier:              "mystrom_root",
		assetType:               "mystrom_root",
		name:                    "myStrom",
		description:             "Root asset for myStrom devices",
	})
	return rootAssetID, err
}

type assetData struct {
	config                  apiserver.Configuration
	projectId               string
	parentFunctionalAssetId *int32
	parentLocationalAssetId *int32
	identifier              string
	assetType               string
	name                    string
	description             string
}

func upsertAsset(d assetData) (created bool, assetID int32, err error) {
	// Get known asset id from configuration
	currentAssetID, err := conf.GetAssetId(context.Background(), d.config, d.projectId, d.identifier)
	if err != nil {
		return false, 0, fmt.Errorf("finding asset ID: %v", err)
	}
	if currentAssetID != nil {
		return false, *currentAssetID, nil
	}

	a := api.Asset{
		ProjectId:               d.projectId,
		GlobalAssetIdentifier:   d.identifier,
		Name:                    *api.NewNullableString(common.Ptr(d.name)),
		AssetType:               d.assetType,
		Description:             *api.NewNullableString(common.Ptr(d.description)),
		ParentFunctionalAssetId: *api.NewNullableInt32(d.parentFunctionalAssetId),
		ParentLocationalAssetId: *api.NewNullableInt32(d.parentLocationalAssetId),
		IsTracker:               *api.NewNullableBool(common.Ptr(false)),
	}
	newID, err := asset.UpsertAsset(a)
	if err != nil {
		return false, 0, fmt.Errorf("upserting asset %+v into Eliona: %v", a, err)
	}
	if newID == nil {
		return false, 0, fmt.Errorf("cannot create asset %s", d.name)
	}

	if err := conf.InsertAsset(context.Background(), d.config, d.projectId, d.identifier, *newID); err != nil {
		return false, 0, fmt.Errorf("inserting asset to config db: %v", err)
	}

	log.Debug("eliona", "Created new asset for project %s and device %s.", d.projectId, d.identifier)

	return true, *newID, nil
}

func notifyUsers(projectId string, assetsCreated int) error {
	users, _, err := client.NewClient().UsersAPI.GetUsers(client.AuthenticationContext()).Execute()
	if err != nil {
		return fmt.Errorf("fetching all users: %v", err)
	}
	for _, user := range users {
		n := api.Notification{
			User:      user.Email,
			ProjectId: *api.NewNullableString(&projectId),
			Message: *api.NewNullableTranslation(&api.Translation{
				De: api.PtrString(fmt.Sprintf("Die kontinuierliche Asset-Erstellung für myStrom hat %v neue Assets hinzugefügt. Diese sind nun im Asset-Management verfügbar.", assetsCreated)),
				En: api.PtrString(fmt.Sprintf("The Continuous Asset Creation for myStrom added %v new assets. They are now available in Asset Management.", assetsCreated)),
			}),
		}
		receipt, _, err := client.NewClient().CommunicationAPI.
			PostNotification(client.AuthenticationContext()).
			Notification(n).
			Execute()
		log.Debug("eliona", "posted notification about CAC: %v", receipt)
		if err != nil {
			return fmt.Errorf("posting notification: %v", err)
		}
	}
	return nil
}
