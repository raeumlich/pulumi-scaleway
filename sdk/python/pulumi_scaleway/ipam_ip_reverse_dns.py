# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpamIpReverseDnsArgs', 'IpamIpReverseDns']

@pulumi.input_type
class IpamIpReverseDnsArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 hostname: pulumi.Input[str],
                 ipam_ip_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpamIpReverseDns resource.
        :param pulumi.Input[str] address: The IP corresponding to the hostname.
        :param pulumi.Input[str] hostname: The reverse domain name.
        :param pulumi.Input[str] ipam_ip_id: The IPAM IP ID.
        :param pulumi.Input[str] region: `region`) The region of the IP reverse DNS.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "ipam_ip_id", ipam_ip_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        The IP corresponding to the hostname.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Input[str]:
        """
        The reverse domain name.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> pulumi.Input[str]:
        """
        The IPAM IP ID.
        """
        return pulumi.get(self, "ipam_ip_id")

    @ipam_ip_id.setter
    def ipam_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ipam_ip_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the IP reverse DNS.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _IpamIpReverseDnsState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ipam_ip_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpamIpReverseDns resources.
        :param pulumi.Input[str] address: The IP corresponding to the hostname.
        :param pulumi.Input[str] hostname: The reverse domain name.
        :param pulumi.Input[str] ipam_ip_id: The IPAM IP ID.
        :param pulumi.Input[str] region: `region`) The region of the IP reverse DNS.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ipam_ip_id is not None:
            pulumi.set(__self__, "ipam_ip_id", ipam_ip_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP corresponding to the hostname.
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

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The IPAM IP ID.
        """
        return pulumi.get(self, "ipam_ip_id")

    @ipam_ip_id.setter
    def ipam_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipam_ip_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the IP reverse DNS.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class IpamIpReverseDns(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ipam_ip_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Scaleway IPAM IP Reverse DNS.

        ## Import

        IPAM IP reverse DNS can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/ipamIpReverseDns:IpamIpReverseDns main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP corresponding to the hostname.
        :param pulumi.Input[str] hostname: The reverse domain name.
        :param pulumi.Input[str] ipam_ip_id: The IPAM IP ID.
        :param pulumi.Input[str] region: `region`) The region of the IP reverse DNS.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpamIpReverseDnsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Scaleway IPAM IP Reverse DNS.

        ## Import

        IPAM IP reverse DNS can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/ipamIpReverseDns:IpamIpReverseDns main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IpamIpReverseDnsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpamIpReverseDnsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ipam_ip_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpamIpReverseDnsArgs.__new__(IpamIpReverseDnsArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            if hostname is None and not opts.urn:
                raise TypeError("Missing required property 'hostname'")
            __props__.__dict__["hostname"] = hostname
            if ipam_ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'ipam_ip_id'")
            __props__.__dict__["ipam_ip_id"] = ipam_ip_id
            __props__.__dict__["region"] = region
        super(IpamIpReverseDns, __self__).__init__(
            'scaleway:index/ipamIpReverseDns:IpamIpReverseDns',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ipam_ip_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'IpamIpReverseDns':
        """
        Get an existing IpamIpReverseDns resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP corresponding to the hostname.
        :param pulumi.Input[str] hostname: The reverse domain name.
        :param pulumi.Input[str] ipam_ip_id: The IPAM IP ID.
        :param pulumi.Input[str] region: `region`) The region of the IP reverse DNS.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpamIpReverseDnsState.__new__(_IpamIpReverseDnsState)

        __props__.__dict__["address"] = address
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["ipam_ip_id"] = ipam_ip_id
        __props__.__dict__["region"] = region
        return IpamIpReverseDns(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The IP corresponding to the hostname.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The reverse domain name.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ipamIpId")
    def ipam_ip_id(self) -> pulumi.Output[str]:
        """
        The IPAM IP ID.
        """
        return pulumi.get(self, "ipam_ip_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region of the IP reverse DNS.
        """
        return pulumi.get(self, "region")

