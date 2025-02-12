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

// Manages a SQLServer cluster within the Yandex Cloud. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-sqlserver/).
//
// Please read [Pricing for Managed Service for SQL Server](https://cloud.yandex.com/docs/managed-sqlserver/pricing#prices) before using SQLServer cluster.
//
// ## Example Usage
//
// {{ tffile "examples/mdb_sqlserver_cluster/r_mdb_sqlserver_cluster_1.tf" }}
//
// ## SQLServer config
//
// If not specified `sqlserverConfig` then does not make any changes.
//
// * maxDegreeOfParallelism - Limits the number of processors to use in parallel plan execution per task. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-max-degree-of-parallelism-server-configuration-option?view=sql-server-2016).
//
// * costThresholdForParallelism - Specifies the threshold at which SQL Server creates and runs parallel plans for queries. SQL Server creates and runs a parallel plan for a query only when the estimated cost to run a serial plan for the same query is higher than the value of the option. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-cost-threshold-for-parallelism-server-configuration-option?view=sql-server-2016).
//
// * auditLevel - Describes how to configure login auditing to monitor SQL Server Database Engine login activity. Possible values:
//   - 0 — do not log login attempts,˚√
//   - 1 — log only failed login attempts,
//   - 2 — log only successful login attempts (not recommended),
//   - 3 — log all login attempts (not recommended). See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/ssms/configure-login-auditing-sql-server-management-studio?view=sql-server-2016).
//
// * fillFactorPercent - Manages the fill factor server configuration option. When an index is created or rebuilt the fill factor determines the percentage of space on each index leaf-level page to be filled with data, reserving the rest as free space for future growth. Values 0 and 100 mean full page usage (no space reserved). See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-fill-factor-server-configuration-option?view=sql-server-2016).
// * optimizeForAdHocWorkloads - Determines whether plans should be cached only after second execution. Allows to avoid SQL cache bloat because of single-use plans. See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/optimize-for-ad-hoc-workloads-server-configuration-option?view=sql-server-2016).
//
// ## Import
//
// A cluster can be imported using the `id` of the resource, e.g.
//
// ```sh
// $ pulumi import yandex:index/mdbSqlserverCluster:MdbSqlserverCluster foo cluster_id
// ```
type MdbSqlserverCluster struct {
	pulumi.CustomResourceState

	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlserverClusterBackupWindowStartOutput `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlserverClusterDatabaseArrayOutput `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds pulumi.StringArrayOutput `pulumi:"hostGroupIds"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlserverClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the SQLServer cluster. Provided by the client when the cluster is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlserverClusterResourcesOutput `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
	Sqlcollation pulumi.StringOutput `pulumi:"sqlcollation"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapOutput `pulumi:"sqlserverConfig"`
	// Status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlserverClusterUserArrayOutput `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewMdbSqlserverCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbSqlserverCluster(ctx *pulumi.Context,
	name string, args *MdbSqlserverClusterArgs, opts ...pulumi.ResourceOption) (*MdbSqlserverCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Databases == nil {
		return nil, errors.New("invalid value for required argument 'Databases'")
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
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MdbSqlserverCluster
	err := ctx.RegisterResource("yandex:index/mdbSqlserverCluster:MdbSqlserverCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbSqlserverCluster gets an existing MdbSqlserverCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbSqlserverCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbSqlserverClusterState, opts ...pulumi.ResourceOption) (*MdbSqlserverCluster, error) {
	var resource MdbSqlserverCluster
	err := ctx.ReadResource("yandex:index/mdbSqlserverCluster:MdbSqlserverCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbSqlserverCluster resources.
type mdbSqlserverClusterState struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbSqlserverClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Creation timestamp of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases []MdbSqlserverClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment *string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds []string `pulumi:"hostGroupIds"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts []MdbSqlserverClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels map[string]string `pulumi:"labels"`
	// Name of the SQLServer cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId *string `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources *MdbSqlserverClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
	Sqlcollation *string `pulumi:"sqlcollation"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users []MdbSqlserverClusterUser `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version *string `pulumi:"version"`
}

type MdbSqlserverClusterState struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlserverClusterBackupWindowStartPtrInput
	// Creation timestamp of the cluster.
	CreatedAt pulumi.StringPtrInput
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlserverClusterDatabaseArrayInput
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringPtrInput
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds pulumi.StringArrayInput
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlserverClusterHostArrayInput
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapInput
	// Name of the SQLServer cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringPtrInput
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlserverClusterResourcesPtrInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
	Sqlcollation pulumi.StringPtrInput
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapInput
	// Status of the cluster.
	Status pulumi.StringPtrInput
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlserverClusterUserArrayInput
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringPtrInput
}

func (MdbSqlserverClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbSqlserverClusterState)(nil)).Elem()
}

type mdbSqlserverClusterArgs struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart *MdbSqlserverClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// A database of the SQLServer cluster. The structure is documented below.
	Databases []MdbSqlserverClusterDatabase `pulumi:"databases"`
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Description of the SQLServer cluster.
	Description *string `pulumi:"description"`
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment string `pulumi:"environment"`
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds []string `pulumi:"hostGroupIds"`
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts []MdbSqlserverClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels map[string]string `pulumi:"labels"`
	// Name of the SQLServer cluster. Provided by the client when the cluster is created.
	Name *string `pulumi:"name"`
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId string `pulumi:"networkId"`
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlserverClusterResources `pulumi:"resources"`
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
	Sqlcollation *string `pulumi:"sqlcollation"`
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig map[string]string `pulumi:"sqlserverConfig"`
	// A user of the SQLServer cluster. The structure is documented below.
	Users []MdbSqlserverClusterUser `pulumi:"users"`
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a MdbSqlserverCluster resource.
type MdbSqlserverClusterArgs struct {
	// Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart MdbSqlserverClusterBackupWindowStartPtrInput
	// A database of the SQLServer cluster. The structure is documented below.
	Databases MdbSqlserverClusterDatabaseArrayInput
	// Inhibits deletion of the cluster. Can be either `true` or `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Description of the SQLServer cluster.
	Description pulumi.StringPtrInput
	// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringInput
	// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// A list of IDs of the host groups hosting VMs of the cluster.
	HostGroupIds pulumi.StringArrayInput
	// A host of the SQLServer cluster. The structure is documented below.
	Hosts MdbSqlserverClusterHostArrayInput
	// A set of key/value label pairs to assign to the SQLServer cluster.
	Labels pulumi.StringMapInput
	// Name of the SQLServer cluster. Provided by the client when the cluster is created.
	Name pulumi.StringPtrInput
	// ID of the network, to which the SQLServer cluster uses.
	NetworkId pulumi.StringInput
	// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
	Resources MdbSqlserverClusterResourcesInput
	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
	Sqlcollation pulumi.StringPtrInput
	// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
	SqlserverConfig pulumi.StringMapInput
	// A user of the SQLServer cluster. The structure is documented below.
	Users MdbSqlserverClusterUserArrayInput
	// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
	Version pulumi.StringInput
}

func (MdbSqlserverClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbSqlserverClusterArgs)(nil)).Elem()
}

