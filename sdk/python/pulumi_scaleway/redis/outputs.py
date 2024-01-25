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
    'ClusterAcl',
    'ClusterPrivateNetwork',
    'ClusterPublicNetwork',
    'GetClusterAclResult',
    'GetClusterPrivateNetworkResult',
    'GetClusterPublicNetworkResult',
]

@pulumi.output_type
class ClusterAcl(dict):
    def __init__(__self__, *,
                 ip: str,
                 description: Optional[str] = None,
                 id: Optional[str] = None):
        """
        :param str ip: The ip range to whitelist
               in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        :param str description: A text describing this rule. Default description: `Allow IP`
               
               > The `acl` conflict with `private_network`. Only one should be specified.
        :param str id: The UUID of the private network resource.
        """
        pulumi.set(__self__, "ip", ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The ip range to whitelist
        in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A text describing this rule. Default description: `Allow IP`

        > The `acl` conflict with `private_network`. Only one should be specified.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class ClusterPrivateNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endpointId":
            suggest = "endpoint_id"
        elif key == "serviceIps":
            suggest = "service_ips"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterPrivateNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterPrivateNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterPrivateNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 endpoint_id: Optional[str] = None,
                 service_ips: Optional[Sequence[str]] = None,
                 zone: Optional[str] = None):
        """
        :param str id: The UUID of the private network resource.
        :param str endpoint_id: The ID of the endpoint.
        :param Sequence[str] service_ips: Endpoint IPv4 addresses
               in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at
               least one IP per node or The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
               
               > The `private_network` conflict with `acl`. Only one should be specified.
        :param str zone: `zone`) The zone in which the
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
    def id(self) -> str:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[str]:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="serviceIps")
    def service_ips(self) -> Optional[Sequence[str]]:
        """
        Endpoint IPv4 addresses
        in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at
        least one IP per node or The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.

        > The `private_network` conflict with `acl`. Only one should be specified.
        """
        return pulumi.get(self, "service_ips")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        `zone`) The zone in which the
        Redis Cluster should be created.
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class ClusterPublicNetwork(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 ips: Optional[Sequence[str]] = None,
                 port: Optional[int] = None):
        """
        :param str id: The UUID of the private network resource.
        :param Sequence[str] ips: Lis of IPv4 address of the endpoint (IP address).
        :param int port: TCP port of the endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ips is not None:
            pulumi.set(__self__, "ips", ips)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The UUID of the private network resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Optional[Sequence[str]]:
        """
        Lis of IPv4 address of the endpoint (IP address).
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        TCP port of the endpoint.
        """
        return pulumi.get(self, "port")


@pulumi.output_type
class GetClusterAclResult(dict):
    def __init__(__self__, *,
                 description: str,
                 id: str,
                 ip: str):
        """
        :param str id: The ID of the Redis cluster.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip", ip)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Redis cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")


@pulumi.output_type
class GetClusterPrivateNetworkResult(dict):
    def __init__(__self__, *,
                 endpoint_id: str,
                 id: str,
                 service_ips: Sequence[str],
                 zone: str):
        """
        :param str id: The ID of the Redis cluster.
        :param str zone: `region`) The zone in which the server exists.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "service_ips", service_ips)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Redis cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceIps")
    def service_ips(self) -> Sequence[str]:
        return pulumi.get(self, "service_ips")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        `region`) The zone in which the server exists.
        """
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetClusterPublicNetworkResult(dict):
    def __init__(__self__, *,
                 id: str,
                 ips: Sequence[str],
                 port: int):
        """
        :param str id: The ID of the Redis cluster.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ips", ips)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Redis cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

