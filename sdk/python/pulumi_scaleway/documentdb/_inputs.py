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
    'ReadReplicaDirectAccessArgs',
    'ReadReplicaPrivateNetworkArgs',
]

@pulumi.input_type
class ReadReplicaDirectAccessArgs:
    def __init__(__self__, *,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] endpoint_id: The ID of the endpoint of the read replica.
        :param pulumi.Input[str] hostname: Hostname of the endpoint. Only one of ip and hostname may be set.
        :param pulumi.Input[str] ip: IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        :param pulumi.Input[str] name: Name of the endpoint.
        :param pulumi.Input[int] port: TCP port of the endpoint.
        """
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the endpoint of the read replica.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the endpoint. Only one of ip and hostname may be set.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the endpoint.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


@pulumi.input_type
class ReadReplicaPrivateNetworkArgs:
    def __init__(__self__, *,
                 private_network_id: pulumi.Input[str],
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 service_ip: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] private_network_id: UUID of the private network to be connected to the read replica.
        :param pulumi.Input[str] endpoint_id: The ID of the endpoint of the read replica.
        :param pulumi.Input[str] hostname: Hostname of the endpoint. Only one of ip and hostname may be set.
        :param pulumi.Input[str] ip: IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        :param pulumi.Input[str] name: Name of the endpoint.
        :param pulumi.Input[int] port: TCP port of the endpoint.
        :param pulumi.Input[str] service_ip: The IP network address within the private subnet. This must be an IPv4 address with a
               CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
        """
        pulumi.set(__self__, "private_network_id", private_network_id)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if service_ip is not None:
            pulumi.set(__self__, "service_ip", service_ip)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Input[str]:
        """
        UUID of the private network to be connected to the read replica.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the endpoint of the read replica.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the endpoint. Only one of ip and hostname may be set.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the endpoint.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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

    @property
    @pulumi.getter(name="serviceIp")
    def service_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP network address within the private subnet. This must be an IPv4 address with a
        CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.
        """
        return pulumi.get(self, "service_ip")

    @service_ip.setter
    def service_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_ip", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


