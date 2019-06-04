// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import { getVersion } from "../../version";

    /**
     * CSIDriver captures information about a Container Storage Interface (CSI) volume driver
     * deployed on the cluster. CSI drivers do not need to create the CSIDriver object directly.
     * Instead they may use the cluster-driver-registrar sidecar container. When deployed with a CSI
     * driver it automatically creates a CSIDriver object representing the driver. Kubernetes attach
     * detach controller uses this object to determine whether attach is required. Kubelet uses this
     * object to determine whether pod information needs to be passed on mount. CSIDriver objects
     * are non-namespaced.
     */
    export class CSIDriver extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"storage.k8s.io/v1beta1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"CSIDriver">;

      /**
       * Standard object metadata. metadata.Name indicates the name of the CSI driver that this
       * object refers to; it MUST be the same name returned by the CSI GetPluginName() call for
       * that driver. The driver name must be 63 characters or less, beginning and ending with an
       * alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between.
       * More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Specification of the CSI Driver.
       */
      public readonly spec: pulumi.Output<outputApi.storage.v1beta1.CSIDriverSpec>;

      /**
       * Get the state of an existing `CSIDriver` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CSIDriver {
          return new CSIDriver(name, undefined, { ...opts, id: id });
      }

      /**
       * Create a storage.v1beta1.CSIDriver resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.storage.v1beta1.CSIDriver, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "storage.k8s.io/v1beta1";
          inputs["kind"] = "CSIDriver";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
        
          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super("kubernetes:storage.k8s.io/v1beta1:CSIDriver", name, inputs, opts);
      }
    }
