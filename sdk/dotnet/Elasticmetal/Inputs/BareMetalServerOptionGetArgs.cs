// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal.Inputs
{

    public sealed class BareMetalServerOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auto expiration date for compatible options
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The id of the private network to attach.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BareMetalServerOptionGetArgs()
        {
        }
        public static new BareMetalServerOptionGetArgs Empty => new BareMetalServerOptionGetArgs();
    }
}