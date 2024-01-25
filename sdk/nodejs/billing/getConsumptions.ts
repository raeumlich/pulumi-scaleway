// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about your Consumptions.
 */
export function getConsumptions(opts?: pulumi.InvokeOptions): Promise<GetConsumptionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:billing/getConsumptions:getConsumptions", {
    }, opts);
}

/**
 * A collection of values returned by getConsumptions.
 */
export interface GetConsumptionsResult {
    /**
     * List of found consumptions
     */
    readonly consumptions: outputs.billing.GetConsumptionsConsumption[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationId: string;
    /**
     * The last consumption update date.
     */
    readonly updatedAt: string;
}
/**
 * Gets information about your Consumptions.
 */
export function getConsumptionsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetConsumptionsResult> {
    return pulumi.output(getConsumptions(opts))
}
