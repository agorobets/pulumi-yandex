// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbMysqlUser
    {
        /// <summary>
        /// Get information about a Yandex Managed MySQL user. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_user/d_mdb_mysql_user_1.tf" }}
        /// </summary>
        public static Task<GetMdbMysqlUserResult> InvokeAsync(GetMdbMysqlUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMdbMysqlUserResult>("yandex:index/getMdbMysqlUser:getMdbMysqlUser", args ?? new GetMdbMysqlUserArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Managed MySQL user. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_user/d_mdb_mysql_user_1.tf" }}
        /// </summary>
        public static Output<GetMdbMysqlUserResult> Invoke(GetMdbMysqlUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbMysqlUserResult>("yandex:index/getMdbMysqlUser:getMdbMysqlUser", args ?? new GetMdbMysqlUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Managed MySQL user. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_user/d_mdb_mysql_user_1.tf" }}
        /// </summary>
        public static Output<GetMdbMysqlUserResult> Invoke(GetMdbMysqlUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbMysqlUserResult>("yandex:index/getMdbMysqlUser:getMdbMysqlUser", args ?? new GetMdbMysqlUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbMysqlUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the MySQL cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the MySQL user.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("permissions")]
        private List<Inputs.GetMdbMysqlUserPermissionArgs>? _permissions;

        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public List<Inputs.GetMdbMysqlUserPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new List<Inputs.GetMdbMysqlUserPermissionArgs>());
            set => _permissions = value;
        }

        public GetMdbMysqlUserArgs()
        {
        }
        public static new GetMdbMysqlUserArgs Empty => new GetMdbMysqlUserArgs();
    }

    public sealed class GetMdbMysqlUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the MySQL cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the MySQL user.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("permissions")]
        private InputList<Inputs.GetMdbMysqlUserPermissionInputArgs>? _permissions;

        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public InputList<Inputs.GetMdbMysqlUserPermissionInputArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.GetMdbMysqlUserPermissionInputArgs>());
            set => _permissions = value;
        }

        public GetMdbMysqlUserInvokeArgs()
        {
        }
        public static new GetMdbMysqlUserInvokeArgs Empty => new GetMdbMysqlUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetMdbMysqlUserResult
    {
        /// <summary>
        /// Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD`
        /// </summary>
        public readonly string AuthenticationPlugin;
        public readonly string ClusterId;
        /// <summary>
        /// User's connection limits. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlUserConnectionLimitResult> ConnectionLimits;
        /// <summary>
        /// List user's global permissions. Allowed values: `REPLICATION_CLIENT`, `REPLICATION_SLAVE`, `PROCESS` or empty list.
        /// </summary>
        public readonly ImmutableArray<string> GlobalPermissions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set of permissions granted to the user. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlUserPermissionResult> Permissions;

        [OutputConstructor]
        private GetMdbMysqlUserResult(
            string authenticationPlugin,

            string clusterId,

            ImmutableArray<Outputs.GetMdbMysqlUserConnectionLimitResult> connectionLimits,

            ImmutableArray<string> globalPermissions,

            string id,

            string name,

            ImmutableArray<Outputs.GetMdbMysqlUserPermissionResult> permissions)
        {
            AuthenticationPlugin = authenticationPlugin;
            ClusterId = clusterId;
            ConnectionLimits = connectionLimits;
            GlobalPermissions = globalPermissions;
            Id = id;
            Name = name;
            Permissions = permissions;
        }
    }
}
