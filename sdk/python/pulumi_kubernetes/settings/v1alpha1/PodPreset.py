# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
import warnings

from ... import tables, version


class PodPreset(pulumi.CustomResource):
    """
    PodPreset is a policy resource that defines additional runtime requirements for a Pod.
    """
    def __init__(self, resource_name, opts=None, metadata=None, spec=None, __name__=None, __opts__=None):
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

        __props__['apiVersion'] = 'settings.k8s.io/v1alpha1'
        __props__['kind'] = 'PodPreset'
        __props__['metadata'] = metadata
        __props__['spec'] = spec

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = version.get_version()

        super(PodPreset, self).__init__(
            "kubernetes:settings.k8s.io/v1alpha1:PodPreset",
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
