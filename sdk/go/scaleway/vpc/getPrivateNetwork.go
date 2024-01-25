// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a private network.
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
//			_, err := vpc.LookupPrivateNetwork(ctx, &vpc.LookupPrivateNetworkArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vpc.LookupPrivateNetwork(ctx, &vpc.LookupPrivateNetworkArgs{
//				Name:  pulumi.StringRef("foobar"),
//				VpcId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vpc.LookupPrivateNetwork(ctx, &vpc.LookupPrivateNetworkArgs{
//				PrivateNetworkId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPrivateNetwork(ctx *pulumi.Context, args *LookupPrivateNetworkArgs, opts ...pulumi.InvokeOption) (*LookupPrivateNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateNetworkResult
	err := ctx.Invoke("scaleway:vpc/getPrivateNetwork:getPrivateNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateNetwork.
type LookupPrivateNetworkArgs struct {
	// Name of the private network. Cannot be used with `privateNetworkId`.
	Name *string `pulumi:"name"`
	// ID of the private network. Cannot be used with `name` and `vpcId`.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// ID of the VPC in which the private network is. Cannot be used with `privateNetworkId`.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getPrivateNetwork.
type LookupPrivateNetworkResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IPv4 subnet associated with the private network.
	Ipv4Subnets []GetPrivateNetworkIpv4Subnet `pulumi:"ipv4Subnets"`
	// The IPv6 subnets associated with the private network.
	Ipv6Subnets      []GetPrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	IsRegional       bool                          `pulumi:"isRegional"`
	Name             *string                       `pulumi:"name"`
	OrganizationId   string                        `pulumi:"organizationId"`
	PrivateNetworkId *string                       `pulumi:"privateNetworkId"`
	ProjectId        *string                       `pulumi:"projectId"`
	Region           string                        `pulumi:"region"`
	Tags             []string                      `pulumi:"tags"`
	UpdatedAt        string                        `pulumi:"updatedAt"`
	VpcId            *string                       `pulumi:"vpcId"`
	Zone             string                        `pulumi:"zone"`
}

func LookupPrivateNetworkOutput(ctx *pulumi.Context, args LookupPrivateNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateNetworkResult, error) {
			args := v.(LookupPrivateNetworkArgs)
			r, err := LookupPrivateNetwork(ctx, &args, opts...)
			var s LookupPrivateNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateNetworkResultOutput)
}

// A collection of arguments for invoking getPrivateNetwork.
type LookupPrivateNetworkOutputArgs struct {
	// Name of the private network. Cannot be used with `privateNetworkId`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the private network. Cannot be used with `name` and `vpcId`.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// The ID of the project the private network is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// ID of the VPC in which the private network is. Cannot be used with `privateNetworkId`.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupPrivateNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateNetwork.
type LookupPrivateNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNetworkResult)(nil)).Elem()
}

func (o LookupPrivateNetworkResultOutput) ToLookupPrivateNetworkResultOutput() LookupPrivateNetworkResultOutput {
	return o
}

func (o LookupPrivateNetworkResultOutput) ToLookupPrivateNetworkResultOutputWithContext(ctx context.Context) LookupPrivateNetworkResultOutput {
	return o
}

func (o LookupPrivateNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IPv4 subnet associated with the private network.
func (o LookupPrivateNetworkResultOutput) Ipv4Subnets() GetPrivateNetworkIpv4SubnetArrayOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) []GetPrivateNetworkIpv4Subnet { return v.Ipv4Subnets }).(GetPrivateNetworkIpv4SubnetArrayOutput)
}

// The IPv6 subnets associated with the private network.
func (o LookupPrivateNetworkResultOutput) Ipv6Subnets() GetPrivateNetworkIpv6SubnetArrayOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) []GetPrivateNetworkIpv6Subnet { return v.Ipv6Subnets }).(GetPrivateNetworkIpv6SubnetArrayOutput)
}

func (o LookupPrivateNetworkResultOutput) IsRegional() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) bool { return v.IsRegional }).(pulumi.BoolOutput)
}

func (o LookupPrivateNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateNetworkResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupPrivateNetworkResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateNetworkResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateNetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupPrivateNetworkResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupPrivateNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupPrivateNetworkResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateNetworkResultOutput{})
}