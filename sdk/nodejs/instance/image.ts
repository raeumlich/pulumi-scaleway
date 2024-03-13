// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Compute Images.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#images-41389b).
 *
 * ## Example Usage
 *
 * ### From a volume
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const volume = new scaleway.instance.Volume("volume", {
 *     type: "b_ssd",
 *     sizeInGb: 20,
 * });
 * const volumeSnapshot = new scaleway.instance.Snapshot("volumeSnapshot", {volumeId: volume.id});
 * const volumeImage = new scaleway.instance.Image("volumeImage", {rootVolumeId: volumeSnapshot.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### From a server
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const server = new scaleway.instance.Server("server", {
 *     image: "ubuntu_jammy",
 *     type: "DEV1-S",
 * });
 * const serverSnapshot = new scaleway.instance.Snapshot("serverSnapshot", {volumeId: scaleway_instance_server.main.root_volume[0].volume_id});
 * const serverImage = new scaleway.instance.Image("serverImage", {rootVolumeId: serverSnapshot.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Images can be imported using the `{zone}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:instance/image:Image main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:instance/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    public readonly additionalVolumeIds!: pulumi.Output<string | undefined>;
    /**
     * The description of the extra volumes attached to the image.
     */
    public /*out*/ readonly additionalVolumes!: pulumi.Output<outputs.instance.ImageAdditionalVolume[]>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    public readonly architecture!: pulumi.Output<string | undefined>;
    /**
     * Date of the volume creation.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * ID of the server the image is based on (in case it is a backup).
     */
    public /*out*/ readonly fromServerId!: pulumi.Output<string>;
    /**
     * Date of volume latest update.
     */
    public /*out*/ readonly modificationDate!: pulumi.Output<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the image is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * The ID of the project the image is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Set to `true` if the image is public.
     */
    public readonly public!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    public readonly rootVolumeId!: pulumi.Output<string>;
    /**
     * State of the volume.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A list of tags to apply to the image.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The zone in which the image should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageState | undefined;
            resourceInputs["additionalVolumeIds"] = state ? state.additionalVolumeIds : undefined;
            resourceInputs["additionalVolumes"] = state ? state.additionalVolumes : undefined;
            resourceInputs["architecture"] = state ? state.architecture : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["fromServerId"] = state ? state.fromServerId : undefined;
            resourceInputs["modificationDate"] = state ? state.modificationDate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["public"] = state ? state.public : undefined;
            resourceInputs["rootVolumeId"] = state ? state.rootVolumeId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            if ((!args || args.rootVolumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootVolumeId'");
            }
            resourceInputs["additionalVolumeIds"] = args ? args.additionalVolumeIds : undefined;
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["public"] = args ? args.public : undefined;
            resourceInputs["rootVolumeId"] = args ? args.rootVolumeId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["additionalVolumes"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["fromServerId"] = undefined /*out*/;
            resourceInputs["modificationDate"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    additionalVolumeIds?: pulumi.Input<string>;
    /**
     * The description of the extra volumes attached to the image.
     */
    additionalVolumes?: pulumi.Input<pulumi.Input<inputs.instance.ImageAdditionalVolume>[]>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * Date of the volume creation.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * ID of the server the image is based on (in case it is a backup).
     */
    fromServerId?: pulumi.Input<string>;
    /**
     * Date of volume latest update.
     */
    modificationDate?: pulumi.Input<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the image is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Set to `true` if the image is public.
     */
    public?: pulumi.Input<boolean>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    rootVolumeId?: pulumi.Input<string>;
    /**
     * State of the volume.
     */
    state?: pulumi.Input<string>;
    /**
     * A list of tags to apply to the image.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone in which the image should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * List of IDs of the snapshots of the additional volumes to be attached to the image.
     *
     * > **Important:** For now it is only possible to have 1 additional_volume.
     */
    additionalVolumeIds?: pulumi.Input<string>;
    /**
     * The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * The name of the image. If not provided it will be randomly generated.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the image is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Set to `true` if the image is public.
     */
    public?: pulumi.Input<boolean>;
    /**
     * The ID of the snapshot of the volume to be used as root in the image.
     */
    rootVolumeId: pulumi.Input<string>;
    /**
     * A list of tags to apply to the image.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone in which the image should be created.
     */
    zone?: pulumi.Input<string>;
}
