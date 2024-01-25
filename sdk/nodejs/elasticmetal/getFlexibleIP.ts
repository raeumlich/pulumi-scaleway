// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a Flexible IP.
 */
export function getFlexibleIP(args?: GetFlexibleIPArgs, opts?: pulumi.InvokeOptions): Promise<GetFlexibleIPResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:elasticmetal/getFlexibleIP:getFlexibleIP", {
        "flexibleIpId": args.flexibleIpId,
        "ipAddress": args.ipAddress,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlexibleIP.
 */
export interface GetFlexibleIPArgs {
    flexibleIpId?: string;
    /**
     * The IP address.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipAddress?: string;
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getFlexibleIP.
 */
export interface GetFlexibleIPResult {
    readonly createdAt: string;
    readonly description: string;
    readonly flexibleIpId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddress?: string;
    readonly isIpv6: boolean;
    /**
     * (Defaults to provider `organizationId`) The ID of the organization the IP is in.
     */
    readonly organizationId: string;
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    readonly projectId: string;
    /**
     * The reverse domain associated with this IP.
     */
    readonly reverse: string;
    /**
     * The associated server ID if any
     */
    readonly serverId: string;
    readonly status: string;
    readonly tags: string[];
    readonly updatedAt: string;
    readonly zone: string;
}
/**
 * Gets information about a Flexible IP.
 */
export function getFlexibleIPOutput(args?: GetFlexibleIPOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlexibleIPResult> {
    return pulumi.output(args).apply((a: any) => getFlexibleIP(a, opts))
}

/**
 * A collection of arguments for invoking getFlexibleIP.
 */
export interface GetFlexibleIPOutputArgs {
    flexibleIpId?: pulumi.Input<string>;
    /**
     * The IP address.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * (Defaults to provider `projectId`) The ID of the project the IP is in.
     */
    projectId?: pulumi.Input<string>;
}