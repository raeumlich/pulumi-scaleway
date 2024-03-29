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
    public sealed class BucketACLAccessControlPolicy
    {
        public readonly ImmutableArray<Outputs.BucketACLAccessControlPolicyGrant> Grants;
        /// <summary>
        /// Configuration block of the bucket project owner's display organization ID.
        /// </summary>
        public readonly Outputs.BucketACLAccessControlPolicyOwner Owner;

        [OutputConstructor]
        private BucketACLAccessControlPolicy(
            ImmutableArray<Outputs.BucketACLAccessControlPolicyGrant> grants,

            Outputs.BucketACLAccessControlPolicyOwner owner)
        {
            Grants = grants;
            Owner = owner;
        }
    }
}
