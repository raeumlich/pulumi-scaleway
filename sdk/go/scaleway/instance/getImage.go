// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about an instance image.
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/instance"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := instance.LookupImage(ctx, &instance.LookupImageArgs{
//				ImageId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
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
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageResult
	err := ctx.Invoke("scaleway:instance/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture *string `pulumi:"architecture"`
	// The image id. Only one of `name` and `imageId` should be specified.
	ImageId *string `pulumi:"imageId"`
	// Use the latest image ID.
	Latest *bool `pulumi:"latest"`
	// The image name. Only one of `name` and `imageId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the image is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// IDs of the additional volumes in this image.
	AdditionalVolumeIds []string `pulumi:"additionalVolumeIds"`
	Architecture        *string  `pulumi:"architecture"`
	// Date of the image creation.
	CreationDate string `pulumi:"creationDate"`
	// ID of the default bootscript for this image.
	DefaultBootscriptId string `pulumi:"defaultBootscriptId"`
	// ID of the server the image if based from.
	FromServerId string `pulumi:"fromServerId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	ImageId *string `pulumi:"imageId"`
	Latest  *bool   `pulumi:"latest"`
	// Date of image latest update.
	ModificationDate string  `pulumi:"modificationDate"`
	Name             *string `pulumi:"name"`
	// The ID of the organization the image is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the project the image is associated with.
	ProjectId string `pulumi:"projectId"`
	// Set to `true` if the image is public.
	Public bool `pulumi:"public"`
	// ID of the root volume in this image.
	RootVolumeId string `pulumi:"rootVolumeId"`
	// State of the image. Possible values are: `available`, `creating` or `error`.
	State string `pulumi:"state"`
	Zone  string `pulumi:"zone"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

// A collection of arguments for invoking getImage.
type LookupImageOutputArgs struct {
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	// The image id. Only one of `name` and `imageId` should be specified.
	ImageId pulumi.StringPtrInput `pulumi:"imageId"`
	// Use the latest image ID.
	Latest pulumi.BoolPtrInput `pulumi:"latest"`
	// The image name. Only one of `name` and `imageId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the image is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

// A collection of values returned by getImage.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// IDs of the additional volumes in this image.
func (o LookupImageResultOutput) AdditionalVolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.AdditionalVolumeIds }).(pulumi.StringArrayOutput)
}

func (o LookupImageResultOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Architecture }).(pulumi.StringPtrOutput)
}

// Date of the image creation.
func (o LookupImageResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// ID of the default bootscript for this image.
func (o LookupImageResultOutput) DefaultBootscriptId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DefaultBootscriptId }).(pulumi.StringOutput)
}

// ID of the server the image if based from.
func (o LookupImageResultOutput) FromServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.FromServerId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

// Date of image latest update.
func (o LookupImageResultOutput) ModificationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ModificationDate }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the image is associated with.
func (o LookupImageResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the image is associated with.
func (o LookupImageResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Set to `true` if the image is public.
func (o LookupImageResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// ID of the root volume in this image.
func (o LookupImageResultOutput) RootVolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.RootVolumeId }).(pulumi.StringOutput)
}

// State of the image. Possible values are: `available`, `creating` or `error`.
func (o LookupImageResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
