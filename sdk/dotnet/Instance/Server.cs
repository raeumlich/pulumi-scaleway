// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance
{
    /// <summary>
    /// Creates and manages Scaleway Compute Instance servers. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#servers-8bf7d7).
    /// 
    /// Please check our [FAQ - Instances](https://www.scaleway.com/en/docs/faq/instances).
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var publicIp = new Scaleway.Instance.IP("publicIp");
    /// 
    ///     var web = new Scaleway.Instance.Server("web", new()
    ///     {
    ///         Type = "DEV1-S",
    ///         Image = "ubuntu_jammy",
    ///         IpId = publicIp.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// ### With additional volumes and tags
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var data = new Scaleway.Instance.Volume("data", new()
    ///     {
    ///         SizeInGb = 100,
    ///         Type = "b_ssd",
    ///     });
    /// 
    ///     var web = new Scaleway.Instance.Server("web", new()
    ///     {
    ///         Type = "DEV1-S",
    ///         Image = "ubuntu_jammy",
    ///         Tags = new[]
    ///         {
    ///             "hello",
    ///             "public",
    ///         },
    ///         RootVolume = new Scaleway.Instance.Inputs.ServerRootVolumeArgs
    ///         {
    ///             DeleteOnTermination = false,
    ///         },
    ///         AdditionalVolumeIds = new[]
    ///         {
    ///             data.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### With a reserved IP
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ip = new Scaleway.Instance.IP("ip");
    /// 
    ///     var web = new Scaleway.Instance.Server("web", new()
    ///     {
    ///         Type = "DEV1-S",
    ///         Image = "f974feac-abae-4365-b988-8ec7d1cec10d",
    ///         Tags = new[]
    ///         {
    ///             "hello",
    ///             "public",
    ///         },
    ///         IpId = ip.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// ### With security group
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var www = new Scaleway.Instance.SecurityGroup("www", new()
    ///     {
    ///         InboundDefaultPolicy = "drop",
    ///         OutboundDefaultPolicy = "accept",
    ///         InboundRules = new[]
    ///         {
    ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
    ///             {
    ///                 Action = "accept",
    ///                 Port = 22,
    ///                 Ip = "212.47.225.64",
    ///             },
    ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
    ///             {
    ///                 Action = "accept",
    ///                 Port = 80,
    ///             },
    ///             new Scaleway.Instance.Inputs.SecurityGroupInboundRuleArgs
    ///             {
    ///                 Action = "accept",
    ///                 Port = 443,
    ///             },
    ///         },
    ///         OutboundRules = new[]
    ///         {
    ///             new Scaleway.Instance.Inputs.SecurityGroupOutboundRuleArgs
    ///             {
    ///                 Action = "drop",
    ///                 IpRange = "10.20.0.0/24",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var web = new Scaleway.Instance.Server("web", new()
    ///     {
    ///         Type = "DEV1-S",
    ///         Image = "ubuntu_jammy",
    ///         SecurityGroupId = www.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// ### With user data and cloud-init
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var web = new Scaleway.Instance.Server("web", new()
    ///     {
    ///         Type = "DEV1-S",
    ///         Image = "ubuntu_jammy",
    ///         UserData = 
    ///         {
    ///             { "foo", "bar" },
    ///             { "cloud-init", File.ReadAllText($"{path.Module}/cloud-init.yml") },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### With private network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn01 = new Scaleway.Vpc.PrivateNetwork("pn01");
    /// 
    ///     var @base = new Scaleway.Instance.Server("base", new()
    ///     {
    ///         Image = "ubuntu_jammy",
    ///         Type = "DEV1-S",
    ///         PrivateNetworks = new[]
    ///         {
    ///             new Scaleway.Instance.Inputs.ServerPrivateNetworkArgs
    ///             {
    ///                 PnId = pn01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Root volume configuration
    /// ### Resized block volume with installed image
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var image = new Scaleway.Instance.Server("image", new()
    ///     {
    ///         Image = "ubuntu_jammy",
    ///         RootVolume = new Scaleway.Instance.Inputs.ServerRootVolumeArgs
    ///         {
    ///             SizeInGb = 100,
    ///             VolumeType = "b_ssd",
    ///         },
    ///         Type = "PRO2-XXS",
    ///     });
    /// 
    /// });
    /// ```
    /// ### From snapshot
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var snapshot = Scaleway.Instance.GetSnapshot.Invoke(new()
    ///     {
    ///         Name = "my_snapshot",
    ///     });
    /// 
    ///     var fromSnapshotVolume = new Scaleway.Instance.Volume("fromSnapshotVolume", new()
    ///     {
    ///         FromSnapshotId = snapshot.Apply(getSnapshotResult =&gt; getSnapshotResult.Id),
    ///         Type = "b_ssd",
    ///     });
    /// 
    ///     var fromSnapshotServer = new Scaleway.Instance.Server("fromSnapshotServer", new()
    ///     {
    ///         Type = "PRO2-XXS",
    ///         RootVolume = new Scaleway.Instance.Inputs.ServerRootVolumeArgs
    ///         {
    ///             VolumeId = fromSnapshotVolume.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Private Network
    /// 
    /// &gt; **Important:** Updates to `private_network` will recreate a new private network interface.
    /// 
    /// - `pn_id` - (Required) The private network ID where to connect.
    /// - `mac_address` The private NIC MAC address.
    /// - `status` The private NIC state.
    /// - `zone` - (Defaults to provider `zone`) The zone in which the server must be created.
    /// 
    /// &gt; **Important:**
    /// 
    /// - You can only attach an instance in the same zone as a private network.
    /// - Instance supports maximum 8 different private networks.
    /// 
    /// ## Import
    /// 
    /// Instance servers can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:instance/server:Server web fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:instance/server:Server")]
    public partial class Server : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
        /// attached to the server. Updates to this field will trigger a stop/start of the server.
        /// 
        /// &gt; **Important:** If this field contains local volumes, the `state` must be set to `stopped`, otherwise it will fail.
        /// 
        /// &gt; **Important:** If this field contains local volumes, you have to first detach them, in one apply, and then delete the volume in another apply.
        /// </summary>
        [Output("additionalVolumeIds")]
        public Output<ImmutableArray<string>> AdditionalVolumeIds { get; private set; } = null!;

        /// <summary>
        /// The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
        /// </summary>
        [Output("bootType")]
        public Output<string?> BootType { get; private set; } = null!;

        /// <summary>
        /// The ID of the bootscript to use  (set boot_type to `bootscript`).
        /// </summary>
        [Output("bootscriptId")]
        public Output<string> BootscriptId { get; private set; } = null!;

        /// <summary>
        /// The cloud init script associated with this server
        /// </summary>
        [Output("cloudInit")]
        public Output<string> CloudInit { get; private set; } = null!;

        /// <summary>
        /// If true a dynamic IP will be attached to the server.
        /// </summary>
        [Output("enableDynamicIp")]
        public Output<bool?> EnableDynamicIp { get; private set; } = null!;

        /// <summary>
        /// Determines if IPv6 is enabled for the server.
        /// </summary>
        [Output("enableIpv6")]
        public Output<bool?> EnableIpv6 { get; private set; } = null!;

        /// <summary>
        /// The UUID or the label of the base image used by the server. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
        /// to find either the right `label` or the right local image `ID` for a given `type`. Optional when creating an instance with an existing root volume.
        /// 
        /// You can check the available labels with our [CLI](https://www.scaleway.com/en/docs/compute/instances/api-cli/creating-managing-instances-with-cliv2/). ```scw marketplace image list```
        /// 
        /// To retrieve more information by label please use: ```scw marketplace image get label=&lt;LABEL&gt;```
        /// </summary>
        [Output("image")]
        public Output<string?> Image { get; private set; } = null!;

        /// <summary>
        /// The ID of the reserved IP that is attached to the server.
        /// </summary>
        [Output("ipId")]
        public Output<string?> IpId { get; private set; } = null!;

        /// <summary>
        /// List of ID of reserved IPs that are attached to the server. Cannot be used with `ip_id`.
        /// 
        /// &gt; `ip_id` to `ip_ids` migration: if moving the ip from the old `ip_id` field to the new `ip_ids`, it should not detach the ip.
        /// </summary>
        [Output("ipIds")]
        public Output<ImmutableArray<string>> IpIds { get; private set; } = null!;

        /// <summary>
        /// The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        /// <summary>
        /// The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Output("ipv6Gateway")]
        public Output<string> Ipv6Gateway { get; private set; } = null!;

        /// <summary>
        /// The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Output("ipv6PrefixLength")]
        public Output<int> Ipv6PrefixLength { get; private set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the server is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        /// 
        /// 
        /// &gt; **Important:** When updating `placement_group_id` the `state` must be set to `stopped`, otherwise it will fail.
        /// </summary>
        [Output("placementGroupId")]
        public Output<string?> PlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// True when the placement group policy is respected.
        /// </summary>
        [Output("placementGroupPolicyRespected")]
        public Output<bool> PlacementGroupPolicyRespected { get; private set; } = null!;

        /// <summary>
        /// The Scaleway internal IP address of the server.
        /// </summary>
        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        /// <summary>
        /// The private network associated with the server.
        /// Use the `pn_id` key to attach a [private_network](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea) on your instance.
        /// </summary>
        [Output("privateNetworks")]
        public Output<ImmutableArray<Outputs.ServerPrivateNetwork>> PrivateNetworks { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The public IP address of the server.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// The list of public IPs of the server.
        /// </summary>
        [Output("publicIps")]
        public Output<ImmutableArray<Outputs.ServerPublicIp>> PublicIps { get; private set; } = null!;

        /// <summary>
        /// If true, the server will be replaced if `type` is changed. Otherwise, the server will migrate.
        /// </summary>
        [Output("replaceOnTypeChange")]
        public Output<bool?> ReplaceOnTypeChange { get; private set; } = null!;

        /// <summary>
        /// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
        /// </summary>
        [Output("rootVolume")]
        public Output<Outputs.ServerRootVolume> RootVolume { get; private set; } = null!;

        /// <summary>
        /// If true, the server will support routed ips only. Changing it to true will migrate the server and its IP to routed type.
        /// 
        /// &gt; **Important:** Enabling routed ip will restart the server
        /// </summary>
        [Output("routedIpEnabled")]
        public Output<bool> RoutedIpEnabled { get; private set; } = null!;

        /// <summary>
        /// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The state of the server. Possible values are: `started`, `stopped` or `standby`.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The commercial type of the server.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// Updates to this field will migrate the server, local storage constraint must be respected. [More info](https://www.scaleway.com/en/docs/compute/instances/api-cli/migrating-instances/).
        /// Use `replace_on_type_change` to trigger replacement instead of migration.
        /// 
        /// &gt; **Important:** If `type` change and migration occurs, the server will be stopped and changed backed to its original state. It will be started again if it was running.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The user data associated with the server.
        /// Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// - Binary files using filebase64.
        /// </summary>
        [Output("userData")]
        public Output<ImmutableDictionary<string, string>> UserData { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("scaleway:instance/server:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:instance/server:Server", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/raeumlich/pulumi-scaleway/releases/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalVolumeIds")]
        private InputList<string>? _additionalVolumeIds;

        /// <summary>
        /// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
        /// attached to the server. Updates to this field will trigger a stop/start of the server.
        /// 
        /// &gt; **Important:** If this field contains local volumes, the `state` must be set to `stopped`, otherwise it will fail.
        /// 
        /// &gt; **Important:** If this field contains local volumes, you have to first detach them, in one apply, and then delete the volume in another apply.
        /// </summary>
        public InputList<string> AdditionalVolumeIds
        {
            get => _additionalVolumeIds ?? (_additionalVolumeIds = new InputList<string>());
            set => _additionalVolumeIds = value;
        }

        /// <summary>
        /// The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
        /// </summary>
        [Input("bootType")]
        public Input<string>? BootType { get; set; }

        /// <summary>
        /// The ID of the bootscript to use  (set boot_type to `bootscript`).
        /// </summary>
        [Input("bootscriptId")]
        public Input<string>? BootscriptId { get; set; }

        /// <summary>
        /// The cloud init script associated with this server
        /// </summary>
        [Input("cloudInit")]
        public Input<string>? CloudInit { get; set; }

        /// <summary>
        /// If true a dynamic IP will be attached to the server.
        /// </summary>
        [Input("enableDynamicIp")]
        public Input<bool>? EnableDynamicIp { get; set; }

        /// <summary>
        /// Determines if IPv6 is enabled for the server.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The UUID or the label of the base image used by the server. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
        /// to find either the right `label` or the right local image `ID` for a given `type`. Optional when creating an instance with an existing root volume.
        /// 
        /// You can check the available labels with our [CLI](https://www.scaleway.com/en/docs/compute/instances/api-cli/creating-managing-instances-with-cliv2/). ```scw marketplace image list```
        /// 
        /// To retrieve more information by label please use: ```scw marketplace image get label=&lt;LABEL&gt;```
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The ID of the reserved IP that is attached to the server.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// List of ID of reserved IPs that are attached to the server. Cannot be used with `ip_id`.
        /// 
        /// &gt; `ip_id` to `ip_ids` migration: if moving the ip from the old `ip_id` field to the new `ip_ids`, it should not detach the ip.
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        /// 
        /// 
        /// &gt; **Important:** When updating `placement_group_id` the `state` must be set to `stopped`, otherwise it will fail.
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        [Input("privateNetworks")]
        private InputList<Inputs.ServerPrivateNetworkArgs>? _privateNetworks;

        /// <summary>
        /// The private network associated with the server.
        /// Use the `pn_id` key to attach a [private_network](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea) on your instance.
        /// </summary>
        public InputList<Inputs.ServerPrivateNetworkArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.ServerPrivateNetworkArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("publicIps")]
        private InputList<Inputs.ServerPublicIpArgs>? _publicIps;

        /// <summary>
        /// The list of public IPs of the server.
        /// </summary>
        public InputList<Inputs.ServerPublicIpArgs> PublicIps
        {
            get => _publicIps ?? (_publicIps = new InputList<Inputs.ServerPublicIpArgs>());
            set => _publicIps = value;
        }

        /// <summary>
        /// If true, the server will be replaced if `type` is changed. Otherwise, the server will migrate.
        /// </summary>
        [Input("replaceOnTypeChange")]
        public Input<bool>? ReplaceOnTypeChange { get; set; }

        /// <summary>
        /// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
        /// </summary>
        [Input("rootVolume")]
        public Input<Inputs.ServerRootVolumeArgs>? RootVolume { get; set; }

        /// <summary>
        /// If true, the server will support routed ips only. Changing it to true will migrate the server and its IP to routed type.
        /// 
        /// &gt; **Important:** Enabling routed ip will restart the server
        /// </summary>
        [Input("routedIpEnabled")]
        public Input<bool>? RoutedIpEnabled { get; set; }

        /// <summary>
        /// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The state of the server. Possible values are: `started`, `stopped` or `standby`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The commercial type of the server.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// Updates to this field will migrate the server, local storage constraint must be respected. [More info](https://www.scaleway.com/en/docs/compute/instances/api-cli/migrating-instances/).
        /// Use `replace_on_type_change` to trigger replacement instead of migration.
        /// 
        /// &gt; **Important:** If `type` change and migration occurs, the server will be stopped and changed backed to its original state. It will be started again if it was running.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userData")]
        private InputMap<string>? _userData;

        /// <summary>
        /// The user data associated with the server.
        /// Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// - Binary files using filebase64.
        /// </summary>
        public InputMap<string> UserData
        {
            get => _userData ?? (_userData = new InputMap<string>());
            set => _userData = value;
        }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ServerArgs()
        {
        }
        public static new ServerArgs Empty => new ServerArgs();
    }

    public sealed class ServerState : global::Pulumi.ResourceArgs
    {
        [Input("additionalVolumeIds")]
        private InputList<string>? _additionalVolumeIds;

        /// <summary>
        /// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
        /// attached to the server. Updates to this field will trigger a stop/start of the server.
        /// 
        /// &gt; **Important:** If this field contains local volumes, the `state` must be set to `stopped`, otherwise it will fail.
        /// 
        /// &gt; **Important:** If this field contains local volumes, you have to first detach them, in one apply, and then delete the volume in another apply.
        /// </summary>
        public InputList<string> AdditionalVolumeIds
        {
            get => _additionalVolumeIds ?? (_additionalVolumeIds = new InputList<string>());
            set => _additionalVolumeIds = value;
        }

        /// <summary>
        /// The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
        /// </summary>
        [Input("bootType")]
        public Input<string>? BootType { get; set; }

        /// <summary>
        /// The ID of the bootscript to use  (set boot_type to `bootscript`).
        /// </summary>
        [Input("bootscriptId")]
        public Input<string>? BootscriptId { get; set; }

        /// <summary>
        /// The cloud init script associated with this server
        /// </summary>
        [Input("cloudInit")]
        public Input<string>? CloudInit { get; set; }

        /// <summary>
        /// If true a dynamic IP will be attached to the server.
        /// </summary>
        [Input("enableDynamicIp")]
        public Input<bool>? EnableDynamicIp { get; set; }

        /// <summary>
        /// Determines if IPv6 is enabled for the server.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The UUID or the label of the base image used by the server. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
        /// to find either the right `label` or the right local image `ID` for a given `type`. Optional when creating an instance with an existing root volume.
        /// 
        /// You can check the available labels with our [CLI](https://www.scaleway.com/en/docs/compute/instances/api-cli/creating-managing-instances-with-cliv2/). ```scw marketplace image list```
        /// 
        /// To retrieve more information by label please use: ```scw marketplace image get label=&lt;LABEL&gt;```
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The ID of the reserved IP that is attached to the server.
        /// </summary>
        [Input("ipId")]
        public Input<string>? IpId { get; set; }

        [Input("ipIds")]
        private InputList<string>? _ipIds;

        /// <summary>
        /// List of ID of reserved IPs that are attached to the server. Cannot be used with `ip_id`.
        /// 
        /// &gt; `ip_id` to `ip_ids` migration: if moving the ip from the old `ip_id` field to the new `ip_ids`, it should not detach the ip.
        /// </summary>
        public InputList<string> IpIds
        {
            get => _ipIds ?? (_ipIds = new InputList<string>());
            set => _ipIds = value;
        }

        /// <summary>
        /// The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        /// <summary>
        /// The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Input("ipv6Gateway")]
        public Input<string>? Ipv6Gateway { get; set; }

        /// <summary>
        /// The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        [Input("ipv6PrefixLength")]
        public Input<int>? Ipv6PrefixLength { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the server is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        /// 
        /// 
        /// &gt; **Important:** When updating `placement_group_id` the `state` must be set to `stopped`, otherwise it will fail.
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        /// <summary>
        /// True when the placement group policy is respected.
        /// </summary>
        [Input("placementGroupPolicyRespected")]
        public Input<bool>? PlacementGroupPolicyRespected { get; set; }

        /// <summary>
        /// The Scaleway internal IP address of the server.
        /// </summary>
        [Input("privateIp")]
        public Input<string>? PrivateIp { get; set; }

        [Input("privateNetworks")]
        private InputList<Inputs.ServerPrivateNetworkGetArgs>? _privateNetworks;

        /// <summary>
        /// The private network associated with the server.
        /// Use the `pn_id` key to attach a [private_network](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea) on your instance.
        /// </summary>
        public InputList<Inputs.ServerPrivateNetworkGetArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.ServerPrivateNetworkGetArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The public IP address of the server.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("publicIps")]
        private InputList<Inputs.ServerPublicIpGetArgs>? _publicIps;

        /// <summary>
        /// The list of public IPs of the server.
        /// </summary>
        public InputList<Inputs.ServerPublicIpGetArgs> PublicIps
        {
            get => _publicIps ?? (_publicIps = new InputList<Inputs.ServerPublicIpGetArgs>());
            set => _publicIps = value;
        }

        /// <summary>
        /// If true, the server will be replaced if `type` is changed. Otherwise, the server will migrate.
        /// </summary>
        [Input("replaceOnTypeChange")]
        public Input<bool>? ReplaceOnTypeChange { get; set; }

        /// <summary>
        /// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
        /// </summary>
        [Input("rootVolume")]
        public Input<Inputs.ServerRootVolumeGetArgs>? RootVolume { get; set; }

        /// <summary>
        /// If true, the server will support routed ips only. Changing it to true will migrate the server and its IP to routed type.
        /// 
        /// &gt; **Important:** Enabling routed ip will restart the server
        /// </summary>
        [Input("routedIpEnabled")]
        public Input<bool>? RoutedIpEnabled { get; set; }

        /// <summary>
        /// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The state of the server. Possible values are: `started`, `stopped` or `standby`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The commercial type of the server.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// Updates to this field will migrate the server, local storage constraint must be respected. [More info](https://www.scaleway.com/en/docs/compute/instances/api-cli/migrating-instances/).
        /// Use `replace_on_type_change` to trigger replacement instead of migration.
        /// 
        /// &gt; **Important:** If `type` change and migration occurs, the server will be stopped and changed backed to its original state. It will be started again if it was running.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userData")]
        private InputMap<string>? _userData;

        /// <summary>
        /// The user data associated with the server.
        /// Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// - Binary files using filebase64.
        /// </summary>
        public InputMap<string> UserData
        {
            get => _userData ?? (_userData = new InputMap<string>());
            set => _userData = value;
        }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public ServerState()
        {
        }
        public static new ServerState Empty => new ServerState();
    }
}