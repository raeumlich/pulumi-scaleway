// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Messaging and queuing SQS Queues.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-queues/)
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainSQS = new scaleway.mnq.SQS("mainSQS", {});
 * const mainSQSCredentials = new scaleway.mnq.SQSCredentials("mainSQSCredentials", {
 *     projectId: mainSQS.projectId,
 *     permissions: {
 *         canManage: true,
 *         canReceive: false,
 *         canPublish: false,
 *     },
 * });
 * const mainSQSQueue = new scaleway.mnq.SQSQueue("mainSQSQueue", {
 *     projectId: mainSQS.projectId,
 *     sqsEndpoint: mainSQS.endpoint,
 *     accessKey: mainSQSCredentials.accessKey,
 *     secretKey: mainSQSCredentials.secretKey,
 * });
 * ```
 */
export class SQSQueue extends pulumi.CustomResource {
    /**
     * Get an existing SQSQueue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SQSQueueState, opts?: pulumi.CustomResourceOptions): SQSQueue {
        return new SQSQueue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/sQSQueue:SQSQueue';

    /**
     * Returns true if the given object is an instance of SQSQueue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SQSQueue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SQSQueue.__pulumiType;
    }

    /**
     * The access key of the SQS queue.
     */
    public readonly accessKey!: pulumi.Output<string>;
    /**
     * Specifies whether to enable content-based deduplication. Defaults to `false`.
     */
    public readonly contentBasedDeduplication!: pulumi.Output<boolean>;
    /**
     * Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
     */
    public readonly fifoQueue!: pulumi.Output<boolean>;
    /**
     * The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
     */
    public readonly messageMaxAge!: pulumi.Output<number | undefined>;
    /**
     * The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
     */
    public readonly messageMaxSize!: pulumi.Output<number | undefined>;
    /**
     * The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
     */
    public readonly receiveWaitTimeSeconds!: pulumi.Output<number | undefined>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret key of the SQS queue.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
     */
    public readonly sqsEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The URL of the queue.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
     */
    public readonly visibilityTimeoutSeconds!: pulumi.Output<number | undefined>;

    /**
     * Create a SQSQueue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SQSQueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SQSQueueArgs | SQSQueueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SQSQueueState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["contentBasedDeduplication"] = state ? state.contentBasedDeduplication : undefined;
            resourceInputs["fifoQueue"] = state ? state.fifoQueue : undefined;
            resourceInputs["messageMaxAge"] = state ? state.messageMaxAge : undefined;
            resourceInputs["messageMaxSize"] = state ? state.messageMaxSize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["receiveWaitTimeSeconds"] = state ? state.receiveWaitTimeSeconds : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["sqsEndpoint"] = state ? state.sqsEndpoint : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["visibilityTimeoutSeconds"] = state ? state.visibilityTimeoutSeconds : undefined;
        } else {
            const args = argsOrState as SQSQueueArgs | undefined;
            if ((!args || args.accessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessKey'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            resourceInputs["fifoQueue"] = args ? args.fifoQueue : undefined;
            resourceInputs["messageMaxAge"] = args ? args.messageMaxAge : undefined;
            resourceInputs["messageMaxSize"] = args ? args.messageMaxSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["receiveWaitTimeSeconds"] = args ? args.receiveWaitTimeSeconds : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["sqsEndpoint"] = args ? args.sqsEndpoint : undefined;
            resourceInputs["visibilityTimeoutSeconds"] = args ? args.visibilityTimeoutSeconds : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SQSQueue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SQSQueue resources.
 */
export interface SQSQueueState {
    /**
     * The access key of the SQS queue.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * Specifies whether to enable content-based deduplication. Defaults to `false`.
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
     */
    fifoQueue?: pulumi.Input<boolean>;
    /**
     * The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
     */
    messageMaxAge?: pulumi.Input<number>;
    /**
     * The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
     */
    messageMaxSize?: pulumi.Input<number>;
    /**
     * The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
     */
    receiveWaitTimeSeconds?: pulumi.Input<number>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SQS queue.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
     */
    sqsEndpoint?: pulumi.Input<string>;
    /**
     * The URL of the queue.
     */
    url?: pulumi.Input<string>;
    /**
     * The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
     */
    visibilityTimeoutSeconds?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SQSQueue resource.
 */
export interface SQSQueueArgs {
    /**
     * The access key of the SQS queue.
     */
    accessKey: pulumi.Input<string>;
    /**
     * Specifies whether to enable content-based deduplication. Defaults to `false`.
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
     */
    fifoQueue?: pulumi.Input<boolean>;
    /**
     * The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
     */
    messageMaxAge?: pulumi.Input<number>;
    /**
     * The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
     */
    messageMaxSize?: pulumi.Input<number>;
    /**
     * The unique name of the sqs queue. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
     */
    receiveWaitTimeSeconds?: pulumi.Input<number>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SQS queue.
     */
    secretKey: pulumi.Input<string>;
    /**
     * The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
     */
    sqsEndpoint?: pulumi.Input<string>;
    /**
     * The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
     */
    visibilityTimeoutSeconds?: pulumi.Input<number>;
}
