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

// Creates and manages a Scaleway Serverless Job Definition. For more information, see [the documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/jobs/v1alpha1).
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
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/serverless"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := serverless.NewJobDefinition(ctx, "main", &serverless.JobDefinitionArgs{
//				CpuLimit:    pulumi.Int(140),
//				MemoryLimit: pulumi.Int(256),
//				ImageUri:    pulumi.String("docker.io/alpine:latest"),
//				Command:     pulumi.String("ls"),
//				Timeout:     pulumi.String("10m"),
//				Env: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
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
// ## Import
//
// Serverless Jobs can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:serverless/jobDefinition:JobDefinition job fr-par/11111111-1111-1111-1111-111111111111
// ```
type JobDefinition struct {
	pulumi.CustomResourceState

	// The command that will be run in the container if specified.
	Command pulumi.StringPtrOutput `pulumi:"command"`
	// The amount of vCPU computing resources to allocate to each container running the job.
	CpuLimit pulumi.IntOutput           `pulumi:"cpuLimit"`
	Cron     JobDefinitionCronPtrOutput `pulumi:"cron"`
	// The description of the job
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The environment variables of the container.
	Env pulumi.StringMapOutput `pulumi:"env"`
	// The uri of the container image that will be used for the job run.
	ImageUri pulumi.StringPtrOutput `pulumi:"imageUri"`
	// The memory computing resources in MB to allocate to each container running the job.
	MemoryLimit pulumi.IntOutput `pulumi:"memoryLimit"`
	// The name of the job.
	Name pulumi.StringOutput `pulumi:"name"`
	// `projectId`) The ID of the project the Job is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region of the Job.
	Region pulumi.StringOutput `pulumi:"region"`
	// The job run timeout, in Go Time format (ex: `2h30m25s`)
	Timeout pulumi.StringOutput `pulumi:"timeout"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CpuLimit == nil {
		return nil, errors.New("invalid value for required argument 'CpuLimit'")
	}
	if args.MemoryLimit == nil {
		return nil, errors.New("invalid value for required argument 'MemoryLimit'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JobDefinition
	err := ctx.RegisterResource("scaleway:serverless/jobDefinition:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("scaleway:serverless/jobDefinition:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
	// The command that will be run in the container if specified.
	Command *string `pulumi:"command"`
	// The amount of vCPU computing resources to allocate to each container running the job.
	CpuLimit *int               `pulumi:"cpuLimit"`
	Cron     *JobDefinitionCron `pulumi:"cron"`
	// The description of the job
	Description *string `pulumi:"description"`
	// The environment variables of the container.
	Env map[string]string `pulumi:"env"`
	// The uri of the container image that will be used for the job run.
	ImageUri *string `pulumi:"imageUri"`
	// The memory computing resources in MB to allocate to each container running the job.
	MemoryLimit *int `pulumi:"memoryLimit"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the Job is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Job.
	Region *string `pulumi:"region"`
	// The job run timeout, in Go Time format (ex: `2h30m25s`)
	Timeout *string `pulumi:"timeout"`
}

