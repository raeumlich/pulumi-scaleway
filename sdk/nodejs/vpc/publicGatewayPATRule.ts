// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway VPC Public Gateway PAT (Port Address Translation).
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1#pat-rules-e75d10).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const sg01 = new scaleway.instance.SecurityGroup("sg01", {
 *     inboundDefaultPolicy: "drop",
 *     outboundDefaultPolicy: "accept",
 *     inboundRules: [{
 *         action: "accept",
 *         port: 22,
 *         protocol: "TCP",
 *     }],
 * });
 * const srv01 = new scaleway.instance.Server("srv01", {
 *     type: "PLAY2-NANO",
 *     image: "ubuntu_jammy",
 *     securityGroupId: sg01.id,
 * });
 * const pn01 = new scaleway.vpc.PrivateNetwork("pn01", {});
 * const pnic01 = new scaleway.instance.PrivateNIC("pnic01", {
 *     serverId: srv01.id,
 *     privateNetworkId: pn01.id,
 * });
 * const dhcp01 = new scaleway.vpc.PublicGatewayDHCP("dhcp01", {subnet: "192.168.0.0/24"});
 * const ip01 = new scaleway.vpc.PublicGatewayIP("ip01", {});
 * const pg01 = new scaleway.vpc.PublicGateway("pg01", {
 *     type: "VPC-GW-S",
 *     ipId: ip01.id,
 * });
 * const gn01 = new scaleway.vpc.GatewayNetwork("gn01", {
 *     gatewayId: pg01.id,
 *     privateNetworkId: pn01.id,
 *     dhcpId: dhcp01.id,
 *     cleanupDhcp: true,
 *     enableMasquerade: true,
 * });
 * const rsv01 = new scaleway.vpc.PublicGatewayDHCPReservation("rsv01", {
 *     gatewayNetworkId: gn01.id,
 *     macAddress: pnic01.macAddress,
 *     ipAddress: "192.168.0.7",
 * });
 * // PAT rule for SSH traffic
 * const pat01 = new scaleway.vpc.PublicGatewayPATRule("pat01", {
 *     gatewayId: pg01.id,
 *     privateIp: rsv01.ipAddress,
 *     privatePort: 22,
 *     publicPort: 2202,
 *     protocol: "tcp",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Public gateway PAT rules config can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:vpc/publicGatewayPATRule:PublicGatewayPATRule main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class PublicGatewayPATRule extends pulumi.CustomResource {
    /**
     * Get an existing PublicGatewayPATRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicGatewayPATRuleState, opts?: pulumi.CustomResourceOptions): PublicGatewayPATRule {
        return new PublicGatewayPATRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:vpc/publicGatewayPATRule:PublicGatewayPATRule';

    /**
     * Returns true if the given object is an instance of PublicGatewayPATRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicGatewayPATRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicGatewayPATRule.__pulumiType;
    }

    /**
     * The date and time of the creation of the pat rule config.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The ID of the public gateway.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * The organization ID the pat rule config is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The Private IP to forward data to (IP address).
     */
    public readonly privateIp!: pulumi.Output<string>;
    /**
     * The Private port to translate to.
     */
    public readonly privatePort!: pulumi.Output<number>;
    /**
     * The Protocol the rule should apply to. Possible values are both, tcp and udp.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The Public port to listen on.
     */
    public readonly publicPort!: pulumi.Output<number>;
    /**
     * The date and time of the last update of the pat rule config.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PublicGatewayPATRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicGatewayPATRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicGatewayPATRuleArgs | PublicGatewayPATRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicGatewayPATRuleState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["privatePort"] = state ? state.privatePort : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["publicPort"] = state ? state.publicPort : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as PublicGatewayPATRuleArgs | undefined;
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.privateIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateIp'");
            }
            if ((!args || args.privatePort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privatePort'");
            }
            if ((!args || args.publicPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicPort'");
            }
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["privateIp"] = args ? args.privateIp : undefined;
            resourceInputs["privatePort"] = args ? args.privatePort : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["publicPort"] = args ? args.publicPort : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicGatewayPATRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicGatewayPATRule resources.
 */
export interface PublicGatewayPATRuleState {
    /**
     * The date and time of the creation of the pat rule config.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The ID of the public gateway.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The organization ID the pat rule config is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The Private IP to forward data to (IP address).
     */
    privateIp?: pulumi.Input<string>;
    /**
     * The Private port to translate to.
     */
    privatePort?: pulumi.Input<number>;
    /**
     * The Protocol the rule should apply to. Possible values are both, tcp and udp.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Public port to listen on.
     */
    publicPort?: pulumi.Input<number>;
    /**
     * The date and time of the last update of the pat rule config.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicGatewayPATRule resource.
 */
export interface PublicGatewayPATRuleArgs {
    /**
     * The ID of the public gateway.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * The Private IP to forward data to (IP address).
     */
    privateIp: pulumi.Input<string>;
    /**
     * The Private port to translate to.
     */
    privatePort: pulumi.Input<number>;
    /**
     * The Protocol the rule should apply to. Possible values are both, tcp and udp.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Public port to listen on.
     */
    publicPort: pulumi.Input<number>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}
