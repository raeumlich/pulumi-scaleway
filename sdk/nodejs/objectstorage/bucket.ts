// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway object storage buckets.
 * For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const someBucket = new scaleway.objectstorage.Bucket("someBucket", {tags: {
 *     key: "value",
 * }});
 * ```
 * ### Creating the bucket in a specific project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const someBucket = new scaleway.objectstorage.Bucket("someBucket", {projectId: "11111111-1111-1111-1111-111111111111"});
 * ```
 * ### Using object lifecycle
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.objectstorage.Bucket("main", {
 *     lifecycleRules: [
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 365,
 *             },
 *             id: "id1",
 *             prefix: "path1/",
 *             transitions: [{
 *                 days: 120,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 50,
 *             },
 *             id: "id2",
 *             prefix: "path2/",
 *         },
 *         {
 *             enabled: false,
 *             expiration: {
 *                 days: 1,
 *             },
 *             id: "id3",
 *             prefix: "path3/",
 *             tags: {
 *                 tagKey: "tagValue",
 *                 terraform: "hashicorp",
 *             },
 *         },
 *         {
 *             enabled: true,
 *             id: "id4",
 *             tags: {
 *                 tag1: "value1",
 *             },
 *             transitions: [{
 *                 days: 1,
 *                 storageClass: "GLACIER",
 *             }],
 *         },
 *         {
 *             abortIncompleteMultipartUploadDays: 30,
 *             enabled: true,
 *         },
 *     ],
 *     region: "fr-par",
 * });
 * ```
 *
 * ## Import
 *
 * Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:objectstorage/bucket:Bucket some_bucket fr-par/some-bucket
 * ```
 *
 *  If you are importing a bucket from a specific project (that is not your default project), you can use the following syntaxbash
 *
 * ```sh
 *  $ pulumi import scaleway:objectstorage/bucket:Bucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
 * ```
 */
export class Bucket extends pulumi.CustomResource {
    /**
     * Get an existing Bucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketState, opts?: pulumi.CustomResourceOptions): Bucket {
        return new Bucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:objectstorage/bucket:Bucket';

    /**
     * Returns true if the given object is an instance of Bucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bucket.__pulumiType;
    }

    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * API URL of the bucket
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    public readonly corsRules!: pulumi.Output<outputs.objectstorage.BucketCorsRule[] | undefined>;
    /**
     * The endpoint URL of the bucket
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.objectstorage.BucketLifecycleRule[] | undefined>;
    /**
     * The name of the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable object lock
     */
    public readonly objectLockEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * The `acl` attribute is deprecated. See scaleway.objectstorage.BucketACL resource documentation.
     * Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list of tags (key / value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    public readonly versioning!: pulumi.Output<outputs.objectstorage.BucketVersioning>;

    /**
     * Create a Bucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketArgs | BucketState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketState | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["apiEndpoint"] = state ? state.apiEndpoint : undefined;
            resourceInputs["corsRules"] = state ? state.corsRules : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectLockEnabled"] = state ? state.objectLockEnabled : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
        } else {
            const args = argsOrState as BucketArgs | undefined;
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["corsRules"] = args ? args.corsRules : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["objectLockEnabled"] = args ? args.objectLockEnabled : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["apiEndpoint"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bucket.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bucket resources.
 */
export interface BucketState {
    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
     */
    acl?: pulumi.Input<string>;
    /**
     * API URL of the bucket
     */
    apiEndpoint?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.objectstorage.BucketCorsRule>[]>;
    /**
     * The endpoint URL of the bucket
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.objectstorage.BucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable object lock
     */
    objectLockEnabled?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * The `acl` attribute is deprecated. See scaleway.objectstorage.BucketACL resource documentation.
     * Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key / value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.objectstorage.BucketVersioning>;
}

/**
 * The set of arguments for constructing a Bucket resource.
 */
export interface BucketArgs {
    /**
     * (Deprecated) The canned ACL you want to apply to the bucket.
     *
     * @deprecated ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
     */
    acl?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    corsRules?: pulumi.Input<pulumi.Input<inputs.objectstorage.BucketCorsRule>[]>;
    /**
     * Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
     */
    lifecycleRules?: pulumi.Input<pulumi.Input<inputs.objectstorage.BucketLifecycleRule>[]>;
    /**
     * The name of the bucket.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable object lock
     */
    objectLockEnabled?: pulumi.Input<boolean>;
    /**
     * `projectId`) The ID of the project the bucket is associated with.
     *
     * The `acl` attribute is deprecated. See scaleway.objectstorage.BucketACL resource documentation.
     * Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of tags (key / value) for the bucket.
     *
     * * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
     * Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    versioning?: pulumi.Input<inputs.objectstorage.BucketVersioning>;
}