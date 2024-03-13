# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetBackendsResult',
    'AwaitableGetBackendsResult',
    'get_backends',
    'get_backends_output',
]

@pulumi.output_type
class GetBackendsResult:
    """
    A collection of values returned by getBackends.
    """
    def __init__(__self__, backends=None, id=None, lb_id=None, name=None, organization_id=None, project_id=None, zone=None):
        if backends and not isinstance(backends, list):
            raise TypeError("Expected argument 'backends' to be a list")
        pulumi.set(__self__, "backends", backends)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lb_id and not isinstance(lb_id, str):
            raise TypeError("Expected argument 'lb_id' to be a str")
        pulumi.set(__self__, "lb_id", lb_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def backends(self) -> Sequence['outputs.GetBackendsBackendResult']:
        """
        List of found backends
        """
        return pulumi.get(self, "backends")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> str:
        return pulumi.get(self, "lb_id")

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
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetBackendsResult(GetBackendsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendsResult(
            backends=self.backends,
            id=self.id,
            lb_id=self.lb_id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            zone=self.zone)


def get_backends(lb_id: Optional[str] = None,
                 name: Optional[str] = None,
                 project_id: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackendsResult:
    """
    Gets information about multiple Load Balancer Backends.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_lbid = scaleway.loadbalancer.get_backends(lb_id=scaleway_lb["lb01"]["id"])
    by_lbid_and_name = scaleway.loadbalancer.get_backends(lb_id=scaleway_lb["lb01"]["id"],
        name="tf-backend-datasource")
    ```
    <!--End PulumiCodeChooser -->


    :param str lb_id: The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
    :param str name: The backend name used as filter. Backends with a name like it are listed.
    :param str zone: `zone`) The zone in which backends exist.
    """
    __args__ = dict()
    __args__['lbId'] = lb_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:loadbalancer/getBackends:getBackends', __args__, opts=opts, typ=GetBackendsResult).value

    return AwaitableGetBackendsResult(
        backends=pulumi.get(__ret__, 'backends'),
        id=pulumi.get(__ret__, 'id'),
        lb_id=pulumi.get(__ret__, 'lb_id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_backends)
def get_backends_output(lb_id: Optional[pulumi.Input[str]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackendsResult]:
    """
    Gets information about multiple Load Balancer Backends.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_lbid = scaleway.loadbalancer.get_backends(lb_id=scaleway_lb["lb01"]["id"])
    by_lbid_and_name = scaleway.loadbalancer.get_backends(lb_id=scaleway_lb["lb01"]["id"],
        name="tf-backend-datasource")
    ```
    <!--End PulumiCodeChooser -->


    :param str lb_id: The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
    :param str name: The backend name used as filter. Backends with a name like it are listed.
    :param str zone: `zone`) The zone in which backends exist.
    """
    ...
