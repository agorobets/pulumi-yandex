// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Compute snapshot. For more information, see [the official documentation](https://cloud.yandex.com/docs/compute/concepts/snapshot).
//
// ## Example Usage
//
// {{ tffile "examples/compute_snapshot/d_compute_snapshot_1.tf" }}
func LookupComputeSnapshot(ctx *pulumi.Context, args *LookupComputeSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupComputeSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComputeSnapshotResult
	err := ctx.Invoke("yandex:index/getComputeSnapshot:getComputeSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeSnapshot.
type LookupComputeSnapshotArgs struct {
	// ID of the folder that the snapshot belongs to.
	FolderId *string `pulumi:"folderId"`
	// The name of the snapshot.
	//
	// > One of `snapshotId` or `name` should be specified.
	Name *string `pulumi:"name"`
	// The ID of a specific snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
}

// A collection of values returned by getComputeSnapshot.
type LookupComputeSnapshotResult struct {
	// Snapshot creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// An optional description of this snapshot.
	Description string `pulumi:"description"`
	// Minimum required size of the disk which is created from this snapshot.
	DiskSize int `pulumi:"diskSize"`
	// ID of the folder that the snapshot belongs to.
	FolderId            string                                 `pulumi:"folderId"`
	HardwareGenerations []GetComputeSnapshotHardwareGeneration `pulumi:"hardwareGenerations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of KMS symmetric key used to encrypt snapshot.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// A map of labels applied to this snapshot.
	Labels map[string]string `pulumi:"labels"`
	Name   string            `pulumi:"name"`
	// License IDs that indicate which licenses are attached to this snapshot.
	ProductIds []string `pulumi:"productIds"`
	SnapshotId string   `pulumi:"snapshotId"`
	// ID of the source disk.
	SourceDiskId string `pulumi:"sourceDiskId"`
	// The status of the snapshot.
	Status string `pulumi:"status"`
	// The size of the snapshot, specified in Gb.
	StorageSize int `pulumi:"storageSize"`
}

func LookupComputeSnapshotOutput(ctx *pulumi.Context, args LookupComputeSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupComputeSnapshotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeSnapshotResultOutput, error) {
			args := v.(LookupComputeSnapshotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getComputeSnapshot:getComputeSnapshot", args, LookupComputeSnapshotResultOutput{}, options).(LookupComputeSnapshotResultOutput), nil
		}).(LookupComputeSnapshotResultOutput)
}

// A collection of arguments for invoking getComputeSnapshot.
type LookupComputeSnapshotOutputArgs struct {
	// ID of the folder that the snapshot belongs to.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// The name of the snapshot.
	//
	// > One of `snapshotId` or `name` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of a specific snapshot.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
}

func (LookupComputeSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getComputeSnapshot.
type LookupComputeSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupComputeSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSnapshotResult)(nil)).Elem()
}

func (o LookupComputeSnapshotResultOutput) ToLookupComputeSnapshotResultOutput() LookupComputeSnapshotResultOutput {
	return o
}

func (o LookupComputeSnapshotResultOutput) ToLookupComputeSnapshotResultOutputWithContext(ctx context.Context) LookupComputeSnapshotResultOutput {
	return o
}

// Snapshot creation timestamp.
func (o LookupComputeSnapshotResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// An optional description of this snapshot.
func (o LookupComputeSnapshotResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.Description }).(pulumi.StringOutput)
}

// Minimum required size of the disk which is created from this snapshot.
func (o LookupComputeSnapshotResultOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) int { return v.DiskSize }).(pulumi.IntOutput)
}

// ID of the folder that the snapshot belongs to.
func (o LookupComputeSnapshotResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupComputeSnapshotResultOutput) HardwareGenerations() GetComputeSnapshotHardwareGenerationArrayOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) []GetComputeSnapshotHardwareGeneration {
		return v.HardwareGenerations
	}).(GetComputeSnapshotHardwareGenerationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupComputeSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of KMS symmetric key used to encrypt snapshot.
func (o LookupComputeSnapshotResultOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

// A map of labels applied to this snapshot.
func (o LookupComputeSnapshotResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupComputeSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// License IDs that indicate which licenses are attached to this snapshot.
func (o LookupComputeSnapshotResultOutput) ProductIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) []string { return v.ProductIds }).(pulumi.StringArrayOutput)
}

func (o LookupComputeSnapshotResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// ID of the source disk.
func (o LookupComputeSnapshotResultOutput) SourceDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.SourceDiskId }).(pulumi.StringOutput)
}

// The status of the snapshot.
func (o LookupComputeSnapshotResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) string { return v.Status }).(pulumi.StringOutput)
}

// The size of the snapshot, specified in Gb.
func (o LookupComputeSnapshotResultOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeSnapshotResult) int { return v.StorageSize }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeSnapshotResultOutput{})
}
