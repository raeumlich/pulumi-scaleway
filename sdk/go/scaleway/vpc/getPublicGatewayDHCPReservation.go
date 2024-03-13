// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a dhcp entries. For further information please check the
// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6)
//
// ## Example Dynamic
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/instance"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainPrivateNetwork, err := vpc.NewPrivateNetwork(ctx, "mainPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			mainServer, err := instance.NewServer(ctx, "mainServer", &instance.ServerArgs{
//				Image: pulumi.String("ubuntu_jammy"),
//				Type:  pulumi.String("DEV1-S"),
//				Zone:  pulumi.String("fr-par-1"),
//			})
//			if err != nil {
//				return err
//			}
//			mainPrivateNIC, err := instance.NewPrivateNIC(ctx, "mainPrivateNIC", &instance.PrivateNICArgs{
//				ServerId:         mainServer.ID(),
//				PrivateNetworkId: mainPrivateNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainPublicGatewayIP, err := vpc.NewPublicGatewayIP(ctx, "mainPublicGatewayIP", nil)
//			if err != nil {
//				return err
//			}
//			mainPublicGatewayDHCP, err := vpc.NewPublicGatewayDHCP(ctx, "mainPublicGatewayDHCP", &vpc.PublicGatewayDHCPArgs{
//				Subnet: pulumi.String("192.168.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			mainPublicGateway, err := vpc.NewPublicGateway(ctx, "mainPublicGateway", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: mainPublicGatewayIP.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainGatewayNetwork, err := vpc.NewGatewayNetwork(ctx, "mainGatewayNetwork", &vpc.GatewayNetworkArgs{
//				GatewayId:        mainPublicGateway.ID(),
//				PrivateNetworkId: mainPrivateNetwork.ID(),
//				DhcpId:           mainPublicGatewayDHCP.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.LookupPublicGatewayDHCPReservationOutput(ctx, vpc.GetPublicGatewayDHCPReservationOutputArgs{
//				MacAddress:       mainPrivateNIC.MacAddress,
//				GatewayNetworkId: mainGatewayNetwork.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Example Static and PAT rule
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/instance"
//	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainPrivateNetwork, err := vpc.NewPrivateNetwork(ctx, "mainPrivateNetwork", nil)
//			if err != nil {
//				return err
//			}
//			mainSecurityGroup, err := instance.NewSecurityGroup(ctx, "mainSecurityGroup", &instance.SecurityGroupArgs{
//				InboundDefaultPolicy:  pulumi.String("drop"),
//				OutboundDefaultPolicy: pulumi.String("accept"),
//				InboundRules: instance.SecurityGroupInboundRuleArray{
//					&instance.SecurityGroupInboundRuleArgs{
//						Action: pulumi.String("accept"),
//						Port:   pulumi.Int(22),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			mainServer, err := instance.NewServer(ctx, "mainServer", &instance.ServerArgs{
//				Image:           pulumi.String("ubuntu_jammy"),
//				Type:            pulumi.String("DEV1-S"),
//				Zone:            pulumi.String("fr-par-1"),
//				SecurityGroupId: mainSecurityGroup.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainPrivateNIC, err := instance.NewPrivateNIC(ctx, "mainPrivateNIC", &instance.PrivateNICArgs{
//				ServerId:         mainServer.ID(),
//				PrivateNetworkId: mainPrivateNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainPublicGatewayIP, err := vpc.NewPublicGatewayIP(ctx, "mainPublicGatewayIP", nil)
//			if err != nil {
//				return err
//			}
//			mainPublicGatewayDHCP, err := vpc.NewPublicGatewayDHCP(ctx, "mainPublicGatewayDHCP", &vpc.PublicGatewayDHCPArgs{
//				Subnet: pulumi.String("192.168.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			mainPublicGateway, err := vpc.NewPublicGateway(ctx, "mainPublicGateway", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: mainPublicGatewayIP.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainGatewayNetwork, err := vpc.NewGatewayNetwork(ctx, "mainGatewayNetwork", &vpc.GatewayNetworkArgs{
//				GatewayId:        mainPublicGateway.ID(),
//				PrivateNetworkId: mainPrivateNetwork.ID(),
//				DhcpId:           mainPublicGatewayDHCP.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			mainPublicGatewayDHCPReservation, err := vpc.NewPublicGatewayDHCPReservation(ctx, "mainPublicGatewayDHCPReservation", &vpc.PublicGatewayDHCPReservationArgs{
//				GatewayNetworkId: mainGatewayNetwork.ID(),
//				MacAddress:       mainPrivateNIC.MacAddress,
//				IpAddress:        pulumi.String("192.168.1.4"),
//			})
//			if err != nil {
//				return err
//			}
//			// ## VPC PAT RULE
//			_, err = vpc.NewPublicGatewayPATRule(ctx, "mainPublicGatewayPATRule", &vpc.PublicGatewayPATRuleArgs{
//				GatewayId:   mainPublicGateway.ID(),
//				PrivateIp:   mainPublicGatewayDHCPReservation.IpAddress,
//				PrivatePort: pulumi.Int(22),
//				PublicPort:  pulumi.Int(2222),
//				Protocol:    pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.LookupPublicGatewayDHCPReservationOutput(ctx, vpc.GetPublicGatewayDHCPReservationOutputArgs{
//				ReservationId: mainPublicGatewayDHCPReservation.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupPublicGatewayDHCPReservation(ctx *pulumi.Context, args *LookupPublicGatewayDHCPReservationArgs, opts ...pulumi.InvokeOption) (*LookupPublicGatewayDHCPReservationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicGatewayDHCPReservationResult
	err := ctx.Invoke("scaleway:vpc/getPublicGatewayDHCPReservation:getPublicGatewayDHCPReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicGatewayDHCPReservation.
type LookupPublicGatewayDHCPReservationArgs struct {
	// The ID of the owning GatewayNetwork.
	//
	// > Only one of `reservationId` or `macAddress` with `gatewayNetworkId` should be specified.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the Reservation to retrieve.
	ReservationId *string `pulumi:"reservationId"`
	// Boolean to wait for macAddress to exist in dhcp.
	WaitForDhcp *bool `pulumi:"waitForDhcp"`
	// `zone`) The zone in which the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getPublicGatewayDHCPReservation.
type LookupPublicGatewayDHCPReservationResult struct {
	CreatedAt        string  `pulumi:"createdAt"`
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	Hostname         string  `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	IpAddress     string  `pulumi:"ipAddress"`
	MacAddress    *string `pulumi:"macAddress"`
	ReservationId *string `pulumi:"reservationId"`
	Type          string  `pulumi:"type"`
	UpdatedAt     string  `pulumi:"updatedAt"`
	WaitForDhcp   *bool   `pulumi:"waitForDhcp"`
	Zone          *string `pulumi:"zone"`
}

func LookupPublicGatewayDHCPReservationOutput(ctx *pulumi.Context, args LookupPublicGatewayDHCPReservationOutputArgs, opts ...pulumi.InvokeOption) LookupPublicGatewayDHCPReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicGatewayDHCPReservationResult, error) {
			args := v.(LookupPublicGatewayDHCPReservationArgs)
			r, err := LookupPublicGatewayDHCPReservation(ctx, &args, opts...)
			var s LookupPublicGatewayDHCPReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicGatewayDHCPReservationResultOutput)
}

// A collection of arguments for invoking getPublicGatewayDHCPReservation.
type LookupPublicGatewayDHCPReservationOutputArgs struct {
	// The ID of the owning GatewayNetwork.
	//
	// > Only one of `reservationId` or `macAddress` with `gatewayNetworkId` should be specified.
	GatewayNetworkId pulumi.StringPtrInput `pulumi:"gatewayNetworkId"`
	// The MAC address of the reservation to retrieve.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The ID of the Reservation to retrieve.
	ReservationId pulumi.StringPtrInput `pulumi:"reservationId"`
	// Boolean to wait for macAddress to exist in dhcp.
	WaitForDhcp pulumi.BoolPtrInput `pulumi:"waitForDhcp"`
	// `zone`) The zone in which the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupPublicGatewayDHCPReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayDHCPReservationArgs)(nil)).Elem()
}

// A collection of values returned by getPublicGatewayDHCPReservation.
type LookupPublicGatewayDHCPReservationResultOutput struct{ *pulumi.OutputState }

func (LookupPublicGatewayDHCPReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayDHCPReservationResult)(nil)).Elem()
}

func (o LookupPublicGatewayDHCPReservationResultOutput) ToLookupPublicGatewayDHCPReservationResultOutput() LookupPublicGatewayDHCPReservationResultOutput {
	return o
}

func (o LookupPublicGatewayDHCPReservationResultOutput) ToLookupPublicGatewayDHCPReservationResultOutputWithContext(ctx context.Context) LookupPublicGatewayDHCPReservationResultOutput {
	return o
}

func (o LookupPublicGatewayDHCPReservationResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) GatewayNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) *string { return v.GatewayNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPublicGatewayDHCPReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) ReservationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) *string { return v.ReservationId }).(pulumi.StringPtrOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) WaitForDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) *bool { return v.WaitForDhcp }).(pulumi.BoolPtrOutput)
}

func (o LookupPublicGatewayDHCPReservationResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayDHCPReservationResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicGatewayDHCPReservationResultOutput{})
}
