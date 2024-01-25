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
    'ClusterAclArgs',
    'ClusterPrivateNetworkArgs',
    'ClusterPublicNetworkArgs',
]

@pulumi.input_type
class ClusterAclArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] ip: The ip range to whitelist
               in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        :param pulumi.Input[str] description: A text describing this rule. Default description: `Allow IP`
               
               > The `acl` conflict with `private_network`. Only one should be specified.
        :param pulumi.Input[str] id: The UUID of the private network resource.
        """
        pulumi.set(__self__, "ip", ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        The ip range to whitelist
        in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A text describing this rule. Default description: `Allow IP`

        > The `acl` conflict with `private_network`. Only one should be specified.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class ClusterPrivateNetworkArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 service_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The UUID of the private network resource.
        :param pulumi.Input[str] endpoint_id: The ID of the endpoint.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_ips: Endpoint IPv4 addresses
               in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at
               least one IP per node or The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
               
               > The `private_network` conflict with `acl`. Only one should be specified.
        :param pulumi.Input[str] zone: `zone`) The zone in which the
               Redis Cluster should be created.
        """
        pulumi.set(__self__, "id", id)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if service_ips is not None:
            pulumi.set(__self__, "service_ips", service_ips)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="serviceIps")
    def service_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Endpoint IPv4 addresses
        in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at
        least one IP per node or The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.

        > The `private_network` conflict with `acl`. Only one should be specified.
        """
        return pulumi.get(self, "service_ips")

    @service_ips.setter
    def service_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "service_ips", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the
        Redis Cluster should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class ClusterPublicNetworkArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] id: The UUID of the private network resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ips: Lis of IPv4 address of the endpoint (IP address).
        :param pulumi.Input[int] port: TCP port of the endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ips is not None:
            pulumi.set(__self__, "ips", ips)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Lis of IPv4 address of the endpoint (IP address).
        """
        return pulumi.get(self, "ips")

    @ips.setter
    def ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ips", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        TCP port of the endpoint.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


