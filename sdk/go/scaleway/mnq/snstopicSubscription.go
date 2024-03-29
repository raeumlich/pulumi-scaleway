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

// Manage Scaleway Messaging and queuing SNS Topic Subscriptions.
// For further information please check
// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
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
//			// For default project in default region
//			mainSNS, err := mnq.NewSNS(ctx, "mainSNS", nil)
//			if err != nil {
//				return err
//			}
//			mainSNSCredentials, err := mnq.NewSNSCredentials(ctx, "mainSNSCredentials", &mnq.SNSCredentialsArgs{
//				ProjectId: mainSNS.ProjectId,
//				Permissions: &mnq.SNSCredentialsPermissionsArgs{
//					CanManage:  pulumi.Bool(true),
//					CanPublish: pulumi.Bool(true),
//					CanReceive: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			topic, err := mnq.NewSNSTopic(ctx, "topic", &mnq.SNSTopicArgs{
//				ProjectId: mainSNS.ProjectId,
//				AccessKey: mainSNSCredentials.AccessKey,
//				SecretKey: mainSNSCredentials.SecretKey,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mnq.NewSNSTopicSubscription(ctx, "mainSNSTopicSubscription", &mnq.SNSTopicSubscriptionArgs{
//				ProjectId: mainSNS.ProjectId,
//				AccessKey: mainSNSCredentials.AccessKey,
//				SecretKey: mainSNSCredentials.SecretKey,
//				TopicId:   topic.ID(),
//				Protocol:  pulumi.String("http"),
//				Endpoint:  pulumi.String("http://example.com"),
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
// SNS topic subscriptions can be imported using the `{region}/{project-id}/{topic-name}/{subscription-id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:mnq/sNSTopicSubscription:SNSTopicSubscription main fr-par/11111111111111111111111111111111/my-topic/11111111111111111111111111111111
// ```
type SNSTopicSubscription struct {
	pulumi.CustomResourceState

	// The access key of the SNS credentials.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// The ARN of the topic subscription
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Endpoint of the subscription
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Protocol of the SNS Topic Subscription.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Activate JSON Redrive Policy.
	RedrivePolicy pulumi.BoolOutput `pulumi:"redrivePolicy"`
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrOutput `pulumi:"snsEndpoint"`
	// The ARN of the topic. Either `topicId` or `topicArn` is required.
	TopicArn pulumi.StringPtrOutput `pulumi:"topicArn"`
	// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
	TopicId pulumi.StringPtrOutput `pulumi:"topicId"`
}

