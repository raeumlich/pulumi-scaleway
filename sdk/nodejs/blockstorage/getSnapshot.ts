// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a Block Snapshot.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mySnapshot = scaleway.blockstorage.getSnapshot({
 *     snapshotId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSnapshot(args?: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:blockstorage/getSnapshot:getSnapshot", {
        "name": args.name,
        "projectId": args.projectId,
        "snapshotId": args.snapshotId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshot.
 */
export interface GetSnapshotArgs {
    /**
     * The name of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the snapshot is associated with.
     */
    projectId?: string;
    /**
     * The ID of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: string;
    /**
     * The ID of the volume from which the snapshot has been created.
     */
    volumeId?: string;
    /**
     * `zone`) The zone in which the snapshot exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getSnapshot.
 */
export interface GetSnapshotResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly projectId?: string;
    readonly snapshotId?: string;
    readonly tags: string[];
    readonly volumeId?: string;
    readonly zone?: string;
}
/**
 * Gets information about a Block Snapshot.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mySnapshot = scaleway.blockstorage.getSnapshot({
 *     snapshotId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSnapshotOutput(args?: GetSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotResult> {
    return pulumi.output(args).apply((a: any) => getSnapshot(a, opts))
}

/**
 * A collection of arguments for invoking getSnapshot.
 */
export interface GetSnapshotOutputArgs {
    /**
     * The name of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the snapshot is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the snapshot. Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The ID of the volume from which the snapshot has been created.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the snapshot exists.
     */
    zone?: pulumi.Input<string>;
}
