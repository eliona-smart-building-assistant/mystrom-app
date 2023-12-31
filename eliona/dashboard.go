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

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
	"github.com/eliona-smart-building-assistant/go-eliona/client"
	"github.com/eliona-smart-building-assistant/go-utils/common"
)

func GetDashboard(projectId string) (api.Dashboard, error) {
	dashboard := api.Dashboard{}
	dashboard.Name = "myStrom"
	dashboard.ProjectId = projectId
	dashboard.Widgets = []api.Widget{}

	switches, _, err := client.NewClient().AssetsAPI.
		GetAssets(client.AuthenticationContext()).
		AssetTypeName("mystrom_switch").
		ProjectId(projectId).
		Execute()
	if err != nil {
		return api.Dashboard{}, fmt.Errorf("fetching switches: %v", err)
	}
	widgetSequence := int32(0)
	for _, sw := range switches {
		switchesData := []api.WidgetData{
			{
				ElementSequence: nullableInt32(1),
				AssetId:         sw.Id,
				Data: map[string]interface{}{
					"attribute":   "relay",
					"description": sw.Name,
					"key":         "_CURRENT",
					"seq":         0,
					"subtype":     "output",
				},
			},
			{
				ElementSequence: nullableInt32(1),
				AssetId:         sw.Id,
				Data: map[string]interface{}{
					"attribute":   "relay",
					"description": sw.Name,
					"key":         "_SETPOINT",
					"seq":         0,
					"subtype":     "output",
				},
			},
			{
				ElementSequence: nullableInt32(2),
				AssetId:         sw.Id,
				Data: map[string]interface{}{
					"aggregatedDataField":  nil,
					"aggregatedDataRaster": nil,
					"aggregatedDataType":   "heap",
					"attribute":            "power",
					"description":          "Power",
					"key":                  "",
					"seq":                  0,
					"subtype":              "input",
				},
			},
			{
				ElementSequence: nullableInt32(3),
				AssetId:         sw.Id,
				Data: map[string]interface{}{
					"aggregatedDataField":  nil,
					"aggregatedDataRaster": nil,
					"aggregatedDataType":   "heap",
					"attribute":            "temperature",
					"description":          "Temperature",
					"key":                  "",
					"seq":                  0,
					"subtype":              "input",
				},
			},
		}
		dashboard.Widgets = append(dashboard.Widgets, api.Widget{
			WidgetTypeName: "myStrom Switch",
			AssetId:        sw.Id,
			Sequence:       nullableInt32(widgetSequence),
			Details: map[string]any{
				"size":     1,
				"timespan": 7,
			},
			Data: switchesData,
		})
		widgetSequence++
	}

	return dashboard, nil
}

func nullableInt32(val int32) api.NullableInt32 {
	return *api.NewNullableInt32(common.Ptr[int32](val))
}
