package utils

import (
	gocontext "context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha2"

	"github.com/prodanlabs/kubeedge-wechat-examples/client"
)

const (
	// MergePatchType is patch type
	MergePatchType = "application/merge-patch+json"
	// ResourceTypeDevices is plural of device resource in apiserver
	ResourceTypeDevices = "devices"
)

// DeviceStatus is used to patch device status
type DeviceStatus struct {
	Status v1alpha2.DeviceStatus `json:"status"`
}

// UpdateDeviceTwinWithDesiredTrack patches the desired state of
// the device twin with the track to play.
func UpdateDeviceTwinWithDesiredTrack(track string) bool {
	status := buildStatusWithDesiredTrack(track)
	deviceStatus := &DeviceStatus{Status: status}
	body, err := json.Marshal(deviceStatus)
	if err != nil {
		log.Printf("Failed to marshal device status %v", deviceStatus)
		return false
	}
	env := GetEnv("IN_CLUSTER", "out-of-cluster")
	kubeConfig, err := client.InitConnect(env)
	if err != nil {
		log.Println("Failed to connect kubernetes")
	}
	crdClient, err := client.NewCRDClient(kubeConfig)
	if err != nil {
		log.Println("crdClient initialization failed")
	}
	// The device id of the speaker
	deviceID := GetEnv("DEVICE_ID", "speaker-01")
	// The default namespace in which the speaker device instance resides
	namespace := GetEnv("DEVICE_NAMESPACE", "default")

	result := crdClient.Patch(MergePatchType).Namespace(namespace).Resource(ResourceTypeDevices).Name(deviceID).Body(body).Do(gocontext.TODO())
	if result.Error() != nil {
		log.Printf("Failed to patch device status %v of device %v in namespace %v \n error:%+v", deviceStatus, deviceID, namespace, result.Error())
		return false
	} else {
		log.Printf("Track [ %s ] will be played on speaker %s", track, deviceID)
	}
	return true
}

func buildStatusWithDesiredTrack(song string) v1alpha2.DeviceStatus {
	metadata := map[string]string{"timestamp": strconv.FormatInt(time.Now().Unix()/1e6, 10),
		"type": "string",
	}
	twins := []v1alpha2.Twin{{PropertyName: "track", Desired: v1alpha2.TwinProperty{Value: song, Metadata: metadata}, Reported: v1alpha2.TwinProperty{Value: "unknown", Metadata: metadata}}}
	devicestatus := v1alpha2.DeviceStatus{Twins: twins}
	return devicestatus
}
