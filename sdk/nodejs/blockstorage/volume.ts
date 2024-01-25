// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Block Volumes.
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const blockVolume = new scaleway.blockstorage.Volume("blockVolume", {
 *     iops: 5000,
 *     sizeInGb: 20,
 * });
 * ```
 *
 * ## Import
 *
 * Block Volumes can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:blockstorage/volume:Volume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:blockstorage/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The maximum IO/s expected, must match available options.
     */
    public readonly iops!: pulumi.Output<number>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
     */
    public readonly sizeInGb!: pulumi.Output<number>;
    /**
     * If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * A list of tags to apply to the volume.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["iops"] = state ? state.iops : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["sizeInGb"] = state ? state.sizeInGb : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.iops === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iops'");
            }
            resourceInputs["iops"] = args ? args.iops : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["sizeInGb"] = args ? args.sizeInGb : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * The maximum IO/s expected, must match available options.
     */
    iops?: pulumi.Input<number>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the volume.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The maximum IO/s expected, must match available options.
     */
    iops: pulumi.Input<number>;
    /**
     * The name of the volume. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the volume is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the volume.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `zone`) The zone in which the volume should be created.
     */
    zone?: pulumi.Input<string>;
}