// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Create and manage Scaleway RDB database privilege.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainInstance = new scaleway.rdb.Instance("mainInstance", {
 *     nodeType: "DB-DEV-S",
 *     engine: "PostgreSQL-11",
 *     isHaCluster: true,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 * });
 * const mainDatabase = new scaleway.rdb.Database("mainDatabase", {instanceId: mainInstance.id});
 * const mainUser = new scaleway.rdb.User("mainUser", {
 *     instanceId: mainInstance.id,
 *     password: "thiZ_is_v&ry_s3cret",
 *     isAdmin: false,
 * });
 * const mainPrivilege = new scaleway.rdb.Privilege("mainPrivilege", {
 *     instanceId: mainInstance.id,
 *     userName: "my-db-user",
 *     databaseName: "my-db-name",
 *     permission: "all",
 * }, {
 *     dependsOn: [
 *         mainUser,
 *         mainDatabase,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:rdb/privilege:Privilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
 * ```
 */
export class Privilege extends pulumi.CustomResource {
    /**
     * Get an existing Privilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivilegeState, opts?: pulumi.CustomResourceOptions): Privilege {
        return new Privilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:rdb/privilege:Privilege';

    /**
     * Returns true if the given object is an instance of Privilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Privilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Privilege.__pulumiType;
    }

    /**
     * Name of the database (e.g. `my-db-name`).
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * UUID of the rdb instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    public readonly permission!: pulumi.Output<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a Privilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivilegeArgs | PrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivilegeState | undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as PrivilegeArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.permission === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permission'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Privilege.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Privilege resources.
 */
export interface PrivilegeState {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName?: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Privilege resource.
 */
export interface PrivilegeArgs {
    /**
     * Name of the database (e.g. `my-db-name`).
     */
    databaseName: pulumi.Input<string>;
    /**
     * UUID of the rdb instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
     */
    permission: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of the user (e.g. `my-db-user`).
     */
    userName: pulumi.Input<string>;
}