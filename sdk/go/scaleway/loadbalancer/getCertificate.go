// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Get information about Scaleway Load-Balancer Certificates.
//
// This data source can prove useful when a module accepts an LB Certificate as an input variable and needs to, for example, determine the security of a certificate for your LB Frontend associated with your domain, etc.
//
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-certificate).
//
// ## Examples
func GetCertificate(ctx *pulumi.Context, args *GetCertificateArgs, opts ...pulumi.InvokeOption) (*GetCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCertificateResult
	err := ctx.Invoke("scaleway:loadbalancer/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type GetCertificateArgs struct {
	// The certificate id.
	// - Only one of `name` and `certificateId` should be specified.
	CertificateId *string `pulumi:"certificateId"`
	// The load-balancer ID this certificate is attached to.
	LbId *string `pulumi:"lbId"`
	// The name of the certificate backend.
	// - When using a certificate `name` you should specify the `lb-id`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCertificate.
type GetCertificateResult struct {
	CertificateId      *string                           `pulumi:"certificateId"`
	CommonName         string                            `pulumi:"commonName"`
	CustomCertificates []GetCertificateCustomCertificate `pulumi:"customCertificates"`
	Fingerprint        string                            `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                      `pulumi:"id"`
	LbId                    *string                     `pulumi:"lbId"`
	Letsencrypts            []GetCertificateLetsencrypt `pulumi:"letsencrypts"`
	Name                    *string                     `pulumi:"name"`
	NotValidAfter           string                      `pulumi:"notValidAfter"`
	NotValidBefore          string                      `pulumi:"notValidBefore"`
	Status                  string                      `pulumi:"status"`
	SubjectAlternativeNames []string                    `pulumi:"subjectAlternativeNames"`
}

func GetCertificateOutput(ctx *pulumi.Context, args GetCertificateOutputArgs, opts ...pulumi.InvokeOption) GetCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCertificateResult, error) {
			args := v.(GetCertificateArgs)
			r, err := GetCertificate(ctx, &args, opts...)
			var s GetCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCertificateResultOutput)
}

// A collection of arguments for invoking getCertificate.
type GetCertificateOutputArgs struct {
	// The certificate id.
	// - Only one of `name` and `certificateId` should be specified.
	CertificateId pulumi.StringPtrInput `pulumi:"certificateId"`
	// The load-balancer ID this certificate is attached to.
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The name of the certificate backend.
	// - When using a certificate `name` you should specify the `lb-id`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getCertificate.
type GetCertificateResultOutput struct{ *pulumi.OutputState }

func (GetCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificateResult)(nil)).Elem()
}

func (o GetCertificateResultOutput) ToGetCertificateResultOutput() GetCertificateResultOutput {
	return o
}

func (o GetCertificateResultOutput) ToGetCertificateResultOutputWithContext(ctx context.Context) GetCertificateResultOutput {
	return o
}

func (o GetCertificateResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificateResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

func (o GetCertificateResultOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.CommonName }).(pulumi.StringOutput)
}

func (o GetCertificateResultOutput) CustomCertificates() GetCertificateCustomCertificateArrayOutput {
	return o.ApplyT(func(v GetCertificateResult) []GetCertificateCustomCertificate { return v.CustomCertificates }).(GetCertificateCustomCertificateArrayOutput)
}

func (o GetCertificateResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCertificateResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificateResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o GetCertificateResultOutput) Letsencrypts() GetCertificateLetsencryptArrayOutput {
	return o.ApplyT(func(v GetCertificateResult) []GetCertificateLetsencrypt { return v.Letsencrypts }).(GetCertificateLetsencryptArrayOutput)
}

func (o GetCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCertificateResultOutput) NotValidAfter() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.NotValidAfter }).(pulumi.StringOutput)
}

func (o GetCertificateResultOutput) NotValidBefore() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.NotValidBefore }).(pulumi.StringOutput)
}

func (o GetCertificateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetCertificateResultOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCertificateResult) []string { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificateResultOutput{})
}