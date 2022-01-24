module github.com/kcp-dev/kcp

go 1.16

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/evanphx/json-patch v5.6.0+incompatible
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.1.2
	github.com/googleapis/gnostic v0.5.5
	github.com/muesli/reflow v0.1.0
	github.com/onsi/gomega v1.10.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/wayneashleyberry/terminal-dimensions v1.0.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.etcd.io/etcd/server/v3 v3.5.0
	go.uber.org/multierr v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.0.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.30.0
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	k8s.io/kubectl v0.0.0
	k8s.io/kubernetes v0.0.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/api => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/api
	k8s.io/apiextensions-apiserver => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/apimachinery
	k8s.io/apiserver => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/apiserver
	k8s.io/cli-runtime => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/cli-runtime
	k8s.io/client-go => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/client-go
	k8s.io/cloud-provider => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/cluster-bootstrap
	k8s.io/code-generator => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/code-generator
	k8s.io/component-base => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/component-base
	k8s.io/component-helpers => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/component-helpers
	k8s.io/controller-manager => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/controller-manager
	k8s.io/cri-api => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/cri-api
	k8s.io/csi-translation-lib => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/csi-translation-lib
	k8s.io/kube-aggregator => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kube-controller-manager
	k8s.io/kube-proxy => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kube-proxy
	k8s.io/kube-scheduler => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kube-scheduler
	k8s.io/kubectl => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kubectl
	k8s.io/kubelet => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/kubelet
	k8s.io/kubernetes => /Users/surbania/projects/kcp/kubernetes
	k8s.io/legacy-cloud-providers => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/legacy-cloud-providers
	k8s.io/metrics => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/metrics
	k8s.io/mount-utils => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/mount-utils
	k8s.io/pod-security-admission => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/pod-security-admission
	k8s.io/sample-apiserver => /Users/surbania/projects/kcp/kubernetes/staging/src/k8s.io/sample-apiserver
)
