// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Domain zone.\
// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/configure-dns-zones/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/dns"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewZone(ctx, "test", &dns.ZoneArgs{
//				Domain:    pulumi.String("scaleway-terraform.com"),
//				Subdomain: pulumi.String("test"),
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
// Zone can be imported using the `{subdomain}.{domain}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:dns/zone:Zone test test.scaleway-terraform.com
//
// ```
type Zone struct {
	pulumi.CustomResourceState

	// The domain where the DNS zone will be created.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Message
	Message pulumi.StringOutput `pulumi:"message"`
	// NameServer list for zone.
	Ns pulumi.StringArrayOutput `pulumi:"ns"`
	// NameServer default list for zone.
	NsDefaults pulumi.StringArrayOutput `pulumi:"nsDefaults"`
	// NameServer master list for zone.
	NsMasters pulumi.StringArrayOutput `pulumi:"nsMasters"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The domain zone status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringOutput `pulumi:"subdomain"`
	// The date and time of the last update of the DNS zone.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Subdomain == nil {
		return nil, errors.New("invalid value for required argument 'Subdomain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Zone
	err := ctx.RegisterResource("scaleway:dns/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("scaleway:dns/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// The domain where the DNS zone will be created.
	Domain *string `pulumi:"domain"`
	// Message
	Message *string `pulumi:"message"`
	// NameServer list for zone.
	Ns []string `pulumi:"ns"`
	// NameServer default list for zone.
	NsDefaults []string `pulumi:"nsDefaults"`
	// NameServer master list for zone.
	NsMasters []string `pulumi:"nsMasters"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The domain zone status.
	Status *string `pulumi:"status"`
	// The subdomain(zone name) to create in the domain.
	Subdomain *string `pulumi:"subdomain"`
	// The date and time of the last update of the DNS zone.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ZoneState struct {
	// The domain where the DNS zone will be created.
	Domain pulumi.StringPtrInput
	// Message
	Message pulumi.StringPtrInput
	// NameServer list for zone.
	Ns pulumi.StringArrayInput
	// NameServer default list for zone.
	NsDefaults pulumi.StringArrayInput
	// NameServer master list for zone.
	NsMasters pulumi.StringArrayInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// The domain zone status.
	Status pulumi.StringPtrInput
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringPtrInput
	// The date and time of the last update of the DNS zone.
	UpdatedAt pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// The domain where the DNS zone will be created.
	Domain string `pulumi:"domain"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The subdomain(zone name) to create in the domain.
	Subdomain string `pulumi:"subdomain"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The domain where the DNS zone will be created.
	Domain pulumi.StringInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//	ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//	ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

// The domain where the DNS zone will be created.
func (o ZoneOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Message
func (o ZoneOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// NameServer list for zone.
func (o ZoneOutput) Ns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.Ns }).(pulumi.StringArrayOutput)
}

// NameServer default list for zone.
func (o ZoneOutput) NsDefaults() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.NsDefaults }).(pulumi.StringArrayOutput)
}

// NameServer master list for zone.
func (o ZoneOutput) NsMasters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.NsMasters }).(pulumi.StringArrayOutput)
}

// `projectId`) The ID of the project the domain is associated with.
func (o ZoneOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The domain zone status.
func (o ZoneOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The subdomain(zone name) to create in the domain.
func (o ZoneOutput) Subdomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Subdomain }).(pulumi.StringOutput)
}

// The date and time of the last update of the DNS zone.
func (o ZoneOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}