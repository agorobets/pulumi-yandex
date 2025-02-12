// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Yandex VPC subnet. For more information, see [Yandex Cloud VPC](https://cloud.yandex.com/docs/vpc/concepts/index).
//
// ## Example Usage
//
// {{ tffile "examples/vpc_subnet/d_vpc_subnet_1.tf" }}
//
// This data source is used to define [VPC Subnets](https://cloud.yandex.com/docs/vpc/concepts/network#subnet) that can be used by other resources.
func LookupVpcSubnet(ctx *pulumi.Context, args *LookupVpcSubnetArgs, opts ...pulumi.InvokeOption) (*LookupVpcSubnetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcSubnetResult
	err := ctx.Invoke("yandex:index/getVpcSubnet:getVpcSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcSubnet.
type LookupVpcSubnetArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Name of the subnet.
	//
	// > One of `subnetId` or `name` should be specified.
	Name *string `pulumi:"name"`
	// Subnet ID.
	SubnetId *string `pulumi:"subnetId"`
}

// A collection of values returned by getVpcSubnet.
type LookupVpcSubnetResult struct {
	// Creation timestamp of this subnet.
	CreatedAt string `pulumi:"createdAt"`
	// Description of the subnet.
	Description string `pulumi:"description"`
	// Options for DHCP client. The structure is documented below.
	DhcpOptions []GetVpcSubnetDhcpOption `pulumi:"dhcpOptions"`
	FolderId    string                   `pulumi:"folderId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Labels to assign to this subnet.
	Labels map[string]string `pulumi:"labels"`
	Name   string            `pulumi:"name"`
	// ID of the network this subnet belongs to.
	NetworkId string `pulumi:"networkId"`
	// ID of the route table to assign to this subnet.
	RouteTableId string `pulumi:"routeTableId"`
	SubnetId     string `pulumi:"subnetId"`
	// The blocks of internal IPv4 addresses owned by this subnet.
	V4CidrBlocks []string `pulumi:"v4CidrBlocks"`
	// The blocks of internal IPv6 addresses owned by this subnet.
	V6CidrBlocks []string `pulumi:"v6CidrBlocks"`
	// Name of the availability zone for this subnet.
	Zone string `pulumi:"zone"`
}

func LookupVpcSubnetOutput(ctx *pulumi.Context, args LookupVpcSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupVpcSubnetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcSubnetResultOutput, error) {
			args := v.(LookupVpcSubnetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getVpcSubnet:getVpcSubnet", args, LookupVpcSubnetResultOutput{}, options).(LookupVpcSubnetResultOutput), nil
		}).(LookupVpcSubnetResultOutput)
}

// A collection of arguments for invoking getVpcSubnet.
type LookupVpcSubnetOutputArgs struct {
	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderId pulumi.StringPtrInput `pulumi:"folderId"`
	// Name of the subnet.
	//
	// > One of `subnetId` or `name` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Subnet ID.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (LookupVpcSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcSubnetArgs)(nil)).Elem()
}

// A collection of values returned by getVpcSubnet.
type LookupVpcSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupVpcSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcSubnetResult)(nil)).Elem()
}

func (o LookupVpcSubnetResultOutput) ToLookupVpcSubnetResultOutput() LookupVpcSubnetResultOutput {
	return o
}

func (o LookupVpcSubnetResultOutput) ToLookupVpcSubnetResultOutputWithContext(ctx context.Context) LookupVpcSubnetResultOutput {
	return o
}

// Creation timestamp of this subnet.
func (o LookupVpcSubnetResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the subnet.
func (o LookupVpcSubnetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.Description }).(pulumi.StringOutput)
}

// Options for DHCP client. The structure is documented below.
func (o LookupVpcSubnetResultOutput) DhcpOptions() GetVpcSubnetDhcpOptionArrayOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) []GetVpcSubnetDhcpOption { return v.DhcpOptions }).(GetVpcSubnetDhcpOptionArrayOutput)
}

func (o LookupVpcSubnetResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.FolderId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Labels to assign to this subnet.
func (o LookupVpcSubnetResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupVpcSubnetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the network this subnet belongs to.
func (o LookupVpcSubnetResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// ID of the route table to assign to this subnet.
func (o LookupVpcSubnetResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

func (o LookupVpcSubnetResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// The blocks of internal IPv4 addresses owned by this subnet.
func (o LookupVpcSubnetResultOutput) V4CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) []string { return v.V4CidrBlocks }).(pulumi.StringArrayOutput)
}

// The blocks of internal IPv6 addresses owned by this subnet.
func (o LookupVpcSubnetResultOutput) V6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) []string { return v.V6CidrBlocks }).(pulumi.StringArrayOutput)
}

// Name of the availability zone for this subnet.
func (o LookupVpcSubnetResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcSubnetResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcSubnetResultOutput{})
}
