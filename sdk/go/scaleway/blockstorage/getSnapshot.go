// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Block Snapshot.
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/blockstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockstorage.LookupSnapshot(ctx, &blockstorage.LookupSnapshotArgs{
//				SnapshotId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
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
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("scaleway:blockstorage/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotArgs struct {
	// The name of the snapshot. Only one of `name` and `snapshotId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the snapshot is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The ID of the snapshot. Only one of `name` and `snapshotId` should be specified.
	SnapshotId *string `pulumi:"snapshotId"`
	// The ID of the volume from which the snapshot has been created.
	VolumeId *string `pulumi:"volumeId"`
	// `zone`) The zone in which the snapshot exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Name       *string  `pulumi:"name"`
	ProjectId  *string  `pulumi:"projectId"`
	SnapshotId *string  `pulumi:"snapshotId"`
	Tags       []string `pulumi:"tags"`
	VolumeId   *string  `pulumi:"volumeId"`
	Zone       *string  `pulumi:"zone"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotOutputArgs struct {
	// The name of the snapshot. Only one of `name` and `snapshotId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the snapshot is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The ID of the snapshot. Only one of `name` and `snapshotId` should be specified.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	// The ID of the volume from which the snapshot has been created.
	VolumeId pulumi.StringPtrInput `pulumi:"volumeId"`
	// `zone`) The zone in which the snapshot exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSnapshotResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupSnapshotResultOutput) VolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.VolumeId }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