type MdbSqlserverClusterInput interface {
	pulumi.Input

	ToMdbSqlserverClusterOutput() MdbSqlserverClusterOutput
	ToMdbSqlserverClusterOutputWithContext(ctx context.Context) MdbSqlserverClusterOutput
}

func (*MdbSqlserverCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbSqlserverCluster)(nil)).Elem()
}

func (i *MdbSqlserverCluster) ToMdbSqlserverClusterOutput() MdbSqlserverClusterOutput {
	return i.ToMdbSqlserverClusterOutputWithContext(context.Background())
}

func (i *MdbSqlserverCluster) ToMdbSqlserverClusterOutputWithContext(ctx context.Context) MdbSqlserverClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlserverClusterOutput)
}

// MdbSqlserverClusterArrayInput is an input type that accepts MdbSqlserverClusterArray and MdbSqlserverClusterArrayOutput values.
// You can construct a concrete instance of `MdbSqlserverClusterArrayInput` via:
//
//	MdbSqlserverClusterArray{ MdbSqlserverClusterArgs{...} }
type MdbSqlserverClusterArrayInput interface {
	pulumi.Input

	ToMdbSqlserverClusterArrayOutput() MdbSqlserverClusterArrayOutput
	ToMdbSqlserverClusterArrayOutputWithContext(context.Context) MdbSqlserverClusterArrayOutput
}

