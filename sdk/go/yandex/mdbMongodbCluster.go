// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"errors"
	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a MongoDB cluster within the Yandex Cloud. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).
//
// ## Example Usage
//
// {{ tffile "examples/mdb_mongodb_cluster/r_mdb_mongodb_cluster_1.tf" }}
//
// ## Import
//
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
// $ pulumi import yandex:index/mdbMongodbCluster:MdbMongodbCluster foo cluster_id
// ```
type MdbMongodbCluster struct {
	pulumi.CustomResourceState

	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig MdbMongodbClusterClusterConfigOutput `pulumi:"clusterConfig"`
	// The ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation timestamp of the key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A database of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases MdbMongodbClusterDatabaseArrayOutput `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the MongoDB cluster.
	Description                   pulumi.StringPtrOutput                               `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   MdbMongodbClusterDiskSizeAutoscalingMongocfgOutput   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     MdbMongodbClusterDiskSizeAutoscalingMongodOutput     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra MdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     MdbMongodbClusterDiskSizeAutoscalingMongosOutput     `pulumi:"diskSizeAutoscalingMongos"`
	// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Health pulumi.StringOutput `pulumi:"health"`
	// A host of the MongoDB cluster. The structure is documented below.
	Hosts MdbMongodbClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow MdbMongodbClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the MongoDB cluster belongs.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources MdbMongodbClusterResourcesPtrOutput `pulumi:"resources"`
	// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongocfg MdbMongodbClusterResourcesMongocfgPtrOutput `pulumi:"resourcesMongocfg"`
	// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongod MdbMongodbClusterResourcesMongodPtrOutput `pulumi:"resourcesMongod"`
	// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongoinfra MdbMongodbClusterResourcesMongoinfraPtrOutput `pulumi:"resourcesMongoinfra"`
	// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongos MdbMongodbClusterResourcesMongosPtrOutput `pulumi:"resourcesMongos"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMongodbClusterRestorePtrOutput `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// MongoDB Cluster mode enabled/disabled.
	Sharded pulumi.BoolOutput `pulumi:"sharded"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Status pulumi.StringOutput `pulumi:"status"`
	// A user of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users MdbMongodbClusterUserArrayOutput `pulumi:"users"`
}

// NewMdbMongodbCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbMongodbCluster(ctx *pulumi.Context,
	name string, args *MdbMongodbClusterArgs, opts ...pulumi.ResourceOption) (*MdbMongodbCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterConfig == nil {
		return nil, errors.New("invalid value for required argument 'ClusterConfig'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Hosts == nil {
		return nil, errors.New("invalid value for required argument 'Hosts'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MdbMongodbCluster
	err := ctx.RegisterResource("yandex:index/mdbMongodbCluster:MdbMongodbCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbMongodbCluster gets an existing MdbMongodbCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbMongodbCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbMongodbClusterState, opts ...pulumi.ResourceOption) (*MdbMongodbCluster, error) {
	var resource MdbMongodbCluster
	err := ctx.ReadResource("yandex:index/mdbMongodbCluster:MdbMongodbCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbMongodbCluster resources.
type mdbMongodbClusterState struct {
	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig *MdbMongodbClusterClusterConfig `pulumi:"clusterConfig"`
	// The ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Creation timestamp of the key.
	CreatedAt *string `pulumi:"createdAt"`
	// A database of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases []MdbMongodbClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the MongoDB cluster.
	Description                   *string                                         `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   *MdbMongodbClusterDiskSizeAutoscalingMongocfg   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     *MdbMongodbClusterDiskSizeAutoscalingMongod     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra *MdbMongodbClusterDiskSizeAutoscalingMongoinfra `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     *MdbMongodbClusterDiskSizeAutoscalingMongos     `pulumi:"diskSizeAutoscalingMongos"`
	// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Health *string `pulumi:"health"`
	// A host of the MongoDB cluster. The structure is documented below.
	Hosts []MdbMongodbClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels map[string]string `pulumi:"labels"`
	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow *MdbMongodbClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the MongoDB cluster belongs.
	NetworkId *string `pulumi:"networkId"`
	// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources *MdbMongodbClusterResources `pulumi:"resources"`
	// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongocfg *MdbMongodbClusterResourcesMongocfg `pulumi:"resourcesMongocfg"`
	// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongod *MdbMongodbClusterResourcesMongod `pulumi:"resourcesMongod"`
	// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongoinfra *MdbMongodbClusterResourcesMongoinfra `pulumi:"resourcesMongoinfra"`
	// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongos *MdbMongodbClusterResourcesMongos `pulumi:"resourcesMongos"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore *MdbMongodbClusterRestore `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// MongoDB Cluster mode enabled/disabled.
	Sharded *bool `pulumi:"sharded"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Status *string `pulumi:"status"`
	// A user of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users []MdbMongodbClusterUser `pulumi:"users"`
}

type MdbMongodbClusterState struct {
	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig MdbMongodbClusterClusterConfigPtrInput
	// The ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Creation timestamp of the key.
	CreatedAt pulumi.StringPtrInput
	// A database of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases MdbMongodbClusterDatabaseArrayInput
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the MongoDB cluster.
	Description                   pulumi.StringPtrInput
	DiskSizeAutoscalingMongocfg   MdbMongodbClusterDiskSizeAutoscalingMongocfgPtrInput
	DiskSizeAutoscalingMongod     MdbMongodbClusterDiskSizeAutoscalingMongodPtrInput
	DiskSizeAutoscalingMongoinfra MdbMongodbClusterDiskSizeAutoscalingMongoinfraPtrInput
	DiskSizeAutoscalingMongos     MdbMongodbClusterDiskSizeAutoscalingMongosPtrInput
	// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Health pulumi.StringPtrInput
	// A host of the MongoDB cluster. The structure is documented below.
	Hosts MdbMongodbClusterHostArrayInput
	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels pulumi.StringMapInput
	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow MdbMongodbClusterMaintenanceWindowPtrInput
	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the MongoDB cluster belongs.
	NetworkId pulumi.StringPtrInput
	// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources MdbMongodbClusterResourcesPtrInput
	// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongocfg MdbMongodbClusterResourcesMongocfgPtrInput
	// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongod MdbMongodbClusterResourcesMongodPtrInput
	// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongoinfra MdbMongodbClusterResourcesMongoinfraPtrInput
	// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongos MdbMongodbClusterResourcesMongosPtrInput
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMongodbClusterRestorePtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// MongoDB Cluster mode enabled/disabled.
	Sharded pulumi.BoolPtrInput
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
	Status pulumi.StringPtrInput
	// A user of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users MdbMongodbClusterUserArrayInput
}

func (MdbMongodbClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMongodbClusterState)(nil)).Elem()
}

type mdbMongodbClusterArgs struct {
	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig MdbMongodbClusterClusterConfig `pulumi:"clusterConfig"`
	// The ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// A database of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases []MdbMongodbClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the MongoDB cluster.
	Description                   *string                                         `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   *MdbMongodbClusterDiskSizeAutoscalingMongocfg   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     *MdbMongodbClusterDiskSizeAutoscalingMongod     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra *MdbMongodbClusterDiskSizeAutoscalingMongoinfra `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     *MdbMongodbClusterDiskSizeAutoscalingMongos     `pulumi:"diskSizeAutoscalingMongos"`
	// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A host of the MongoDB cluster. The structure is documented below.
	Hosts []MdbMongodbClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels map[string]string `pulumi:"labels"`
	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow *MdbMongodbClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the MongoDB cluster belongs.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources *MdbMongodbClusterResources `pulumi:"resources"`
	// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongocfg *MdbMongodbClusterResourcesMongocfg `pulumi:"resourcesMongocfg"`
	// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongod *MdbMongodbClusterResourcesMongod `pulumi:"resourcesMongod"`
	// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongoinfra *MdbMongodbClusterResourcesMongoinfra `pulumi:"resourcesMongoinfra"`
	// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongos *MdbMongodbClusterResourcesMongos `pulumi:"resourcesMongos"`
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore *MdbMongodbClusterRestore `pulumi:"restore"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A user of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users []MdbMongodbClusterUser `pulumi:"users"`
}

// The set of arguments for constructing a MdbMongodbCluster resource.
type MdbMongodbClusterArgs struct {
	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig MdbMongodbClusterClusterConfigInput
	// The ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// A database of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases MdbMongodbClusterDatabaseArrayInput
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the MongoDB cluster.
	Description                   pulumi.StringPtrInput
	DiskSizeAutoscalingMongocfg   MdbMongodbClusterDiskSizeAutoscalingMongocfgPtrInput
	DiskSizeAutoscalingMongod     MdbMongodbClusterDiskSizeAutoscalingMongodPtrInput
	DiskSizeAutoscalingMongoinfra MdbMongodbClusterDiskSizeAutoscalingMongoinfraPtrInput
	DiskSizeAutoscalingMongos     MdbMongodbClusterDiskSizeAutoscalingMongosPtrInput
	// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A host of the MongoDB cluster. The structure is documented below.
	Hosts MdbMongodbClusterHostArrayInput
	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels pulumi.StringMapInput
	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow MdbMongodbClusterMaintenanceWindowPtrInput
	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the MongoDB cluster belongs.
	NetworkId pulumi.StringInput
	// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources MdbMongodbClusterResourcesPtrInput
	// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongocfg MdbMongodbClusterResourcesMongocfgPtrInput
	// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongod MdbMongodbClusterResourcesMongodPtrInput
	// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongoinfra MdbMongodbClusterResourcesMongoinfraPtrInput
	// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
	ResourcesMongos MdbMongodbClusterResourcesMongosPtrInput
	// The cluster will be created from the specified backup. The structure is documented below.
	Restore MdbMongodbClusterRestorePtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// A user of the MongoDB cluster. The structure is documented below.
	//
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users MdbMongodbClusterUserArrayInput
}

func (MdbMongodbClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbMongodbClusterArgs)(nil)).Elem()
}

type MdbMongodbClusterInput interface {
	pulumi.Input

	ToMdbMongodbClusterOutput() MdbMongodbClusterOutput
	ToMdbMongodbClusterOutputWithContext(ctx context.Context) MdbMongodbClusterOutput
}

func (*MdbMongodbCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMongodbCluster)(nil)).Elem()
}

func (i *MdbMongodbCluster) ToMdbMongodbClusterOutput() MdbMongodbClusterOutput {
	return i.ToMdbMongodbClusterOutputWithContext(context.Background())
}

func (i *MdbMongodbCluster) ToMdbMongodbClusterOutputWithContext(ctx context.Context) MdbMongodbClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMongodbClusterOutput)
}

