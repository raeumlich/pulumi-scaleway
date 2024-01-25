// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Add members to an IAM group.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
//
// ## Example Usage
// ### Application Membership
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			group, err := iam.NewGroup(ctx, "group", &iam.GroupArgs{
//				ExternalMembership: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			app, err := iam.NewApplication(ctx, "app", nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewGroupMembership(ctx, "member", &iam.GroupMembershipArgs{
//				GroupId:       group.ID(),
//				ApplicationId: app.ID(),
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
// IAM group memberships can be imported using two format- For user`{group_id}/user/{user_id}` - For application`{group_id}/app/{application_id}` bash
//
// ```sh
//
//	$ pulumi import scaleway:iam/groupMembership:GroupMembership app 11111111-1111-1111-1111-111111111111/app/11111111-1111-1111-1111-111111111111
//
// ```
type GroupMembership struct {
	pulumi.CustomResourceState

	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMembership
	err := ctx.RegisterResource("scaleway:iam/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("scaleway:iam/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// The ID of the application that will be added to the group.
	ApplicationId *string `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId *string `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId *string `pulumi:"userId"`
}

type GroupMembershipState struct {
	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrInput
	// ID of the group to add members to.
	GroupId pulumi.StringPtrInput
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// The ID of the application that will be added to the group.
	ApplicationId *string `pulumi:"applicationId"`
	// ID of the group to add members to.
	GroupId string `pulumi:"groupId"`
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// The ID of the application that will be added to the group.
	ApplicationId pulumi.StringPtrInput
	// ID of the group to add members to.
	GroupId pulumi.StringInput
	// The ID of the user that will be added to the group
	//
	// - > Only one of `applicationId` or `userId` must be specified
	UserId pulumi.StringPtrInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}

type GroupMembershipInput interface {
	pulumi.Input

	ToGroupMembershipOutput() GroupMembershipOutput
	ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput
}

func (*GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (i *GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i *GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

// GroupMembershipArrayInput is an input type that accepts GroupMembershipArray and GroupMembershipArrayOutput values.
// You can construct a concrete instance of `GroupMembershipArrayInput` via:
//
//	GroupMembershipArray{ GroupMembershipArgs{...} }
type GroupMembershipArrayInput interface {
	pulumi.Input

	ToGroupMembershipArrayOutput() GroupMembershipArrayOutput
	ToGroupMembershipArrayOutputWithContext(context.Context) GroupMembershipArrayOutput
}

type GroupMembershipArray []GroupMembershipInput

func (GroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return i.ToGroupMembershipArrayOutputWithContext(context.Background())
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipArrayOutput)
}

// GroupMembershipMapInput is an input type that accepts GroupMembershipMap and GroupMembershipMapOutput values.
// You can construct a concrete instance of `GroupMembershipMapInput` via:
//
//	GroupMembershipMap{ "key": GroupMembershipArgs{...} }
type GroupMembershipMapInput interface {
	pulumi.Input

	ToGroupMembershipMapOutput() GroupMembershipMapOutput
	ToGroupMembershipMapOutputWithContext(context.Context) GroupMembershipMapOutput
}

type GroupMembershipMap map[string]GroupMembershipInput

func (GroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (i GroupMembershipMap) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return i.ToGroupMembershipMapOutputWithContext(context.Background())
}

func (i GroupMembershipMap) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipMapOutput)
}

type GroupMembershipOutput struct{ *pulumi.OutputState }

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil)).Elem()
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

// The ID of the application that will be added to the group.
func (o GroupMembershipOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// ID of the group to add members to.
func (o GroupMembershipOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The ID of the user that will be added to the group
//
// - > Only one of `applicationId` or `userId` must be specified
func (o GroupMembershipOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembership) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type GroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) Index(i pulumi.IntInput) GroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].([]*GroupMembership)[vs[1].(int)]
	}).(GroupMembershipOutput)
}

type GroupMembershipMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembership)(nil)).Elem()
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMembership {
		return vs[0].(map[string]*GroupMembership)[vs[1].(string)]
	}).(GroupMembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipInput)(nil)).Elem(), &GroupMembership{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipArrayInput)(nil)).Elem(), GroupMembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipMapInput)(nil)).Elem(), GroupMembershipMap{})
	pulumi.RegisterOutputType(GroupMembershipOutput{})
	pulumi.RegisterOutputType(GroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipMapOutput{})
}