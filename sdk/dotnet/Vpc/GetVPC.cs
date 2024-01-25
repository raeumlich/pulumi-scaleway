// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Vpc
{
    public static class GetVPC
    {
        /// <summary>
        /// Gets information about a Scaleway Virtual Private Cloud.
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
        ///     var byName = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var @default = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         IsDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVPCResult> InvokeAsync(GetVPCArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVPCResult>("scaleway:vpc/getVPC:getVPC", args ?? new GetVPCArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Scaleway Virtual Private Cloud.
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
        ///     var byName = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         Name = "foobar",
        ///     });
        /// 
        ///     var byId = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         VpcId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var @default = Scaleway.Vpc.GetVPC.Invoke(new()
        ///     {
        ///         IsDefault = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVPCResult> Invoke(GetVPCInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVPCResult>("scaleway:vpc/getVPC:getVPC", args ?? new GetVPCInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVPCArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// To get default VPC's information.
        /// </summary>
        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        /// <summary>
        /// Name of the VPC. One of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the organization the VPC is associated with.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// ID of the VPC. One of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVPCArgs()
        {
        }
        public static new GetVPCArgs Empty => new GetVPCArgs();
    }

    public sealed class GetVPCInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// To get default VPC's information.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Name of the VPC. One of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the organization the VPC is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the VPC is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the VPC. One of `name` and `vpc_id` should be specified.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetVPCInvokeArgs()
        {
        }
        public static new GetVPCInvokeArgs Empty => new GetVPCInvokeArgs();
    }


    [OutputType]
    public sealed class GetVPCResult
    {
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IsDefault;
        public readonly string? Name;
        public readonly string OrganizationId;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string? VpcId;

        [OutputConstructor]
        private GetVPCResult(
            string createdAt,

            string id,

            bool? isDefault,

            string? name,

            string organizationId,

            string? projectId,

            string? region,

            ImmutableArray<string> tags,

            string updatedAt,

            string? vpcId)
        {
            CreatedAt = createdAt;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
            UpdatedAt = updatedAt;
            VpcId = vpcId;
        }
    }
}
