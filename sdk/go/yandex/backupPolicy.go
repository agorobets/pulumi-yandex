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

// Allows management of [Yandex Cloud Backup Policy](https://yandex.cloud/docs/backup/concepts/policy).
//
// > Cloud Backup Provider must be activated in order to manipulate with policies. Active it either by UI Console or by `yc` command.
//
// ## Example Usage
//
// {{ tffile "examples/backup_policy/r_backup_policy_1.tf" }}
//
// For the full policy attributes, take a look at the following example:
//
// {{ tffile "examples/backup_policy/r_backup_policy_2.tf" }}
//
// ## Defined types
//
// ### interval_type
//
// A string type, that accepts values in the format of: `number` + `time type`, where `time type` might be:
//
// - `s` — seconds
// - `m` — minutes
// - `h` — hours
// - `d` — days
// - `w` — weekdays
// - `M` — months
//
// Example of interval value: `"5m", "10d", "2M", "5w"`
//
// ### dayType
//
// A string type, that accepts the following values: `"ALWAYS_INCREMENTAL"`, `"ALWAYS_FULL"`, `"WEEKLY_FULL_DAILY_INCREMENTAL"`, `'WEEKLY_INCREMENTAL"`.
//
// ### backupSetType
//
// `"TYPE_AUTO"`, `"TYPE_FULL"`, `"TYPE_INCREMENTAL"`, `'TYPE_DIFFERENTIAL"`.
type BackupPolicy struct {
	pulumi.CustomResourceState

	// — The name of generated archives.
	ArchiveName pulumi.StringPtrOutput `pulumi:"archiveName"`
	// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
	Cbt pulumi.StringPtrOutput `pulumi:"cbt"`
	// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
	Compression       pulumi.StringPtrOutput `pulumi:"compression"`
	CreatedAt         pulumi.StringOutput    `pulumi:"createdAt"`
	Enabled           pulumi.BoolOutput      `pulumi:"enabled"`
	FastBackupEnabled pulumi.BoolPtrOutput   `pulumi:"fastBackupEnabled"`
	FolderId          pulumi.StringOutput    `pulumi:"folderId"`
	// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled pulumi.BoolPtrOutput `pulumi:"multiVolumeSnapshottingEnabled"`
	// — Name of the policy
	Name pulumi.StringOutput `pulumi:"name"`
	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled pulumi.BoolPtrOutput `pulumi:"performanceWindowEnabled"`
	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings pulumi.BoolPtrOutput `pulumi:"preserveFileSecuritySettings"`
	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled pulumi.BoolPtrOutput `pulumi:"quiesceSnapshottingEnabled"`
	// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
	Reattempts BackupPolicyReattemptsOutput `pulumi:"reattempts"`
	// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
	Retention BackupPolicyRetentionOutput `pulumi:"retention"`
	// — Schedule settings for creating backups on the host.
	Scheduling BackupPolicySchedulingOutput `pulumi:"scheduling"`
	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled pulumi.BoolPtrOutput `pulumi:"silentModeEnabled"`
	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes pulumi.StringPtrOutput `pulumi:"splittingBytes"`
	UpdatedAt      pulumi.StringOutput    `pulumi:"updatedAt"`
	// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
	VmSnapshotReattempts BackupPolicyVmSnapshotReattemptsOutput `pulumi:"vmSnapshotReattempts"`
	// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
	VssProvider pulumi.StringPtrOutput `pulumi:"vssProvider"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Reattempts == nil {
		return nil, errors.New("invalid value for required argument 'Reattempts'")
	}
	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	if args.Scheduling == nil {
		return nil, errors.New("invalid value for required argument 'Scheduling'")
	}
	if args.VmSnapshotReattempts == nil {
		return nil, errors.New("invalid value for required argument 'VmSnapshotReattempts'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("yandex:index/backupPolicy:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("yandex:index/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// — The name of generated archives.
	ArchiveName *string `pulumi:"archiveName"`
	// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
	Cbt *string `pulumi:"cbt"`
	// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
	Compression       *string `pulumi:"compression"`
	CreatedAt         *string `pulumi:"createdAt"`
	Enabled           *bool   `pulumi:"enabled"`
	FastBackupEnabled *bool   `pulumi:"fastBackupEnabled"`
	FolderId          *string `pulumi:"folderId"`
	// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
	Format *string `pulumi:"format"`
	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled *bool `pulumi:"multiVolumeSnapshottingEnabled"`
	// — Name of the policy
	Name *string `pulumi:"name"`
	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled *bool `pulumi:"performanceWindowEnabled"`
	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings *bool `pulumi:"preserveFileSecuritySettings"`
	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled *bool `pulumi:"quiesceSnapshottingEnabled"`
	// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
	Reattempts *BackupPolicyReattempts `pulumi:"reattempts"`
	// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
	Retention *BackupPolicyRetention `pulumi:"retention"`
	// — Schedule settings for creating backups on the host.
	Scheduling *BackupPolicyScheduling `pulumi:"scheduling"`
	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled *bool `pulumi:"silentModeEnabled"`
	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes *string `pulumi:"splittingBytes"`
	UpdatedAt      *string `pulumi:"updatedAt"`
	// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
	VmSnapshotReattempts *BackupPolicyVmSnapshotReattempts `pulumi:"vmSnapshotReattempts"`
	// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
	VssProvider *string `pulumi:"vssProvider"`
}

type BackupPolicyState struct {
	// — The name of generated archives.
	ArchiveName pulumi.StringPtrInput
	// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
	Cbt pulumi.StringPtrInput
	// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
	Compression       pulumi.StringPtrInput
	CreatedAt         pulumi.StringPtrInput
	Enabled           pulumi.BoolPtrInput
	FastBackupEnabled pulumi.BoolPtrInput
	FolderId          pulumi.StringPtrInput
	// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
	Format pulumi.StringPtrInput
	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled pulumi.BoolPtrInput
	// — Name of the policy
	Name pulumi.StringPtrInput
	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled pulumi.BoolPtrInput
	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings pulumi.BoolPtrInput
	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled pulumi.BoolPtrInput
	// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
	Reattempts BackupPolicyReattemptsPtrInput
	// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
	Retention BackupPolicyRetentionPtrInput
	// — Schedule settings for creating backups on the host.
	Scheduling BackupPolicySchedulingPtrInput
	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled pulumi.BoolPtrInput
	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes pulumi.StringPtrInput
	UpdatedAt      pulumi.StringPtrInput
	// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
	VmSnapshotReattempts BackupPolicyVmSnapshotReattemptsPtrInput
	// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
	VssProvider pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// — The name of generated archives.
	ArchiveName *string `pulumi:"archiveName"`
	// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
	Cbt *string `pulumi:"cbt"`
	// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
	Compression       *string `pulumi:"compression"`
	FastBackupEnabled *bool   `pulumi:"fastBackupEnabled"`
	FolderId          *string `pulumi:"folderId"`
	// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
	Format *string `pulumi:"format"`
	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled *bool `pulumi:"multiVolumeSnapshottingEnabled"`
	// — Name of the policy
	Name *string `pulumi:"name"`
	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled *bool `pulumi:"performanceWindowEnabled"`
	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings *bool `pulumi:"preserveFileSecuritySettings"`
	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled *bool `pulumi:"quiesceSnapshottingEnabled"`
	// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
	Reattempts BackupPolicyReattempts `pulumi:"reattempts"`
	// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
	Retention BackupPolicyRetention `pulumi:"retention"`
	// — Schedule settings for creating backups on the host.
	Scheduling BackupPolicyScheduling `pulumi:"scheduling"`
	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled *bool `pulumi:"silentModeEnabled"`
	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes *string `pulumi:"splittingBytes"`
	// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
	VmSnapshotReattempts BackupPolicyVmSnapshotReattempts `pulumi:"vmSnapshotReattempts"`
	// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
	VssProvider *string `pulumi:"vssProvider"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// — The name of generated archives.
	ArchiveName pulumi.StringPtrInput
	// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
	Cbt pulumi.StringPtrInput
	// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
	Compression       pulumi.StringPtrInput
	FastBackupEnabled pulumi.BoolPtrInput
	FolderId          pulumi.StringPtrInput
	// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
	Format pulumi.StringPtrInput
	// — If true, snapshots of multiple volumes will be taken simultaneously.
	MultiVolumeSnapshottingEnabled pulumi.BoolPtrInput
	// — Name of the policy
	Name pulumi.StringPtrInput
	// — Time windows for performance limitations of backup.
	PerformanceWindowEnabled pulumi.BoolPtrInput
	// — Preserves file security settings. It's better to set this option to true.
	PreserveFileSecuritySettings pulumi.BoolPtrInput
	// — If true, a quiesced snapshot of the virtual machine will be taken.
	QuiesceSnapshottingEnabled pulumi.BoolPtrInput
	// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
	Reattempts BackupPolicyReattemptsInput
	// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
	Retention BackupPolicyRetentionInput
	// — Schedule settings for creating backups on the host.
	Scheduling BackupPolicySchedulingInput
	// — if true, a user interaction will be avoided when possible.
	SilentModeEnabled pulumi.BoolPtrInput
	// — determines the size to split backups. It's better to leave this option unchanged.
	SplittingBytes pulumi.StringPtrInput
	// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
	VmSnapshotReattempts BackupPolicyVmSnapshotReattemptsInput
	// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
	VssProvider pulumi.StringPtrInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

// BackupPolicyArrayInput is an input type that accepts BackupPolicyArray and BackupPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPolicyArrayInput` via:
//
//	BackupPolicyArray{ BackupPolicyArgs{...} }
type BackupPolicyArrayInput interface {
	pulumi.Input

	ToBackupPolicyArrayOutput() BackupPolicyArrayOutput
	ToBackupPolicyArrayOutputWithContext(context.Context) BackupPolicyArrayOutput
}

type BackupPolicyArray []BackupPolicyInput

func (BackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return i.ToBackupPolicyArrayOutputWithContext(context.Background())
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyArrayOutput)
}

