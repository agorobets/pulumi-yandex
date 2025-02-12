// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Managed ClickHouse cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).
//
// ## Example Usage
//
// {{ tffile "examples/mdb_clickhouse_cluster/d_mdb_clickhouse_cluster_1.tf" }}
func LookupMdbClickhouseCluster(ctx *pulumi.Context, args *LookupMdbClickhouseClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbClickhouseClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMdbClickhouseClusterResult
	err := ctx.Invoke("yandex:index/getMdbClickhouseCluster:getMdbClickhouseCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterArgs struct {
	// Access policy to the ClickHouse cluster. The structure is documented below.
	Access        *GetMdbClickhouseClusterAccess `pulumi:"access"`
	AdminPassword *string                        `pulumi:"adminPassword"`
	// The period in days during which backups are stored.
	BackupRetainPeriodDays *int `pulumi:"backupRetainPeriodDays"`
	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart *GetMdbClickhouseClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Configuration of the ClickHouse subcluster. The structure is documented below.
	Clickhouse   *GetMdbClickhouseClusterClickhouse   `pulumi:"clickhouse"`
	CloudStorage *GetMdbClickhouseClusterCloudStorage `pulumi:"cloudStorage"`
	// The ID of the ClickHouse cluster.
	ClusterId            *string `pulumi:"clusterId"`
	CopySchemaOnNewHosts *bool   `pulumi:"copySchemaOnNewHosts"`
	// Creation timestamp of the key.
	CreatedAt *string `pulumi:"createdAt"`
	// A database of the ClickHouse cluster. The structure is documented below.
	Databases          []GetMdbClickhouseClusterDatabase `pulumi:"databases"`
	DeletionProtection *bool                             `pulumi:"deletionProtection"`
	// Description of the shard group.
	Description *string `pulumi:"description"`
	// Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's used ZooKeeper with placement on separate hosts.
	EmbeddedKeeper *bool `pulumi:"embeddedKeeper"`
	// Deployment environment of the ClickHouse cluster.
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A set of protobuf or cap'n proto format schemas. The structure is documented below.
	FormatSchemas []GetMdbClickhouseClusterFormatSchema `pulumi:"formatSchemas"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A host of the ClickHouse cluster. The structure is documented below.
	Hosts []GetMdbClickhouseClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the ClickHouse cluster.
	Labels            map[string]string                         `pulumi:"labels"`
	MaintenanceWindow *GetMdbClickhouseClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// A group of machine learning models. The structure is documented below.
	MlModels []GetMdbClickhouseClusterMlModel `pulumi:"mlModels"`
	// The name of the ClickHouse cluster.
	//
	// > Either `clusterId` or `name` should be specified.
	Name *string `pulumi:"name"`
	// ID of the network, to which the ClickHouse cluster belongs.
	NetworkId *string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	ServiceAccountId *string  `pulumi:"serviceAccountId"`
	// A group of clickhouse shards. The structure is documented below.
	ShardGroups []GetMdbClickhouseClusterShardGroup `pulumi:"shardGroups"`
	// A shard of the ClickHouse cluster. The structure is documented below.
	Shards []GetMdbClickhouseClusterShard `pulumi:"shards"`
	// Grants `admin` user database management permission.
	SqlDatabaseManagement *bool `pulumi:"sqlDatabaseManagement"`
	// Enables `admin` user with user management permission.
	SqlUserManagement *bool `pulumi:"sqlUserManagement"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// A user of the ClickHouse cluster. The structure is documented below.
	Users   []GetMdbClickhouseClusterUser `pulumi:"users"`
	Version *string                       `pulumi:"version"`
	// Configuration of the ZooKeeper subcluster. The structure is documented below.
	Zookeeper *GetMdbClickhouseClusterZookeeper `pulumi:"zookeeper"`
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResult struct {
	// Access policy to the ClickHouse cluster. The structure is documented below.
	Access        GetMdbClickhouseClusterAccess `pulumi:"access"`
	AdminPassword *string                       `pulumi:"adminPassword"`
	// The period in days during which backups are stored.
	BackupRetainPeriodDays *int `pulumi:"backupRetainPeriodDays"`
	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart GetMdbClickhouseClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Configuration of the ClickHouse subcluster. The structure is documented below.
	Clickhouse           GetMdbClickhouseClusterClickhouse   `pulumi:"clickhouse"`
	CloudStorage         GetMdbClickhouseClusterCloudStorage `pulumi:"cloudStorage"`
	ClusterId            string                              `pulumi:"clusterId"`
	CopySchemaOnNewHosts *bool                               `pulumi:"copySchemaOnNewHosts"`
	// Creation timestamp of the key.
	CreatedAt string `pulumi:"createdAt"`
	// A database of the ClickHouse cluster. The structure is documented below.
	Databases          []GetMdbClickhouseClusterDatabase `pulumi:"databases"`
	DeletionProtection bool                              `pulumi:"deletionProtection"`
	// Description of the shard group.
	Description *string `pulumi:"description"`
	// Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's used ZooKeeper with placement on separate hosts.
	EmbeddedKeeper bool `pulumi:"embeddedKeeper"`
	// Deployment environment of the ClickHouse cluster.
	Environment *string `pulumi:"environment"`
	FolderId    string  `pulumi:"folderId"`
	// A set of protobuf or cap'n proto format schemas. The structure is documented below.
	FormatSchemas []GetMdbClickhouseClusterFormatSchema `pulumi:"formatSchemas"`
	// Aggregated health of the cluster.
	Health string `pulumi:"health"`
	// A host of the ClickHouse cluster. The structure is documented below.
	Hosts []GetMdbClickhouseClusterHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of key/value label pairs to assign to the ClickHouse cluster.
	Labels            map[string]string                        `pulumi:"labels"`
	MaintenanceWindow GetMdbClickhouseClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// A group of machine learning models. The structure is documented below.
	MlModels []GetMdbClickhouseClusterMlModel `pulumi:"mlModels"`
	// (Optional) Name for the rule.
	Name *string `pulumi:"name"`
	// ID of the network, to which the ClickHouse cluster belongs.
	NetworkId *string `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	ServiceAccountId *string  `pulumi:"serviceAccountId"`
	// A group of clickhouse shards. The structure is documented below.
	ShardGroups []GetMdbClickhouseClusterShardGroup `pulumi:"shardGroups"`
	// A shard of the ClickHouse cluster. The structure is documented below.
	Shards []GetMdbClickhouseClusterShard `pulumi:"shards"`
	// Grants `admin` user database management permission.
	SqlDatabaseManagement bool `pulumi:"sqlDatabaseManagement"`
	// Enables `admin` user with user management permission.
	SqlUserManagement bool `pulumi:"sqlUserManagement"`
	// Status of the cluster.
	Status string `pulumi:"status"`
	// A user of the ClickHouse cluster. The structure is documented below.
	Users   []GetMdbClickhouseClusterUser `pulumi:"users"`
	Version string                        `pulumi:"version"`
	// Configuration of the ZooKeeper subcluster. The structure is documented below.
	Zookeeper GetMdbClickhouseClusterZookeeper `pulumi:"zookeeper"`
}

func LookupMdbClickhouseClusterOutput(ctx *pulumi.Context, args LookupMdbClickhouseClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbClickhouseClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMdbClickhouseClusterResultOutput, error) {
			args := v.(LookupMdbClickhouseClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getMdbClickhouseCluster:getMdbClickhouseCluster", args, LookupMdbClickhouseClusterResultOutput{}, options).(LookupMdbClickhouseClusterResultOutput), nil
		}).(LookupMdbClickhouseClusterResultOutput)
}

// A collection of arguments for invoking getMdbClickhouseCluster.
type LookupMdbClickhouseClusterOutputArgs struct {
	// Access policy to the ClickHouse cluster. The structure is documented below.
	Access        GetMdbClickhouseClusterAccessPtrInput `pulumi:"access"`
	AdminPassword pulumi.StringPtrInput                 `pulumi:"adminPassword"`
	// The period in days during which backups are stored.
	BackupRetainPeriodDays pulumi.IntPtrInput `pulumi:"backupRetainPeriodDays"`
	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart GetMdbClickhouseClusterBackupWindowStartPtrInput `pulumi:"backupWindowStart"`
	// Configuration of the ClickHouse subcluster. The structure is documented below.
	Clickhouse   GetMdbClickhouseClusterClickhousePtrInput   `pulumi:"clickhouse"`
	CloudStorage GetMdbClickhouseClusterCloudStoragePtrInput `pulumi:"cloudStorage"`
	// The ID of the ClickHouse cluster.
	ClusterId            pulumi.StringPtrInput `pulumi:"clusterId"`
	CopySchemaOnNewHosts pulumi.BoolPtrInput   `pulumi:"copySchemaOnNewHosts"`
	// Creation timestamp of the key.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// A database of the ClickHouse cluster. The structure is documented below.
	Databases          GetMdbClickhouseClusterDatabaseArrayInput `pulumi:"databases"`
	DeletionProtection pulumi.BoolPtrInput                       `pulumi:"deletionProtection"`
	// Description of the shard group.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's used ZooKeeper with placement on separate hosts.
	EmbeddedKeeper pulumi.BoolPtrInput `pulumi:"embeddedKeeper"`
	// Deployment environment of the ClickHouse cluster.
	Environment pulumi.StringPtrInput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// A set of protobuf or cap'n proto format schemas. The structure is documented below.
	FormatSchemas GetMdbClickhouseClusterFormatSchemaArrayInput `pulumi:"formatSchemas"`
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput `pulumi:"health"`
	// A host of the ClickHouse cluster. The structure is documented below.
	Hosts GetMdbClickhouseClusterHostArrayInput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the ClickHouse cluster.
	Labels            pulumi.StringMapInput                            `pulumi:"labels"`
	MaintenanceWindow GetMdbClickhouseClusterMaintenanceWindowPtrInput `pulumi:"maintenanceWindow"`
	// A group of machine learning models. The structure is documented below.
	MlModels GetMdbClickhouseClusterMlModelArrayInput `pulumi:"mlModels"`
	// The name of the ClickHouse cluster.
	//
	// > Either `clusterId` or `name` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the network, to which the ClickHouse cluster belongs.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	ServiceAccountId pulumi.StringPtrInput   `pulumi:"serviceAccountId"`
	// A group of clickhouse shards. The structure is documented below.
	ShardGroups GetMdbClickhouseClusterShardGroupArrayInput `pulumi:"shardGroups"`
	// A shard of the ClickHouse cluster. The structure is documented below.
	Shards GetMdbClickhouseClusterShardArrayInput `pulumi:"shards"`
	// Grants `admin` user database management permission.
	SqlDatabaseManagement pulumi.BoolPtrInput `pulumi:"sqlDatabaseManagement"`
	// Enables `admin` user with user management permission.
	SqlUserManagement pulumi.BoolPtrInput `pulumi:"sqlUserManagement"`
	// Status of the cluster.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// A user of the ClickHouse cluster. The structure is documented below.
	Users   GetMdbClickhouseClusterUserArrayInput `pulumi:"users"`
	Version pulumi.StringPtrInput                 `pulumi:"version"`
	// Configuration of the ZooKeeper subcluster. The structure is documented below.
	Zookeeper GetMdbClickhouseClusterZookeeperPtrInput `pulumi:"zookeeper"`
}

func (LookupMdbClickhouseClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbClickhouseCluster.
type LookupMdbClickhouseClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbClickhouseClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbClickhouseClusterResult)(nil)).Elem()
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutput() LookupMdbClickhouseClusterResultOutput {
	return o
}

func (o LookupMdbClickhouseClusterResultOutput) ToLookupMdbClickhouseClusterResultOutputWithContext(ctx context.Context) LookupMdbClickhouseClusterResultOutput {
	return o
}

// Access policy to the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Access() GetMdbClickhouseClusterAccessOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterAccess { return v.Access }).(GetMdbClickhouseClusterAccessOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

// The period in days during which backups are stored.
func (o LookupMdbClickhouseClusterResultOutput) BackupRetainPeriodDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *int { return v.BackupRetainPeriodDays }).(pulumi.IntPtrOutput)
}

// Time to start the daily backup, in the UTC timezone. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) BackupWindowStart() GetMdbClickhouseClusterBackupWindowStartOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterBackupWindowStart {
		return v.BackupWindowStart
	}).(GetMdbClickhouseClusterBackupWindowStartOutput)
}

// Configuration of the ClickHouse subcluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Clickhouse() GetMdbClickhouseClusterClickhouseOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterClickhouse { return v.Clickhouse }).(GetMdbClickhouseClusterClickhouseOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) CloudStorage() GetMdbClickhouseClusterCloudStorageOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterCloudStorage { return v.CloudStorage }).(GetMdbClickhouseClusterCloudStorageOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) CopySchemaOnNewHosts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *bool { return v.CopySchemaOnNewHosts }).(pulumi.BoolPtrOutput)
}

// Creation timestamp of the key.
func (o LookupMdbClickhouseClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Databases() GetMdbClickhouseClusterDatabaseArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterDatabase { return v.Databases }).(GetMdbClickhouseClusterDatabaseArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the shard group.
func (o LookupMdbClickhouseClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's used ZooKeeper with placement on separate hosts.
func (o LookupMdbClickhouseClusterResultOutput) EmbeddedKeeper() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.EmbeddedKeeper }).(pulumi.BoolOutput)
}

// Deployment environment of the ClickHouse cluster.
func (o LookupMdbClickhouseClusterResultOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// A set of protobuf or cap'n proto format schemas. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) FormatSchemas() GetMdbClickhouseClusterFormatSchemaArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterFormatSchema { return v.FormatSchemas }).(GetMdbClickhouseClusterFormatSchemaArrayOutput)
}

// Aggregated health of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// A host of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Hosts() GetMdbClickhouseClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterHost { return v.Hosts }).(GetMdbClickhouseClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbClickhouseClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the ClickHouse cluster.
func (o LookupMdbClickhouseClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) MaintenanceWindow() GetMdbClickhouseClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterMaintenanceWindow {
		return v.MaintenanceWindow
	}).(GetMdbClickhouseClusterMaintenanceWindowOutput)
}

// A group of machine learning models. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) MlModels() GetMdbClickhouseClusterMlModelArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterMlModel { return v.MlModels }).(GetMdbClickhouseClusterMlModelArrayOutput)
}

// (Optional) Name for the rule.
func (o LookupMdbClickhouseClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// ID of the network, to which the ClickHouse cluster belongs.
func (o LookupMdbClickhouseClusterResultOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) ServiceAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) *string { return v.ServiceAccountId }).(pulumi.StringPtrOutput)
}

// A group of clickhouse shards. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) ShardGroups() GetMdbClickhouseClusterShardGroupArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterShardGroup { return v.ShardGroups }).(GetMdbClickhouseClusterShardGroupArrayOutput)
}

// A shard of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Shards() GetMdbClickhouseClusterShardArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterShard { return v.Shards }).(GetMdbClickhouseClusterShardArrayOutput)
}

// Grants `admin` user database management permission.
func (o LookupMdbClickhouseClusterResultOutput) SqlDatabaseManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlDatabaseManagement }).(pulumi.BoolOutput)
}

// Enables `admin` user with user management permission.
func (o LookupMdbClickhouseClusterResultOutput) SqlUserManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) bool { return v.SqlUserManagement }).(pulumi.BoolOutput)
}

// Status of the cluster.
func (o LookupMdbClickhouseClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// A user of the ClickHouse cluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Users() GetMdbClickhouseClusterUserArrayOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) []GetMdbClickhouseClusterUser { return v.Users }).(GetMdbClickhouseClusterUserArrayOutput)
}

func (o LookupMdbClickhouseClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

// Configuration of the ZooKeeper subcluster. The structure is documented below.
func (o LookupMdbClickhouseClusterResultOutput) Zookeeper() GetMdbClickhouseClusterZookeeperOutput {
	return o.ApplyT(func(v LookupMdbClickhouseClusterResult) GetMdbClickhouseClusterZookeeper { return v.Zookeeper }).(GetMdbClickhouseClusterZookeeperOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbClickhouseClusterResultOutput{})
}
