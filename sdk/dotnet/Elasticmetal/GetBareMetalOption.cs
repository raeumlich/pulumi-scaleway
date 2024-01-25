// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal
{
    public static class GetBareMetalOption
    {
        /// <summary>
        /// Gets information about a baremetal option.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.Elasticmetal.GetBareMetalOption.Invoke(new()
        ///     {
        ///         Name = "Remote Access",
        ///     });
        /// 
        ///     var byId = Scaleway.Elasticmetal.GetBareMetalOption.Invoke(new()
        ///     {
        ///         OptionId = "931df052-d713-4674-8b58-96a63244c8e2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBareMetalOptionResult> InvokeAsync(GetBareMetalOptionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBareMetalOptionResult>("scaleway:elasticmetal/getBareMetalOption:getBareMetalOption", args ?? new GetBareMetalOptionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a baremetal option.
        /// For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
        ///     var byName = Scaleway.Elasticmetal.GetBareMetalOption.Invoke(new()
        ///     {
        ///         Name = "Remote Access",
        ///     });
        /// 
        ///     var byId = Scaleway.Elasticmetal.GetBareMetalOption.Invoke(new()
        ///     {
        ///         OptionId = "931df052-d713-4674-8b58-96a63244c8e2",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBareMetalOptionResult> Invoke(GetBareMetalOptionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBareMetalOptionResult>("scaleway:elasticmetal/getBareMetalOption:getBareMetalOption", args ?? new GetBareMetalOptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBareMetalOptionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The option name. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The option id. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("optionId")]
        public string? OptionId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the option exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetBareMetalOptionArgs()
        {
        }
        public static new GetBareMetalOptionArgs Empty => new GetBareMetalOptionArgs();
    }

    public sealed class GetBareMetalOptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The option name. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The option id. Only one of `name` and `option_id` should be specified.
        /// </summary>
        [Input("optionId")]
        public Input<string>? OptionId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the option exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBareMetalOptionInvokeArgs()
        {
        }
        public static new GetBareMetalOptionInvokeArgs Empty => new GetBareMetalOptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetBareMetalOptionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Is false if the option could not be added or removed.
        /// </summary>
        public readonly bool Manageable;
        /// <summary>
        /// The name of the option.
        /// </summary>
        public readonly string? Name;
        public readonly string? OptionId;
        public readonly string Zone;

        [OutputConstructor]
        private GetBareMetalOptionResult(
            string id,

            bool manageable,

            string? name,

            string? optionId,

            string zone)
        {
            Id = id;
            Manageable = manageable;
            Name = name;
            OptionId = optionId;
            Zone = zone;
        }
    }
}
