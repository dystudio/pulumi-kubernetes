module github.com/pulumi/pulumi-kubernetes

go 1.13

require (
	cloud.google.com/go v0.45.1 // indirect
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/aws/aws-sdk-go v1.25.3 // indirect
	github.com/cbroglie/mustache v1.0.1
	github.com/cheggaaa/pb v1.0.27 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/googleapis/gnostic v0.2.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jinzhu/copier v0.0.0-20180308034124-7e38e58719c3
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.1
	github.com/stretchr/testify v1.4.1-0.20191106224347-f1bd0923b832
	golang.org/x/net v0.0.0-20191009170851-d66e71096ffb // indirect
	google.golang.org/grpc v1.23.0
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	k8s.io/kubectl v0.17.0
	sigs.k8s.io/yaml v1.1.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.1+incompatible

replace github.com/pulumi/pulumi => ../pulumi

replace github.com/pulumi/pulumi-terraform-bridge => ../pulumi-terraform-bridge
