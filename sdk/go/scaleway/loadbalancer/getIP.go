// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Load Balancer IP.
//
// ## Example Usage
//
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
//			_, err := loadbalancer.LookupIP(ctx, &loadbalancer.LookupIPArgs{
//				IpId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIP(ctx *pulumi.Context, args *LookupIPArgs, opts ...pulumi.InvokeOption) (*LookupIPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIPResult
	err := ctx.Invoke("scaleway:loadbalancer/getIP:getIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIP.
type LookupIPArgs struct {
	// The IP address.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpAddress *string `pulumi:"ipAddress"`
	// The IP ID.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpId *string `pulumi:"ipId"`
	// The ID of the project the LB IP associated with.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getIP.
type LookupIPResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	IpAddress *string `pulumi:"ipAddress"`
	IpId      *string `pulumi:"ipId"`
	// The associated load-balancer ID if any
	LbId string `pulumi:"lbId"`
	// (Defaults to provider `organizationId`) The ID of the organization the LB IP is associated with.
	OrganizationId string `pulumi:"organizationId"`
	ProjectId      string `pulumi:"projectId"`
	Region         string `pulumi:"region"`
	// The reverse domain associated with this IP.
	Reverse string `pulumi:"reverse"`
	Zone    string `pulumi:"zone"`
}

func LookupIPOutput(ctx *pulumi.Context, args LookupIPOutputArgs, opts ...pulumi.InvokeOption) LookupIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPResult, error) {
			args := v.(LookupIPArgs)
			r, err := LookupIP(ctx, &args, opts...)
			var s LookupIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPResultOutput)
}

// A collection of arguments for invoking getIP.
type LookupIPOutputArgs struct {
	// The IP address.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// The IP ID.
	// Only one of `ipAddress` and `ipId` should be specified.
	IpId pulumi.StringPtrInput `pulumi:"ipId"`
	// The ID of the project the LB IP associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPArgs)(nil)).Elem()
}

// A collection of values returned by getIP.
type LookupIPResultOutput struct{ *pulumi.OutputState }

func (LookupIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPResult)(nil)).Elem()
}

func (o LookupIPResultOutput) ToLookupIPResultOutput() LookupIPResultOutput {
	return o
}

func (o LookupIPResultOutput) ToLookupIPResultOutputWithContext(ctx context.Context) LookupIPResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupIPResultOutput) IpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.IpId }).(pulumi.StringPtrOutput)
}

// The associated load-balancer ID if any
func (o LookupIPResultOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.LbId }).(pulumi.StringOutput)
}

// (Defaults to provider `organizationId`) The ID of the organization the LB IP is associated with.
func (o LookupIPResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Region }).(pulumi.StringOutput)
}

// The reverse domain associated with this IP.
func (o LookupIPResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPResultOutput{})
}