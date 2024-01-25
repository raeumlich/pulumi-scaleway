// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Load-Balancer ACLs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls).
//
// ## Example Usage
// ### Basic
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
//			_, err := loadbalancer.NewACL(ctx, "acl01", &loadbalancer.ACLArgs{
//				FrontendId:  pulumi.Any(scaleway_lb_frontend.Frt01.Id),
//				Description: pulumi.String("Exclude well-known IPs"),
//				Index:       pulumi.Int(0),
//				Action: &loadbalancer.ACLActionArgs{
//					Type: pulumi.String("allow"),
//				},
//				Match: &loadbalancer.ACLMatchArgs{
//					IpSubnets: pulumi.StringArray{
//						pulumi.String("192.168.0.1"),
//						pulumi.String("192.168.0.2"),
//						pulumi.String("192.168.10.0/24"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Load-Balancer ACL can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:loadbalancer/aCL:ACL acl01 fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type ACL struct {
	pulumi.CustomResourceState

	// Action to undertake when an ACL filter matches.
	Action ACLActionOutput `pulumi:"action"`
	// Date and time of ACL's creation (RFC 3339 format)
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ACL description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringOutput `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntOutput `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match ACLMatchPtrOutput `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// Date and time of ACL's update (RFC 3339 format)
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewACL registers a new resource with the given unique name, arguments, and options.
func NewACL(ctx *pulumi.Context,
	name string, args *ACLArgs, opts ...pulumi.ResourceOption) (*ACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FrontendId == nil {
		return nil, errors.New("invalid value for required argument 'FrontendId'")
	}
	if args.Index == nil {
		return nil, errors.New("invalid value for required argument 'Index'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ACL
	err := ctx.RegisterResource("scaleway:loadbalancer/aCL:ACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetACL gets an existing ACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ACLState, opts ...pulumi.ResourceOption) (*ACL, error) {
	var resource ACL
	err := ctx.ReadResource("scaleway:loadbalancer/aCL:ACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ACL resources.
type aclState struct {
	// Action to undertake when an ACL filter matches.
	Action *ACLAction `pulumi:"action"`
	// Date and time of ACL's creation (RFC 3339 format)
	CreatedAt *string `pulumi:"createdAt"`
	// The ACL description.
	Description *string `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId *string `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index *int `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match *ACLMatch `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// Date and time of ACL's update (RFC 3339 format)
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ACLState struct {
	// Action to undertake when an ACL filter matches.
	Action ACLActionPtrInput
	// Date and time of ACL's creation (RFC 3339 format)
	CreatedAt pulumi.StringPtrInput
	// The ACL description.
	Description pulumi.StringPtrInput
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringPtrInput
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntPtrInput
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match ACLMatchPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// Date and time of ACL's update (RFC 3339 format)
	UpdatedAt pulumi.StringPtrInput
}

func (ACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// Action to undertake when an ACL filter matches.
	Action ACLAction `pulumi:"action"`
	// The ACL description.
	Description *string `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId string `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match *ACLMatch `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ACL resource.
type ACLArgs struct {
	// Action to undertake when an ACL filter matches.
	Action ACLActionInput
	// The ACL description.
	Description pulumi.StringPtrInput
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringInput
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntInput
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match ACLMatchPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
}

func (ACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type ACLInput interface {
	pulumi.Input

	ToACLOutput() ACLOutput
	ToACLOutputWithContext(ctx context.Context) ACLOutput
}

func (*ACL) ElementType() reflect.Type {
	return reflect.TypeOf((**ACL)(nil)).Elem()
}

func (i *ACL) ToACLOutput() ACLOutput {
	return i.ToACLOutputWithContext(context.Background())
}

func (i *ACL) ToACLOutputWithContext(ctx context.Context) ACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACLOutput)
}

// ACLArrayInput is an input type that accepts ACLArray and ACLArrayOutput values.
// You can construct a concrete instance of `ACLArrayInput` via:
//
//	ACLArray{ ACLArgs{...} }
type ACLArrayInput interface {
	pulumi.Input

	ToACLArrayOutput() ACLArrayOutput
	ToACLArrayOutputWithContext(context.Context) ACLArrayOutput
}

