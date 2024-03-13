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
    'GetVPCResult',
    'AwaitableGetVPCResult',
    'get_vpc',
    'get_vpc_output',
]

@pulumi.output_type
class GetVPCResult:
    """
    A collection of values returned by getVPC.
    """
    def __init__(__self__, created_at=None, id=None, is_default=None, name=None, organization_id=None, project_id=None, region=None, tags=None, updated_at=None, vpc_id=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[bool]:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetVPCResult(GetVPCResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPCResult(
            created_at=self.created_at,
            id=self.id,
            is_default=self.is_default,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            region=self.region,
            tags=self.tags,
            updated_at=self.updated_at,
            vpc_id=self.vpc_id)


def get_vpc(is_default: Optional[bool] = None,
            name: Optional[str] = None,
            organization_id: Optional[str] = None,
            project_id: Optional[str] = None,
            region: Optional[str] = None,
            vpc_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPCResult:
    """
    Gets information about a Scaleway Virtual Private Cloud.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.vpc.get_vpc(name="foobar")
    by_id = scaleway.vpc.get_vpc(vpc_id="11111111-1111-1111-1111-111111111111")
    default = scaleway.vpc.get_vpc(is_default=True)
    ```
    <!--End PulumiCodeChooser -->


    :param bool is_default: To get default VPC's information.
    :param str name: Name of the VPC. One of `name` and `vpc_id` should be specified.
    :param str organization_id: The ID of the organization the VPC is associated with.
    :param str project_id: `project_id`) The ID of the project the VPC is associated with.
    :param str vpc_id: ID of the VPC. One of `name` and `vpc_id` should be specified.
    """
    __args__ = dict()
    __args__['isDefault'] = is_default
    __args__['name'] = name
    __args__['organizationId'] = organization_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:vpc/getVPC:getVPC', __args__, opts=opts, typ=GetVPCResult).value

    return AwaitableGetVPCResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_vpc)
def get_vpc_output(is_default: Optional[pulumi.Input[Optional[bool]]] = None,
                   name: Optional[pulumi.Input[Optional[str]]] = None,
                   organization_id: Optional[pulumi.Input[Optional[str]]] = None,
                   project_id: Optional[pulumi.Input[Optional[str]]] = None,
                   region: Optional[pulumi.Input[Optional[str]]] = None,
                   vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVPCResult]:
    """
    Gets information about a Scaleway Virtual Private Cloud.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.vpc.get_vpc(name="foobar")
    by_id = scaleway.vpc.get_vpc(vpc_id="11111111-1111-1111-1111-111111111111")
    default = scaleway.vpc.get_vpc(is_default=True)
    ```
    <!--End PulumiCodeChooser -->


    :param bool is_default: To get default VPC's information.
    :param str name: Name of the VPC. One of `name` and `vpc_id` should be specified.
    :param str organization_id: The ID of the organization the VPC is associated with.
    :param str project_id: `project_id`) The ID of the project the VPC is associated with.
    :param str vpc_id: ID of the VPC. One of `name` and `vpc_id` should be specified.
    """
    ...
