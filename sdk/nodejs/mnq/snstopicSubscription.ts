// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage Scaleway Messaging and queuing SNS Topic Subscriptions.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // For default project in default region
 * const mainSNS = new scaleway.mnq.SNS("mainSNS", {});
 * const mainSNSCredentials = new scaleway.mnq.SNSCredentials("mainSNSCredentials", {
 *     projectId: mainSNS.projectId,
 *     permissions: {
 *         canManage: true,
 *         canPublish: true,
 *         canReceive: true,
 *     },
 * });
 * const topic = new scaleway.mnq.SNSTopic("topic", {
 *     projectId: mainSNS.projectId,
 *     accessKey: mainSNSCredentials.accessKey,
 *     secretKey: mainSNSCredentials.secretKey,
 * });
 * const mainSNSTopicSubscription = new scaleway.mnq.SNSTopicSubscription("mainSNSTopicSubscription", {
 *     projectId: mainSNS.projectId,
 *     accessKey: mainSNSCredentials.accessKey,
 *     secretKey: mainSNSCredentials.secretKey,
 *     topicId: topic.id,
 *     protocol: "http",
 *     endpoint: "http://example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * SNS topic subscriptions can be imported using the `{region}/{project-id}/{topic-name}/{subscription-id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:mnq/sNSTopicSubscription:SNSTopicSubscription main fr-par/11111111111111111111111111111111/my-topic/11111111111111111111111111111111
 * ```
 */
export class SNSTopicSubscription extends pulumi.CustomResource {
    /**
     * Get an existing SNSTopicSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SNSTopicSubscriptionState, opts?: pulumi.CustomResourceOptions): SNSTopicSubscription {
        return new SNSTopicSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/sNSTopicSubscription:SNSTopicSubscription';

    /**
     * Returns true if the given object is an instance of SNSTopicSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SNSTopicSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SNSTopicSubscription.__pulumiType;
    }

    /**
     * The access key of the SNS credentials.
     */
    public readonly accessKey!: pulumi.Output<string>;
    /**
     * The ARN of the topic subscription
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Endpoint of the subscription
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Protocol of the SNS Topic Subscription.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Activate JSON Redrive Policy.
     */
    public readonly redrivePolicy!: pulumi.Output<boolean>;
    /**
     * `region`). The region
     * in which sns is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret key of the SNS credentials.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    public readonly snsEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the topic. Either `topicId` or `topicArn` is required.
     */
    public readonly topicArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
     */
    public readonly topicId!: pulumi.Output<string | undefined>;

    /**
     * Create a SNSTopicSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SNSTopicSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SNSTopicSubscriptionArgs | SNSTopicSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SNSTopicSubscriptionState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["redrivePolicy"] = state ? state.redrivePolicy : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["snsEndpoint"] = state ? state.snsEndpoint : undefined;
            resourceInputs["topicArn"] = state ? state.topicArn : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
        } else {
            const args = argsOrState as SNSTopicSubscriptionArgs | undefined;
            if ((!args || args.accessKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessKey'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            resourceInputs["accessKey"] = args?.accessKey ? pulumi.secret(args.accessKey) : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["redrivePolicy"] = args ? args.redrivePolicy : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["snsEndpoint"] = args ? args.snsEndpoint : undefined;
            resourceInputs["topicArn"] = args ? args.topicArn : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SNSTopicSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SNSTopicSubscription resources.
 */
export interface SNSTopicSubscriptionState {
    /**
     * The access key of the SNS credentials.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The ARN of the topic subscription
     */
    arn?: pulumi.Input<string>;
    /**
     * Endpoint of the subscription
     */
    endpoint?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Protocol of the SNS Topic Subscription.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Activate JSON Redrive Policy.
     */
    redrivePolicy?: pulumi.Input<boolean>;
    /**
     * `region`). The region
     * in which sns is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SNS credentials.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    snsEndpoint?: pulumi.Input<string>;
    /**
     * The ARN of the topic. Either `topicId` or `topicArn` is required.
     */
    topicArn?: pulumi.Input<string>;
    /**
     * The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
     */
    topicId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SNSTopicSubscription resource.
 */
export interface SNSTopicSubscriptionArgs {
    /**
     * The access key of the SNS credentials.
     */
    accessKey: pulumi.Input<string>;
    /**
     * Endpoint of the subscription
     */
    endpoint?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Protocol of the SNS Topic Subscription.
     */
    protocol: pulumi.Input<string>;
    /**
     * Activate JSON Redrive Policy.
     */
    redrivePolicy?: pulumi.Input<boolean>;
    /**
     * `region`). The region
     * in which sns is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret key of the SNS credentials.
     */
    secretKey: pulumi.Input<string>;
    /**
     * The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
     */
    snsEndpoint?: pulumi.Input<string>;
    /**
     * The ARN of the topic. Either `topicId` or `topicArn` is required.
     */
    topicArn?: pulumi.Input<string>;
    /**
     * The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
     */
    topicId?: pulumi.Input<string>;
}
