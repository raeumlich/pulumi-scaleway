// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Webhosting
{
    public static class GetWebHosting
    {
        /// <summary>
        /// Gets information about a webhosting.
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
        ///     var byDomain = Scaleway.Webhosting.GetWebHosting.Invoke(new()
        ///     {
        ///         Domain = "foobar.com",
        ///     });
        /// 
        ///     var byId = Scaleway.Webhosting.GetWebHosting.Invoke(new()
        ///     {
        ///         WebhostingId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetWebHostingResult> InvokeAsync(GetWebHostingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebHostingResult>("scaleway:webhosting/getWebHosting:getWebHosting", args ?? new GetWebHostingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a webhosting.
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
        ///     var byDomain = Scaleway.Webhosting.GetWebHosting.Invoke(new()
        ///     {
        ///         Domain = "foobar.com",
        ///     });
        /// 
        ///     var byId = Scaleway.Webhosting.GetWebHosting.Invoke(new()
        ///     {
        ///         WebhostingId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetWebHostingResult> Invoke(GetWebHostingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebHostingResult>("scaleway:webhosting/getWebHosting:getWebHosting", args ?? new GetWebHostingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebHostingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// The ID of the organization the hosting is associated with.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the hosting is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The hosting id. Only one of `domain` and `webhosting_id` should be specified.
        /// </summary>
        [Input("webhostingId")]
        public string? WebhostingId { get; set; }

        public GetWebHostingArgs()
        {
        }
        public static new GetWebHostingArgs Empty => new GetWebHostingArgs();
    }

    public sealed class GetWebHostingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of the organization the hosting is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the hosting is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The hosting id. Only one of `domain` and `webhosting_id` should be specified.
        /// </summary>
        [Input("webhostingId")]
        public Input<string>? WebhostingId { get; set; }

        public GetWebHostingInvokeArgs()
        {
        }
        public static new GetWebHostingInvokeArgs Empty => new GetWebHostingInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebHostingResult
    {
        public readonly ImmutableArray<Outputs.GetWebHostingCpanelUrlResult> CpanelUrls;
        public readonly string CreatedAt;
        public readonly string DnsStatus;
        public readonly string? Domain;
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OfferId;
        public readonly string OfferName;
        public readonly ImmutableArray<string> OptionIds;
        public readonly ImmutableArray<Outputs.GetWebHostingOptionResult> Options;
        public readonly string OrganizationId;
        public readonly string PlatformHostname;
        public readonly int PlatformNumber;
        public readonly string? ProjectId;
        public readonly string Region;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string Username;
        public readonly string? WebhostingId;

        [OutputConstructor]
        private GetWebHostingResult(
            ImmutableArray<Outputs.GetWebHostingCpanelUrlResult> cpanelUrls,

            string createdAt,

            string dnsStatus,

            string? domain,

            string email,

            string id,

            string offerId,

            string offerName,

            ImmutableArray<string> optionIds,

            ImmutableArray<Outputs.GetWebHostingOptionResult> options,

            string organizationId,

            string platformHostname,

            int platformNumber,

            string? projectId,

            string region,

            string status,

            ImmutableArray<string> tags,

            string updatedAt,

            string username,

            string? webhostingId)
        {
            CpanelUrls = cpanelUrls;
            CreatedAt = createdAt;
            DnsStatus = dnsStatus;
            Domain = domain;
            Email = email;
            Id = id;
            OfferId = offerId;
            OfferName = offerName;
            OptionIds = optionIds;
            Options = options;
            OrganizationId = organizationId;
            PlatformHostname = platformHostname;
            PlatformNumber = platformNumber;
            ProjectId = projectId;
            Region = region;
            Status = status;
            Tags = tags;
            UpdatedAt = updatedAt;
            Username = username;
            WebhostingId = webhostingId;
        }
    }
}
