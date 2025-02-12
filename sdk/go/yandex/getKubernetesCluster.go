// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex Cloud Managed Kubernetes Cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kubernetes/concepts/#kubernetes-cluster).
//
// ## Example Usage
//
// {{ tffile "examples/kubernetes_cluster/d_kubernetes_cluster_1.tf" }}
func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("yandex:index/getKubernetesCluster:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterArgs struct {
	// ID of a specific Kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Name of a specific Kubernetes cluster.
	//
	// > One of `clusterId` or `name` should be specified.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// IP range for allocating pod addresses.
	ClusterIpv4Range string `pulumi:"clusterIpv4Range"`
	// Identical to clusterIpv4Range but for the IPv6 protocol.
	ClusterIpv6Range string `pulumi:"clusterIpv6Range"`
	// The Kubernetes cluster creation timestamp.
	CreatedAt string `pulumi:"createdAt"`
	// A description of the Kubernetes cluster.
	Description string `pulumi:"description"`
	// (Optional) ID of the folder default Log group of which should be used to collect logs.
	FolderId string `pulumi:"folderId"`
	// Health of the Kubernetes cluster.
	Health string `pulumi:"health"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// cluster KMS provider parameters.
	KmsProviders []GetKubernetesClusterKmsProvider `pulumi:"kmsProviders"`
	// A set of key/value label pairs to assign to the Kubernetes cluster.
	Labels map[string]string `pulumi:"labels"`
	// (Optional) ID of the Yandex Cloud Logging [Log group](https://cloud.yandex.com/docs/logging/concepts/log-group).
	LogGroupId string `pulumi:"logGroupId"`
	// Kubernetes master configuration options. The structure is documented below.
	Masters []GetKubernetesClusterMaster `pulumi:"masters"`
	Name    string                       `pulumi:"name"`
	// The ID of the cluster network.
	NetworkId string `pulumi:"networkId"`
	// (Optional) Network Implementation options. The structure is documented below.
	NetworkImplementations []GetKubernetesClusterNetworkImplementation `pulumi:"networkImplementations"`
	// Network policy provider for the cluster, if present. Possible values: `CALICO`.
	NetworkPolicyProvider string `pulumi:"networkPolicyProvider"`
	// Size of the masks that are assigned to each node in the cluster.
	NodeIpv4CidrMaskSize int `pulumi:"nodeIpv4CidrMaskSize"`
	// Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.
	NodeServiceAccountId string `pulumi:"nodeServiceAccountId"`
	// Cluster release channel.
	ReleaseChannel string `pulumi:"releaseChannel"`
	// Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from.
	ServiceIpv4Range string `pulumi:"serviceIpv4Range"`
	// Identical to serviceIpv4Range but for the IPv6 protocol.
	ServiceIpv6Range string `pulumi:"serviceIpv6Range"`
	// Status of the Kubernetes cluster.
	Status string `pulumi:"status"`
}

func LookupKubernetesClusterOutput(ctx *pulumi.Context, args LookupKubernetesClusterOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubernetesClusterResultOutput, error) {
			args := v.(LookupKubernetesClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getKubernetesCluster:getKubernetesCluster", args, LookupKubernetesClusterResultOutput{}, options).(LookupKubernetesClusterResultOutput), nil
		}).(LookupKubernetesClusterResultOutput)
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterOutputArgs struct {
	// ID of a specific Kubernetes cluster.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// Name of a specific Kubernetes cluster.
	//
	// > One of `clusterId` or `name` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupKubernetesClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesClusterResult)(nil)).Elem()
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutput() LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) ToLookupKubernetesClusterResultOutputWithContext(ctx context.Context) LookupKubernetesClusterResultOutput {
	return o
}

func (o LookupKubernetesClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// IP range for allocating pod addresses.
func (o LookupKubernetesClusterResultOutput) ClusterIpv4Range() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ClusterIpv4Range }).(pulumi.StringOutput)
}

// Identical to clusterIpv4Range but for the IPv6 protocol.
func (o LookupKubernetesClusterResultOutput) ClusterIpv6Range() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ClusterIpv6Range }).(pulumi.StringOutput)
}

// The Kubernetes cluster creation timestamp.
func (o LookupKubernetesClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description of the Kubernetes cluster.
func (o LookupKubernetesClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// (Optional) ID of the folder default Log group of which should be used to collect logs.
func (o LookupKubernetesClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// Health of the Kubernetes cluster.
func (o LookupKubernetesClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKubernetesClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// cluster KMS provider parameters.
func (o LookupKubernetesClusterResultOutput) KmsProviders() GetKubernetesClusterKmsProviderArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []GetKubernetesClusterKmsProvider { return v.KmsProviders }).(GetKubernetesClusterKmsProviderArrayOutput)
}

// A set of key/value label pairs to assign to the Kubernetes cluster.
func (o LookupKubernetesClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// (Optional) ID of the Yandex Cloud Logging [Log group](https://cloud.yandex.com/docs/logging/concepts/log-group).
func (o LookupKubernetesClusterResultOutput) LogGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.LogGroupId }).(pulumi.StringOutput)
}

// Kubernetes master configuration options. The structure is documented below.
func (o LookupKubernetesClusterResultOutput) Masters() GetKubernetesClusterMasterArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []GetKubernetesClusterMaster { return v.Masters }).(GetKubernetesClusterMasterArrayOutput)
}

func (o LookupKubernetesClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the cluster network.
func (o LookupKubernetesClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// (Optional) Network Implementation options. The structure is documented below.
func (o LookupKubernetesClusterResultOutput) NetworkImplementations() GetKubernetesClusterNetworkImplementationArrayOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) []GetKubernetesClusterNetworkImplementation {
		return v.NetworkImplementations
	}).(GetKubernetesClusterNetworkImplementationArrayOutput)
}

// Network policy provider for the cluster, if present. Possible values: `CALICO`.
func (o LookupKubernetesClusterResultOutput) NetworkPolicyProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.NetworkPolicyProvider }).(pulumi.StringOutput)
}

// Size of the masks that are assigned to each node in the cluster.
func (o LookupKubernetesClusterResultOutput) NodeIpv4CidrMaskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) int { return v.NodeIpv4CidrMaskSize }).(pulumi.IntOutput)
}

// Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.
func (o LookupKubernetesClusterResultOutput) NodeServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.NodeServiceAccountId }).(pulumi.StringOutput)
}

// Cluster release channel.
func (o LookupKubernetesClusterResultOutput) ReleaseChannel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ReleaseChannel }).(pulumi.StringOutput)
}

// Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have `edit` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.
func (o LookupKubernetesClusterResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// IP range Kubernetes services Kubernetes cluster IP addresses will be allocated from.
func (o LookupKubernetesClusterResultOutput) ServiceIpv4Range() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ServiceIpv4Range }).(pulumi.StringOutput)
}

// Identical to serviceIpv4Range but for the IPv6 protocol.
func (o LookupKubernetesClusterResultOutput) ServiceIpv6Range() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.ServiceIpv6Range }).(pulumi.StringOutput)
}

// Status of the Kubernetes cluster.
func (o LookupKubernetesClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesClusterResultOutput{})
}
