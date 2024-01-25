// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal.Outputs
{

    [OutputType]
    public sealed class GetFlexibleIPsIpMacAddressResult
    {
        /// <summary>
        /// The date on which the flexible IP was created (RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The MAC address ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The MAC address of the Virtual MAC.
        /// </summary>
        public readonly string MacAddress;
        /// <summary>
        /// The type of virtual MAC.
        /// </summary>
        public readonly string MacType;
        /// <summary>
        /// The status of virtual MAC.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The date on which the flexible IP was last updated (RFC 3339 format).
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// `zone`) The zone in which IPs exist.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetFlexibleIPsIpMacAddressResult(
            string createdAt,

            string id,

            string macAddress,

            string macType,

            string status,

            string updatedAt,

            string zone)
        {
            CreatedAt = createdAt;
            Id = id;
            MacAddress = macAddress;
            MacType = macType;
            Status = status;
            UpdatedAt = updatedAt;
            Zone = zone;
        }
    }
}
