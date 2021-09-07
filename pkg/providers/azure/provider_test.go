package azure

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
	"github.com/c2fo/cloud-finder/pkg/logging"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// From Azure Docs: https://docs.microsoft.com/en-us/azure/virtual-machines/windows/instance-metadata-service?tabs=linux
const mockResp = `{
    "compute": {
        "azEnvironment": "AZUREPUBLICCLOUD",
        "isHostCompatibilityLayerVm": "true",
        "licenseType":  "",
        "location": "westus",
        "name": "examplevmname",
        "offer": "UbuntuServer",
        "osProfile": {
            "adminUsername": "admin",
            "computerName": "examplevmname",
            "disablePasswordAuthentication": "true"
        },
        "osType": "Linux",
        "placementGroupId": "f67c14ab-e92c-408c-ae2d-da15866ec79a",
        "plan": {
            "name": "planName",
            "product": "planProduct",
            "publisher": "planPublisher"
        },
        "platformFaultDomain": "36",
        "platformUpdateDomain": "42",
        "publicKeys": [{
                "keyData": "ssh-rsa 0",
                "path": "/home/user/.ssh/authorized_keys0"
            },
            {
                "keyData": "ssh-rsa 1",
                "path": "/home/user/.ssh/authorized_keys1"
            }
        ],
        "publisher": "Canonical",
        "resourceGroupName": "macikgo-test-may-23",
        "resourceId": "/subscriptions/xxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx/resourceGroups/macikgo-test-may-23/providers/Microsoft.Compute/virtualMachines/examplevmname",
        "securityProfile": {
            "secureBootEnabled": "true",
            "virtualTpmEnabled": "false"
        },
        "sku": "18.04-LTS",
        "storageProfile": {
            "dataDisks": [{
                "caching": "None",
                "createOption": "Empty",
                "diskSizeGB": "1024",
                "image": {
                    "uri": ""
                },
                "lun": "0",
                "managedDisk": {
                    "id": "/subscriptions/xxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx/resourceGroups/macikgo-test-may-23/providers/Microsoft.Compute/disks/exampledatadiskname",
                    "storageAccountType": "Standard_LRS"
                },
                "name": "exampledatadiskname",
                "vhd": {
                    "uri": ""
                },
                "writeAcceleratorEnabled": "false"
            }],
            "imageReference": {
                "id": "",
                "offer": "UbuntuServer",
                "publisher": "Canonical",
                "sku": "16.04.0-LTS",
                "version": "latest"
            },
            "osDisk": {
                "caching": "ReadWrite",
                "createOption": "FromImage",
                "diskSizeGB": "30",
                "diffDiskSettings": {
                    "option": "Local"
                },
                "encryptionSettings": {
                    "enabled": "false"
                },
                "image": {
                    "uri": ""
                },
                "managedDisk": {
                    "id": "/subscriptions/xxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx/resourceGroups/macikgo-test-may-23/providers/Microsoft.Compute/disks/exampleosdiskname",
                    "storageAccountType": "Standard_LRS"
                },
                "name": "exampleosdiskname",
                "osType": "Linux",
                "vhd": {
                    "uri": ""
                },
                "writeAcceleratorEnabled": "false"
            }
        },
        "subscriptionId": "xxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
        "tags": "baz:bash;foo:bar",
        "version": "15.05.22",
        "vmId": "02aab8a4-74ef-476e-8182-f6d2ba4166a6",
        "vmScaleSetName": "crpteste9vflji9",
        "vmSize": "Standard_A3",
        "zone": ""
    },
    "network": {
        "interface": [{
            "ipv4": {
               "ipAddress": [{
                    "privateIpAddress": "10.144.133.132",
                    "publicIpAddress": ""
                }],
                "subnet": [{
                    "address": "10.144.133.128",
                    "prefix": "26"
                }]
            },
            "ipv6": {
                "ipAddress": [
                 ]
            },
            "macAddress": "0011AAFFBB22"
        }]
    }
}`

func TestProviderName(t *testing.T) {
	p := Provider{}
	assert.Equal(t, "azure", p.Name())
}

func TestAzureProviderImplementsProvider(t *testing.T) {
	assert.Implements(t, (*provider.Provider)(nil), new(Provider))
}

func registerHTTPMockResponse(method, relativePath, response string) {
	fullPath := strings.Join([]string{baseURL, relativePath}, "")
	httpmock.RegisterResponder(method, fullPath, httpmock.NewStringResponder(http.StatusOK, response))
}

func withTestRoutes(t *testing.T, resp string, f func(t *testing.T)) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	registerHTTPMockResponse("GET", "/metadata/instance?api-version=2020-09-01", resp)

	f(t)
}

func TestAzureProvider(t *testing.T) {
	logging.EnableDebug()
	azureProvider := &Provider{}
	providerOptions := &provider.Options{HTTPTimeout: 5 * time.Second}

	// GitHub Actions runs on Azure so need to activate the mock for a Azure error response
	httpmock.Activate()
	result := azureProvider.Check(&provider.Options{HTTPTimeout: 100 * time.Millisecond})
	httpmock.DeactivateAndReset()
	assert.Nil(t, result)

	withTestRoutes(t, mockResp, func(t *testing.T) {
		result := azureProvider.Check(providerOptions)
		assert.NotNil(t, result, "Result should not be nil.")

		azureresult, ok := result.(Result)
		assert.True(t, ok, "Result should be an Azure Result.")

		assert.Equal(t, azureresult.Compute.AZEnvironment, "AZUREPUBLICCLOUD")
		assert.Equal(t, azureresult.Compute.Location, "westus")
		assert.Equal(t, azureresult.Compute.OSType, "Linux")
		assert.Equal(t, azureresult.Network.Interface[0].MAC, "0011AAFFBB22")
	})

	withTestRoutes(t, `non json response`, func(t *testing.T) {
		result := azureProvider.Check(providerOptions)
		assert.Nil(t, result, "Result should not be nil.")
	})
}
