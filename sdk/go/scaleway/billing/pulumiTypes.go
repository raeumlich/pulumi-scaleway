// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

var _ = internal.GetEnvOrDefault

type GetConsumptionsConsumption struct {
	// The category of the consumption.
	Category string `pulumi:"category"`
	// The description of the consumption.
	Description string `pulumi:"description"`
	// The unique identifier of the product.
	OperationPath string `pulumi:"operationPath"`
	// The project ID of the consumption.
	ProjectId string `pulumi:"projectId"`
	// The monetary value of the consumption.
	Value string `pulumi:"value"`
}

// GetConsumptionsConsumptionInput is an input type that accepts GetConsumptionsConsumptionArgs and GetConsumptionsConsumptionOutput values.
// You can construct a concrete instance of `GetConsumptionsConsumptionInput` via:
//
//	GetConsumptionsConsumptionArgs{...}
type GetConsumptionsConsumptionInput interface {
	pulumi.Input

	ToGetConsumptionsConsumptionOutput() GetConsumptionsConsumptionOutput
	ToGetConsumptionsConsumptionOutputWithContext(context.Context) GetConsumptionsConsumptionOutput
}

type GetConsumptionsConsumptionArgs struct {
	// The category of the consumption.
	Category pulumi.StringInput `pulumi:"category"`
	// The description of the consumption.
	Description pulumi.StringInput `pulumi:"description"`
	// The unique identifier of the product.
	OperationPath pulumi.StringInput `pulumi:"operationPath"`
	// The project ID of the consumption.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The monetary value of the consumption.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetConsumptionsConsumptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConsumptionsConsumption)(nil)).Elem()
}

func (i GetConsumptionsConsumptionArgs) ToGetConsumptionsConsumptionOutput() GetConsumptionsConsumptionOutput {
	return i.ToGetConsumptionsConsumptionOutputWithContext(context.Background())
}

func (i GetConsumptionsConsumptionArgs) ToGetConsumptionsConsumptionOutputWithContext(ctx context.Context) GetConsumptionsConsumptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetConsumptionsConsumptionOutput)
}

// GetConsumptionsConsumptionArrayInput is an input type that accepts GetConsumptionsConsumptionArray and GetConsumptionsConsumptionArrayOutput values.
// You can construct a concrete instance of `GetConsumptionsConsumptionArrayInput` via:
//
//	GetConsumptionsConsumptionArray{ GetConsumptionsConsumptionArgs{...} }
type GetConsumptionsConsumptionArrayInput interface {
	pulumi.Input

	ToGetConsumptionsConsumptionArrayOutput() GetConsumptionsConsumptionArrayOutput
	ToGetConsumptionsConsumptionArrayOutputWithContext(context.Context) GetConsumptionsConsumptionArrayOutput
}

type GetConsumptionsConsumptionArray []GetConsumptionsConsumptionInput

func (GetConsumptionsConsumptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetConsumptionsConsumption)(nil)).Elem()
}

func (i GetConsumptionsConsumptionArray) ToGetConsumptionsConsumptionArrayOutput() GetConsumptionsConsumptionArrayOutput {
	return i.ToGetConsumptionsConsumptionArrayOutputWithContext(context.Background())
}

func (i GetConsumptionsConsumptionArray) ToGetConsumptionsConsumptionArrayOutputWithContext(ctx context.Context) GetConsumptionsConsumptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetConsumptionsConsumptionArrayOutput)
}

type GetConsumptionsConsumptionOutput struct{ *pulumi.OutputState }

func (GetConsumptionsConsumptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConsumptionsConsumption)(nil)).Elem()
}

func (o GetConsumptionsConsumptionOutput) ToGetConsumptionsConsumptionOutput() GetConsumptionsConsumptionOutput {
	return o
}

func (o GetConsumptionsConsumptionOutput) ToGetConsumptionsConsumptionOutputWithContext(ctx context.Context) GetConsumptionsConsumptionOutput {
	return o
}

// The category of the consumption.
func (o GetConsumptionsConsumptionOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsConsumption) string { return v.Category }).(pulumi.StringOutput)
}

// The description of the consumption.
func (o GetConsumptionsConsumptionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsConsumption) string { return v.Description }).(pulumi.StringOutput)
}

