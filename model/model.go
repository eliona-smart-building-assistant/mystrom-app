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

package model

import (
	"context"
	"fmt"
	"mystrom/apiserver"
	assetupsert "mystrom/asset-upsert"
	"mystrom/conf"

	"github.com/eliona-smart-building-assistant/go-eliona/utils"
	"github.com/eliona-smart-building-assistant/go-utils/common"
)

type Switch struct {
	ID   string `eliona:"id"`
	Name string `eliona:"name,filterable"`

	Power float32 `eliona:"power" subtype:"input"`
	Temp  float32 `eliona:"temperature" subtype:"input"`

	Relay int `eliona:"relay" subtype:"output"`

	Config *apiserver.Configuration
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

func (s *Switch) GetName() string {
	return s.Name
}

func (s *Switch) GetDescription() string {
	return ""
}

func (s *Switch) GetAssetType() string {
	return "mystrom_switch"
}

func (s *Switch) GetGAI() string {
	return s.GetAssetType() + "_" + s.ID
}

func (s *Switch) GetAssetID(projectID string) (*int32, error) {
	return conf.GetAssetId(context.Background(), *s.Config, projectID, s.GetGAI())
}

func (s *Switch) SetAssetID(assetID int32, projectID string) error {
	if err := conf.InsertAsset(context.Background(), *s.Config, projectID, s.GetGAI(), assetID, s.ID); err != nil {
		return fmt.Errorf("inserting asset to Config db: %v", err)
	}
	return nil
}

func (s *Switch) GetLocationalChildren() []assetupsert.LocationalNode {
	return []assetupsert.LocationalNode{}
}

func (s *Switch) GetFunctionalChildren() []assetupsert.FunctionalNode {
	return []assetupsert.FunctionalNode{}
}

type Room struct {
	ID   string
	Name string

	Config *apiserver.Configuration // get rid of this

	Switches []Switch
}

func (r *Room) GetName() string {
	return r.Name
}

func (r *Room) GetDescription() string {
	return ""
}

func (r *Room) GetAssetType() string {
	return "mystrom_room"
}

func (r *Room) GetGAI() string {
	return r.GetAssetType() + "_" + r.ID
}

func (r *Room) GetAssetID(projectID string) (*int32, error) {
	return conf.GetAssetId(context.Background(), *r.Config, projectID, r.GetGAI())
}

func (r *Room) SetAssetID(assetID int32, projectID string) error {
	if err := conf.InsertAsset(context.Background(), *r.Config, projectID, r.GetGAI(), assetID, r.ID); err != nil {
		return fmt.Errorf("inserting asset to Config db: %v", err)
	}
	return nil
}

func (r *Room) GetLocationalChildren() []assetupsert.LocationalNode {
	locationalChildren := make([]assetupsert.LocationalNode, len(r.Switches))
	for i, sw := range r.Switches {
		locationalChildren[i] = &sw
	}
	return locationalChildren
}

func (r *Room) GetFunctionalChildren() []assetupsert.FunctionalNode {
	functionalChildren := make([]assetupsert.FunctionalNode, len(r.Switches))
	for i, sw := range r.Switches {
		functionalChildren[i] = &sw
	}
	return functionalChildren
}

type Root struct {
	Rooms    map[string]Room
	Switches []Switch

	Config *apiserver.Configuration
}

func (r *Root) GetDevices() []Switch {
	return r.Switches
}

func (r *Root) GetName() string {
	return "myStrom"
}

func (r *Root) GetDescription() string {
	return "Root asset for myStrom devices"
}

func (r *Root) GetAssetType() string {
	return "mystrom_root"
}

func (r *Root) GetGAI() string {
	return r.GetAssetType()
}

func (r *Root) GetAssetID(projectID string) (*int32, error) {
	return conf.GetAssetId(context.Background(), *r.Config, projectID, r.GetGAI())
}

func (r *Root) SetAssetID(assetID int32, projectID string) error {
	if err := conf.InsertAsset(context.Background(), *r.Config, projectID, r.GetGAI(), assetID, ""); err != nil {
		return fmt.Errorf("inserting asset to Config db: %v", err)
	}
	return nil
}

func (r *Root) GetLocationalChildren() []assetupsert.LocationalNode {
	locationalChildren := make([]assetupsert.LocationalNode, 0)
	for _, room := range r.Rooms {
		locationalChildren = append(locationalChildren, &room)
	}
	return locationalChildren
}

func (r *Root) GetFunctionalChildren() []assetupsert.FunctionalNode {
	functionalChildren := make([]assetupsert.FunctionalNode, len(r.Switches))
	for i, sw := range r.Switches {
		functionalChildren[i] = &sw
	}
	return functionalChildren
}

//

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
