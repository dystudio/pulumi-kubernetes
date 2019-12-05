// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// APIGroupList is a list of APIGroup, to allow clients to discover the API at /apis.
type APIGroupList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// groups is a list of APIGroup.
	Groups []APIGroup `pulumi:"groups"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

}

var _APIGroupListType = reflect.TypeOf((*APIGroupList)(nil)).Elem()

// APIGroupListInput represents an input type that resolves to a APIGroupList.
type APIGroupListInput interface {
	ElementType() reflect.Type

	ToAPIGroupListOutput() APIGroupListOutput
	ToAPIGroupListOutputWithContext(ctx context.Context) APIGroupListOutput
}

// APIGroupListArgs is a APIGroupListInput whose fields are all Input types.
type APIGroupListArgs struct {
	// groups is a list of APIGroup.
	Groups APIGroupArrayInput `pulumi:"groups"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

}

func (a APIGroupListArgs) ElementType() reflect.Type {
	return _APIGroupListType
}

func (a APIGroupListArgs) ToAPIGroupListOutput() APIGroupListOutput {
	return pulumi.ToOutput(a).(APIGroupListOutput)
}

func (a APIGroupListArgs) ToAPIGroupListOutputWithContext(ctx context.Context) APIGroupListOutput {
	return pulumi.ToOutputWithContext(ctx, a).(APIGroupListOutput)
}

// APIGroupListOutput is an output type that resolves to a Input.
type APIGroupListOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(APIGroupListOutput{}) }

func (APIGroupListOutput) ElementType() reflect.Type {
	return _APIGroupListType
}

func (o APIGroupListOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v APIGroupList) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o APIGroupListOutput) Groups() APIGroupArrayOutput {
	return o.Apply(func(v APIGroupList) []APIGroup {
		return v.Groups
	}).(APIGroupArrayOutput)
}

func (o APIGroupListOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v APIGroupList) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}
