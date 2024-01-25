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
// ### Basic
//
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
//			_, err = iot.NewDevice(ctx, "mainDevice", &iot.DeviceArgs{
//				HubId: mainHub.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With custom certificate
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-local/sdk/go/local"
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
//			deviceCert, err := local.LookupFile(ctx, &local.LookupFileArgs{
//				Filename: "device-certificate.pem",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iot.NewDevice(ctx, "mainDevice", &iot.DeviceArgs{
//				HubId: mainHub.ID(),
//				Certificate: &iot.DeviceCertificateArgs{
//					Crt: *pulumi.String(deviceCert.Content),
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
//
// ## Import
//
// IoT devices can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:iot/device:Device device01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type Device struct {
	pulumi.CustomResourceState

	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrOutput `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrOutput `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate DeviceCertificateOutput `pulumi:"certificate"`
	// The date and time the device was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringOutput `pulumi:"hubId"`
	// The current connection status of the device.
	IsConnected pulumi.BoolOutput `pulumi:"isConnected"`
	// The last MQTT activity of the device.
	LastActivityAt pulumi.StringOutput `pulumi:"lastActivityAt"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters DeviceMessageFiltersPtrOutput `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The current status of the device.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date and time the device resource was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubId == nil {
		return nil, errors.New("invalid value for required argument 'HubId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("scaleway:iot/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("scaleway:iot/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure *bool `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections *bool `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate *DeviceCertificate `pulumi:"certificate"`
	// The date and time the device was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the IoT device (e.g. `living room`).
	Description *string `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId *string `pulumi:"hubId"`
	// The current connection status of the device.
	IsConnected *bool `pulumi:"isConnected"`
	// The last MQTT activity of the device.
	LastActivityAt *string `pulumi:"lastActivityAt"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters *DeviceMessageFilters `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name *string `pulumi:"name"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The current status of the device.
	Status *string `pulumi:"status"`
	// The date and time the device resource was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DeviceState struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrInput
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrInput
	// The certificate bundle of the device.
	Certificate DeviceCertificatePtrInput
	// The date and time the device was created.
	CreatedAt pulumi.StringPtrInput
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrInput
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringPtrInput
	// The current connection status of the device.
	IsConnected pulumi.BoolPtrInput
	// The last MQTT activity of the device.
	LastActivityAt pulumi.StringPtrInput
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters DeviceMessageFiltersPtrInput
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The current status of the device.
	Status pulumi.StringPtrInput
	// The date and time the device resource was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure *bool `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections *bool `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate *DeviceCertificate `pulumi:"certificate"`
	// The description of the IoT device (e.g. `living room`).
	Description *string `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId string `pulumi:"hubId"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters *DeviceMessageFilters `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name *string `pulumi:"name"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrInput
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrInput
	// The certificate bundle of the device.
	Certificate DeviceCertificatePtrInput
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrInput
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringInput
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters DeviceMessageFiltersPtrInput
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

// DeviceArrayInput is an input type that accepts DeviceArray and DeviceArrayOutput values.
// You can construct a concrete instance of `DeviceArrayInput` via:
//
//	DeviceArray{ DeviceArgs{...} }
type DeviceArrayInput interface {
	pulumi.Input

	ToDeviceArrayOutput() DeviceArrayOutput
	ToDeviceArrayOutputWithContext(context.Context) DeviceArrayOutput
}

type DeviceArray []DeviceInput

func (DeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (i DeviceArray) ToDeviceArrayOutput() DeviceArrayOutput {
	return i.ToDeviceArrayOutputWithContext(context.Background())
}

func (i DeviceArray) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceArrayOutput)
}

// DeviceMapInput is an input type that accepts DeviceMap and DeviceMapOutput values.
// You can construct a concrete instance of `DeviceMapInput` via:
//
//	DeviceMap{ "key": DeviceArgs{...} }
type DeviceMapInput interface {
	pulumi.Input

	ToDeviceMapOutput() DeviceMapOutput
	ToDeviceMapOutputWithContext(context.Context) DeviceMapOutput
}

type DeviceMap map[string]DeviceInput

func (DeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (i DeviceMap) ToDeviceMapOutput() DeviceMapOutput {
	return i.ToDeviceMapOutputWithContext(context.Background())
}

func (i DeviceMap) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceMapOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
//
// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
func (o DeviceOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.BoolPtrOutput { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

// Allow more than one simultaneous connection using the same device credentials.
//
// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
func (o DeviceOutput) AllowMultipleConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.BoolPtrOutput { return v.AllowMultipleConnections }).(pulumi.BoolPtrOutput)
}

// The certificate bundle of the device.
func (o DeviceOutput) Certificate() DeviceCertificateOutput {
	return o.ApplyT(func(v *Device) DeviceCertificateOutput { return v.Certificate }).(DeviceCertificateOutput)
}

// The date and time the device was created.
func (o DeviceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the IoT device (e.g. `living room`).
func (o DeviceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the hub on which this device will be created.
func (o DeviceOutput) HubId() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.HubId }).(pulumi.StringOutput)
}

// The current connection status of the device.
func (o DeviceOutput) IsConnected() pulumi.BoolOutput {
	return o.ApplyT(func(v *Device) pulumi.BoolOutput { return v.IsConnected }).(pulumi.BoolOutput)
}

// The last MQTT activity of the device.
func (o DeviceOutput) LastActivityAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.LastActivityAt }).(pulumi.StringOutput)
}

// Rules that define which messages are authorized or denied based on their topic.
func (o DeviceOutput) MessageFilters() DeviceMessageFiltersPtrOutput {
	return o.ApplyT(func(v *Device) DeviceMessageFiltersPtrOutput { return v.MessageFilters }).(DeviceMessageFiltersPtrOutput)
}

// The name of the IoT device you want to create (e.g. `my-device`).
//
// > **Important:** Updates to `name` will destroy and recreate a new resource.
func (o DeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o DeviceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current status of the device.
func (o DeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date and time the device resource was updated.
func (o DeviceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Device {
		return vs[0].([]*Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceMapOutput struct{ *pulumi.OutputState }

func (DeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (o DeviceMapOutput) ToDeviceMapOutput() DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) MapIndex(k pulumi.StringInput) DeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Device {
		return vs[0].(map[string]*Device)[vs[1].(string)]
	}).(DeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceArrayInput)(nil)).Elem(), DeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceMapInput)(nil)).Elem(), DeviceMap{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceMapOutput{})
}
