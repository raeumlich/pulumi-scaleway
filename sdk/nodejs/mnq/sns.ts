// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Activate Scaleway Messaging and queuing SNS for a project.
 * For further information please check
 * our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
 *
 * ## Example Usage
 * ### Basic
 *
 * Activate SNS for default project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.mnq.SNS("main", {});
 * ```
 *
 * Activate SNS for a specific project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const project = scaleway.account.getProject({
 *     name: "default",
 * });
 * // For specific project in default region
 * const forProject = new scaleway.mnq.SNS("forProject", {projectId: project.then(project => project.id)});
 * ```
 *
 * ## Import
 *
 * SNS status can be imported using the `{region}/{project_id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:mnq/sNS:SNS main fr-par/11111111111111111111111111111111
 * ```
 */
export class SNS extends pulumi.CustomResource {
    /**
     * Get an existing SNS resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SNSState, opts?: pulumi.CustomResourceOptions): SNS {
        return new SNS(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:mnq/sNS:SNS';

    /**
     * Returns true if the given object is an instance of SNS.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SNS {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SNS.__pulumiType;
    }

    /**
     * The endpoint of the SNS service for this project.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the sns will be enabled for.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region
     * in which sns will be enabled.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a SNS resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SNSArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SNSArgs | SNSState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SNSState | undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as SNSArgs | undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SNS.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SNS resources.
 */
export interface SNSState {
    /**
     * The endpoint of the SNS service for this project.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the sns will be enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which sns will be enabled.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SNS resource.
 */
export interface SNSArgs {
    /**
     * `projectId`) The ID of the project the sns will be enabled for.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region
     * in which sns will be enabled.
     */
    region?: pulumi.Input<string>;
}
