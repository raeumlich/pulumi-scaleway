// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage.Outputs
{

    [OutputType]
    public sealed class GetBucketLifecycleRuleTransitionResult
    {
        /// <summary>
        /// Specifies the number of days after object creation when the specific rule action takes effect
        /// </summary>
        public readonly int Days;
        /// <summary>
        /// Specifies the Scaleway Object Storage class to which you want the object to transition
        /// </summary>
        public readonly string StorageClass;

        [OutputConstructor]
        private GetBucketLifecycleRuleTransitionResult(
            int days,

            string storageClass)
        {
            Days = days;
            StorageClass = storageClass;
        }
    }
}
