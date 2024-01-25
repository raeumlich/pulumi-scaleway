// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package instance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about an instance IP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/instance"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := instance.LookupIP(ctx, &instance.LookupIPArgs{
//				Id: pulumi.StringRef("fr-par-1/11111111-1111-1111-1111-111111111111"),
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
	err := ctx.Invoke("scaleway:instance/getIP:getIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIP.
type LookupIPArgs struct {
	// The IPv4 address to retrieve
	// Only one of `address` and `id` should be specified.
	Address *string `pulumi:"address"`
	// The ID of the IP address to retrieve
	// Only one of `address` and `id` should be specified.
	Id *string `pulumi:"id"`
}

// A collection of values returned by getIP.
type LookupIPResult struct {
	// The IP address.
	Address *string `pulumi:"address"`
	// The ID of the IP.
	Id *string `pulumi:"id"`
	// The organization ID the IP is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The IP Prefix.
	Prefix    string `pulumi:"prefix"`
	ProjectId string `pulumi:"projectId"`
	// The reverse dns attached to this IP
	Reverse  string   `pulumi:"reverse"`
	ServerId string   `pulumi:"serverId"`
	Tags     []string `pulumi:"tags"`
	// The type of the IP
	Type string `pulumi:"type"`
	Zone string `pulumi:"zone"`
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
	// The IPv4 address to retrieve
	// Only one of `address` and `id` should be specified.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The ID of the IP address to retrieve
	// Only one of `address` and `id` should be specified.
	Id pulumi.StringPtrInput `pulumi:"id"`
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

// The IP address.
func (o LookupIPResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The ID of the IP.
func (o LookupIPResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The organization ID the IP is associated with.
func (o LookupIPResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The IP Prefix.
func (o LookupIPResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Prefix }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse dns attached to this IP
func (o LookupIPResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Reverse }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.ServerId }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIPResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the IP
func (o LookupIPResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPResultOutput{})
}
