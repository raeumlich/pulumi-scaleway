// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const db = scaleway.documentdb.getInstance({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:documentdb/getInstance:getInstance", {
        "instanceId": args.instanceId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * The DocumentDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: string;
    /**
     * The name of the DocumentDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the DocumentDB instance is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the DocumentDB instance exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly engine: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly isHaCluster: boolean;
    readonly name?: string;
    readonly nodeType: string;
    readonly password: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly tags: string[];
    readonly telemetryEnabled: boolean;
    readonly userName: string;
    readonly volumeSizeInGb: number;
    readonly volumeType: string;
}
/**
 * Gets information about an DocumentDB instance. For further information see our [developers website](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const db = scaleway.documentdb.getInstance({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * The DocumentDB instance ID.
     * Only one of `name` and `instanceId` should be specified.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the DocumentDB instance.
     * Only one of `name` and `instanceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the DocumentDB instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the DocumentDB instance exists.
     */
    region?: pulumi.Input<string>;
}
