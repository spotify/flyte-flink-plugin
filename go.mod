module github.com/spotify/flyte-flink-plugin

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.3
	github.com/flyteorg/flyteidl v0.21.25
	github.com/flyteorg/flyteplugins v0.10.3
	github.com/flyteorg/flytestdlib v0.4.12
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-version v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/spotify/flink-on-k8s-operator v0.3.8
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.9.6
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.8.3
)
