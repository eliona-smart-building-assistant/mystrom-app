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
	"fmt"
	"mystrom/apiserver"
	assetupsert "mystrom/asset-upsert"

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
	"github.com/eliona-smart-building-assistant/go-eliona/client"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func CreateAssets(config apiserver.Configuration, root assetupsert.Root) error {
	for _, projectId := range *config.ProjectIDs {
		rootAssetID, err := assetupsert.CreateRoot(root, projectId)
		if err != nil {
			return fmt.Errorf("upserting root asset: %v", err)
		}
		for _, fc := range root.GetFunctionalChildren() {
			if fc == nil {
				continue
			}
			if err := assetupsert.TraverseFunctionalTree(fc, projectId, rootAssetID, rootAssetID); err != nil {
				return fmt.Errorf("functional tree traversal: %v", err)
			}
		}

		for _, lc := range root.GetLocationalChildren() {
			if lc == nil {
				continue
			}
			if err := assetupsert.TraverseLocationalTree(lc, projectId, rootAssetID, rootAssetID); err != nil {
				return fmt.Errorf("locational tree traversal: %v", err)
			}
		}
	}
	return nil
}

// TODO: Notify users. Currently not implemented.
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
