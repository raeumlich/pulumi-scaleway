// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Elasticmetal
{
    public static class GetFlexibleIPs
    {
        /// <summary>
        /// Gets information about multiple Flexible IPs.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fipsByTags = Scaleway.Elasticmetal.GetFlexibleIPs.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "a tag",
        ///         },
        ///     });
        /// 
        ///     var myOffer = Scaleway.Elasticmetal.GetBareMetalOffer.Invoke(new()
        ///     {
        ///         Name = "EM-B112X-SSD",
        ///     });
        /// 
        ///     var @base = new Scaleway.Elasticmetal.BareMetalServer("base", new()
        ///     {
        ///         Offer = myOffer.Apply(getBareMetalOfferResult =&gt; getBareMetalOfferResult.OfferId),
        ///         InstallConfigAfterward = true,
        ///     });
        /// 
        ///     var first = new Scaleway.Elasticmetal.FlexibleIP("first", new()
        ///     {
        ///         ServerId = @base.Id,
        ///         Tags = new[]
        ///         {
        ///             "foo",
        ///             "first",
        ///         },
        ///     });
        /// 
        ///     var second = new Scaleway.Elasticmetal.FlexibleIP("second", new()
        ///     {
        ///         ServerId = @base.Id,
        ///         Tags = new[]
        ///         {
        ///             "foo",
        ///             "second",
        ///         },
        ///     });
        /// 
        ///     var fipsByServerId = Scaleway.Elasticmetal.GetFlexibleIPs.Invoke(new()
        ///     {
        ///         ServerIds = new[]
        ///         {
        ///             @base.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFlexibleIPsResult> InvokeAsync(GetFlexibleIPsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlexibleIPsResult>("scaleway:elasticmetal/getFlexibleIPs:getFlexibleIPs", args ?? new GetFlexibleIPsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about multiple Flexible IPs.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fipsByTags = Scaleway.Elasticmetal.GetFlexibleIPs.Invoke(new()
        ///     {
        ///         Tags = new[]
        ///         {
        ///             "a tag",
        ///         },
        ///     });
        /// 
        ///     var myOffer = Scaleway.Elasticmetal.GetBareMetalOffer.Invoke(new()
        ///     {
        ///         Name = "EM-B112X-SSD",
        ///     });
        /// 
        ///     var @base = new Scaleway.Elasticmetal.BareMetalServer("base", new()
        ///     {
        ///         Offer = myOffer.Apply(getBareMetalOfferResult =&gt; getBareMetalOfferResult.OfferId),
        ///         InstallConfigAfterward = true,
        ///     });
        /// 
        ///     var first = new Scaleway.Elasticmetal.FlexibleIP("first", new()
        ///     {
        ///         ServerId = @base.Id,
        ///         Tags = new[]
        ///         {
        ///             "foo",
        ///             "first",
        ///         },
        ///     });
        /// 
        ///     var second = new Scaleway.Elasticmetal.FlexibleIP("second", new()
        ///     {
        ///         ServerId = @base.Id,
        ///         Tags = new[]
        ///         {
        ///             "foo",
        ///             "second",
        ///         },
        ///     });
        /// 
        ///     var fipsByServerId = Scaleway.Elasticmetal.GetFlexibleIPs.Invoke(new()
        ///     {
        ///         ServerIds = new[]
        ///         {
        ///             @base.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFlexibleIPsResult> Invoke(GetFlexibleIPsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlexibleIPsResult>("scaleway:elasticmetal/getFlexibleIPs:getFlexibleIPs", args ?? new GetFlexibleIPsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlexibleIPsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Defaults to provider `project_id`) The ID of the project the IP is in.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("serverIds")]
        private List<string>? _serverIds;

        /// <summary>
        /// List of server IDs used as filter. IPs with these exact server IDs are listed.
        /// </summary>
        public List<string> ServerIds
        {
            get => _serverIds ?? (_serverIds = new List<string>());
            set => _serverIds = value;
        }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// List of tags used as filter. IPs with these exact tags are listed.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which IPs exist.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetFlexibleIPsArgs()
        {
        }
        public static new GetFlexibleIPsArgs Empty => new GetFlexibleIPsArgs();
    }

    public sealed class GetFlexibleIPsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Defaults to provider `project_id`) The ID of the project the IP is in.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("serverIds")]
        private InputList<string>? _serverIds;

        /// <summary>
        /// List of server IDs used as filter. IPs with these exact server IDs are listed.
        /// </summary>
        public InputList<string> ServerIds
        {
            get => _serverIds ?? (_serverIds = new InputList<string>());
            set => _serverIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags used as filter. IPs with these exact tags are listed.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which IPs exist.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetFlexibleIPsInvokeArgs()
        {
        }
        public static new GetFlexibleIPsInvokeArgs Empty => new GetFlexibleIPsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlexibleIPsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of found flexible IPS
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFlexibleIPsIpResult> Ips;
        /// <summary>
        /// (Defaults to provider `organization_id`) The ID of the organization the IP is in.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// (Defaults to provider `project_id`) The ID of the project the IP is in.
        /// </summary>
        public readonly string ProjectId;
        public readonly ImmutableArray<string> ServerIds;
        /// <summary>
        /// The list of tags which are attached to the flexible IP.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// (Defaults to provider `zone`) The zone in which the MAC address exist.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetFlexibleIPsResult(
            string id,

            ImmutableArray<Outputs.GetFlexibleIPsIpResult> ips,

            string organizationId,

            string projectId,

            ImmutableArray<string> serverIds,

            ImmutableArray<string> tags,

            string zone)
        {
            Id = id;
            Ips = ips;
            OrganizationId = organizationId;
            ProjectId = projectId;
            ServerIds = serverIds;
            Tags = tags;
            Zone = zone;
        }
    }
}
