// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about multiple Load Balancer ACLs.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byFrontID = scaleway.loadbalancer.getACLs({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const byFrontIDAndName = scaleway.loadbalancer.getACLs({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 *     name: "tf-acls-datasource",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getACLs(args: GetACLsArgs, opts?: pulumi.InvokeOptions): Promise<GetACLsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:loadbalancer/getACLs:getACLs", {
        "frontendId": args.frontendId,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getACLs.
 */
export interface GetACLsArgs {
    /**
     * The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * > **Important:** LB Frontends' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     */
    frontendId: string;
    /**
     * The ACL name used as filter. ACLs with a name like it are listed.
     */
    name?: string;
    projectId?: string;
    /**
     * `zone`) The zone in which ACLs exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getACLs.
 */
export interface GetACLsResult {
    /**
     * List of found ACLs
     */
    readonly acls: outputs.loadbalancer.GetACLsAcl[];
    readonly frontendId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer ACLs.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byFrontID = scaleway.loadbalancer.getACLs({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const byFrontIDAndName = scaleway.loadbalancer.getACLs({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 *     name: "tf-acls-datasource",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getACLsOutput(args: GetACLsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetACLsResult> {
    return pulumi.output(args).apply((a: any) => getACLs(a, opts))
}

/**
 * A collection of arguments for invoking getACLs.
 */
export interface GetACLsOutputArgs {
    /**
     * The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * > **Important:** LB Frontends' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     */
    frontendId: pulumi.Input<string>;
    /**
     * The ACL name used as filter. ACLs with a name like it are listed.
     */
    name?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which ACLs exist.
     */
    zone?: pulumi.Input<string>;
}
