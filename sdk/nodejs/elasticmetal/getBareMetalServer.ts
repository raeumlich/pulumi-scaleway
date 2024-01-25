// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about a baremetal server.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.elasticmetal.getBareMetalServer({
 *     name: "foobar",
 *     zone: "fr-par-2",
 * });
 * const byId = scaleway.elasticmetal.getBareMetalServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getBareMetalServer(args?: GetBareMetalServerArgs, opts?: pulumi.InvokeOptions): Promise<GetBareMetalServerResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:elasticmetal/getBareMetalServer:getBareMetalServer", {
        "name": args.name,
        "projectId": args.projectId,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBareMetalServer.
 */
export interface GetBareMetalServerArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the baremetal server is associated with.
     */
    projectId?: string;
    serverId?: string;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getBareMetalServer.
 */
export interface GetBareMetalServerResult {
    readonly description: string;
    readonly domain: string;
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly installConfigAfterward: boolean;
    readonly ips: outputs.elasticmetal.GetBareMetalServerIp[];
    readonly ipv4s: outputs.elasticmetal.GetBareMetalServerIpv4[];
    readonly ipv6s: outputs.elasticmetal.GetBareMetalServerIpv6[];
    readonly name?: string;
    readonly offer: string;
    readonly offerId: string;
    readonly offerName: string;
    readonly options: outputs.elasticmetal.GetBareMetalServerOption[];
    readonly organizationId: string;
    readonly os: string;
    readonly osName: string;
    readonly password: string;
    readonly privateNetworks: outputs.elasticmetal.GetBareMetalServerPrivateNetwork[];
    readonly projectId?: string;
    readonly reinstallOnConfigChanges: boolean;
    readonly serverId?: string;
    readonly servicePassword: string;
    readonly serviceUser: string;
    readonly sshKeyIds: string[];
    readonly tags: string[];
    readonly user: string;
    readonly zone?: string;
}
/**
 * Gets information about a baremetal server.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.elasticmetal.getBareMetalServer({
 *     name: "foobar",
 *     zone: "fr-par-2",
 * });
 * const byId = scaleway.elasticmetal.getBareMetalServer({
 *     serverId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getBareMetalServerOutput(args?: GetBareMetalServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBareMetalServerResult> {
    return pulumi.output(args).apply((a: any) => getBareMetalServer(a, opts))
}

/**
 * A collection of arguments for invoking getBareMetalServer.
 */
export interface GetBareMetalServerOutputArgs {
    /**
     * The server name. Only one of `name` and `serverId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the baremetal server is associated with.
     */
    projectId?: pulumi.Input<string>;
    serverId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the server exists.
     */
    zone?: pulumi.Input<string>;
}
