package azure

import (
	"fmt"
	"strings"

	"github.com/c2fo/cloud-finder/pkg/cloudfinder/provider"
)

type diskInfo struct {
	Caching      string `json:"caching"`
	CreateOption string `json:"createOption"`
	DiskSizeGB   string `json:"diskSizeGB"`
	Image        struct {
		URI string `json:"uri"`
	} `json:"Image"`
	ManagedDisk struct {
		ID                 string `json:"id"`
		StorageAccountType string `json:"storageAccountType"`
	} `json:"managedDisk"`
	Name string `json:"name"`
	VHD  struct {
		URI string `json:"uri"`
	} `json:"vhd"`
	WriteAcceleratorEnabled string `json:"writeAcceleratorEnabled"`
}

// Result is the response from the Azure Metadata service and implements
// provider.Result
type Result struct {
	Compute struct {
		AZEnvironment              string `json:"azEnvironment"`
		IsHostCompatibilityLayerVM string `json:"isHostCompatibilityLayerVm"`
		Location                   string `json:"location"`
		Name                       string `json:"name"`
		Offer                      string `json:"offer"`
		OSProfile                  struct {
			AdminUsername                 string `json:"adminUsername"`
			ComputerName                  string `json:"computerName"`
			DisablePasswordAuthentication string `json:"disablePasswordAuthentication"`
		}
		OSType           string `json:"osType"`
		PlacementGroupID string `json:"placementGroupId"`
		Plan             struct {
			Name      string `json:"name"`
			Product   string `json:"product"`
			Publisher string `json:"publisher"`
		}
		PlatformFaultDomain  string `json:"platformFaultDomain"`
		PlatformUpdateDomain string `json:"platformUpdateDomain"`
		PublicKeys           []struct {
			KeyData string `json:"keyData"`
			Path    string `json:"path"`
		} `json:"publicKeys"`
		Publisher         string `json:"publisher"`
		ResourceGroupName string `json:"resourceGroupName"`
		ResourceID        string `json:"resourceId"`
		SecurityProfile   struct {
			SecureBootEnabled string `json:"secureBootEnabled"`
			VirtualTPMEnabled string `json:"virtualTpmEnabled"`
		}
		SKU            string `json:"sku"`
		StorageProfile struct {
			DataDisks []struct {
				diskInfo
				LUN string `json:"lun"`
			} `json:"dataDisks"`
			ImageReference struct {
				ID        string `json:"string"`
				Offer     string `json:"offer"`
				Publisher string `json:"publisher"`
				SKU       string `json:"sku"`
				Version   string `json:"version"`
			} `json:"imageReference"`
			OSDisk struct {
				diskInfo
				DiffDiskSettings struct {
					Option string `json:"local"`
				} `json:"diffDiskSettings"`
				EncryptionSettings struct {
					Enabled string `json:"enabled"`
				} `json:"encryptionSettings"`
				OSType string `json:"osType"`
			}
		} `json:"storageProfile"`
		SubscriptionID string `json:"subscriptionId"`
		Tags           string `json:"tags"`
		Version        string `json:"version"`
		VMID           string `json:"vmId"`
		VMScaleSetName string `json:"vmScaleSetName"`
		VMSize         string `json:"vmSize"`
		Zone           string `json:"zone"`
	} `json:"compute"`
	Network struct {
		Interface []struct {
			IPV4 struct {
				IPAddress []struct {
					PrivateIPAddress string `json:"privateIpAddress"`
					PublicIPAddress  string `json:"publicIpAddress"`
				} `json:"ipAddress"`
				Subnet []struct {
					Address string `json:"address"`
					Prefix  string `json:"prefix"`
				} `json:"subnet"`
			} `json:"ipv4"`
			IPV6 struct {
			} `json:"ipv6"`
			MAC string `json:"macAddress"`
		} `json:"interface"`
	} `json:"network"`
}

// Provider returns the Provider that made the Result.
func (r Result) Provider() provider.Provider {
	return &Provider{}
}

func (r Result) properties() map[string]string {
	props := make(map[string]string)
	props["CF_CLOUD"] = strings.ToUpper(r.Provider().Name())
	if len(r.Network.Interface) > 0 {
		ifc := r.Network.Interface[0]
		if len(ifc.IPV4.IPAddress) > 0 {
			props["AZURE_PRIVATE_IPV4"] = ifc.IPV4.IPAddress[0].PrivateIPAddress
		}
		props["AZURE_MAC"] = ifc.MAC
	}
	props["AZURE_LOCATION"] = r.Compute.Location
	props["AZURE_ZONE"] = r.Compute.Zone
	props["AZURE_TAGS"] = r.Compute.Tags
	props["AZURE_VM_SIZE"] = r.Compute.VMSize
	props["AZURE_VM_ID"] = r.Compute.VMID
	return props
}

// ToEval returns a string which should be able to be eval'd in a shell
func (r Result) ToEval() string {
	exports := make([]string, 0)
	for k, v := range r.properties() {
		exports = append(exports, fmt.Sprintf("export %s=%s", k, v))
	}
	return strings.Join(exports, "\n")
}

func (r Result) String() string {
	items := make([]string, 0)
	for k, v := range r.properties() {
		items = append(items, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(items, "\n")
}
