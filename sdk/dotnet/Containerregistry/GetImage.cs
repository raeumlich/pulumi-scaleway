// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Containerregistry
{
    public static class GetImage
    {
        /// <summary>
        /// Gets information about a registry image.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myImage = Scaleway.Containerregistry.GetImage.Invoke(new()
        ///     {
        ///         ImageId = "11111111-1111-1111-1111-111111111111",
        ///         NamespaceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("scaleway:containerregistry/getImage:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a registry image.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myImage = Scaleway.Containerregistry.GetImage.Invoke(new()
        ///     {
        ///         ImageId = "11111111-1111-1111-1111-111111111111",
        ///         NamespaceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("scaleway:containerregistry/getImage:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The image ID.
        /// Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("imageId")]
        public string? ImageId { get; set; }

        /// <summary>
        /// The image name.
        /// Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The namespace ID in which the image is.
        /// </summary>
        [Input("namespaceId")]
        public string? NamespaceId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the image exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags associated with the registry image
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The image ID.
        /// Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The image name.
        /// Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace ID in which the image is.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the image exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the registry image
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ImageId;
        public readonly string? Name;
        public readonly string NamespaceId;
        /// <summary>
        /// The organization ID the image is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string ProjectId;
        public readonly string Region;
        /// <summary>
        /// The size of the registry image.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The tags associated with the registry image
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The date the image of the last update
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The privacy policy of the registry image.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetImageResult(
            string id,

            string? imageId,

            string? name,

            string namespaceId,

            string organizationId,

            string projectId,

            string region,

            int size,

            ImmutableArray<string> tags,

            string updatedAt,

            string visibility)
        {
            Id = id;
            ImageId = imageId;
            Name = name;
            NamespaceId = namespaceId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Size = size;
            Tags = tags;
            UpdatedAt = updatedAt;
            Visibility = visibility;
        }
    }
}