// The unique identifier of the product.
func (o GetConsumptionsConsumptionOutput) OperationPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsConsumption) string { return v.OperationPath }).(pulumi.StringOutput)
}

// The project ID of the consumption.
func (o GetConsumptionsConsumptionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsConsumption) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The monetary value of the consumption.
func (o GetConsumptionsConsumptionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsumptionsConsumption) string { return v.Value }).(pulumi.StringOutput)
}

type GetConsumptionsConsumptionArrayOutput struct{ *pulumi.OutputState }

func (GetConsumptionsConsumptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetConsumptionsConsumption)(nil)).Elem()
}

func (o GetConsumptionsConsumptionArrayOutput) ToGetConsumptionsConsumptionArrayOutput() GetConsumptionsConsumptionArrayOutput {
	return o
}

func (o GetConsumptionsConsumptionArrayOutput) ToGetConsumptionsConsumptionArrayOutputWithContext(ctx context.Context) GetConsumptionsConsumptionArrayOutput {
	return o
}

func (o GetConsumptionsConsumptionArrayOutput) Index(i pulumi.IntInput) GetConsumptionsConsumptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetConsumptionsConsumption {
		return vs[0].([]GetConsumptionsConsumption)[vs[1].(int)]
	}).(GetConsumptionsConsumptionOutput)
}

type GetInvoicesInvoice struct {
	// The payment time limit, set according to the Organization's payment conditions (RFC 3339 format).
	DueDate string `pulumi:"dueDate"`
	// The associated invoice ID.
	Id string `pulumi:"id"`
	// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
	InvoiceType string `pulumi:"invoiceType"`
	// The date when the invoice was sent to the customer (RFC 3339 format).
	IssuedDate string `pulumi:"issuedDate"`
	// The invoice number.
	Number int `pulumi:"number"`
	// The start date of the billing period (RFC 3339 format).
	StartDate string `pulumi:"startDate"`
	// The total amount, taxed.
	TotalTaxed string `pulumi:"totalTaxed"`
	// The total amount, untaxed.
	TotalUntaxed string `pulumi:"totalUntaxed"`
}

// GetInvoicesInvoiceInput is an input type that accepts GetInvoicesInvoiceArgs and GetInvoicesInvoiceOutput values.
// You can construct a concrete instance of `GetInvoicesInvoiceInput` via:
//
//	GetInvoicesInvoiceArgs{...}
type GetInvoicesInvoiceInput interface {
	pulumi.Input

	ToGetInvoicesInvoiceOutput() GetInvoicesInvoiceOutput
	ToGetInvoicesInvoiceOutputWithContext(context.Context) GetInvoicesInvoiceOutput
}

type GetInvoicesInvoiceArgs struct {
	// The payment time limit, set according to the Organization's payment conditions (RFC 3339 format).
	DueDate pulumi.StringInput `pulumi:"dueDate"`
	// The associated invoice ID.
	Id pulumi.StringInput `pulumi:"id"`
	// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
	InvoiceType pulumi.StringInput `pulumi:"invoiceType"`
	// The date when the invoice was sent to the customer (RFC 3339 format).
	IssuedDate pulumi.StringInput `pulumi:"issuedDate"`
	// The invoice number.
	Number pulumi.IntInput `pulumi:"number"`
	// The start date of the billing period (RFC 3339 format).
	StartDate pulumi.StringInput `pulumi:"startDate"`
	// The total amount, taxed.
	TotalTaxed pulumi.StringInput `pulumi:"totalTaxed"`
	// The total amount, untaxed.
	TotalUntaxed pulumi.StringInput `pulumi:"totalUntaxed"`
}

func (GetInvoicesInvoiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInvoicesInvoice)(nil)).Elem()
}

func (i GetInvoicesInvoiceArgs) ToGetInvoicesInvoiceOutput() GetInvoicesInvoiceOutput {
	return i.ToGetInvoicesInvoiceOutputWithContext(context.Background())
}

func (i GetInvoicesInvoiceArgs) ToGetInvoicesInvoiceOutputWithContext(ctx context.Context) GetInvoicesInvoiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInvoicesInvoiceOutput)
}

