// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage
{
    public static class GetBucketPolicy
    {
        /// <summary>
        /// Gets information about the Bucket's policy.
        /// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
        ///     var main = Scaleway.Objectstorage.GetBucketPolicy.Invoke(new()
        ///     {
        ///         Bucket = "bucket.test.com",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketPolicyResult> InvokeAsync(GetBucketPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketPolicyResult>("scaleway:objectstorage/getBucketPolicy:getBucketPolicy", args ?? new GetBucketPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about the Bucket's policy.
        /// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
        ///     var main = Scaleway.Objectstorage.GetBucketPolicy.Invoke(new()
        ///     {
        ///         Bucket = "bucket.test.com",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBucketPolicyResult> Invoke(GetBucketPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketPolicyResult>("scaleway:objectstorage/getBucketPolicy:getBucketPolicy", args ?? new GetBucketPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Object Storage exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetBucketPolicyArgs()
        {
        }
        public static new GetBucketPolicyArgs Empty => new GetBucketPolicyArgs();
    }

    public sealed class GetBucketPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Object Storage exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetBucketPolicyInvokeArgs()
        {
        }
        public static new GetBucketPolicyInvokeArgs Empty => new GetBucketPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketPolicyResult
    {
        public readonly string Bucket;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The bucket's policy in JSON format.
        /// </summary>
        public readonly string Policy;
        public readonly string? ProjectId;
        public readonly string? Region;

        [OutputConstructor]
        private GetBucketPolicyResult(
            string bucket,

            string id,

            string policy,

            string? projectId,

            string? region)
        {
            Bucket = bucket;
            Id = id;
            Policy = policy;
            ProjectId = projectId;
            Region = region;
        }
    }
}