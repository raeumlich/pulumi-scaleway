// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Ipam.Inputs
{

    public sealed class IPReverseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Request a specific IP in the requested source pool.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The reverse domain name.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        public IPReverseArgs()
        {
        }
        public static new IPReverseArgs Empty => new IPReverseArgs();
    }
}
