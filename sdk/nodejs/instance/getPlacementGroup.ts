// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets information about a Security Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.instance.getPlacementGroup({
 *     placementGroupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getPlacementGroup(args?: GetPlacementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPlacementGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:instance/getPlacementGroup:getPlacementGroup", {
        "name": args.name,
        "placementGroupId": args.placementGroupId,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlacementGroup.
 */
export interface GetPlacementGroupArgs {
    /**
     * The placement group name. Only one of `name` and `placementGroupId` should be specified.
     */
    name?: string;
    /**
     * The placement group id. Only one of `name` and `placementGroupId` should be specified.
     */
    placementGroupId?: string;
    /**
     * `projectId`) The ID of the project the placement group is associated with.
     */
    projectId?: string;
    /**
     * `zone`) The zone in which the placement group exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getPlacementGroup.
 */
export interface GetPlacementGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The organization ID the placement group is associated with.
     */
    readonly organizationId: string;
    readonly placementGroupId?: string;
    /**
     * The [policy mode](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group.
     */
    readonly policyMode: string;
    /**
     * Is true when the policy is respected.
     */
    readonly policyRespected: boolean;
    /**
     * The [policy type](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) of the placement group.
     */
    readonly policyType: string;
    readonly projectId: string;
    /**
     * A list of tags to apply to the placement group.
     */
    readonly tags: string[];
    readonly zone?: string;
}
/**
 * Gets information about a Security Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myKey = scaleway.instance.getPlacementGroup({
 *     placementGroupId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getPlacementGroupOutput(args?: GetPlacementGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlacementGroupResult> {
    return pulumi.output(args).apply((a: any) => getPlacementGroup(a, opts))
}

/**
 * A collection of arguments for invoking getPlacementGroup.
 */
export interface GetPlacementGroupOutputArgs {
    /**
     * The placement group name. Only one of `name` and `placementGroupId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The placement group id. Only one of `name` and `placementGroupId` should be specified.
     */
    placementGroupId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the placement group is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the placement group exists.
     */
    zone?: pulumi.Input<string>;
}
