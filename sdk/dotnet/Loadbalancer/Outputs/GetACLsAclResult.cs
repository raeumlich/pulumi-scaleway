// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer.Outputs
{

    [OutputType]
    public sealed class GetACLsAclResult
    {
        /// <summary>
        /// The action that has been undertaken when an ACL filter had matched.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetACLsAclActionResult> Actions;
        /// <summary>
        /// The date at which the ACL was created (RFC 3339 format).
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the ACL resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
        /// &gt; **Important:** LB Frontends' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
        /// </summary>
        public readonly string FrontendId;
        /// <summary>
        /// The associated ACL ID.
        /// &gt; **Important:** LB ACLs' IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The order between the ACLs.
        /// </summary>
        public readonly int Index;
        /// <summary>
        /// The ACL match rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetACLsAclMatchResult> Matches;
        /// <summary>
        /// The ACL name used as filter. ACLs with a name like it are listed.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The date at which the ACL was last updated (RFC 3339 format).
        /// </summary>
        public readonly string UpdateAt;

        [OutputConstructor]
        private GetACLsAclResult(
            ImmutableArray<Outputs.GetACLsAclActionResult> actions,

            string createdAt,

            string description,

            string frontendId,

            string id,

            int index,

            ImmutableArray<Outputs.GetACLsAclMatchResult> matches,

            string name,

            string updateAt)
        {
            Actions = actions;
            CreatedAt = createdAt;
            Description = description;
            FrontendId = frontendId;
            Id = id;
            Index = index;
            Matches = matches;
            Name = name;
            UpdateAt = updateAt;
        }
    }
}
