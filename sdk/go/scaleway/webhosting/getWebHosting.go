// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webhosting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a webhosting.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/webhosting"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := webhosting.LookupWebHosting(ctx, &webhosting.LookupWebHostingArgs{
//				Domain: pulumi.StringRef("foobar.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = webhosting.LookupWebHosting(ctx, &webhosting.LookupWebHostingArgs{
//				WebhostingId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupWebHosting(ctx *pulumi.Context, args *LookupWebHostingArgs, opts ...pulumi.InvokeOption) (*LookupWebHostingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebHostingResult
	err := ctx.Invoke("scaleway:webhosting/getWebHosting:getWebHosting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebHosting.
type LookupWebHostingArgs struct {
	// The hosting domain name. Only one of `domain` and `webhostingId` should be specified.
	Domain *string `pulumi:"domain"`
	// The ID of the organization the hosting is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the hosting is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The hosting id. Only one of `domain` and `webhostingId` should be specified.
	WebhostingId *string `pulumi:"webhostingId"`
}

// A collection of values returned by getWebHosting.
type LookupWebHostingResult struct {
	CpanelUrls []GetWebHostingCpanelUrl `pulumi:"cpanelUrls"`
	CreatedAt  string                   `pulumi:"createdAt"`
	DnsStatus  string                   `pulumi:"dnsStatus"`
	Domain     *string                  `pulumi:"domain"`
	Email      string                   `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                `pulumi:"id"`
	OfferId          string                `pulumi:"offerId"`
	OfferName        string                `pulumi:"offerName"`
	OptionIds        []string              `pulumi:"optionIds"`
	Options          []GetWebHostingOption `pulumi:"options"`
	OrganizationId   string                `pulumi:"organizationId"`
	PlatformHostname string                `pulumi:"platformHostname"`
	PlatformNumber   int                   `pulumi:"platformNumber"`
	ProjectId        *string               `pulumi:"projectId"`
	Region           string                `pulumi:"region"`
	Status           string                `pulumi:"status"`
	Tags             []string              `pulumi:"tags"`
	UpdatedAt        string                `pulumi:"updatedAt"`
	Username         string                `pulumi:"username"`
	WebhostingId     *string               `pulumi:"webhostingId"`
}

func LookupWebHostingOutput(ctx *pulumi.Context, args LookupWebHostingOutputArgs, opts ...pulumi.InvokeOption) LookupWebHostingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebHostingResult, error) {
			args := v.(LookupWebHostingArgs)
			r, err := LookupWebHosting(ctx, &args, opts...)
			var s LookupWebHostingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebHostingResultOutput)
}

// A collection of arguments for invoking getWebHosting.
type LookupWebHostingOutputArgs struct {
	// The hosting domain name. Only one of `domain` and `webhostingId` should be specified.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// The ID of the organization the hosting is associated with.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the hosting is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The hosting id. Only one of `domain` and `webhostingId` should be specified.
	WebhostingId pulumi.StringPtrInput `pulumi:"webhostingId"`
}

func (LookupWebHostingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebHostingArgs)(nil)).Elem()
}

// A collection of values returned by getWebHosting.
type LookupWebHostingResultOutput struct{ *pulumi.OutputState }

func (LookupWebHostingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebHostingResult)(nil)).Elem()
}

func (o LookupWebHostingResultOutput) ToLookupWebHostingResultOutput() LookupWebHostingResultOutput {
	return o
}

func (o LookupWebHostingResultOutput) ToLookupWebHostingResultOutputWithContext(ctx context.Context) LookupWebHostingResultOutput {
	return o
}

func (o LookupWebHostingResultOutput) CpanelUrls() GetWebHostingCpanelUrlArrayOutput {
	return o.ApplyT(func(v LookupWebHostingResult) []GetWebHostingCpanelUrl { return v.CpanelUrls }).(GetWebHostingCpanelUrlArrayOutput)
}

func (o LookupWebHostingResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) DnsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.DnsStatus }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebHostingResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o LookupWebHostingResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWebHostingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.OfferId }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) OfferName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.OfferName }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) OptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebHostingResult) []string { return v.OptionIds }).(pulumi.StringArrayOutput)
}

func (o LookupWebHostingResultOutput) Options() GetWebHostingOptionArrayOutput {
	return o.ApplyT(func(v LookupWebHostingResult) []GetWebHostingOption { return v.Options }).(GetWebHostingOptionArrayOutput)
}

func (o LookupWebHostingResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) PlatformHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.PlatformHostname }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) PlatformNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWebHostingResult) int { return v.PlatformNumber }).(pulumi.IntOutput)
}

func (o LookupWebHostingResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebHostingResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupWebHostingResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebHostingResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupWebHostingResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebHostingResult) string { return v.Username }).(pulumi.StringOutput)
}

func (o LookupWebHostingResultOutput) WebhostingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebHostingResult) *string { return v.WebhostingId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebHostingResultOutput{})
}