// BackupPolicyMapInput is an input type that accepts BackupPolicyMap and BackupPolicyMapOutput values.
// You can construct a concrete instance of `BackupPolicyMapInput` via:
//
//	BackupPolicyMap{ "key": BackupPolicyArgs{...} }
type BackupPolicyMapInput interface {
	pulumi.Input

	ToBackupPolicyMapOutput() BackupPolicyMapOutput
	ToBackupPolicyMapOutputWithContext(context.Context) BackupPolicyMapOutput
}

type BackupPolicyMap map[string]BackupPolicyInput

func (BackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyMap) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return i.ToBackupPolicyMapOutputWithContext(context.Background())
}

func (i BackupPolicyMap) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMapOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

// — The name of generated archives.
func (o BackupPolicyOutput) ArchiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.ArchiveName }).(pulumi.StringPtrOutput)
}

// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
func (o BackupPolicyOutput) Cbt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.Cbt }).(pulumi.StringPtrOutput)
}

// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
func (o BackupPolicyOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o BackupPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o BackupPolicyOutput) FastBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.FastBackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o BackupPolicyOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
func (o BackupPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// — If true, snapshots of multiple volumes will be taken simultaneously.
func (o BackupPolicyOutput) MultiVolumeSnapshottingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.MultiVolumeSnapshottingEnabled }).(pulumi.BoolPtrOutput)
}

