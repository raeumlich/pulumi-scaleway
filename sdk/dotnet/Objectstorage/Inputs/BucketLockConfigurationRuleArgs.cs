// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage.Inputs
{

    public sealed class BucketLockConfigurationRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default retention for the lock.
        /// </summary>
        [Input("defaultRetention", required: true)]
        public Input<Inputs.BucketLockConfigurationRuleDefaultRetentionArgs> DefaultRetention { get; set; } = null!;

        public BucketLockConfigurationRuleArgs()
        {
        }
        public static new BucketLockConfigurationRuleArgs Empty => new BucketLockConfigurationRuleArgs();
    }
}