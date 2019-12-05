// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Event is a report of an event somewhere in the cluster.
type Event struct {
	pulumi.CustomResourceState

	// What action was taken/failed regarding to the Regarding object.
	Action pulumi.StringOutput `pulumi:"action"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// The number of times this event has occurred.
	Count pulumi.IntOutput `pulumi:"count"`

	// Time when this Event was first observed.
	EventTime pulumi.StringOutput `pulumi:"eventTime"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp pulumi.StringOutput `pulumi:"firstTimestamp"`

	// The object that this event is about.
	InvolvedObject ObjectReferenceOutput `pulumi:"involvedObject"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp pulumi.StringOutput `pulumi:"lastTimestamp"`

	// A human-readable description of the status of this operation.
	Message pulumi.StringOutput `pulumi:"message"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// This should be a short, machine understandable string that gives the reason for the transition
	// into the object's current status.
	Reason pulumi.StringOutput `pulumi:"reason"`

	// Optional secondary object for more complex actions.
	Related ObjectReferenceOutput `pulumi:"related"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent pulumi.StringOutput `pulumi:"reportingComponent"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringOutput `pulumi:"reportingInstance"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesOutput `pulumi:"series"`

	// The component reporting this event. Should be a short machine understandable string.
	Source EventSourceOutput `pulumi:"source"`

	// Type of this event (Normal, Warning), new types could be added in the future
	Type pulumi.StringOutput `pulumi:"type"`

}

// EventArgs is the set of arguments needed to create a Event resource.
type EventArgs struct {
	// The object that this event is about.
	InvolvedObject ObjectReferenceInput `pulumi:"involvedObject"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// What action was taken/failed regarding to the Regarding object.
	Action pulumi.StringInput `pulumi:"action"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// The number of times this event has occurred.
	Count pulumi.IntInput `pulumi:"count"`

	// Time when this Event was first observed.
	EventTime pulumi.StringInput `pulumi:"eventTime"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp pulumi.StringInput `pulumi:"firstTimestamp"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp pulumi.StringInput `pulumi:"lastTimestamp"`

	// A human-readable description of the status of this operation.
	Message pulumi.StringInput `pulumi:"message"`

	// This should be a short, machine understandable string that gives the reason for the transition
	// into the object's current status.
	Reason pulumi.StringInput `pulumi:"reason"`

	// Optional secondary object for more complex actions.
	Related ObjectReferenceInput `pulumi:"related"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent pulumi.StringInput `pulumi:"reportingComponent"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringInput `pulumi:"reportingInstance"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesInput `pulumi:"series"`

	// The component reporting this event. Should be a short machine understandable string.
	Source EventSourceInput `pulumi:"source"`

	// Type of this event (Normal, Warning), new types could be added in the future
	Type pulumi.StringInput `pulumi:"type"`

}

