// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway object storage buckets.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/objectstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewBucket(ctx, "someBucket", &objectstorage.BucketArgs{
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("value"),
//				},
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
// ### Creating the bucket in a specific project
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/objectstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewBucket(ctx, "someBucket", &objectstorage.BucketArgs{
//				ProjectId: pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// ### Using object lifecycle
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/objectstorage"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewBucket(ctx, "main", &objectstorage.BucketArgs{
//				LifecycleRules: objectstorage.BucketLifecycleRuleArray{
//					&objectstorage.BucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Expiration: &objectstorage.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(365),
//						},
//						Id:     pulumi.String("id1"),
//						Prefix: pulumi.String("path1/"),
//						Transitions: objectstorage.BucketLifecycleRuleTransitionArray{
//							&objectstorage.BucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(120),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&objectstorage.BucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Expiration: &objectstorage.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(50),
//						},
//						Id:     pulumi.String("id2"),
//						Prefix: pulumi.String("path2/"),
//					},
//					&objectstorage.BucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(false),
//						Expiration: &objectstorage.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(1),
//						},
//						Id:     pulumi.String("id3"),
//						Prefix: pulumi.String("path3/"),
//						Tags: pulumi.StringMap{
//							"tagKey":    pulumi.String("tagValue"),
//							"terraform": pulumi.String("hashicorp"),
//						},
//					},
//					&objectstorage.BucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Id:      pulumi.String("id4"),
//						Tags: pulumi.StringMap{
//							"tag1": pulumi.String("value1"),
//						},
//						Transitions: objectstorage.BucketLifecycleRuleTransitionArray{
//							&objectstorage.BucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(1),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&objectstorage.BucketLifecycleRuleArgs{
//						AbortIncompleteMultipartUploadDays: pulumi.Int(30),
//						Enabled:                            pulumi.Bool(true),
//					},
//				},
//				Region: pulumi.String("fr-par"),
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
// Buckets can be imported using the `{region}/{bucketName}` identifier, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:objectstorage/bucket:Bucket some_bucket fr-par/some-bucket
// ```
//
// If you are importing a bucket from a specific project (that is not your default project), you can use the following syntax:
//
// bash
//
// ```sh
// $ pulumi import scaleway:objectstorage/bucket:Bucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// API URL of the bucket
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrOutput `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningOutput `pulumi:"versioning"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bucket
	err := ctx.RegisterResource("scaleway:objectstorage/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("scaleway:objectstorage/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl *string `pulumi:"acl"`
	// API URL of the bucket
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint *string `pulumi:"endpoint"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region *string `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags map[string]string `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
}

type BucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrInput
	// API URL of the bucket
	ApiEndpoint pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// The endpoint URL of the bucket
	Endpoint pulumi.StringPtrInput
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringPtrInput
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringPtrInput
	// A list of tags (key / value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl *string `pulumi:"acl"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region *string `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags map[string]string `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *BucketVersioning `pulumi:"versioning"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules BucketCorsRuleArrayInput
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules BucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringPtrInput
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringPtrInput
	// A list of tags (key / value) for the bucket.
	//
	// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
	// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
	Tags pulumi.StringMapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning BucketVersioningPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// (Deprecated) The canned ACL you want to apply to the bucket.
//
// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
func (o BucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// API URL of the bucket
func (o BucketOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
func (o BucketOutput) CorsRules() BucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketCorsRuleArrayOutput { return v.CorsRules }).(BucketCorsRuleArrayOutput)
}

// The endpoint URL of the bucket
func (o BucketOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
func (o BucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
func (o BucketOutput) LifecycleRules() BucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(BucketLifecycleRuleArrayOutput)
}

// The name of the bucket.
func (o BucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable object lock
func (o BucketOutput) ObjectLockEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.ObjectLockEnabled }).(pulumi.BoolPtrOutput)
}

// `projectId`) The ID of the project the bucket is associated with.
//
// The `acl` attribute is deprecated. See objectstorage.BucketACL resource documentation.
// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
func (o BucketOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
func (o BucketOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of tags (key / value) for the bucket.
//
// * > **Important:** The Scaleway console does not support `key/value` tags yet, so only the tags' values will be displayed.
// Keep in mind that if you make any change to your bucket's tags using the console, it will overwrite them with the format `value/value`.
func (o BucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
func (o BucketOutput) Versioning() BucketVersioningOutput {
	return o.ApplyT(func(v *Bucket) BucketVersioningOutput { return v.Versioning }).(BucketVersioningOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
