// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Loadbalancer
{
    public static class GetBackend
    {
        /// <summary>
        /// Get information about Scaleway Load-Balancer Backends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
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
        ///     var mainIP = new Scaleway.Loadbalancer.IP("mainIP");
        /// 
        ///     var mainLoadBalancer = new Scaleway.Loadbalancer.LoadBalancer("mainLoadBalancer", new()
        ///     {
        ///         IpId = mainIP.Id,
        ///         Type = "LB-S",
        ///     });
        /// 
        ///     var mainBackend = new Scaleway.Loadbalancer.Backend("mainBackend", new()
        ///     {
        ///         LbId = mainLoadBalancer.Id,
        ///         ForwardProtocol = "http",
        ///         ForwardPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancer.GetBackend.Invoke(new()
        ///     {
        ///         BackendId = mainBackend.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancer.GetBackend.Invoke(new()
        ///     {
        ///         Name = mainBackend.Name,
        ///         LbId = mainLoadBalancer.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBackendResult> InvokeAsync(GetBackendArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackendResult>("scaleway:loadbalancer/getBackend:getBackend", args ?? new GetBackendArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about Scaleway Load-Balancer Backends.
        /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-backends).
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
        ///     var mainIP = new Scaleway.Loadbalancer.IP("mainIP");
        /// 
        ///     var mainLoadBalancer = new Scaleway.Loadbalancer.LoadBalancer("mainLoadBalancer", new()
        ///     {
        ///         IpId = mainIP.Id,
        ///         Type = "LB-S",
        ///     });
        /// 
        ///     var mainBackend = new Scaleway.Loadbalancer.Backend("mainBackend", new()
        ///     {
        ///         LbId = mainLoadBalancer.Id,
        ///         ForwardProtocol = "http",
        ///         ForwardPort = 80,
        ///     });
        /// 
        ///     var byID = Scaleway.Loadbalancer.GetBackend.Invoke(new()
        ///     {
        ///         BackendId = mainBackend.Id,
        ///     });
        /// 
        ///     var byName = Scaleway.Loadbalancer.GetBackend.Invoke(new()
        ///     {
        ///         Name = mainBackend.Name,
        ///         LbId = mainLoadBalancer.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBackendResult> Invoke(GetBackendInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendResult>("scaleway:loadbalancer/getBackend:getBackend", args ?? new GetBackendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend id.
        /// - Only one of `name` and `backend_id` should be specified.
        /// </summary>
        [Input("backendId")]
        public string? BackendId { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// </summary>
        [Input("lbId")]
        public string? LbId { get; set; }

        /// <summary>
        /// The name of the backend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetBackendArgs()
        {
        }
        public static new GetBackendArgs Empty => new GetBackendArgs();
    }

    public sealed class GetBackendInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backend id.
        /// - Only one of `name` and `backend_id` should be specified.
        /// </summary>
        [Input("backendId")]
        public Input<string>? BackendId { get; set; }

        /// <summary>
        /// The load-balancer ID this backend is attached to.
        /// </summary>
        [Input("lbId")]
        public Input<string>? LbId { get; set; }

        /// <summary>
        /// The name of the backend.
        /// - When using the `name` you should specify the `lb-id`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetBackendInvokeArgs()
        {
        }
        public static new GetBackendInvokeArgs Empty => new GetBackendInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackendResult
    {
        public readonly string? BackendId;
        public readonly string FailoverHost;
        public readonly int ForwardPort;
        public readonly string ForwardPortAlgorithm;
        public readonly string ForwardProtocol;
        public readonly string HealthCheckDelay;
        public readonly ImmutableArray<Outputs.GetBackendHealthCheckHttpResult> HealthCheckHttp;
        public readonly ImmutableArray<Outputs.GetBackendHealthCheckHttpResult> HealthCheckHttps;
        public readonly int HealthCheckMaxRetries;
        public readonly int HealthCheckPort;
        public readonly bool HealthCheckSendProxy;
        public readonly ImmutableArray<Outputs.GetBackendHealthCheckTcpResult> HealthCheckTcps;
        public readonly string HealthCheckTimeout;
        public readonly string HealthCheckTransientDelay;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IgnoreSslServerVerify;
        public readonly string? LbId;
        public readonly int MaxConnections;
        public readonly int MaxRetries;
        public readonly string? Name;
        public readonly string OnMarkedDownAction;
        public readonly string ProxyProtocol;
        public readonly int RedispatchAttemptCount;
        public readonly bool SendProxyV2;
        public readonly ImmutableArray<string> ServerIps;
        public readonly bool SslBridging;
        public readonly string StickySessions;
        public readonly string StickySessionsCookieName;
        public readonly string TimeoutConnect;
        public readonly string TimeoutQueue;
        public readonly string TimeoutServer;
        public readonly string TimeoutTunnel;

        [OutputConstructor]
        private GetBackendResult(
            string? backendId,

            string failoverHost,

            int forwardPort,

            string forwardPortAlgorithm,

            string forwardProtocol,

            string healthCheckDelay,

            ImmutableArray<Outputs.GetBackendHealthCheckHttpResult> healthCheckHttp,

            ImmutableArray<Outputs.GetBackendHealthCheckHttpResult> healthCheckHttps,

            int healthCheckMaxRetries,

            int healthCheckPort,

            bool healthCheckSendProxy,

            ImmutableArray<Outputs.GetBackendHealthCheckTcpResult> healthCheckTcps,

            string healthCheckTimeout,

            string healthCheckTransientDelay,

            string id,

            bool ignoreSslServerVerify,

            string? lbId,

            int maxConnections,

            int maxRetries,

            string? name,

            string onMarkedDownAction,

            string proxyProtocol,

            int redispatchAttemptCount,

            bool sendProxyV2,

            ImmutableArray<string> serverIps,

            bool sslBridging,

            string stickySessions,

            string stickySessionsCookieName,

            string timeoutConnect,

            string timeoutQueue,

            string timeoutServer,

            string timeoutTunnel)
        {
            BackendId = backendId;
            FailoverHost = failoverHost;
            ForwardPort = forwardPort;
            ForwardPortAlgorithm = forwardPortAlgorithm;
            ForwardProtocol = forwardProtocol;
            HealthCheckDelay = healthCheckDelay;
            HealthCheckHttp = healthCheckHttp;
            HealthCheckHttps = healthCheckHttps;
            HealthCheckMaxRetries = healthCheckMaxRetries;
            HealthCheckPort = healthCheckPort;
            HealthCheckSendProxy = healthCheckSendProxy;
            HealthCheckTcps = healthCheckTcps;
            HealthCheckTimeout = healthCheckTimeout;
            HealthCheckTransientDelay = healthCheckTransientDelay;
            Id = id;
            IgnoreSslServerVerify = ignoreSslServerVerify;
            LbId = lbId;
            MaxConnections = maxConnections;
            MaxRetries = maxRetries;
            Name = name;
            OnMarkedDownAction = onMarkedDownAction;
            ProxyProtocol = proxyProtocol;
            RedispatchAttemptCount = redispatchAttemptCount;
            SendProxyV2 = sendProxyV2;
            ServerIps = serverIps;
            SslBridging = sslBridging;
            StickySessions = stickySessions;
            StickySessionsCookieName = stickySessionsCookieName;
            TimeoutConnect = timeoutConnect;
            TimeoutQueue = timeoutQueue;
            TimeoutServer = timeoutServer;
            TimeoutTunnel = timeoutTunnel;
        }
    }
}
