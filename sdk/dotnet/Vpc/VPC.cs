// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Vpc
{
    /// <summary>
    /// Creates and manages Scaleway Virtual Private Clouds.
    /// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/vpc/concepts/).
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
    ///     var vpc01 = new Scaleway.Vpc.VPC("vpc01", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             "demo",
    ///             "terraform",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPCs can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:vpc/vPC:VPC vpc_demo fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:vpc/vPC:VPC")]
    public partial class VPC : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date and time of VPC's creation (RFC 3339 format).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Defines whether the VPC is the default one for its Project.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC. If not provided it will be randomly generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the VPC is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`) The region of the VPC.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the VPC.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date and time of VPC's last update (RFC 3339 format).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a VPC resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPC(string name, VPCArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:vpc/vPC:VPC", name, args ?? new VPCArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPC(string name, Input<string> id, VPCState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:vpc/vPC:VPC", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VPC resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPC Get(string name, Input<string> id, VPCState? state = null, CustomResourceOptions? options = null)
        {
            return new VPC(name, id, state, options);
        }
    }

    public sealed class VPCArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPC. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the VPC.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the VPC.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public VPCArgs()
        {
        }
        public static new VPCArgs Empty => new VPCArgs();
    }

    public sealed class VPCState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date and time of VPC's creation (RFC 3339 format).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Defines whether the VPC is the default one for its Project.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the VPC. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the VPC is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region of the VPC.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the VPC.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Date and time of VPC's last update (RFC 3339 format).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public VPCState()
        {
        }
        public static new VPCState Empty => new VPCState();
    }
}
