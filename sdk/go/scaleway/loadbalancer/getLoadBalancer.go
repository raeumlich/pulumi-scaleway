// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Load Balancer.
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
//			_, err := loadbalancer.LookupLoadBalancer(ctx, &loadbalancer.LookupLoadBalancerArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = loadbalancer.LookupLoadBalancer(ctx, &loadbalancer.LookupLoadBalancerArgs{
//				LbId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
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
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("scaleway:loadbalancer/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	LbId *string `pulumi:"lbId"`
	// The load balancer name.
	Name *string `pulumi:"name"`
	// The ID of the project the LB is associated with.
	ProjectId *string `pulumi:"projectId"`
	ReleaseIp *bool   `pulumi:"releaseIp"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResult struct {
	AssignFlexibleIp bool   `pulumi:"assignFlexibleIp"`
	Description      string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The load-balancer public IP Address.
	IpAddress             string                          `pulumi:"ipAddress"`
	IpId                  string                          `pulumi:"ipId"`
	LbId                  *string                         `pulumi:"lbId"`
	Name                  *string                         `pulumi:"name"`
	OrganizationId        string                          `pulumi:"organizationId"`
	PrivateNetworks       []GetLoadBalancerPrivateNetwork `pulumi:"privateNetworks"`
	ProjectId             *string                         `pulumi:"projectId"`
	Region                string                          `pulumi:"region"`
	ReleaseIp             *bool                           `pulumi:"releaseIp"`
	SslCompatibilityLevel string                          `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancer.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer.
	Type string `pulumi:"type"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone *string `pulumi:"zone"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerOutputArgs struct {
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The load balancer name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the LB is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	ReleaseIp pulumi.BoolPtrInput   `pulumi:"releaseIp"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) AssignFlexibleIp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) bool { return v.AssignFlexibleIp }).(pulumi.BoolOutput)
}

func (o LookupLoadBalancerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLoadBalancerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The load-balancer public IP Address.
func (o LookupLoadBalancerResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.IpId }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) PrivateNetworks() GetLoadBalancerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []GetLoadBalancerPrivateNetwork { return v.PrivateNetworks }).(GetLoadBalancerPrivateNetworkArrayOutput)
}

func (o LookupLoadBalancerResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) ReleaseIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *bool { return v.ReleaseIp }).(pulumi.BoolPtrOutput)
}

func (o LookupLoadBalancerResultOutput) SslCompatibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.SslCompatibilityLevel }).(pulumi.StringOutput)
}

// The tags associated with the load-balancer.
func (o LookupLoadBalancerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the load-balancer.
func (o LookupLoadBalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

// (Defaults to provider `zone`) The zone in which the LB exists.
func (o LookupLoadBalancerResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
