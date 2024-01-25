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
    public sealed class GetBareMetalOfferDiskResult
    {
        /// <summary>
        /// Capacity of the memory in GB.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// Type of memory.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBareMetalOfferDiskResult(
            int capacity,

            string type)
        {
            Capacity = capacity;
            Type = type;
        }
    }
}
