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

// Manage Scaleway Messaging and queuing SNS Topics.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/mnq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainSNS, err := mnq.NewSNS(ctx, "mainSNS", nil)
//			if err != nil {
//				return err
//			}
//			mainSNSCredentials, err := mnq.NewSNSCredentials(ctx, "mainSNSCredentials", &mnq.SNSCredentialsArgs{
//				ProjectId: mainSNS.ProjectId,
//				Permissions: &mnq.SNSCredentialsPermissionsArgs{
//					CanManage: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mnq.NewSNSTopic(ctx, "topic", &mnq.SNSTopicArgs{
//				ProjectId: mainSNS.ProjectId,
//				AccessKey: mainSNSCredentials.AccessKey,
//				SecretKey: mainSNSCredentials.SecretKey,
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
// SNS topic can be imported using the `{region}/{project-id}/{topic-name}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:mnq/sNSTopic:SNSTopic main fr-par/11111111111111111111111111111111/my-topic
//
// ```
type SNSTopic struct {
	pulumi.CustomResourceState

	// The access key of the SNS credentials.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// The ARN of the topic
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolOutput `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolOutput `pulumi:"fifoTopic"`
	// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner pulumi.StringOutput `pulumi:"owner"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrOutput `pulumi:"snsEndpoint"`
}

// NewSNSTopic registers a new resource with the given unique name, arguments, and options.
func NewSNSTopic(ctx *pulumi.Context,
	name string, args *SNSTopicArgs, opts ...pulumi.ResourceOption) (*SNSTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.AccessKey != nil {
		args.AccessKey = pulumi.ToSecret(args.AccessKey).(pulumi.StringInput)
	}
	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SNSTopic
	err := ctx.RegisterResource("scaleway:mnq/sNSTopic:SNSTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSNSTopic gets an existing SNSTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSNSTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SNSTopicState, opts ...pulumi.ResourceOption) (*SNSTopic, error) {
	var resource SNSTopic
	err := ctx.ReadResource("scaleway:mnq/sNSTopic:SNSTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SNSTopic resources.
type snstopicState struct {
	// The access key of the SNS credentials.
	AccessKey *string `pulumi:"accessKey"`
	// The ARN of the topic
	Arn *string `pulumi:"arn"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
	FifoTopic *bool `pulumi:"fifoTopic"`
	// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner *string `pulumi:"owner"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sns is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey *string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
}

type SNSTopicState struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringPtrInput
	// The ARN of the topic
	Arn pulumi.StringPtrInput
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolPtrInput
	// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Owner of the SNS topic, should have format 'project-${project_id}'
	Owner pulumi.StringPtrInput
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringPtrInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
}

func (SNSTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*snstopicState)(nil)).Elem()
}

type snstopicArgs struct {
	// The access key of the SNS credentials.
	AccessKey string `pulumi:"accessKey"`
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
	FifoTopic *bool `pulumi:"fifoTopic"`
	// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region
	// in which sns is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
}

// The set of arguments for constructing a SNSTopic resource.
type SNSTopicArgs struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringInput
	// Specifies whether to enable content-based deduplication.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
	FifoTopic pulumi.BoolPtrInput
	// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringPtrInput
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
}

func (SNSTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snstopicArgs)(nil)).Elem()
}

type SNSTopicInput interface {
	pulumi.Input

	ToSNSTopicOutput() SNSTopicOutput
	ToSNSTopicOutputWithContext(ctx context.Context) SNSTopicOutput
}

func (*SNSTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**SNSTopic)(nil)).Elem()
}

func (i *SNSTopic) ToSNSTopicOutput() SNSTopicOutput {
	return i.ToSNSTopicOutputWithContext(context.Background())
}

func (i *SNSTopic) ToSNSTopicOutputWithContext(ctx context.Context) SNSTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicOutput)
}

// SNSTopicArrayInput is an input type that accepts SNSTopicArray and SNSTopicArrayOutput values.
// You can construct a concrete instance of `SNSTopicArrayInput` via:
//
//	SNSTopicArray{ SNSTopicArgs{...} }
type SNSTopicArrayInput interface {
	pulumi.Input

	ToSNSTopicArrayOutput() SNSTopicArrayOutput
	ToSNSTopicArrayOutputWithContext(context.Context) SNSTopicArrayOutput
}

type SNSTopicArray []SNSTopicInput

func (SNSTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SNSTopic)(nil)).Elem()
}

