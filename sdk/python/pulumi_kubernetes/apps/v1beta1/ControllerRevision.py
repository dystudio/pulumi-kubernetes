# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
import warnings

from ... import tables, version


class ControllerRevision(pulumi.CustomResource):
    """
    DEPRECATED - This group version of ControllerRevision is deprecated by
    apps/v1beta2/ControllerRevision. See the release notes for more information. ControllerRevision
    implements an immutable snapshot of state data. Clients are responsible for serializing and
    deserializing the objects that contain their internal state. Once a ControllerRevision has been
    successfully created, it can not be updated. The API Server will fail validation of all requests
    that attempt to mutate the Data field. ControllerRevisions may, however, be deleted. Note that,
    due to its use by both the DaemonSet and StatefulSet controllers for update and rollback, this
    object is beta. However, it may be subject to name and representation changes in future
    releases, and clients should not depend on its stability. It is primarily for internal use by
    controllers.
    """
    def __init__(self, resource_name, opts=None, data=None, metadata=None, revision=None, __name__=None, __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'apps/v1beta1'
        __props__['kind'] = 'ControllerRevision'
        if revision is None:
            raise TypeError('Missing required property revision')
        __props__['revision'] = revision
        __props__['data'] = data
        __props__['metadata'] = metadata

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = version.get_version()

        super(ControllerRevision, self).__init__(
            "kubernetes:apps/v1beta1:ControllerRevision",
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
