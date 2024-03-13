// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Gets local image ID of an image from its label name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.instance.getMarketplaceImage({
 *     label: "ubuntu_jammy",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMarketplaceImage(args: GetMarketplaceImageArgs, opts?: pulumi.InvokeOptions): Promise<GetMarketplaceImageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:instance/getMarketplaceImage:getMarketplaceImage", {
        "instanceType": args.instanceType,
        "label": args.label,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getMarketplaceImage.
 */
export interface GetMarketplaceImageArgs {
    /**
     * The instance type the image is compatible with.
     * You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
     */
    instanceType?: string;
    /**
     * Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&per_page=100)
     * to find the right `label`.
     */
    label: string;
    /**
     * `zone`) The zone in which the image exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getMarketplaceImage.
 */
export interface GetMarketplaceImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceType?: string;
    readonly label: string;
    readonly zone: string;
}
/**
 * Gets local image ID of an image from its label name.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myImage = scaleway.instance.getMarketplaceImage({
 *     label: "ubuntu_jammy",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMarketplaceImageOutput(args: GetMarketplaceImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMarketplaceImageResult> {
    return pulumi.output(args).apply((a: any) => getMarketplaceImage(a, opts))
}

/**
 * A collection of arguments for invoking getMarketplaceImage.
 */
export interface GetMarketplaceImageOutputArgs {
    /**
     * The instance type the image is compatible with.
     * You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Exact label of the desired image. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&per_page=100)
     * to find the right `label`.
     */
    label: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the image exists.
     */
    zone?: pulumi.Input<string>;
}
