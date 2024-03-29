// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Vpc
{
    public static class GetGatewayNetwork
    {
        /// <summary>
        /// Gets information about a gateway network.
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
        ///     var main = new Scaleway.Vpc.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///         DhcpId = scaleway_vpc_public_gateway_dhcp.Dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var byId = Scaleway.Vpc.GetGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayNetworkId = main.Id,
        ///     });
        /// 
        ///     var byGatewayAndPn = Scaleway.Vpc.GetGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetGatewayNetworkResult> InvokeAsync(GetGatewayNetworkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayNetworkResult>("scaleway:vpc/getGatewayNetwork:getGatewayNetwork", args ?? new GetGatewayNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a gateway network.
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
        ///     var main = new Scaleway.Vpc.GatewayNetwork("main", new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///         DhcpId = scaleway_vpc_public_gateway_dhcp.Dhcp01.Id,
        ///         CleanupDhcp = true,
        ///         EnableMasquerade = true,
        ///     });
        /// 
        ///     var byId = Scaleway.Vpc.GetGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayNetworkId = main.Id,
        ///     });
        /// 
        ///     var byGatewayAndPn = Scaleway.Vpc.GetGatewayNetwork.Invoke(new()
        ///     {
        ///         GatewayId = scaleway_vpc_public_gateway.Pg01.Id,
        ///         PrivateNetworkId = scaleway_vpc_private_network.Pn01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetGatewayNetworkResult> Invoke(GetGatewayNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayNetworkResult>("scaleway:vpc/getGatewayNetwork:getGatewayNetwork", args ?? new GetGatewayNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the public gateway DHCP config
        /// </summary>
        [Input("dhcpId")]
        public string? DhcpId { get; set; }

        /// <summary>
        /// If masquerade is enabled on requested network
        /// </summary>
        [Input("enableMasquerade")]
        public bool? EnableMasquerade { get; set; }

        /// <summary>
        /// ID of the public gateway the gateway network is linked to
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        /// <summary>
        /// ID of the gateway network.
        /// 
        /// &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
        /// </summary>
        [Input("gatewayNetworkId")]
        public string? GatewayNetworkId { get; set; }

        /// <summary>
        /// ID of the private network the gateway network is linked to
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        public GetGatewayNetworkArgs()
        {
        }
        public static new GetGatewayNetworkArgs Empty => new GetGatewayNetworkArgs();
    }

    public sealed class GetGatewayNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the public gateway DHCP config
        /// </summary>
        [Input("dhcpId")]
        public Input<string>? DhcpId { get; set; }

        /// <summary>
        /// If masquerade is enabled on requested network
        /// </summary>
        [Input("enableMasquerade")]
        public Input<bool>? EnableMasquerade { get; set; }

        /// <summary>
        /// ID of the public gateway the gateway network is linked to
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// ID of the gateway network.
        /// 
        /// &gt; Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
        /// </summary>
        [Input("gatewayNetworkId")]
        public Input<string>? GatewayNetworkId { get; set; }

        /// <summary>
        /// ID of the private network the gateway network is linked to
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        public GetGatewayNetworkInvokeArgs()
        {
        }
        public static new GetGatewayNetworkInvokeArgs Empty => new GetGatewayNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayNetworkResult
    {
        public readonly bool CleanupDhcp;
        public readonly string CreatedAt;
        public readonly string? DhcpId;
        public readonly bool EnableDhcp;
        public readonly bool? EnableMasquerade;
        public readonly string? GatewayId;
        public readonly string? GatewayNetworkId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetGatewayNetworkIpamConfigResult> IpamConfigs;
        public readonly string MacAddress;
        public readonly string? PrivateNetworkId;
        public readonly string StaticAddress;
        public readonly string Status;
        public readonly string UpdatedAt;
        public readonly string Zone;

        [OutputConstructor]
        private GetGatewayNetworkResult(
            bool cleanupDhcp,

            string createdAt,

            string? dhcpId,

            bool enableDhcp,

            bool? enableMasquerade,

            string? gatewayId,

            string? gatewayNetworkId,

            string id,

            ImmutableArray<Outputs.GetGatewayNetworkIpamConfigResult> ipamConfigs,

            string macAddress,

            string? privateNetworkId,

            string staticAddress,

            string status,

            string updatedAt,

            string zone)
        {
            CleanupDhcp = cleanupDhcp;
            CreatedAt = createdAt;
            DhcpId = dhcpId;
            EnableDhcp = enableDhcp;
            EnableMasquerade = enableMasquerade;
            GatewayId = gatewayId;
            GatewayNetworkId = gatewayNetworkId;
            Id = id;
            IpamConfigs = ipamConfigs;
            MacAddress = macAddress;
            PrivateNetworkId = privateNetworkId;
            StaticAddress = staticAddress;
            Status = status;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
