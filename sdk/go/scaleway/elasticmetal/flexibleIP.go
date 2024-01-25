// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway flexible IPs.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := elasticmetal.NewFlexibleIP(ctx, "main", &elasticmetal.FlexibleIPArgs{
//				Reverse: pulumi.String("my-reverse.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With zone
//
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
//			_, err := elasticmetal.NewFlexibleIP(ctx, "main", &elasticmetal.FlexibleIPArgs{
//				Zone: pulumi.String("fr-par-2"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With IPv6
//
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
//			_, err := elasticmetal.NewFlexibleIP(ctx, "main", &elasticmetal.FlexibleIPArgs{
//				IsIpv6: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With baremetal server
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/account"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/elasticmetal"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainSSHKey, err := account.NewSSHKey(ctx, "mainSSHKey", &account.SSHKeyArgs{
//				PublicKey: pulumi.String("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILHy/M5FVm5ydLGcal3e5LNcfTalbeN7QL/ZGCvDEdqJ foobar@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			byId, err := elasticmetal.GetBareMetalOS(ctx, &elasticmetal.GetBareMetalOSArgs{
//				Zone:    pulumi.StringRef("fr-par-2"),
//				Name:    pulumi.StringRef("Ubuntu"),
//				Version: pulumi.StringRef("20.04 LTS (Focal Fossa)"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			myOffer, err := elasticmetal.GetBareMetalOffer(ctx, &elasticmetal.GetBareMetalOfferArgs{
//				Zone: pulumi.StringRef("fr-par-2"),
//				Name: pulumi.StringRef("EM-A210R-HDD"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			base, err := elasticmetal.NewBareMetalServer(ctx, "base", &elasticmetal.BareMetalServerArgs{
//				Zone:      pulumi.String("fr-par-2"),
//				Offer:     *pulumi.String(myOffer.OfferId),
//				Os:        *pulumi.String(byId.OsId),
//				SshKeyIds: mainSSHKey.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticmetal.NewFlexibleIP(ctx, "mainFlexibleIP", &elasticmetal.FlexibleIPArgs{
//				ServerId: base.ID(),
//				Zone:     pulumi.String("fr-par-2"),
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
// Flexible IPs can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:elasticmetal/flexibleIP:FlexibleIP main fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type FlexibleIP struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the Flexible IP (Format ISO 8601).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A description of the flexible IP.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP address of the Flexible IP.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Defines whether the flexible IP has an IPv6 address.
	IsIpv6 pulumi.BoolPtrOutput `pulumi:"isIpv6"`
	// The organization of the Flexible IP.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The project of the Flexible IP.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The ID of the associated server.
	ServerId pulumi.StringPtrOutput `pulumi:"serverId"`
	// The status of the flexible IP.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of tags to apply to the flexible IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the Flexible IP (Format ISO 8601).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The zone of the Flexible IP.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFlexibleIP registers a new resource with the given unique name, arguments, and options.
func NewFlexibleIP(ctx *pulumi.Context,
	name string, args *FlexibleIPArgs, opts ...pulumi.ResourceOption) (*FlexibleIP, error) {
	if args == nil {
		args = &FlexibleIPArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlexibleIP
	err := ctx.RegisterResource("scaleway:elasticmetal/flexibleIP:FlexibleIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexibleIP gets an existing FlexibleIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexibleIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexibleIPState, opts ...pulumi.ResourceOption) (*FlexibleIP, error) {
	var resource FlexibleIP
	err := ctx.ReadResource("scaleway:elasticmetal/flexibleIP:FlexibleIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexibleIP resources.
type flexibleIPState struct {
	// The date and time of the creation of the Flexible IP (Format ISO 8601).
	CreatedAt *string `pulumi:"createdAt"`
	// A description of the flexible IP.
	Description *string `pulumi:"description"`
	// The IP address of the Flexible IP.
	IpAddress *string `pulumi:"ipAddress"`
	// Defines whether the flexible IP has an IPv6 address.
	IsIpv6 *bool `pulumi:"isIpv6"`
	// The organization of the Flexible IP.
	OrganizationId *string `pulumi:"organizationId"`
	// The project of the Flexible IP.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse *string `pulumi:"reverse"`
	// The ID of the associated server.
	ServerId *string `pulumi:"serverId"`
	// The status of the flexible IP.
	Status *string `pulumi:"status"`
	// A list of tags to apply to the flexible IP.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the Flexible IP (Format ISO 8601).
	UpdatedAt *string `pulumi:"updatedAt"`
	// The zone of the Flexible IP.
	Zone *string `pulumi:"zone"`
}

type FlexibleIPState struct {
	// The date and time of the creation of the Flexible IP (Format ISO 8601).
	CreatedAt pulumi.StringPtrInput
	// A description of the flexible IP.
	Description pulumi.StringPtrInput
	// The IP address of the Flexible IP.
	IpAddress pulumi.StringPtrInput
	// Defines whether the flexible IP has an IPv6 address.
	IsIpv6 pulumi.BoolPtrInput
	// The organization of the Flexible IP.
	OrganizationId pulumi.StringPtrInput
	// The project of the Flexible IP.
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringPtrInput
	// The ID of the associated server.
	ServerId pulumi.StringPtrInput
	// The status of the flexible IP.
	Status pulumi.StringPtrInput
	// A list of tags to apply to the flexible IP.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the Flexible IP (Format ISO 8601).
	UpdatedAt pulumi.StringPtrInput
	// The zone of the Flexible IP.
	Zone pulumi.StringPtrInput
}

func (FlexibleIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIPState)(nil)).Elem()
}

type flexibleIPArgs struct {
	// A description of the flexible IP.
	Description *string `pulumi:"description"`
	// Defines whether the flexible IP has an IPv6 address.
	IsIpv6 *bool `pulumi:"isIpv6"`
	// The project of the Flexible IP.
	ProjectId *string `pulumi:"projectId"`
	// The reverse domain associated with this flexible IP.
	Reverse *string `pulumi:"reverse"`
	// The ID of the associated server.
	ServerId *string `pulumi:"serverId"`
	// A list of tags to apply to the flexible IP.
	Tags []string `pulumi:"tags"`
	// The zone of the Flexible IP.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a FlexibleIP resource.
type FlexibleIPArgs struct {
	// A description of the flexible IP.
	Description pulumi.StringPtrInput
	// Defines whether the flexible IP has an IPv6 address.
	IsIpv6 pulumi.BoolPtrInput
	// The project of the Flexible IP.
	ProjectId pulumi.StringPtrInput
	// The reverse domain associated with this flexible IP.
	Reverse pulumi.StringPtrInput
	// The ID of the associated server.
	ServerId pulumi.StringPtrInput
	// A list of tags to apply to the flexible IP.
	Tags pulumi.StringArrayInput
	// The zone of the Flexible IP.
	Zone pulumi.StringPtrInput
}

func (FlexibleIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexibleIPArgs)(nil)).Elem()
}

type FlexibleIPInput interface {
	pulumi.Input

	ToFlexibleIPOutput() FlexibleIPOutput
	ToFlexibleIPOutputWithContext(ctx context.Context) FlexibleIPOutput
}

func (*FlexibleIP) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIP)(nil)).Elem()
}

