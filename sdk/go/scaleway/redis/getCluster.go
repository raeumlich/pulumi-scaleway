// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Redis cluster. For further information check our [api documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816)
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redis.LookupCluster(ctx, &redis.LookupClusterArgs{
//				ClusterId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
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
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("scaleway:redis/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The Redis cluster ID.
	// Only one of the `name` and `clusterId` should be specified.
	ClusterId *string `pulumi:"clusterId"`
	// The name of the Redis cluster.
	// Only one of the `name` and `clusterId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the Redis cluster is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The zone in which the server exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	Acls        []GetClusterAcl `pulumi:"acls"`
	Certificate string          `pulumi:"certificate"`
	ClusterId   *string         `pulumi:"clusterId"`
	ClusterSize int             `pulumi:"clusterSize"`
	CreatedAt   string          `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                     `pulumi:"id"`
	Name            *string                    `pulumi:"name"`
	NodeType        string                     `pulumi:"nodeType"`
	Password        string                     `pulumi:"password"`
	PrivateNetworks []GetClusterPrivateNetwork `pulumi:"privateNetworks"`
	ProjectId       *string                    `pulumi:"projectId"`
	PublicNetworks  []GetClusterPublicNetwork  `pulumi:"publicNetworks"`
	Settings        map[string]string          `pulumi:"settings"`
	Tags            []string                   `pulumi:"tags"`
	TlsEnabled      bool                       `pulumi:"tlsEnabled"`
	UpdatedAt       string                     `pulumi:"updatedAt"`
	UserName        string                     `pulumi:"userName"`
	Version         string                     `pulumi:"version"`
	Zone            *string                    `pulumi:"zone"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// The Redis cluster ID.
	// Only one of the `name` and `clusterId` should be specified.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The name of the Redis cluster.
	// Only one of the `name` and `clusterId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the Redis cluster is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The zone in which the server exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) Acls() GetClusterAclArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterAcl { return v.Acls }).(GetClusterAclArrayOutput)
}

func (o LookupClusterResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) PrivateNetworks() GetClusterPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterPrivateNetwork { return v.PrivateNetworks }).(GetClusterPrivateNetworkArrayOutput)
}

func (o LookupClusterResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) PublicNetworks() GetClusterPublicNetworkArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterPublicNetwork { return v.PublicNetworks }).(GetClusterPublicNetworkArrayOutput)
}

func (o LookupClusterResultOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) TlsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.TlsEnabled }).(pulumi.BoolOutput)
}

func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
