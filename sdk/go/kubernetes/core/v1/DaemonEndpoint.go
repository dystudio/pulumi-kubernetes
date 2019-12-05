// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DaemonEndpoint contains information about a single Daemon endpoint.
type DaemonEndpoint struct {
	// Port number of the given endpoint.
	Port int `pulumi:"Port"`

}

var _DaemonEndpointType = reflect.TypeOf((*DaemonEndpoint)(nil)).Elem()

// DaemonEndpointInput represents an input type that resolves to a DaemonEndpoint.
type DaemonEndpointInput interface {
	ElementType() reflect.Type

	ToDaemonEndpointOutput() DaemonEndpointOutput
	ToDaemonEndpointOutputWithContext(ctx context.Context) DaemonEndpointOutput
}

// DaemonEndpointArgs is a DaemonEndpointInput whose fields are all Input types.
type DaemonEndpointArgs struct {
	// Port number of the given endpoint.
	Port pulumi.IntInput `pulumi:"Port"`

}

func (a DaemonEndpointArgs) ElementType() reflect.Type {
	return _DaemonEndpointType
}

func (a DaemonEndpointArgs) ToDaemonEndpointOutput() DaemonEndpointOutput {
	return pulumi.ToOutput(a).(DaemonEndpointOutput)
}

func (a DaemonEndpointArgs) ToDaemonEndpointOutputWithContext(ctx context.Context) DaemonEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DaemonEndpointOutput)
}

// DaemonEndpointOutput is an output type that resolves to a Input.
type DaemonEndpointOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DaemonEndpointOutput{}) }

func (DaemonEndpointOutput) ElementType() reflect.Type {
	return _DaemonEndpointType
}

func (o DaemonEndpointOutput) Port() pulumi.IntOutput {
	return o.Apply(func(v DaemonEndpoint) int {
		return v.Port
	}).(pulumi.IntOutput)
}