func (i SNSTopicArray) ToSNSTopicArrayOutput() SNSTopicArrayOutput {
	return i.ToSNSTopicArrayOutputWithContext(context.Background())
}

func (i SNSTopicArray) ToSNSTopicArrayOutputWithContext(ctx context.Context) SNSTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicArrayOutput)
}

// SNSTopicMapInput is an input type that accepts SNSTopicMap and SNSTopicMapOutput values.
// You can construct a concrete instance of `SNSTopicMapInput` via:
//
//	SNSTopicMap{ "key": SNSTopicArgs{...} }
type SNSTopicMapInput interface {
	pulumi.Input

	ToSNSTopicMapOutput() SNSTopicMapOutput
	ToSNSTopicMapOutputWithContext(context.Context) SNSTopicMapOutput
}

type SNSTopicMap map[string]SNSTopicInput

func (SNSTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SNSTopic)(nil)).Elem()
}

func (i SNSTopicMap) ToSNSTopicMapOutput() SNSTopicMapOutput {
	return i.ToSNSTopicMapOutputWithContext(context.Background())
}

func (i SNSTopicMap) ToSNSTopicMapOutputWithContext(ctx context.Context) SNSTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicMapOutput)
}

type SNSTopicOutput struct{ *pulumi.OutputState }

func (SNSTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SNSTopic)(nil)).Elem()
}

func (o SNSTopicOutput) ToSNSTopicOutput() SNSTopicOutput {
	return o
}

func (o SNSTopicOutput) ToSNSTopicOutputWithContext(ctx context.Context) SNSTopicOutput {
	return o
}

// The access key of the SNS credentials.
func (o SNSTopicOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// The ARN of the topic
func (o SNSTopicOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether to enable content-based deduplication.
func (o SNSTopicOutput) ContentBasedDeduplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.BoolOutput { return v.ContentBasedDeduplication }).(pulumi.BoolOutput)
}

// Whether the topic is a FIFO. If true, the topic name must end with .fifo.
func (o SNSTopicOutput) FifoTopic() pulumi.BoolOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.BoolOutput { return v.FifoTopic }).(pulumi.BoolOutput)
}

// The unique name of the sns topic. Either `name` or `namePrefix` is required. Conflicts with `namePrefix`.
func (o SNSTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o SNSTopicOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Owner of the SNS topic, should have format 'project-${project_id}'
func (o SNSTopicOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the sns is enabled for.
func (o SNSTopicOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region
// in which sns is enabled.
func (o SNSTopicOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret key of the SNS credentials.
func (o SNSTopicOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
func (o SNSTopicOutput) SnsEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SNSTopic) pulumi.StringPtrOutput { return v.SnsEndpoint }).(pulumi.StringPtrOutput)
}

type SNSTopicArrayOutput struct{ *pulumi.OutputState }

func (SNSTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SNSTopic)(nil)).Elem()
}

func (o SNSTopicArrayOutput) ToSNSTopicArrayOutput() SNSTopicArrayOutput {
	return o
}

func (o SNSTopicArrayOutput) ToSNSTopicArrayOutputWithContext(ctx context.Context) SNSTopicArrayOutput {
	return o
}

func (o SNSTopicArrayOutput) Index(i pulumi.IntInput) SNSTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SNSTopic {
		return vs[0].([]*SNSTopic)[vs[1].(int)]
	}).(SNSTopicOutput)
}

type SNSTopicMapOutput struct{ *pulumi.OutputState }

func (SNSTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SNSTopic)(nil)).Elem()
}

func (o SNSTopicMapOutput) ToSNSTopicMapOutput() SNSTopicMapOutput {
	return o
}

func (o SNSTopicMapOutput) ToSNSTopicMapOutputWithContext(ctx context.Context) SNSTopicMapOutput {
	return o
}

func (o SNSTopicMapOutput) MapIndex(k pulumi.StringInput) SNSTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SNSTopic {
		return vs[0].(map[string]*SNSTopic)[vs[1].(string)]
	}).(SNSTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicInput)(nil)).Elem(), &SNSTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicArrayInput)(nil)).Elem(), SNSTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicMapInput)(nil)).Elem(), SNSTopicMap{})
	pulumi.RegisterOutputType(SNSTopicOutput{})
	pulumi.RegisterOutputType(SNSTopicArrayOutput{})
	pulumi.RegisterOutputType(SNSTopicMapOutput{})
}
