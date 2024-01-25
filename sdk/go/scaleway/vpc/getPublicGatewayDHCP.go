// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a public gateway DHCP.
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
//			main, err := vpc.NewPublicGatewayDHCP(ctx, "main", &vpc.PublicGatewayDHCPArgs{
//				Subnet: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.LookupPublicGatewayDHCPOutput(ctx, vpc.GetPublicGatewayDHCPOutputArgs{
//				DhcpId: main.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupPublicGatewayDHCP(ctx *pulumi.Context, args *LookupPublicGatewayDHCPArgs, opts ...pulumi.InvokeOption) (*LookupPublicGatewayDHCPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicGatewayDHCPResult
	err := ctx.Invoke("scaleway:vpc/getPublicGatewayDHCP:getPublicGatewayDHCP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicGatewayDHCP.
type LookupPublicGatewayDHCPArgs struct {
	DhcpId string `pulumi:"dhcpId"`
}

// A collection of values returned by getPublicGatewayDHCP.
type LookupPublicGatewayDHCPResult struct {
	Address             string   `pulumi:"address"`
	CreatedAt           string   `pulumi:"createdAt"`
	DhcpId              string   `pulumi:"dhcpId"`
	DnsLocalName        string   `pulumi:"dnsLocalName"`
	DnsSearches         []string `pulumi:"dnsSearches"`
	DnsServersOverrides []string `pulumi:"dnsServersOverrides"`
	EnableDynamic       bool     `pulumi:"enableDynamic"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	OrganizationId   string `pulumi:"organizationId"`
	PoolHigh         string `pulumi:"poolHigh"`
	PoolLow          string `pulumi:"poolLow"`
	ProjectId        string `pulumi:"projectId"`
	PushDefaultRoute bool   `pulumi:"pushDefaultRoute"`
	PushDnsServer    bool   `pulumi:"pushDnsServer"`
	RebindTimer      int    `pulumi:"rebindTimer"`
	RenewTimer       int    `pulumi:"renewTimer"`
	Subnet           string `pulumi:"subnet"`
	UpdatedAt        string `pulumi:"updatedAt"`
	ValidLifetime    int    `pulumi:"validLifetime"`
	Zone             string `pulumi:"zone"`
}

func LookupPublicGatewayDHCPOutput(ctx *pulumi.Context, args LookupPublicGatewayDHCPOutputArgs, opts ...pulumi.InvokeOption) LookupPublicGatewayDHCPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicGatewayDHCPResult, error) {
			args := v.(LookupPublicGatewayDHCPArgs)
			r, err := LookupPublicGatewayDHCP(ctx, &args, opts...)
			var s LookupPublicGatewayDHCPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicGatewayDHCPResultOutput)
}

// A collection of arguments for invoking getPublicGatewayDHCP.
type LookupPublicGatewayDHCPOutputArgs struct {
	DhcpId pulumi.StringInput `pulumi:"dhcpId"`
}

func (LookupPublicGatewayDHCPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayDHCPArgs)(nil)).Elem()
}

// A collection of values returned by getPublicGatewayDHCP.
type LookupPublicGatewayDHCPResultOutput struct{ *pulumi.OutputState }

func (LookupPublicGatewayDHCPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayDHCPResult)(nil)).Elem()
}

func (o LookupPublicGatewayDHCPResultOutput) ToLookupPublicGatewayDHCPResultOutput() LookupPublicGatewayDHCPResultOutput {
	return o
}

func (o LookupPublicGatewayDHCPResultOutput) ToLookupPublicGatewayDHCPResultOutputWithContext(ctx context.Context) LookupPublicGatewayDHCPResultOutput {
	return o
}

func (o LookupPublicGatewayDHCPResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) DhcpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.DhcpId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) DnsLocalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.DnsLocalName }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) DnsSearches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) []string { return v.DnsSearches }).(pulumi.StringArrayOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) DnsServersOverrides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) []string { return v.DnsServersOverrides }).(pulumi.StringArrayOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) EnableDynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) bool { return v.EnableDynamic }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPublicGatewayDHCPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) PoolHigh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.PoolHigh }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) PoolLow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.PoolLow }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) PushDefaultRoute() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) bool { return v.PushDefaultRoute }).(pulumi.BoolOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) PushDnsServer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) bool { return v.PushDnsServer }).(pulumi.BoolOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) RebindTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) int { return v.RebindTimer }).(pulumi.IntOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) RenewTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) int { return v.RenewTimer }).(pulumi.IntOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) ValidLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) int { return v.ValidLifetime }).(pulumi.IntOutput)
}

func (o LookupPublicGatewayDHCPResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicGatewayDHCPResultOutput{})
}