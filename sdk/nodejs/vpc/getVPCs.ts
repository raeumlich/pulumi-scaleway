// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about multiple Virtual Private Clouds.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.vpc.getVPCs({
 *     name: "tf-vpc-datasource",
 *     region: "nl-ams",
 * });
 * ```
 */
export function getVPCs(args?: GetVPCsArgs, opts?: pulumi.InvokeOptions): Promise<GetVPCsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:vpc/getVPCs:getVPCs", {
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVPCs.
 */
export interface GetVPCsArgs {
    /**
     * The VPC name used as filter. VPCs with a name like it are listed.
     */
    name?: string;
    /**
     * The ID of the project the VPC is associated with.
     */
    projectId?: string;
    /**
     * `region`). The region in which vpcs exist.
     */
    region?: string;
    /**
     * List of tags used as filter. VPCs with these exact tags are listed.
     */
    tags?: string[];
}

/**
 * A collection of values returned by getVPCs.
 */
export interface GetVPCsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The organization ID the VPC is associated with.
     */
    readonly organizationId: string;
    /**
     * The ID of the project the VPC is associated with.
     */
    readonly projectId: string;
    readonly region: string;
    readonly tags?: string[];
    /**
     * List of found vpcs
     */
    readonly vpcs: outputs.vpc.GetVPCsVpc[];
}
/**
 * Gets information about multiple Virtual Private Clouds.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.vpc.getVPCs({
 *     name: "tf-vpc-datasource",
 *     region: "nl-ams",
 * });
 * ```
 */
export function getVPCsOutput(args?: GetVPCsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVPCsResult> {
    return pulumi.output(args).apply((a: any) => getVPCs(a, opts))
}

/**
 * A collection of arguments for invoking getVPCs.
 */
export interface GetVPCsOutputArgs {
    /**
     * The VPC name used as filter. VPCs with a name like it are listed.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the VPC is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which vpcs exist.
     */
    region?: pulumi.Input<string>;
    /**
     * List of tags used as filter. VPCs with these exact tags are listed.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}