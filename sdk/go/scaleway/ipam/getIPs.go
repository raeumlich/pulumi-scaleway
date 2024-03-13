// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about multiple IPs managed by IPAM service.
//
// ## Examples
//
// ### By tag
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/ipam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ipam.GetIPs(ctx, &ipam.GetIPsArgs{
//				Tags: []string{
//					"tag",
//				},
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
//
// ### By type and resource
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/ipam"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/redis"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpc01, err := vpc.NewVPC(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", &vpc.PrivateNetworkArgs{
//				VpcId: vpc01.ID(),
//				Ipv4Subnet: &vpc.PrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.32.0/22"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			redis01, err := redis.NewCluster(ctx, "redis01", &redis.ClusterArgs{
//				Version:     pulumi.String("7.0.5"),
//				NodeType:    pulumi.String("RED1-XS"),
//				UserName:    pulumi.String("my_initial_user"),
//				Password:    pulumi.String("thiZ_is_v&ry_s3cret"),
//				ClusterSize: pulumi.Int(3),
//				PrivateNetworks: redis.ClusterPrivateNetworkArray{
//					&redis.ClusterPrivateNetworkArgs{
//						Id: pn01.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = ipam.GetIPsOutput(ctx, ipam.GetIPsOutputArgs{
//				Type: pulumi.String("ipv4"),
//				Resource: &ipam.GetIPsResourceArgs{
//					Id:   redis01.ID(),
//					Type: pulumi.String("redis_cluster"),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetIPs(ctx *pulumi.Context, args *GetIPsArgs, opts ...pulumi.InvokeOption) (*GetIPsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIPsResult
	err := ctx.Invoke("scaleway:ipam/getIPs:getIPs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIPs.
type GetIPsArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource.
	Attached *bool `pulumi:"attached"`
	// The Mac Address used as filter.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network used as filter.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the project used as filter.
	ProjectId *string `pulumi:"projectId"`
	// The region used as filter.
	Region *string `pulumi:"region"`
	// Filter by resource ID, type or name.
	Resource *GetIPsResource `pulumi:"resource"`
	// The tags used as filter.
	Tags []string `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type *string `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal *string `pulumi:"zonal"`
}

// A collection of values returned by getIPs.
type GetIPsResult struct {
	Attached *bool `pulumi:"attached"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of found IPs
	Ips []GetIPsIp `pulumi:"ips"`
	// The mac address.
	MacAddress       *string `pulumi:"macAddress"`
	OrganizationId   string  `pulumi:"organizationId"`
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The ID of the project the server is associated with.
	ProjectId string `pulumi:"projectId"`
	// The region in which the IP is.
	Region string `pulumi:"region"`
	// The list of public IPs of the server.
	Resource *GetIPsResource `pulumi:"resource"`
	// The tags associated with the IP.
	Tags []string `pulumi:"tags"`
	// The type of resource.
	Type  *string `pulumi:"type"`
	Zonal string  `pulumi:"zonal"`
}

func GetIPsOutput(ctx *pulumi.Context, args GetIPsOutputArgs, opts ...pulumi.InvokeOption) GetIPsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIPsResult, error) {
			args := v.(GetIPsArgs)
			r, err := GetIPs(ctx, &args, opts...)
			var s GetIPsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIPsResultOutput)
}

// A collection of arguments for invoking getIPs.
type GetIPsOutputArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource.
	Attached pulumi.BoolPtrInput `pulumi:"attached"`
	// The Mac Address used as filter.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the private network used as filter.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// The ID of the project used as filter.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region used as filter.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Filter by resource ID, type or name.
	Resource GetIPsResourcePtrInput `pulumi:"resource"`
	// The tags used as filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal pulumi.StringPtrInput `pulumi:"zonal"`
}

func (GetIPsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIPsArgs)(nil)).Elem()
}

// A collection of values returned by getIPs.
type GetIPsResultOutput struct{ *pulumi.OutputState }

func (GetIPsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIPsResult)(nil)).Elem()
}

func (o GetIPsResultOutput) ToGetIPsResultOutput() GetIPsResultOutput {
	return o
}

func (o GetIPsResultOutput) ToGetIPsResultOutputWithContext(ctx context.Context) GetIPsResultOutput {
	return o
}

func (o GetIPsResultOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetIPsResult) *bool { return v.Attached }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIPsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of found IPs
func (o GetIPsResultOutput) Ips() GetIPsIpArrayOutput {
	return o.ApplyT(func(v GetIPsResult) []GetIPsIp { return v.Ips }).(GetIPsIpArrayOutput)
}

// The mac address.
func (o GetIPsResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIPsResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o GetIPsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GetIPsResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIPsResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

// The ID of the project the server is associated with.
func (o GetIPsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which the IP is.
func (o GetIPsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.Region }).(pulumi.StringOutput)
}

// The list of public IPs of the server.
func (o GetIPsResultOutput) Resource() GetIPsResourcePtrOutput {
	return o.ApplyT(func(v GetIPsResult) *GetIPsResource { return v.Resource }).(GetIPsResourcePtrOutput)
}

// The tags associated with the IP.
func (o GetIPsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIPsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of resource.
func (o GetIPsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIPsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetIPsResultOutput) Zonal() pulumi.StringOutput {
	return o.ApplyT(func(v GetIPsResult) string { return v.Zonal }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIPsResultOutput{})
}
