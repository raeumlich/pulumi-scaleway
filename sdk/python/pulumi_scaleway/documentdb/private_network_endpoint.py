# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateNetworkEndpointArgs', 'PrivateNetworkEndpoint']

@pulumi.input_type
class PrivateNetworkEndpointArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 private_network_id: pulumi.Input[str],
                 ip_net: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateNetworkEndpoint resource.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
        :param pulumi.Input[str] private_network_id: The ID of the private network.
        :param pulumi.Input[str] ip_net: The IP network address within the private subnet. This must be an IPv4 address with a
               CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
        :param pulumi.Input[int] port: Port in the Private Network.
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "private_network_id", private_network_id)
        if ip_net is not None:
            pulumi.set(__self__, "ip_net", ip_net)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the documentdb instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the private network.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="ipNet")
    def ip_net(self) -> Optional[pulumi.Input[str]]:
        """
        The IP network address within the private subnet. This must be an IPv4 address with a
        CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.
        """
        return pulumi.get(self, "ip_net")

    @ip_net.setter
    def ip_net(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_net", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port in the Private Network.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _PrivateNetworkEndpointState:
    def __init__(__self__, *,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_net: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateNetworkEndpoint resources.
        :param pulumi.Input[str] hostname: Hostname of the endpoint.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
        :param pulumi.Input[str] ip: IPv4 address on the network.
        :param pulumi.Input[str] ip_net: The IP network address within the private subnet. This must be an IPv4 address with a
               CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
        :param pulumi.Input[str] name: Name of the endpoint.
        :param pulumi.Input[int] port: Port in the Private Network.
        :param pulumi.Input[str] private_network_id: The ID of the private network.
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_net is not None:
            pulumi.set(__self__, "ip_net", ip_net)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the endpoint.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the documentdb instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address on the network.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipNet")
    def ip_net(self) -> Optional[pulumi.Input[str]]:
        """
        The IP network address within the private subnet. This must be an IPv4 address with a
        CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.
        """
        return pulumi.get(self, "ip_net")

    @ip_net.setter
    def ip_net(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_net", value)

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
        Port in the Private Network.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the private network.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class PrivateNetworkEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_net: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:documentdb/privateNetworkEndpoint:PrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
        :param pulumi.Input[str] ip_net: The IP network address within the private subnet. This must be an IPv4 address with a
               CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
        :param pulumi.Input[int] port: Port in the Private Network.
        :param pulumi.Input[str] private_network_id: The ID of the private network.
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateNetworkEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:documentdb/privateNetworkEndpoint:PrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param PrivateNetworkEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateNetworkEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip_net: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateNetworkEndpointArgs.__new__(PrivateNetworkEndpointArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["ip_net"] = ip_net
            __props__.__dict__["port"] = port
            if private_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'private_network_id'")
            __props__.__dict__["private_network_id"] = private_network_id
            __props__.__dict__["region"] = region
            __props__.__dict__["zone"] = zone
            __props__.__dict__["hostname"] = None
            __props__.__dict__["ip"] = None
            __props__.__dict__["name"] = None
        super(PrivateNetworkEndpoint, __self__).__init__(
            'scaleway:documentdb/privateNetworkEndpoint:PrivateNetworkEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip_net: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            private_network_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'PrivateNetworkEndpoint':
        """
        Get an existing PrivateNetworkEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: Hostname of the endpoint.
        :param pulumi.Input[str] instance_id: UUID of the documentdb instance.
        :param pulumi.Input[str] ip: IPv4 address on the network.
        :param pulumi.Input[str] ip_net: The IP network address within the private subnet. This must be an IPv4 address with a
               CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
               service if not set.
        :param pulumi.Input[str] name: Name of the endpoint.
        :param pulumi.Input[int] port: Port in the Private Network.
        :param pulumi.Input[str] private_network_id: The ID of the private network.
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] zone: The zone you want to attach the resource to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateNetworkEndpointState.__new__(_PrivateNetworkEndpointState)

        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip_net"] = ip_net
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["private_network_id"] = private_network_id
        __props__.__dict__["region"] = region
        __props__.__dict__["zone"] = zone
        return PrivateNetworkEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Hostname of the endpoint.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the documentdb instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IPv4 address on the network.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipNet")
    def ip_net(self) -> pulumi.Output[str]:
        """
        The IP network address within the private subnet. This must be an IPv4 address with a
        CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        service if not set.
        """
        return pulumi.get(self, "ip_net")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the endpoint.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port in the Private Network.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the private network.
        """
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone you want to attach the resource to
        """
        return pulumi.get(self, "zone")

