// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myOffer = scaleway.elasticmetal.getBareMetalOffer({
 *     offerId: "25dcf38b-c90c-4b18-97a2-6956e9d1e113",
 *     zone: "fr-par-2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBareMetalOffer(args?: GetBareMetalOfferArgs, opts?: pulumi.InvokeOptions): Promise<GetBareMetalOfferResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:elasticmetal/getBareMetalOffer:getBareMetalOffer", {
        "includeDisabled": args.includeDisabled,
        "name": args.name,
        "offerId": args.offerId,
        "subscriptionPeriod": args.subscriptionPeriod,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getBareMetalOffer.
 */
export interface GetBareMetalOfferArgs {
    includeDisabled?: boolean;
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: string;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: string;
    /**
     * Period of subscription the desired offer. Should be `hourly` or `monthly`.
     */
    subscriptionPeriod?: string;
    /**
     * `zone`) The zone in which the offer should be created.
     */
    zone?: string;
}

/**
 * A collection of values returned by getBareMetalOffer.
 */
export interface GetBareMetalOfferResult {
    /**
     * Available Bandwidth with the offer.
     */
    readonly bandwidth: number;
    /**
     * Commercial range of the offer.
     */
    readonly commercialRange: string;
    /**
     * A list of cpu specifications. (Structure is documented below.)
     */
    readonly cpus: outputs.elasticmetal.GetBareMetalOfferCpus[];
    /**
     * A list of disk specifications. (Structure is documented below.)
     */
    readonly disks: outputs.elasticmetal.GetBareMetalOfferDisk[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeDisabled?: boolean;
    /**
     * A list of memory specifications. (Structure is documented below.)
     */
    readonly memories: outputs.elasticmetal.GetBareMetalOfferMemory[];
    /**
     * Name of the CPU.
     */
    readonly name?: string;
    readonly offerId?: string;
    /**
     * Stock status for this offer. Possible values are: `empty`, `low` or `available`.
     */
    readonly stock: string;
    readonly subscriptionPeriod?: string;
    readonly zone: string;
}
/**
 * Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myOffer = scaleway.elasticmetal.getBareMetalOffer({
 *     offerId: "25dcf38b-c90c-4b18-97a2-6956e9d1e113",
 *     zone: "fr-par-2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBareMetalOfferOutput(args?: GetBareMetalOfferOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBareMetalOfferResult> {
    return pulumi.output(args).apply((a: any) => getBareMetalOffer(a, opts))
}

/**
 * A collection of arguments for invoking getBareMetalOffer.
 */
export interface GetBareMetalOfferOutputArgs {
    includeDisabled?: pulumi.Input<boolean>;
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: pulumi.Input<string>;
    /**
     * Period of subscription the desired offer. Should be `hourly` or `monthly`.
     */
    subscriptionPeriod?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the offer should be created.
     */
    zone?: pulumi.Input<string>;
}
