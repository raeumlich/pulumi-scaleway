// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about Scaleway Load-Balancer Certificates.
 *
 * This data source can prove useful when a module accepts an LB Certificate as an input variable and needs to, for example, determine the security of a certificate for your LB Frontend associated with your domain, etc.
 *
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).
 *
 * ## Examples
 */
export function getCertificate(args?: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:loadbalancer/getCertificate:getCertificate", {
        "certificateId": args.certificateId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * The certificate id.
     * - Only one of `name` and `certificateId` should be specified.
     */
    certificateId?: string;
    /**
     * The load-balancer ID this certificate is attached to.
     */
    lbId?: string;
    /**
     * The name of the certificate backend.
     * - When using a certificate `name` you should specify the `lb-id`
     */
    name?: string;
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    readonly certificateId?: string;
    readonly commonName: string;
    readonly customCertificates: outputs.loadbalancer.GetCertificateCustomCertificate[];
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbId?: string;
    readonly letsencrypts: outputs.loadbalancer.GetCertificateLetsencrypt[];
    readonly name?: string;
    readonly notValidAfter: string;
    readonly notValidBefore: string;
    readonly status: string;
    readonly subjectAlternativeNames: string[];
}
/**
 * Get information about Scaleway Load-Balancer Certificates.
 *
 * This data source can prove useful when a module accepts an LB Certificate as an input variable and needs to, for example, determine the security of a certificate for your LB Frontend associated with your domain, etc.
 *
 * For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).
 *
 * ## Examples
 */
export function getCertificateOutput(args?: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply((a: any) => getCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateOutputArgs {
    /**
     * The certificate id.
     * - Only one of `name` and `certificateId` should be specified.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The load-balancer ID this certificate is attached to.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The name of the certificate backend.
     * - When using a certificate `name` you should specify the `lb-id`
     */
    name?: pulumi.Input<string>;
}