type JobDefinitionState struct {
	// The command that will be run in the container if specified.
	Command pulumi.StringPtrInput
	// The amount of vCPU computing resources to allocate to each container running the job.
	CpuLimit pulumi.IntPtrInput
	Cron     JobDefinitionCronPtrInput
	// The description of the job
	Description pulumi.StringPtrInput
	// The environment variables of the container.
	Env pulumi.StringMapInput
	// The uri of the container image that will be used for the job run.
	ImageUri pulumi.StringPtrInput
	// The memory computing resources in MB to allocate to each container running the job.
	MemoryLimit pulumi.IntPtrInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the Job is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Job.
	Region pulumi.StringPtrInput
	// The job run timeout, in Go Time format (ex: `2h30m25s`)
	Timeout pulumi.StringPtrInput
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// The command that will be run in the container if specified.
	Command *string `pulumi:"command"`
	// The amount of vCPU computing resources to allocate to each container running the job.
	CpuLimit int                `pulumi:"cpuLimit"`
	Cron     *JobDefinitionCron `pulumi:"cron"`
	// The description of the job
	Description *string `pulumi:"description"`
	// The environment variables of the container.
	Env map[string]string `pulumi:"env"`
	// The uri of the container image that will be used for the job run.
	ImageUri *string `pulumi:"imageUri"`
	// The memory computing resources in MB to allocate to each container running the job.
	MemoryLimit int `pulumi:"memoryLimit"`
	// The name of the job.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the Job is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Job.
	Region *string `pulumi:"region"`
	// The job run timeout, in Go Time format (ex: `2h30m25s`)
	Timeout *string `pulumi:"timeout"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// The command that will be run in the container if specified.
	Command pulumi.StringPtrInput
	// The amount of vCPU computing resources to allocate to each container running the job.
	CpuLimit pulumi.IntInput
	Cron     JobDefinitionCronPtrInput
	// The description of the job
	Description pulumi.StringPtrInput
	// The environment variables of the container.
	Env pulumi.StringMapInput
	// The uri of the container image that will be used for the job run.
	ImageUri pulumi.StringPtrInput
	// The memory computing resources in MB to allocate to each container running the job.
	MemoryLimit pulumi.IntInput
	// The name of the job.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the Job is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Job.
	Region pulumi.StringPtrInput
	// The job run timeout, in Go Time format (ex: `2h30m25s`)
	Timeout pulumi.StringPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

// JobDefinitionArrayInput is an input type that accepts JobDefinitionArray and JobDefinitionArrayOutput values.
// You can construct a concrete instance of `JobDefinitionArrayInput` via:
//
//	JobDefinitionArray{ JobDefinitionArgs{...} }
type JobDefinitionArrayInput interface {
	pulumi.Input

	ToJobDefinitionArrayOutput() JobDefinitionArrayOutput
	ToJobDefinitionArrayOutputWithContext(context.Context) JobDefinitionArrayOutput
}

type JobDefinitionArray []JobDefinitionInput

func (JobDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobDefinition)(nil)).Elem()
}

func (i JobDefinitionArray) ToJobDefinitionArrayOutput() JobDefinitionArrayOutput {
	return i.ToJobDefinitionArrayOutputWithContext(context.Background())
}

func (i JobDefinitionArray) ToJobDefinitionArrayOutputWithContext(ctx context.Context) JobDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionArrayOutput)
}

// JobDefinitionMapInput is an input type that accepts JobDefinitionMap and JobDefinitionMapOutput values.
// You can construct a concrete instance of `JobDefinitionMapInput` via:
//
//	JobDefinitionMap{ "key": JobDefinitionArgs{...} }
type JobDefinitionMapInput interface {
	pulumi.Input

	ToJobDefinitionMapOutput() JobDefinitionMapOutput
	ToJobDefinitionMapOutputWithContext(context.Context) JobDefinitionMapOutput
}

type JobDefinitionMap map[string]JobDefinitionInput

func (JobDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobDefinition)(nil)).Elem()
}

func (i JobDefinitionMap) ToJobDefinitionMapOutput() JobDefinitionMapOutput {
	return i.ToJobDefinitionMapOutputWithContext(context.Background())
}

func (i JobDefinitionMap) ToJobDefinitionMapOutputWithContext(ctx context.Context) JobDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionMapOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

// The command that will be run in the container if specified.
func (o JobDefinitionOutput) Command() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.Command }).(pulumi.StringPtrOutput)
}

// The amount of vCPU computing resources to allocate to each container running the job.
func (o JobDefinitionOutput) CpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.IntOutput { return v.CpuLimit }).(pulumi.IntOutput)
}

func (o JobDefinitionOutput) Cron() JobDefinitionCronPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionCronPtrOutput { return v.Cron }).(JobDefinitionCronPtrOutput)
}

// The description of the job
func (o JobDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment variables of the container.
func (o JobDefinitionOutput) Env() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringMapOutput { return v.Env }).(pulumi.StringMapOutput)
}

// The uri of the container image that will be used for the job run.
func (o JobDefinitionOutput) ImageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.ImageUri }).(pulumi.StringPtrOutput)
}

// The memory computing resources in MB to allocate to each container running the job.
func (o JobDefinitionOutput) MemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.IntOutput { return v.MemoryLimit }).(pulumi.IntOutput)
}

// The name of the job.
func (o JobDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the Job is associated with.
func (o JobDefinitionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region of the Job.
func (o JobDefinitionOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The job run timeout, in Go Time format (ex: `2h30m25s`)
func (o JobDefinitionOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

type JobDefinitionArrayOutput struct{ *pulumi.OutputState }

func (JobDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobDefinition)(nil)).Elem()
}

func (o JobDefinitionArrayOutput) ToJobDefinitionArrayOutput() JobDefinitionArrayOutput {
	return o
}

func (o JobDefinitionArrayOutput) ToJobDefinitionArrayOutputWithContext(ctx context.Context) JobDefinitionArrayOutput {
	return o
}

func (o JobDefinitionArrayOutput) Index(i pulumi.IntInput) JobDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *JobDefinition {
		return vs[0].([]*JobDefinition)[vs[1].(int)]
	}).(JobDefinitionOutput)
}

type JobDefinitionMapOutput struct{ *pulumi.OutputState }

func (JobDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobDefinition)(nil)).Elem()
}

func (o JobDefinitionMapOutput) ToJobDefinitionMapOutput() JobDefinitionMapOutput {
	return o
}

func (o JobDefinitionMapOutput) ToJobDefinitionMapOutputWithContext(ctx context.Context) JobDefinitionMapOutput {
	return o
}

func (o JobDefinitionMapOutput) MapIndex(k pulumi.StringInput) JobDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *JobDefinition {
		return vs[0].(map[string]*JobDefinition)[vs[1].(string)]
	}).(JobDefinitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobDefinitionInput)(nil)).Elem(), &JobDefinition{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobDefinitionArrayInput)(nil)).Elem(), JobDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobDefinitionMapInput)(nil)).Elem(), JobDefinitionMap{})
	pulumi.RegisterOutputType(JobDefinitionOutput{})
	pulumi.RegisterOutputType(JobDefinitionArrayOutput{})
	pulumi.RegisterOutputType(JobDefinitionMapOutput{})
}
