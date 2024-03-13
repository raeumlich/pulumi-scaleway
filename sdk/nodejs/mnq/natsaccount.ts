// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Messaging and queuing Nats Accounts.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)
 * To use Scaleway's provider with official nats jetstream provider, check out the corresponding guide
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.mnq.NATSAccount("main", {});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Namespaces can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:mnq/nATSAccount:NATSAccount main fr-par/11111111111111111111111111111111
 * ```
 */
export class NATSAccount extends pulumi.CustomResource {
    /**
     * Get an existing NATSAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NATSAccountState, opts?: pulumi.CustomResourceOptions): NATSAccount {
        return new NATSAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/nATSAccount:NATSAccount';

    /**
     * Returns true if the given object is an instance of NATSAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NATSAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NATSAccount.__pulumiType;
    }

    /**
     * The endpoint of the NATS service for this account.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The unique name of the nats account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the
     * account is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a NATSAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NATSAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NATSAccountArgs | NATSAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NATSAccountState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as NATSAccountArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NATSAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NATSAccount resources.
 */
export interface NATSAccountState {
    /**
     * The endpoint of the NATS service for this account.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The unique name of the nats account.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the
     * account is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NATSAccount resource.
 */
export interface NATSAccountArgs {
    /**
     * The unique name of the nats account.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the
     * account is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which the account should be created.
     */
    region?: pulumi.Input<string>;
}
