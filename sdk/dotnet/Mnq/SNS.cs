// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Mnq
{
    /// <summary>
    /// Activate Scaleway Messaging and queuing SNS for a project.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// Activate SNS for default project
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
    ///     var main = new Scaleway.Mnq.SNS("main");
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Activate SNS for a specific project
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
    ///     var project = Scaleway.Account.GetProject.Invoke(new()
    ///     {
    ///         Name = "default",
    ///     });
    /// 
    ///     // For specific project in default region
    ///     var forProject = new Scaleway.Mnq.SNS("forProject", new()
    ///     {
    ///         ProjectId = project.Apply(getProjectResult =&gt; getProjectResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SNS status can be imported using the `{region}/{project_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:mnq/sNS:SNS main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:mnq/sNS:SNS")]
    public partial class SNS : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the SNS service for this project.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a SNS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SNS(string name, SNSArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/sNS:SNS", name, args ?? new SNSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SNS(string name, Input<string> id, SNSState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/sNS:SNS", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SNS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SNS Get(string name, Input<string> id, SNSState? state = null, CustomResourceOptions? options = null)
        {
            return new SNS(name, id, state, options);
        }
    }

    public sealed class SNSArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SNSArgs()
        {
        }
        public static new SNSArgs Empty => new SNSArgs();
    }

    public sealed class SNSState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the SNS service for this project.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sns will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sns will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SNSState()
        {
        }
        public static new SNSState Empty => new SNSState();
    }
}
