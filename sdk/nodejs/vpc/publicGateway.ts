// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway VPC Public Gateway.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.vpc.PublicGateway("main", {
 *     tags: [
 *         "demo",
 *         "terraform",
 *     ],
 *     type: "VPC-GW-S",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Public gateway can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:vpc/publicGateway:PublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class PublicGateway extends pulumi.CustomResource {
    /**
     * Get an existing PublicGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicGatewayState, opts?: pulumi.CustomResourceOptions): PublicGateway {
        return new PublicGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:vpc/publicGateway:PublicGateway';

    /**
     * Returns true if the given object is an instance of PublicGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicGateway.__pulumiType;
    }

    /**
     * Enable SSH bastion on the gateway
     */
    public readonly bastionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The port on which the SSH bastion will listen.
     */
    public readonly bastionPort!: pulumi.Output<number>;
    /**
     * The date and time of the creation of the public gateway.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Enable SMTP on the gateway
     */
    public readonly enableSmtp!: pulumi.Output<boolean>;
    /**
     * attach an existing flexible IP to the gateway
     */
    public readonly ipId!: pulumi.Output<string>;
    /**
     * The name of the public gateway. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the public gateway is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the public gateway is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The status of the public gateway.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags associated with the public gateway.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The gateway type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the public gateway.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * override the gateway's default recursive DNS servers, if DNS features are enabled.
     */
    public readonly upstreamDnsServers!: pulumi.Output<string[] | undefined>;
    /**
     * `zone`) The zone in which the public gateway should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PublicGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicGatewayArgs | PublicGatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicGatewayState | undefined;
            resourceInputs["bastionEnabled"] = state ? state.bastionEnabled : undefined;
            resourceInputs["bastionPort"] = state ? state.bastionPort : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enableSmtp"] = state ? state.enableSmtp : undefined;
            resourceInputs["ipId"] = state ? state.ipId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["upstreamDnsServers"] = state ? state.upstreamDnsServers : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as PublicGatewayArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["bastionEnabled"] = args ? args.bastionEnabled : undefined;
            resourceInputs["bastionPort"] = args ? args.bastionPort : undefined;
            resourceInputs["enableSmtp"] = args ? args.enableSmtp : undefined;
            resourceInputs["ipId"] = args ? args.ipId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["upstreamDnsServers"] = args ? args.upstreamDnsServers : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicGateway resources.
 */
export interface PublicGatewayState {
    /**
     * Enable SSH bastion on the gateway
     */
    bastionEnabled?: pulumi.Input<boolean>;
    /**
     * The port on which the SSH bastion will listen.
     */
    bastionPort?: pulumi.Input<number>;
    /**
     * The date and time of the creation of the public gateway.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Enable SMTP on the gateway
     */
    enableSmtp?: pulumi.Input<boolean>;
    /**
     * attach an existing flexible IP to the gateway
     */
    ipId?: pulumi.Input<string>;
    /**
     * The name of the public gateway. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the public gateway is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the public gateway is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The status of the public gateway.
     */
    status?: pulumi.Input<string>;
    /**
     * The tags associated with the public gateway.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The gateway type.
     */
    type?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the public gateway.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * override the gateway's default recursive DNS servers, if DNS features are enabled.
     */
    upstreamDnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the public gateway should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicGateway resource.
 */
export interface PublicGatewayArgs {
    /**
     * Enable SSH bastion on the gateway
     */
    bastionEnabled?: pulumi.Input<boolean>;
    /**
     * The port on which the SSH bastion will listen.
     */
    bastionPort?: pulumi.Input<number>;
    /**
     * Enable SMTP on the gateway
     */
    enableSmtp?: pulumi.Input<boolean>;
    /**
     * attach an existing flexible IP to the gateway
     */
    ipId?: pulumi.Input<string>;
    /**
     * The name of the public gateway. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the public gateway is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The tags associated with the public gateway.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The gateway type.
     */
    type: pulumi.Input<string>;
    /**
     * override the gateway's default recursive DNS servers, if DNS features are enabled.
     */
    upstreamDnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the public gateway should be created.
     */
    zone?: pulumi.Input<string>;
}
