// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Sysctl defines a kernel parameter to be set
type Sysctl struct {
	// Name of a property to set
	Name string `pulumi:"name"`

	// Value of a property to set
	Value string `pulumi:"value"`

}

var _SysctlType = reflect.TypeOf((*Sysctl)(nil)).Elem()

// SysctlInput represents an input type that resolves to a Sysctl.
type SysctlInput interface {
	ElementType() reflect.Type

	ToSysctlOutput() SysctlOutput
	ToSysctlOutputWithContext(ctx context.Context) SysctlOutput
}

// SysctlArgs is a SysctlInput whose fields are all Input types.
type SysctlArgs struct {
	// Name of a property to set
	Name pulumi.StringInput `pulumi:"name"`

	// Value of a property to set
	Value pulumi.StringInput `pulumi:"value"`

}

func (a SysctlArgs) ElementType() reflect.Type {
	return _SysctlType
}

func (a SysctlArgs) ToSysctlOutput() SysctlOutput {
	return pulumi.ToOutput(a).(SysctlOutput)
}

func (a SysctlArgs) ToSysctlOutputWithContext(ctx context.Context) SysctlOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SysctlOutput)
}

// SysctlOutput is an output type that resolves to a Input.
type SysctlOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(SysctlOutput{}) }

func (SysctlOutput) ElementType() reflect.Type {
	return _SysctlType
}

func (o SysctlOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v Sysctl) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o SysctlOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v Sysctl) string {
		return v.Value
	}).(pulumi.StringOutput)
}

var _SysctlArrayType = reflect.TypeOf((*[]Sysctl)(nil)).Elem()

type SysctlArrayInput interface {
	ElementType() reflect.Type

	ToSysctlArrayOutput() SysctlArrayOutput
	ToSysctlArrayOutputWithContext(ctx context.Context) SysctlArrayOutput
}

type SysctlArray []SysctlInput

func (a SysctlArray) ElementType() reflect.Type {
	return _SysctlArrayType
}

func (a SysctlArray) ToSysctlArrayOutput() SysctlArrayOutput {
	return pulumi.ToOutput(a).(SysctlArrayOutput)
}

func (a SysctlArray) ToSysctlArrayOutputWithContext(ctx context.Context) SysctlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SysctlArrayOutput)
}

type SysctlArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(SysctlArrayOutput{}) }

func (SysctlArrayOutput) ElementType() reflect.Type {
	return _SysctlArrayType
}

func (o SysctlArrayOutput) Index(i pulumi.IntInput) SysctlOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) Sysctl {
		return vs[0].([]Sysctl)[vs[1].(int)]
	}).(SysctlOutput)
}