// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple Load Balancer Frontends.
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
//			_, err := loadbalancer.GetFrontends(ctx, &loadbalancer.GetFrontendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = loadbalancer.GetFrontends(ctx, &loadbalancer.GetFrontendsArgs{
//				LbId: scaleway_lb.Lb01.Id,
//				Name: pulumi.StringRef("tf-frontend-datasource"),
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
func GetFrontends(ctx *pulumi.Context, args *GetFrontendsArgs, opts ...pulumi.InvokeOption) (*GetFrontendsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFrontendsResult
	err := ctx.Invoke("scaleway:loadbalancer/getFrontends:getFrontends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFrontends.
type GetFrontendsArgs struct {
	// The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
	LbId string `pulumi:"lbId"`
	// The frontend name used as filter. Frontends with a name like it are listed.
	Name      *string `pulumi:"name"`
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which frontends exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getFrontends.
type GetFrontendsResult struct {
	// List of found frontends
	Frontends []GetFrontendsFrontend `pulumi:"frontends"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	LbId           string  `pulumi:"lbId"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	Zone           string  `pulumi:"zone"`
}

func GetFrontendsOutput(ctx *pulumi.Context, args GetFrontendsOutputArgs, opts ...pulumi.InvokeOption) GetFrontendsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFrontendsResult, error) {
			args := v.(GetFrontendsArgs)
			r, err := GetFrontends(ctx, &args, opts...)
			var s GetFrontendsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFrontendsResultOutput)
}

// A collection of arguments for invoking getFrontends.
type GetFrontendsOutputArgs struct {
	// The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
	LbId pulumi.StringInput `pulumi:"lbId"`
	// The frontend name used as filter. Frontends with a name like it are listed.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which frontends exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetFrontendsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFrontendsArgs)(nil)).Elem()
}

// A collection of values returned by getFrontends.
type GetFrontendsResultOutput struct{ *pulumi.OutputState }

func (GetFrontendsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFrontendsResult)(nil)).Elem()
}

func (o GetFrontendsResultOutput) ToGetFrontendsResultOutput() GetFrontendsResultOutput {
	return o
}

func (o GetFrontendsResultOutput) ToGetFrontendsResultOutputWithContext(ctx context.Context) GetFrontendsResultOutput {
	return o
}

// List of found frontends
func (o GetFrontendsResultOutput) Frontends() GetFrontendsFrontendArrayOutput {
	return o.ApplyT(func(v GetFrontendsResult) []GetFrontendsFrontend { return v.Frontends }).(GetFrontendsFrontendArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFrontendsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.LbId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFrontendsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetFrontendsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetFrontendsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetFrontendsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFrontendsResultOutput{})
}
