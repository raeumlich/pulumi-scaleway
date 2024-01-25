// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a gateway network.
//
// ## Example Usage
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
//			main, err := vpc.NewGatewayNetwork(ctx, "main", &vpc.GatewayNetworkArgs{
//				GatewayId:        pulumi.Any(scaleway_vpc_public_gateway.Pg01.Id),
//				PrivateNetworkId: pulumi.Any(scaleway_vpc_private_network.Pn01.Id),
//				DhcpId:           pulumi.Any(scaleway_vpc_public_gateway_dhcp.Dhcp01.Id),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.LookupGatewayNetworkOutput(ctx, vpc.GetGatewayNetworkOutputArgs{
//				GatewayNetworkId: main.ID(),
//			}, nil)
//			_, err = vpc.LookupGatewayNetwork(ctx, &vpc.LookupGatewayNetworkArgs{
//				GatewayId:        pulumi.StringRef(scaleway_vpc_public_gateway.Pg01.Id),
//				PrivateNetworkId: pulumi.StringRef(scaleway_vpc_private_network.Pn01.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGatewayNetwork(ctx *pulumi.Context, args *LookupGatewayNetworkArgs, opts ...pulumi.InvokeOption) (*LookupGatewayNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayNetworkResult
	err := ctx.Invoke("scaleway:vpc/getGatewayNetwork:getGatewayNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayNetwork.
type LookupGatewayNetworkArgs struct {
	// ID of the public gateway DHCP config
	DhcpId *string `pulumi:"dhcpId"`
	// If masquerade is enabled on requested network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// ID of the public gateway the gateway network is linked to
	GatewayId *string `pulumi:"gatewayId"`
	// ID of the gateway network.
	//
	// > Only one of `gatewayNetworkId` or filters should be specified. You can use all the filters you want.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// ID of the private network the gateway network is linked to
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
}

// A collection of values returned by getGatewayNetwork.
type LookupGatewayNetworkResult struct {
	CleanupDhcp      bool    `pulumi:"cleanupDhcp"`
	CreatedAt        string  `pulumi:"createdAt"`
	DhcpId           *string `pulumi:"dhcpId"`
	EnableDhcp       bool    `pulumi:"enableDhcp"`
	EnableMasquerade *bool   `pulumi:"enableMasquerade"`
	GatewayId        *string `pulumi:"gatewayId"`
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                        `pulumi:"id"`
	IpamConfigs      []GetGatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	MacAddress       string                        `pulumi:"macAddress"`
	PrivateNetworkId *string                       `pulumi:"privateNetworkId"`
	StaticAddress    string                        `pulumi:"staticAddress"`
	Status           string                        `pulumi:"status"`
	UpdatedAt        string                        `pulumi:"updatedAt"`
	Zone             string                        `pulumi:"zone"`
}

func LookupGatewayNetworkOutput(ctx *pulumi.Context, args LookupGatewayNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayNetworkResult, error) {
			args := v.(LookupGatewayNetworkArgs)
			r, err := LookupGatewayNetwork(ctx, &args, opts...)
			var s LookupGatewayNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayNetworkResultOutput)
}

// A collection of arguments for invoking getGatewayNetwork.
type LookupGatewayNetworkOutputArgs struct {
	// ID of the public gateway DHCP config
	DhcpId pulumi.StringPtrInput `pulumi:"dhcpId"`
	// If masquerade is enabled on requested network
	EnableMasquerade pulumi.BoolPtrInput `pulumi:"enableMasquerade"`
	// ID of the public gateway the gateway network is linked to
	GatewayId pulumi.StringPtrInput `pulumi:"gatewayId"`
	// ID of the gateway network.
	//
	// > Only one of `gatewayNetworkId` or filters should be specified. You can use all the filters you want.
	GatewayNetworkId pulumi.StringPtrInput `pulumi:"gatewayNetworkId"`
	// ID of the private network the gateway network is linked to
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
}

func (LookupGatewayNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayNetwork.
type LookupGatewayNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayNetworkResult)(nil)).Elem()
}

func (o LookupGatewayNetworkResultOutput) ToLookupGatewayNetworkResultOutput() LookupGatewayNetworkResultOutput {
	return o
}

func (o LookupGatewayNetworkResultOutput) ToLookupGatewayNetworkResultOutputWithContext(ctx context.Context) LookupGatewayNetworkResultOutput {
	return o
}

func (o LookupGatewayNetworkResultOutput) CleanupDhcp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) bool { return v.CleanupDhcp }).(pulumi.BoolOutput)
}

func (o LookupGatewayNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) DhcpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) *string { return v.DhcpId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayNetworkResultOutput) EnableDhcp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) bool { return v.EnableDhcp }).(pulumi.BoolOutput)
}

func (o LookupGatewayNetworkResultOutput) EnableMasquerade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) *bool { return v.EnableMasquerade }).(pulumi.BoolPtrOutput)
}

func (o LookupGatewayNetworkResultOutput) GatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) *string { return v.GatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayNetworkResultOutput) GatewayNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) *string { return v.GatewayNetworkId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) IpamConfigs() GetGatewayNetworkIpamConfigArrayOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) []GetGatewayNetworkIpamConfig { return v.IpamConfigs }).(GetGatewayNetworkIpamConfigArrayOutput)
}

func (o LookupGatewayNetworkResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayNetworkResultOutput) StaticAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.StaticAddress }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupGatewayNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayNetworkResultOutput{})
}
