// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway IAM API Keys. For more information, please
// check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#api-keys-3665ae)
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewApplication(ctx, "ciCd", nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewAPIKey(ctx, "main", &iam.APIKeyArgs{
//				ApplicationId: pulumi.Any(scaleway_iam_application.Main.Id),
//				Description:   pulumi.String("a description"),
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
// Api keys can be imported using the `{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:iam/aPIKey:APIKey main 11111111111111111111
// ```
type APIKey struct {
	pulumi.CustomResourceState

	// The access key of the iam api key.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// ID of the application attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	ApplicationId pulumi.StringPtrOutput `pulumi:"applicationId"`
	// The date and time of the creation of the iam api key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The IP Address of the device which created the API key.
	CreationIp pulumi.StringOutput `pulumi:"creationIp"`
	// The default project ID to use with object storage.
	DefaultProjectId pulumi.StringOutput `pulumi:"defaultProjectId"`
	// The description of the iam api key.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the iam api key is editable.
	Editable pulumi.BoolOutput `pulumi:"editable"`
	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The secret Key of the iam api key.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// The date and time of the last update of the iam api key.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// ID of the user attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewAPIKey registers a new resource with the given unique name, arguments, and options.
func NewAPIKey(ctx *pulumi.Context,
	name string, args *APIKeyArgs, opts ...pulumi.ResourceOption) (*APIKey, error) {
	if args == nil {
		args = &APIKeyArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource APIKey
	err := ctx.RegisterResource("scaleway:iam/aPIKey:APIKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIKey gets an existing APIKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIKeyState, opts ...pulumi.ResourceOption) (*APIKey, error) {
	var resource APIKey
	err := ctx.ReadResource("scaleway:iam/aPIKey:APIKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIKey resources.
type apikeyState struct {
	// The access key of the iam api key.
	AccessKey *string `pulumi:"accessKey"`
	// ID of the application attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	ApplicationId *string `pulumi:"applicationId"`
	// The date and time of the creation of the iam api key.
	CreatedAt *string `pulumi:"createdAt"`
	// The IP Address of the device which created the API key.
	CreationIp *string `pulumi:"creationIp"`
	// The default project ID to use with object storage.
	DefaultProjectId *string `pulumi:"defaultProjectId"`
	// The description of the iam api key.
	Description *string `pulumi:"description"`
	// Whether the iam api key is editable.
	Editable *bool `pulumi:"editable"`
	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The secret Key of the iam api key.
	SecretKey *string `pulumi:"secretKey"`
	// The date and time of the last update of the iam api key.
	UpdatedAt *string `pulumi:"updatedAt"`
	// ID of the user attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	UserId *string `pulumi:"userId"`
}

type APIKeyState struct {
	// The access key of the iam api key.
	AccessKey pulumi.StringPtrInput
	// ID of the application attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	ApplicationId pulumi.StringPtrInput
	// The date and time of the creation of the iam api key.
	CreatedAt pulumi.StringPtrInput
	// The IP Address of the device which created the API key.
	CreationIp pulumi.StringPtrInput
	// The default project ID to use with object storage.
	DefaultProjectId pulumi.StringPtrInput
	// The description of the iam api key.
	Description pulumi.StringPtrInput
	// Whether the iam api key is editable.
	Editable pulumi.BoolPtrInput
	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	ExpiresAt pulumi.StringPtrInput
	// The secret Key of the iam api key.
	SecretKey pulumi.StringPtrInput
	// The date and time of the last update of the iam api key.
	UpdatedAt pulumi.StringPtrInput
	// ID of the user attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	UserId pulumi.StringPtrInput
}

func (APIKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyState)(nil)).Elem()
}

type apikeyArgs struct {
	// ID of the application attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	ApplicationId *string `pulumi:"applicationId"`
	// The default project ID to use with object storage.
	DefaultProjectId *string `pulumi:"defaultProjectId"`
	// The description of the iam api key.
	Description *string `pulumi:"description"`
	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	ExpiresAt *string `pulumi:"expiresAt"`
	// ID of the user attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a APIKey resource.
type APIKeyArgs struct {
	// ID of the application attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	ApplicationId pulumi.StringPtrInput
	// The default project ID to use with object storage.
	DefaultProjectId pulumi.StringPtrInput
	// The description of the iam api key.
	Description pulumi.StringPtrInput
	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	ExpiresAt pulumi.StringPtrInput
	// ID of the user attached to the api key.
	// Only one of the `applicationId` and `userId` should be specified.
	UserId pulumi.StringPtrInput
}

func (APIKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyArgs)(nil)).Elem()
}

type APIKeyInput interface {
	pulumi.Input

	ToAPIKeyOutput() APIKeyOutput
	ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput
}

func (*APIKey) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKey)(nil)).Elem()
}

func (i *APIKey) ToAPIKeyOutput() APIKeyOutput {
	return i.ToAPIKeyOutputWithContext(context.Background())
}

func (i *APIKey) ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyOutput)
}

