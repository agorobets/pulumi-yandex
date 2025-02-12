// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbMysqlCluster
    {
        /// <summary>
        /// Get information about a Yandex Managed MySQL cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_cluster/d_mdb_mysql_cluster_1.tf" }}
        /// </summary>
        public static Task<GetMdbMysqlClusterResult> InvokeAsync(GetMdbMysqlClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMdbMysqlClusterResult>("yandex:index/getMdbMysqlCluster:getMdbMysqlCluster", args ?? new GetMdbMysqlClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Managed MySQL cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_cluster/d_mdb_mysql_cluster_1.tf" }}
        /// </summary>
        public static Output<GetMdbMysqlClusterResult> Invoke(GetMdbMysqlClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbMysqlClusterResult>("yandex:index/getMdbMysqlCluster:getMdbMysqlCluster", args ?? new GetMdbMysqlClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Yandex Managed MySQL cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/).
        /// 
        /// ## Example Usage
        /// 
        /// {{ tffile "examples/mdb_mysql_cluster/d_mdb_mysql_cluster_1.tf" }}
        /// </summary>
        public static Output<GetMdbMysqlClusterResult> Invoke(GetMdbMysqlClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbMysqlClusterResult>("yandex:index/getMdbMysqlCluster:getMdbMysqlCluster", args ?? new GetMdbMysqlClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbMysqlClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Access policy to the MySQL cluster. The structure is documented below.
        /// </summary>
        [Input("access")]
        public Inputs.GetMdbMysqlClusterAccessArgs? Access { get; set; }

        /// <summary>
        /// The ID of the MySQL cluster.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("deletionProtection")]
        public bool? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the MySQL cluster.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the MySQL cluster.
        /// </summary>
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        [Input("mysqlConfig")]
        private Dictionary<string, string>? _mysqlConfig;

        /// <summary>
        /// MySQL cluster config.
        /// </summary>
        public Dictionary<string, string> MysqlConfig
        {
            get => _mysqlConfig ?? (_mysqlConfig = new Dictionary<string, string>());
            set => _mysqlConfig = value;
        }

        /// <summary>
        /// The name of the MySQL cluster.
        /// 
        /// &gt; Either `cluster_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetMdbMysqlClusterArgs()
        {
        }
        public static new GetMdbMysqlClusterArgs Empty => new GetMdbMysqlClusterArgs();
    }

    public sealed class GetMdbMysqlClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Access policy to the MySQL cluster. The structure is documented below.
        /// </summary>
        [Input("access")]
        public Input<Inputs.GetMdbMysqlClusterAccessInputArgs>? Access { get; set; }

        /// <summary>
        /// The ID of the MySQL cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the MySQL cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the MySQL cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("mysqlConfig")]
        private InputMap<string>? _mysqlConfig;

        /// <summary>
        /// MySQL cluster config.
        /// </summary>
        public InputMap<string> MysqlConfig
        {
            get => _mysqlConfig ?? (_mysqlConfig = new InputMap<string>());
            set => _mysqlConfig = value;
        }

        /// <summary>
        /// The name of the MySQL cluster.
        /// 
        /// &gt; Either `cluster_id` or `name` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMdbMysqlClusterInvokeArgs()
        {
        }
        public static new GetMdbMysqlClusterInvokeArgs Empty => new GetMdbMysqlClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetMdbMysqlClusterResult
    {
        /// <summary>
        /// Access policy to the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly Outputs.GetMdbMysqlClusterAccessResult Access;
        /// <summary>
        /// The period in days during which backups are stored.
        /// </summary>
        public readonly int BackupRetainPeriodDays;
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterBackupWindowStartResult> BackupWindowStarts;
        public readonly string ClusterId;
        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A database of the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterDatabaseResult> Databases;
        public readonly bool DeletionProtection;
        /// <summary>
        /// Description of the MySQL cluster.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Deployment environment of the MySQL cluster.
        /// </summary>
        public readonly string Environment;
        public readonly string FolderId;
        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        public readonly string Health;
        public readonly ImmutableArray<string> HostGroupIds;
        /// <summary>
        /// A host of the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterHostResult> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of key/value label pairs to assign to the MySQL cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Maintenance window settings of the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterMaintenanceWindowResult> MaintenanceWindows;
        /// <summary>
        /// MySQL cluster config.
        /// </summary>
        public readonly ImmutableDictionary<string, string> MysqlConfig;
        /// <summary>
        /// The name of the database.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the network, to which the MySQL cluster belongs.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// Cluster performance diagnostics settings. The structure is documented below. [YC Documentation](https://cloud.yandex.com/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterPerformanceDiagnosticResult> PerformanceDiagnostics;
        /// <summary>
        /// Resources allocated to hosts of the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterResourceResult> Resources;
        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// Status of the cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A user of the MySQL cluster. The structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMdbMysqlClusterUserResult> Users;
        /// <summary>
        /// Version of the MySQL cluster.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetMdbMysqlClusterResult(
            Outputs.GetMdbMysqlClusterAccessResult access,

            int backupRetainPeriodDays,

            ImmutableArray<Outputs.GetMdbMysqlClusterBackupWindowStartResult> backupWindowStarts,

            string clusterId,

            string createdAt,

            ImmutableArray<Outputs.GetMdbMysqlClusterDatabaseResult> databases,

            bool deletionProtection,

            string? description,

            string environment,

            string folderId,

            string health,

            ImmutableArray<string> hostGroupIds,

            ImmutableArray<Outputs.GetMdbMysqlClusterHostResult> hosts,

            string id,

            ImmutableDictionary<string, string>? labels,

            ImmutableArray<Outputs.GetMdbMysqlClusterMaintenanceWindowResult> maintenanceWindows,

            ImmutableDictionary<string, string> mysqlConfig,

            string name,

            string networkId,

            ImmutableArray<Outputs.GetMdbMysqlClusterPerformanceDiagnosticResult> performanceDiagnostics,

            ImmutableArray<Outputs.GetMdbMysqlClusterResourceResult> resources,

            ImmutableArray<string> securityGroupIds,

            string status,

            ImmutableArray<Outputs.GetMdbMysqlClusterUserResult> users,

            string version)
        {
            Access = access;
            BackupRetainPeriodDays = backupRetainPeriodDays;
            BackupWindowStarts = backupWindowStarts;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Databases = databases;
            DeletionProtection = deletionProtection;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            Health = health;
            HostGroupIds = hostGroupIds;
            Hosts = hosts;
            Id = id;
            Labels = labels;
            MaintenanceWindows = maintenanceWindows;
            MysqlConfig = mysqlConfig;
            Name = name;
            NetworkId = networkId;
            PerformanceDiagnostics = performanceDiagnostics;
            Resources = resources;
            SecurityGroupIds = securityGroupIds;
            Status = status;
            Users = users;
            Version = version;
        }
    }
}
