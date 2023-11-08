package eliona

import (
	"context"
	"fmt"
	"mystrom/apiserver"
	"mystrom/broker"
	"mystrom/conf"

	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func UpsertSwitchData(config apiserver.Configuration, assets []broker.Switch) error {
	for _, projectId := range *config.ProjectIDs {
		for _, a := range assets {
			log.Debug("Eliona", "upserting data for asset: config %d and asset '%v'", config.Id, a.Id())
			assetId, err := conf.GetAssetId(context.Background(), config, projectId, a.Id())
			if err != nil {
				return err
			}
			if assetId == nil {
				return fmt.Errorf("unable to find asset ID")
			}

			data := asset.Data{
				AssetId: *assetId,
				Data:    a,
			}
			if asset.UpsertAssetDataIfAssetExists(data); err != nil {
				return fmt.Errorf("upserting data: %v", err)
			}
		}
	}
	return nil
}
