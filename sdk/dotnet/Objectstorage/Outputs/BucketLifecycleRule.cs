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
    public sealed class BucketLifecycleRule
    {
        /// <summary>
        /// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
        /// 
        /// * &gt; **Important:** It's not recommended using `prefix` for `AbortIncompleteMultipartUpload` as any incomplete multipart upload will be billed
        /// </summary>
        public readonly int? AbortIncompleteMultipartUploadDays;
        /// <summary>
        /// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Specifies a period in the object's expire (documented below).
        /// </summary>
        public readonly Outputs.BucketLifecycleRuleExpiration? Expiration;
        /// <summary>
        /// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// Specifies object tags key and value.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Specifies a period in the object's transitions (documented below).
        /// 
        /// At least one of `abort_incomplete_multipart_upload_days`, `expiration`, `transition` must be specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketLifecycleRuleTransition> Transitions;

        [OutputConstructor]
        private BucketLifecycleRule(
            int? abortIncompleteMultipartUploadDays,

            bool enabled,

            Outputs.BucketLifecycleRuleExpiration? expiration,

            string? id,

            string? prefix,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.BucketLifecycleRuleTransition> transitions)
        {
            AbortIncompleteMultipartUploadDays = abortIncompleteMultipartUploadDays;
            Enabled = enabled;
            Expiration = expiration;
            Id = id;
            Prefix = prefix;
            Tags = tags;
            Transitions = transitions;
        }
    }
}
