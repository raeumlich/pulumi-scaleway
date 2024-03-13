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
    'IPResourceArgs',
    'IPReverseArgs',
    'IPSourceArgs',
    'GetIPResourceArgs',
    'GetIPsResourceArgs',
]

@pulumi.input_type
class IPResourceArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of the resource that the IP is bound to.
        :param pulumi.Input[str] mac_address: The MAC Address of the resource the IP is attached to.
        :param pulumi.Input[str] name: The name of the resource the IP is attached to.
        :param pulumi.Input[str] type: The type of resource the IP is attached to.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource that the IP is bound to.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC Address of the resource the IP is attached to.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource the IP is attached to.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resource the IP is attached to.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class IPReverseArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] address: Request a specific IP in the requested source pool.
        :param pulumi.Input[str] hostname: The reverse domain name.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Request a specific IP in the requested source pool.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The reverse domain name.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)


@pulumi.input_type
class IPSourceArgs:
    def __init__(__self__, *,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 zonal: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] private_network_id: The private network the IP lives in if the IP is a private IP.
        :param pulumi.Input[str] subnet_id: The private network subnet the IP lives in if the IP is a private IP in a private network.
        :param pulumi.Input[str] zonal: The zone the IP lives in if the IP is a public zoned one
        """
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if zonal is not None:
            pulumi.set(__self__, "zonal", zonal)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The private network the IP lives in if the IP is a private IP.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The private network subnet the IP lives in if the IP is a private IP in a private network.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def zonal(self) -> Optional[pulumi.Input[str]]:
        """
        The zone the IP lives in if the IP is a public zoned one
        """
        return pulumi.get(self, "zonal")

    @zonal.setter
    def zonal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zonal", value)


@pulumi.input_type
class GetIPResourceArgs:
    def __init__(__self__, *,
                 type: str,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        :param str type: The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        :param str id: The ID of the resource that the IP is bound to.
        :param str name: The name of the resource to get the IP from.
        """
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: str):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the resource that the IP is bound to.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource to get the IP from.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class GetIPsResourceArgs:
    def __init__(__self__, *,
                 type: str,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        """
        :param str type: The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        :param str id: The ID of the resource that the IP is bound to.
        :param str name: The name of the resource to get the IP from.
        """
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: str):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the resource that the IP is bound to.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource to get the IP from.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[str]):
        pulumi.set(self, "name", value)


