// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway VPC Public Gateway IP.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.vpc.PublicGatewayIP("main", {reverse: "tf.example.com"});
 * const tfA = new scaleway.dns.Record("tfA", {
 *     data: main.address,
 *     dnsZone: "example.com",
 *     priority: 1,
 *     ttl: 3600,
 *     type: "A",
 * });
 * ```
 *
 * ## Import
 *
 * Public gateway can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:vpc/publicGatewayIP:PublicGatewayIP main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class PublicGatewayIP extends pulumi.CustomResource {
    /**
     * Get an existing PublicGatewayIP resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicGatewayIPState, opts?: pulumi.CustomResourceOptions): PublicGatewayIP {
        return new PublicGatewayIP(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:vpc/publicGatewayIP:PublicGatewayIP';

    /**
     * Returns true if the given object is an instance of PublicGatewayIP.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicGatewayIP {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicGatewayIP.__pulumiType;
    }

    /**
     * The IP address itself.
     */
    public /*out*/ readonly address!: pulumi.Output<string>;
    /**
     * The date and time of the creation of the public gateway ip.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The organization ID the public gateway ip is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the public gateway ip is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The reverse domain name for the IP address
     */
    public readonly reverse!: pulumi.Output<string>;
    /**
     * The tags associated with the public gateway IP.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time of the last update of the public gateway ip.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the public gateway ip should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a PublicGatewayIP resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PublicGatewayIPArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicGatewayIPArgs | PublicGatewayIPState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PublicGatewayIPState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["reverse"] = state ? state.reverse : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as PublicGatewayIPArgs | undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["reverse"] = args ? args.reverse : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PublicGatewayIP.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicGatewayIP resources.
 */
export interface PublicGatewayIPState {
    /**
     * The IP address itself.
     */
    address?: pulumi.Input<string>;
    /**
     * The date and time of the creation of the public gateway ip.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The organization ID the public gateway ip is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the public gateway ip is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The reverse domain name for the IP address
     */
    reverse?: pulumi.Input<string>;
    /**
     * The tags associated with the public gateway IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time of the last update of the public gateway ip.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the public gateway ip should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublicGatewayIP resource.
 */
export interface PublicGatewayIPArgs {
    /**
     * `projectId`) The ID of the project the public gateway ip is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The reverse domain name for the IP address
     */
    reverse?: pulumi.Input<string>;
    /**
     * The tags associated with the public gateway IP.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the public gateway ip should be created.
     */
    zone?: pulumi.Input<string>;
}
