module github.com/prodanlabs/kubeedge-wechat-examples

go 1.15

require (
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/golang/protobuf v1.4.3
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/iris/v12 v12.2.0-alpha2
	github.com/kubeedge/kubeedge v1.5.0
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/yaotian/gowechat v0.0.0-20180129120839-42de4f86cb0f
	github.com/yudai/pp v2.0.1+incompatible // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v0.19.3
)

replace (
	github.com/kubeedge/beehive => /opt/workspaces/beehive
	github.com/kubeedge/viaduct => /opt/workspaces/viaduct
	k8s.io/api => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/api
	k8s.io/apiextensions-apiserver => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/apimachinery
	k8s.io/apiserver => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/apiserver
	k8s.io/cli-runtime => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/cli-runtime
	k8s.io/client-go => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/client-go
	k8s.io/cloud-provider => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/cluster-bootstrap
	k8s.io/code-generator => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/code-generator
	k8s.io/component-base => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/component-base
	k8s.io/cri-api => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/cri-api
	k8s.io/csi-translation-lib => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/csi-translation-lib
	k8s.io/klog/v2 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/klog/v2
	k8s.io/kube-aggregator v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kube-controller-manager
	k8s.io/kube-openapi v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kube-openapi
	k8s.io/kube-proxy v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kube-proxy
	k8s.io/kube-scheduler => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kube-scheduler
	k8s.io/kubectl v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kubectl
	k8s.io/kubelet => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/kubelet
	k8s.io/legacy-cloud-providers v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/legacy-cloud-providers
	k8s.io/metrics v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/metrics
	k8s.io/sample-apiserver v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/sample-apiserver
	k8s.io/utils v0.0.0 => /opt/workspaces/kubernetes-release-1.19/vendor/k8s.io/utils
)