// NewEvent creates a Event resource with the given unique name, arguments, and options.
func NewEvent(ctx *pulumi.Context, name string, args *EventArgs, opts ...pulumi.ResourceOption) (*Event, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.InvolvedObject == nil {
		return nil, errors.New("missing required argument 'InvolvedObject'")
	}
	if args == nil || args.Metadata == nil {
		return nil, errors.New("missing required argument 'Metadata'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("Event")
		inputs["involvedObject"] = args.InvolvedObject.ToObjectReferenceOutput()
		inputs["metadata"] = args.Metadata.ToObjectMetaOutput()
		if i := args.Action; i != nil {
			inputs["action"] = i.ToStringOutput()
		}
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Count; i != nil {
			inputs["count"] = i.ToIntOutput()
		}
		if i := args.EventTime; i != nil {
			inputs["eventTime"] = i.ToStringOutput()
		}
		if i := args.FirstTimestamp; i != nil {
			inputs["firstTimestamp"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.LastTimestamp; i != nil {
			inputs["lastTimestamp"] = i.ToStringOutput()
		}
		if i := args.Message; i != nil {
			inputs["message"] = i.ToStringOutput()
		}
		if i := args.Reason; i != nil {
			inputs["reason"] = i.ToStringOutput()
		}
		if i := args.Related; i != nil {
			inputs["related"] = i.ToObjectReferenceOutput()
		}
		if i := args.ReportingComponent; i != nil {
			inputs["reportingComponent"] = i.ToStringOutput()
		}
		if i := args.ReportingInstance; i != nil {
			inputs["reportingInstance"] = i.ToStringOutput()
		}
		if i := args.Series; i != nil {
			inputs["series"] = i.ToEventSeriesOutput()
		}
		if i := args.Source; i != nil {
			inputs["source"] = i.ToEventSourceOutput()
		}
		if i := args.Type; i != nil {
			inputs["type"] = i.ToStringOutput()
		}
	}
	var resource Event
	err := ctx.RegisterResource("kubernetes:core/v1:Event", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEvent gets an existing Event resource's state with the given name and ID.
func GetEvent(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*Event, error) {
	var resource Event
	err := ctx.ReadResource("kubernetes:core/v1:Event", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Event is a report of an event somewhere in the cluster.
type EventProperty struct {
	// What action was taken/failed regarding to the Regarding object.
	Action *string `pulumi:"action"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// The number of times this event has occurred.
	Count *int `pulumi:"count"`

	// Time when this Event was first observed.
	EventTime *string `pulumi:"eventTime"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp *string `pulumi:"firstTimestamp"`

	// The object that this event is about.
	InvolvedObject ObjectReference `pulumi:"involvedObject"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp *string `pulumi:"lastTimestamp"`

	// A human-readable description of the status of this operation.
	Message *string `pulumi:"message"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMeta `pulumi:"metadata"`

	// This should be a short, machine understandable string that gives the reason for the transition
	// into the object's current status.
	Reason *string `pulumi:"reason"`

	// Optional secondary object for more complex actions.
	Related *ObjectReference `pulumi:"related"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent *string `pulumi:"reportingComponent"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance *string `pulumi:"reportingInstance"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `pulumi:"series"`

	// The component reporting this event. Should be a short machine understandable string.
	Source *EventSource `pulumi:"source"`

	// Type of this event (Normal, Warning), new types could be added in the future
	Type *string `pulumi:"type"`

}

var _EventPropertyType = reflect.TypeOf((*EventProperty)(nil)).Elem()

// EventPropertyInput represents an input type that resolves to a EventProperty.
type EventPropertyInput interface {
	ElementType() reflect.Type

	ToEventPropertyOutput() EventPropertyOutput
	ToEventPropertyOutputWithContext(ctx context.Context) EventPropertyOutput
}

// EventPropertyArgs is a EventPropertyInput whose fields are all Input types.
type EventPropertyArgs struct {
	// The object that this event is about.
	InvolvedObject ObjectReferenceInput `pulumi:"involvedObject"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// What action was taken/failed regarding to the Regarding object.
	Action pulumi.StringInput `pulumi:"action"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// The number of times this event has occurred.
	Count pulumi.IntInput `pulumi:"count"`

	// Time when this Event was first observed.
	EventTime pulumi.StringInput `pulumi:"eventTime"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp pulumi.StringInput `pulumi:"firstTimestamp"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp pulumi.StringInput `pulumi:"lastTimestamp"`

	// A human-readable description of the status of this operation.
	Message pulumi.StringInput `pulumi:"message"`

	// This should be a short, machine understandable string that gives the reason for the transition
	// into the object's current status.
	Reason pulumi.StringInput `pulumi:"reason"`

	// Optional secondary object for more complex actions.
	Related ObjectReferenceInput `pulumi:"related"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent pulumi.StringInput `pulumi:"reportingComponent"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance pulumi.StringInput `pulumi:"reportingInstance"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesInput `pulumi:"series"`

	// The component reporting this event. Should be a short machine understandable string.
	Source EventSourceInput `pulumi:"source"`

	// Type of this event (Normal, Warning), new types could be added in the future
	Type pulumi.StringInput `pulumi:"type"`

}

func (a EventPropertyArgs) ElementType() reflect.Type {
	return _EventPropertyType
}

func (a EventPropertyArgs) ToEventPropertyOutput() EventPropertyOutput {
	return pulumi.ToOutput(a).(EventPropertyOutput)
}

func (a EventPropertyArgs) ToEventPropertyOutputWithContext(ctx context.Context) EventPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventPropertyOutput)
}

// EventPropertyOutput is an output type that resolves to a Input.
type EventPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EventPropertyOutput{}) }

func (EventPropertyOutput) ElementType() reflect.Type {
	return _EventPropertyType
}

func (o EventPropertyOutput) Action() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.Action
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) Count() pulumi.IntOutput {
	return o.Apply(func(v EventProperty) *int {
		return v.Count
	}).(pulumi.IntOutput)
}

func (o EventPropertyOutput) EventTime() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.EventTime
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) FirstTimestamp() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.FirstTimestamp
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) InvolvedObject() ObjectReferenceOutput {
	return o.Apply(func(v EventProperty) ObjectReference {
		return v.InvolvedObject
	}).(ObjectReferenceOutput)
}

func (o EventPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) LastTimestamp() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.LastTimestamp
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v EventProperty) metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o EventPropertyOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) Related() ObjectReferenceOutput {
	return o.Apply(func(v EventProperty) *ObjectReference {
		return v.Related
	}).(ObjectReferenceOutput)
}

func (o EventPropertyOutput) ReportingComponent() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.ReportingComponent
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) ReportingInstance() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.ReportingInstance
	}).(pulumi.StringOutput)
}

func (o EventPropertyOutput) Series() EventSeriesOutput {
	return o.Apply(func(v EventProperty) *EventSeries {
		return v.Series
	}).(EventSeriesOutput)
}

func (o EventPropertyOutput) Source() EventSourceOutput {
	return o.Apply(func(v EventProperty) *EventSource {
		return v.Source
	}).(EventSourceOutput)
}

func (o EventPropertyOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v EventProperty) *string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _EventPropertyArrayType = reflect.TypeOf((*[]EventProperty)(nil)).Elem()

type EventPropertyArrayInput interface {
	ElementType() reflect.Type

	ToEventPropertyArrayOutput() EventPropertyArrayOutput
	ToEventPropertyArrayOutputWithContext(ctx context.Context) EventPropertyArrayOutput
}

type EventPropertyArray []EventPropertyInput

func (a EventPropertyArray) ElementType() reflect.Type {
	return _EventPropertyArrayType
}

func (a EventPropertyArray) ToEventPropertyArrayOutput() EventPropertyArrayOutput {
	return pulumi.ToOutput(a).(EventPropertyArrayOutput)
}

func (a EventPropertyArray) ToEventPropertyArrayOutputWithContext(ctx context.Context) EventPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventPropertyArrayOutput)
}

type EventPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EventPropertyArrayOutput{}) }

func (EventPropertyArrayOutput) ElementType() reflect.Type {
	return _EventPropertyArrayType
}

func (o EventPropertyArrayOutput) Index(i pulumi.IntInput) EventPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EventProperty {
		return vs[0].([]EventProperty)[vs[1].(int)]
	}).(EventPropertyOutput)
}