// NewSNSTopicSubscription registers a new resource with the given unique name, arguments, and options.
func NewSNSTopicSubscription(ctx *pulumi.Context,
	name string, args *SNSTopicSubscriptionArgs, opts ...pulumi.ResourceOption) (*SNSTopicSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessKey == nil {
		return nil, errors.New("invalid value for required argument 'AccessKey'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
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
	var resource SNSTopicSubscription
	err := ctx.RegisterResource("scaleway:mnq/sNSTopicSubscription:SNSTopicSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSNSTopicSubscription gets an existing SNSTopicSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSNSTopicSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SNSTopicSubscriptionState, opts ...pulumi.ResourceOption) (*SNSTopicSubscription, error) {
	var resource SNSTopicSubscription
	err := ctx.ReadResource("scaleway:mnq/sNSTopicSubscription:SNSTopicSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SNSTopicSubscription resources.
type snstopicSubscriptionState struct {
	// The access key of the SNS credentials.
	AccessKey *string `pulumi:"accessKey"`
	// The ARN of the topic subscription
	Arn *string `pulumi:"arn"`
	// Endpoint of the subscription
	Endpoint *string `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// Protocol of the SNS Topic Subscription.
	Protocol *string `pulumi:"protocol"`
	// Activate JSON Redrive Policy.
	RedrivePolicy *bool `pulumi:"redrivePolicy"`
	// `region`). The region
	// in which sns is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey *string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
	// The ARN of the topic. Either `topicId` or `topicArn` is required.
	TopicArn *string `pulumi:"topicArn"`
	// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
	TopicId *string `pulumi:"topicId"`
}

type SNSTopicSubscriptionState struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringPtrInput
	// The ARN of the topic subscription
	Arn pulumi.StringPtrInput
	// Endpoint of the subscription
	Endpoint pulumi.StringPtrInput
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringPtrInput
	// Protocol of the SNS Topic Subscription.
	Protocol pulumi.StringPtrInput
	// Activate JSON Redrive Policy.
	RedrivePolicy pulumi.BoolPtrInput
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringPtrInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
	// The ARN of the topic. Either `topicId` or `topicArn` is required.
	TopicArn pulumi.StringPtrInput
	// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
	TopicId pulumi.StringPtrInput
}

func (SNSTopicSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*snstopicSubscriptionState)(nil)).Elem()
}

type snstopicSubscriptionArgs struct {
	// The access key of the SNS credentials.
	AccessKey string `pulumi:"accessKey"`
	// Endpoint of the subscription
	Endpoint *string `pulumi:"endpoint"`
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId *string `pulumi:"projectId"`
	// Protocol of the SNS Topic Subscription.
	Protocol string `pulumi:"protocol"`
	// Activate JSON Redrive Policy.
	RedrivePolicy *bool `pulumi:"redrivePolicy"`
	// `region`). The region
	// in which sns is enabled.
	Region *string `pulumi:"region"`
	// The secret key of the SNS credentials.
	SecretKey string `pulumi:"secretKey"`
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint *string `pulumi:"snsEndpoint"`
	// The ARN of the topic. Either `topicId` or `topicArn` is required.
	TopicArn *string `pulumi:"topicArn"`
	// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
	TopicId *string `pulumi:"topicId"`
}

// The set of arguments for constructing a SNSTopicSubscription resource.
type SNSTopicSubscriptionArgs struct {
	// The access key of the SNS credentials.
	AccessKey pulumi.StringInput
	// Endpoint of the subscription
	Endpoint pulumi.StringPtrInput
	// `projectId`) The ID of the project the sns is enabled for.
	ProjectId pulumi.StringPtrInput
	// Protocol of the SNS Topic Subscription.
	Protocol pulumi.StringInput
	// Activate JSON Redrive Policy.
	RedrivePolicy pulumi.BoolPtrInput
	// `region`). The region
	// in which sns is enabled.
	Region pulumi.StringPtrInput
	// The secret key of the SNS credentials.
	SecretKey pulumi.StringInput
	// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
	SnsEndpoint pulumi.StringPtrInput
	// The ARN of the topic. Either `topicId` or `topicArn` is required.
	TopicArn pulumi.StringPtrInput
	// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
	TopicId pulumi.StringPtrInput
}

func (SNSTopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snstopicSubscriptionArgs)(nil)).Elem()
}

type SNSTopicSubscriptionInput interface {
	pulumi.Input

	ToSNSTopicSubscriptionOutput() SNSTopicSubscriptionOutput
	ToSNSTopicSubscriptionOutputWithContext(ctx context.Context) SNSTopicSubscriptionOutput
}

func (*SNSTopicSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**SNSTopicSubscription)(nil)).Elem()
}

func (i *SNSTopicSubscription) ToSNSTopicSubscriptionOutput() SNSTopicSubscriptionOutput {
	return i.ToSNSTopicSubscriptionOutputWithContext(context.Background())
}

func (i *SNSTopicSubscription) ToSNSTopicSubscriptionOutputWithContext(ctx context.Context) SNSTopicSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicSubscriptionOutput)
}

// SNSTopicSubscriptionArrayInput is an input type that accepts SNSTopicSubscriptionArray and SNSTopicSubscriptionArrayOutput values.
// You can construct a concrete instance of `SNSTopicSubscriptionArrayInput` via:
//
//	SNSTopicSubscriptionArray{ SNSTopicSubscriptionArgs{...} }
type SNSTopicSubscriptionArrayInput interface {
	pulumi.Input

	ToSNSTopicSubscriptionArrayOutput() SNSTopicSubscriptionArrayOutput
	ToSNSTopicSubscriptionArrayOutputWithContext(context.Context) SNSTopicSubscriptionArrayOutput
}

type SNSTopicSubscriptionArray []SNSTopicSubscriptionInput

func (SNSTopicSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SNSTopicSubscription)(nil)).Elem()
}

func (i SNSTopicSubscriptionArray) ToSNSTopicSubscriptionArrayOutput() SNSTopicSubscriptionArrayOutput {
	return i.ToSNSTopicSubscriptionArrayOutputWithContext(context.Background())
}

func (i SNSTopicSubscriptionArray) ToSNSTopicSubscriptionArrayOutputWithContext(ctx context.Context) SNSTopicSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicSubscriptionArrayOutput)
}

// SNSTopicSubscriptionMapInput is an input type that accepts SNSTopicSubscriptionMap and SNSTopicSubscriptionMapOutput values.
// You can construct a concrete instance of `SNSTopicSubscriptionMapInput` via:
//
//	SNSTopicSubscriptionMap{ "key": SNSTopicSubscriptionArgs{...} }
type SNSTopicSubscriptionMapInput interface {
	pulumi.Input

	ToSNSTopicSubscriptionMapOutput() SNSTopicSubscriptionMapOutput
	ToSNSTopicSubscriptionMapOutputWithContext(context.Context) SNSTopicSubscriptionMapOutput
}

type SNSTopicSubscriptionMap map[string]SNSTopicSubscriptionInput

func (SNSTopicSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SNSTopicSubscription)(nil)).Elem()
}

func (i SNSTopicSubscriptionMap) ToSNSTopicSubscriptionMapOutput() SNSTopicSubscriptionMapOutput {
	return i.ToSNSTopicSubscriptionMapOutputWithContext(context.Background())
}

func (i SNSTopicSubscriptionMap) ToSNSTopicSubscriptionMapOutputWithContext(ctx context.Context) SNSTopicSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SNSTopicSubscriptionMapOutput)
}

type SNSTopicSubscriptionOutput struct{ *pulumi.OutputState }

func (SNSTopicSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SNSTopicSubscription)(nil)).Elem()
}

func (o SNSTopicSubscriptionOutput) ToSNSTopicSubscriptionOutput() SNSTopicSubscriptionOutput {
	return o
}

func (o SNSTopicSubscriptionOutput) ToSNSTopicSubscriptionOutputWithContext(ctx context.Context) SNSTopicSubscriptionOutput {
	return o
}

// The access key of the SNS credentials.
func (o SNSTopicSubscriptionOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// The ARN of the topic subscription
func (o SNSTopicSubscriptionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Endpoint of the subscription
func (o SNSTopicSubscriptionOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// `projectId`) The ID of the project the sns is enabled for.
func (o SNSTopicSubscriptionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Protocol of the SNS Topic Subscription.
func (o SNSTopicSubscriptionOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Activate JSON Redrive Policy.
func (o SNSTopicSubscriptionOutput) RedrivePolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.BoolOutput { return v.RedrivePolicy }).(pulumi.BoolOutput)
}

// `region`). The region
// in which sns is enabled.
func (o SNSTopicSubscriptionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret key of the SNS credentials.
func (o SNSTopicSubscriptionOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
func (o SNSTopicSubscriptionOutput) SnsEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringPtrOutput { return v.SnsEndpoint }).(pulumi.StringPtrOutput)
}

// The ARN of the topic. Either `topicId` or `topicArn` is required.
func (o SNSTopicSubscriptionOutput) TopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringPtrOutput { return v.TopicArn }).(pulumi.StringPtrOutput)
}

// The ID of the topic. Either `topicId` or `topicArn` is required. Conflicts with `topicArn`.
func (o SNSTopicSubscriptionOutput) TopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SNSTopicSubscription) pulumi.StringPtrOutput { return v.TopicId }).(pulumi.StringPtrOutput)
}

type SNSTopicSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SNSTopicSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SNSTopicSubscription)(nil)).Elem()
}

func (o SNSTopicSubscriptionArrayOutput) ToSNSTopicSubscriptionArrayOutput() SNSTopicSubscriptionArrayOutput {
	return o
}

func (o SNSTopicSubscriptionArrayOutput) ToSNSTopicSubscriptionArrayOutputWithContext(ctx context.Context) SNSTopicSubscriptionArrayOutput {
	return o
}

func (o SNSTopicSubscriptionArrayOutput) Index(i pulumi.IntInput) SNSTopicSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SNSTopicSubscription {
		return vs[0].([]*SNSTopicSubscription)[vs[1].(int)]
	}).(SNSTopicSubscriptionOutput)
}

type SNSTopicSubscriptionMapOutput struct{ *pulumi.OutputState }

func (SNSTopicSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SNSTopicSubscription)(nil)).Elem()
}

func (o SNSTopicSubscriptionMapOutput) ToSNSTopicSubscriptionMapOutput() SNSTopicSubscriptionMapOutput {
	return o
}

func (o SNSTopicSubscriptionMapOutput) ToSNSTopicSubscriptionMapOutputWithContext(ctx context.Context) SNSTopicSubscriptionMapOutput {
	return o
}

func (o SNSTopicSubscriptionMapOutput) MapIndex(k pulumi.StringInput) SNSTopicSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SNSTopicSubscription {
		return vs[0].(map[string]*SNSTopicSubscription)[vs[1].(string)]
	}).(SNSTopicSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicSubscriptionInput)(nil)).Elem(), &SNSTopicSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicSubscriptionArrayInput)(nil)).Elem(), SNSTopicSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SNSTopicSubscriptionMapInput)(nil)).Elem(), SNSTopicSubscriptionMap{})
	pulumi.RegisterOutputType(SNSTopicSubscriptionOutput{})
	pulumi.RegisterOutputType(SNSTopicSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SNSTopicSubscriptionMapOutput{})
}