// MdbMongodbClusterArrayInput is an input type that accepts MdbMongodbClusterArray and MdbMongodbClusterArrayOutput values.
// You can construct a concrete instance of `MdbMongodbClusterArrayInput` via:
//
//	MdbMongodbClusterArray{ MdbMongodbClusterArgs{...} }
type MdbMongodbClusterArrayInput interface {
	pulumi.Input

	ToMdbMongodbClusterArrayOutput() MdbMongodbClusterArrayOutput
	ToMdbMongodbClusterArrayOutputWithContext(context.Context) MdbMongodbClusterArrayOutput
}

type MdbMongodbClusterArray []MdbMongodbClusterInput

func (MdbMongodbClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMongodbCluster)(nil)).Elem()
}

func (i MdbMongodbClusterArray) ToMdbMongodbClusterArrayOutput() MdbMongodbClusterArrayOutput {
	return i.ToMdbMongodbClusterArrayOutputWithContext(context.Background())
}

func (i MdbMongodbClusterArray) ToMdbMongodbClusterArrayOutputWithContext(ctx context.Context) MdbMongodbClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMongodbClusterArrayOutput)
}

// MdbMongodbClusterMapInput is an input type that accepts MdbMongodbClusterMap and MdbMongodbClusterMapOutput values.
// You can construct a concrete instance of `MdbMongodbClusterMapInput` via:
//
//	MdbMongodbClusterMap{ "key": MdbMongodbClusterArgs{...} }
type MdbMongodbClusterMapInput interface {
	pulumi.Input

	ToMdbMongodbClusterMapOutput() MdbMongodbClusterMapOutput
	ToMdbMongodbClusterMapOutputWithContext(context.Context) MdbMongodbClusterMapOutput
}

