# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

from ... import tables, version


class ClusterRoleBinding(pulumi.CustomResource):
    """
    ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole
    in the global namespace, and adds who information via Subject.
    """
    def __init__(self, __name__, __opts__=None, metadata=None, role_ref=None, subjects=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'rbac.authorization.k8s.io/v1'
        __props__['kind'] = 'ClusterRoleBinding'
        if role_ref is None:
            raise TypeError('Missing required property role_ref')
        __props__['roleRef'] = role_ref
        __props__['metadata'] = metadata
        __props__['subjects'] = subjects

        if __opts__ is None:
            __opts__ = pulumi.ResourceOptions()
        if __opts__.version is None:
            __opts__.version = version.get_version()

        super(ClusterRoleBinding, self).__init__(
            "kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBinding",
            __name__,
            __props__,
            __opts__)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
