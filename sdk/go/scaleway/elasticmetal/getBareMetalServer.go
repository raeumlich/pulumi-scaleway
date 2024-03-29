// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a baremetal server.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := elasticmetal.LookupBareMetalServer(ctx, &elasticmetal.LookupBareMetalServerArgs{
//				Name: pulumi.StringRef("foobar"),
//				Zone: pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = elasticmetal.LookupBareMetalServer(ctx, &elasticmetal.LookupBareMetalServerArgs{
//				ServerId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
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
func LookupBareMetalServer(ctx *pulumi.Context, args *LookupBareMetalServerArgs, opts ...pulumi.InvokeOption) (*LookupBareMetalServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBareMetalServerResult
	err := ctx.Invoke("scaleway:elasticmetal/getBareMetalServer:getBareMetalServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBareMetalServer.
type LookupBareMetalServerArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the baremetal server is associated with.
	ProjectId *string `pulumi:"projectId"`
	ServerId  *string `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getBareMetalServer.
type LookupBareMetalServerResult struct {
	Description string `pulumi:"description"`
	Domain      string `pulumi:"domain"`
	Hostname    string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string                             `pulumi:"id"`
	InstallConfigAfterward   bool                               `pulumi:"installConfigAfterward"`
	Ips                      []GetBareMetalServerIp             `pulumi:"ips"`
	Ipv4s                    []GetBareMetalServerIpv4           `pulumi:"ipv4s"`
	Ipv6s                    []GetBareMetalServerIpv6           `pulumi:"ipv6s"`
	Name                     *string                            `pulumi:"name"`
	Offer                    string                             `pulumi:"offer"`
	OfferId                  string                             `pulumi:"offerId"`
	OfferName                string                             `pulumi:"offerName"`
	Options                  []GetBareMetalServerOption         `pulumi:"options"`
	OrganizationId           string                             `pulumi:"organizationId"`
	Os                       string                             `pulumi:"os"`
	OsName                   string                             `pulumi:"osName"`
	Password                 string                             `pulumi:"password"`
	PrivateNetworks          []GetBareMetalServerPrivateNetwork `pulumi:"privateNetworks"`
	ProjectId                *string                            `pulumi:"projectId"`
	ReinstallOnConfigChanges bool                               `pulumi:"reinstallOnConfigChanges"`
	ServerId                 *string                            `pulumi:"serverId"`
	ServicePassword          string                             `pulumi:"servicePassword"`
	ServiceUser              string                             `pulumi:"serviceUser"`
	SshKeyIds                []string                           `pulumi:"sshKeyIds"`
	Tags                     []string                           `pulumi:"tags"`
	User                     string                             `pulumi:"user"`
	Zone                     *string                            `pulumi:"zone"`
}

func LookupBareMetalServerOutput(ctx *pulumi.Context, args LookupBareMetalServerOutputArgs, opts ...pulumi.InvokeOption) LookupBareMetalServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBareMetalServerResult, error) {
			args := v.(LookupBareMetalServerArgs)
			r, err := LookupBareMetalServer(ctx, &args, opts...)
			var s LookupBareMetalServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBareMetalServerResultOutput)
}

// A collection of arguments for invoking getBareMetalServer.
type LookupBareMetalServerOutputArgs struct {
	// The server name. Only one of `name` and `serverId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the baremetal server is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	ServerId  pulumi.StringPtrInput `pulumi:"serverId"`
	// `zone`) The zone in which the server exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupBareMetalServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalServerArgs)(nil)).Elem()
}

// A collection of values returned by getBareMetalServer.
type LookupBareMetalServerResultOutput struct{ *pulumi.OutputState }

func (LookupBareMetalServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalServerResult)(nil)).Elem()
}

func (o LookupBareMetalServerResultOutput) ToLookupBareMetalServerResultOutput() LookupBareMetalServerResultOutput {
	return o
}

func (o LookupBareMetalServerResultOutput) ToLookupBareMetalServerResultOutputWithContext(ctx context.Context) LookupBareMetalServerResultOutput {
	return o
}

func (o LookupBareMetalServerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBareMetalServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) InstallConfigAfterward() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) bool { return v.InstallConfigAfterward }).(pulumi.BoolOutput)
}

func (o LookupBareMetalServerResultOutput) Ips() GetBareMetalServerIpArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []GetBareMetalServerIp { return v.Ips }).(GetBareMetalServerIpArrayOutput)
}

func (o LookupBareMetalServerResultOutput) Ipv4s() GetBareMetalServerIpv4ArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []GetBareMetalServerIpv4 { return v.Ipv4s }).(GetBareMetalServerIpv4ArrayOutput)
}

func (o LookupBareMetalServerResultOutput) Ipv6s() GetBareMetalServerIpv6ArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []GetBareMetalServerIpv6 { return v.Ipv6s }).(GetBareMetalServerIpv6ArrayOutput)
}

func (o LookupBareMetalServerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupBareMetalServerResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Offer }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.OfferId }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) OfferName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.OfferName }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Options() GetBareMetalServerOptionArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []GetBareMetalServerOption { return v.Options }).(GetBareMetalServerOptionArrayOutput)
}

func (o LookupBareMetalServerResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Os }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.OsName }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) PrivateNetworks() GetBareMetalServerPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []GetBareMetalServerPrivateNetwork { return v.PrivateNetworks }).(GetBareMetalServerPrivateNetworkArrayOutput)
}

func (o LookupBareMetalServerResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupBareMetalServerResultOutput) ReinstallOnConfigChanges() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) bool { return v.ReinstallOnConfigChanges }).(pulumi.BoolOutput)
}

func (o LookupBareMetalServerResultOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) *string { return v.ServerId }).(pulumi.StringPtrOutput)
}

func (o LookupBareMetalServerResultOutput) ServicePassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.ServicePassword }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) ServiceUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.ServiceUser }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) SshKeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []string { return v.SshKeyIds }).(pulumi.StringArrayOutput)
}

func (o LookupBareMetalServerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupBareMetalServerResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) string { return v.User }).(pulumi.StringOutput)
}

func (o LookupBareMetalServerResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBareMetalServerResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBareMetalServerResultOutput{})
}