func (i *FlexibleIP) ToFlexibleIPOutput() FlexibleIPOutput {
	return i.ToFlexibleIPOutputWithContext(context.Background())
}

func (i *FlexibleIP) ToFlexibleIPOutputWithContext(ctx context.Context) FlexibleIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIPOutput)
}

// FlexibleIPArrayInput is an input type that accepts FlexibleIPArray and FlexibleIPArrayOutput values.
// You can construct a concrete instance of `FlexibleIPArrayInput` via:
//
//	FlexibleIPArray{ FlexibleIPArgs{...} }
type FlexibleIPArrayInput interface {
	pulumi.Input

	ToFlexibleIPArrayOutput() FlexibleIPArrayOutput
	ToFlexibleIPArrayOutputWithContext(context.Context) FlexibleIPArrayOutput
}

type FlexibleIPArray []FlexibleIPInput

func (FlexibleIPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIP)(nil)).Elem()
}

func (i FlexibleIPArray) ToFlexibleIPArrayOutput() FlexibleIPArrayOutput {
	return i.ToFlexibleIPArrayOutputWithContext(context.Background())
}

func (i FlexibleIPArray) ToFlexibleIPArrayOutputWithContext(ctx context.Context) FlexibleIPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIPArrayOutput)
}

// FlexibleIPMapInput is an input type that accepts FlexibleIPMap and FlexibleIPMapOutput values.
// You can construct a concrete instance of `FlexibleIPMapInput` via:
//
//	FlexibleIPMap{ "key": FlexibleIPArgs{...} }
type FlexibleIPMapInput interface {
	pulumi.Input

	ToFlexibleIPMapOutput() FlexibleIPMapOutput
	ToFlexibleIPMapOutputWithContext(context.Context) FlexibleIPMapOutput
}

