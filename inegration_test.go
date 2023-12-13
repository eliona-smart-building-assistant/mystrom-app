package main

import (
	"github.com/eliona-smart-building-assistant/app-integration-tests/app"
	"github.com/eliona-smart-building-assistant/app-integration-tests/assert"
	"github.com/eliona-smart-building-assistant/app-integration-tests/test"
	"testing"
)

func TestApp(t *testing.T) {
	app.StartApp()
	test.AppWorks(t)
	t.Run("TestAssetTypes", assetTypes)
	t.Run("TestWidgetTypes", widgetTypes)
	t.Run("TestSchema", schema)
	app.StopApp()
}

func schema(t *testing.T) {
	t.Parallel()

	assert.SchemaExists(t, "mystrom", []string{"configuration", "asset"})
}

func assetTypes(t *testing.T) {
	t.Parallel()

	assert.AssetTypeExists(t, "mystrom_room", []string{})
	assert.AssetTypeExists(t, "mystrom_root", []string{})
	assert.AssetTypeExists(t, "mystrom_switch", []string{})
}

func widgetTypes(t *testing.T) {
	t.Parallel()

	assert.WidgetTypeExists(t, "myStrom Switch")
}
