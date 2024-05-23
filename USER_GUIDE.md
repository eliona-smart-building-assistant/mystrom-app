# myStrom app

### Eliona App for myStrom smart switches integration

> The Smart WiFi Switch that switches connected devices on and off, automatically saves energy and protects your home. Designed in Switzerland.

This application provides direct access to myStrom smart switches via Eliona. It enables users to check the current states of switches, view statistical data, manage the smart switches, and crucially, link myStrom devices with systems from various other brands.

## Installation

The myStrom App is installed via the App Store in Eliona.

## Assets

The myStrom App automatically creates all the necessary asset types and assets.

### Structure assets

The `myStrom Root` and `myStrom Room` asset types are created just to create a structure in Eliona.

### Devices

- *Switch*: A smart WiFi switch.

| Attribute     | Description   | Subtype |
|---------------|---------------|---------|
| `Power` | Power  | input   |
| `Temp` | Temp  | input   |
| `RelayState` | Relay State  | input   |
| `Relay`      | Relay        | output  |

## Configuration

The myStrom App is configured by defining authentication credentials. Each configuration requires the following data:

| Attribute        | Description                                               |
|------------------|-----------------------------------------------------------|
| `apiKey`       | API key provided by myStrom                        |
| `enable`         | Flag to enable or disable fetching from this API          |
| `refreshInterval`| Interval in seconds for device discovery. This is an expensive operation, should be no lower than 3600 s |
| `dataPollInterval` | Frequency of polling for data updates in seconds.|
| `requestTimeout` | API query timeout in seconds                              |
| `assetFilter`    | Filter for asset creation, more details can be found in app's README |
| `projectIDs`     | List of Eliona project ids for which this device should collect data. For each project id, all assets are automatically created in Eliona. |

The configuration is done via a corresponding JSON structure. As an example, the following JSON structure can be used to define an endpoint for app permissions:

```
{
  "apiKey": "api.key",
  "enable": true,
  "refreshInterval": 3600,
  "dataPollInterval": 60,
  "requestTimeout": 120,
  "assetFilter": [],
  "projectIDs": [
    "10"
  ]
}
```

Configurations can be created using this structure in Eliona under `Apps > myStrom > Settings`. To do this, select the /configs endpoint with the POST method.

After completing configuration, the app starts Continuous Asset Creation. When all discovered devices are created, user is notified about that in Eliona's notification system.