type ACLArray []ACLInput

func (ACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ACL)(nil)).Elem()
}

func (i ACLArray) ToACLArrayOutput() ACLArrayOutput {
	return i.ToACLArrayOutputWithContext(context.Background())
}

func (i ACLArray) ToACLArrayOutputWithContext(ctx context.Context) ACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACLArrayOutput)
}

// ACLMapInput is an input type that accepts ACLMap and ACLMapOutput values.
// You can construct a concrete instance of `ACLMapInput` via:
//
//	ACLMap{ "key": ACLArgs{...} }
type ACLMapInput interface {
	pulumi.Input

	ToACLMapOutput() ACLMapOutput
	ToACLMapOutputWithContext(context.Context) ACLMapOutput
}

type ACLMap map[string]ACLInput

func (ACLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ACL)(nil)).Elem()
}

func (i ACLMap) ToACLMapOutput() ACLMapOutput {
	return i.ToACLMapOutputWithContext(context.Background())
}

func (i ACLMap) ToACLMapOutputWithContext(ctx context.Context) ACLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACLMapOutput)
}

type ACLOutput struct{ *pulumi.OutputState }

func (ACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACL)(nil)).Elem()
}

func (o ACLOutput) ToACLOutput() ACLOutput {
	return o
}

func (o ACLOutput) ToACLOutputWithContext(ctx context.Context) ACLOutput {
	return o
}

// Action to undertake when an ACL filter matches.
func (o ACLOutput) Action() ACLActionOutput {
	return o.ApplyT(func(v *ACL) ACLActionOutput { return v.Action }).(ACLActionOutput)
}

// Date and time of ACL's creation (RFC 3339 format)
func (o ACLOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ACL) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ACL description.
func (o ACLOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACL) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The load-balancer Frontend ID to attach the ACL to.
func (o ACLOutput) FrontendId() pulumi.StringOutput {
	return o.ApplyT(func(v *ACL) pulumi.StringOutput { return v.FrontendId }).(pulumi.StringOutput)
}

// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
func (o ACLOutput) Index() pulumi.IntOutput {
	return o.ApplyT(func(v *ACL) pulumi.IntOutput { return v.Index }).(pulumi.IntOutput)
}

// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
func (o ACLOutput) Match() ACLMatchPtrOutput {
	return o.ApplyT(func(v *ACL) ACLMatchPtrOutput { return v.Match }).(ACLMatchPtrOutput)
}

// The ACL name. If not provided it will be randomly generated.
func (o ACLOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ACL) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Date and time of ACL's update (RFC 3339 format)
func (o ACLOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ACL) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ACLArrayOutput struct{ *pulumi.OutputState }

func (ACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ACL)(nil)).Elem()
}

func (o ACLArrayOutput) ToACLArrayOutput() ACLArrayOutput {
	return o
}

func (o ACLArrayOutput) ToACLArrayOutputWithContext(ctx context.Context) ACLArrayOutput {
	return o
}

func (o ACLArrayOutput) Index(i pulumi.IntInput) ACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ACL {
		return vs[0].([]*ACL)[vs[1].(int)]
	}).(ACLOutput)
}

type ACLMapOutput struct{ *pulumi.OutputState }

func (ACLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ACL)(nil)).Elem()
}

func (o ACLMapOutput) ToACLMapOutput() ACLMapOutput {
	return o
}

func (o ACLMapOutput) ToACLMapOutputWithContext(ctx context.Context) ACLMapOutput {
	return o
}

func (o ACLMapOutput) MapIndex(k pulumi.StringInput) ACLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ACL {
		return vs[0].(map[string]*ACL)[vs[1].(string)]
	}).(ACLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ACLInput)(nil)).Elem(), &ACL{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACLArrayInput)(nil)).Elem(), ACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACLMapInput)(nil)).Elem(), ACLMap{})
	pulumi.RegisterOutputType(ACLOutput{})
	pulumi.RegisterOutputType(ACLArrayOutput{})
	pulumi.RegisterOutputType(ACLMapOutput{})
}
