# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'PolicyRuleArgs',
]

@pulumi.input_type
class PolicyRuleArgs:
    def __init__(__self__, *,
                 permission_set_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permission_set_names: Names of permission sets bound to the rule.
               
               **_TIP:_**  You can use the Scaleway CLI to list the permissions details. e.g:
               
               ```shell
               $ scw iam permission-set list
               ```
        :param pulumi.Input[str] organization_id: ID of organization scoped to the rule, this can be used to create a rule for all projects in an organization.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_ids: List of project IDs scoped to the rule.
               
               > **Important** One of `organization_id` or `project_ids`  must be set per rule.
        """
        pulumi.set(__self__, "permission_set_names", permission_set_names)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_ids is not None:
            pulumi.set(__self__, "project_ids", project_ids)

    @property
    @pulumi.getter(name="permissionSetNames")
    def permission_set_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Names of permission sets bound to the rule.

        **_TIP:_**  You can use the Scaleway CLI to list the permissions details. e.g:

        ```shell
        $ scw iam permission-set list
        ```
        """
        return pulumi.get(self, "permission_set_names")

    @permission_set_names.setter
    def permission_set_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permission_set_names", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of organization scoped to the rule, this can be used to create a rule for all projects in an organization.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectIds")
    def project_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of project IDs scoped to the rule.

        > **Important** One of `organization_id` or `project_ids`  must be set per rule.
        """
        return pulumi.get(self, "project_ids")

    @project_ids.setter
    def project_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_ids", value)


