// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a function.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myFunction = scaleway.serverless.getFunction({
 *     functionId: "11111111-1111-1111-1111-111111111111",
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:serverless/getFunction:getFunction", {
        "functionId": args.functionId,
        "name": args.name,
        "namespaceId": args.namespaceId,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunction.
 */
export interface GetFunctionArgs {
    /**
     * The function id. Only one of `name` and `functionId` should be specified.
     */
    functionId?: string;
    /**
     * The function name. Only one of `name` and `namespaceId` should be specified.
     */
    name?: string;
    /**
     * The namespace id associated with this function.
     */
    namespaceId: string;
    /**
     * The ID of the project the function is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the function exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getFunction.
 */
export interface GetFunctionResult {
    readonly cpuLimit: number;
    readonly deploy: boolean;
    readonly description: string;
    readonly domainName: string;
    readonly environmentVariables: {[key: string]: string};
    readonly functionId?: string;
    readonly handler: string;
    readonly httpOption: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly maxScale: number;
    readonly memoryLimit: number;
    readonly minScale: number;
    readonly name?: string;
    readonly namespaceId: string;
    readonly organizationId: string;
    readonly privacy: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly runtime: string;
    readonly secretEnvironmentVariables: {[key: string]: string};
    readonly timeout: number;
    readonly zipFile: string;
    readonly zipHash: string;
}
/**
 * Gets information about a function.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myFunction = scaleway.serverless.getFunction({
 *     functionId: "11111111-1111-1111-1111-111111111111",
 *     namespaceId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getFunctionOutput(args: GetFunctionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionResult> {
    return pulumi.output(args).apply((a: any) => getFunction(a, opts))
}

/**
 * A collection of arguments for invoking getFunction.
 */
export interface GetFunctionOutputArgs {
    /**
     * The function id. Only one of `name` and `functionId` should be specified.
     */
    functionId?: pulumi.Input<string>;
    /**
     * The function name. Only one of `name` and `namespaceId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace id associated with this function.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * The ID of the project the function is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the function exists.
     */
    region?: pulumi.Input<string>;
}
