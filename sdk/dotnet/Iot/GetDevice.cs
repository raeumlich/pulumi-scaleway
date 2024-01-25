// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Iot
{
    public static class GetDevice
    {
        /// <summary>
        /// Gets information about an IOT Device.
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
        ///     var myDevice = Scaleway.Iot.GetDevice.Invoke(new()
        ///     {
        ///         DeviceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDeviceResult> InvokeAsync(GetDeviceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeviceResult>("scaleway:iot/getDevice:getDevice", args ?? new GetDeviceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an IOT Device.
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
        ///     var myDevice = Scaleway.Iot.GetDevice.Invoke(new()
        ///     {
        ///         DeviceId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDeviceResult> Invoke(GetDeviceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeviceResult>("scaleway:iot/getDevice:getDevice", args ?? new GetDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeviceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device ID.
        /// Only one of the `name` and `device_id` should be specified.
        /// </summary>
        [Input("deviceId")]
        public string? DeviceId { get; set; }

        /// <summary>
        /// The hub ID.
        /// </summary>
        [Input("hubId")]
        public string? HubId { get; set; }

        /// <summary>
        /// The name of the Hub.
        /// Only one of the `name` and `device_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the hub exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetDeviceArgs()
        {
        }
        public static new GetDeviceArgs Empty => new GetDeviceArgs();
    }

    public sealed class GetDeviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device ID.
        /// Only one of the `name` and `device_id` should be specified.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The hub ID.
        /// </summary>
        [Input("hubId")]
        public Input<string>? HubId { get; set; }

        /// <summary>
        /// The name of the Hub.
        /// Only one of the `name` and `device_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the hub exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetDeviceInvokeArgs()
        {
        }
        public static new GetDeviceInvokeArgs Empty => new GetDeviceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeviceResult
    {
        public readonly bool AllowInsecure;
        public readonly bool AllowMultipleConnections;
        public readonly ImmutableArray<Outputs.GetDeviceCertificateResult> Certificates;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string? DeviceId;
        public readonly string HubId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsConnected;
        public readonly string LastActivityAt;
        public readonly ImmutableArray<Outputs.GetDeviceMessageFilterResult> MessageFilters;
        public readonly string? Name;
        public readonly string? Region;
        public readonly string Status;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetDeviceResult(
            bool allowInsecure,

            bool allowMultipleConnections,

            ImmutableArray<Outputs.GetDeviceCertificateResult> certificates,

            string createdAt,

            string description,

            string? deviceId,

            string hubId,

            string id,

            bool isConnected,

            string lastActivityAt,

            ImmutableArray<Outputs.GetDeviceMessageFilterResult> messageFilters,

            string? name,

            string? region,

            string status,

            string updatedAt)
        {
            AllowInsecure = allowInsecure;
            AllowMultipleConnections = allowMultipleConnections;
            Certificates = certificates;
            CreatedAt = createdAt;
            Description = description;
            DeviceId = deviceId;
            HubId = hubId;
            Id = id;
            IsConnected = isConnected;
            LastActivityAt = lastActivityAt;
            MessageFilters = messageFilters;
            Name = name;
            Region = region;
            Status = status;
            UpdatedAt = updatedAt;
        }
    }
}