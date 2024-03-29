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
    public sealed class ACLAction
    {
        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public readonly ImmutableArray<Outputs.ACLActionRedirect> Redirects;
        /// <summary>
        /// The redirect type. Possible values are: `location` or `scheme`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ACLAction(
            ImmutableArray<Outputs.ACLActionRedirect> redirects,

            string type)
        {
            Redirects = redirects;
            Type = type;
        }
    }
}
