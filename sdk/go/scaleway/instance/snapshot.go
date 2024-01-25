// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Compute Snapshots.
// For more information,
// see [the documentation](https://developers.scaleway.com/en/products/instance/api/#snapshots-756fae).
//
// ## Example Usage
//
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
//			_, err := instance.NewSnapshot(ctx, "main", &instance.SnapshotArgs{
//				VolumeId: pulumi.String("11111111-1111-1111-1111-111111111111"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example with Unified type snapshot
//
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
//			mainVolume, err := instance.NewVolume(ctx, "mainVolume", &instance.VolumeArgs{
//				Type:     pulumi.String("l_ssd"),
//				SizeInGb: pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			mainServer, err := instance.NewServer(ctx, "mainServer", &instance.ServerArgs{
//				Image: pulumi.String("ubuntu_jammy"),
//				Type:  pulumi.String("DEV1-S"),
//				RootVolume: &instance.ServerRootVolumeArgs{
//					SizeInGb:   pulumi.Int(10),
//					VolumeType: pulumi.String("l_ssd"),
//				},
//				AdditionalVolumeIds: pulumi.StringArray{
//					mainVolume.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = instance.NewSnapshot(ctx, "mainSnapshot", &instance.SnapshotArgs{
//				VolumeId: mainVolume.ID(),
//				Type:     pulumi.String("unified"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				mainServer,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example importing a local qcow2 file
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/instance"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/objectstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucket, err := objectstorage.NewBucket(ctx, "bucket", nil)
//			if err != nil {
//				return err
//			}
//			qcow, err := objectstorage.NewItem(ctx, "qcow", &objectstorage.ItemArgs{
//				Bucket: bucket.Name,
//				Key:    pulumi.String("server.qcow2"),
//				File:   pulumi.String("myqcow.qcow2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = instance.NewSnapshot(ctx, "snapshot", &instance.SnapshotArgs{
//				Type: pulumi.String("unified"),
//				Import: &instance.SnapshotImportArgs{
//					Bucket: qcow.Bucket,
//					Key:    qcow.Key,
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
//
// ## Import
//
// Snapshots can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:instance/snapshot:Snapshot main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The snapshot creation time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Import a snapshot from a qcow2 file located in a bucket
	Import SnapshotImportPtrOutput `pulumi:"import"`
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the snapshot is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the snapshot is
	// associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// (Optional) The size of the snapshot.
	SizeInGb pulumi.IntOutput `pulumi:"sizeInGb"`
	// A list of tags to apply to the snapshot.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
	// Updates to this field will recreate a new resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringPtrOutput `pulumi:"volumeId"`
	// `zone`) The zone in which
	// the snapshot should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		args = &SnapshotArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("scaleway:instance/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("scaleway:instance/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The snapshot creation time.
	CreatedAt *string `pulumi:"createdAt"`
	// Import a snapshot from a qcow2 file located in a bucket
	Import *SnapshotImport `pulumi:"import"`
	// The name of the snapshot. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the snapshot is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the snapshot is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// (Optional) The size of the snapshot.
	SizeInGb *int `pulumi:"sizeInGb"`
	// A list of tags to apply to the snapshot.
	Tags []string `pulumi:"tags"`
	// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
	// Updates to this field will recreate a new resource.
	Type *string `pulumi:"type"`
	// The ID of the volume to take a snapshot from.
	VolumeId *string `pulumi:"volumeId"`
	// `zone`) The zone in which
	// the snapshot should be created.
	Zone *string `pulumi:"zone"`
}

type SnapshotState struct {
	// The snapshot creation time.
	CreatedAt pulumi.StringPtrInput
	// Import a snapshot from a qcow2 file located in a bucket
	Import SnapshotImportPtrInput
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the snapshot is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the snapshot is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// (Optional) The size of the snapshot.
	SizeInGb pulumi.IntPtrInput
	// A list of tags to apply to the snapshot.
	Tags pulumi.StringArrayInput
	// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
	// Updates to this field will recreate a new resource.
	Type pulumi.StringPtrInput
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringPtrInput
	// `zone`) The zone in which
	// the snapshot should be created.
	Zone pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Import a snapshot from a qcow2 file located in a bucket
	Import *SnapshotImport `pulumi:"import"`
	// The name of the snapshot. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the snapshot is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// A list of tags to apply to the snapshot.
	Tags []string `pulumi:"tags"`
	// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
	// Updates to this field will recreate a new resource.
	Type *string `pulumi:"type"`
	// The ID of the volume to take a snapshot from.
	VolumeId *string `pulumi:"volumeId"`
	// `zone`) The zone in which
	// the snapshot should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Import a snapshot from a qcow2 file located in a bucket
	Import SnapshotImportPtrInput
	// The name of the snapshot. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the snapshot is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// A list of tags to apply to the snapshot.
	Tags pulumi.StringArrayInput
	// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
	// Updates to this field will recreate a new resource.
	Type pulumi.StringPtrInput
	// The ID of the volume to take a snapshot from.
	VolumeId pulumi.StringPtrInput
	// `zone`) The zone in which
	// the snapshot should be created.
	Zone pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The snapshot creation time.
func (o SnapshotOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Import a snapshot from a qcow2 file located in a bucket
func (o SnapshotOutput) Import() SnapshotImportPtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotImportPtrOutput { return v.Import }).(SnapshotImportPtrOutput)
}

// The name of the snapshot. If not provided it will be randomly generated.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the snapshot is associated with.
func (o SnapshotOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the snapshot is
// associated with.
func (o SnapshotOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// (Optional) The size of the snapshot.
func (o SnapshotOutput) SizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.SizeInGb }).(pulumi.IntOutput)
}

// A list of tags to apply to the snapshot.
func (o SnapshotOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The snapshot's volume type.  The possible values are: `bSsd` (Block SSD), `lSsd` (Local SSD) and `unified`.
// Updates to this field will recreate a new resource.
func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The ID of the volume to take a snapshot from.
func (o SnapshotOutput) VolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.VolumeId }).(pulumi.StringPtrOutput)
}

// `zone`) The zone in which
// the snapshot should be created.
func (o SnapshotOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}