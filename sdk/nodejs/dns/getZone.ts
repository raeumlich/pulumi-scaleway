// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a domain zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.dns.getZone({
 *     domain: "scaleway-terraform.com",
 *     subdomain: "test",
 * });
 * ```
 */
export function getZone(args?: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:dns/getZone:getZone", {
        "domain": args.domain,
        "subdomain": args.subdomain,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    /**
     * The domain where the DNS zone will be created.
     */
    domain?: string;
    /**
     * The subdomain(zone name) to create in the domain.
     */
    subdomain?: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Message
     */
    readonly message: string;
    /**
     * NameServer list for zone.
     */
    readonly ns: string[];
    /**
     * NameServer default list for zone.
     */
    readonly nsDefaults: string[];
    /**
     * NameServer master list for zone.
     */
    readonly nsMasters: string[];
    readonly projectId: string;
    /**
     * The domain zone status.
     */
    readonly status: string;
    readonly subdomain?: string;
    /**
     * The date and time of the last update of the DNS zone.
     */
    readonly updatedAt: string;
}
/**
 * Gets information about a domain zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.dns.getZone({
 *     domain: "scaleway-terraform.com",
 *     subdomain: "test",
 * });
 * ```
 */
export function getZoneOutput(args?: GetZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneResult> {
    return pulumi.output(args).apply((a: any) => getZone(a, opts))
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    /**
     * The domain where the DNS zone will be created.
     */
    domain?: pulumi.Input<string>;
    /**
     * The subdomain(zone name) to create in the domain.
     */
    subdomain?: pulumi.Input<string>;
}