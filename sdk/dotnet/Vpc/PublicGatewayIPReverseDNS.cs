// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Vpc
{
    /// <summary>
    /// Manages Scaleway VPC Public Gateways IPs reverse DNS.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
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
    ///     var mainPublicGatewayIP = new Scaleway.Vpc.PublicGatewayIP("mainPublicGatewayIP");
    /// 
    ///     var tfA = new Scaleway.Dns.Record("tfA", new()
    ///     {
    ///         DnsZone = "example.com",
    ///         Type = "A",
    ///         Data = mainPublicGatewayIP.Address,
    ///         Ttl = 3600,
    ///         Priority = 1,
    ///     });
    /// 
    ///     var mainPublicGatewayIPReverseDNS = new Scaleway.Vpc.PublicGatewayIPReverseDNS("mainPublicGatewayIPReverseDNS", new()
    ///     {
    ///         GatewayIpId = mainPublicGatewayIP.Id,
    ///         Reverse = "tf.example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Public gateway IPs reverse DNS can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:vpc/publicGatewayIPReverseDNS:PublicGatewayIPReverseDNS reverse fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:vpc/publicGatewayIPReverseDNS:PublicGatewayIPReverseDNS")]
    public partial class PublicGatewayIPReverseDNS : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Output("gatewayIpId")]
        public Output<string> GatewayIpId { get; private set; } = null!;

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a PublicGatewayIPReverseDNS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicGatewayIPReverseDNS(string name, PublicGatewayIPReverseDNSArgs args, CustomResourceOptions? options = null)
            : base("scaleway:vpc/publicGatewayIPReverseDNS:PublicGatewayIPReverseDNS", name, args ?? new PublicGatewayIPReverseDNSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicGatewayIPReverseDNS(string name, Input<string> id, PublicGatewayIPReverseDNSState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:vpc/publicGatewayIPReverseDNS:PublicGatewayIPReverseDNS", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/raeumlich/pulumi-scaleway/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PublicGatewayIPReverseDNS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicGatewayIPReverseDNS Get(string name, Input<string> id, PublicGatewayIPReverseDNSState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicGatewayIPReverseDNS(name, id, state, options);
        }
    }

    public sealed class PublicGatewayIPReverseDNSArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Input("gatewayIpId", required: true)]
        public Input<string> GatewayIpId { get; set; } = null!;

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Input("reverse", required: true)]
        public Input<string> Reverse { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PublicGatewayIPReverseDNSArgs()
        {
        }
        public static new PublicGatewayIPReverseDNSArgs Empty => new PublicGatewayIPReverseDNSArgs();
    }

    public sealed class PublicGatewayIPReverseDNSState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Input("gatewayIpId")]
        public Input<string>? GatewayIpId { get; set; }

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public PublicGatewayIPReverseDNSState()
        {
        }
        public static new PublicGatewayIPReverseDNSState Empty => new PublicGatewayIPReverseDNSState();
    }
}
