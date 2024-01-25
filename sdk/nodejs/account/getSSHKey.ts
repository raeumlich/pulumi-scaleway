// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.account.getSSHKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getSSHKey(args?: GetSSHKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetSSHKeyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:account/getSSHKey:getSSHKey", {
        "name": args.name,
        "projectId": args.projectId,
        "sshKeyId": args.sshKeyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSSHKey.
 */
export interface GetSSHKeyArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: string;
    /**
     * `projectId`) The ID of the project the SSH key is associated with.
     */
    projectId?: string;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: string;
}

/**
 * A collection of values returned by getSSHKey.
 */
export interface GetSSHKeyResult {
    readonly createdAt: string;
    readonly disabled: boolean;
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The ID of the organization the SSH key is associated with.
     */
    readonly organizationId: string;
    readonly projectId?: string;
    /**
     * The SSH public key string
     */
    readonly publicKey: string;
    readonly sshKeyId?: string;
    readonly updatedAt: string;
}
/**
 * Use this data source to get SSH key information based on its ID or name.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.account.getSSHKey({
 *     sshKeyId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getSSHKeyOutput(args?: GetSSHKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSSHKeyResult> {
    return pulumi.output(args).apply((a: any) => getSSHKey(a, opts))
}

/**
 * A collection of arguments for invoking getSSHKey.
 */
export interface GetSSHKeyOutputArgs {
    /**
     * The SSH key name. Only one of `name` and `sshKeyId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH key is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The SSH key id. Only one of `name` and `sshKeyId` should be specified.
     */
    sshKeyId?: pulumi.Input<string>;
}
