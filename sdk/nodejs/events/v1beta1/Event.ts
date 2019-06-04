// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import { getVersion } from "../../version";

    /**
     * Event is a report of an event somewhere in the cluster. It generally denotes some state
     * change in the system.
     */
    export class Event extends pulumi.CustomResource {
      /**
       * What action was taken/failed regarding to the regarding object.
       */
      public readonly action: pulumi.Output<string>;

      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"events.k8s.io/v1beta1">;

      /**
       * Deprecated field assuring backward compatibility with core.v1 Event type
       */
      public readonly deprecatedCount: pulumi.Output<number>;

      /**
       * Deprecated field assuring backward compatibility with core.v1 Event type
       */
      public readonly deprecatedFirstTimestamp: pulumi.Output<string>;

      /**
       * Deprecated field assuring backward compatibility with core.v1 Event type
       */
      public readonly deprecatedLastTimestamp: pulumi.Output<string>;

      /**
       * Deprecated field assuring backward compatibility with core.v1 Event type
       */
      public readonly deprecatedSource: pulumi.Output<outputApi.core.v1.EventSource>;

      /**
       * Required. Time when this Event was first observed.
       */
      public readonly eventTime: pulumi.Output<string>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"Event">;

      
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Optional. A human-readable description of the status of this operation. Maximal length of
       * the note is 1kB, but libraries should be prepared to handle values up to 64kB.
       */
      public readonly note: pulumi.Output<string>;

      /**
       * Why the action was taken.
       */
      public readonly reason: pulumi.Output<string>;

      /**
       * The object this Event is about. In most cases it's an Object reporting controller
       * implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted
       * because it acts on some changes in a ReplicaSet object.
       */
      public readonly regarding: pulumi.Output<outputApi.core.v1.ObjectReference>;

      /**
       * Optional secondary object for more complex actions. E.g. when regarding object triggers a
       * creation or deletion of related object.
       */
      public readonly related: pulumi.Output<outputApi.core.v1.ObjectReference>;

      /**
       * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
       */
      public readonly reportingController: pulumi.Output<string>;

      /**
       * ID of the controller instance, e.g. `kubelet-xyzf`.
       */
      public readonly reportingInstance: pulumi.Output<string>;

      /**
       * Data about the Event series this event represents or nil if it's a singleton Event.
       */
      public readonly series: pulumi.Output<outputApi.events.v1beta1.EventSeries>;

      /**
       * Type of this event (Normal, Warning), new types could be added in the future.
       */
      public readonly type: pulumi.Output<string>;

      /**
       * Get the state of an existing `Event` resource, as identified by `id`.
       * Typically this ID  is of the form <namespace>/<name>; if <namespace> is omitted, then (per
       * Kubernetes convention) the ID becomes default/<name>.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form
       *  <namespace>/<name> or <name>.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Event {
          return new Event(name, undefined, { ...opts, id: id });
      }

      /**
       * Create a events.v1beta1.Event resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.events.v1beta1.Event, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["action"] = args && args.action || undefined;
          inputs["apiVersion"] = "events.k8s.io/v1beta1";
          inputs["deprecatedCount"] = args && args.deprecatedCount || undefined;
          inputs["deprecatedFirstTimestamp"] = args && args.deprecatedFirstTimestamp || undefined;
          inputs["deprecatedLastTimestamp"] = args && args.deprecatedLastTimestamp || undefined;
          inputs["deprecatedSource"] = args && args.deprecatedSource || undefined;
          inputs["eventTime"] = args && args.eventTime || undefined;
          inputs["kind"] = "Event";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["note"] = args && args.note || undefined;
          inputs["reason"] = args && args.reason || undefined;
          inputs["regarding"] = args && args.regarding || undefined;
          inputs["related"] = args && args.related || undefined;
          inputs["reportingController"] = args && args.reportingController || undefined;
          inputs["reportingInstance"] = args && args.reportingInstance || undefined;
          inputs["series"] = args && args.series || undefined;
          inputs["type"] = args && args.type || undefined;
        
          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super("kubernetes:events.k8s.io/v1beta1:Event", name, inputs, opts);
      }
    }