type MdbMongodbClusterMap map[string]MdbMongodbClusterInput

func (MdbMongodbClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMongodbCluster)(nil)).Elem()
}

func (i MdbMongodbClusterMap) ToMdbMongodbClusterMapOutput() MdbMongodbClusterMapOutput {
	return i.ToMdbMongodbClusterMapOutputWithContext(context.Background())
}

func (i MdbMongodbClusterMap) ToMdbMongodbClusterMapOutputWithContext(ctx context.Context) MdbMongodbClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbMongodbClusterMapOutput)
}

type MdbMongodbClusterOutput struct{ *pulumi.OutputState }

func (MdbMongodbClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbMongodbCluster)(nil)).Elem()
}

func (o MdbMongodbClusterOutput) ToMdbMongodbClusterOutput() MdbMongodbClusterOutput {
	return o
}

func (o MdbMongodbClusterOutput) ToMdbMongodbClusterOutputWithContext(ctx context.Context) MdbMongodbClusterOutput {
	return o
}

// Configuration of the MongoDB subcluster. The structure is documented below.
func (o MdbMongodbClusterOutput) ClusterConfig() MdbMongodbClusterClusterConfigOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterClusterConfigOutput { return v.ClusterConfig }).(MdbMongodbClusterClusterConfigOutput)
}

// The ID of the cluster.
func (o MdbMongodbClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation timestamp of the key.
func (o MdbMongodbClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the MongoDB cluster. The structure is documented below.
//
// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
func (o MdbMongodbClusterOutput) Databases() MdbMongodbClusterDatabaseArrayOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterDatabaseArrayOutput { return v.Databases }).(MdbMongodbClusterDatabaseArrayOutput)
}

// Inhibits deletion of the cluster. Can be either `true` or `false`.
func (o MdbMongodbClusterOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the MongoDB cluster.
func (o MdbMongodbClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MdbMongodbClusterOutput) DiskSizeAutoscalingMongocfg() MdbMongodbClusterDiskSizeAutoscalingMongocfgOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterDiskSizeAutoscalingMongocfgOutput {
		return v.DiskSizeAutoscalingMongocfg
	}).(MdbMongodbClusterDiskSizeAutoscalingMongocfgOutput)
}

func (o MdbMongodbClusterOutput) DiskSizeAutoscalingMongod() MdbMongodbClusterDiskSizeAutoscalingMongodOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterDiskSizeAutoscalingMongodOutput {
		return v.DiskSizeAutoscalingMongod
	}).(MdbMongodbClusterDiskSizeAutoscalingMongodOutput)
}

func (o MdbMongodbClusterOutput) DiskSizeAutoscalingMongoinfra() MdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput {
		return v.DiskSizeAutoscalingMongoinfra
	}).(MdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput)
}

func (o MdbMongodbClusterOutput) DiskSizeAutoscalingMongos() MdbMongodbClusterDiskSizeAutoscalingMongosOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterDiskSizeAutoscalingMongosOutput {
		return v.DiskSizeAutoscalingMongos
	}).(MdbMongodbClusterDiskSizeAutoscalingMongosOutput)
}

// Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
func (o MdbMongodbClusterOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
func (o MdbMongodbClusterOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see `health` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
func (o MdbMongodbClusterOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// A host of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) Hosts() MdbMongodbClusterHostArrayOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterHostArrayOutput { return v.Hosts }).(MdbMongodbClusterHostArrayOutput)
}

// A set of key/value label pairs to assign to the MongoDB cluster.
func (o MdbMongodbClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Maintenance window settings of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) MaintenanceWindow() MdbMongodbClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterMaintenanceWindowOutput { return v.MaintenanceWindow }).(MdbMongodbClusterMaintenanceWindowOutput)
}

// Name of the MongoDB cluster. Provided by the client when the cluster is created.
func (o MdbMongodbClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the MongoDB cluster belongs.
func (o MdbMongodbClusterOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
//
// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
func (o MdbMongodbClusterOutput) Resources() MdbMongodbClusterResourcesPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterResourcesPtrOutput { return v.Resources }).(MdbMongodbClusterResourcesPtrOutput)
}

