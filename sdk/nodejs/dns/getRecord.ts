// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about a domain record.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byContent = scaleway.dns.getRecord({
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     name: "www",
 *     type: "A",
 * });
 * const byId = scaleway.dns.getRecord({
 *     dnsZone: "domain.tld",
 *     recordId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRecord(args?: GetRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:dns/getRecord:getRecord", {
        "data": args.data,
        "dnsZone": args.dnsZone,
        "name": args.name,
        "projectId": args.projectId,
        "recordId": args.recordId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordArgs {
    /**
     * The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     * Cannot be used with `recordId`.
     */
    data?: string;
    /**
     * The IP address.
     */
    dnsZone?: string;
    /**
     * The name of the record (can be an empty string for a root record).
     * Cannot be used with `recordId`.
     */
    name?: string;
    /**
     * `projectId`) The ID of the project the domain is associated with.
     */
    projectId?: string;
    /**
     * The record ID.
     * Cannot be used with `name`, `type` and `data`.
     */
    recordId?: string;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     * Cannot be used with `recordId`.
     */
    type?: string;
}

/**
 * A collection of values returned by getRecord.
 */
export interface GetRecordResult {
    readonly data?: string;
    readonly dnsZone?: string;
    readonly fqdn: string;
    /**
     * Dynamic record base on user geolocalisation (More information about dynamic records)
     */
    readonly geoIps: outputs.dns.GetRecordGeoIp[];
    /**
     * Dynamic record base on URL resolve (More information about dynamic records)
     */
    readonly httpServices: outputs.dns.GetRecordHttpService[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keepEmptyZone: boolean;
    readonly name?: string;
    /**
     * The priority of the record (mostly used with an `MX` record)
     */
    readonly priority: number;
    readonly projectId?: string;
    readonly recordId?: string;
    readonly rootZone: boolean;
    /**
     * Time To Live of the record in seconds.
     */
    readonly ttl: number;
    readonly type?: string;
    /**
     * Dynamic record based on the client’s (resolver) subnet (More information about dynamic records)
     */
    readonly views: outputs.dns.GetRecordView[];
    /**
     * Dynamic record base on IP weights (More information about dynamic records)
     */
    readonly weighteds: outputs.dns.GetRecordWeighted[];
}
/**
 * Gets information about a domain record.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byContent = scaleway.dns.getRecord({
 *     data: "1.2.3.4",
 *     dnsZone: "domain.tld",
 *     name: "www",
 *     type: "A",
 * });
 * const byId = scaleway.dns.getRecord({
 *     dnsZone: "domain.tld",
 *     recordId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRecordOutput(args?: GetRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecordResult> {
    return pulumi.output(args).apply((a: any) => getRecord(a, opts))
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordOutputArgs {
    /**
     * The content of the record (an IPv4 for an `A`, a string for a `TXT`...).
     * Cannot be used with `recordId`.
     */
    data?: pulumi.Input<string>;
    /**
     * The IP address.
     */
    dnsZone?: pulumi.Input<string>;
    /**
     * The name of the record (can be an empty string for a root record).
     * Cannot be used with `recordId`.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the domain is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The record ID.
     * Cannot be used with `name`, `type` and `data`.
     */
    recordId?: pulumi.Input<string>;
    /**
     * The type of the record (`A`, `AAAA`, `MX`, `CNAME`, `DNAME`, `ALIAS`, `NS`, `PTR`, `SRV`, `TXT`, `TLSA`, or `CAA`).
     * Cannot be used with `recordId`.
     */
    type?: pulumi.Input<string>;
}
