// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ipam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about IP managed by IPAM service. IPAM service is used for dhcp bundled in VPCs' private networks.
//
// ## Examples
func LookupIP(ctx *pulumi.Context, args *LookupIPArgs, opts ...pulumi.InvokeOption) (*LookupIPResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIPResult
	err := ctx.Invoke("scaleway:ipam/getIP:getIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIP.
type LookupIPArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipamIpId`.
	Attached *bool `pulumi:"attached"`
	// The IPAM IP ID. Cannot be used with the rest of the arguments.
	IpamIpId *string `pulumi:"ipamIpId"`
	// The Mac Address linked to the IP. Cannot be used with `ipamIpId`.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network the IP belong to. Cannot be used with `ipamIpId`.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the IP exists.
	Region *string `pulumi:"region"`
	// Filter by resource ID, type or name. Cannot be used with `ipamIpId`.
	// If specified, `type` is required, and at least one of `id` or `name` must be set.
	Resource *GetIPResource `pulumi:"resource"`
	// The tags associated with the IP. Cannot be used with `ipamIpId`.
	// As datasource only returns one IP, the search with given tags must return only one result.
	Tags []string `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type *string `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal *string `pulumi:"zonal"`
}

// A collection of values returned by getIP.
type LookupIPResult struct {
	// The IP address.
	Address string `pulumi:"address"`
	// the IP address with a CIDR notation.
	AddressCidr string `pulumi:"addressCidr"`
	Attached    *bool  `pulumi:"attached"`
	// The provider-assigned unique ID for this managed resource.
	Id               string         `pulumi:"id"`
	IpamIpId         *string        `pulumi:"ipamIpId"`
	MacAddress       *string        `pulumi:"macAddress"`
	OrganizationId   string         `pulumi:"organizationId"`
	PrivateNetworkId *string        `pulumi:"privateNetworkId"`
	ProjectId        string         `pulumi:"projectId"`
	Region           string         `pulumi:"region"`
	Resource         *GetIPResource `pulumi:"resource"`
	Tags             []string       `pulumi:"tags"`
	Type             *string        `pulumi:"type"`
	Zonal            string         `pulumi:"zonal"`
}

func LookupIPOutput(ctx *pulumi.Context, args LookupIPOutputArgs, opts ...pulumi.InvokeOption) LookupIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPResult, error) {
			args := v.(LookupIPArgs)
			r, err := LookupIP(ctx, &args, opts...)
			var s LookupIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPResultOutput)
}

// A collection of arguments for invoking getIP.
type LookupIPOutputArgs struct {
	// Defines whether to filter only for IPs which are attached to a resource. Cannot be used with `ipamIpId`.
	Attached pulumi.BoolPtrInput `pulumi:"attached"`
	// The IPAM IP ID. Cannot be used with the rest of the arguments.
	IpamIpId pulumi.StringPtrInput `pulumi:"ipamIpId"`
	// The Mac Address linked to the IP. Cannot be used with `ipamIpId`.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the private network the IP belong to. Cannot be used with `ipamIpId`.
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the IP exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Filter by resource ID, type or name. Cannot be used with `ipamIpId`.
	// If specified, `type` is required, and at least one of `id` or `name` must be set.
	Resource GetIPResourcePtrInput `pulumi:"resource"`
	// The tags associated with the IP. Cannot be used with `ipamIpId`.
	// As datasource only returns one IP, the search with given tags must return only one result.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1#pkg-constants) with type list.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Only IPs that are zonal, and in this zone, will be returned.
	Zonal pulumi.StringPtrInput `pulumi:"zonal"`
}

func (LookupIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPArgs)(nil)).Elem()
}

// A collection of values returned by getIP.
type LookupIPResultOutput struct{ *pulumi.OutputState }

func (LookupIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPResult)(nil)).Elem()
}

func (o LookupIPResultOutput) ToLookupIPResultOutput() LookupIPResultOutput {
	return o
}

func (o LookupIPResultOutput) ToLookupIPResultOutputWithContext(ctx context.Context) LookupIPResultOutput {
	return o
}

// The IP address.
func (o LookupIPResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Address }).(pulumi.StringOutput)
}

// the IP address with a CIDR notation.
func (o LookupIPResultOutput) AddressCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.AddressCidr }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *bool { return v.Attached }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIPResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) IpamIpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.IpamIpId }).(pulumi.StringPtrOutput)
}

func (o LookupIPResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupIPResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupIPResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupIPResultOutput) Resource() GetIPResourcePtrOutput {
	return o.ApplyT(func(v LookupIPResult) *GetIPResource { return v.Resource }).(GetIPResourcePtrOutput)
}

func (o LookupIPResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIPResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupIPResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupIPResultOutput) Zonal() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPResult) string { return v.Zonal }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPResultOutput{})
}