// Resources allocated to `mongocfg` hosts of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) ResourcesMongocfg() MdbMongodbClusterResourcesMongocfgPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterResourcesMongocfgPtrOutput { return v.ResourcesMongocfg }).(MdbMongodbClusterResourcesMongocfgPtrOutput)
}

// Resources allocated to `mongod` hosts of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) ResourcesMongod() MdbMongodbClusterResourcesMongodPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterResourcesMongodPtrOutput { return v.ResourcesMongod }).(MdbMongodbClusterResourcesMongodPtrOutput)
}

// Resources allocated to `mongoinfra` hosts of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) ResourcesMongoinfra() MdbMongodbClusterResourcesMongoinfraPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterResourcesMongoinfraPtrOutput { return v.ResourcesMongoinfra }).(MdbMongodbClusterResourcesMongoinfraPtrOutput)
}

// Resources allocated to `mongos` hosts of the MongoDB cluster. The structure is documented below.
func (o MdbMongodbClusterOutput) ResourcesMongos() MdbMongodbClusterResourcesMongosPtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterResourcesMongosPtrOutput { return v.ResourcesMongos }).(MdbMongodbClusterResourcesMongosPtrOutput)
}

// The cluster will be created from the specified backup. The structure is documented below.
func (o MdbMongodbClusterOutput) Restore() MdbMongodbClusterRestorePtrOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterRestorePtrOutput { return v.Restore }).(MdbMongodbClusterRestorePtrOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o MdbMongodbClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// MongoDB Cluster mode enabled/disabled.
func (o MdbMongodbClusterOutput) Sharded() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.BoolOutput { return v.Sharded }).(pulumi.BoolOutput)
}

// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).
func (o MdbMongodbClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A user of the MongoDB cluster. The structure is documented below.
//
// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
func (o MdbMongodbClusterOutput) Users() MdbMongodbClusterUserArrayOutput {
	return o.ApplyT(func(v *MdbMongodbCluster) MdbMongodbClusterUserArrayOutput { return v.Users }).(MdbMongodbClusterUserArrayOutput)
}

type MdbMongodbClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbMongodbClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbMongodbCluster)(nil)).Elem()
}

func (o MdbMongodbClusterArrayOutput) ToMdbMongodbClusterArrayOutput() MdbMongodbClusterArrayOutput {
	return o
}

func (o MdbMongodbClusterArrayOutput) ToMdbMongodbClusterArrayOutputWithContext(ctx context.Context) MdbMongodbClusterArrayOutput {
	return o
}

func (o MdbMongodbClusterArrayOutput) Index(i pulumi.IntInput) MdbMongodbClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbMongodbCluster {
		return vs[0].([]*MdbMongodbCluster)[vs[1].(int)]
	}).(MdbMongodbClusterOutput)
}

type MdbMongodbClusterMapOutput struct{ *pulumi.OutputState }

func (MdbMongodbClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbMongodbCluster)(nil)).Elem()
}

func (o MdbMongodbClusterMapOutput) ToMdbMongodbClusterMapOutput() MdbMongodbClusterMapOutput {
	return o
}

func (o MdbMongodbClusterMapOutput) ToMdbMongodbClusterMapOutputWithContext(ctx context.Context) MdbMongodbClusterMapOutput {
	return o
}

func (o MdbMongodbClusterMapOutput) MapIndex(k pulumi.StringInput) MdbMongodbClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbMongodbCluster {
		return vs[0].(map[string]*MdbMongodbCluster)[vs[1].(string)]
	}).(MdbMongodbClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMongodbClusterInput)(nil)).Elem(), &MdbMongodbCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMongodbClusterArrayInput)(nil)).Elem(), MdbMongodbClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbMongodbClusterMapInput)(nil)).Elem(), MdbMongodbClusterMap{})
	pulumi.RegisterOutputType(MdbMongodbClusterOutput{})
	pulumi.RegisterOutputType(MdbMongodbClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbMongodbClusterMapOutput{})
}
