// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Load-Balancers IPs.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/loadbalancer"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := loadbalancer.NewIP(ctx, "ip", &loadbalancer.IPArgs{
//				Reverse: pulumi.String("my-reverse.com"),
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
// IPs can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:loadbalancer/iP:IP ip01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type IP struct {
	pulumi.CustomResourceState

	// The IP Address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The associated load-balance ID if any
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region of the resource
	Region pulumi.StringOutput `pulumi:"region"`
	// The reverse domain associated with this IP.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIP registers a new resource with the given unique name, arguments, and options.
func NewIP(ctx *pulumi.Context,
	name string, args *IPArgs, opts ...pulumi.ResourceOption) (*IP, error) {
	if args == nil {
		args = &IPArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IP
	err := ctx.RegisterResource("scaleway:loadbalancer/iP:IP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIP gets an existing IP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPState, opts ...pulumi.ResourceOption) (*IP, error) {
	var resource IP
	err := ctx.ReadResource("scaleway:loadbalancer/iP:IP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IP resources.
type ipState struct {
	// The IP Address
	IpAddress *string `pulumi:"ipAddress"`
	// The associated load-balance ID if any
	LbId *string `pulumi:"lbId"`
	// The organization_id you want to attach the resource to
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The region of the resource
	Region *string `pulumi:"region"`
	// The reverse domain associated with this IP.
	Reverse *string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

type IPState struct {
	// The IP Address
	IpAddress pulumi.StringPtrInput
	// The associated load-balance ID if any
	LbId pulumi.StringPtrInput
	// The organization_id you want to attach the resource to
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The region of the resource
	Region pulumi.StringPtrInput
	// The reverse domain associated with this IP.
	Reverse pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (IPState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipState)(nil)).Elem()
}

type ipArgs struct {
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this IP.
	Reverse *string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a IP resource.
type IPArgs struct {
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this IP.
	Reverse pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (IPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipArgs)(nil)).Elem()
}

type IPInput interface {
	pulumi.Input

	ToIPOutput() IPOutput
	ToIPOutputWithContext(ctx context.Context) IPOutput
}

func (*IP) ElementType() reflect.Type {
	return reflect.TypeOf((**IP)(nil)).Elem()
}

func (i *IP) ToIPOutput() IPOutput {
	return i.ToIPOutputWithContext(context.Background())
}

func (i *IP) ToIPOutputWithContext(ctx context.Context) IPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPOutput)
}

// IPArrayInput is an input type that accepts IPArray and IPArrayOutput values.
// You can construct a concrete instance of `IPArrayInput` via:
//
//	IPArray{ IPArgs{...} }
type IPArrayInput interface {
	pulumi.Input

	ToIPArrayOutput() IPArrayOutput
	ToIPArrayOutputWithContext(context.Context) IPArrayOutput
}

type IPArray []IPInput

func (IPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IP)(nil)).Elem()
}

func (i IPArray) ToIPArrayOutput() IPArrayOutput {
	return i.ToIPArrayOutputWithContext(context.Background())
}

func (i IPArray) ToIPArrayOutputWithContext(ctx context.Context) IPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPArrayOutput)
}

// IPMapInput is an input type that accepts IPMap and IPMapOutput values.
// You can construct a concrete instance of `IPMapInput` via:
//
//	IPMap{ "key": IPArgs{...} }
type IPMapInput interface {
	pulumi.Input

	ToIPMapOutput() IPMapOutput
	ToIPMapOutputWithContext(context.Context) IPMapOutput
}

type IPMap map[string]IPInput

func (IPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IP)(nil)).Elem()
}

func (i IPMap) ToIPMapOutput() IPMapOutput {
	return i.ToIPMapOutputWithContext(context.Background())
}

func (i IPMap) ToIPMapOutputWithContext(ctx context.Context) IPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPMapOutput)
}

type IPOutput struct{ *pulumi.OutputState }

func (IPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IP)(nil)).Elem()
}

func (o IPOutput) ToIPOutput() IPOutput {
	return o
}

func (o IPOutput) ToIPOutputWithContext(ctx context.Context) IPOutput {
	return o
}

// The IP Address
func (o IPOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The associated load-balance ID if any
func (o IPOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// The organization_id you want to attach the resource to
func (o IPOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the IP is associated with.
func (o IPOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region of the resource
func (o IPOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The reverse domain associated with this IP.
func (o IPOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// `zone`) The zone in which the IP should be reserved.
func (o IPOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IPArrayOutput struct{ *pulumi.OutputState }

func (IPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IP)(nil)).Elem()
}

func (o IPArrayOutput) ToIPArrayOutput() IPArrayOutput {
	return o
}

func (o IPArrayOutput) ToIPArrayOutputWithContext(ctx context.Context) IPArrayOutput {
	return o
}

func (o IPArrayOutput) Index(i pulumi.IntInput) IPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IP {
		return vs[0].([]*IP)[vs[1].(int)]
	}).(IPOutput)
}

type IPMapOutput struct{ *pulumi.OutputState }

func (IPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IP)(nil)).Elem()
}

func (o IPMapOutput) ToIPMapOutput() IPMapOutput {
	return o
}

func (o IPMapOutput) ToIPMapOutputWithContext(ctx context.Context) IPMapOutput {
	return o
}

func (o IPMapOutput) MapIndex(k pulumi.StringInput) IPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IP {
		return vs[0].(map[string]*IP)[vs[1].(string)]
	}).(IPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPInput)(nil)).Elem(), &IP{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPArrayInput)(nil)).Elem(), IPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPMapInput)(nil)).Elem(), IPMap{})
	pulumi.RegisterOutputType(IPOutput{})
	pulumi.RegisterOutputType(IPArrayOutput{})
	pulumi.RegisterOutputType(IPMapOutput{})
}
