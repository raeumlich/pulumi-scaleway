// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about multiple Load Balancer IPs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.loadbalancer.getIPs({
 *     ipCidrRange: "0.0.0.0/0",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getIPs(args?: GetIPsArgs, opts?: pulumi.InvokeOptions): Promise<GetIPsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:loadbalancer/getIPs:getIPs", {
        "ipCidrRange": args.ipCidrRange,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getIPs.
 */
export interface GetIPsArgs {
    /**
     * The IP CIDR range used as a filter. IPs within a CIDR block like it are listed.
     */
    ipCidrRange?: string;
    /**
     * The ID of the project the load-balancer is associated with.
     */
    projectId?: string;
    /**
     * `zone`) The zone in which IPs exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getIPs.
 */
export interface GetIPsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipCidrRange?: string;
    /**
     * List of found IPs
     */
    readonly ips: outputs.loadbalancer.GetIPsIp[];
    /**
     * The organization ID the load-balancer is associated with.
     */
    readonly organizationId: string;
    /**
     * The ID of the project the load-balancer is associated with.
     */
    readonly projectId: string;
    /**
     * The zone in which the load-balancer is.
     */
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer IPs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.loadbalancer.getIPs({
 *     ipCidrRange: "0.0.0.0/0",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getIPsOutput(args?: GetIPsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIPsResult> {
    return pulumi.output(args).apply((a: any) => getIPs(a, opts))
}

/**
 * A collection of arguments for invoking getIPs.
 */
export interface GetIPsOutputArgs {
    /**
     * The IP CIDR range used as a filter. IPs within a CIDR block like it are listed.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * The ID of the project the load-balancer is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which IPs exist.
     */
    zone?: pulumi.Input<string>;
}
