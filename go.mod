module knative.dev/eventing-ceph

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/go-cmp v0.5.4
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/procfs v0.0.11 // indirect
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
	knative.dev/eventing v0.20.1-0.20210128132430-1725902f7e39
	knative.dev/hack v0.0.0-20210120165453-8d623a0af457
	knative.dev/pkg v0.0.0-20210127163530-0d31134d5f4e
	sigs.k8s.io/structured-merge-diff/v3 v3.0.1-0.20200706213357-43c19bbb7fba // indirect
)

replace (
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.6
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	gomodules.xyz/jsonpatch/v2 => github.com/gomodules/jsonpatch/v2 v2.1.0
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/apiserver => k8s.io/apiserver v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/code-generator => k8s.io/code-generator v0.18.8
	vbom.ml/util => github.com/fvbommel/util v0.0.2
)
