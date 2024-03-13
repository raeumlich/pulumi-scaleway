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
    'GetDeviceResult',
    'AwaitableGetDeviceResult',
    'get_device',
    'get_device_output',
]

@pulumi.output_type
class GetDeviceResult:
    """
    A collection of values returned by getDevice.
    """
    def __init__(__self__, allow_insecure=None, allow_multiple_connections=None, certificates=None, created_at=None, description=None, device_id=None, hub_id=None, id=None, is_connected=None, last_activity_at=None, message_filters=None, name=None, region=None, status=None, updated_at=None):
        if allow_insecure and not isinstance(allow_insecure, bool):
            raise TypeError("Expected argument 'allow_insecure' to be a bool")
        pulumi.set(__self__, "allow_insecure", allow_insecure)
        if allow_multiple_connections and not isinstance(allow_multiple_connections, bool):
            raise TypeError("Expected argument 'allow_multiple_connections' to be a bool")
        pulumi.set(__self__, "allow_multiple_connections", allow_multiple_connections)
        if certificates and not isinstance(certificates, list):
            raise TypeError("Expected argument 'certificates' to be a list")
        pulumi.set(__self__, "certificates", certificates)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if hub_id and not isinstance(hub_id, str):
            raise TypeError("Expected argument 'hub_id' to be a str")
        pulumi.set(__self__, "hub_id", hub_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_connected and not isinstance(is_connected, bool):
            raise TypeError("Expected argument 'is_connected' to be a bool")
        pulumi.set(__self__, "is_connected", is_connected)
        if last_activity_at and not isinstance(last_activity_at, str):
            raise TypeError("Expected argument 'last_activity_at' to be a str")
        pulumi.set(__self__, "last_activity_at", last_activity_at)
        if message_filters and not isinstance(message_filters, list):
            raise TypeError("Expected argument 'message_filters' to be a list")
        pulumi.set(__self__, "message_filters", message_filters)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="allowInsecure")
    def allow_insecure(self) -> bool:
        return pulumi.get(self, "allow_insecure")

    @property
    @pulumi.getter(name="allowMultipleConnections")
    def allow_multiple_connections(self) -> bool:
        return pulumi.get(self, "allow_multiple_connections")

    @property
    @pulumi.getter
    def certificates(self) -> Sequence['outputs.GetDeviceCertificateResult']:
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[str]:
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> str:
        return pulumi.get(self, "hub_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isConnected")
    def is_connected(self) -> bool:
        return pulumi.get(self, "is_connected")

    @property
    @pulumi.getter(name="lastActivityAt")
    def last_activity_at(self) -> str:
        return pulumi.get(self, "last_activity_at")

    @property
    @pulumi.getter(name="messageFilters")
    def message_filters(self) -> Sequence['outputs.GetDeviceMessageFilterResult']:
        return pulumi.get(self, "message_filters")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetDeviceResult(GetDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceResult(
            allow_insecure=self.allow_insecure,
            allow_multiple_connections=self.allow_multiple_connections,
            certificates=self.certificates,
            created_at=self.created_at,
            description=self.description,
            device_id=self.device_id,
            hub_id=self.hub_id,
            id=self.id,
            is_connected=self.is_connected,
            last_activity_at=self.last_activity_at,
            message_filters=self.message_filters,
            name=self.name,
            region=self.region,
            status=self.status,
            updated_at=self.updated_at)


def get_device(device_id: Optional[str] = None,
               hub_id: Optional[str] = None,
               name: Optional[str] = None,
               region: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceResult:
    """
    Gets information about an IOT Device.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_device = scaleway.iot.get_device(device_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str device_id: The device ID.
           Only one of the `name` and `device_id` should be specified.
    :param str hub_id: The hub ID.
    :param str name: The name of the Hub.
           Only one of the `name` and `device_id` should be specified.
    :param str region: `region`) The region in which the hub exists.
    """
    __args__ = dict()
    __args__['deviceId'] = device_id
    __args__['hubId'] = hub_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:iot/getDevice:getDevice', __args__, opts=opts, typ=GetDeviceResult).value

    return AwaitableGetDeviceResult(
        allow_insecure=pulumi.get(__ret__, 'allow_insecure'),
        allow_multiple_connections=pulumi.get(__ret__, 'allow_multiple_connections'),
        certificates=pulumi.get(__ret__, 'certificates'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        device_id=pulumi.get(__ret__, 'device_id'),
        hub_id=pulumi.get(__ret__, 'hub_id'),
        id=pulumi.get(__ret__, 'id'),
        is_connected=pulumi.get(__ret__, 'is_connected'),
        last_activity_at=pulumi.get(__ret__, 'last_activity_at'),
        message_filters=pulumi.get(__ret__, 'message_filters'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_device)
def get_device_output(device_id: Optional[pulumi.Input[Optional[str]]] = None,
                      hub_id: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      region: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeviceResult]:
    """
    Gets information about an IOT Device.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_device = scaleway.iot.get_device(device_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str device_id: The device ID.
           Only one of the `name` and `device_id` should be specified.
    :param str hub_id: The hub ID.
    :param str name: The name of the Hub.
           Only one of the `name` and `device_id` should be specified.
    :param str region: `region`) The region in which the hub exists.
    """
    ...
