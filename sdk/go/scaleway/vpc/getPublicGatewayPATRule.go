// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/raeumlich/pulumi-scaleway/sdk/go/scaleway/internal"
)

// Gets information about a public gateway PAT rule. For further information please check the
// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
//
// ## Example Usage
//
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
//			sg01, err := instance.NewSecurityGroup(ctx, "sg01", &instance.SecurityGroupArgs{
//				InboundDefaultPolicy:  pulumi.String("drop"),
//				OutboundDefaultPolicy: pulumi.String("accept"),
//				InboundRules: instance.SecurityGroupInboundRuleArray{
//					&instance.SecurityGroupInboundRuleArgs{
//						Action:   pulumi.String("accept"),
//						Port:     pulumi.Int(22),
//						Protocol: pulumi.String("TCP"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			srv01, err := instance.NewServer(ctx, "srv01", &instance.ServerArgs{
//				Type:            pulumi.String("PLAY2-NANO"),
//				Image:           pulumi.String("ubuntu_jammy"),
//				SecurityGroupId: sg01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			pn01, err := vpc.NewPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			pnic01, err := instance.NewPrivateNIC(ctx, "pnic01", &instance.PrivateNICArgs{
//				ServerId:         srv01.ID(),
//				PrivateNetworkId: pn01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			dhcp01, err := vpc.NewPublicGatewayDHCP(ctx, "dhcp01", &vpc.PublicGatewayDHCPArgs{
//				Subnet: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := vpc.NewPublicGatewayIP(ctx, "ip01", nil)
//			if err != nil {
//				return err
//			}
//			pg01, err := vpc.NewPublicGateway(ctx, "pg01", &vpc.PublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: ip01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			gn01, err := vpc.NewGatewayNetwork(ctx, "gn01", &vpc.GatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				DhcpId:           dhcp01.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			rsv01, err := vpc.NewPublicGatewayDHCPReservation(ctx, "rsv01", &vpc.PublicGatewayDHCPReservationArgs{
//				GatewayNetworkId: gn01.ID(),
//				MacAddress:       pnic01.MacAddress,
//				IpAddress:        pulumi.String("192.168.0.7"),
//			})
//			if err != nil {
//				return err
//			}
//			pat01, err := vpc.NewPublicGatewayPATRule(ctx, "pat01", &vpc.PublicGatewayPATRuleArgs{
//				GatewayId:   pg01.ID(),
//				PrivateIp:   rsv01.IpAddress,
//				PrivatePort: pulumi.Int(22),
//				PublicPort:  pulumi.Int(2202),
//				Protocol:    pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.LookupPublicGatewayPATRuleOutput(ctx, vpc.GetPublicGatewayPATRuleOutputArgs{
//				PatRuleId: pat01.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupPublicGatewayPATRule(ctx *pulumi.Context, args *LookupPublicGatewayPATRuleArgs, opts ...pulumi.InvokeOption) (*LookupPublicGatewayPATRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPublicGatewayPATRuleResult
	err := ctx.Invoke("scaleway:vpc/getPublicGatewayPATRule:getPublicGatewayPATRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicGatewayPATRule.
type LookupPublicGatewayPATRuleArgs struct {
	// The ID of the PAT rule to retrieve
	PatRuleId string `pulumi:"patRuleId"`
	// `zone`) The zone in which
	// the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getPublicGatewayPATRule.
type LookupPublicGatewayPATRuleResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The ID of the public gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	OrganizationId string `pulumi:"organizationId"`
	PatRuleId      string `pulumi:"patRuleId"`
	// The Private IP to forward data to (IP address).
	PrivateIp string `pulumi:"privateIp"`
	// The Private port to translate to.
	PrivatePort int `pulumi:"privatePort"`
	// The Protocol the rule should apply to. Possible values are both, tcp and udp.
	Protocol string `pulumi:"protocol"`
	// The Public port to listen on.
	PublicPort int     `pulumi:"publicPort"`
	UpdatedAt  string  `pulumi:"updatedAt"`
	Zone       *string `pulumi:"zone"`
}

func LookupPublicGatewayPATRuleOutput(ctx *pulumi.Context, args LookupPublicGatewayPATRuleOutputArgs, opts ...pulumi.InvokeOption) LookupPublicGatewayPATRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicGatewayPATRuleResult, error) {
			args := v.(LookupPublicGatewayPATRuleArgs)
			r, err := LookupPublicGatewayPATRule(ctx, &args, opts...)
			var s LookupPublicGatewayPATRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicGatewayPATRuleResultOutput)
}

// A collection of arguments for invoking getPublicGatewayPATRule.
type LookupPublicGatewayPATRuleOutputArgs struct {
	// The ID of the PAT rule to retrieve
	PatRuleId pulumi.StringInput `pulumi:"patRuleId"`
	// `zone`) The zone in which
	// the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupPublicGatewayPATRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayPATRuleArgs)(nil)).Elem()
}

// A collection of values returned by getPublicGatewayPATRule.
type LookupPublicGatewayPATRuleResultOutput struct{ *pulumi.OutputState }

func (LookupPublicGatewayPATRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicGatewayPATRuleResult)(nil)).Elem()
}

func (o LookupPublicGatewayPATRuleResultOutput) ToLookupPublicGatewayPATRuleResultOutput() LookupPublicGatewayPATRuleResultOutput {
	return o
}

func (o LookupPublicGatewayPATRuleResultOutput) ToLookupPublicGatewayPATRuleResultOutputWithContext(ctx context.Context) LookupPublicGatewayPATRuleResultOutput {
	return o
}

func (o LookupPublicGatewayPATRuleResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the public gateway.
func (o LookupPublicGatewayPATRuleResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPublicGatewayPATRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayPATRuleResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayPATRuleResultOutput) PatRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.PatRuleId }).(pulumi.StringOutput)
}

// The Private IP to forward data to (IP address).
func (o LookupPublicGatewayPATRuleResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// The Private port to translate to.
func (o LookupPublicGatewayPATRuleResultOutput) PrivatePort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) int { return v.PrivatePort }).(pulumi.IntOutput)
}

// The Protocol the rule should apply to. Possible values are both, tcp and udp.
func (o LookupPublicGatewayPATRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// The Public port to listen on.
func (o LookupPublicGatewayPATRuleResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupPublicGatewayPATRuleResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupPublicGatewayPATRuleResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicGatewayPATRuleResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicGatewayPATRuleResultOutput{})
}
