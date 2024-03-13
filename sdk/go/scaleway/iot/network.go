// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/iot"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainHub, err := iot.NewHub(ctx, "mainHub", &iot.HubArgs{
//				ProductPlan: pulumi.String("plan_shared"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iot.NewNetwork(ctx, "mainNetwork", &iot.NetworkArgs{
//				HubId: mainHub.ID(),
//				Type:  pulumi.String("sigfox"),
//			})
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
// ## Import
//
// IoT Networks can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:iot/network:Network net01 fr-par/11111111-1111-1111-1111-111111111111
// ```
type Network struct {
	pulumi.CustomResourceState

	// The date and time the Network was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The endpoint to use when interacting with the network.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringOutput `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The endpoint key to keep secret.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrOutput `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubId == nil {
		return nil, errors.New("invalid value for required argument 'HubId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Network
	err := ctx.RegisterResource("scaleway:iot/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("scaleway:iot/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// The date and time the Network was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The endpoint to use when interacting with the network.
	Endpoint *string `pulumi:"endpoint"`
	// The hub ID to which the Network will be attached to.
	HubId *string `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name *string `pulumi:"name"`
	// The endpoint key to keep secret.
	Secret *string `pulumi:"secret"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix *string `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type *string `pulumi:"type"`
}

type NetworkState struct {
	// The date and time the Network was created.
	CreatedAt pulumi.StringPtrInput
	// The endpoint to use when interacting with the network.
	Endpoint pulumi.StringPtrInput
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringPtrInput
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringPtrInput
	// The endpoint key to keep secret.
	Secret pulumi.StringPtrInput
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrInput
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// The hub ID to which the Network will be attached to.
	HubId string `pulumi:"hubId"`
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name *string `pulumi:"name"`
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix *string `pulumi:"topicPrefix"`
	// The network type to create (e.g. `sigfox`).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The hub ID to which the Network will be attached to.
	HubId pulumi.StringInput
	// The name of the IoT Network you want to create (e.g. `my-net`).
	Name pulumi.StringPtrInput
	// The prefix that will be prepended to all topics for this Network.
	TopicPrefix pulumi.StringPtrInput
	// The network type to create (e.g. `sigfox`).
	Type pulumi.StringInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

// NetworkArrayInput is an input type that accepts NetworkArray and NetworkArrayOutput values.
// You can construct a concrete instance of `NetworkArrayInput` via:
//
//	NetworkArray{ NetworkArgs{...} }
type NetworkArrayInput interface {
	pulumi.Input

	ToNetworkArrayOutput() NetworkArrayOutput
	ToNetworkArrayOutputWithContext(context.Context) NetworkArrayOutput
}

type NetworkArray []NetworkInput

func (NetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network)(nil)).Elem()
}

func (i NetworkArray) ToNetworkArrayOutput() NetworkArrayOutput {
	return i.ToNetworkArrayOutputWithContext(context.Background())
}

func (i NetworkArray) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkArrayOutput)
}

// NetworkMapInput is an input type that accepts NetworkMap and NetworkMapOutput values.
// You can construct a concrete instance of `NetworkMapInput` via:
//
//	NetworkMap{ "key": NetworkArgs{...} }
type NetworkMapInput interface {
	pulumi.Input

	ToNetworkMapOutput() NetworkMapOutput
	ToNetworkMapOutputWithContext(context.Context) NetworkMapOutput
}

type NetworkMap map[string]NetworkInput

func (NetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network)(nil)).Elem()
}

func (i NetworkMap) ToNetworkMapOutput() NetworkMapOutput {
	return i.ToNetworkMapOutputWithContext(context.Background())
}

func (i NetworkMap) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMapOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

// The date and time the Network was created.
func (o NetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The endpoint to use when interacting with the network.
func (o NetworkOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The hub ID to which the Network will be attached to.
func (o NetworkOutput) HubId() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.HubId }).(pulumi.StringOutput)
}

// The name of the IoT Network you want to create (e.g. `my-net`).
func (o NetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The endpoint key to keep secret.
func (o NetworkOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// The prefix that will be prepended to all topics for this Network.
func (o NetworkOutput) TopicPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) pulumi.StringPtrOutput { return v.TopicPrefix }).(pulumi.StringPtrOutput)
}

// The network type to create (e.g. `sigfox`).
func (o NetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Network) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type NetworkArrayOutput struct{ *pulumi.OutputState }

func (NetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Network)(nil)).Elem()
}

func (o NetworkArrayOutput) ToNetworkArrayOutput() NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) Index(i pulumi.IntInput) NetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Network {
		return vs[0].([]*Network)[vs[1].(int)]
	}).(NetworkOutput)
}

type NetworkMapOutput struct{ *pulumi.OutputState }

func (NetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Network)(nil)).Elem()
}

func (o NetworkMapOutput) ToNetworkMapOutput() NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) MapIndex(k pulumi.StringInput) NetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Network {
		return vs[0].(map[string]*Network)[vs[1].(string)]
	}).(NetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInput)(nil)).Elem(), &Network{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkArrayInput)(nil)).Elem(), NetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkMapInput)(nil)).Elem(), NetworkMap{})
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkArrayOutput{})
	pulumi.RegisterOutputType(NetworkMapOutput{})
}
