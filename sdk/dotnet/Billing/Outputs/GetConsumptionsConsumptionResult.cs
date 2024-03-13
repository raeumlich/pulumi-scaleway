// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Billing.Outputs
{

    [OutputType]
    public sealed class GetConsumptionsConsumptionResult
    {
        /// <summary>
        /// The consumed quantity.
        /// </summary>
        public readonly string BilledQuantity;
        /// <summary>
        /// The name of the consumption category.
        /// </summary>
        public readonly string CategoryName;
        /// <summary>
        /// The product name.
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// `project_id`) The ID of the project the consumption list is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The unique identifier of the product.
        /// </summary>
        public readonly string Sku;
        /// <summary>
        /// The unit of consumed quantity.
        /// </summary>
        public readonly string Unit;
        /// <summary>
        /// The monetary value of the consumption.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetConsumptionsConsumptionResult(
            string billedQuantity,

            string categoryName,

            string productName,

            string projectId,

            string sku,

            string unit,

            string value)
        {
            BilledQuantity = billedQuantity;
            CategoryName = categoryName;
            ProductName = productName;
            ProjectId = projectId;
            Sku = sku;
            Unit = unit;
            Value = value;
        }
    }
}