// — Name of the policy
func (o BackupPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// — Time windows for performance limitations of backup.
func (o BackupPolicyOutput) PerformanceWindowEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.PerformanceWindowEnabled }).(pulumi.BoolPtrOutput)
}

// — Preserves file security settings. It's better to set this option to true.
func (o BackupPolicyOutput) PreserveFileSecuritySettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.PreserveFileSecuritySettings }).(pulumi.BoolPtrOutput)
}

// — If true, a quiesced snapshot of the virtual machine will be taken.
func (o BackupPolicyOutput) QuiesceSnapshottingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.QuiesceSnapshottingEnabled }).(pulumi.BoolPtrOutput)
}

// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
func (o BackupPolicyOutput) Reattempts() BackupPolicyReattemptsOutput {
	return o.ApplyT(func(v *BackupPolicy) BackupPolicyReattemptsOutput { return v.Reattempts }).(BackupPolicyReattemptsOutput)
}

// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
func (o BackupPolicyOutput) Retention() BackupPolicyRetentionOutput {
	return o.ApplyT(func(v *BackupPolicy) BackupPolicyRetentionOutput { return v.Retention }).(BackupPolicyRetentionOutput)
}

// — Schedule settings for creating backups on the host.
func (o BackupPolicyOutput) Scheduling() BackupPolicySchedulingOutput {
	return o.ApplyT(func(v *BackupPolicy) BackupPolicySchedulingOutput { return v.Scheduling }).(BackupPolicySchedulingOutput)
}

// — if true, a user interaction will be avoided when possible.
func (o BackupPolicyOutput) SilentModeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.BoolPtrOutput { return v.SilentModeEnabled }).(pulumi.BoolPtrOutput)
}

// — determines the size to split backups. It's better to leave this option unchanged.
func (o BackupPolicyOutput) SplittingBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.SplittingBytes }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
func (o BackupPolicyOutput) VmSnapshotReattempts() BackupPolicyVmSnapshotReattemptsOutput {
	return o.ApplyT(func(v *BackupPolicy) BackupPolicyVmSnapshotReattemptsOutput { return v.VmSnapshotReattempts }).(BackupPolicyVmSnapshotReattemptsOutput)
}

// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
func (o BackupPolicyOutput) VssProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.VssProvider }).(pulumi.StringPtrOutput)
}

type BackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) Index(i pulumi.IntInput) BackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].([]*BackupPolicy)[vs[1].(int)]
	}).(BackupPolicyOutput)
}

type BackupPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].(map[string]*BackupPolicy)[vs[1].(string)]
	}).(BackupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyInput)(nil)).Elem(), &BackupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyArrayInput)(nil)).Elem(), BackupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyMapInput)(nil)).Elem(), BackupPolicyMap{})
	pulumi.RegisterOutputType(BackupPolicyOutput{})
	pulumi.RegisterOutputType(BackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyMapOutput{})
}
