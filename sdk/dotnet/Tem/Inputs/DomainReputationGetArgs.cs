// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Tem.Inputs
{

    public sealed class DomainReputationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The previously-calculated domain's reputation score.
        /// </summary>
        [Input("previousScore")]
        public Input<int>? PreviousScore { get; set; }

        /// <summary>
        /// The time and date the previous reputation score was calculated.
        /// </summary>
        [Input("previousScoredAt")]
        public Input<string>? PreviousScoredAt { get; set; }

        /// <summary>
        /// A range from 0 to 100 that determines your domain's reputation score.
        /// </summary>
        [Input("score")]
        public Input<int>? Score { get; set; }

        /// <summary>
        /// The time and date the score was calculated.
        /// </summary>
        [Input("scoredAt")]
        public Input<string>? ScoredAt { get; set; }

        /// <summary>
        /// The status of the domain's reputation.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DomainReputationGetArgs()
        {
        }
        public static new DomainReputationGetArgs Empty => new DomainReputationGetArgs();
    }
}
