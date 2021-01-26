package client

import (
	"flag"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var (
	KubeQPS         = float32(5.000000)
	KubeBurst       = 10
	KubeContentType = "application/vnd.kubernetes.protobuf"
	kubeconfig      *string
)

// InitConnect 连接k8s api
func InitConnect(env string) (*rest.Config, error) {

	if env == "out-of-cluster" {
		if kubeconfig != nil {
			config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
			if err != nil {
				panic(err.Error())
			}
			config.QPS = KubeQPS
			config.Burst = KubeBurst
			config.ContentType = KubeContentType
			return config, err
		}
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		config.QPS = KubeQPS
		config.Burst = KubeBurst
		config.ContentType = KubeContentType
		return config, err

	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		config.QPS = KubeQPS
		config.Burst = KubeBurst
		config.ContentType = KubeContentType
		return config, err
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
