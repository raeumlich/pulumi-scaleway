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
    public sealed class GetFrontendAclActionResult
    {
        public readonly ImmutableArray<Outputs.GetFrontendAclActionRedirectResult> Redirects;
        public readonly string Type;

        [OutputConstructor]
        private GetFrontendAclActionResult(
            ImmutableArray<Outputs.GetFrontendAclActionRedirectResult> redirects,

            string type)
        {
            Redirects = redirects;
            Type = type;
        }
    }
}
