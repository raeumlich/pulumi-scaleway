// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway VPC Private Networks.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc/api/#private-networks-ac2df4).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewPrivateNetwork(ctx, "pnPriv", &vpc.PrivateNetworkArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("demo"),
//					pulumi.String("terraform"),
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
// ### With subnets
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewPrivateNetwork(ctx, "pnPriv", &vpc.PrivateNetworkArgs{
//				Ipv4Subnet: &vpc.PrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("192.168.0.0/24"),
//				},
//				Ipv6Subnets: vpc.PrivateNetworkIpv6SubnetArray{
//					&vpc.PrivateNetworkIpv6SubnetArgs{
//						Subnet: pulumi.String("fd46:78ab:30b8:177c::/64"),
//					},
//					&vpc.PrivateNetworkIpv6SubnetArgs{
//						Subnet: pulumi.String("fd46:78ab:30b8:c7df::/64"),
//					},
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("demo"),
//					pulumi.String("terraform"),
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
// Private networks can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:vpc/privateNetwork:PrivateNetwork vpc_demo fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type PrivateNetwork struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the subnet.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The IPv4 subnet to associate with the private network.
	Ipv4Subnet PrivateNetworkIpv4SubnetOutput `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets PrivateNetworkIpv6SubnetArrayOutput `pulumi:"ipv6Subnets"`
	// The private networks are necessarily regional now.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolOutput `pulumi:"isRegional"`
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the private network is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region of the private network.
	Region pulumi.StringOutput `pulumi:"region"`
	// The tags associated with the private network.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the subnet.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The VPC in which to create the private network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewPrivateNetwork registers a new resource with the given unique name, arguments, and options.
func NewPrivateNetwork(ctx *pulumi.Context,
	name string, args *PrivateNetworkArgs, opts ...pulumi.ResourceOption) (*PrivateNetwork, error) {
	if args == nil {
		args = &PrivateNetworkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateNetwork
	err := ctx.RegisterResource("scaleway:vpc/privateNetwork:PrivateNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateNetwork gets an existing PrivateNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateNetworkState, opts ...pulumi.ResourceOption) (*PrivateNetwork, error) {
	var resource PrivateNetwork
	err := ctx.ReadResource("scaleway:vpc/privateNetwork:PrivateNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateNetwork resources.
type privateNetworkState struct {
	// The date and time of the creation of the subnet.
	CreatedAt *string `pulumi:"createdAt"`
	// The IPv4 subnet to associate with the private network.
	Ipv4Subnet *PrivateNetworkIpv4Subnet `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets []PrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	// The private networks are necessarily regional now.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional *bool `pulumi:"isRegional"`
	// The name of the private network. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the private network is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the private network.
	Region *string `pulumi:"region"`
	// The tags associated with the private network.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the subnet.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The VPC in which to create the private network.
	VpcId *string `pulumi:"vpcId"`
	// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone *string `pulumi:"zone"`
}

type PrivateNetworkState struct {
	// The date and time of the creation of the subnet.
	CreatedAt pulumi.StringPtrInput
	// The IPv4 subnet to associate with the private network.
	Ipv4Subnet PrivateNetworkIpv4SubnetPtrInput
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets PrivateNetworkIpv6SubnetArrayInput
	// The private networks are necessarily regional now.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolPtrInput
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the private network is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the private network.
	Region pulumi.StringPtrInput
	// The tags associated with the private network.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the subnet.
	UpdatedAt pulumi.StringPtrInput
	// The VPC in which to create the private network.
	VpcId pulumi.StringPtrInput
	// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone pulumi.StringPtrInput
}

func (PrivateNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateNetworkState)(nil)).Elem()
}

