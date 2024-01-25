// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Messaging and queuing SQS Credentials.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
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
 *         canManage: false,
 *         canReceive: true,
 *         canPublish: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * SQS credentials can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:mnq/sQSCredentials:SQSCredentials main fr-par/11111111111111111111111111111111
 * ```
 */
export class SQSCredentials extends pulumi.CustomResource {
    /**
     * Get an existing SQSCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SQSCredentialsState, opts?: pulumi.CustomResourceOptions): SQSCredentials {
        return new SQSCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/sQSCredentials:SQSCredentials';

    /**
     * Returns true if the given object is an instance of SQSCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SQSCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SQSCredentials.__pulumiType;
    }

    /**
     * The ID of the key.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * The unique name of the sqs credentials.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    public readonly permissions!: pulumi.Output<outputs.mnq.SQSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret value of the key.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;

    /**
     * Create a SQSCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SQSCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SQSCredentialsArgs | SQSCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SQSCredentialsState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
        } else {
            const args = argsOrState as SQSCredentialsArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SQSCredentials.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SQSCredentials resources.
 */
export interface SQSCredentialsState {
    /**
     * The ID of the key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The unique name of the sqs credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    permissions?: pulumi.Input<inputs.mnq.SQSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret value of the key.
     */
    secretKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SQSCredentials resource.
 */
export interface SQSCredentialsArgs {
    /**
     * The unique name of the sqs credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    permissions?: pulumi.Input<inputs.mnq.SQSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sqs is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which sqs is enabled.
     */
    region?: pulumi.Input<string>;
}