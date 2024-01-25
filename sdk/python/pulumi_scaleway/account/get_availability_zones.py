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
    'GetAvailabilityZonesResult',
    'AwaitableGetAvailabilityZonesResult',
    'get_availability_zones',
    'get_availability_zones_output',
]

@pulumi.output_type
class GetAvailabilityZonesResult:
    """
    A collection of values returned by getAvailabilityZones.
    """
    def __init__(__self__, id=None, region=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        """
        List of availability zones by regions
        """
        return pulumi.get(self, "zones")


class AwaitableGetAvailabilityZonesResult(GetAvailabilityZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailabilityZonesResult(
            id=self.id,
            region=self.region,
            zones=self.zones)


def get_availability_zones(region: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAvailabilityZonesResult:
    """
    Use this data source to get the available zones information based on its Region.

    For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
    you can choose the location that better fits your need (country, latency, …).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.account.get_availability_zones(region="nl-ams")
    ```


    :param str region: Region is represented as a Geographical area such as France. Defaults: `fr-par`.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:account/getAvailabilityZones:getAvailabilityZones', __args__, opts=opts, typ=GetAvailabilityZonesResult).value

    return AwaitableGetAvailabilityZonesResult(
        id=pulumi.get(__ret__, 'id'),
        region=pulumi.get(__ret__, 'region'),
        zones=pulumi.get(__ret__, 'zones'))


@_utilities.lift_output_func(get_availability_zones)
def get_availability_zones_output(region: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAvailabilityZonesResult]:
    """
    Use this data source to get the available zones information based on its Region.

    For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
    you can choose the location that better fits your need (country, latency, …).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.account.get_availability_zones(region="nl-ams")
    ```


    :param str region: Region is represented as a Geographical area such as France. Defaults: `fr-par`.
    """
    ...
