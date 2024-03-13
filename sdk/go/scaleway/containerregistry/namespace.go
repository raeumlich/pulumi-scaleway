// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Container Registry.
// For more information see [the documentation](https://developers.scaleway.com/en/products/registry/api/).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/containerregistry"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containerregistry.NewNamespace(ctx, "main", &containerregistry.NamespaceArgs{
//				Description: pulumi.String("Main container registry"),
//				IsPublic:    pulumi.Bool(false),
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
// Namespaces can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:containerregistry/namespace:Namespace main fr-par/11111111-1111-1111-1111-111111111111
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// The description of the namespace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Endpoint reachable by Docker.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// The unique name of the namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		args = &NamespaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("scaleway:containerregistry/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("scaleway:containerregistry/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Endpoint reachable by Docker.
	Endpoint *string `pulumi:"endpoint"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic *bool `pulumi:"isPublic"`
	// The unique name of the namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name *string `pulumi:"name"`
	// The organization ID the namespace is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
}

type NamespaceState struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Endpoint reachable by Docker.
	Endpoint pulumi.StringPtrInput
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrInput
	// The unique name of the namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringPtrInput
	// The organization ID the namespace is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic *bool `pulumi:"isPublic"`
	// The unique name of the namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// Whether the images stored in the namespace should be downloadable publicly (docker pull).
	IsPublic pulumi.BoolPtrInput
	// The unique name of the namespace.
	//
	// > **Important** Updates to `name` will recreate the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the namespace is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace should be created.
	Region pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//	NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//	NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// The description of the namespace.
func (o NamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Endpoint reachable by Docker.
func (o NamespaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Whether the images stored in the namespace should be downloadable publicly (docker pull).
func (o NamespaceOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// The unique name of the namespace.
//
// > **Important** Updates to `name` will recreate the namespace.
func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the namespace is associated with.
func (o NamespaceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the namespace is associated with.
func (o NamespaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region in which the namespace should be created.
func (o NamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].([]*Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Namespace {
		return vs[0].(map[string]*Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceInput)(nil)).Elem(), &Namespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceArrayInput)(nil)).Elem(), NamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceMapInput)(nil)).Elem(), NamespaceMap{})
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
