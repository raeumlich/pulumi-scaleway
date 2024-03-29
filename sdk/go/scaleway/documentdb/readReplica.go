// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway DocumentDB Database read replicas.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
//
// ## Example Usage
//
// ### Basic
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/documentdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := documentdb.NewReadReplica(ctx, "replica", &documentdb.ReadReplicaArgs{
//				DirectAccess: nil,
//				InstanceId:   pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// ### Private network
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/documentdb"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pn, err := vpc.NewPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = documentdb.NewReadReplica(ctx, "replica", &documentdb.ReadReplicaArgs{
//				InstanceId: pulumi.Any(scaleway_rdb_instance.Instance.Id),
//				PrivateNetwork: &documentdb.ReadReplicaPrivateNetworkArgs{
//					PrivateNetworkId: pn.ID(),
//					ServiceIp:        pulumi.String("192.168.1.254/24"),
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
// Database Read replica can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:documentdb/readReplica:ReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
// ```
type ReadReplica struct {
	pulumi.CustomResourceState

	// Creates a direct access endpoint to documentdb replica.
	DirectAccess ReadReplicaDirectAccessPtrOutput `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork ReadReplicaPrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewReadReplica registers a new resource with the given unique name, arguments, and options.
func NewReadReplica(ctx *pulumi.Context,
	name string, args *ReadReplicaArgs, opts ...pulumi.ResourceOption) (*ReadReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReadReplica
	err := ctx.RegisterResource("scaleway:documentdb/readReplica:ReadReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadReplica gets an existing ReadReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadReplicaState, opts ...pulumi.ResourceOption) (*ReadReplica, error) {
	var resource ReadReplica
	err := ctx.ReadResource("scaleway:documentdb/readReplica:ReadReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadReplica resources.
type readReplicaState struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess *ReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId *string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *ReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

type ReadReplicaState struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess ReadReplicaDirectAccessPtrInput
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringPtrInput
	// Create an endpoint in a private network.
	PrivateNetwork ReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (ReadReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*readReplicaState)(nil)).Elem()
}

type readReplicaArgs struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess *ReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *ReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ReadReplica resource.
type ReadReplicaArgs struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess ReadReplicaDirectAccessPtrInput
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringInput
	// Create an endpoint in a private network.
	PrivateNetwork ReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (ReadReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readReplicaArgs)(nil)).Elem()
}

type ReadReplicaInput interface {
	pulumi.Input

	ToReadReplicaOutput() ReadReplicaOutput
	ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput
}

func (*ReadReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadReplica)(nil)).Elem()
}

func (i *ReadReplica) ToReadReplicaOutput() ReadReplicaOutput {
	return i.ToReadReplicaOutputWithContext(context.Background())
}

func (i *ReadReplica) ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaOutput)
}

// ReadReplicaArrayInput is an input type that accepts ReadReplicaArray and ReadReplicaArrayOutput values.
// You can construct a concrete instance of `ReadReplicaArrayInput` via:
//
//	ReadReplicaArray{ ReadReplicaArgs{...} }
type ReadReplicaArrayInput interface {
	pulumi.Input

	ToReadReplicaArrayOutput() ReadReplicaArrayOutput
	ToReadReplicaArrayOutputWithContext(context.Context) ReadReplicaArrayOutput
}

type ReadReplicaArray []ReadReplicaInput

func (ReadReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadReplica)(nil)).Elem()
}

func (i ReadReplicaArray) ToReadReplicaArrayOutput() ReadReplicaArrayOutput {
	return i.ToReadReplicaArrayOutputWithContext(context.Background())
}

func (i ReadReplicaArray) ToReadReplicaArrayOutputWithContext(ctx context.Context) ReadReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaArrayOutput)
}

// ReadReplicaMapInput is an input type that accepts ReadReplicaMap and ReadReplicaMapOutput values.
// You can construct a concrete instance of `ReadReplicaMapInput` via:
//
//	ReadReplicaMap{ "key": ReadReplicaArgs{...} }
type ReadReplicaMapInput interface {
	pulumi.Input

	ToReadReplicaMapOutput() ReadReplicaMapOutput
	ToReadReplicaMapOutputWithContext(context.Context) ReadReplicaMapOutput
}

type ReadReplicaMap map[string]ReadReplicaInput

func (ReadReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadReplica)(nil)).Elem()
}

func (i ReadReplicaMap) ToReadReplicaMapOutput() ReadReplicaMapOutput {
	return i.ToReadReplicaMapOutputWithContext(context.Background())
}

func (i ReadReplicaMap) ToReadReplicaMapOutputWithContext(ctx context.Context) ReadReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadReplicaMapOutput)
}

type ReadReplicaOutput struct{ *pulumi.OutputState }

func (ReadReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadReplica)(nil)).Elem()
}

func (o ReadReplicaOutput) ToReadReplicaOutput() ReadReplicaOutput {
	return o
}

func (o ReadReplicaOutput) ToReadReplicaOutputWithContext(ctx context.Context) ReadReplicaOutput {
	return o
}

// Creates a direct access endpoint to documentdb replica.
func (o ReadReplicaOutput) DirectAccess() ReadReplicaDirectAccessPtrOutput {
	return o.ApplyT(func(v *ReadReplica) ReadReplicaDirectAccessPtrOutput { return v.DirectAccess }).(ReadReplicaDirectAccessPtrOutput)
}

// UUID of the documentdb instance.
//
// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
func (o ReadReplicaOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadReplica) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Create an endpoint in a private network.
func (o ReadReplicaOutput) PrivateNetwork() ReadReplicaPrivateNetworkPtrOutput {
	return o.ApplyT(func(v *ReadReplica) ReadReplicaPrivateNetworkPtrOutput { return v.PrivateNetwork }).(ReadReplicaPrivateNetworkPtrOutput)
}

// `region`) The region
// in which the Database read replica should be created.
func (o ReadReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ReadReplicaArrayOutput struct{ *pulumi.OutputState }

func (ReadReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadReplica)(nil)).Elem()
}

func (o ReadReplicaArrayOutput) ToReadReplicaArrayOutput() ReadReplicaArrayOutput {
	return o
}

func (o ReadReplicaArrayOutput) ToReadReplicaArrayOutputWithContext(ctx context.Context) ReadReplicaArrayOutput {
	return o
}

func (o ReadReplicaArrayOutput) Index(i pulumi.IntInput) ReadReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReadReplica {
		return vs[0].([]*ReadReplica)[vs[1].(int)]
	}).(ReadReplicaOutput)
}

type ReadReplicaMapOutput struct{ *pulumi.OutputState }

func (ReadReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadReplica)(nil)).Elem()
}

func (o ReadReplicaMapOutput) ToReadReplicaMapOutput() ReadReplicaMapOutput {
	return o
}

func (o ReadReplicaMapOutput) ToReadReplicaMapOutputWithContext(ctx context.Context) ReadReplicaMapOutput {
	return o
}

func (o ReadReplicaMapOutput) MapIndex(k pulumi.StringInput) ReadReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReadReplica {
		return vs[0].(map[string]*ReadReplica)[vs[1].(string)]
	}).(ReadReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaInput)(nil)).Elem(), &ReadReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaArrayInput)(nil)).Elem(), ReadReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadReplicaMapInput)(nil)).Elem(), ReadReplicaMap{})
	pulumi.RegisterOutputType(ReadReplicaOutput{})
	pulumi.RegisterOutputType(ReadReplicaArrayOutput{})
	pulumi.RegisterOutputType(ReadReplicaMapOutput{})
}
