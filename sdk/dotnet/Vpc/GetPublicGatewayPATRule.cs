// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Vpc
{
    public static class GetPublicGatewayPATRule
    {
        /// <summary>
        /// Gets information about a public gateway PAT rule. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sg01 = new Scaleway.Instance.SecurityGroup("sg01", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///                 Protocol = "TCP",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var srv01 = new Scaleway.Instance.Server("srv01", new()
        ///     {
        ///         Type = "PLAY2-NANO",
        ///         Image = "ubuntu_jammy",
        ///         SecurityGroupId = sg01.Id,
        ///     });
        /// 
        ///     var pn01 = new Scaleway.Vpc.PrivateNetwork("pn01");
        /// 
        ///     var pnic01 = new Scaleway.Instance.PrivateNIC("pnic01", new()
        ///     {
        ///         ServerId = srv01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///     });
        /// 
        ///     var dhcp01 = new Scaleway.Vpc.PublicGatewayDHCP("dhcp01", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var ip01 = new Scaleway.Vpc.PublicGatewayIP("ip01");
        /// 
        ///     var pg01 = new Scaleway.Vpc.PublicGateway("pg01", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = ip01.Id,
        ///     });
        /// 
        ///     var gn01 = new Scaleway.Vpc.GatewayNetwork("gn01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///         DhcpId = dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var rsv01 = new Scaleway.Vpc.PublicGatewayDHCPReservation("rsv01", new()
        ///     {
        ///         GatewayNetworkId = gn01.Id,
        ///         MacAddress = pnic01.MacAddress,
        ///         IpAddress = "192.168.0.7",
        ///     });
        /// 
        ///     var pat01 = new Scaleway.Vpc.PublicGatewayPATRule("pat01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateIp = rsv01.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2202,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var main = Scaleway.Vpc.GetPublicGatewayPATRule.Invoke(new()
        ///     {
        ///         PatRuleId = pat01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPublicGatewayPATRuleResult> InvokeAsync(GetPublicGatewayPATRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublicGatewayPATRuleResult>("scaleway:vpc/getPublicGatewayPATRule:getPublicGatewayPATRule", args ?? new GetPublicGatewayPATRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a public gateway PAT rule. For further information please check the
        /// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sg01 = new Scaleway.Instance.SecurityGroup("sg01", new()
        ///     {
        ///         InboundDefaultPolicy = "drop",
        ///         OutboundDefaultPolicy = "accept",
        ///         InboundRules = new[]
        ///         {
        ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
        ///             {
        ///                 Action = "accept",
        ///                 Port = 22,
        ///                 Protocol = "TCP",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var srv01 = new Scaleway.Instance.Server("srv01", new()
        ///     {
        ///         Type = "PLAY2-NANO",
        ///         Image = "ubuntu_jammy",
        ///         SecurityGroupId = sg01.Id,
        ///     });
        /// 
        ///     var pn01 = new Scaleway.Vpc.PrivateNetwork("pn01");
        /// 
        ///     var pnic01 = new Scaleway.Instance.PrivateNIC("pnic01", new()
        ///     {
        ///         ServerId = srv01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///     });
        /// 
        ///     var dhcp01 = new Scaleway.Vpc.PublicGatewayDHCP("dhcp01", new()
        ///     {
        ///         Subnet = "192.168.0.0/24",
        ///     });
        /// 
        ///     var ip01 = new Scaleway.Vpc.PublicGatewayIP("ip01");
        /// 
        ///     var pg01 = new Scaleway.Vpc.PublicGateway("pg01", new()
        ///     {
        ///         Type = "VPC-GW-S",
        ///         IpId = ip01.Id,
        ///     });
        /// 
        ///     var gn01 = new Scaleway.Vpc.GatewayNetwork("gn01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateNetworkId = pn01.Id,
        ///         DhcpId = dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var rsv01 = new Scaleway.Vpc.PublicGatewayDHCPReservation("rsv01", new()
        ///     {
        ///         GatewayNetworkId = gn01.Id,
        ///         MacAddress = pnic01.MacAddress,
        ///         IpAddress = "192.168.0.7",
        ///     });
        /// 
        ///     var pat01 = new Scaleway.Vpc.PublicGatewayPATRule("pat01", new()
        ///     {
        ///         GatewayId = pg01.Id,
        ///         PrivateIp = rsv01.IpAddress,
        ///         PrivatePort = 22,
        ///         PublicPort = 2202,
        ///         Protocol = "tcp",
        ///     });
        /// 
        ///     var main = Scaleway.Vpc.GetPublicGatewayPATRule.Invoke(new()
        ///     {
        ///         PatRuleId = pat01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPublicGatewayPATRuleResult> Invoke(GetPublicGatewayPATRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublicGatewayPATRuleResult>("scaleway:vpc/getPublicGatewayPATRule:getPublicGatewayPATRule", args ?? new GetPublicGatewayPATRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicGatewayPATRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PAT rule to retrieve
        /// </summary>
        [Input("patRuleId", required: true)]
        public string PatRuleId { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which
        /// the image exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetPublicGatewayPATRuleArgs()
        {
        }
        public static new GetPublicGatewayPATRuleArgs Empty => new GetPublicGatewayPATRuleArgs();
    }

    public sealed class GetPublicGatewayPATRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the PAT rule to retrieve
        /// </summary>
        [Input("patRuleId", required: true)]
        public Input<string> PatRuleId { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which
        /// the image exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetPublicGatewayPATRuleInvokeArgs()
        {
        }
        public static new GetPublicGatewayPATRuleInvokeArgs Empty => new GetPublicGatewayPATRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublicGatewayPATRuleResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        public readonly string PatRuleId;
        /// <summary>
        /// The Private IP to forward data to (IP address).
        /// </summary>
        public readonly string PrivateIp;
        /// <summary>
        /// The Private port to translate to.
        /// </summary>
        public readonly int PrivatePort;
        /// <summary>
        /// The Protocol the rule should apply to. Possible values are both, tcp and udp.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The Public port to listen on.
        /// </summary>
        public readonly int PublicPort;
        public readonly string UpdatedAt;
        public readonly string? Zone;

        [OutputConstructor]
        private GetPublicGatewayPATRuleResult(
            string createdAt,

            string gatewayId,

            string id,

            string organizationId,

            string patRuleId,

            string privateIp,

            int privatePort,

            string protocol,

            int publicPort,

            string updatedAt,

            string? zone)
        {
            CreatedAt = createdAt;
            GatewayId = gatewayId;
            Id = id;
            OrganizationId = organizationId;
            PatRuleId = patRuleId;
            PrivateIp = privateIp;
            PrivatePort = privatePort;
            Protocol = protocol;
            PublicPort = publicPort;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
