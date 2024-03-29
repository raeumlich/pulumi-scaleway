// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Get information about Scaleway Load-Balancer Frontends.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
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
//			ip01, err := loadbalancer.NewIP(ctx, "ip01", nil)
//			if err != nil {
//				return err
//			}
//			lb01, err := loadbalancer.NewLoadBalancer(ctx, "lb01", &loadbalancer.LoadBalancerArgs{
//				IpId: ip01.ID(),
//				Type: pulumi.String("lb-s"),
//			})
//			if err != nil {
//				return err
//			}
//			bkd01, err := loadbalancer.NewBackend(ctx, "bkd01", &loadbalancer.BackendArgs{
//				LbId:            lb01.ID(),
//				ForwardProtocol: pulumi.String("tcp"),
//				ForwardPort:     pulumi.Int(80),
//				ProxyProtocol:   pulumi.String("none"),
//			})
//			if err != nil {
//				return err
//			}
//			frt01, err := loadbalancer.NewFrontend(ctx, "frt01", &loadbalancer.FrontendArgs{
//				LbId:        lb01.ID(),
//				BackendId:   bkd01.ID(),
//				InboundPort: pulumi.Int(80),
//			})
//			if err != nil {
//				return err
//			}
//			_ = loadbalancer.LookupFrontendOutput(ctx, loadbalancer.GetFrontendOutputArgs{
//				FrontendId: frt01.ID(),
//			}, nil)
//			_ = loadbalancer.LookupFrontendOutput(ctx, loadbalancer.GetFrontendOutputArgs{
//				Name: frt01.Name,
//				LbId: lb01.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupFrontend(ctx *pulumi.Context, args *LookupFrontendArgs, opts ...pulumi.InvokeOption) (*LookupFrontendResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFrontendResult
	err := ctx.Invoke("scaleway:loadbalancer/getFrontend:getFrontend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFrontend.
type LookupFrontendArgs struct {
	// The frontend id.
	// - Only one of `name` and `frontendId` should be specified.
	FrontendId *string `pulumi:"frontendId"`
	// The load-balancer ID this frontend is attached to.
	LbId *string `pulumi:"lbId"`
	// The name of the frontend.
	// - When using the `name` you should specify the `lb-id`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getFrontend.
type LookupFrontendResult struct {
	Acls           []GetFrontendAcl `pulumi:"acls"`
	BackendId      string           `pulumi:"backendId"`
	CertificateId  string           `pulumi:"certificateId"`
	CertificateIds []string         `pulumi:"certificateIds"`
	EnableHttp3    bool             `pulumi:"enableHttp3"`
	ExternalAcls   bool             `pulumi:"externalAcls"`
	FrontendId     *string          `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	InboundPort   int     `pulumi:"inboundPort"`
	LbId          *string `pulumi:"lbId"`
	Name          *string `pulumi:"name"`
	TimeoutClient string  `pulumi:"timeoutClient"`
}

func LookupFrontendOutput(ctx *pulumi.Context, args LookupFrontendOutputArgs, opts ...pulumi.InvokeOption) LookupFrontendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFrontendResult, error) {
			args := v.(LookupFrontendArgs)
			r, err := LookupFrontend(ctx, &args, opts...)
			var s LookupFrontendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFrontendResultOutput)
}

// A collection of arguments for invoking getFrontend.
type LookupFrontendOutputArgs struct {
	// The frontend id.
	// - Only one of `name` and `frontendId` should be specified.
	FrontendId pulumi.StringPtrInput `pulumi:"frontendId"`
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The name of the frontend.
	// - When using the `name` you should specify the `lb-id`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupFrontendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendArgs)(nil)).Elem()
}

// A collection of values returned by getFrontend.
type LookupFrontendResultOutput struct{ *pulumi.OutputState }

func (LookupFrontendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendResult)(nil)).Elem()
}

func (o LookupFrontendResultOutput) ToLookupFrontendResultOutput() LookupFrontendResultOutput {
	return o
}

func (o LookupFrontendResultOutput) ToLookupFrontendResultOutputWithContext(ctx context.Context) LookupFrontendResultOutput {
	return o
}

func (o LookupFrontendResultOutput) Acls() GetFrontendAclArrayOutput {
	return o.ApplyT(func(v LookupFrontendResult) []GetFrontendAcl { return v.Acls }).(GetFrontendAclArrayOutput)
}

func (o LookupFrontendResultOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendResult) string { return v.BackendId }).(pulumi.StringOutput)
}

func (o LookupFrontendResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o LookupFrontendResultOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFrontendResult) []string { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

func (o LookupFrontendResultOutput) EnableHttp3() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFrontendResult) bool { return v.EnableHttp3 }).(pulumi.BoolOutput)
}

func (o LookupFrontendResultOutput) ExternalAcls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFrontendResult) bool { return v.ExternalAcls }).(pulumi.BoolOutput)
}

func (o LookupFrontendResultOutput) FrontendId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendResult) *string { return v.FrontendId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFrontendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFrontendResultOutput) InboundPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFrontendResult) int { return v.InboundPort }).(pulumi.IntOutput)
}

func (o LookupFrontendResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupFrontendResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFrontendResultOutput) TimeoutClient() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendResult) string { return v.TimeoutClient }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFrontendResultOutput{})
}