type FlexibleIPMap map[string]FlexibleIPInput

func (FlexibleIPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIP)(nil)).Elem()
}

func (i FlexibleIPMap) ToFlexibleIPMapOutput() FlexibleIPMapOutput {
	return i.ToFlexibleIPMapOutputWithContext(context.Background())
}

func (i FlexibleIPMap) ToFlexibleIPMapOutputWithContext(ctx context.Context) FlexibleIPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexibleIPMapOutput)
}

type FlexibleIPOutput struct{ *pulumi.OutputState }

func (FlexibleIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexibleIP)(nil)).Elem()
}

func (o FlexibleIPOutput) ToFlexibleIPOutput() FlexibleIPOutput {
	return o
}

func (o FlexibleIPOutput) ToFlexibleIPOutputWithContext(ctx context.Context) FlexibleIPOutput {
	return o
}

// The date and time of the creation of the Flexible IP (Format ISO 8601).
func (o FlexibleIPOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description of the flexible IP.
func (o FlexibleIPOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IP address of the Flexible IP.
func (o FlexibleIPOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Defines whether the flexible IP has an IPv6 address.
func (o FlexibleIPOutput) IsIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.BoolPtrOutput { return v.IsIpv6 }).(pulumi.BoolPtrOutput)
}

// The organization of the Flexible IP.
func (o FlexibleIPOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The project of the Flexible IP.
func (o FlexibleIPOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse domain associated with this flexible IP.
func (o FlexibleIPOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The ID of the associated server.
func (o FlexibleIPOutput) ServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringPtrOutput { return v.ServerId }).(pulumi.StringPtrOutput)
}

// The status of the flexible IP.
func (o FlexibleIPOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A list of tags to apply to the flexible IP.
func (o FlexibleIPOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the Flexible IP (Format ISO 8601).
func (o FlexibleIPOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The zone of the Flexible IP.
func (o FlexibleIPOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *FlexibleIP) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type FlexibleIPArrayOutput struct{ *pulumi.OutputState }

func (FlexibleIPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FlexibleIP)(nil)).Elem()
}

func (o FlexibleIPArrayOutput) ToFlexibleIPArrayOutput() FlexibleIPArrayOutput {
	return o
}

func (o FlexibleIPArrayOutput) ToFlexibleIPArrayOutputWithContext(ctx context.Context) FlexibleIPArrayOutput {
	return o
}

func (o FlexibleIPArrayOutput) Index(i pulumi.IntInput) FlexibleIPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FlexibleIP {
		return vs[0].([]*FlexibleIP)[vs[1].(int)]
	}).(FlexibleIPOutput)
}

type FlexibleIPMapOutput struct{ *pulumi.OutputState }

func (FlexibleIPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FlexibleIP)(nil)).Elem()
}

func (o FlexibleIPMapOutput) ToFlexibleIPMapOutput() FlexibleIPMapOutput {
	return o
}

func (o FlexibleIPMapOutput) ToFlexibleIPMapOutputWithContext(ctx context.Context) FlexibleIPMapOutput {
	return o
}

func (o FlexibleIPMapOutput) MapIndex(k pulumi.StringInput) FlexibleIPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FlexibleIP {
		return vs[0].(map[string]*FlexibleIP)[vs[1].(string)]
	}).(FlexibleIPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIPInput)(nil)).Elem(), &FlexibleIP{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIPArrayInput)(nil)).Elem(), FlexibleIPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlexibleIPMapInput)(nil)).Elem(), FlexibleIPMap{})
	pulumi.RegisterOutputType(FlexibleIPOutput{})
	pulumi.RegisterOutputType(FlexibleIPArrayOutput{})
	pulumi.RegisterOutputType(FlexibleIPMapOutput{})
}
