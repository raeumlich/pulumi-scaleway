// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cockpit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// ## Import
//
// Cockpits can be imported using the `{project_id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:cockpit/cockpit:Cockpit main 11111111-1111-1111-1111-111111111111
// ```
type Cockpit struct {
	pulumi.CustomResourceState

	// Endpoints.
	Endpoints CockpitEndpointArrayOutput `pulumi:"endpoints"`
	// Name or ID of the plan to use.
	Plan pulumi.StringPtrOutput `pulumi:"plan"`
	// The ID of the current plan.
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Push_url
	PushUrls CockpitPushUrlArrayOutput `pulumi:"pushUrls"`
}

// NewCockpit registers a new resource with the given unique name, arguments, and options.
func NewCockpit(ctx *pulumi.Context,
	name string, args *CockpitArgs, opts ...pulumi.ResourceOption) (*Cockpit, error) {
	if args == nil {
		args = &CockpitArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cockpit
	err := ctx.RegisterResource("scaleway:cockpit/cockpit:Cockpit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCockpit gets an existing Cockpit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCockpit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CockpitState, opts ...pulumi.ResourceOption) (*Cockpit, error) {
	var resource Cockpit
	err := ctx.ReadResource("scaleway:cockpit/cockpit:Cockpit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cockpit resources.
type cockpitState struct {
	// Endpoints.
	Endpoints []CockpitEndpoint `pulumi:"endpoints"`
	// Name or ID of the plan to use.
	Plan *string `pulumi:"plan"`
	// The ID of the current plan.
	PlanId *string `pulumi:"planId"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Push_url
	PushUrls []CockpitPushUrl `pulumi:"pushUrls"`
}

type CockpitState struct {
	// Endpoints.
	Endpoints CockpitEndpointArrayInput
	// Name or ID of the plan to use.
	Plan pulumi.StringPtrInput
	// The ID of the current plan.
	PlanId pulumi.StringPtrInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// Push_url
	PushUrls CockpitPushUrlArrayInput
}

func (CockpitState) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitState)(nil)).Elem()
}

type cockpitArgs struct {
	// Name or ID of the plan to use.
	Plan *string `pulumi:"plan"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
}

// The set of arguments for constructing a Cockpit resource.
type CockpitArgs struct {
	// Name or ID of the plan to use.
	Plan pulumi.StringPtrInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
}

func (CockpitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitArgs)(nil)).Elem()
}

type CockpitInput interface {
	pulumi.Input

	ToCockpitOutput() CockpitOutput
	ToCockpitOutputWithContext(ctx context.Context) CockpitOutput
}

func (*Cockpit) ElementType() reflect.Type {
	return reflect.TypeOf((**Cockpit)(nil)).Elem()
}

func (i *Cockpit) ToCockpitOutput() CockpitOutput {
	return i.ToCockpitOutputWithContext(context.Background())
}

func (i *Cockpit) ToCockpitOutputWithContext(ctx context.Context) CockpitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitOutput)
}

// CockpitArrayInput is an input type that accepts CockpitArray and CockpitArrayOutput values.
// You can construct a concrete instance of `CockpitArrayInput` via:
//
//	CockpitArray{ CockpitArgs{...} }
type CockpitArrayInput interface {
	pulumi.Input

	ToCockpitArrayOutput() CockpitArrayOutput
	ToCockpitArrayOutputWithContext(context.Context) CockpitArrayOutput
}

type CockpitArray []CockpitInput

func (CockpitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cockpit)(nil)).Elem()
}

func (i CockpitArray) ToCockpitArrayOutput() CockpitArrayOutput {
	return i.ToCockpitArrayOutputWithContext(context.Background())
}

func (i CockpitArray) ToCockpitArrayOutputWithContext(ctx context.Context) CockpitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitArrayOutput)
}

// CockpitMapInput is an input type that accepts CockpitMap and CockpitMapOutput values.
// You can construct a concrete instance of `CockpitMapInput` via:
//
//	CockpitMap{ "key": CockpitArgs{...} }
type CockpitMapInput interface {
	pulumi.Input

	ToCockpitMapOutput() CockpitMapOutput
	ToCockpitMapOutputWithContext(context.Context) CockpitMapOutput
}

type CockpitMap map[string]CockpitInput

func (CockpitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cockpit)(nil)).Elem()
}

func (i CockpitMap) ToCockpitMapOutput() CockpitMapOutput {
	return i.ToCockpitMapOutputWithContext(context.Background())
}

func (i CockpitMap) ToCockpitMapOutputWithContext(ctx context.Context) CockpitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitMapOutput)
}

type CockpitOutput struct{ *pulumi.OutputState }

func (CockpitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cockpit)(nil)).Elem()
}

func (o CockpitOutput) ToCockpitOutput() CockpitOutput {
	return o
}

func (o CockpitOutput) ToCockpitOutputWithContext(ctx context.Context) CockpitOutput {
	return o
}

// Endpoints.
func (o CockpitOutput) Endpoints() CockpitEndpointArrayOutput {
	return o.ApplyT(func(v *Cockpit) CockpitEndpointArrayOutput { return v.Endpoints }).(CockpitEndpointArrayOutput)
}

// Name or ID of the plan to use.
func (o CockpitOutput) Plan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cockpit) pulumi.StringPtrOutput { return v.Plan }).(pulumi.StringPtrOutput)
}

// The ID of the current plan.
func (o CockpitOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cockpit) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the cockpit is associated with.
func (o CockpitOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cockpit) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Push_url
func (o CockpitOutput) PushUrls() CockpitPushUrlArrayOutput {
	return o.ApplyT(func(v *Cockpit) CockpitPushUrlArrayOutput { return v.PushUrls }).(CockpitPushUrlArrayOutput)
}

type CockpitArrayOutput struct{ *pulumi.OutputState }

func (CockpitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cockpit)(nil)).Elem()
}

func (o CockpitArrayOutput) ToCockpitArrayOutput() CockpitArrayOutput {
	return o
}

func (o CockpitArrayOutput) ToCockpitArrayOutputWithContext(ctx context.Context) CockpitArrayOutput {
	return o
}

func (o CockpitArrayOutput) Index(i pulumi.IntInput) CockpitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cockpit {
		return vs[0].([]*Cockpit)[vs[1].(int)]
	}).(CockpitOutput)
}

type CockpitMapOutput struct{ *pulumi.OutputState }

func (CockpitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cockpit)(nil)).Elem()
}

func (o CockpitMapOutput) ToCockpitMapOutput() CockpitMapOutput {
	return o
}

func (o CockpitMapOutput) ToCockpitMapOutputWithContext(ctx context.Context) CockpitMapOutput {
	return o
}

func (o CockpitMapOutput) MapIndex(k pulumi.StringInput) CockpitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cockpit {
		return vs[0].(map[string]*Cockpit)[vs[1].(string)]
	}).(CockpitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitInput)(nil)).Elem(), &Cockpit{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitArrayInput)(nil)).Elem(), CockpitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitMapInput)(nil)).Elem(), CockpitMap{})
	pulumi.RegisterOutputType(CockpitOutput{})
	pulumi.RegisterOutputType(CockpitArrayOutput{})
	pulumi.RegisterOutputType(CockpitMapOutput{})
}
