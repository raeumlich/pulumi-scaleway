// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a Block Volume.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myVolume = scaleway.blockstorage.getVolume({
 *     volumeId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVolume(args?: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:blockstorage/getVolume:getVolume", {
        "name": args.name,
        "projectId": args.projectId,
        "volumeId": args.volumeId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    /**
     * The name of the volume. Only one of `name` and `volumeId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the volume is associated with.
     */
    projectId?: string;
    /**
     * The ID of the volume. Only one of `name` and `volumeId` should be specified.
     */
    volumeId?: string;
    /**
     * `zone`) The zone in which the volume exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly iops: number;
    readonly name?: string;
    readonly projectId?: string;
    readonly sizeInGb: number;
    readonly snapshotId: string;
    readonly tags: string[];
    readonly volumeId?: string;
    readonly zone?: string;
}
/**
 * Gets information about a Block Volume.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myVolume = scaleway.blockstorage.getVolume({
 *     volumeId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getVolumeOutput(args?: GetVolumeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply((a: any) => getVolume(a, opts))
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeOutputArgs {
    /**
     * The name of the volume. Only one of `name` and `volumeId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The ID of the volume. Only one of `name` and `volumeId` should be specified.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the volume exists.
     */
    zone?: pulumi.Input<string>;
}
