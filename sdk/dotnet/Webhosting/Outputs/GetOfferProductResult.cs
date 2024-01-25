// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Webhosting.Outputs
{

    [OutputType]
    public sealed class GetOfferProductResult
    {
        /// <summary>
        /// The quota of databases.
        /// </summary>
        public readonly int DatabasesQuota;
        /// <summary>
        /// The quota of email accounts.
        /// </summary>
        public readonly int EmailAccountsQuota;
        /// <summary>
        /// The quota of email storage.
        /// </summary>
        public readonly int EmailStorageQuota;
        /// <summary>
        /// The quota of hosting storage.
        /// </summary>
        public readonly int HostingStorageQuota;
        /// <summary>
        /// The offer name. Only one of `name` and `offer_id` should be specified.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The product option.
        /// </summary>
        public readonly bool Option;
        /// <summary>
        /// The capacity of the memory in GB.
        /// </summary>
        public readonly int Ram;
        /// <summary>
        /// If support is included.
        /// </summary>
        public readonly bool SupportIncluded;
        /// <summary>
        /// The number of cores.
        /// </summary>
        public readonly int VCpu;

        [OutputConstructor]
        private GetOfferProductResult(
            int databasesQuota,

            int emailAccountsQuota,

            int emailStorageQuota,

            int hostingStorageQuota,

            string name,

            bool option,

            int ram,

            bool supportIncluded,

            int vCpu)
        {
            DatabasesQuota = databasesQuota;
            EmailAccountsQuota = emailAccountsQuota;
            EmailStorageQuota = emailStorageQuota;
            HostingStorageQuota = hostingStorageQuota;
            Name = name;
            Option = option;
            Ram = ram;
            SupportIncluded = supportIncluded;
            VCpu = vCpu;
        }
    }
}
