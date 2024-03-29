// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about multiple IPs managed by IPAM service.
 *
 * ## Examples
 *
 * ### By tag
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byTag = scaleway.ipam.getIPs({
 *     tags: ["tag"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### By type and resource
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const vpc01 = new scaleway.vpc.VPC("vpc01", {});
 * const pn01 = new scaleway.vpc.PrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const redis01 = new scaleway.redis.Cluster("redis01", {
 *     version: "7.0.5",
 *     nodeType: "RED1-XS",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     clusterSize: 3,
 *     privateNetworks: [{
 *         id: pn01.id,
 *     }],
 * });
 * const byTypeAndResource = scaleway.ipam.getIPsOutput({
 *     type: "ipv4",
 *     resource: {
 *         id: redis01.id,
 *         type: "redis_cluster",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIPs(args?: GetIPsArgs, opts?: pulumi.InvokeOptions): Promise<GetIPsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:ipam/getIPs:getIPs", {
        "attached": args.attached,
        "macAddress": args.macAddress,
        "privateNetworkId": args.privateNetworkId,
        "projectId": args.projectId,
        "region": args.region,
        "resource": args.resource,
        "tags": args.tags,
        "type": args.type,
        "zonal": args.zonal,
    }, opts);
}

/**
 * A collection of arguments for invoking getIPs.
 */
export interface GetIPsArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource.
     */
    attached?: boolean;
    /**
     * The Mac Address used as filter.
     */
    macAddress?: string;
    /**
     * The ID of the private network used as filter.
     */
    privateNetworkId?: string;
    /**
     * The ID of the project used as filter.
     */
    projectId?: string;
    /**
     * The region used as filter.
     */
    region?: string;
    /**
     * Filter by resource ID, type or name.
     */
    resource?: inputs.ipam.GetIPsResource;
    /**
     * The tags used as filter.
     */
    tags?: string[];
    /**
     * The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
     */
    type?: string;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: string;
}

/**
 * A collection of values returned by getIPs.
 */
export interface GetIPsResult {
    readonly attached?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of found IPs
     */
    readonly ips: outputs.ipam.GetIPsIp[];
    /**
     * The mac address.
     */
    readonly macAddress?: string;
    readonly organizationId: string;
    readonly privateNetworkId?: string;
    /**
     * The ID of the project the server is associated with.
     */
    readonly projectId: string;
    /**
     * The region in which the IP is.
     */
    readonly region: string;
    /**
     * The list of public IPs of the server.
     */
    readonly resource?: outputs.ipam.GetIPsResource;
    /**
     * The tags associated with the IP.
     */
    readonly tags?: string[];
    /**
     * The type of resource.
     */
    readonly type?: string;
    readonly zonal: string;
}
/**
 * Gets information about multiple IPs managed by IPAM service.
 *
 * ## Examples
 *
 * ### By tag
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byTag = scaleway.ipam.getIPs({
 *     tags: ["tag"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### By type and resource
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const vpc01 = new scaleway.vpc.VPC("vpc01", {});
 * const pn01 = new scaleway.vpc.PrivateNetwork("pn01", {
 *     vpcId: vpc01.id,
 *     ipv4Subnet: {
 *         subnet: "172.16.32.0/22",
 *     },
 * });
 * const redis01 = new scaleway.redis.Cluster("redis01", {
 *     version: "7.0.5",
 *     nodeType: "RED1-XS",
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     clusterSize: 3,
 *     privateNetworks: [{
 *         id: pn01.id,
 *     }],
 * });
 * const byTypeAndResource = scaleway.ipam.getIPsOutput({
 *     type: "ipv4",
 *     resource: {
 *         id: redis01.id,
 *         type: "redis_cluster",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getIPsOutput(args?: GetIPsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIPsResult> {
    return pulumi.output(args).apply((a: any) => getIPs(a, opts))
}

/**
 * A collection of arguments for invoking getIPs.
 */
export interface GetIPsOutputArgs {
    /**
     * Defines whether to filter only for IPs which are attached to a resource.
     */
    attached?: pulumi.Input<boolean>;
    /**
     * The Mac Address used as filter.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The ID of the private network used as filter.
     */
    privateNetworkId?: pulumi.Input<string>;
    /**
     * The ID of the project used as filter.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region used as filter.
     */
    region?: pulumi.Input<string>;
    /**
     * Filter by resource ID, type or name.
     */
    resource?: pulumi.Input<inputs.ipam.GetIPsResourceArgs>;
    /**
     * The tags used as filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
     */
    type?: pulumi.Input<string>;
    /**
     * Only IPs that are zonal, and in this zone, will be returned.
     */
    zonal?: pulumi.Input<string>;
}
