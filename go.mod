module knative.dev/eventing-ceph

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/google/go-cmp v0.5.5
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/eventing v0.21.1-0.20210330142021-dca8b0acf039
	knative.dev/hack v0.0.0-20210325223819-b6ab329907d3
	knative.dev/pkg v0.0.0-20210331065221-952fdd90dbb0
)

replace (
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.6
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
	gomodules.xyz/jsonpatch/v2 => github.com/gomodules/jsonpatch/v2 v2.1.0
	vbom.ml/util => github.com/fvbommel/util v0.0.2
)
