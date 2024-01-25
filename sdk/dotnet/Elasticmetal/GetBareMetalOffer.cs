// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal
{
    public static class GetBareMetalOffer
    {
        /// <summary>
        /// Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myOffer = Scaleway.Elasticmetal.GetBareMetalOffer.Invoke(new()
        ///     {
        ///         OfferId = "25dcf38b-c90c-4b18-97a2-6956e9d1e113",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBareMetalOfferResult> InvokeAsync(GetBareMetalOfferArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBareMetalOfferResult>("scaleway:elasticmetal/getBareMetalOffer:getBareMetalOffer", args ?? new GetBareMetalOfferArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myOffer = Scaleway.Elasticmetal.GetBareMetalOffer.Invoke(new()
        ///     {
        ///         OfferId = "25dcf38b-c90c-4b18-97a2-6956e9d1e113",
        ///         Zone = "fr-par-2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBareMetalOfferResult> Invoke(GetBareMetalOfferInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBareMetalOfferResult>("scaleway:elasticmetal/getBareMetalOffer:getBareMetalOffer", args ?? new GetBareMetalOfferInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBareMetalOfferArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeDisabled")]
        public bool? IncludeDisabled { get; set; }

        /// <summary>
        /// The offer name. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The offer id. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("offerId")]
        public string? OfferId { get; set; }

        /// <summary>
        /// Period of subscription the desired offer. Should be `hourly` or `monthly`.
        /// </summary>
        [Input("subscriptionPeriod")]
        public string? SubscriptionPeriod { get; set; }

        /// <summary>
        /// `zone`) The zone in which the offer should be created.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetBareMetalOfferArgs()
        {
        }
        public static new GetBareMetalOfferArgs Empty => new GetBareMetalOfferArgs();
    }

    public sealed class GetBareMetalOfferInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeDisabled")]
        public Input<bool>? IncludeDisabled { get; set; }

        /// <summary>
        /// The offer name. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The offer id. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        [Input("offerId")]
        public Input<string>? OfferId { get; set; }

        /// <summary>
        /// Period of subscription the desired offer. Should be `hourly` or `monthly`.
        /// </summary>
        [Input("subscriptionPeriod")]
        public Input<string>? SubscriptionPeriod { get; set; }

        /// <summary>
        /// `zone`) The zone in which the offer should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBareMetalOfferInvokeArgs()
        {
        }
        public static new GetBareMetalOfferInvokeArgs Empty => new GetBareMetalOfferInvokeArgs();
    }


    [OutputType]
    public sealed class GetBareMetalOfferResult
    {
        /// <summary>
        /// Available Bandwidth with the offer.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// Commercial range of the offer.
        /// </summary>
        public readonly string CommercialRange;
        /// <summary>
        /// A list of cpu specifications. (Structure is documented below.)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBareMetalOfferCpusResult> Cpus;
        /// <summary>
        /// A list of disk specifications. (Structure is documented below.)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBareMetalOfferDiskResult> Disks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeDisabled;
        /// <summary>
        /// A list of memory specifications. (Structure is documented below.)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBareMetalOfferMemoryResult> Memories;
        /// <summary>
        /// Name of the CPU.
        /// </summary>
        public readonly string? Name;
        public readonly string? OfferId;
        /// <summary>
        /// Stock status for this offer. Possible values are: `empty`, `low` or `available`.
        /// </summary>
        public readonly string Stock;
        public readonly string? SubscriptionPeriod;
        public readonly string Zone;

        [OutputConstructor]
        private GetBareMetalOfferResult(
            int bandwidth,

            string commercialRange,

            ImmutableArray<Outputs.GetBareMetalOfferCpusResult> cpus,

            ImmutableArray<Outputs.GetBareMetalOfferDiskResult> disks,

            string id,

            bool? includeDisabled,

            ImmutableArray<Outputs.GetBareMetalOfferMemoryResult> memories,

            string? name,

            string? offerId,

            string stock,

            string? subscriptionPeriod,

            string zone)
        {
            Bandwidth = bandwidth;
            CommercialRange = commercialRange;
            Cpus = cpus;
            Disks = disks;
            Id = id;
            IncludeDisabled = includeDisabled;
            Memories = memories;
            Name = name;
            OfferId = offerId;
            Stock = stock;
            SubscriptionPeriod = subscriptionPeriod;
            Zone = zone;
        }
    }
}
