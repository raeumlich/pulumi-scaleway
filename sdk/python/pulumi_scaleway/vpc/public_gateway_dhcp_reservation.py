# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PublicGatewayDHCPReservationArgs', 'PublicGatewayDHCPReservation']

@pulumi.input_type
class PublicGatewayDHCPReservationArgs:
    def __init__(__self__, *,
                 gateway_network_id: pulumi.Input[str],
                 ip_address: pulumi.Input[str],
                 mac_address: pulumi.Input[str],
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PublicGatewayDHCPReservation resource.
        :param pulumi.Input[str] gateway_network_id: The ID of the owning GatewayNetwork.
        :param pulumi.Input[str] ip_address: The IP address to give to the machine (IP address).
        :param pulumi.Input[str] mac_address: The MAC address to give a static entry to.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway DHCP config should be created.
        """
        pulumi.set(__self__, "gateway_network_id", gateway_network_id)
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "mac_address", mac_address)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="gatewayNetworkId")
    def gateway_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the owning GatewayNetwork.
        """
        return pulumi.get(self, "gateway_network_id")

    @gateway_network_id.setter
    def gateway_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_network_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        """
        The IP address to give to the machine (IP address).
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> pulumi.Input[str]:
        """
        The MAC address to give a static entry to.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "mac_address", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the public gateway DHCP config should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _PublicGatewayDHCPReservationState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 gateway_network_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PublicGatewayDHCPReservation resources.
        :param pulumi.Input[str] created_at: The date and time of the creation of the public gateway DHCP config.
        :param pulumi.Input[str] gateway_network_id: The ID of the owning GatewayNetwork.
        :param pulumi.Input[str] hostname: The Hostname of the client machine.
        :param pulumi.Input[str] ip_address: The IP address to give to the machine (IP address).
        :param pulumi.Input[str] mac_address: The MAC address to give a static entry to.
        :param pulumi.Input[str] type: The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the public gateway DHCP config.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway DHCP config should be created.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if gateway_network_id is not None:
            pulumi.set(__self__, "gateway_network_id", gateway_network_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the public gateway DHCP config.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="gatewayNetworkId")
    def gateway_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the owning GatewayNetwork.
        """
        return pulumi.get(self, "gateway_network_id")

    @gateway_network_id.setter
    def gateway_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_network_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The Hostname of the client machine.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address to give to the machine (IP address).
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC address to give a static entry to.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the public gateway DHCP config.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the public gateway DHCP config should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class PublicGatewayDHCPReservation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_network_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages the [Scaleway DHCP Reservations](https://www.scaleway.com/en/docs/network/vpc/concepts/#dhcp).

        The static associations are used to assign IP addresses based on the MAC addresses of the Instance.

        Statically assigned IP addresses should fall within the configured subnet, but be outside of the dynamic range.

        For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544) and [configuration guide](https://www.scaleway.com/en/docs/network/vpc/how-to/configure-a-public-gateway/#how-to-review-and-configure-dhcp).

        [DHCP reservations](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6) hold both dynamic DHCP leases (IP addresses dynamically assigned by the gateway to instances) and static user-created DHCP reservations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_private_network = scaleway.vpc.PrivateNetwork("mainPrivateNetwork")
        main_server = scaleway.instance.Server("mainServer",
            image="ubuntu_jammy",
            type="DEV1-S",
            zone="fr-par-1",
            private_networks=[scaleway.instance.ServerPrivateNetworkArgs(
                pn_id=main_private_network.id,
            )])
        main_public_gateway_ip = scaleway.vpc.PublicGatewayIP("mainPublicGatewayIP")
        main_public_gateway_dhcp = scaleway.vpc.PublicGatewayDHCP("mainPublicGatewayDHCP", subnet="192.168.1.0/24")
        main_public_gateway = scaleway.vpc.PublicGateway("mainPublicGateway",
            type="VPC-GW-S",
            ip_id=main_public_gateway_ip.id)
        main_gateway_network = scaleway.vpc.GatewayNetwork("mainGatewayNetwork",
            gateway_id=main_public_gateway.id,
            private_network_id=main_private_network.id,
            dhcp_id=main_public_gateway_dhcp.id,
            cleanup_dhcp=True,
            enable_masquerade=True,
            opts=pulumi.ResourceOptions(depends_on=[
                    main_public_gateway_ip,
                    main_private_network,
                ]))
        main_public_gateway_dhcp_reservation = scaleway.vpc.PublicGatewayDHCPReservation("mainPublicGatewayDHCPReservation",
            gateway_network_id=main_gateway_network.id,
            mac_address=main_server.private_networks[0].mac_address,
            ip_address="192.168.1.1")
        ```

        ## Import

        Public gateway DHCP Reservation config can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:vpc/publicGatewayDHCPReservation:PublicGatewayDHCPReservation main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_network_id: The ID of the owning GatewayNetwork.
        :param pulumi.Input[str] ip_address: The IP address to give to the machine (IP address).
        :param pulumi.Input[str] mac_address: The MAC address to give a static entry to.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway DHCP config should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PublicGatewayDHCPReservationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages the [Scaleway DHCP Reservations](https://www.scaleway.com/en/docs/network/vpc/concepts/#dhcp).

        The static associations are used to assign IP addresses based on the MAC addresses of the Instance.

        Statically assigned IP addresses should fall within the configured subnet, but be outside of the dynamic range.

        For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544) and [configuration guide](https://www.scaleway.com/en/docs/network/vpc/how-to/configure-a-public-gateway/#how-to-review-and-configure-dhcp).

        [DHCP reservations](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6) hold both dynamic DHCP leases (IP addresses dynamically assigned by the gateway to instances) and static user-created DHCP reservations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_private_network = scaleway.vpc.PrivateNetwork("mainPrivateNetwork")
        main_server = scaleway.instance.Server("mainServer",
            image="ubuntu_jammy",
            type="DEV1-S",
            zone="fr-par-1",
            private_networks=[scaleway.instance.ServerPrivateNetworkArgs(
                pn_id=main_private_network.id,
            )])
        main_public_gateway_ip = scaleway.vpc.PublicGatewayIP("mainPublicGatewayIP")
        main_public_gateway_dhcp = scaleway.vpc.PublicGatewayDHCP("mainPublicGatewayDHCP", subnet="192.168.1.0/24")
        main_public_gateway = scaleway.vpc.PublicGateway("mainPublicGateway",
            type="VPC-GW-S",
            ip_id=main_public_gateway_ip.id)
        main_gateway_network = scaleway.vpc.GatewayNetwork("mainGatewayNetwork",
            gateway_id=main_public_gateway.id,
            private_network_id=main_private_network.id,
            dhcp_id=main_public_gateway_dhcp.id,
            cleanup_dhcp=True,
            enable_masquerade=True,
            opts=pulumi.ResourceOptions(depends_on=[
                    main_public_gateway_ip,
                    main_private_network,
                ]))
        main_public_gateway_dhcp_reservation = scaleway.vpc.PublicGatewayDHCPReservation("mainPublicGatewayDHCPReservation",
            gateway_network_id=main_gateway_network.id,
            mac_address=main_server.private_networks[0].mac_address,
            ip_address="192.168.1.1")
        ```

        ## Import

        Public gateway DHCP Reservation config can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:vpc/publicGatewayDHCPReservation:PublicGatewayDHCPReservation main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param PublicGatewayDHCPReservationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PublicGatewayDHCPReservationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_network_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 mac_address: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PublicGatewayDHCPReservationArgs.__new__(PublicGatewayDHCPReservationArgs)

            if gateway_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_network_id'")
            __props__.__dict__["gateway_network_id"] = gateway_network_id
            if ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'ip_address'")
            __props__.__dict__["ip_address"] = ip_address
            if mac_address is None and not opts.urn:
                raise TypeError("Missing required property 'mac_address'")
            __props__.__dict__["mac_address"] = mac_address
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["hostname"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["updated_at"] = None
        super(PublicGatewayDHCPReservation, __self__).__init__(
            'scaleway:vpc/publicGatewayDHCPReservation:PublicGatewayDHCPReservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            gateway_network_id: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            mac_address: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'PublicGatewayDHCPReservation':
        """
        Get an existing PublicGatewayDHCPReservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time of the creation of the public gateway DHCP config.
        :param pulumi.Input[str] gateway_network_id: The ID of the owning GatewayNetwork.
        :param pulumi.Input[str] hostname: The Hostname of the client machine.
        :param pulumi.Input[str] ip_address: The IP address to give to the machine (IP address).
        :param pulumi.Input[str] mac_address: The MAC address to give a static entry to.
        :param pulumi.Input[str] type: The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the public gateway DHCP config.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway DHCP config should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PublicGatewayDHCPReservationState.__new__(_PublicGatewayDHCPReservationState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["gateway_network_id"] = gateway_network_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["mac_address"] = mac_address
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["zone"] = zone
        return PublicGatewayDHCPReservation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the public gateway DHCP config.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="gatewayNetworkId")
    def gateway_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the owning GatewayNetwork.
        """
        return pulumi.get(self, "gateway_network_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        The Hostname of the client machine.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP address to give to the machine (IP address).
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> pulumi.Output[str]:
        """
        The MAC address to give a static entry to.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the public gateway DHCP config.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the public gateway DHCP config should be created.
        """
        return pulumi.get(self, "zone")

