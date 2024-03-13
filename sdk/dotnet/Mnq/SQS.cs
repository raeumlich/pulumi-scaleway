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
    /// Activate Scaleway Messaging and queuing SQS for a project.
    /// For further information please check
    /// our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// Activate SQS for default project
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
    ///     var main = new Scaleway.Mnq.SQS("main");
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// Activate SQS for a specific project
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
    ///     var forProject = new Scaleway.Mnq.SQS("forProject", new()
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
    /// SQS status can be imported using the `{region}/{project_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:mnq/sQS:SQS main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:mnq/sQS:SQS")]
    public partial class SQS : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint of the SQS service for this project.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a SQS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SQS(string name, SQSArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/sQS:SQS", name, args ?? new SQSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SQS(string name, Input<string> id, SQSState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:mnq/sQS:SQS", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SQS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SQS Get(string name, Input<string> id, SQSState? state = null, CustomResourceOptions? options = null)
        {
            return new SQS(name, id, state, options);
        }
    }

    public sealed class SQSArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SQSArgs()
        {
        }
        public static new SQSArgs Empty => new SQSArgs();
    }

    public sealed class SQSState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint of the SQS service for this project.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the sqs will be enabled for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which sqs will be enabled.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public SQSState()
        {
        }
        public static new SQSState Empty => new SQSState();
    }
}
