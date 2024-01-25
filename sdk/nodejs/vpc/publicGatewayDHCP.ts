// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway VPC Public Gateway DHCP.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.vpc.PublicGatewayDHCP("main", {subnet: "192.168.1.0/24"});
 * ```
 *
 * ## Import
 *
 * Public gateway DHCP config can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:vpc/publicGatewayDHCP:PublicGatewayDHCP main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class PublicGatewayDHCP extends pulumi.CustomResource {
    /**
     * Get an existing PublicGatewayDHCP resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicGatewayDHCPState, opts?: pulumi.CustomResourceOptions): PublicGatewayDHCP {
        return new PublicGatewayDHCP(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:vpc/publicGatewayDHCP:PublicGatewayDHCP';

    /**
     * Returns true if the given object is an instance of PublicGatewayDHCP.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicGatewayDHCP {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicGatewayDHCP.__pulumiType;
    }

    /**
     * The IP address of the public gateway DHCP config.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The date and time of the creation of the public gateway DHCP config.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
     */
    public readonly dnsLocalName!: pulumi.Output<string>;
    /**
     * Additional DNS search paths
     */
    public readonly dnsSearches!: pulumi.Output<string[]>;
    /**
     * Override the DNS server list pushed to DHCP clients, instead of the gateway itself
     */
    public readonly dnsServersOverrides!: pulumi.Output<string[]>;
    /**
     * Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
     */
    public readonly enableDynamic!: pulumi.Output<boolean>;
    /**
     * The organization ID the public gateway DHCP config is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
     */
    public readonly poolHigh!: pulumi.Output<string>;
    /**
     * Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
     */
    public readonly poolLow!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the public gateway DHCP config is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
     */
    public readonly pushDefaultRoute!: pulumi.Output<boolean>;
    /**
     * Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
     */
    public readonly pushDnsServer!: pulumi.Output<boolean>;
    /**
     * After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
     */
    public readonly rebindTimer!: pulumi.Output<number>;
    /**
     * After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
     */
    public readonly renewTimer!: pulumi.Output<number>;
    /**
     * The subnet to associate with the public gateway DHCP config.
     */
    public readonly subnet!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the public gateway DHCP config.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
     */
    public readonly validLifetime!: pulumi.Output<number>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PublicGatewayDHCP resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicGatewayDHCPArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicGatewayDHCPArgs | PublicGatewayDHCPState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicGatewayDHCPState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dnsLocalName"] = state ? state.dnsLocalName : undefined;
            resourceInputs["dnsSearches"] = state ? state.dnsSearches : undefined;
            resourceInputs["dnsServersOverrides"] = state ? state.dnsServersOverrides : undefined;
            resourceInputs["enableDynamic"] = state ? state.enableDynamic : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["poolHigh"] = state ? state.poolHigh : undefined;
            resourceInputs["poolLow"] = state ? state.poolLow : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pushDefaultRoute"] = state ? state.pushDefaultRoute : undefined;
            resourceInputs["pushDnsServer"] = state ? state.pushDnsServer : undefined;
            resourceInputs["rebindTimer"] = state ? state.rebindTimer : undefined;
            resourceInputs["renewTimer"] = state ? state.renewTimer : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["validLifetime"] = state ? state.validLifetime : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as PublicGatewayDHCPArgs | undefined;
            if ((!args || args.subnet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnet'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["dnsLocalName"] = args ? args.dnsLocalName : undefined;
            resourceInputs["dnsSearches"] = args ? args.dnsSearches : undefined;
            resourceInputs["dnsServersOverrides"] = args ? args.dnsServersOverrides : undefined;
            resourceInputs["enableDynamic"] = args ? args.enableDynamic : undefined;
            resourceInputs["poolHigh"] = args ? args.poolHigh : undefined;
            resourceInputs["poolLow"] = args ? args.poolLow : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["pushDefaultRoute"] = args ? args.pushDefaultRoute : undefined;
            resourceInputs["pushDnsServer"] = args ? args.pushDnsServer : undefined;
            resourceInputs["rebindTimer"] = args ? args.rebindTimer : undefined;
            resourceInputs["renewTimer"] = args ? args.renewTimer : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["validLifetime"] = args ? args.validLifetime : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicGatewayDHCP.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicGatewayDHCP resources.
 */
export interface PublicGatewayDHCPState {
    /**
     * The IP address of the public gateway DHCP config.
     */
    address?: pulumi.Input<string>;
    /**
     * The date and time of the creation of the public gateway DHCP config.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
     */
    dnsLocalName?: pulumi.Input<string>;
    /**
     * Additional DNS search paths
     */
    dnsSearches?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Override the DNS server list pushed to DHCP clients, instead of the gateway itself
     */
    dnsServersOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
     */
    enableDynamic?: pulumi.Input<boolean>;
    /**
     * The organization ID the public gateway DHCP config is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
     */
    poolHigh?: pulumi.Input<string>;
    /**
     * Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
     */
    poolLow?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the public gateway DHCP config is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
     */
    pushDefaultRoute?: pulumi.Input<boolean>;
    /**
     * Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
     */
    pushDnsServer?: pulumi.Input<boolean>;
    /**
     * After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
     */
    rebindTimer?: pulumi.Input<number>;
    /**
     * After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
     */
    renewTimer?: pulumi.Input<number>;
    /**
     * The subnet to associate with the public gateway DHCP config.
     */
    subnet?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the public gateway DHCP config.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
     */
    validLifetime?: pulumi.Input<number>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicGatewayDHCP resource.
 */
export interface PublicGatewayDHCPArgs {
    /**
     * The IP address of the public gateway DHCP config.
     */
    address?: pulumi.Input<string>;
    /**
     * TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
     */
    dnsLocalName?: pulumi.Input<string>;
    /**
     * Additional DNS search paths
     */
    dnsSearches?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Override the DNS server list pushed to DHCP clients, instead of the gateway itself
     */
    dnsServersOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
     */
    enableDynamic?: pulumi.Input<boolean>;
    /**
     * High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
     */
    poolHigh?: pulumi.Input<string>;
    /**
     * Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
     */
    poolLow?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the public gateway DHCP config is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
     */
    pushDefaultRoute?: pulumi.Input<boolean>;
    /**
     * Whether the gateway should push custom DNS servers to clients. This allows for instance hostname > IP resolution. Defaults to `true`.
     */
    pushDnsServer?: pulumi.Input<boolean>;
    /**
     * After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `validLifetime`. Defaults to 51m (3060s).
     */
    rebindTimer?: pulumi.Input<number>;
    /**
     * After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebindTimer`. Defaults to 50m (3000s).
     */
    renewTimer?: pulumi.Input<number>;
    /**
     * The subnet to associate with the public gateway DHCP config.
     */
    subnet: pulumi.Input<string>;
    /**
     * For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
     */
    validLifetime?: pulumi.Input<number>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}
