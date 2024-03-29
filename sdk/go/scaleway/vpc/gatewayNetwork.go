// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway VPC Public Gateway Network.
// It allows attaching Private Networks to the VPC Public Gateway and your DHCP config
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#step-3-attach-private-networks-to-the-vpc-public-gateway).
//
// ## Example Usage
//
// ### Create a gateway network with IPAM config
//
// <!--Start PulumiCodeChooser -->
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
//			vpc01, err := vpc.NewVPC(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", &vpc.PrivateNetworkArgs{
//				Ipv4Subnet: &vpc.PrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.64.0/22"),
//				},
//				VpcId: vpc01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := vpc.NewPublicGateway(ctx, "pg01", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewGatewayNetwork(ctx, "main", &vpc.GatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableMasquerade: pulumi.Bool(true),
//				IpamConfigs: vpc.GatewayNetworkIpamConfigArray{
//					&vpc.GatewayNetworkIpamConfigArgs{
//						PushDefaultRoute: pulumi.Bool(true),
//					},
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
// ### Create a gateway network with a booked IPAM IP
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/ipam"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpc01, err := vpc.NewVPC(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", &vpc.PrivateNetworkArgs{
//				Ipv4Subnet: &vpc.PrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.64.0/22"),
//				},
//				VpcId: vpc01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := ipam.NewIP(ctx, "ip01", &ipam.IPArgs{
//				Address: pulumi.String("172.16.64.7"),
//				Sources: ipam.IPSourceArray{
//					&ipam.IPSourceArgs{
//						PrivateNetworkId: pn01.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := vpc.NewPublicGateway(ctx, "pg01", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewGatewayNetwork(ctx, "main", &vpc.GatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableMasquerade: pulumi.Bool(true),
//				IpamConfigs: vpc.GatewayNetworkIpamConfigArray{
//					&vpc.GatewayNetworkIpamConfigArgs{
//						PushDefaultRoute: pulumi.Bool(true),
//						IpamIpId:         ip01.ID(),
//					},
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
// ### Create a gateway network with DHCP
//
// <!--Start PulumiCodeChooser -->
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
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			gw01, err := vpc.NewPublicGatewayIP(ctx, "gw01", nil)
//			if err != nil {
//				return err
//			}
//			dhcp01, err := vpc.NewPublicGatewayDHCP(ctx, "dhcp01", &vpc.PublicGatewayDHCPArgs{
//				Subnet:           pulumi.String("192.168.1.0/24"),
//				PushDefaultRoute: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := vpc.NewPublicGateway(ctx, "pg01", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: gw01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewGatewayNetwork(ctx, "main", &vpc.GatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				DhcpId:           dhcp01.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
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
// ### Create a gateway network with a static IP address
//
// <!--Start PulumiCodeChooser -->
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
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			pg01, err := vpc.NewPublicGateway(ctx, "pg01", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewGatewayNetwork(ctx, "main", &vpc.GatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableDhcp:       pulumi.Bool(false),
//				EnableMasquerade: pulumi.Bool(true),
//				StaticAddress:    pulumi.String("192.168.1.42/24"),
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
// Gateway network can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:vpc/gatewayNetwork:GatewayNetwork main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type GatewayNetwork struct {
	pulumi.CustomResourceState

	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrOutput `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrOutput `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrOutput `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrOutput `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs GatewayNetworkIpamConfigArrayOutput `pulumi:"ipamConfigs"`
	// The mac address of the creation of the gateway network.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The ID of the private network.
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringOutput `pulumi:"staticAddress"`
	// The status of the Public Gateway's connection to the Private Network.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date and time of the last update of the gateway network.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewGatewayNetwork registers a new resource with the given unique name, arguments, and options.
func NewGatewayNetwork(ctx *pulumi.Context,
	name string, args *GatewayNetworkArgs, opts ...pulumi.ResourceOption) (*GatewayNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GatewayNetwork
	err := ctx.RegisterResource("scaleway:vpc/gatewayNetwork:GatewayNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayNetwork gets an existing GatewayNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayNetworkState, opts ...pulumi.ResourceOption) (*GatewayNetwork, error) {
	var resource GatewayNetwork
	err := ctx.ReadResource("scaleway:vpc/gatewayNetwork:GatewayNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayNetwork resources.
type gatewayNetworkState struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs []GatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	// The mac address of the creation of the gateway network.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress *string `pulumi:"staticAddress"`
	// The status of the Public Gateway's connection to the Private Network.
	Status *string `pulumi:"status"`
	// The date and time of the last update of the gateway network.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the gateway network should be created.
	Zone *string `pulumi:"zone"`
}

type GatewayNetworkState struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrInput
	// The date and time of the creation of the gateway network.
	CreatedAt pulumi.StringPtrInput
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway.
	GatewayId pulumi.StringPtrInput
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs GatewayNetworkIpamConfigArrayInput
	// The mac address of the creation of the gateway network.
	MacAddress pulumi.StringPtrInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringPtrInput
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringPtrInput
	// The status of the Public Gateway's connection to the Private Network.
	Status pulumi.StringPtrInput
	// The date and time of the last update of the gateway network.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringPtrInput
}

func (GatewayNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayNetworkState)(nil)).Elem()
}

type gatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId string `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs []GatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	// The ID of the private network.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress *string `pulumi:"staticAddress"`
	// `zone`) The zone in which the gateway network should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a GatewayNetwork resource.
type GatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrInput
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway.
	GatewayId pulumi.StringInput
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs GatewayNetworkIpamConfigArrayInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringInput
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringPtrInput
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringPtrInput
}

func (GatewayNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayNetworkArgs)(nil)).Elem()
}

type GatewayNetworkInput interface {
	pulumi.Input

	ToGatewayNetworkOutput() GatewayNetworkOutput
	ToGatewayNetworkOutputWithContext(ctx context.Context) GatewayNetworkOutput
}

func (*GatewayNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayNetwork)(nil)).Elem()
}

func (i *GatewayNetwork) ToGatewayNetworkOutput() GatewayNetworkOutput {
	return i.ToGatewayNetworkOutputWithContext(context.Background())
}

func (i *GatewayNetwork) ToGatewayNetworkOutputWithContext(ctx context.Context) GatewayNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayNetworkOutput)
}

// GatewayNetworkArrayInput is an input type that accepts GatewayNetworkArray and GatewayNetworkArrayOutput values.
// You can construct a concrete instance of `GatewayNetworkArrayInput` via:
//
//	GatewayNetworkArray{ GatewayNetworkArgs{...} }
type GatewayNetworkArrayInput interface {
	pulumi.Input

	ToGatewayNetworkArrayOutput() GatewayNetworkArrayOutput
	ToGatewayNetworkArrayOutputWithContext(context.Context) GatewayNetworkArrayOutput
}

type GatewayNetworkArray []GatewayNetworkInput

func (GatewayNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayNetwork)(nil)).Elem()
}

func (i GatewayNetworkArray) ToGatewayNetworkArrayOutput() GatewayNetworkArrayOutput {
	return i.ToGatewayNetworkArrayOutputWithContext(context.Background())
}

func (i GatewayNetworkArray) ToGatewayNetworkArrayOutputWithContext(ctx context.Context) GatewayNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayNetworkArrayOutput)
}

// GatewayNetworkMapInput is an input type that accepts GatewayNetworkMap and GatewayNetworkMapOutput values.
// You can construct a concrete instance of `GatewayNetworkMapInput` via:
//
//	GatewayNetworkMap{ "key": GatewayNetworkArgs{...} }
type GatewayNetworkMapInput interface {
	pulumi.Input

	ToGatewayNetworkMapOutput() GatewayNetworkMapOutput
	ToGatewayNetworkMapOutputWithContext(context.Context) GatewayNetworkMapOutput
}

type GatewayNetworkMap map[string]GatewayNetworkInput

func (GatewayNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayNetwork)(nil)).Elem()
}

func (i GatewayNetworkMap) ToGatewayNetworkMapOutput() GatewayNetworkMapOutput {
	return i.ToGatewayNetworkMapOutputWithContext(context.Background())
}

func (i GatewayNetworkMap) ToGatewayNetworkMapOutputWithContext(ctx context.Context) GatewayNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayNetworkMapOutput)
}

type GatewayNetworkOutput struct{ *pulumi.OutputState }

func (GatewayNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayNetwork)(nil)).Elem()
}

func (o GatewayNetworkOutput) ToGatewayNetworkOutput() GatewayNetworkOutput {
	return o
}

func (o GatewayNetworkOutput) ToGatewayNetworkOutputWithContext(ctx context.Context) GatewayNetworkOutput {
	return o
}

// Remove DHCP config on this network on destroy. It requires DHCP id.
func (o GatewayNetworkOutput) CleanupDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.BoolPtrOutput { return v.CleanupDhcp }).(pulumi.BoolPtrOutput)
}

// The date and time of the creation of the gateway network.
func (o GatewayNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o GatewayNetworkOutput) DhcpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringPtrOutput { return v.DhcpId }).(pulumi.StringPtrOutput)
}

// Enable DHCP config on this network. It requires DHCP id.
func (o GatewayNetworkOutput) EnableDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.BoolPtrOutput { return v.EnableDhcp }).(pulumi.BoolPtrOutput)
}

// Enable masquerade on this network
func (o GatewayNetworkOutput) EnableMasquerade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.BoolPtrOutput { return v.EnableMasquerade }).(pulumi.BoolPtrOutput)
}

// The ID of the public gateway.
func (o GatewayNetworkOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o GatewayNetworkOutput) IpamConfigs() GatewayNetworkIpamConfigArrayOutput {
	return o.ApplyT(func(v *GatewayNetwork) GatewayNetworkIpamConfigArrayOutput { return v.IpamConfigs }).(GatewayNetworkIpamConfigArrayOutput)
}

// The mac address of the creation of the gateway network.
func (o GatewayNetworkOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The ID of the private network.
func (o GatewayNetworkOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o GatewayNetworkOutput) StaticAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.StaticAddress }).(pulumi.StringOutput)
}

// The status of the Public Gateway's connection to the Private Network.
func (o GatewayNetworkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date and time of the last update of the gateway network.
func (o GatewayNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the gateway network should be created.
func (o GatewayNetworkOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayNetwork) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type GatewayNetworkArrayOutput struct{ *pulumi.OutputState }

func (GatewayNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayNetwork)(nil)).Elem()
}

func (o GatewayNetworkArrayOutput) ToGatewayNetworkArrayOutput() GatewayNetworkArrayOutput {
	return o
}

func (o GatewayNetworkArrayOutput) ToGatewayNetworkArrayOutputWithContext(ctx context.Context) GatewayNetworkArrayOutput {
	return o
}

func (o GatewayNetworkArrayOutput) Index(i pulumi.IntInput) GatewayNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayNetwork {
		return vs[0].([]*GatewayNetwork)[vs[1].(int)]
	}).(GatewayNetworkOutput)
}

type GatewayNetworkMapOutput struct{ *pulumi.OutputState }

func (GatewayNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayNetwork)(nil)).Elem()
}

func (o GatewayNetworkMapOutput) ToGatewayNetworkMapOutput() GatewayNetworkMapOutput {
	return o
}

func (o GatewayNetworkMapOutput) ToGatewayNetworkMapOutputWithContext(ctx context.Context) GatewayNetworkMapOutput {
	return o
}

func (o GatewayNetworkMapOutput) MapIndex(k pulumi.StringInput) GatewayNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayNetwork {
		return vs[0].(map[string]*GatewayNetwork)[vs[1].(string)]
	}).(GatewayNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayNetworkInput)(nil)).Elem(), &GatewayNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayNetworkArrayInput)(nil)).Elem(), GatewayNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayNetworkMapInput)(nil)).Elem(), GatewayNetworkMap{})
	pulumi.RegisterOutputType(GatewayNetworkOutput{})
	pulumi.RegisterOutputType(GatewayNetworkArrayOutput{})
	pulumi.RegisterOutputType(GatewayNetworkMapOutput{})
}
