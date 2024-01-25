// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway VPC Public Gateway IP.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/dns"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := vpc.NewPublicGatewayIP(ctx, "main", &vpc.PublicGatewayIPArgs{
//				Reverse: pulumi.String("tf.example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewRecord(ctx, "tfA", &dns.RecordArgs{
//				Data:     main.Address,
//				DnsZone:  pulumi.String("example.com"),
//				Priority: pulumi.Int(1),
//				Ttl:      pulumi.Int(3600),
//				Type:     pulumi.String("A"),
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
// Public gateway can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:vpc/publicGatewayIP:PublicGatewayIP main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type PublicGatewayIP struct {
	pulumi.CustomResourceState

	// The IP address itself.
	Address pulumi.StringOutput `pulumi:"address"`
	// The date and time of the creation of the public gateway ip.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The organization ID the public gateway ip is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the public gateway ip.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewPublicGatewayIP registers a new resource with the given unique name, arguments, and options.
func NewPublicGatewayIP(ctx *pulumi.Context,
	name string, args *PublicGatewayIPArgs, opts ...pulumi.ResourceOption) (*PublicGatewayIP, error) {
	if args == nil {
		args = &PublicGatewayIPArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicGatewayIP
	err := ctx.RegisterResource("scaleway:vpc/publicGatewayIP:PublicGatewayIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicGatewayIP gets an existing PublicGatewayIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicGatewayIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicGatewayIPState, opts ...pulumi.ResourceOption) (*PublicGatewayIP, error) {
	var resource PublicGatewayIP
	err := ctx.ReadResource("scaleway:vpc/publicGatewayIP:PublicGatewayIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicGatewayIP resources.
type publicGatewayIPState struct {
	// The IP address itself.
	Address *string `pulumi:"address"`
	// The date and time of the creation of the public gateway ip.
	CreatedAt *string `pulumi:"createdAt"`
	// The organization ID the public gateway ip is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the public gateway ip.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone *string `pulumi:"zone"`
}

type PublicGatewayIPState struct {
	// The IP address itself.
	Address pulumi.StringPtrInput
	// The date and time of the creation of the public gateway ip.
	CreatedAt pulumi.StringPtrInput
	// The organization ID the public gateway ip is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the public gateway ip.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringPtrInput
}

func (PublicGatewayIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicGatewayIPState)(nil)).Elem()
}

type publicGatewayIPArgs struct {
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain name for the IP address
	Reverse *string `pulumi:"reverse"`
	// The tags associated with the public gateway IP.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the public gateway ip should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a PublicGatewayIP resource.
type PublicGatewayIPArgs struct {
	// `projectId`) The ID of the project the public gateway ip is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse domain name for the IP address
	Reverse pulumi.StringPtrInput
	// The tags associated with the public gateway IP.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the public gateway ip should be created.
	Zone pulumi.StringPtrInput
}

func (PublicGatewayIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicGatewayIPArgs)(nil)).Elem()
}

type PublicGatewayIPInput interface {
	pulumi.Input

	ToPublicGatewayIPOutput() PublicGatewayIPOutput
	ToPublicGatewayIPOutputWithContext(ctx context.Context) PublicGatewayIPOutput
}

func (*PublicGatewayIP) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicGatewayIP)(nil)).Elem()
}

func (i *PublicGatewayIP) ToPublicGatewayIPOutput() PublicGatewayIPOutput {
	return i.ToPublicGatewayIPOutputWithContext(context.Background())
}

func (i *PublicGatewayIP) ToPublicGatewayIPOutputWithContext(ctx context.Context) PublicGatewayIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIPOutput)
}

// PublicGatewayIPArrayInput is an input type that accepts PublicGatewayIPArray and PublicGatewayIPArrayOutput values.
// You can construct a concrete instance of `PublicGatewayIPArrayInput` via:
//
//	PublicGatewayIPArray{ PublicGatewayIPArgs{...} }
type PublicGatewayIPArrayInput interface {
	pulumi.Input

	ToPublicGatewayIPArrayOutput() PublicGatewayIPArrayOutput
	ToPublicGatewayIPArrayOutputWithContext(context.Context) PublicGatewayIPArrayOutput
}

type PublicGatewayIPArray []PublicGatewayIPInput

func (PublicGatewayIPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicGatewayIP)(nil)).Elem()
}

