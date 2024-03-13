// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer
{
    public static class GetFrontend
    {
        /// <summary>
        /// Get information about Scaleway Load-Balancer Frontends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
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
        ///     var ip01 = new Scaleway.Loadbalancer.IP("ip01");
        /// 
        ///     var lb01 = new Scaleway.Loadbalancer.LoadBalancer("lb01", new()
        ///     {
        ///         IpId = ip01.Id,
        ///         Type = "lb-s",
        ///     });
        /// 
        ///     var bkd01 = new Scaleway.Loadbalancer.Backend("bkd01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         ForwardProtocol = "tcp",
        ///         ForwardPort = 80,
        ///         ProxyProtocol = "none",
        ///     });
        /// 
        ///     var frt01 = new Scaleway.Loadbalancer.Frontend("frt01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         BackendId = bkd01.Id,
        ///         InboundPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancer.GetFrontend.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancer.GetFrontend.Invoke(new()
        ///     {
        ///         Name = frt01.Name,
        ///         LbId = lb01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFrontendResult> InvokeAsync(GetFrontendArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFrontendResult>("scaleway:loadbalancer/getFrontend:getFrontend", args ?? new GetFrontendArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load-Balancer Frontends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
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
        ///     var ip01 = new Scaleway.Loadbalancer.IP("ip01");
        /// 
        ///     var lb01 = new Scaleway.Loadbalancer.LoadBalancer("lb01", new()
        ///     {
        ///         IpId = ip01.Id,
        ///         Type = "lb-s",
        ///     });
        /// 
        ///     var bkd01 = new Scaleway.Loadbalancer.Backend("bkd01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         ForwardProtocol = "tcp",
        ///         ForwardPort = 80,
        ///         ProxyProtocol = "none",
        ///     });
        /// 
        ///     var frt01 = new Scaleway.Loadbalancer.Frontend("frt01", new()
        ///     {
        ///         LbId = lb01.Id,
        ///         BackendId = bkd01.Id,
        ///         InboundPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancer.GetFrontend.Invoke(new()
        ///     {
        ///         FrontendId = frt01.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancer.GetFrontend.Invoke(new()
        ///     {
        ///         Name = frt01.Name,
        ///         LbId = lb01.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFrontendResult> Invoke(GetFrontendInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFrontendResult>("scaleway:loadbalancer/getFrontend:getFrontend", args ?? new GetFrontendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFrontendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend id.
        /// - Only one of `name` and `frontend_id` should be specified.
        /// </summary>
        [Input("frontendId")]
        public string? FrontendId { get; set; }

        /// <summary>
        /// The load-balancer ID this frontend is attached to.
        /// </summary>
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The name of the frontend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetFrontendArgs()
        {
        }
        public static new GetFrontendArgs Empty => new GetFrontendArgs();
    }

    public sealed class GetFrontendInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The frontend id.
        /// - Only one of `name` and `frontend_id` should be specified.
        /// </summary>
        [Input("frontendId")]
        public Input<string>? FrontendId { get; set; }

        /// <summary>
        /// The load-balancer ID this frontend is attached to.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the frontend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetFrontendInvokeArgs()
        {
        }
        public static new GetFrontendInvokeArgs Empty => new GetFrontendInvokeArgs();
    }


    [OutputType]
    public sealed class GetFrontendResult
    {
        public readonly ImmutableArray<Outputs.GetFrontendAclResult> Acls;
        public readonly string BackendId;
        public readonly string CertificateId;
        public readonly ImmutableArray<string> CertificateIds;
        public readonly bool EnableHttp3;
        public readonly bool ExternalAcls;
        public readonly string? FrontendId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InboundPort;
        public readonly string? LbId;
        public readonly string? Name;
        public readonly string TimeoutClient;

        [OutputConstructor]
        private GetFrontendResult(
            ImmutableArray<Outputs.GetFrontendAclResult> acls,

            string backendId,

            string certificateId,

            ImmutableArray<string> certificateIds,

            bool enableHttp3,

            bool externalAcls,

            string? frontendId,

            string id,

            int inboundPort,

            string? lbId,

            string? name,

            string timeoutClient)
        {
            Acls = acls;
            BackendId = backendId;
            CertificateId = certificateId;
            CertificateIds = certificateIds;
            EnableHttp3 = enableHttp3;
            ExternalAcls = externalAcls;
            FrontendId = frontendId;
            Id = id;
            InboundPort = inboundPort;
            LbId = lbId;
            Name = name;
            TimeoutClient = timeoutClient;
        }
    }
}
