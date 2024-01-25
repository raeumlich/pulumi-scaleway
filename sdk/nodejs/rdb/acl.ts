// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Database instance authorized IPs.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.rdb.ACL("main", {
 *     instanceId: scaleway_rdb_instance.main.id,
 *     aclRules: [{
 *         ip: "1.2.3.4/32",
 *         description: "foo",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Database Instance can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:rdb/aCL:ACL acl01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ACL extends pulumi.CustomResource {
    /**
     * Get an existing ACL resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ACLState, opts?: pulumi.CustomResourceOptions): ACL {
        return new ACL(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:rdb/aCL:ACL';

    /**
     * Returns true if the given object is an instance of ACL.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ACL {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ACL.__pulumiType;
    }

    /**
     * A list of ACLs (structure is described below)
     */
    public readonly aclRules!: pulumi.Output<outputs.rdb.ACLAclRule[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a ACL resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ACLArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ACLArgs | ACLState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ACLState | undefined;
            resourceInputs["aclRules"] = state ? state.aclRules : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ACLArgs | undefined;
            if ((!args || args.aclRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclRules'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["aclRules"] = args ? args.aclRules : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ACL.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ACL resources.
 */
export interface ACLState {
    /**
     * A list of ACLs (structure is described below)
     */
    aclRules?: pulumi.Input<pulumi.Input<inputs.rdb.ACLAclRule>[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ACL resource.
 */
export interface ACLArgs {
    /**
     * A list of ACLs (structure is described below)
     */
    aclRules: pulumi.Input<pulumi.Input<inputs.rdb.ACLAclRule>[]>;
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database ACL.
     */
    instanceId: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}