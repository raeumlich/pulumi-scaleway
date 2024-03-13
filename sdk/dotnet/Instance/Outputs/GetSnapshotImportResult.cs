// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Instance.Outputs
{

    [OutputType]
    public sealed class GetSnapshotImportResult
    {
        /// <summary>
        /// Bucket containing qcow
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Key of the qcow file in the specified bucket
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private GetSnapshotImportResult(
            string bucket,

            string key)
        {
            Bucket = bucket;
            Key = key;
        }
    }
}
