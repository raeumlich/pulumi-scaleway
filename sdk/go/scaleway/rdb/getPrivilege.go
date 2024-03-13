// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about the privilege on RDB database.
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/rdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rdb.LookupPrivilege(ctx, &rdb.LookupPrivilegeArgs{
//				DatabaseName: "my-database",
//				InstanceId:   "11111111-1111-111111111111",
//				UserName:     "my-user",
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
func LookupPrivilege(ctx *pulumi.Context, args *LookupPrivilegeArgs, opts ...pulumi.InvokeOption) (*LookupPrivilegeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivilegeResult
	err := ctx.Invoke("scaleway:rdb/getPrivilege:getPrivilege", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivilege.
type LookupPrivilegeArgs struct {
	// The database name.
	DatabaseName string `pulumi:"databaseName"`
	// The RDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// The user name.
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getPrivilege.
type LookupPrivilegeResult struct {
	DatabaseName string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
	// , `custom` and `none`.
	Permission string  `pulumi:"permission"`
	Region     *string `pulumi:"region"`
	UserName   string  `pulumi:"userName"`
}

func LookupPrivilegeOutput(ctx *pulumi.Context, args LookupPrivilegeOutputArgs, opts ...pulumi.InvokeOption) LookupPrivilegeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivilegeResult, error) {
			args := v.(LookupPrivilegeArgs)
			r, err := LookupPrivilege(ctx, &args, opts...)
			var s LookupPrivilegeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivilegeResultOutput)
}

// A collection of arguments for invoking getPrivilege.
type LookupPrivilegeOutputArgs struct {
	// The database name.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The RDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The user name.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupPrivilegeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivilegeArgs)(nil)).Elem()
}

// A collection of values returned by getPrivilege.
type LookupPrivilegeResultOutput struct{ *pulumi.OutputState }

func (LookupPrivilegeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivilegeResult)(nil)).Elem()
}

func (o LookupPrivilegeResultOutput) ToLookupPrivilegeResultOutput() LookupPrivilegeResultOutput {
	return o
}

func (o LookupPrivilegeResultOutput) ToLookupPrivilegeResultOutputWithContext(ctx context.Context) LookupPrivilegeResultOutput {
	return o
}

func (o LookupPrivilegeResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivilegeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivilegeResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The permission for this user on the database. Possible values are `readonly`, `readwrite`, `all`
// , `custom` and `none`.
func (o LookupPrivilegeResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) string { return v.Permission }).(pulumi.StringOutput)
}

func (o LookupPrivilegeResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupPrivilegeResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivilegeResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivilegeResultOutput{})
}
