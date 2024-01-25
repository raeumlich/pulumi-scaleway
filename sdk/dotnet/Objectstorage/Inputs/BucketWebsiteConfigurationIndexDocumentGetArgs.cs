// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Objectstorage.Inputs
{

    public sealed class BucketWebsiteConfigurationIndexDocumentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A suffix that is appended to a request that is for a directory on the website endpoint.
        /// 
        /// &gt; **Important:** The suffix must not be empty and must not include a slash character. The routing is not supported.
        /// </summary>
        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public BucketWebsiteConfigurationIndexDocumentGetArgs()
        {
        }
        public static new BucketWebsiteConfigurationIndexDocumentGetArgs Empty => new BucketWebsiteConfigurationIndexDocumentGetArgs();
    }
}
