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

	"github.com/eliona-smart-building-assistant/go-eliona/utils"
	"github.com/eliona-smart-building-assistant/go-utils/common"
)

type Switch struct {
	ID   string `eliona:"id" subtype:"info"`
	Name string `eliona:"name,filterable" subtype:"info"`

	Power      float32 `eliona:"power" subtype:"input"`
	Temp       float32 `eliona:"temperature" subtype:"input"`
	RelayState int     `eliona:"relay_state" subtype:"input"`

	Relay int `eliona:"relay" subtype:"output"`
}

func (s *Switch) AssetType() string {
	return "mystrom_switch"
}

func (s *Switch) Id() string {
	return s.AssetType() + "_" + s.ID
}

func GetDevices(config apiserver.Configuration) ([]Switch, error) {
	return nil, nil
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
