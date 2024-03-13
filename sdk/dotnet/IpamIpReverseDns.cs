// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway
{
    /// <summary>
    /// Manages Scaleway IPAM IP Reverse DNS.
    /// 
    /// ## Import
    /// 
    /// IPAM IP reverse DNS can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/ipamIpReverseDns:IpamIpReverseDns main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/ipamIpReverseDns:IpamIpReverseDns")]
    public partial class IpamIpReverseDns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP corresponding to the hostname.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The reverse domain name.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The IPAM IP ID.
        /// </summary>
        [Output("ipamIpId")]
        public Output<string> IpamIpId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region of the IP reverse DNS.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a IpamIpReverseDns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpamIpReverseDns(string name, IpamIpReverseDnsArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/ipamIpReverseDns:IpamIpReverseDns", name, args ?? new IpamIpReverseDnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpamIpReverseDns(string name, Input<string> id, IpamIpReverseDnsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/ipamIpReverseDns:IpamIpReverseDns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpamIpReverseDns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpamIpReverseDns Get(string name, Input<string> id, IpamIpReverseDnsState? state = null, CustomResourceOptions? options = null)
        {
            return new IpamIpReverseDns(name, id, state, options);
        }
    }

    public sealed class IpamIpReverseDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP corresponding to the hostname.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// The reverse domain name.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// The IPAM IP ID.
        /// </summary>
        [Input("ipamIpId", required: true)]
        public Input<string> IpamIpId { get; set; } = null!;

        /// <summary>
        /// `region`) The region of the IP reverse DNS.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public IpamIpReverseDnsArgs()
        {
        }
        public static new IpamIpReverseDnsArgs Empty => new IpamIpReverseDnsArgs();
    }

    public sealed class IpamIpReverseDnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP corresponding to the hostname.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The reverse domain name.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The IPAM IP ID.
        /// </summary>
        [Input("ipamIpId")]
        public Input<string>? IpamIpId { get; set; }

        /// <summary>
        /// `region`) The region of the IP reverse DNS.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public IpamIpReverseDnsState()
        {
        }
        public static new IpamIpReverseDnsState Empty => new IpamIpReverseDnsState();
    }
}
