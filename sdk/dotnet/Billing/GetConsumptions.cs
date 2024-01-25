// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Billing
{
    public static class GetConsumptions
    {
        /// <summary>
        /// Gets information about your Consumptions.
        /// </summary>
        public static Task<GetConsumptionsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConsumptionsResult>("scaleway:billing/getConsumptions:getConsumptions", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Gets information about your Consumptions.
        /// </summary>
        public static Output<GetConsumptionsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConsumptionsResult>("scaleway:billing/getConsumptions:getConsumptions", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetConsumptionsResult
    {
        /// <summary>
        /// List of found consumptions
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConsumptionsConsumptionResult> Consumptions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OrganizationId;
        /// <summary>
        /// The last consumption update date.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetConsumptionsResult(
            ImmutableArray<Outputs.GetConsumptionsConsumptionResult> consumptions,

            string id,

            string organizationId,

            string updatedAt)
        {
            Consumptions = consumptions;
            Id = id;
            OrganizationId = organizationId;
            UpdatedAt = updatedAt;
        }
    }
}