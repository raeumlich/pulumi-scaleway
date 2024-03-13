// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Virtual Private Clouds.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewVPC(ctx, "vpc01", &vpc.VPCArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("demo"),
//					pulumi.String("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// VPCs can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:vpc/vPC:VPC vpc_demo fr-par/11111111-1111-1111-1111-111111111111
// ```
type VPC struct {
	pulumi.CustomResourceState

	// Date and time of VPC's creation (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Defines whether the VPC is the default one for its Project.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The name of the VPC. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the VPC is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region of the VPC.
	Region pulumi.StringOutput `pulumi:"region"`
	// The tags associated with the VPC.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time of VPC's last update (RFC 3339 format).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewVPC registers a new resource with the given unique name, arguments, and options.
func NewVPC(ctx *pulumi.Context,
	name string, args *VPCArgs, opts ...pulumi.ResourceOption) (*VPC, error) {
	if args == nil {
		args = &VPCArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VPC
	err := ctx.RegisterResource("scaleway:vpc/vPC:VPC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPC gets an existing VPC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCState, opts ...pulumi.ResourceOption) (*VPC, error) {
	var resource VPC
	err := ctx.ReadResource("scaleway:vpc/vPC:VPC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPC resources.
type vpcState struct {
	// Date and time of VPC's creation (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// Defines whether the VPC is the default one for its Project.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the VPC. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the VPC is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the VPC.
	Region *string `pulumi:"region"`
	// The tags associated with the VPC.
	Tags []string `pulumi:"tags"`
	// Date and time of VPC's last update (RFC 3339 format).
	UpdatedAt *string `pulumi:"updatedAt"`
}

type VPCState struct {
	// Date and time of VPC's creation (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// Defines whether the VPC is the default one for its Project.
	IsDefault pulumi.BoolPtrInput
	// The name of the VPC. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the VPC is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the VPC.
	Region pulumi.StringPtrInput
	// The tags associated with the VPC.
	Tags pulumi.StringArrayInput
	// Date and time of VPC's last update (RFC 3339 format).
	UpdatedAt pulumi.StringPtrInput
}

func (VPCState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// The name of the VPC. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the VPC.
	Region *string `pulumi:"region"`
	// The tags associated with the VPC.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a VPC resource.
type VPCArgs struct {
	// The name of the VPC. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the VPC.
	Region pulumi.StringPtrInput
	// The tags associated with the VPC.
	Tags pulumi.StringArrayInput
}

func (VPCArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VPCInput interface {
	pulumi.Input

	ToVPCOutput() VPCOutput
	ToVPCOutputWithContext(ctx context.Context) VPCOutput
}

func (*VPC) ElementType() reflect.Type {
	return reflect.TypeOf((**VPC)(nil)).Elem()
}

func (i *VPC) ToVPCOutput() VPCOutput {
	return i.ToVPCOutputWithContext(context.Background())
}

func (i *VPC) ToVPCOutputWithContext(ctx context.Context) VPCOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCOutput)
}

// VPCArrayInput is an input type that accepts VPCArray and VPCArrayOutput values.
// You can construct a concrete instance of `VPCArrayInput` via:
//
//	VPCArray{ VPCArgs{...} }
type VPCArrayInput interface {
	pulumi.Input

	ToVPCArrayOutput() VPCArrayOutput
	ToVPCArrayOutputWithContext(context.Context) VPCArrayOutput
}

type VPCArray []VPCInput

func (VPCArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VPC)(nil)).Elem()
}

func (i VPCArray) ToVPCArrayOutput() VPCArrayOutput {
	return i.ToVPCArrayOutputWithContext(context.Background())
}

func (i VPCArray) ToVPCArrayOutputWithContext(ctx context.Context) VPCArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCArrayOutput)
}

// VPCMapInput is an input type that accepts VPCMap and VPCMapOutput values.
// You can construct a concrete instance of `VPCMapInput` via:
//
//	VPCMap{ "key": VPCArgs{...} }
type VPCMapInput interface {
	pulumi.Input

	ToVPCMapOutput() VPCMapOutput
	ToVPCMapOutputWithContext(context.Context) VPCMapOutput
}

type VPCMap map[string]VPCInput

func (VPCMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VPC)(nil)).Elem()
}

func (i VPCMap) ToVPCMapOutput() VPCMapOutput {
	return i.ToVPCMapOutputWithContext(context.Background())
}

func (i VPCMap) ToVPCMapOutputWithContext(ctx context.Context) VPCMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCMapOutput)
}

type VPCOutput struct{ *pulumi.OutputState }

func (VPCOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPC)(nil)).Elem()
}

func (o VPCOutput) ToVPCOutput() VPCOutput {
	return o
}

func (o VPCOutput) ToVPCOutputWithContext(ctx context.Context) VPCOutput {
	return o
}

// Date and time of VPC's creation (RFC 3339 format).
func (o VPCOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Defines whether the VPC is the default one for its Project.
func (o VPCOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *VPC) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The name of the VPC. If not provided it will be randomly generated.
func (o VPCOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the VPC is associated with.
func (o VPCOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the VPC is associated with.
func (o VPCOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region of the VPC.
func (o VPCOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The tags associated with the VPC.
func (o VPCOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Date and time of VPC's last update (RFC 3339 format).
func (o VPCOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VPC) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type VPCArrayOutput struct{ *pulumi.OutputState }

func (VPCArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VPC)(nil)).Elem()
}

func (o VPCArrayOutput) ToVPCArrayOutput() VPCArrayOutput {
	return o
}

func (o VPCArrayOutput) ToVPCArrayOutputWithContext(ctx context.Context) VPCArrayOutput {
	return o
}

func (o VPCArrayOutput) Index(i pulumi.IntInput) VPCOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VPC {
		return vs[0].([]*VPC)[vs[1].(int)]
	}).(VPCOutput)
}

type VPCMapOutput struct{ *pulumi.OutputState }

func (VPCMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VPC)(nil)).Elem()
}

func (o VPCMapOutput) ToVPCMapOutput() VPCMapOutput {
	return o
}

func (o VPCMapOutput) ToVPCMapOutputWithContext(ctx context.Context) VPCMapOutput {
	return o
}

func (o VPCMapOutput) MapIndex(k pulumi.StringInput) VPCOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VPC {
		return vs[0].(map[string]*VPC)[vs[1].(string)]
	}).(VPCOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPCInput)(nil)).Elem(), &VPC{})
	pulumi.RegisterInputType(reflect.TypeOf((*VPCArrayInput)(nil)).Elem(), VPCArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VPCMapInput)(nil)).Elem(), VPCMap{})
	pulumi.RegisterOutputType(VPCOutput{})
	pulumi.RegisterOutputType(VPCArrayOutput{})
	pulumi.RegisterOutputType(VPCMapOutput{})
}
