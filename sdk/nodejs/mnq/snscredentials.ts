// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates and manages Scaleway Messaging and queuing SNS Credentials.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainSNS = new scaleway.mnq.SNS("mainSNS", {});
 * const mainSNSCredentials = new scaleway.mnq.SNSCredentials("mainSNSCredentials", {
 *     projectId: mainSNS.projectId,
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
 * SNS credentials can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:mnq/sNSCredentials:SNSCredentials main fr-par/11111111111111111111111111111111
 * ```
 */
export class SNSCredentials extends pulumi.CustomResource {
    /**
     * Get an existing SNSCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SNSCredentialsState, opts?: pulumi.CustomResourceOptions): SNSCredentials {
        return new SNSCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/sNSCredentials:SNSCredentials';

    /**
     * Returns true if the given object is an instance of SNSCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SNSCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SNSCredentials.__pulumiType;
    }

    /**
     * The ID of the key.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * The unique name of the sns credentials.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    public readonly permissions!: pulumi.Output<outputs.mnq.SNSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which sns is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret value of the key.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;

    /**
     * Create a SNSCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SNSCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SNSCredentialsArgs | SNSCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SNSCredentialsState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
        } else {
            const args = argsOrState as SNSCredentialsArgs | undefined;
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
        super(SNSCredentials.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SNSCredentials resources.
 */
export interface SNSCredentialsState {
    /**
     * The ID of the key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The unique name of the sns credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    permissions?: pulumi.Input<inputs.mnq.SNSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which sns is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret value of the key.
     */
    secretKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SNSCredentials resource.
 */
export interface SNSCredentialsArgs {
    /**
     * The unique name of the sns credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated to these credentials. Only one of permissions may be set.
     */
    permissions?: pulumi.Input<inputs.mnq.SNSCredentialsPermissions>;
    /**
     * `projectId`) The ID of the project the sns is enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which sns is enabled.
     */
    region?: pulumi.Input<string>;
}