// APIKeyArrayInput is an input type that accepts APIKeyArray and APIKeyArrayOutput values.
// You can construct a concrete instance of `APIKeyArrayInput` via:
//
//	APIKeyArray{ APIKeyArgs{...} }
type APIKeyArrayInput interface {
	pulumi.Input

	ToAPIKeyArrayOutput() APIKeyArrayOutput
	ToAPIKeyArrayOutputWithContext(context.Context) APIKeyArrayOutput
}

type APIKeyArray []APIKeyInput

func (APIKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIKey)(nil)).Elem()
}

func (i APIKeyArray) ToAPIKeyArrayOutput() APIKeyArrayOutput {
	return i.ToAPIKeyArrayOutputWithContext(context.Background())
}

func (i APIKeyArray) ToAPIKeyArrayOutputWithContext(ctx context.Context) APIKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyArrayOutput)
}

// APIKeyMapInput is an input type that accepts APIKeyMap and APIKeyMapOutput values.
// You can construct a concrete instance of `APIKeyMapInput` via:
//
//	APIKeyMap{ "key": APIKeyArgs{...} }
type APIKeyMapInput interface {
	pulumi.Input

	ToAPIKeyMapOutput() APIKeyMapOutput
	ToAPIKeyMapOutputWithContext(context.Context) APIKeyMapOutput
}

type APIKeyMap map[string]APIKeyInput

func (APIKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIKey)(nil)).Elem()
}

func (i APIKeyMap) ToAPIKeyMapOutput() APIKeyMapOutput {
	return i.ToAPIKeyMapOutputWithContext(context.Background())
}

func (i APIKeyMap) ToAPIKeyMapOutputWithContext(ctx context.Context) APIKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyMapOutput)
}

type APIKeyOutput struct{ *pulumi.OutputState }

func (APIKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKey)(nil)).Elem()
}

func (o APIKeyOutput) ToAPIKeyOutput() APIKeyOutput {
	return o
}

func (o APIKeyOutput) ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput {
	return o
}

// The access key of the iam api key.
func (o APIKeyOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// ID of the application attached to the api key.
// Only one of the `applicationId` and `userId` should be specified.
func (o APIKeyOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringPtrOutput { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

// The date and time of the creation of the iam api key.
func (o APIKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The IP Address of the device which created the API key.
func (o APIKeyOutput) CreationIp() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.CreationIp }).(pulumi.StringOutput)
}

// The default project ID to use with object storage.
func (o APIKeyOutput) DefaultProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.DefaultProjectId }).(pulumi.StringOutput)
}

// The description of the iam api key.
func (o APIKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the iam api key is editable.
func (o APIKeyOutput) Editable() pulumi.BoolOutput {
	return o.ApplyT(func(v *APIKey) pulumi.BoolOutput { return v.Editable }).(pulumi.BoolOutput)
}

// The date and time of the expiration of the iam api key. Please note that in case of change,
// the resource will be recreated.
func (o APIKeyOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// The secret Key of the iam api key.
func (o APIKeyOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

// The date and time of the last update of the iam api key.
func (o APIKeyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// ID of the user attached to the api key.
// Only one of the `applicationId` and `userId` should be specified.
func (o APIKeyOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APIKey) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type APIKeyArrayOutput struct{ *pulumi.OutputState }

func (APIKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIKey)(nil)).Elem()
}

func (o APIKeyArrayOutput) ToAPIKeyArrayOutput() APIKeyArrayOutput {
	return o
}

func (o APIKeyArrayOutput) ToAPIKeyArrayOutputWithContext(ctx context.Context) APIKeyArrayOutput {
	return o
}

func (o APIKeyArrayOutput) Index(i pulumi.IntInput) APIKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APIKey {
		return vs[0].([]*APIKey)[vs[1].(int)]
	}).(APIKeyOutput)
}

type APIKeyMapOutput struct{ *pulumi.OutputState }

func (APIKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIKey)(nil)).Elem()
}

func (o APIKeyMapOutput) ToAPIKeyMapOutput() APIKeyMapOutput {
	return o
}

func (o APIKeyMapOutput) ToAPIKeyMapOutputWithContext(ctx context.Context) APIKeyMapOutput {
	return o
}

func (o APIKeyMapOutput) MapIndex(k pulumi.StringInput) APIKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APIKey {
		return vs[0].(map[string]*APIKey)[vs[1].(string)]
	}).(APIKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyInput)(nil)).Elem(), &APIKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyArrayInput)(nil)).Elem(), APIKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIKeyMapInput)(nil)).Elem(), APIKeyMap{})
	pulumi.RegisterOutputType(APIKeyOutput{})
	pulumi.RegisterOutputType(APIKeyArrayOutput{})
	pulumi.RegisterOutputType(APIKeyMapOutput{})
}
