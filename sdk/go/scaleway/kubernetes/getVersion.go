// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a Kubernetes version.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/k8s/api).
//
// You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.
//
// ## Example Usage
// ### Use the latest version
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/kubernetes"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kubernetes.GetVersion(ctx, &kubernetes.GetVersionArgs{
//				Name: "latest",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Use a specific version
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/kubernetes"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kubernetes.GetVersion(ctx, &kubernetes.GetVersionArgs{
//				Name: "1.26.0",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVersion(ctx *pulumi.Context, args *GetVersionArgs, opts ...pulumi.InvokeOption) (*GetVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVersionResult
	err := ctx.Invoke("scaleway:kubernetes/getVersion:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVersion.
type GetVersionArgs struct {
	// The name of the Kubernetes version.
	Name string `pulumi:"name"`
	// `region`) The region in which the version exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVersion.
type GetVersionResult struct {
	// The list of supported Container Network Interface (CNI) plugins for this version.
	AvailableCnis []string `pulumi:"availableCnis"`
	// The list of supported container runtimes for this version.
	AvailableContainerRuntimes []string `pulumi:"availableContainerRuntimes"`
	// The list of supported feature gates for this version.
	AvailableFeatureGates []string `pulumi:"availableFeatureGates"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
}

func GetVersionOutput(ctx *pulumi.Context, args GetVersionOutputArgs, opts ...pulumi.InvokeOption) GetVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVersionResult, error) {
			args := v.(GetVersionArgs)
			r, err := GetVersion(ctx, &args, opts...)
			var s GetVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVersionResultOutput)
}

// A collection of arguments for invoking getVersion.
type GetVersionOutputArgs struct {
	// The name of the Kubernetes version.
	Name pulumi.StringInput `pulumi:"name"`
	// `region`) The region in which the version exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVersionArgs)(nil)).Elem()
}

// A collection of values returned by getVersion.
type GetVersionResultOutput struct{ *pulumi.OutputState }

func (GetVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVersionResult)(nil)).Elem()
}

func (o GetVersionResultOutput) ToGetVersionResultOutput() GetVersionResultOutput {
	return o
}

func (o GetVersionResultOutput) ToGetVersionResultOutputWithContext(ctx context.Context) GetVersionResultOutput {
	return o
}

// The list of supported Container Network Interface (CNI) plugins for this version.
func (o GetVersionResultOutput) AvailableCnis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVersionResult) []string { return v.AvailableCnis }).(pulumi.StringArrayOutput)
}

// The list of supported container runtimes for this version.
func (o GetVersionResultOutput) AvailableContainerRuntimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVersionResult) []string { return v.AvailableContainerRuntimes }).(pulumi.StringArrayOutput)
}

// The list of supported feature gates for this version.
func (o GetVersionResultOutput) AvailableFeatureGates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVersionResult) []string { return v.AvailableFeatureGates }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetVersionResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVersionResultOutput{})
}