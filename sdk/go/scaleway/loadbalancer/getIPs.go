// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer IPs.
//
// ## Example Usage
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
//			_, err := loadbalancer.GetIPs(ctx, &loadbalancer.GetIPsArgs{
//				IpCidrRange: pulumi.StringRef("0.0.0.0/0"),
//				Zone:        pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetIPs(ctx *pulumi.Context, args *GetIPsArgs, opts ...pulumi.InvokeOption) (*GetIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIPsResult
	err := ctx.Invoke("scaleway:loadbalancer/getIPs:getIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIPs.
type GetIPsArgs struct {
	// The IP CIDR range used as a filter. IPs within a CIDR block like it are listed.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// The ID of the project the load-balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which IPs exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getIPs.
type GetIPsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// List of found IPs
	Ips []GetIPsIp `pulumi:"ips"`
	// The organization ID the load-balancer is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the project the load-balancer is associated with.
	ProjectId string `pulumi:"projectId"`
	// The zone in which the load-balancer is.
	Zone string `pulumi:"zone"`
}

func GetIPsOutput(ctx *pulumi.Context, args GetIPsOutputArgs, opts ...pulumi.InvokeOption) GetIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIPsResult, error) {
			args := v.(GetIPsArgs)
			r, err := GetIPs(ctx, &args, opts...)
			var s GetIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIPsResultOutput)
}

// A collection of arguments for invoking getIPs.
type GetIPsOutputArgs struct {
	// The IP CIDR range used as a filter. IPs within a CIDR block like it are listed.
	IpCidrRange pulumi.StringPtrInput `pulumi:"ipCidrRange"`
	// The ID of the project the load-balancer is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which IPs exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIPsArgs)(nil)).Elem()
}

// A collection of values returned by getIPs.
type GetIPsResultOutput struct{ *pulumi.OutputState }

func (GetIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIPsResult)(nil)).Elem()
}

func (o GetIPsResultOutput) ToGetIPsResultOutput() GetIPsResultOutput {
	return o
}

func (o GetIPsResultOutput) ToGetIPsResultOutputWithContext(ctx context.Context) GetIPsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIPsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIPsResultOutput) IpCidrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIPsResult) *string { return v.IpCidrRange }).(pulumi.StringPtrOutput)
}

// List of found IPs
func (o GetIPsResultOutput) Ips() GetIPsIpArrayOutput {
	return o.ApplyT(func(v GetIPsResult) []GetIPsIp { return v.Ips }).(GetIPsIpArrayOutput)
}

// The organization ID the load-balancer is associated with.
func (o GetIPsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the load-balancer is associated with.
func (o GetIPsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The zone in which the load-balancer is.
func (o GetIPsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIPsResultOutput{})
}
