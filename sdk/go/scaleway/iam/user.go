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

// Creates and manages Scaleway IAM Users.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/iam/#path-users-list-users-of-an-organization).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewUser(ctx, "basic", &iam.UserArgs{
//				Email: pulumi.String("test@test.com"),
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
// IAM users can be imported using the `{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:iam/user:User basic 11111111-1111-1111-1111-111111111111
//
// ```
type User struct {
	pulumi.CustomResourceState

	// The ID of the account root user associated with the user.
	AccountRootUserId pulumi.StringOutput `pulumi:"accountRootUserId"`
	// The date and time of the creation of the iam user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether the iam user is deletable.
	Deletable pulumi.BoolOutput `pulumi:"deletable"`
	// The email of the IAM user.
	Email pulumi.StringOutput `pulumi:"email"`
	// The date of the last login.
	LastLoginAt pulumi.StringOutput `pulumi:"lastLoginAt"`
	// Whether the MFA is enabled.
	Mfa pulumi.BoolOutput `pulumi:"mfa"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status pulumi.StringOutput `pulumi:"status"`
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type pulumi.StringOutput `pulumi:"type"`
	// The date and time of the last update of the iam user.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("scaleway:iam/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("scaleway:iam/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The ID of the account root user associated with the user.
	AccountRootUserId *string `pulumi:"accountRootUserId"`
	// The date and time of the creation of the iam user.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether the iam user is deletable.
	Deletable *bool `pulumi:"deletable"`
	// The email of the IAM user.
	Email *string `pulumi:"email"`
	// The date of the last login.
	LastLoginAt *string `pulumi:"lastLoginAt"`
	// Whether the MFA is enabled.
	Mfa *bool `pulumi:"mfa"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status *string `pulumi:"status"`
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type *string `pulumi:"type"`
	// The date and time of the last update of the iam user.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type UserState struct {
	// The ID of the account root user associated with the user.
	AccountRootUserId pulumi.StringPtrInput
	// The date and time of the creation of the iam user.
	CreatedAt pulumi.StringPtrInput
	// Whether the iam user is deletable.
	Deletable pulumi.BoolPtrInput
	// The email of the IAM user.
	Email pulumi.StringPtrInput
	// The date of the last login.
	LastLoginAt pulumi.StringPtrInput
	// Whether the MFA is enabled.
	Mfa pulumi.BoolPtrInput
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringPtrInput
	// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Status pulumi.StringPtrInput
	// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
	Type pulumi.StringPtrInput
	// The date and time of the last update of the iam user.
	UpdatedAt pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The email of the IAM user.
	Email string `pulumi:"email"`
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId *string `pulumi:"organizationId"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The email of the IAM user.
	Email pulumi.StringInput
	// `organizationId`) The ID of the organization the user is associated with.
	OrganizationId pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// The ID of the account root user associated with the user.
func (o UserOutput) AccountRootUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.AccountRootUserId }).(pulumi.StringOutput)
}

// The date and time of the creation of the iam user.
func (o UserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Whether the iam user is deletable.
func (o UserOutput) Deletable() pulumi.BoolOutput {
	return o.ApplyT(func(v *User) pulumi.BoolOutput { return v.Deletable }).(pulumi.BoolOutput)
}

// The email of the IAM user.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The date of the last login.
func (o UserOutput) LastLoginAt() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.LastLoginAt }).(pulumi.StringOutput)
}

// Whether the MFA is enabled.
func (o UserOutput) Mfa() pulumi.BoolOutput {
	return o.ApplyT(func(v *User) pulumi.BoolOutput { return v.Mfa }).(pulumi.BoolOutput)
}

// `organizationId`) The ID of the organization the user is associated with.
func (o UserOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The status of user invitation. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
func (o UserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The type of user. Check the possible values in the [api doc](https://www.scaleway.com/en/developers/api/iam/#path-users-get-a-given-user).
func (o UserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date and time of the last update of the iam user.
func (o UserOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}