// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Documentdb
{
    public static class GetLoadBalancerEndpoint
    {
        /// <summary>
        /// Gets information about an DocumentDB load balancer endpoint.
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
        ///     var myEndpoint = Scaleway.Documentdb.GetLoadBalancerEndpoint.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLoadBalancerEndpointResult> InvokeAsync(GetLoadBalancerEndpointArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerEndpointResult>("scaleway:documentdb/getLoadBalancerEndpoint:getLoadBalancerEndpoint", args ?? new GetLoadBalancerEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an DocumentDB load balancer endpoint.
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
        ///     var myEndpoint = Scaleway.Documentdb.GetLoadBalancerEndpoint.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLoadBalancerEndpointResult> Invoke(GetLoadBalancerEndpointInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoadBalancerEndpointResult>("scaleway:documentdb/getLoadBalancerEndpoint:getLoadBalancerEndpoint", args ?? new GetLoadBalancerEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancerEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DocumentDB Instance on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        /// <summary>
        /// The ID of the project the DocumentDB endpoint is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the DocumentDB endpoint exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetLoadBalancerEndpointArgs()
        {
        }
        public static new GetLoadBalancerEndpointArgs Empty => new GetLoadBalancerEndpointArgs();
    }

    public sealed class GetLoadBalancerEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DocumentDB Instance on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instance_name` and `instance_id` should be specified.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The ID of the project the DocumentDB endpoint is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the DocumentDB endpoint exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetLoadBalancerEndpointInvokeArgs()
        {
        }
        public static new GetLoadBalancerEndpointInvokeArgs Empty => new GetLoadBalancerEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoadBalancerEndpointResult
    {
        /// <summary>
        /// The hostname of your endpoint.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InstanceName;
        /// <summary>
        /// The IP of your load balancer service.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The name of your load balancer service.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port of your load balancer service.
        /// </summary>
        public readonly int Port;
        public readonly string ProjectId;
        public readonly string Region;

        [OutputConstructor]
        private GetLoadBalancerEndpointResult(
            string hostname,

            string id,

            string instanceId,

            string instanceName,

            string ip,

            string name,

            int port,

            string projectId,

            string region)
        {
            Hostname = hostname;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Ip = ip;
            Name = name;
            Port = port;
            ProjectId = projectId;
            Region = region;
        }
    }
}