type privateNetworkArgs struct {
	// The IPv4 subnet to associate with the private network.
	Ipv4Subnet *PrivateNetworkIpv4Subnet `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets []PrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	// The private networks are necessarily regional now.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional *bool `pulumi:"isRegional"`
	// The name of the private network. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the private network.
	Region *string `pulumi:"region"`
	// The tags associated with the private network.
	Tags []string `pulumi:"tags"`
	// The VPC in which to create the private network.
	VpcId *string `pulumi:"vpcId"`
	// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a PrivateNetwork resource.
type PrivateNetworkArgs struct {
	// The IPv4 subnet to associate with the private network.
	Ipv4Subnet PrivateNetworkIpv4SubnetPtrInput
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets PrivateNetworkIpv6SubnetArrayInput
	// The private networks are necessarily regional now.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolPtrInput
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the private network.
	Region pulumi.StringPtrInput
	// The tags associated with the private network.
	Tags pulumi.StringArrayInput
	// The VPC in which to create the private network.
	VpcId pulumi.StringPtrInput
	// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone pulumi.StringPtrInput
}

func (PrivateNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateNetworkArgs)(nil)).Elem()
}

type PrivateNetworkInput interface {
	pulumi.Input

	ToPrivateNetworkOutput() PrivateNetworkOutput
	ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput
}

func (*PrivateNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateNetwork)(nil)).Elem()
}

func (i *PrivateNetwork) ToPrivateNetworkOutput() PrivateNetworkOutput {
	return i.ToPrivateNetworkOutputWithContext(context.Background())
}

func (i *PrivateNetwork) ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkOutput)
}

// PrivateNetworkArrayInput is an input type that accepts PrivateNetworkArray and PrivateNetworkArrayOutput values.
// You can construct a concrete instance of `PrivateNetworkArrayInput` via:
//
//	PrivateNetworkArray{ PrivateNetworkArgs{...} }
type PrivateNetworkArrayInput interface {
	pulumi.Input

	ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput
	ToPrivateNetworkArrayOutputWithContext(context.Context) PrivateNetworkArrayOutput
}

type PrivateNetworkArray []PrivateNetworkInput

func (PrivateNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateNetwork)(nil)).Elem()
}

func (i PrivateNetworkArray) ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput {
	return i.ToPrivateNetworkArrayOutputWithContext(context.Background())
}

func (i PrivateNetworkArray) ToPrivateNetworkArrayOutputWithContext(ctx context.Context) PrivateNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkArrayOutput)
}

// PrivateNetworkMapInput is an input type that accepts PrivateNetworkMap and PrivateNetworkMapOutput values.
// You can construct a concrete instance of `PrivateNetworkMapInput` via:
//
//	PrivateNetworkMap{ "key": PrivateNetworkArgs{...} }
type PrivateNetworkMapInput interface {
	pulumi.Input

	ToPrivateNetworkMapOutput() PrivateNetworkMapOutput
	ToPrivateNetworkMapOutputWithContext(context.Context) PrivateNetworkMapOutput
}

type PrivateNetworkMap map[string]PrivateNetworkInput

func (PrivateNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateNetwork)(nil)).Elem()
}

func (i PrivateNetworkMap) ToPrivateNetworkMapOutput() PrivateNetworkMapOutput {
	return i.ToPrivateNetworkMapOutputWithContext(context.Background())
}

func (i PrivateNetworkMap) ToPrivateNetworkMapOutputWithContext(ctx context.Context) PrivateNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateNetworkMapOutput)
}

type PrivateNetworkOutput struct{ *pulumi.OutputState }

func (PrivateNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkOutput) ToPrivateNetworkOutput() PrivateNetworkOutput {
	return o
}

func (o PrivateNetworkOutput) ToPrivateNetworkOutputWithContext(ctx context.Context) PrivateNetworkOutput {
	return o
}

// The date and time of the creation of the subnet.
func (o PrivateNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The IPv4 subnet to associate with the private network.
func (o PrivateNetworkOutput) Ipv4Subnet() PrivateNetworkIpv4SubnetOutput {
	return o.ApplyT(func(v *PrivateNetwork) PrivateNetworkIpv4SubnetOutput { return v.Ipv4Subnet }).(PrivateNetworkIpv4SubnetOutput)
}

// The IPv6 subnets to associate with the private network.
func (o PrivateNetworkOutput) Ipv6Subnets() PrivateNetworkIpv6SubnetArrayOutput {
	return o.ApplyT(func(v *PrivateNetwork) PrivateNetworkIpv6SubnetArrayOutput { return v.Ipv6Subnets }).(PrivateNetworkIpv6SubnetArrayOutput)
}

// The private networks are necessarily regional now.
//
// Deprecated: This field is deprecated and will be removed in the next major version
func (o PrivateNetworkOutput) IsRegional() pulumi.BoolOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.BoolOutput { return v.IsRegional }).(pulumi.BoolOutput)
}

// The name of the private network. If not provided it will be randomly generated.
func (o PrivateNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the private network is associated with.
func (o PrivateNetworkOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the private network is associated with.
func (o PrivateNetworkOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region of the private network.
func (o PrivateNetworkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The tags associated with the private network.
func (o PrivateNetworkOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the subnet.
func (o PrivateNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The VPC in which to create the private network.
func (o PrivateNetworkOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// please use `region` instead - (Defaults to provider `zone`) The zone in which the private network should be created.
//
// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
func (o PrivateNetworkOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateNetwork) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type PrivateNetworkArrayOutput struct{ *pulumi.OutputState }

func (PrivateNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkArrayOutput) ToPrivateNetworkArrayOutput() PrivateNetworkArrayOutput {
	return o
}

func (o PrivateNetworkArrayOutput) ToPrivateNetworkArrayOutputWithContext(ctx context.Context) PrivateNetworkArrayOutput {
	return o
}

func (o PrivateNetworkArrayOutput) Index(i pulumi.IntInput) PrivateNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateNetwork {
		return vs[0].([]*PrivateNetwork)[vs[1].(int)]
	}).(PrivateNetworkOutput)
}

type PrivateNetworkMapOutput struct{ *pulumi.OutputState }

func (PrivateNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateNetwork)(nil)).Elem()
}

func (o PrivateNetworkMapOutput) ToPrivateNetworkMapOutput() PrivateNetworkMapOutput {
	return o
}

func (o PrivateNetworkMapOutput) ToPrivateNetworkMapOutputWithContext(ctx context.Context) PrivateNetworkMapOutput {
	return o
}

func (o PrivateNetworkMapOutput) MapIndex(k pulumi.StringInput) PrivateNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateNetwork {
		return vs[0].(map[string]*PrivateNetwork)[vs[1].(string)]
	}).(PrivateNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkInput)(nil)).Elem(), &PrivateNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkArrayInput)(nil)).Elem(), PrivateNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateNetworkMapInput)(nil)).Elem(), PrivateNetworkMap{})
	pulumi.RegisterOutputType(PrivateNetworkOutput{})
	pulumi.RegisterOutputType(PrivateNetworkArrayOutput{})
	pulumi.RegisterOutputType(PrivateNetworkMapOutput{})
}
