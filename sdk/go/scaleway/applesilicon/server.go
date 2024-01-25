// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applesilicon

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Apple silicon M1. For more information,
// see [the documentation](https://www.scaleway.com/en/docs/compute/apple-silicon/concepts).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/applesilicon"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := applesilicon.NewServer(ctx, "server", &applesilicon.ServerArgs{
//				Type: pulumi.String("M1-M"),
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
// Instance servers can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:applesilicon/server:Server main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type Server struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the Apple Silicon server.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringOutput `pulumi:"deletableAt"`
	// IPv4 address of the server (IPv4 address).
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The name of the server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The state of the server.
	State pulumi.StringOutput `pulumi:"state"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// URL of the VNC.
	VncUrl pulumi.StringOutput `pulumi:"vncUrl"`
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("scaleway:applesilicon/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("scaleway:applesilicon/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The date and time of the creation of the Apple Silicon server.
	CreatedAt *string `pulumi:"createdAt"`
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt *string `pulumi:"deletableAt"`
	// IPv4 address of the server (IPv4 address).
	Ip *string `pulumi:"ip"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The organization ID the server is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The state of the server.
	State *string `pulumi:"state"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type *string `pulumi:"type"`
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt *string `pulumi:"updatedAt"`
	// URL of the VNC.
	VncUrl *string `pulumi:"vncUrl"`
	// `zone`) The zone in which
	// the server should be created.
	Zone *string `pulumi:"zone"`
}

type ServerState struct {
	// The date and time of the creation of the Apple Silicon server.
	CreatedAt pulumi.StringPtrInput
	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt pulumi.StringPtrInput
	// IPv4 address of the server (IPv4 address).
	Ip pulumi.StringPtrInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The organization ID the server is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The state of the server.
	State pulumi.StringPtrInput
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringPtrInput
	// The date and time of the last update of the Apple Silicon server.
	UpdatedAt pulumi.StringPtrInput
	// URL of the VNC.
	VncUrl pulumi.StringPtrInput
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The name of the server.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId *string `pulumi:"projectId"`
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type string `pulumi:"type"`
	// `zone`) The zone in which
	// the server should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The name of the server.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the server is
	// associated with.
	ProjectId pulumi.StringPtrInput
	// The commercial type of the server. You find all the available types on
	// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
	// resource.
	Type pulumi.StringInput
	// `zone`) The zone in which
	// the server should be created.
	Zone pulumi.StringPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

// ServerArrayInput is an input type that accepts ServerArray and ServerArrayOutput values.
// You can construct a concrete instance of `ServerArrayInput` via:
//
//	ServerArray{ ServerArgs{...} }
type ServerArrayInput interface {
	pulumi.Input

	ToServerArrayOutput() ServerArrayOutput
	ToServerArrayOutputWithContext(context.Context) ServerArrayOutput
}

type ServerArray []ServerInput

func (ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (i ServerArray) ToServerArrayOutput() ServerArrayOutput {
	return i.ToServerArrayOutputWithContext(context.Background())
}

func (i ServerArray) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerArrayOutput)
}

// ServerMapInput is an input type that accepts ServerMap and ServerMapOutput values.
// You can construct a concrete instance of `ServerMapInput` via:
//
//	ServerMap{ "key": ServerArgs{...} }
type ServerMapInput interface {
	pulumi.Input

	ToServerMapOutput() ServerMapOutput
	ToServerMapOutputWithContext(context.Context) ServerMapOutput
}

type ServerMap map[string]ServerInput

func (ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (i ServerMap) ToServerMapOutput() ServerMapOutput {
	return i.ToServerMapOutputWithContext(context.Background())
}

func (i ServerMap) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerMapOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

// The date and time of the creation of the Apple Silicon server.
func (o ServerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The minimal date and time on which you can delete this server due to Apple licence
func (o ServerOutput) DeletableAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.DeletableAt }).(pulumi.StringOutput)
}

// IPv4 address of the server (IPv4 address).
func (o ServerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The name of the server.
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the server is associated with.
func (o ServerOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the server is
// associated with.
func (o ServerOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The state of the server.
func (o ServerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The commercial type of the server. You find all the available types on
// the [pricing page](https://www.scaleway.com/en/pricing/#apple-silicon). Updates to this field will recreate a new
// resource.
func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the Apple Silicon server.
func (o ServerOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// URL of the VNC.
func (o ServerOutput) VncUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.VncUrl }).(pulumi.StringOutput)
}

// `zone`) The zone in which
// the server should be created.
func (o ServerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ServerArrayOutput struct{ *pulumi.OutputState }

func (ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (o ServerArrayOutput) ToServerArrayOutput() ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) Index(i pulumi.IntInput) ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Server {
		return vs[0].([]*Server)[vs[1].(int)]
	}).(ServerOutput)
}

type ServerMapOutput struct{ *pulumi.OutputState }

func (ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (o ServerMapOutput) ToServerMapOutput() ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return o
}

func (o ServerMapOutput) MapIndex(k pulumi.StringInput) ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Server {
		return vs[0].(map[string]*Server)[vs[1].(string)]
	}).(ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerArrayInput)(nil)).Elem(), ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerMapInput)(nil)).Elem(), ServerMap{})
	pulumi.RegisterOutputType(ServerOutput{})
	pulumi.RegisterOutputType(ServerArrayOutput{})
	pulumi.RegisterOutputType(ServerMapOutput{})
}
