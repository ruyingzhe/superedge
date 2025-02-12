/*
Copyright 2020 The SuperEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constant

const (
	ACTION_CHANGE = "change"
	ACTION_REVERT = "revert"
)

const (
	NODE_ROLE_NODE   = "node"
	NODE_ROLE_MASTER = "master"
)

const (
	KUBELET_STATUS_CMD  = "systemctl status kubelet.service"
	KUBELET_RESTART_CMD = "systemctl restart kubelet.service"
)

const (
	EDGE_NODE_KEY                 = "superedge.io/edge"
	EDGE_CHANGE_NODE_KEY          = "superedge.io/change"
	KUBERNETES_DEFAULT_ROLE_LABEL = "node-role.kubernetes.io/master"
)

const (
	LiteAPIServerAddr    = "https://127.0.0.1:51003"
	LiteAPIServerPodName = "lite-apiserver"

	KubeletHealthzURl = "http://127.0.0.1:10248/healthz"
)

const (
	EDGE_CERT_CM = "edge-cert"
)

const (
	MasterHostsFilePath         = "/etc/hosts"
	KubeletStartEnvFile         = "/etc/sysconfig/kubelet"
	KubeadmKubeletEdgeCert      = "/etc/kubernetes/edge/"
	KubeadmKubeletConfig        = "/etc/kubernetes/kubelet.conf"
	EdgeadmKubeletConfig        = "/etc/kubernetes/edge/kubelet.config"
	LiteAPIServerTemplatePath   = "/etc/kubernetes/edge/lite-apiserver.yaml"
	LiteAPIServerYamlPath       = "/etc/kubernetes/manifests/lite-api-server.yaml"
	KubeAPIServerBackUPYamlPath = "/etc/kubernetes/kube-apiserver-backup.yaml"
	KubeAPIServerSrcYamlPath    = "/etc/kubernetes/manifests/kube-apiserver.yaml"
)

const (
	CHANGE_KUBELET_KUBECONFIG_ARGS = `KUBELET_KUBECONFIG_ARGS="--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/edge/kubelet.config"`
)
