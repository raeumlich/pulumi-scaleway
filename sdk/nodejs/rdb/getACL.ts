// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about the RDB instance network Access Control List.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myAcl = scaleway.rdb.getACL({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getACL(args: GetACLArgs, opts?: pulumi.InvokeOptions): Promise<GetACLResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:rdb/getACL:getACL", {
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getACL.
 */
export interface GetACLArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: string;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: string;
}

/**
 * A collection of values returned by getACL.
 */
export interface GetACLResult {
    /**
     * A list of ACLs rules (structure is described below)
     */
    readonly aclRules: outputs.rdb.GetACLAclRule[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region?: string;
}
/**
 * Gets information about the RDB instance network Access Control List.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myAcl = scaleway.rdb.getACL({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getACLOutput(args: GetACLOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetACLResult> {
    return pulumi.output(args).apply((a: any) => getACL(a, opts))
}

/**
 * A collection of arguments for invoking getACL.
 */
export interface GetACLOutputArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}
