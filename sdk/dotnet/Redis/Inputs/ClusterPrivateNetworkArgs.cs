// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Redis.Inputs
{

    public sealed class ClusterPrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The UUID of the Private Network resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("serviceIps")]
        private InputList<string>? _serviceIps;

        /// <summary>
        /// Endpoint IPv4 addresses in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation). You must provide at least one IP per node.
        /// Keep in mind that in Cluster mode you cannot edit your Private Network after its creation so if you want to be able to
        /// scale your Cluster horizontally (adding nodes) later, you should provide more IPs than nodes.
        /// If not set, the IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// 
        /// &gt; The `private_network` conflicts with `acl`. Only one should be specified.
        /// 
        /// &gt; **Important:** The way to use private networks differs whether you are using Redis in Standalone or Cluster mode.
        /// 
        /// - Standalone mode (`cluster_size` = 1) : you can attach as many Private Networks as you want (each must be a separate
        /// block). If you detach your only private network, your cluster won't be reachable until you define a new Private or
        /// Public Network. You can modify your `private_network` and its specs, you can have both a Private and Public Network side
        /// by side.
        /// 
        /// - Cluster mode (`cluster_size` &gt; 2) : you can define a single Private Network as you create your Cluster, you won't be
        /// able to edit or detach it afterward, unless you create another Cluster. This also means that, if you are using a static
        /// configuration (`service_ips`), you won't be able to scale your Cluster horizontally (add more nodes) since it would
        /// require updating the private network to add IPs.
        /// Your `service_ips` must be listed as follows:
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public InputList<string> ServiceIps
        {
            get => _serviceIps ?? (_serviceIps = new InputList<string>());
            set => _serviceIps = value;
        }

        /// <summary>
        /// `zone`) The zone in which the
        /// Redis Cluster should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ClusterPrivateNetworkArgs()
        {
        }
        public static new ClusterPrivateNetworkArgs Empty => new ClusterPrivateNetworkArgs();
    }
}
