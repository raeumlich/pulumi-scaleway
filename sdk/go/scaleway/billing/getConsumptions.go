// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about your Consumptions.
func GetConsumptions(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetConsumptionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetConsumptionsResult
	err := ctx.Invoke("scaleway:billing/getConsumptions:getConsumptions", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getConsumptions.
type GetConsumptionsResult struct {
	// List of found consumptions
	Consumptions []GetConsumptionsConsumption `pulumi:"consumptions"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	OrganizationId string `pulumi:"organizationId"`
	// The last consumption update date.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetConsumptionsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetConsumptionsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetConsumptionsResult, error) {
		r, err := GetConsumptions(ctx, opts...)
		var s GetConsumptionsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetConsumptionsResultOutput)
}

// A collection of values returned by getConsumptions.
type GetConsumptionsResultOutput struct{ *pulumi.OutputState }

func (GetConsumptionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConsumptionsResult)(nil)).Elem()
}

func (o GetConsumptionsResultOutput) ToGetConsumptionsResultOutput() GetConsumptionsResultOutput {
	return o
}

func (o GetConsumptionsResultOutput) ToGetConsumptionsResultOutputWithContext(ctx context.Context) GetConsumptionsResultOutput {
	return o
}

// List of found consumptions
func (o GetConsumptionsResultOutput) Consumptions() GetConsumptionsConsumptionArrayOutput {
	return o.ApplyT(func(v GetConsumptionsResult) []GetConsumptionsConsumption { return v.Consumptions }).(GetConsumptionsConsumptionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetConsumptionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetConsumptionsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The last consumption update date.
func (o GetConsumptionsResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConsumptionsResultOutput{})
}