// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// ConfigMap holds configuration data for pods to consume.
type ConfigMap struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored
	// in BinaryData must not overlap with the ones in the Data field, this is enforced during
	// validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapOutput `pulumi:"binaryData"`

	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in
	// Data must not overlap with the keys in the BinaryData field, this is enforced during validation
	// process.
	Data pulumi.StringMapOutput `pulumi:"data"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

}

// ConfigMapArgs is the set of arguments needed to create a ConfigMap resource.
type ConfigMapArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored
	// in BinaryData must not overlap with the ones in the Data field, this is enforced during
	// validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapInput `pulumi:"binaryData"`

	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in
	// Data must not overlap with the keys in the BinaryData field, this is enforced during validation
	// process.
	Data pulumi.StringMapInput `pulumi:"data"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

// NewConfigMap creates a ConfigMap resource with the given unique name, arguments, and options.
func NewConfigMap(ctx *pulumi.Context, name string, args *ConfigMapArgs, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("ConfigMap")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.BinaryData; i != nil {
			inputs["binaryData"] = i.ToStringMapOutput()
		}
		if i := args.Data; i != nil {
			inputs["data"] = i.ToStringMapOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
	}
	var resource ConfigMap
	err := ctx.RegisterResource("kubernetes:core/v1:ConfigMap", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMap gets an existing ConfigMap resource's state with the given name and ID.
func GetConfigMap(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	var resource ConfigMap
	err := ctx.ReadResource("kubernetes:core/v1:ConfigMap", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// ConfigMap holds configuration data for pods to consume.
type ConfigMapProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored
	// in BinaryData must not overlap with the ones in the Data field, this is enforced during
	// validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData map[string]string `pulumi:"binaryData"`

	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in
	// Data must not overlap with the keys in the BinaryData field, this is enforced during validation
	// process.
	Data map[string]string `pulumi:"data"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

}

var _ConfigMapPropertyType = reflect.TypeOf((*ConfigMapProperty)(nil)).Elem()

// ConfigMapPropertyInput represents an input type that resolves to a ConfigMapProperty.
type ConfigMapPropertyInput interface {
	ElementType() reflect.Type

	ToConfigMapPropertyOutput() ConfigMapPropertyOutput
	ToConfigMapPropertyOutputWithContext(ctx context.Context) ConfigMapPropertyOutput
}

// ConfigMapPropertyArgs is a ConfigMapPropertyInput whose fields are all Input types.
type ConfigMapPropertyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored
	// in BinaryData must not overlap with the ones in the Data field, this is enforced during
	// validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapInput `pulumi:"binaryData"`

	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_'
	// or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in
	// Data must not overlap with the keys in the BinaryData field, this is enforced during validation
	// process.
	Data pulumi.StringMapInput `pulumi:"data"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

func (a ConfigMapPropertyArgs) ElementType() reflect.Type {
	return _ConfigMapPropertyType
}

func (a ConfigMapPropertyArgs) ToConfigMapPropertyOutput() ConfigMapPropertyOutput {
	return pulumi.ToOutput(a).(ConfigMapPropertyOutput)
}

func (a ConfigMapPropertyArgs) ToConfigMapPropertyOutputWithContext(ctx context.Context) ConfigMapPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ConfigMapPropertyOutput)
}

// ConfigMapPropertyOutput is an output type that resolves to a Input.
type ConfigMapPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ConfigMapPropertyOutput{}) }

func (ConfigMapPropertyOutput) ElementType() reflect.Type {
	return _ConfigMapPropertyType
}

func (o ConfigMapPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v ConfigMapProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o ConfigMapPropertyOutput) BinaryData() pulumi.StringMapOutput {
	return o.Apply(func(v ConfigMapProperty) map[string]string {
		return v.BinaryData
	}).(pulumi.StringMapOutput)
}

func (o ConfigMapPropertyOutput) Data() pulumi.StringMapOutput {
	return o.Apply(func(v ConfigMapProperty) map[string]string {
		return v.Data
	}).(pulumi.StringMapOutput)
}

func (o ConfigMapPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v ConfigMapProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o ConfigMapPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v ConfigMapProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

var _ConfigMapPropertyArrayType = reflect.TypeOf((*[]ConfigMapProperty)(nil)).Elem()

type ConfigMapPropertyArrayInput interface {
	ElementType() reflect.Type

	ToConfigMapPropertyArrayOutput() ConfigMapPropertyArrayOutput
	ToConfigMapPropertyArrayOutputWithContext(ctx context.Context) ConfigMapPropertyArrayOutput
}

type ConfigMapPropertyArray []ConfigMapPropertyInput

func (a ConfigMapPropertyArray) ElementType() reflect.Type {
	return _ConfigMapPropertyArrayType
}

func (a ConfigMapPropertyArray) ToConfigMapPropertyArrayOutput() ConfigMapPropertyArrayOutput {
	return pulumi.ToOutput(a).(ConfigMapPropertyArrayOutput)
}

func (a ConfigMapPropertyArray) ToConfigMapPropertyArrayOutputWithContext(ctx context.Context) ConfigMapPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ConfigMapPropertyArrayOutput)
}

type ConfigMapPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ConfigMapPropertyArrayOutput{}) }

func (ConfigMapPropertyArrayOutput) ElementType() reflect.Type {
	return _ConfigMapPropertyArrayType
}

func (o ConfigMapPropertyArrayOutput) Index(i pulumi.IntInput) ConfigMapPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ConfigMapProperty {
		return vs[0].([]ConfigMapProperty)[vs[1].(int)]
	}).(ConfigMapPropertyOutput)
}