// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package account

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Use this data source to get SSH key information based on its ID or name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/account"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := account.LookupSSHKey(ctx, &account.LookupSSHKeyArgs{
//				SshKeyId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSSHKey(ctx *pulumi.Context, args *LookupSSHKeyArgs, opts ...pulumi.InvokeOption) (*LookupSSHKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSSHKeyResult
	err := ctx.Invoke("scaleway:account/getSSHKey:getSSHKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSSHKey.
type LookupSSHKeyArgs struct {
	// The SSH key name. Only one of `name` and `sshKeyId` should be specified.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the SSH key is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The SSH key id. Only one of `name` and `sshKeyId` should be specified.
	SshKeyId *string `pulumi:"sshKeyId"`
}

// A collection of values returned by getSSHKey.
type LookupSSHKeyResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Disabled    bool   `pulumi:"disabled"`
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The ID of the organization the SSH key is associated with.
	OrganizationId string  `pulumi:"organizationId"`
	ProjectId      *string `pulumi:"projectId"`
	// The SSH public key string
	PublicKey string  `pulumi:"publicKey"`
	SshKeyId  *string `pulumi:"sshKeyId"`
	UpdatedAt string  `pulumi:"updatedAt"`
}

func LookupSSHKeyOutput(ctx *pulumi.Context, args LookupSSHKeyOutputArgs, opts ...pulumi.InvokeOption) LookupSSHKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSSHKeyResult, error) {
			args := v.(LookupSSHKeyArgs)
			r, err := LookupSSHKey(ctx, &args, opts...)
			var s LookupSSHKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSSHKeyResultOutput)
}

// A collection of arguments for invoking getSSHKey.
type LookupSSHKeyOutputArgs struct {
	// The SSH key name. Only one of `name` and `sshKeyId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `projectId`) The ID of the project the SSH key is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The SSH key id. Only one of `name` and `sshKeyId` should be specified.
	SshKeyId pulumi.StringPtrInput `pulumi:"sshKeyId"`
}

func (LookupSSHKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSSHKeyArgs)(nil)).Elem()
}

// A collection of values returned by getSSHKey.
type LookupSSHKeyResultOutput struct{ *pulumi.OutputState }

func (LookupSSHKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSSHKeyResult)(nil)).Elem()
}

func (o LookupSSHKeyResultOutput) ToLookupSSHKeyResultOutput() LookupSSHKeyResultOutput {
	return o
}

func (o LookupSSHKeyResultOutput) ToLookupSSHKeyResultOutputWithContext(ctx context.Context) LookupSSHKeyResultOutput {
	return o
}

func (o LookupSSHKeyResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupSSHKeyResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupSSHKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSSHKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSSHKeyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the SSH key is associated with.
func (o LookupSSHKeyResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupSSHKeyResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The SSH public key string
func (o LookupSSHKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupSSHKeyResultOutput) SshKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) *string { return v.SshKeyId }).(pulumi.StringPtrOutput)
}

func (o LookupSSHKeyResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSSHKeyResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSSHKeyResultOutput{})
}