func (i PublicGatewayIPArray) ToPublicGatewayIPArrayOutput() PublicGatewayIPArrayOutput {
	return i.ToPublicGatewayIPArrayOutputWithContext(context.Background())
}

func (i PublicGatewayIPArray) ToPublicGatewayIPArrayOutputWithContext(ctx context.Context) PublicGatewayIPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIPArrayOutput)
}

// PublicGatewayIPMapInput is an input type that accepts PublicGatewayIPMap and PublicGatewayIPMapOutput values.
// You can construct a concrete instance of `PublicGatewayIPMapInput` via:
//
//	PublicGatewayIPMap{ "key": PublicGatewayIPArgs{...} }
type PublicGatewayIPMapInput interface {
	pulumi.Input

	ToPublicGatewayIPMapOutput() PublicGatewayIPMapOutput
	ToPublicGatewayIPMapOutputWithContext(context.Context) PublicGatewayIPMapOutput
}

type PublicGatewayIPMap map[string]PublicGatewayIPInput

func (PublicGatewayIPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicGatewayIP)(nil)).Elem()
}

func (i PublicGatewayIPMap) ToPublicGatewayIPMapOutput() PublicGatewayIPMapOutput {
	return i.ToPublicGatewayIPMapOutputWithContext(context.Background())
}

func (i PublicGatewayIPMap) ToPublicGatewayIPMapOutputWithContext(ctx context.Context) PublicGatewayIPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicGatewayIPMapOutput)
}

type PublicGatewayIPOutput struct{ *pulumi.OutputState }

func (PublicGatewayIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicGatewayIP)(nil)).Elem()
}

func (o PublicGatewayIPOutput) ToPublicGatewayIPOutput() PublicGatewayIPOutput {
	return o
}

func (o PublicGatewayIPOutput) ToPublicGatewayIPOutputWithContext(ctx context.Context) PublicGatewayIPOutput {
	return o
}

// The IP address itself.
func (o PublicGatewayIPOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The date and time of the creation of the public gateway ip.
func (o PublicGatewayIPOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The organization ID the public gateway ip is associated with.
func (o PublicGatewayIPOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the public gateway ip is associated with.
func (o PublicGatewayIPOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain name for the IP address
func (o PublicGatewayIPOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The tags associated with the public gateway IP.
func (o PublicGatewayIPOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the public gateway ip.
func (o PublicGatewayIPOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the public gateway ip should be created.
func (o PublicGatewayIPOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicGatewayIP) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type PublicGatewayIPArrayOutput struct{ *pulumi.OutputState }

func (PublicGatewayIPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicGatewayIP)(nil)).Elem()
}

func (o PublicGatewayIPArrayOutput) ToPublicGatewayIPArrayOutput() PublicGatewayIPArrayOutput {
	return o
}

func (o PublicGatewayIPArrayOutput) ToPublicGatewayIPArrayOutputWithContext(ctx context.Context) PublicGatewayIPArrayOutput {
	return o
}

func (o PublicGatewayIPArrayOutput) Index(i pulumi.IntInput) PublicGatewayIPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicGatewayIP {
		return vs[0].([]*PublicGatewayIP)[vs[1].(int)]
	}).(PublicGatewayIPOutput)
}

type PublicGatewayIPMapOutput struct{ *pulumi.OutputState }

func (PublicGatewayIPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicGatewayIP)(nil)).Elem()
}

func (o PublicGatewayIPMapOutput) ToPublicGatewayIPMapOutput() PublicGatewayIPMapOutput {
	return o
}

func (o PublicGatewayIPMapOutput) ToPublicGatewayIPMapOutputWithContext(ctx context.Context) PublicGatewayIPMapOutput {
	return o
}

func (o PublicGatewayIPMapOutput) MapIndex(k pulumi.StringInput) PublicGatewayIPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicGatewayIP {
		return vs[0].(map[string]*PublicGatewayIP)[vs[1].(string)]
	}).(PublicGatewayIPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIPInput)(nil)).Elem(), &PublicGatewayIP{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIPArrayInput)(nil)).Elem(), PublicGatewayIPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicGatewayIPMapInput)(nil)).Elem(), PublicGatewayIPMap{})
	pulumi.RegisterOutputType(PublicGatewayIPOutput{})
	pulumi.RegisterOutputType(PublicGatewayIPArrayOutput{})
	pulumi.RegisterOutputType(PublicGatewayIPMapOutput{})
}