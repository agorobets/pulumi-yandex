// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetLoggingGroup
    {
        /// <summary>
        /// Get information about a Yandex Cloud Logging group. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/logging_group/d_logging_group_1.tf" }}
        /// </summary>
        public static Task<GetLoggingGroupResult> InvokeAsync(GetLoggingGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoggingGroupResult>("yandex:index/getLoggingGroup:getLoggingGroup", args ?? new GetLoggingGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Logging group. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/logging_group/d_logging_group_1.tf" }}
        /// </summary>
        public static Output<GetLoggingGroupResult> Invoke(GetLoggingGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoggingGroupResult>("yandex:index/getLoggingGroup:getLoggingGroup", args ?? new GetLoggingGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Cloud Logging group. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/logging_group/d_logging_group_1.tf" }}
        /// </summary>
        public static Output<GetLoggingGroupResult> Invoke(GetLoggingGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoggingGroupResult>("yandex:index/getLoggingGroup:getLoggingGroup", args ?? new GetLoggingGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoggingGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        /// 
        /// &gt; If `group_id` is not specified `name` and `folder_id` will be used to designate Yandex Cloud Logging group.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        /// <summary>
        /// The Yandex Cloud Logging group ID.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// The Yandex Cloud Logging group name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetLoggingGroupArgs()
        {
        }
        public static new GetLoggingGroupArgs Empty => new GetLoggingGroupArgs();
    }

    public sealed class GetLoggingGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        /// 
        /// &gt; If `group_id` is not specified `name` and `folder_id` will be used to designate Yandex Cloud Logging group.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// The Yandex Cloud Logging group ID.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The Yandex Cloud Logging group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetLoggingGroupInvokeArgs()
        {
        }
        public static new GetLoggingGroupInvokeArgs Empty => new GetLoggingGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoggingGroupResult
    {
        /// <summary>
        /// ID of the cloud that the Yandex Cloud Logging group belong to.
        /// </summary>
        public readonly string CloudId;
        /// <summary>
        /// The Yandex Cloud Logging group creation timestamp.
        /// </summary>
        public readonly string CreatedAt;
        public readonly string DataStream;
        /// <summary>
        /// The Yandex Cloud Logging group description.
        /// </summary>
        public readonly string Description;
        public readonly string FolderId;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs assigned to the Yandex Cloud Logging group.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        /// <summary>
        /// The Yandex Cloud Logging group log entries retention period.
        /// </summary>
        public readonly string RetentionPeriod;
        /// <summary>
        /// The Yandex Cloud Logging group status.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetLoggingGroupResult(
            string cloudId,

            string createdAt,

            string dataStream,

            string description,

            string folderId,

            string groupId,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            string retentionPeriod,

            string status)
        {
            CloudId = cloudId;
            CreatedAt = createdAt;
            DataStream = dataStream;
            Description = description;
            FolderId = folderId;
            GroupId = groupId;
            Id = id;
            Labels = labels;
            Name = name;
            RetentionPeriod = retentionPeriod;
            Status = status;
        }
    }
}
