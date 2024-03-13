// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Compute Snapshots.
 * For more information,
 * see [the documentation](https://developers.scaleway.com/en/products/instance/api/#snapshots-756fae).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.instance.Snapshot("main", {volumeId: "11111111-1111-1111-1111-111111111111"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Example with Unified type snapshot
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainVolume = new scaleway.instance.Volume("mainVolume", {
 *     type: "l_ssd",
 *     sizeInGb: 10,
 * });
 * const mainServer = new scaleway.instance.Server("mainServer", {
 *     image: "ubuntu_jammy",
 *     type: "DEV1-S",
 *     rootVolume: {
 *         sizeInGb: 10,
 *         volumeType: "l_ssd",
 *     },
 *     additionalVolumeIds: [mainVolume.id],
 * });
 * const mainSnapshot = new scaleway.instance.Snapshot("mainSnapshot", {
 *     volumeId: mainVolume.id,
 *     type: "unified",
 * }, {
 *     dependsOn: [mainServer],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Example importing a local qcow2 file
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const bucket = new scaleway.objectstorage.Bucket("bucket", {});
 * const qcow = new scaleway.objectstorage.Item("qcow", {
 *     bucket: bucket.name,
 *     key: "server.qcow2",
 *     file: "myqcow.qcow2",
 * });
 * const snapshot = new scaleway.instance.Snapshot("snapshot", {
 *     type: "unified",
 *     "import": {
 *         bucket: qcow.bucket,
 *         key: qcow.key,
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Snapshots can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:instance/snapshot:Snapshot main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:instance/snapshot:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * The snapshot creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Import a snapshot from a qcow2 file located in a bucket
     */
    public readonly import!: pulumi.Output<outputs.instance.SnapshotImport | undefined>;
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the snapshot is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the snapshot is
     * associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * (Optional) The size of the snapshot.
     */
    public /*out*/ readonly sizeInGb!: pulumi.Output<number>;
    /**
     * A list of tags to apply to the snapshot.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
     * Updates to this field will recreate a new resource.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    public readonly volumeId!: pulumi.Output<string | undefined>;
    /**
     * `zone`) The zone in which
     * the snapshot should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["import"] = state ? state.import : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["sizeInGb"] = state ? state.sizeInGb : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            resourceInputs["import"] = args ? args.import : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["sizeInGb"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * The snapshot creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Import a snapshot from a qcow2 file located in a bucket
     */
    import?: pulumi.Input<inputs.instance.SnapshotImport>;
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the snapshot is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the snapshot is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * (Optional) The size of the snapshot.
     */
    sizeInGb?: pulumi.Input<number>;
    /**
     * A list of tags to apply to the snapshot.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
     * Updates to this field will recreate a new resource.
     */
    type?: pulumi.Input<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * Import a snapshot from a qcow2 file located in a bucket
     */
    import?: pulumi.Input<inputs.instance.SnapshotImport>;
    /**
     * The name of the snapshot. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the snapshot is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the snapshot.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
     * Updates to this field will recreate a new resource.
     */
    type?: pulumi.Input<string>;
    /**
     * The ID of the volume to take a snapshot from.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the snapshot should be created.
     */
    zone?: pulumi.Input<string>;
}
