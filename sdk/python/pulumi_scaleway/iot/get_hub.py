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
    'GetHubResult',
    'AwaitableGetHubResult',
    'get_hub',
    'get_hub_output',
]

@pulumi.output_type
class GetHubResult:
    """
    A collection of values returned by getHub.
    """
    def __init__(__self__, connected_device_count=None, created_at=None, device_auto_provisioning=None, device_count=None, disable_events=None, enabled=None, endpoint=None, events_topic_prefix=None, hub_ca=None, hub_ca_challenge=None, hub_id=None, id=None, name=None, organization_id=None, product_plan=None, project_id=None, region=None, status=None, updated_at=None):
        if connected_device_count and not isinstance(connected_device_count, int):
            raise TypeError("Expected argument 'connected_device_count' to be a int")
        pulumi.set(__self__, "connected_device_count", connected_device_count)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if device_auto_provisioning and not isinstance(device_auto_provisioning, bool):
            raise TypeError("Expected argument 'device_auto_provisioning' to be a bool")
        pulumi.set(__self__, "device_auto_provisioning", device_auto_provisioning)
        if device_count and not isinstance(device_count, int):
            raise TypeError("Expected argument 'device_count' to be a int")
        pulumi.set(__self__, "device_count", device_count)
        if disable_events and not isinstance(disable_events, bool):
            raise TypeError("Expected argument 'disable_events' to be a bool")
        pulumi.set(__self__, "disable_events", disable_events)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if events_topic_prefix and not isinstance(events_topic_prefix, str):
            raise TypeError("Expected argument 'events_topic_prefix' to be a str")
        pulumi.set(__self__, "events_topic_prefix", events_topic_prefix)
        if hub_ca and not isinstance(hub_ca, str):
            raise TypeError("Expected argument 'hub_ca' to be a str")
        pulumi.set(__self__, "hub_ca", hub_ca)
        if hub_ca_challenge and not isinstance(hub_ca_challenge, str):
            raise TypeError("Expected argument 'hub_ca_challenge' to be a str")
        pulumi.set(__self__, "hub_ca_challenge", hub_ca_challenge)
        if hub_id and not isinstance(hub_id, str):
            raise TypeError("Expected argument 'hub_id' to be a str")
        pulumi.set(__self__, "hub_id", hub_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if product_plan and not isinstance(product_plan, str):
            raise TypeError("Expected argument 'product_plan' to be a str")
        pulumi.set(__self__, "product_plan", product_plan)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
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
    @pulumi.getter(name="connectedDeviceCount")
    def connected_device_count(self) -> int:
        return pulumi.get(self, "connected_device_count")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deviceAutoProvisioning")
    def device_auto_provisioning(self) -> bool:
        return pulumi.get(self, "device_auto_provisioning")

    @property
    @pulumi.getter(name="deviceCount")
    def device_count(self) -> int:
        return pulumi.get(self, "device_count")

    @property
    @pulumi.getter(name="disableEvents")
    def disable_events(self) -> bool:
        return pulumi.get(self, "disable_events")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="eventsTopicPrefix")
    def events_topic_prefix(self) -> str:
        return pulumi.get(self, "events_topic_prefix")

    @property
    @pulumi.getter(name="hubCa")
    def hub_ca(self) -> str:
        return pulumi.get(self, "hub_ca")

    @property
    @pulumi.getter(name="hubCaChallenge")
    def hub_ca_challenge(self) -> str:
        return pulumi.get(self, "hub_ca_challenge")

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> Optional[str]:
        return pulumi.get(self, "hub_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="productPlan")
    def product_plan(self) -> str:
        return pulumi.get(self, "product_plan")

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
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetHubResult(GetHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHubResult(
            connected_device_count=self.connected_device_count,
            created_at=self.created_at,
            device_auto_provisioning=self.device_auto_provisioning,
            device_count=self.device_count,
            disable_events=self.disable_events,
            enabled=self.enabled,
            endpoint=self.endpoint,
            events_topic_prefix=self.events_topic_prefix,
            hub_ca=self.hub_ca,
            hub_ca_challenge=self.hub_ca_challenge,
            hub_id=self.hub_id,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            product_plan=self.product_plan,
            project_id=self.project_id,
            region=self.region,
            status=self.status,
            updated_at=self.updated_at)


def get_hub(hub_id: Optional[str] = None,
            name: Optional[str] = None,
            project_id: Optional[str] = None,
            region: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHubResult:
    """
    Gets information about an IOT Hub.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_hub = scaleway.iot.get_hub(hub_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str hub_id: The Hub ID.
           Only one of the `name` and `hub_id` should be specified.
    :param str name: The name of the Hub.
           Only one of the `name` and `hub_id` should be specified.
    :param str project_id: The ID of the project the hub is associated with.
    :param str region: `region`) The region in which the hub exists.
    """
    __args__ = dict()
    __args__['hubId'] = hub_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:iot/getHub:getHub', __args__, opts=opts, typ=GetHubResult).value

    return AwaitableGetHubResult(
        connected_device_count=pulumi.get(__ret__, 'connected_device_count'),
        created_at=pulumi.get(__ret__, 'created_at'),
        device_auto_provisioning=pulumi.get(__ret__, 'device_auto_provisioning'),
        device_count=pulumi.get(__ret__, 'device_count'),
        disable_events=pulumi.get(__ret__, 'disable_events'),
        enabled=pulumi.get(__ret__, 'enabled'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        events_topic_prefix=pulumi.get(__ret__, 'events_topic_prefix'),
        hub_ca=pulumi.get(__ret__, 'hub_ca'),
        hub_ca_challenge=pulumi.get(__ret__, 'hub_ca_challenge'),
        hub_id=pulumi.get(__ret__, 'hub_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        product_plan=pulumi.get(__ret__, 'product_plan'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_hub)
def get_hub_output(hub_id: Optional[pulumi.Input[Optional[str]]] = None,
                   name: Optional[pulumi.Input[Optional[str]]] = None,
                   project_id: Optional[pulumi.Input[Optional[str]]] = None,
                   region: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHubResult]:
    """
    Gets information about an IOT Hub.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_hub = scaleway.iot.get_hub(hub_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str hub_id: The Hub ID.
           Only one of the `name` and `hub_id` should be specified.
    :param str name: The name of the Hub.
           Only one of the `name` and `hub_id` should be specified.
    :param str project_id: The ID of the project the hub is associated with.
    :param str region: `region`) The region in which the hub exists.
    """
    ...