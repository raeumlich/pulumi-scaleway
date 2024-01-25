// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Use this data source to get information on an existing IAM user based on its ID or email address.
// For more information,
// see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#users-06bdcf).
//
// ## Example Usage
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
//			_, err := iam.LookupUser(ctx, &iam.LookupUserArgs{
//				UserId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.LookupUser(ctx, &iam.LookupUserArgs{
//				Email: pulumi.StringRef("foo@bar.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("scaleway:iam/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The email address of the IAM user. Only one of the `email` and `userId` should be specified.
	Email *string `pulumi:"email"`
	// `organizationId`) The ID of the
	// organization the user is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The ID of the IAM user. Only one of the `email` and `userId` should be specified.
	UserId *string `pulumi:"userId"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	Email *string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	OrganizationId *string `pulumi:"organizationId"`
	UserId         *string `pulumi:"userId"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The email address of the IAM user. Only one of the `email` and `userId` should be specified.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// `organizationId`) The ID of the
	// organization the user is associated with.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// The ID of the IAM user. Only one of the `email` and `userId` should be specified.
	UserId pulumi.StringPtrInput `pulumi:"userId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}