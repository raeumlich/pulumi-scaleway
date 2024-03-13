// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Compute Instance Volumes.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39).
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
//			_, err := instance.NewVolume(ctx, "serverVolume", &instance.VolumeArgs{
//				SizeInGb: pulumi.Int(20),
//				Type:     pulumi.String("l_ssd"),
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
// volumes can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:instance/volume:Volume server_volume fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type Volume struct {
	pulumi.CustomResourceState

	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	FromSnapshotId pulumi.StringPtrOutput `pulumi:"fromSnapshotId"`
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the volume is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The id of the associated server.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	SizeInGb pulumi.IntPtrOutput `pulumi:"sizeInGb"`
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
	Type pulumi.StringOutput `pulumi:"type"`
	// `zone`) The zone in which the volume should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("scaleway:instance/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("scaleway:instance/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	FromSnapshotId *string `pulumi:"fromSnapshotId"`
	// The name of the volume. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the volume is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The id of the associated server.
	ServerId *string `pulumi:"serverId"`
	// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	SizeInGb *int `pulumi:"sizeInGb"`
	// A list of tags to apply to the volume.
	Tags []string `pulumi:"tags"`
	// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
	Type *string `pulumi:"type"`
	// `zone`) The zone in which the volume should be created.
	Zone *string `pulumi:"zone"`
}

type VolumeState struct {
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	FromSnapshotId pulumi.StringPtrInput
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the volume is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringPtrInput
	// The id of the associated server.
	ServerId pulumi.StringPtrInput
	// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	SizeInGb pulumi.IntPtrInput
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayInput
	// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
	Type pulumi.StringPtrInput
	// `zone`) The zone in which the volume should be created.
	Zone pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	FromSnapshotId *string `pulumi:"fromSnapshotId"`
	// The name of the volume. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	SizeInGb *int `pulumi:"sizeInGb"`
	// A list of tags to apply to the volume.
	Tags []string `pulumi:"tags"`
	// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
	Type string `pulumi:"type"`
	// `zone`) The zone in which the volume should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	FromSnapshotId pulumi.StringPtrInput
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringPtrInput
	// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
	SizeInGb pulumi.IntPtrInput
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayInput
	// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
	Type pulumi.StringInput
	// `zone`) The zone in which the volume should be created.
	Zone pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

// VolumeArrayInput is an input type that accepts VolumeArray and VolumeArrayOutput values.
// You can construct a concrete instance of `VolumeArrayInput` via:
//
//	VolumeArray{ VolumeArgs{...} }
type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

// VolumeMapInput is an input type that accepts VolumeMap and VolumeMapOutput values.
// You can construct a concrete instance of `VolumeMapInput` via:
//
//	VolumeMap{ "key": VolumeArgs{...} }
type VolumeMapInput interface {
	pulumi.Input

	ToVolumeMapOutput() VolumeMapOutput
	ToVolumeMapOutputWithContext(context.Context) VolumeMapOutput
}

type VolumeMap map[string]VolumeInput

func (VolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (i VolumeMap) ToVolumeMapOutput() VolumeMapOutput {
	return i.ToVolumeMapOutputWithContext(context.Background())
}

func (i VolumeMap) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMapOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// If set, the new volume will be created from this snapshot. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
func (o VolumeOutput) FromSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.FromSnapshotId }).(pulumi.StringPtrOutput)
}

// The name of the volume. If not provided it will be randomly generated.
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the volume is associated with.
func (o VolumeOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the volume is associated with.
func (o VolumeOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The id of the associated server.
func (o VolumeOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The size of the volume. Only one of `sizeInGb` and `fromSnapshotId` should be specified.
func (o VolumeOutput) SizeInGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.SizeInGb }).(pulumi.IntPtrOutput)
}

// A list of tags to apply to the volume.
func (o VolumeOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the volume. The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD), `scratch` (Local Scratch SSD).
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// `zone`) The zone in which the volume should be created.
func (o VolumeOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].([]*Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMapOutput struct{ *pulumi.OutputState }

func (VolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (o VolumeMapOutput) ToVolumeMapOutput() VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) MapIndex(k pulumi.StringInput) VolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].(map[string]*Volume)[vs[1].(string)]
	}).(VolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), &Volume{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMapInput)(nil)).Elem(), VolumeMap{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMapOutput{})
}
