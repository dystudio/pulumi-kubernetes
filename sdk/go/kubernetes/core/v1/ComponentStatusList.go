// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Status of all the conditions for the component as a list of ComponentStatus objects.
type ComponentStatusList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// List of ComponentStatus objects.
	Items ComponentStatusPropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard list metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// ComponentStatusListArgs is the set of arguments needed to create a ComponentStatusList resource.
type ComponentStatusListArgs struct {
	// List of ComponentStatus objects.
	Items ComponentStatusPropertyArrayInput `pulumi:"items"`

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

	// Standard list metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewComponentStatusList creates a ComponentStatusList resource with the given unique name, arguments, and options.
func NewComponentStatusList(ctx *pulumi.Context, name string, args *ComponentStatusListArgs, opts ...pulumi.ResourceOption) (*ComponentStatusList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("ComponentStatusList")
		inputs["items"] = args.Items.ToComponentStatusPropertyArrayOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToListMetaOutput()
		}
	}
	var resource ComponentStatusList
	err := ctx.RegisterResource("kubernetes:core/v1:ComponentStatusList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponentStatusList gets an existing ComponentStatusList resource's state with the given name and ID.
func GetComponentStatusList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*ComponentStatusList, error) {
	var resource ComponentStatusList
	err := ctx.ReadResource("kubernetes:core/v1:ComponentStatusList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
