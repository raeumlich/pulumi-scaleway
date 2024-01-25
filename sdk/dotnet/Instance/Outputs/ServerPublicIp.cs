// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance.Outputs
{

    [OutputType]
    public sealed class ServerPublicIp
    {
        /// <summary>
        /// The address of the IP
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The ID of the IP
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ServerPublicIp(
            string? address,

            string? id)
        {
            Address = address;
            Id = id;
        }
    }
}
