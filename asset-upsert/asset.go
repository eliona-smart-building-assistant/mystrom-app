package assetupsert

import (
	"fmt"

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v2"
	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-utils/common"
)

type Root interface {
	LocationalNode
	FunctionalNode
}

type LocationalNode interface {
	Asset
	GetLocationalChildren() []LocationalNode
}

type FunctionalNode interface {
	Asset
	GetFunctionalChildren() []FunctionalNode
}

type Asset interface {
	GetName() string
	GetDescription() string
	GetAssetType() string
	GetGAI() string

	GetAssetID(projectID string) (*int32, error)
	SetAssetID(assetID int32, projectID string) error
}

func TraverseLocationalTree(node LocationalNode, projectId string, locationalParentAssetId, functionalParentAssetId *int32) error {
	currentAssetId, err := createAsset(node, projectId, locationalParentAssetId, functionalParentAssetId)
	if err != nil {
		return err
	}

	for _, child := range node.GetLocationalChildren() {
		if child == nil {
			continue
		}
		err := TraverseLocationalTree(child, projectId, currentAssetId, functionalParentAssetId)
		if err != nil {
			return err
		}
	}
	return nil
}

func TraverseFunctionalTree(node FunctionalNode, projectId string, locationalParentAssetId, functionalParentAssetId *int32) error {
	currentAssetId, err := createAsset(node, projectId, locationalParentAssetId, functionalParentAssetId)
	if err != nil {
		return err
	}

	for _, child := range node.GetFunctionalChildren() {
		if child == nil {
			continue
		}
		err := TraverseFunctionalTree(child, projectId, locationalParentAssetId, currentAssetId)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateRoot(ast Asset, projectId string) (*int32, error) {
	return createAsset(ast, projectId, nil, nil)
}

func createAsset(ast Asset, projectId string, locationalParentAssetId *int32, functionalParentAssetId *int32) (*int32, error) {
	originalAssetID, err := ast.GetAssetID(projectId)
	if err != nil {
		return nil, fmt.Errorf("getting asset id: %v", err)
	}
	a := api.Asset{
		Id:                      *api.NewNullableInt32(originalAssetID),
		ProjectId:               projectId,
		GlobalAssetIdentifier:   ast.GetGAI(),
		Name:                    *api.NewNullableString(common.Ptr(ast.GetName())),
		AssetType:               ast.GetAssetType(),
		Description:             *api.NewNullableString(common.Ptr(ast.GetDescription())),
		ParentFunctionalAssetId: *api.NewNullableInt32(functionalParentAssetId),
		ParentLocationalAssetId: *api.NewNullableInt32(locationalParentAssetId),
	}
	assetID, err := asset.UpsertAsset(a)
	if err != nil {
		return nil, fmt.Errorf("upserting asset %+v into Eliona: %v", a, err)
	}
	if assetID == nil {
		return nil, fmt.Errorf("cannot create asset %s", ast.GetName())
	}

	if err := ast.SetAssetID(*assetID, projectId); err != nil {
		return nil, fmt.Errorf("setting asset id: %v", err)
	}
	return assetID, nil
}