// GetInvoicesInvoiceArrayInput is an input type that accepts GetInvoicesInvoiceArray and GetInvoicesInvoiceArrayOutput values.
// You can construct a concrete instance of `GetInvoicesInvoiceArrayInput` via:
//
//	GetInvoicesInvoiceArray{ GetInvoicesInvoiceArgs{...} }
type GetInvoicesInvoiceArrayInput interface {
	pulumi.Input

	ToGetInvoicesInvoiceArrayOutput() GetInvoicesInvoiceArrayOutput
	ToGetInvoicesInvoiceArrayOutputWithContext(context.Context) GetInvoicesInvoiceArrayOutput
}

type GetInvoicesInvoiceArray []GetInvoicesInvoiceInput

func (GetInvoicesInvoiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInvoicesInvoice)(nil)).Elem()
}

func (i GetInvoicesInvoiceArray) ToGetInvoicesInvoiceArrayOutput() GetInvoicesInvoiceArrayOutput {
	return i.ToGetInvoicesInvoiceArrayOutputWithContext(context.Background())
}

func (i GetInvoicesInvoiceArray) ToGetInvoicesInvoiceArrayOutputWithContext(ctx context.Context) GetInvoicesInvoiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInvoicesInvoiceArrayOutput)
}

type GetInvoicesInvoiceOutput struct{ *pulumi.OutputState }

func (GetInvoicesInvoiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInvoicesInvoice)(nil)).Elem()
}

func (o GetInvoicesInvoiceOutput) ToGetInvoicesInvoiceOutput() GetInvoicesInvoiceOutput {
	return o
}

func (o GetInvoicesInvoiceOutput) ToGetInvoicesInvoiceOutputWithContext(ctx context.Context) GetInvoicesInvoiceOutput {
	return o
}

// The payment time limit, set according to the Organization's payment conditions (RFC 3339 format).
func (o GetInvoicesInvoiceOutput) DueDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.DueDate }).(pulumi.StringOutput)
}

// The associated invoice ID.
func (o GetInvoicesInvoiceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.Id }).(pulumi.StringOutput)
}

// Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
func (o GetInvoicesInvoiceOutput) InvoiceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.InvoiceType }).(pulumi.StringOutput)
}

// The date when the invoice was sent to the customer (RFC 3339 format).
func (o GetInvoicesInvoiceOutput) IssuedDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.IssuedDate }).(pulumi.StringOutput)
}

// The invoice number.
func (o GetInvoicesInvoiceOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) int { return v.Number }).(pulumi.IntOutput)
}

// The start date of the billing period (RFC 3339 format).
func (o GetInvoicesInvoiceOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.StartDate }).(pulumi.StringOutput)
}

// The total amount, taxed.
func (o GetInvoicesInvoiceOutput) TotalTaxed() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.TotalTaxed }).(pulumi.StringOutput)
}

// The total amount, untaxed.
func (o GetInvoicesInvoiceOutput) TotalUntaxed() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvoicesInvoice) string { return v.TotalUntaxed }).(pulumi.StringOutput)
}

type GetInvoicesInvoiceArrayOutput struct{ *pulumi.OutputState }

func (GetInvoicesInvoiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInvoicesInvoice)(nil)).Elem()
}

func (o GetInvoicesInvoiceArrayOutput) ToGetInvoicesInvoiceArrayOutput() GetInvoicesInvoiceArrayOutput {
	return o
}

func (o GetInvoicesInvoiceArrayOutput) ToGetInvoicesInvoiceArrayOutputWithContext(ctx context.Context) GetInvoicesInvoiceArrayOutput {
	return o
}

func (o GetInvoicesInvoiceArrayOutput) Index(i pulumi.IntInput) GetInvoicesInvoiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInvoicesInvoice {
		return vs[0].([]GetInvoicesInvoice)[vs[1].(int)]
	}).(GetInvoicesInvoiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetConsumptionsConsumptionInput)(nil)).Elem(), GetConsumptionsConsumptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetConsumptionsConsumptionArrayInput)(nil)).Elem(), GetConsumptionsConsumptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInvoicesInvoiceInput)(nil)).Elem(), GetInvoicesInvoiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInvoicesInvoiceArrayInput)(nil)).Elem(), GetInvoicesInvoiceArray{})
	pulumi.RegisterOutputType(GetConsumptionsConsumptionOutput{})
	pulumi.RegisterOutputType(GetConsumptionsConsumptionArrayOutput{})
	pulumi.RegisterOutputType(GetInvoicesInvoiceOutput{})
	pulumi.RegisterOutputType(GetInvoicesInvoiceArrayOutput{})
}
