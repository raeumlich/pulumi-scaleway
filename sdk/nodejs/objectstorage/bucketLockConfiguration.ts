// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Object bucket lock configuration resource.
 * For more information, see [Setting up object lock](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/).
 *
 * ## Example Usage
 * ### Configure an Object Lock for a new bucket
 *
 * Please note that `objectLockEnabled` must be set to `true` before configuring the lock.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainBucket = new scaleway.objectstorage.Bucket("mainBucket", {
 *     acl: "public-read",
 *     objectLockEnabled: true,
 * });
 * const mainBucketLockConfiguration = new scaleway.objectstorage.BucketLockConfiguration("mainBucketLockConfiguration", {
 *     bucket: mainBucket.name,
 *     rule: {
 *         defaultRetention: {
 *             mode: "GOVERNANCE",
 *             days: 1,
 *         },
 *     },
 * });
 * ```
 * ### Configure an Object Lock for an existing bucket
 *
 * You should [contact Scaleway support](https://console.scaleway.com/support/tickets/create) to enable object lock on an existing bucket.
 *
 * ## Import
 *
 * Bucket lock configurations can be imported using the `{region}/{bucketName}` identifier, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket
 * ```
 *
 *  ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash
 *
 * ```sh
 *  $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
 * ```
 */
export class BucketLockConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing BucketLockConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketLockConfigurationState, opts?: pulumi.CustomResourceOptions): BucketLockConfiguration {
        return new BucketLockConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration';

    /**
     * Returns true if the given object is an instance of BucketLockConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketLockConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketLockConfiguration.__pulumiType;
    }

    /**
     * The bucket's name or regional ID.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the Object Lock rule for the specified object.
     */
    public readonly rule!: pulumi.Output<outputs.objectstorage.BucketLockConfigurationRule>;

    /**
     * Create a BucketLockConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketLockConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketLockConfigurationArgs | BucketLockConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketLockConfigurationState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rule"] = state ? state.rule : undefined;
        } else {
            const args = argsOrState as BucketLockConfigurationArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.rule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rule'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rule"] = args ? args.rule : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketLockConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketLockConfiguration resources.
 */
export interface BucketLockConfigurationState {
    /**
     * The bucket's name or regional ID.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the Object Lock rule for the specified object.
     */
    rule?: pulumi.Input<inputs.objectstorage.BucketLockConfigurationRule>;
}

/**
 * The set of arguments for constructing a BucketLockConfiguration resource.
 */
export interface BucketLockConfigurationArgs {
    /**
     * The bucket's name or regional ID.
     */
    bucket: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the Object Lock rule for the specified object.
     */
    rule: pulumi.Input<inputs.objectstorage.BucketLockConfigurationRule>;
}
