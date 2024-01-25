// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about multiple Load Balancer Routes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byFrontendID = scaleway.loadbalancer.getRoutes({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const myKey = scaleway.loadbalancer.getRoutes({
 *     frontendId: "11111111-1111-1111-1111-111111111111",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getRoutes(args?: GetRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetRoutesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:loadbalancer/getRoutes:getRoutes", {
        "frontendId": args.frontendId,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getRoutes.
 */
export interface GetRoutesArgs {
    /**
     * The frontend ID origin of redirection used as a filter. routes with a frontend ID like it are listed.
     */
    frontendId?: string;
    projectId?: string;
    /**
     * `zone`) The zone in which routes exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getRoutes.
 */
export interface GetRoutesResult {
    readonly frontendId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationId: string;
    readonly projectId: string;
    /**
     * List of found routes
     */
    readonly routes: outputs.loadbalancer.GetRoutesRoute[];
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer Routes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byFrontendID = scaleway.loadbalancer.getRoutes({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const myKey = scaleway.loadbalancer.getRoutes({
 *     frontendId: "11111111-1111-1111-1111-111111111111",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getRoutesOutput(args?: GetRoutesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoutesResult> {
    return pulumi.output(args).apply((a: any) => getRoutes(a, opts))
}

/**
 * A collection of arguments for invoking getRoutes.
 */
export interface GetRoutesOutputArgs {
    /**
     * The frontend ID origin of redirection used as a filter. routes with a frontend ID like it are listed.
     */
    frontendId?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which routes exist.
     */
    zone?: pulumi.Input<string>;
}