type MdbSqlserverClusterArray []MdbSqlserverClusterInput

func (MdbSqlserverClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbSqlserverCluster)(nil)).Elem()
}

func (i MdbSqlserverClusterArray) ToMdbSqlserverClusterArrayOutput() MdbSqlserverClusterArrayOutput {
	return i.ToMdbSqlserverClusterArrayOutputWithContext(context.Background())
}

func (i MdbSqlserverClusterArray) ToMdbSqlserverClusterArrayOutputWithContext(ctx context.Context) MdbSqlserverClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlserverClusterArrayOutput)
}

// MdbSqlserverClusterMapInput is an input type that accepts MdbSqlserverClusterMap and MdbSqlserverClusterMapOutput values.
// You can construct a concrete instance of `MdbSqlserverClusterMapInput` via:
//
//	MdbSqlserverClusterMap{ "key": MdbSqlserverClusterArgs{...} }
type MdbSqlserverClusterMapInput interface {
	pulumi.Input

	ToMdbSqlserverClusterMapOutput() MdbSqlserverClusterMapOutput
	ToMdbSqlserverClusterMapOutputWithContext(context.Context) MdbSqlserverClusterMapOutput
}

type MdbSqlserverClusterMap map[string]MdbSqlserverClusterInput

func (MdbSqlserverClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbSqlserverCluster)(nil)).Elem()
}

func (i MdbSqlserverClusterMap) ToMdbSqlserverClusterMapOutput() MdbSqlserverClusterMapOutput {
	return i.ToMdbSqlserverClusterMapOutputWithContext(context.Background())
}

func (i MdbSqlserverClusterMap) ToMdbSqlserverClusterMapOutputWithContext(ctx context.Context) MdbSqlserverClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbSqlserverClusterMapOutput)
}

type MdbSqlserverClusterOutput struct{ *pulumi.OutputState }

func (MdbSqlserverClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbSqlserverCluster)(nil)).Elem()
}

func (o MdbSqlserverClusterOutput) ToMdbSqlserverClusterOutput() MdbSqlserverClusterOutput {
	return o
}

func (o MdbSqlserverClusterOutput) ToMdbSqlserverClusterOutputWithContext(ctx context.Context) MdbSqlserverClusterOutput {
	return o
}

// Time to start the daily backup, in the UTC. The structure is documented below.
func (o MdbSqlserverClusterOutput) BackupWindowStart() MdbSqlserverClusterBackupWindowStartOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) MdbSqlserverClusterBackupWindowStartOutput { return v.BackupWindowStart }).(MdbSqlserverClusterBackupWindowStartOutput)
}

// Creation timestamp of the cluster.
func (o MdbSqlserverClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A database of the SQLServer cluster. The structure is documented below.
func (o MdbSqlserverClusterOutput) Databases() MdbSqlserverClusterDatabaseArrayOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) MdbSqlserverClusterDatabaseArrayOutput { return v.Databases }).(MdbSqlserverClusterDatabaseArrayOutput)
}

// Inhibits deletion of the cluster. Can be either `true` or `false`.
func (o MdbSqlserverClusterOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Description of the SQLServer cluster.
func (o MdbSqlserverClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Deployment environment of the SQLServer cluster. (PRODUCTION, PRESTABLE)
func (o MdbSqlserverClusterOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
func (o MdbSqlserverClusterOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster.
func (o MdbSqlserverClusterOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// A list of IDs of the host groups hosting VMs of the cluster.
func (o MdbSqlserverClusterOutput) HostGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringArrayOutput { return v.HostGroupIds }).(pulumi.StringArrayOutput)
}

// A host of the SQLServer cluster. The structure is documented below.
func (o MdbSqlserverClusterOutput) Hosts() MdbSqlserverClusterHostArrayOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) MdbSqlserverClusterHostArrayOutput { return v.Hosts }).(MdbSqlserverClusterHostArrayOutput)
}

