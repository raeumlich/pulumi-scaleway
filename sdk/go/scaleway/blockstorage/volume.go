// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Block Volumes.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/blockstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockstorage.NewVolume(ctx, "blockVolume", &blockstorage.VolumeArgs{
//				Iops:     pulumi.Int(5000),
//				SizeInGb: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Block Volumes can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:blockstorage/volume:Volume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type Volume struct {
	pulumi.CustomResourceState

	// The maximum IO/s expected, must match available options.
	Iops pulumi.IntOutput `pulumi:"iops"`
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
	SizeInGb pulumi.IntOutput `pulumi:"sizeInGb"`
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// `zone`) The zone in which the volume should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Iops == nil {
		return nil, errors.New("invalid value for required argument 'Iops'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("scaleway:blockstorage/volume:Volume", name, args, &resource, opts...)
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
	err := ctx.ReadResource("scaleway:blockstorage/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// The maximum IO/s expected, must match available options.
	Iops *int `pulumi:"iops"`
	// The name of the volume. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
	SizeInGb *int `pulumi:"sizeInGb"`
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of tags to apply to the volume.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the volume should be created.
	Zone *string `pulumi:"zone"`
}

type VolumeState struct {
	// The maximum IO/s expected, must match available options.
	Iops pulumi.IntPtrInput
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringPtrInput
	// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
	SizeInGb pulumi.IntPtrInput
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
	SnapshotId pulumi.StringPtrInput
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the volume should be created.
	Zone pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The maximum IO/s expected, must match available options.
	Iops int `pulumi:"iops"`
	// The name of the volume. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
	SizeInGb *int `pulumi:"sizeInGb"`
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of tags to apply to the volume.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the volume should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The maximum IO/s expected, must match available options.
	Iops pulumi.IntInput
	// The name of the volume. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the volume is associated with.
	ProjectId pulumi.StringPtrInput
	// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
	SizeInGb pulumi.IntPtrInput
	// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
	SnapshotId pulumi.StringPtrInput
	// A list of tags to apply to the volume.
	Tags pulumi.StringArrayInput
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

// The maximum IO/s expected, must match available options.
func (o VolumeOutput) Iops() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.Iops }).(pulumi.IntOutput)
}

// The name of the volume. If not provided it will be randomly generated.
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the volume is associated with.
func (o VolumeOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The size of the volume. Only one of `sizeInGb`, and `snapshotId` should be specified.
func (o VolumeOutput) SizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.SizeInGb }).(pulumi.IntOutput)
}

// If set, the new volume will be created from this snapshot. Only one of `sizeInGb`, `snapshotId` should be specified.
func (o VolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// A list of tags to apply to the volume.
func (o VolumeOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
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
