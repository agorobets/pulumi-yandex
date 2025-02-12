// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// Manages a Redis cluster within the Yandex Cloud. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts).
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/mdb_redis_cluster/r_mdb_redis_cluster_1.tf" }}
    /// 
    /// Example of creating a sharded Redis Cluster.
    /// 
    /// {{ tffile "examples/mdb_redis_cluster/r_mdb_redis_cluster_2.tf" }}
    /// 
    /// ## Import
    /// 
    /// A cluster can be imported using the `id` of the resource, e.g.
    /// 
    /// ```sh
    /// $ pulumi import yandex:index/mdbRedisCluster:MdbRedisCluster foo cluster_id
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/mdbRedisCluster:MdbRedisCluster")]
    public partial class MdbRedisCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access policy to the Redis cluster. The structure is documented below.
        /// </summary>
        [Output("access")]
        public Output<Outputs.MdbRedisClusterAccess> Access { get; private set; } = null!;

        /// <summary>
        /// Announce fqdn instead of ip address.
        /// </summary>
        [Output("announceHostnames")]
        public Output<bool?> AnnounceHostnames { get; private set; } = null!;

        /// <summary>
        /// Configuration of the Redis cluster. The structure is documented below.
        /// </summary>
        [Output("config")]
        public Output<Outputs.MdbRedisClusterConfig> Config { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Inhibits deletion of the cluster. Can be either `true` or `false`.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Description of the Redis cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("diskSizeAutoscaling")]
        public Output<Outputs.MdbRedisClusterDiskSizeAutoscaling> DiskSizeAutoscaling { get; private set; } = null!;

        /// <summary>
        /// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        /// <summary>
        /// A host of the Redis cluster. The structure is documented below.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<Outputs.MdbRedisClusterHost>> Hosts { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the Redis cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("maintenanceWindow")]
        public Output<Outputs.MdbRedisClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// Name of the Redis cluster. Provided by the client when the cluster is created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the network, to which the Redis cluster belongs.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Persistence mode. Possible values: `ON`, `OFF`.
        /// </summary>
        [Output("persistenceMode")]
        public Output<string> PersistenceMode { get; private set; } = null!;

        /// <summary>
        /// Resources allocated to hosts of the Redis cluster. The structure is documented below.
        /// </summary>
        [Output("resources")]
        public Output<Outputs.MdbRedisClusterResources> Resources { get; private set; } = null!;

        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Redis Cluster mode enabled/disabled. Enables sharding when cluster non-sharded. If cluster is sharded - disabling is not allowed.
        /// </summary>
        [Output("sharded")]
        public Output<bool?> Sharded { get; private set; } = null!;

        /// <summary>
        /// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// TLS support mode enabled/disabled.
        /// </summary>
        [Output("tlsEnabled")]
        public Output<bool> TlsEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a MdbRedisCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbRedisCluster(string name, MdbRedisClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbRedisCluster:MdbRedisCluster", name, args ?? new MdbRedisClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbRedisCluster(string name, Input<string> id, MdbRedisClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbRedisCluster:MdbRedisCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/agorobets/pulumi-yandex/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbRedisCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbRedisCluster Get(string name, Input<string> id, MdbRedisClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbRedisCluster(name, id, state, options);
        }
    }

    public sealed class MdbRedisClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbRedisClusterAccessArgs>? Access { get; set; }

        /// <summary>
        /// Announce fqdn instead of ip address.
        /// </summary>
        [Input("announceHostnames")]
        public Input<bool>? AnnounceHostnames { get; set; }

        /// <summary>
        /// Configuration of the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.MdbRedisClusterConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// Inhibits deletion of the cluster. Can be either `true` or `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the Redis cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskSizeAutoscaling")]
        public Input<Inputs.MdbRedisClusterDiskSizeAutoscalingArgs>? DiskSizeAutoscaling { get; set; }

        /// <summary>
        /// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("hosts", required: true)]
        private InputList<Inputs.MdbRedisClusterHostArgs>? _hosts;

        /// <summary>
        /// A host of the Redis cluster. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbRedisClusterHostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbRedisClusterHostArgs>());
            set => _hosts = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Redis cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("maintenanceWindow")]
        public Input<Inputs.MdbRedisClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Name of the Redis cluster. Provided by the client when the cluster is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network, to which the Redis cluster belongs.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Persistence mode. Possible values: `ON`, `OFF`.
        /// </summary>
        [Input("persistenceMode")]
        public Input<string>? PersistenceMode { get; set; }

        /// <summary>
        /// Resources allocated to hosts of the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.MdbRedisClusterResourcesArgs> Resources { get; set; } = null!;

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Redis Cluster mode enabled/disabled. Enables sharding when cluster non-sharded. If cluster is sharded - disabling is not allowed.
        /// </summary>
        [Input("sharded")]
        public Input<bool>? Sharded { get; set; }

        /// <summary>
        /// TLS support mode enabled/disabled.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        public MdbRedisClusterArgs()
        {
        }
        public static new MdbRedisClusterArgs Empty => new MdbRedisClusterArgs();
    }

    public sealed class MdbRedisClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbRedisClusterAccessGetArgs>? Access { get; set; }

        /// <summary>
        /// Announce fqdn instead of ip address.
        /// </summary>
        [Input("announceHostnames")]
        public Input<bool>? AnnounceHostnames { get; set; }

        /// <summary>
        /// Configuration of the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("config")]
        public Input<Inputs.MdbRedisClusterConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// Creation timestamp of the key.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Inhibits deletion of the cluster. Can be either `true` or `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Description of the Redis cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskSizeAutoscaling")]
        public Input<Inputs.MdbRedisClusterDiskSizeAutoscalingGetArgs>? DiskSizeAutoscaling { get; set; }

        /// <summary>
        /// Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("hosts")]
        private InputList<Inputs.MdbRedisClusterHostGetArgs>? _hosts;

        /// <summary>
        /// A host of the Redis cluster. The structure is documented below.
        /// </summary>
        public InputList<Inputs.MdbRedisClusterHostGetArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbRedisClusterHostGetArgs>());
            set => _hosts = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the Redis cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("maintenanceWindow")]
        public Input<Inputs.MdbRedisClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Name of the Redis cluster. Provided by the client when the cluster is created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network, to which the Redis cluster belongs.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Persistence mode. Possible values: `ON`, `OFF`.
        /// </summary>
        [Input("persistenceMode")]
        public Input<string>? PersistenceMode { get; set; }

        /// <summary>
        /// Resources allocated to hosts of the Redis cluster. The structure is documented below.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.MdbRedisClusterResourcesGetArgs>? Resources { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A set of ids of security groups assigned to hosts of the cluster.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Redis Cluster mode enabled/disabled. Enables sharding when cluster non-sharded. If cluster is sharded - disabling is not allowed.
        /// </summary>
        [Input("sharded")]
        public Input<bool>? Sharded { get; set; }

        /// <summary>
        /// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TLS support mode enabled/disabled.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

        public MdbRedisClusterState()
        {
        }
        public static new MdbRedisClusterState Empty => new MdbRedisClusterState();
    }
}