// A set of key/value label pairs to assign to the SQLServer cluster.
func (o MdbSqlserverClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the SQLServer cluster. Provided by the client when the cluster is created.
func (o MdbSqlserverClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network, to which the SQLServer cluster uses.
func (o MdbSqlserverClusterOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Resources allocated to hosts of the SQLServer cluster. The structure is documented below.
func (o MdbSqlserverClusterOutput) Resources() MdbSqlserverClusterResourcesOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) MdbSqlserverClusterResourcesOutput { return v.Resources }).(MdbSqlserverClusterResourcesOutput)
}

// A set of ids of security groups assigned to hosts of the cluster.
func (o MdbSqlserverClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// SQL Collation cluster will be created with. This attribute cannot be changed when cluster is created!
func (o MdbSqlserverClusterOutput) Sqlcollation() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Sqlcollation }).(pulumi.StringOutput)
}

// SQLServer cluster config. Detail info in "SQLServer config" section (documented below).
func (o MdbSqlserverClusterOutput) SqlserverConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringMapOutput { return v.SqlserverConfig }).(pulumi.StringMapOutput)
}

// Status of the cluster.
func (o MdbSqlserverClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A user of the SQLServer cluster. The structure is documented below.
func (o MdbSqlserverClusterOutput) Users() MdbSqlserverClusterUserArrayOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) MdbSqlserverClusterUserArrayOutput { return v.Users }).(MdbSqlserverClusterUserArrayOutput)
}

// Version of the SQLServer cluster. (2016sp2std, 2016sp2ent)
func (o MdbSqlserverClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbSqlserverCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type MdbSqlserverClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbSqlserverClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbSqlserverCluster)(nil)).Elem()
}

func (o MdbSqlserverClusterArrayOutput) ToMdbSqlserverClusterArrayOutput() MdbSqlserverClusterArrayOutput {
	return o
}

func (o MdbSqlserverClusterArrayOutput) ToMdbSqlserverClusterArrayOutputWithContext(ctx context.Context) MdbSqlserverClusterArrayOutput {
	return o
}

func (o MdbSqlserverClusterArrayOutput) Index(i pulumi.IntInput) MdbSqlserverClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbSqlserverCluster {
		return vs[0].([]*MdbSqlserverCluster)[vs[1].(int)]
	}).(MdbSqlserverClusterOutput)
}

type MdbSqlserverClusterMapOutput struct{ *pulumi.OutputState }

func (MdbSqlserverClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbSqlserverCluster)(nil)).Elem()
}

func (o MdbSqlserverClusterMapOutput) ToMdbSqlserverClusterMapOutput() MdbSqlserverClusterMapOutput {
	return o
}

func (o MdbSqlserverClusterMapOutput) ToMdbSqlserverClusterMapOutputWithContext(ctx context.Context) MdbSqlserverClusterMapOutput {
	return o
}

func (o MdbSqlserverClusterMapOutput) MapIndex(k pulumi.StringInput) MdbSqlserverClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbSqlserverCluster {
		return vs[0].(map[string]*MdbSqlserverCluster)[vs[1].(string)]
	}).(MdbSqlserverClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbSqlserverClusterInput)(nil)).Elem(), &MdbSqlserverCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbSqlserverClusterArrayInput)(nil)).Elem(), MdbSqlserverClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbSqlserverClusterMapInput)(nil)).Elem(), MdbSqlserverClusterMap{})
	pulumi.RegisterOutputType(MdbSqlserverClusterOutput{})
	pulumi.RegisterOutputType(MdbSqlserverClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbSqlserverClusterMapOutput{})
}
