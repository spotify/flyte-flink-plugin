module github.com/spotify/flyte-flink-plugin

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.3
	github.com/flyteorg/flyteidl v0.21.23
	github.com/flyteorg/flyteplugins v0.9.4
	github.com/flyteorg/flytestdlib v0.4.10
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-version v1.4.0
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/onsi/gomega v1.16.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0 // indirect
	go.uber.org/zap v1.18.1 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver v0.21.3 // indirect
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	k8s.io/utils v0.0.0-20210722164352-7f3ee0f31471 // indirect
	sigs.k8s.io/controller-runtime v0.9.6
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.8.3
)
