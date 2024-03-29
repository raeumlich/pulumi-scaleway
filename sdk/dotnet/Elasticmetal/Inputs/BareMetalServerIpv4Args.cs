// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal.Inputs
{

    public sealed class BareMetalServerIpv4Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address of the IPv6.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The id of the private network to attach.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The reverse of the IPv6.
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// The type of the IPv6.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BareMetalServerIpv4Args()
        {
        }
        public static new BareMetalServerIpv4Args Empty => new BareMetalServerIpv4Args();
    }
}
