// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mnq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Messaging and queuing Nats Credentials.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)
//
// ## Example Usage
//
// ### Basic
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/mnq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainNATSAccount, err := mnq.NewNATSAccount(ctx, "mainNATSAccount", nil)
//			if err != nil {
//				return err
//			}
//			_, err = mnq.NewNATSCredentials(ctx, "mainNATSCredentials", &mnq.NATSCredentialsArgs{
//				AccountId: mainNATSAccount.ID(),
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
// Namespaces can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:mnq/nATSCredentials:NATSCredentials main fr-par/11111111111111111111111111111111
// ```
type NATSCredentials struct {
	pulumi.CustomResourceState

	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The content of the credentials file.
	File pulumi.StringOutput `pulumi:"file"`
	// The unique name of the nats credentials.
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewNATSCredentials registers a new resource with the given unique name, arguments, and options.
func NewNATSCredentials(ctx *pulumi.Context,
	name string, args *NATSCredentialsArgs, opts ...pulumi.ResourceOption) (*NATSCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NATSCredentials
	err := ctx.RegisterResource("scaleway:mnq/nATSCredentials:NATSCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNATSCredentials gets an existing NATSCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNATSCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NATSCredentialsState, opts ...pulumi.ResourceOption) (*NATSCredentials, error) {
	var resource NATSCredentials
	err := ctx.ReadResource("scaleway:mnq/nATSCredentials:NATSCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NATSCredentials resources.
type natscredentialsState struct {
	// The ID of the nats account the credentials are generated from
	AccountId *string `pulumi:"accountId"`
	// The content of the credentials file.
	File *string `pulumi:"file"`
	// The unique name of the nats credentials.
	Name *string `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region *string `pulumi:"region"`
}

type NATSCredentialsState struct {
	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringPtrInput
	// The content of the credentials file.
	File pulumi.StringPtrInput
	// The unique name of the nats credentials.
	Name pulumi.StringPtrInput
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringPtrInput
}

func (NATSCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*natscredentialsState)(nil)).Elem()
}

type natscredentialsArgs struct {
	// The ID of the nats account the credentials are generated from
	AccountId string `pulumi:"accountId"`
	// The unique name of the nats credentials.
	Name *string `pulumi:"name"`
	// `region`). The region
	// in which the account exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a NATSCredentials resource.
type NATSCredentialsArgs struct {
	// The ID of the nats account the credentials are generated from
	AccountId pulumi.StringInput
	// The unique name of the nats credentials.
	Name pulumi.StringPtrInput
	// `region`). The region
	// in which the account exists.
	Region pulumi.StringPtrInput
}

func (NATSCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natscredentialsArgs)(nil)).Elem()
}

type NATSCredentialsInput interface {
	pulumi.Input

	ToNATSCredentialsOutput() NATSCredentialsOutput
	ToNATSCredentialsOutputWithContext(ctx context.Context) NATSCredentialsOutput
}

func (*NATSCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((**NATSCredentials)(nil)).Elem()
}

func (i *NATSCredentials) ToNATSCredentialsOutput() NATSCredentialsOutput {
	return i.ToNATSCredentialsOutputWithContext(context.Background())
}

func (i *NATSCredentials) ToNATSCredentialsOutputWithContext(ctx context.Context) NATSCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NATSCredentialsOutput)
}

// NATSCredentialsArrayInput is an input type that accepts NATSCredentialsArray and NATSCredentialsArrayOutput values.
// You can construct a concrete instance of `NATSCredentialsArrayInput` via:
//
//	NATSCredentialsArray{ NATSCredentialsArgs{...} }
type NATSCredentialsArrayInput interface {
	pulumi.Input

	ToNATSCredentialsArrayOutput() NATSCredentialsArrayOutput
	ToNATSCredentialsArrayOutputWithContext(context.Context) NATSCredentialsArrayOutput
}

type NATSCredentialsArray []NATSCredentialsInput

func (NATSCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NATSCredentials)(nil)).Elem()
}

func (i NATSCredentialsArray) ToNATSCredentialsArrayOutput() NATSCredentialsArrayOutput {
	return i.ToNATSCredentialsArrayOutputWithContext(context.Background())
}

func (i NATSCredentialsArray) ToNATSCredentialsArrayOutputWithContext(ctx context.Context) NATSCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NATSCredentialsArrayOutput)
}

// NATSCredentialsMapInput is an input type that accepts NATSCredentialsMap and NATSCredentialsMapOutput values.
// You can construct a concrete instance of `NATSCredentialsMapInput` via:
//
//	NATSCredentialsMap{ "key": NATSCredentialsArgs{...} }
type NATSCredentialsMapInput interface {
	pulumi.Input

	ToNATSCredentialsMapOutput() NATSCredentialsMapOutput
	ToNATSCredentialsMapOutputWithContext(context.Context) NATSCredentialsMapOutput
}

type NATSCredentialsMap map[string]NATSCredentialsInput

func (NATSCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NATSCredentials)(nil)).Elem()
}

func (i NATSCredentialsMap) ToNATSCredentialsMapOutput() NATSCredentialsMapOutput {
	return i.ToNATSCredentialsMapOutputWithContext(context.Background())
}

func (i NATSCredentialsMap) ToNATSCredentialsMapOutputWithContext(ctx context.Context) NATSCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NATSCredentialsMapOutput)
}

type NATSCredentialsOutput struct{ *pulumi.OutputState }

func (NATSCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NATSCredentials)(nil)).Elem()
}

func (o NATSCredentialsOutput) ToNATSCredentialsOutput() NATSCredentialsOutput {
	return o
}

func (o NATSCredentialsOutput) ToNATSCredentialsOutputWithContext(ctx context.Context) NATSCredentialsOutput {
	return o
}

// The ID of the nats account the credentials are generated from
func (o NATSCredentialsOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *NATSCredentials) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The content of the credentials file.
func (o NATSCredentialsOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v *NATSCredentials) pulumi.StringOutput { return v.File }).(pulumi.StringOutput)
}

// The unique name of the nats credentials.
func (o NATSCredentialsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NATSCredentials) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`). The region
// in which the account exists.
func (o NATSCredentialsOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NATSCredentials) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type NATSCredentialsArrayOutput struct{ *pulumi.OutputState }

func (NATSCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NATSCredentials)(nil)).Elem()
}

func (o NATSCredentialsArrayOutput) ToNATSCredentialsArrayOutput() NATSCredentialsArrayOutput {
	return o
}

func (o NATSCredentialsArrayOutput) ToNATSCredentialsArrayOutputWithContext(ctx context.Context) NATSCredentialsArrayOutput {
	return o
}

func (o NATSCredentialsArrayOutput) Index(i pulumi.IntInput) NATSCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NATSCredentials {
		return vs[0].([]*NATSCredentials)[vs[1].(int)]
	}).(NATSCredentialsOutput)
}

type NATSCredentialsMapOutput struct{ *pulumi.OutputState }

func (NATSCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NATSCredentials)(nil)).Elem()
}

func (o NATSCredentialsMapOutput) ToNATSCredentialsMapOutput() NATSCredentialsMapOutput {
	return o
}

func (o NATSCredentialsMapOutput) ToNATSCredentialsMapOutputWithContext(ctx context.Context) NATSCredentialsMapOutput {
	return o
}

func (o NATSCredentialsMapOutput) MapIndex(k pulumi.StringInput) NATSCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NATSCredentials {
		return vs[0].(map[string]*NATSCredentials)[vs[1].(string)]
	}).(NATSCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NATSCredentialsInput)(nil)).Elem(), &NATSCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*NATSCredentialsArrayInput)(nil)).Elem(), NATSCredentialsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NATSCredentialsMapInput)(nil)).Elem(), NATSCredentialsMap{})
	pulumi.RegisterOutputType(NATSCredentialsOutput{})
	pulumi.RegisterOutputType(NATSCredentialsArrayOutput{})
	pulumi.RegisterOutputType(NATSCredentialsMapOutput{})
}
