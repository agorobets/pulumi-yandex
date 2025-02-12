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
    /// Allows management of [Yandex Cloud Backup Policy](https://yandex.cloud/docs/backup/concepts/policy).
    /// 
    /// &gt; Cloud Backup Provider must be activated in order to manipulate with policies. Active it either by UI Console or by `yc` command.
    /// 
    /// ## Example Usage
    /// 
    /// {{ tffile "examples/backup_policy/r_backup_policy_1.tf" }}
    /// 
    /// For the full policy attributes, take a look at the following example:
    /// 
    /// {{ tffile "examples/backup_policy/r_backup_policy_2.tf" }}
    /// 
    /// ## Defined types
    /// 
    /// ### interval_type 
    /// 
    /// A string type, that accepts values in the format of: `number` + `time type`, where `time type` might be:
    /// 
    /// - `s` — seconds
    /// - `m` — minutes
    /// - `h` — hours
    /// - `d` — days
    /// - `w` — weekdays
    /// - `M` — months
    /// 
    /// Example of interval value: `"5m", "10d", "2M", "5w"`
    /// 
    /// ### day_type
    /// 
    /// A string type, that accepts the following values: `"ALWAYS_INCREMENTAL"`, `"ALWAYS_FULL"`, `"WEEKLY_FULL_DAILY_INCREMENTAL"`, `'WEEKLY_INCREMENTAL"`.
    /// 
    /// ### backup_set_type
    /// 
    /// `"TYPE_AUTO"`, `"TYPE_FULL"`, `"TYPE_INCREMENTAL"`, `'TYPE_DIFFERENTIAL"`.
    /// </summary>
    [YandexResourceType("yandex:index/backupPolicy:BackupPolicy")]
    public partial class BackupPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// — The name of generated archives.
        /// </summary>
        [Output("archiveName")]
        public Output<string?> ArchiveName { get; private set; } = null!;

        /// <summary>
        /// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
        /// </summary>
        [Output("cbt")]
        public Output<string?> Cbt { get; private set; } = null!;

        /// <summary>
        /// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
        /// </summary>
        [Output("compression")]
        public Output<string?> Compression { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("fastBackupEnabled")]
        public Output<bool?> FastBackupEnabled { get; private set; } = null!;

        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

        /// <summary>
        /// — If true, snapshots of multiple volumes will be taken simultaneously.
        /// </summary>
        [Output("multiVolumeSnapshottingEnabled")]
        public Output<bool?> MultiVolumeSnapshottingEnabled { get; private set; } = null!;

        /// <summary>
        /// — Name of the policy
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// — Time windows for performance limitations of backup.
        /// </summary>
        [Output("performanceWindowEnabled")]
        public Output<bool?> PerformanceWindowEnabled { get; private set; } = null!;

        /// <summary>
        /// — Preserves file security settings. It's better to set this option to true.
        /// </summary>
        [Output("preserveFileSecuritySettings")]
        public Output<bool?> PreserveFileSecuritySettings { get; private set; } = null!;

        /// <summary>
        /// — If true, a quiesced snapshot of the virtual machine will be taken.
        /// </summary>
        [Output("quiesceSnapshottingEnabled")]
        public Output<bool?> QuiesceSnapshottingEnabled { get; private set; } = null!;

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
        /// </summary>
        [Output("reattempts")]
        public Output<Outputs.BackupPolicyReattempts> Reattempts { get; private set; } = null!;

        /// <summary>
        /// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
        /// </summary>
        [Output("retention")]
        public Output<Outputs.BackupPolicyRetention> Retention { get; private set; } = null!;

        /// <summary>
        /// — Schedule settings for creating backups on the host.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.BackupPolicyScheduling> Scheduling { get; private set; } = null!;

        /// <summary>
        /// — if true, a user interaction will be avoided when possible.
        /// </summary>
        [Output("silentModeEnabled")]
        public Output<bool?> SilentModeEnabled { get; private set; } = null!;

        /// <summary>
        /// — determines the size to split backups. It's better to leave this option unchanged.
        /// </summary>
        [Output("splittingBytes")]
        public Output<string?> SplittingBytes { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
        /// </summary>
        [Output("vmSnapshotReattempts")]
        public Output<Outputs.BackupPolicyVmSnapshotReattempts> VmSnapshotReattempts { get; private set; } = null!;

        /// <summary>
        /// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
        /// </summary>
        [Output("vssProvider")]
        public Output<string?> VssProvider { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPolicy(string name, BackupPolicyArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/backupPolicy:BackupPolicy", name, args ?? new BackupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPolicy(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/backupPolicy:BackupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPolicy Get(string name, Input<string> id, BackupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupPolicy(name, id, state, options);
        }
    }

    public sealed class BackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// — The name of generated archives.
        /// </summary>
        [Input("archiveName")]
        public Input<string>? ArchiveName { get; set; }

        /// <summary>
        /// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
        /// </summary>
        [Input("cbt")]
        public Input<string>? Cbt { get; set; }

        /// <summary>
        /// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        [Input("fastBackupEnabled")]
        public Input<bool>? FastBackupEnabled { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// — If true, snapshots of multiple volumes will be taken simultaneously.
        /// </summary>
        [Input("multiVolumeSnapshottingEnabled")]
        public Input<bool>? MultiVolumeSnapshottingEnabled { get; set; }

        /// <summary>
        /// — Name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// — Time windows for performance limitations of backup.
        /// </summary>
        [Input("performanceWindowEnabled")]
        public Input<bool>? PerformanceWindowEnabled { get; set; }

        /// <summary>
        /// — Preserves file security settings. It's better to set this option to true.
        /// </summary>
        [Input("preserveFileSecuritySettings")]
        public Input<bool>? PreserveFileSecuritySettings { get; set; }

        /// <summary>
        /// — If true, a quiesced snapshot of the virtual machine will be taken.
        /// </summary>
        [Input("quiesceSnapshottingEnabled")]
        public Input<bool>? QuiesceSnapshottingEnabled { get; set; }

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
        /// </summary>
        [Input("reattempts", required: true)]
        public Input<Inputs.BackupPolicyReattemptsArgs> Reattempts { get; set; } = null!;

        /// <summary>
        /// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
        /// </summary>
        [Input("retention", required: true)]
        public Input<Inputs.BackupPolicyRetentionArgs> Retention { get; set; } = null!;

        /// <summary>
        /// — Schedule settings for creating backups on the host.
        /// </summary>
        [Input("scheduling", required: true)]
        public Input<Inputs.BackupPolicySchedulingArgs> Scheduling { get; set; } = null!;

        /// <summary>
        /// — if true, a user interaction will be avoided when possible.
        /// </summary>
        [Input("silentModeEnabled")]
        public Input<bool>? SilentModeEnabled { get; set; }

        /// <summary>
        /// — determines the size to split backups. It's better to leave this option unchanged.
        /// </summary>
        [Input("splittingBytes")]
        public Input<string>? SplittingBytes { get; set; }

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
        /// </summary>
        [Input("vmSnapshotReattempts", required: true)]
        public Input<Inputs.BackupPolicyVmSnapshotReattemptsArgs> VmSnapshotReattempts { get; set; } = null!;

        /// <summary>
        /// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
        /// </summary>
        [Input("vssProvider")]
        public Input<string>? VssProvider { get; set; }

        public BackupPolicyArgs()
        {
        }
        public static new BackupPolicyArgs Empty => new BackupPolicyArgs();
    }

    public sealed class BackupPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// — The name of generated archives.
        /// </summary>
        [Input("archiveName")]
        public Input<string>? ArchiveName { get; set; }

        /// <summary>
        /// — Configuration of Changed Block Tracking. Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
        /// </summary>
        [Input("cbt")]
        public Input<string>? Cbt { get; set; }

        /// <summary>
        /// — Archive compression level. Affects CPU. Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
        /// </summary>
        [Input("compression")]
        public Input<string>? Compression { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("fastBackupEnabled")]
        public Input<bool>? FastBackupEnabled { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`. Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// — If true, snapshots of multiple volumes will be taken simultaneously.
        /// </summary>
        [Input("multiVolumeSnapshottingEnabled")]
        public Input<bool>? MultiVolumeSnapshottingEnabled { get; set; }

        /// <summary>
        /// — Name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// — Time windows for performance limitations of backup.
        /// </summary>
        [Input("performanceWindowEnabled")]
        public Input<bool>? PerformanceWindowEnabled { get; set; }

        /// <summary>
        /// — Preserves file security settings. It's better to set this option to true.
        /// </summary>
        [Input("preserveFileSecuritySettings")]
        public Input<bool>? PreserveFileSecuritySettings { get; set; }

        /// <summary>
        /// — If true, a quiesced snapshot of the virtual machine will be taken.
        /// </summary>
        [Input("quiesceSnapshottingEnabled")]
        public Input<bool>? QuiesceSnapshottingEnabled { get; set; }

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make backup at the host. This attribute consists of the following parameters:
        /// </summary>
        [Input("reattempts")]
        public Input<Inputs.BackupPolicyReattemptsGetArgs>? Reattempts { get; set; }

        /// <summary>
        /// — Retention policy for backups. Allows to setup backups lifecycle. This attribute consists of the following parameters:
        /// </summary>
        [Input("retention")]
        public Input<Inputs.BackupPolicyRetentionGetArgs>? Retention { get; set; }

        /// <summary>
        /// — Schedule settings for creating backups on the host.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.BackupPolicySchedulingGetArgs>? Scheduling { get; set; }

        /// <summary>
        /// — if true, a user interaction will be avoided when possible.
        /// </summary>
        [Input("silentModeEnabled")]
        public Input<bool>? SilentModeEnabled { get; set; }

        /// <summary>
        /// — determines the size to split backups. It's better to leave this option unchanged.
        /// </summary>
        [Input("splittingBytes")]
        public Input<string>? SplittingBytes { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// — Amount of reattempts that should be performed while trying to make snapshot. This attribute consists of the following parameters:
        /// </summary>
        [Input("vmSnapshotReattempts")]
        public Input<Inputs.BackupPolicyVmSnapshotReattemptsGetArgs>? VmSnapshotReattempts { get; set; }

        /// <summary>
        /// — Settings for the volume shadow copy service. Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
        /// </summary>
        [Input("vssProvider")]
        public Input<string>? VssProvider { get; set; }

        public BackupPolicyState()
        {
        }
        public static new BackupPolicyState Empty => new BackupPolicyState();
    }
}
