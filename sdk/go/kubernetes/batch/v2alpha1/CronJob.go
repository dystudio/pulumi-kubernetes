// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// CronJob represents the configuration of a single cron job.
type CronJob struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Specification of the desired behavior of a cron job, including the schedule. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecOutput `pulumi:"spec"`

	// Current status of a cron job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status CronJobStatusOutput `pulumi:"status"`

}

// CronJobArgs is the set of arguments needed to create a CronJob resource.
type CronJobArgs struct {
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

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Specification of the desired behavior of a cron job, including the schedule. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecInput `pulumi:"spec"`

}

// NewCronJob creates a CronJob resource with the given unique name, arguments, and options.
func NewCronJob(ctx *pulumi.Context, name string, args *CronJobArgs, opts ...pulumi.ResourceOption) (*CronJob, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("batch/v2alpha1")
		args.Kind = pulumi.String("CronJob")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Spec; i != nil {
			inputs["spec"] = i.ToCronJobSpecOutput()
		}
	}
	var resource CronJob
	err := ctx.RegisterResource("kubernetes:batch/v2alpha1:CronJob", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJob gets an existing CronJob resource's state with the given name and ID.
func GetCronJob(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*CronJob, error) {
	var resource CronJob
	err := ctx.ReadResource("kubernetes:batch/v2alpha1:CronJob", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// CronJob represents the configuration of a single cron job.
type CronJobProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Specification of the desired behavior of a cron job, including the schedule. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *CronJobSpec `pulumi:"spec"`

	// Current status of a cron job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status CronJobStatus `pulumi:"status"`

}

var _CronJobPropertyType = reflect.TypeOf((*CronJobProperty)(nil)).Elem()

// CronJobPropertyInput represents an input type that resolves to a CronJobProperty.
type CronJobPropertyInput interface {
	ElementType() reflect.Type

	ToCronJobPropertyOutput() CronJobPropertyOutput
	ToCronJobPropertyOutputWithContext(ctx context.Context) CronJobPropertyOutput
}

// CronJobPropertyArgs is a CronJobPropertyInput whose fields are all Input types.
type CronJobPropertyArgs struct {
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

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Specification of the desired behavior of a cron job, including the schedule. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecInput `pulumi:"spec"`

}

func (a CronJobPropertyArgs) ElementType() reflect.Type {
	return _CronJobPropertyType
}

func (a CronJobPropertyArgs) ToCronJobPropertyOutput() CronJobPropertyOutput {
	return pulumi.ToOutput(a).(CronJobPropertyOutput)
}

func (a CronJobPropertyArgs) ToCronJobPropertyOutputWithContext(ctx context.Context) CronJobPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CronJobPropertyOutput)
}

// CronJobPropertyOutput is an output type that resolves to a Input.
type CronJobPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CronJobPropertyOutput{}) }

func (CronJobPropertyOutput) ElementType() reflect.Type {
	return _CronJobPropertyType
}

func (o CronJobPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v CronJobProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o CronJobPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v CronJobProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o CronJobPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v CronJobProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o CronJobPropertyOutput) Spec() CronJobSpecOutput {
	return o.Apply(func(v CronJobProperty) *CronJobSpec {
		return v.Spec
	}).(CronJobSpecOutput)
}

func (o CronJobPropertyOutput) Status() CronJobStatusOutput {
	return o.Apply(func(v CronJobProperty) CronJobStatus {
		return v.Status
	}).(CronJobStatusOutput)
}

var _CronJobPropertyArrayType = reflect.TypeOf((*[]CronJobProperty)(nil)).Elem()

type CronJobPropertyArrayInput interface {
	ElementType() reflect.Type

	ToCronJobPropertyArrayOutput() CronJobPropertyArrayOutput
	ToCronJobPropertyArrayOutputWithContext(ctx context.Context) CronJobPropertyArrayOutput
}

type CronJobPropertyArray []CronJobPropertyInput

func (a CronJobPropertyArray) ElementType() reflect.Type {
	return _CronJobPropertyArrayType
}

func (a CronJobPropertyArray) ToCronJobPropertyArrayOutput() CronJobPropertyArrayOutput {
	return pulumi.ToOutput(a).(CronJobPropertyArrayOutput)
}

func (a CronJobPropertyArray) ToCronJobPropertyArrayOutputWithContext(ctx context.Context) CronJobPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CronJobPropertyArrayOutput)
}

type CronJobPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CronJobPropertyArrayOutput{}) }

func (CronJobPropertyArrayOutput) ElementType() reflect.Type {
	return _CronJobPropertyArrayType
}

func (o CronJobPropertyArrayOutput) Index(i pulumi.IntInput) CronJobPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) CronJobProperty {
		return vs[0].([]CronJobProperty)[vs[1].(int)]
	}).(CronJobPropertyOutput)
}