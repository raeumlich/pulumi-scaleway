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

// Creates and manages Scaleway DocumentDB database.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db).
//
// ## Example Usage
// ### Basic
//
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
//			_, err := documentdb.NewDatabase(ctx, "main", &documentdb.DatabaseArgs{
//				InstanceId: pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// DocumentDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:documentdb/database:Database mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
//
// ```
type Database struct {
	pulumi.CustomResourceState

	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed pulumi.BoolOutput `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the owner of the database.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Size in gigabytes of the database.
	Size pulumi.StringOutput `pulumi:"size"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("scaleway:documentdb/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("scaleway:documentdb/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId *string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed *bool `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner *string `pulumi:"owner"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Size in gigabytes of the database.
	Size *string `pulumi:"size"`
}

type DatabaseState struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringPtrInput
	// Whether the database is managed or not.
	Managed pulumi.BoolPtrInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The name of the owner of the database.
	Owner pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Size in gigabytes of the database.
	Size pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The project_id you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The project_id you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// UUID of the documentdb instance.
//
// > **Important:** Updates to `instanceId` will recreate the Database.
func (o DatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o DatabaseOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolOutput { return v.Managed }).(pulumi.BoolOutput)
}

// Name of the database (e.g. `my-new-database`).
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the owner of the database.
func (o DatabaseOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The project_id you want to attach the resource to
func (o DatabaseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o DatabaseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Size in gigabytes of the database.
func (o DatabaseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
