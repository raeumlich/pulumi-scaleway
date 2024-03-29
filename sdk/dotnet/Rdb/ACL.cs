// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Rdb
{
    /// <summary>
    /// Creates and manages Scaleway Database instance authorized IPs.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
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
    ///     var main = new Scaleway.Rdb.ACL("main", new()
    ///     {
    ///         InstanceId = scaleway_rdb_instance.Main.Id,
    ///         AclRules = new[]
    ///         {
    ///             new Scaleway.Rdb.Inputs.ACLAclRuleArgs
    ///             {
    ///                 Ip = "1.2.3.4/32",
    ///                 Description = "foo",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Database Instance can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:rdb/aCL:ACL acl01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:rdb/aCL:ACL")]
    public partial class ACL : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        [Output("aclRules")]
        public Output<ImmutableArray<Outputs.ACLAclRule>> AclRules { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a ACL resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ACL(string name, ACLArgs args, CustomResourceOptions? options = null)
            : base("scaleway:rdb/aCL:ACL", name, args ?? new ACLArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ACL(string name, Input<string> id, ACLState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:rdb/aCL:ACL", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ACL resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ACL Get(string name, Input<string> id, ACLState? state = null, CustomResourceOptions? options = null)
        {
            return new ACL(name, id, state, options);
        }
    }

    public sealed class ACLArgs : global::Pulumi.ResourceArgs
    {
        [Input("aclRules", required: true)]
        private InputList<Inputs.ACLAclRuleArgs>? _aclRules;

        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        public InputList<Inputs.ACLAclRuleArgs> AclRules
        {
            get => _aclRules ?? (_aclRules = new InputList<Inputs.ACLAclRuleArgs>());
            set => _aclRules = value;
        }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ACLArgs()
        {
        }
        public static new ACLArgs Empty => new ACLArgs();
    }

    public sealed class ACLState : global::Pulumi.ResourceArgs
    {
        [Input("aclRules")]
        private InputList<Inputs.ACLAclRuleGetArgs>? _aclRules;

        /// <summary>
        /// A list of ACLs (structure is described below)
        /// </summary>
        public InputList<Inputs.ACLAclRuleGetArgs> AclRules
        {
            get => _aclRules ?? (_aclRules = new InputList<Inputs.ACLAclRuleGetArgs>());
            set => _aclRules = value;
        }

        /// <summary>
        /// UUID of the rdb instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database ACL.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ACLState()
        {
        }
        public static new ACLState Empty => new ACLState();
    }
}
