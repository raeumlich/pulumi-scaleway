// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serverless

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Creates and manages Scaleway Container Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
//
// For more information consult
// the [documentation](https://www.scaleway.com/en/docs/serverless/containers/)
// .
//
// For more details about the limitation
// check [containers-limitations](https://www.scaleway.com/en/docs/compute/containers/reference-content/containers-limitations/)
// .
//
// You can check also
// our [containers cron api documentation](https://developers.scaleway.com/en/products/containers/api/#crons-942bf4).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/serverless"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainContainerNamespace, err := serverless.NewContainerNamespace(ctx, "mainContainerNamespace", nil)
//			if err != nil {
//				return err
//			}
//			mainContainer, err := serverless.NewContainer(ctx, "mainContainer", &serverless.ContainerArgs{
//				NamespaceId: mainContainerNamespace.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"address": map[string]interface{}{
//					"city":    "Paris",
//					"country": "FR",
//				},
//				"age":       23,
//				"firstName": "John",
//				"isAlive":   true,
//				"lastName":  "Smith",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = serverless.NewContainerCron(ctx, "mainContainerCron", &serverless.ContainerCronArgs{
//				ContainerId: mainContainer.ID(),
//				Schedule:    pulumi.String("5 4 1 * *"),
//				Args:        pulumi.String(json0),
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
// Container Cron can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:serverless/containerCron:ContainerCron main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type ContainerCron struct {
	pulumi.CustomResourceState

	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	Args pulumi.StringOutput `pulumi:"args"`
	// The container ID to link with your cron.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The name of the container cron. If not provided, the name is generated.
	// during
	Name pulumi.StringOutput `pulumi:"name"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The cron status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewContainerCron registers a new resource with the given unique name, arguments, and options.
func NewContainerCron(ctx *pulumi.Context,
	name string, args *ContainerCronArgs, opts ...pulumi.ResourceOption) (*ContainerCron, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerCron
	err := ctx.RegisterResource("scaleway:serverless/containerCron:ContainerCron", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerCron gets an existing ContainerCron resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerCron(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerCronState, opts ...pulumi.ResourceOption) (*ContainerCron, error) {
	var resource ContainerCron
	err := ctx.ReadResource("scaleway:serverless/containerCron:ContainerCron", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerCron resources.
type containerCronState struct {
	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	Args *string `pulumi:"args"`
	// The container ID to link with your cron.
	ContainerId *string `pulumi:"containerId"`
	// The name of the container cron. If not provided, the name is generated.
	// during
	Name *string `pulumi:"name"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region *string `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule *string `pulumi:"schedule"`
	// The cron status.
	Status *string `pulumi:"status"`
}

type ContainerCronState struct {
	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	Args pulumi.StringPtrInput
	// The container ID to link with your cron.
	ContainerId pulumi.StringPtrInput
	// The name of the container cron. If not provided, the name is generated.
	// during
	Name pulumi.StringPtrInput
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringPtrInput
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringPtrInput
	// The cron status.
	Status pulumi.StringPtrInput
}

func (ContainerCronState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerCronState)(nil)).Elem()
}

type containerCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	Args string `pulumi:"args"`
	// The container ID to link with your cron.
	ContainerId string `pulumi:"containerId"`
	// The name of the container cron. If not provided, the name is generated.
	// during
	Name *string `pulumi:"name"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region *string `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule string `pulumi:"schedule"`
}

// The set of arguments for constructing a ContainerCron resource.
type ContainerCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	Args pulumi.StringInput
	// The container ID to link with your cron.
	ContainerId pulumi.StringInput
	// The name of the container cron. If not provided, the name is generated.
	// during
	Name pulumi.StringPtrInput
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringPtrInput
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringInput
}

func (ContainerCronArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerCronArgs)(nil)).Elem()
}

type ContainerCronInput interface {
	pulumi.Input

	ToContainerCronOutput() ContainerCronOutput
	ToContainerCronOutputWithContext(ctx context.Context) ContainerCronOutput
}

func (*ContainerCron) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerCron)(nil)).Elem()
}

func (i *ContainerCron) ToContainerCronOutput() ContainerCronOutput {
	return i.ToContainerCronOutputWithContext(context.Background())
}

func (i *ContainerCron) ToContainerCronOutputWithContext(ctx context.Context) ContainerCronOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCronOutput)
}

// ContainerCronArrayInput is an input type that accepts ContainerCronArray and ContainerCronArrayOutput values.
// You can construct a concrete instance of `ContainerCronArrayInput` via:
//
//	ContainerCronArray{ ContainerCronArgs{...} }
type ContainerCronArrayInput interface {
	pulumi.Input

	ToContainerCronArrayOutput() ContainerCronArrayOutput
	ToContainerCronArrayOutputWithContext(context.Context) ContainerCronArrayOutput
}

type ContainerCronArray []ContainerCronInput

func (ContainerCronArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerCron)(nil)).Elem()
}

func (i ContainerCronArray) ToContainerCronArrayOutput() ContainerCronArrayOutput {
	return i.ToContainerCronArrayOutputWithContext(context.Background())
}

func (i ContainerCronArray) ToContainerCronArrayOutputWithContext(ctx context.Context) ContainerCronArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCronArrayOutput)
}

// ContainerCronMapInput is an input type that accepts ContainerCronMap and ContainerCronMapOutput values.
// You can construct a concrete instance of `ContainerCronMapInput` via:
//
//	ContainerCronMap{ "key": ContainerCronArgs{...} }
type ContainerCronMapInput interface {
	pulumi.Input

	ToContainerCronMapOutput() ContainerCronMapOutput
	ToContainerCronMapOutputWithContext(context.Context) ContainerCronMapOutput
}

type ContainerCronMap map[string]ContainerCronInput

func (ContainerCronMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerCron)(nil)).Elem()
}

func (i ContainerCronMap) ToContainerCronMapOutput() ContainerCronMapOutput {
	return i.ToContainerCronMapOutputWithContext(context.Background())
}

func (i ContainerCronMap) ToContainerCronMapOutputWithContext(ctx context.Context) ContainerCronMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCronMapOutput)
}

type ContainerCronOutput struct{ *pulumi.OutputState }

func (ContainerCronOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerCron)(nil)).Elem()
}

func (o ContainerCronOutput) ToContainerCronOutput() ContainerCronOutput {
	return o
}

func (o ContainerCronOutput) ToContainerCronOutputWithContext(ctx context.Context) ContainerCronOutput {
	return o
}

// The key-value mapping to define arguments that will be passed to your container’s event object
// during
func (o ContainerCronOutput) Args() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.Args }).(pulumi.StringOutput)
}

// The container ID to link with your cron.
func (o ContainerCronOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// The name of the container cron. If not provided, the name is generated.
// during
func (o ContainerCronOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Defaults to provider `region`) The region
// in where the job was created.
func (o ContainerCronOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
// executed.
func (o ContainerCronOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The cron status.
func (o ContainerCronOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerCron) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ContainerCronArrayOutput struct{ *pulumi.OutputState }

func (ContainerCronArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerCron)(nil)).Elem()
}

func (o ContainerCronArrayOutput) ToContainerCronArrayOutput() ContainerCronArrayOutput {
	return o
}

func (o ContainerCronArrayOutput) ToContainerCronArrayOutputWithContext(ctx context.Context) ContainerCronArrayOutput {
	return o
}

func (o ContainerCronArrayOutput) Index(i pulumi.IntInput) ContainerCronOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerCron {
		return vs[0].([]*ContainerCron)[vs[1].(int)]
	}).(ContainerCronOutput)
}

type ContainerCronMapOutput struct{ *pulumi.OutputState }

func (ContainerCronMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerCron)(nil)).Elem()
}

func (o ContainerCronMapOutput) ToContainerCronMapOutput() ContainerCronMapOutput {
	return o
}

func (o ContainerCronMapOutput) ToContainerCronMapOutputWithContext(ctx context.Context) ContainerCronMapOutput {
	return o
}

func (o ContainerCronMapOutput) MapIndex(k pulumi.StringInput) ContainerCronOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerCron {
		return vs[0].(map[string]*ContainerCron)[vs[1].(string)]
	}).(ContainerCronOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerCronInput)(nil)).Elem(), &ContainerCron{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerCronArrayInput)(nil)).Elem(), ContainerCronArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerCronMapInput)(nil)).Elem(), ContainerCronMap{})
	pulumi.RegisterOutputType(ContainerCronOutput{})
	pulumi.RegisterOutputType(ContainerCronArrayOutput{})
	pulumi.RegisterOutputType(ContainerCronMapOutput{})
}
