// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer.Inputs
{

    public sealed class FrontendAclGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.FrontendAclActionGetArgs> Action { get; set; } = null!;

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.FrontendAclMatchGetArgs> Match { get; set; } = null!;

        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public FrontendAclGetArgs()
        {
        }
        public static new FrontendAclGetArgs Empty => new FrontendAclGetArgs();
    }
}
