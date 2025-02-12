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

// Manages a VPC Private Endpoint within the Yandex Cloud. For more information, see [the official documentation](https://yandex.cloud/docs/vpc/concepts/private-endpoint).
//
// * How-to Guides
//   - [Cloud Networking](https://cloud.yandex.com/docs/vpc/)
//
// ## Example Usage
//
// {{ tffile "examples/vpc_private_endpoint/r_vpc_private_endpoint_1.tf" }}
//
// ## Import
//
// Private endpoint can be imported using the `id` of the resource, e.g.
//
// ```sh
// $ pulumi import yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint pe private_endpoint_id
// ```
type VpcPrivateEndpoint struct {
	pulumi.CustomResourceState

	// Creation timestamp of the key.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Private endpoint DNS options block.
	DnsOptions VpcPrivateEndpointDnsOptionsOutput `pulumi:"dnsOptions"`
	// Private endpoint address specification block.
	EndpointAddress VpcPrivateEndpointEndpointAddressOutput `pulumi:"endpointAddress"`
	// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Labels to apply to this resource. A list of key/value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the private endpoint. Provided by the client when the private endpoint is created.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network which private endpoint belongs to.
	NetworkId     pulumi.StringOutput                   `pulumi:"networkId"`
	ObjectStorage VpcPrivateEndpointObjectStorageOutput `pulumi:"objectStorage"`
	// Status of the private endpoint.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewVpcPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcPrivateEndpoint(ctx *pulumi.Context,
	name string, args *VpcPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*VpcPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.ObjectStorage == nil {
		return nil, errors.New("invalid value for required argument 'ObjectStorage'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPrivateEndpoint
	err := ctx.RegisterResource("yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPrivateEndpoint gets an existing VpcPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPrivateEndpointState, opts ...pulumi.ResourceOption) (*VpcPrivateEndpoint, error) {
	var resource VpcPrivateEndpoint
	err := ctx.ReadResource("yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPrivateEndpoint resources.
type vpcPrivateEndpointState struct {
	// Creation timestamp of the key.
	CreatedAt *string `pulumi:"createdAt"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Private endpoint DNS options block.
	DnsOptions *VpcPrivateEndpointDnsOptions `pulumi:"dnsOptions"`
	// Private endpoint address specification block.
	EndpointAddress *VpcPrivateEndpointEndpointAddress `pulumi:"endpointAddress"`
	// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Labels to apply to this resource. A list of key/value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the private endpoint. Provided by the client when the private endpoint is created.
	Name *string `pulumi:"name"`
	// ID of the network which private endpoint belongs to.
	NetworkId     *string                          `pulumi:"networkId"`
	ObjectStorage *VpcPrivateEndpointObjectStorage `pulumi:"objectStorage"`
	// Status of the private endpoint.
	Status *string `pulumi:"status"`
}

type VpcPrivateEndpointState struct {
	// Creation timestamp of the key.
	CreatedAt pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Private endpoint DNS options block.
	DnsOptions VpcPrivateEndpointDnsOptionsPtrInput
	// Private endpoint address specification block.
	EndpointAddress VpcPrivateEndpointEndpointAddressPtrInput
	// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Labels to apply to this resource. A list of key/value pairs.
	Labels pulumi.StringMapInput
	// Name of the private endpoint. Provided by the client when the private endpoint is created.
	Name pulumi.StringPtrInput
	// ID of the network which private endpoint belongs to.
	NetworkId     pulumi.StringPtrInput
	ObjectStorage VpcPrivateEndpointObjectStoragePtrInput
	// Status of the private endpoint.
	Status pulumi.StringPtrInput
}

func (VpcPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPrivateEndpointState)(nil)).Elem()
}

type vpcPrivateEndpointArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Private endpoint DNS options block.
	DnsOptions *VpcPrivateEndpointDnsOptions `pulumi:"dnsOptions"`
	// Private endpoint address specification block.
	EndpointAddress *VpcPrivateEndpointEndpointAddress `pulumi:"endpointAddress"`
	// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId *string `pulumi:"folderId"`
	// Labels to apply to this resource. A list of key/value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the private endpoint. Provided by the client when the private endpoint is created.
	Name *string `pulumi:"name"`
	// ID of the network which private endpoint belongs to.
	NetworkId     string                          `pulumi:"networkId"`
	ObjectStorage VpcPrivateEndpointObjectStorage `pulumi:"objectStorage"`
}

// The set of arguments for constructing a VpcPrivateEndpoint resource.
type VpcPrivateEndpointArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Private endpoint DNS options block.
	DnsOptions VpcPrivateEndpointDnsOptionsPtrInput
	// Private endpoint address specification block.
	EndpointAddress VpcPrivateEndpointEndpointAddressPtrInput
	// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderId pulumi.StringPtrInput
	// Labels to apply to this resource. A list of key/value pairs.
	Labels pulumi.StringMapInput
	// Name of the private endpoint. Provided by the client when the private endpoint is created.
	Name pulumi.StringPtrInput
	// ID of the network which private endpoint belongs to.
	NetworkId     pulumi.StringInput
	ObjectStorage VpcPrivateEndpointObjectStorageInput
}

func (VpcPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPrivateEndpointArgs)(nil)).Elem()
}

type VpcPrivateEndpointInput interface {
	pulumi.Input

	ToVpcPrivateEndpointOutput() VpcPrivateEndpointOutput
	ToVpcPrivateEndpointOutputWithContext(ctx context.Context) VpcPrivateEndpointOutput
}

func (*VpcPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPrivateEndpoint)(nil)).Elem()
}

func (i *VpcPrivateEndpoint) ToVpcPrivateEndpointOutput() VpcPrivateEndpointOutput {
	return i.ToVpcPrivateEndpointOutputWithContext(context.Background())
}

func (i *VpcPrivateEndpoint) ToVpcPrivateEndpointOutputWithContext(ctx context.Context) VpcPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateEndpointOutput)
}

// VpcPrivateEndpointArrayInput is an input type that accepts VpcPrivateEndpointArray and VpcPrivateEndpointArrayOutput values.
// You can construct a concrete instance of `VpcPrivateEndpointArrayInput` via:
//
//	VpcPrivateEndpointArray{ VpcPrivateEndpointArgs{...} }
type VpcPrivateEndpointArrayInput interface {
	pulumi.Input

	ToVpcPrivateEndpointArrayOutput() VpcPrivateEndpointArrayOutput
	ToVpcPrivateEndpointArrayOutputWithContext(context.Context) VpcPrivateEndpointArrayOutput
}

type VpcPrivateEndpointArray []VpcPrivateEndpointInput

func (VpcPrivateEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPrivateEndpoint)(nil)).Elem()
}

func (i VpcPrivateEndpointArray) ToVpcPrivateEndpointArrayOutput() VpcPrivateEndpointArrayOutput {
	return i.ToVpcPrivateEndpointArrayOutputWithContext(context.Background())
}

func (i VpcPrivateEndpointArray) ToVpcPrivateEndpointArrayOutputWithContext(ctx context.Context) VpcPrivateEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateEndpointArrayOutput)
}

// VpcPrivateEndpointMapInput is an input type that accepts VpcPrivateEndpointMap and VpcPrivateEndpointMapOutput values.
// You can construct a concrete instance of `VpcPrivateEndpointMapInput` via:
//
//	VpcPrivateEndpointMap{ "key": VpcPrivateEndpointArgs{...} }
type VpcPrivateEndpointMapInput interface {
	pulumi.Input

	ToVpcPrivateEndpointMapOutput() VpcPrivateEndpointMapOutput
	ToVpcPrivateEndpointMapOutputWithContext(context.Context) VpcPrivateEndpointMapOutput
}

type VpcPrivateEndpointMap map[string]VpcPrivateEndpointInput

func (VpcPrivateEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPrivateEndpoint)(nil)).Elem()
}

func (i VpcPrivateEndpointMap) ToVpcPrivateEndpointMapOutput() VpcPrivateEndpointMapOutput {
	return i.ToVpcPrivateEndpointMapOutputWithContext(context.Background())
}

func (i VpcPrivateEndpointMap) ToVpcPrivateEndpointMapOutputWithContext(ctx context.Context) VpcPrivateEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateEndpointMapOutput)
}

type VpcPrivateEndpointOutput struct{ *pulumi.OutputState }

func (VpcPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPrivateEndpoint)(nil)).Elem()
}

func (o VpcPrivateEndpointOutput) ToVpcPrivateEndpointOutput() VpcPrivateEndpointOutput {
	return o
}

func (o VpcPrivateEndpointOutput) ToVpcPrivateEndpointOutputWithContext(ctx context.Context) VpcPrivateEndpointOutput {
	return o
}

// Creation timestamp of the key.
func (o VpcPrivateEndpointOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o VpcPrivateEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Private endpoint DNS options block.
func (o VpcPrivateEndpointOutput) DnsOptions() VpcPrivateEndpointDnsOptionsOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) VpcPrivateEndpointDnsOptionsOutput { return v.DnsOptions }).(VpcPrivateEndpointDnsOptionsOutput)
}

// Private endpoint address specification block.
func (o VpcPrivateEndpointOutput) EndpointAddress() VpcPrivateEndpointEndpointAddressOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) VpcPrivateEndpointEndpointAddressOutput { return v.EndpointAddress }).(VpcPrivateEndpointEndpointAddressOutput)
}

// ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
func (o VpcPrivateEndpointOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Labels to apply to this resource. A list of key/value pairs.
func (o VpcPrivateEndpointOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the private endpoint. Provided by the client when the private endpoint is created.
func (o VpcPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network which private endpoint belongs to.
func (o VpcPrivateEndpointOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o VpcPrivateEndpointOutput) ObjectStorage() VpcPrivateEndpointObjectStorageOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) VpcPrivateEndpointObjectStorageOutput { return v.ObjectStorage }).(VpcPrivateEndpointObjectStorageOutput)
}

// Status of the private endpoint.
func (o VpcPrivateEndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateEndpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type VpcPrivateEndpointArrayOutput struct{ *pulumi.OutputState }

func (VpcPrivateEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPrivateEndpoint)(nil)).Elem()
}

func (o VpcPrivateEndpointArrayOutput) ToVpcPrivateEndpointArrayOutput() VpcPrivateEndpointArrayOutput {
	return o
}

func (o VpcPrivateEndpointArrayOutput) ToVpcPrivateEndpointArrayOutputWithContext(ctx context.Context) VpcPrivateEndpointArrayOutput {
	return o
}

func (o VpcPrivateEndpointArrayOutput) Index(i pulumi.IntInput) VpcPrivateEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPrivateEndpoint {
		return vs[0].([]*VpcPrivateEndpoint)[vs[1].(int)]
	}).(VpcPrivateEndpointOutput)
}

type VpcPrivateEndpointMapOutput struct{ *pulumi.OutputState }

func (VpcPrivateEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPrivateEndpoint)(nil)).Elem()
}

func (o VpcPrivateEndpointMapOutput) ToVpcPrivateEndpointMapOutput() VpcPrivateEndpointMapOutput {
	return o
}

func (o VpcPrivateEndpointMapOutput) ToVpcPrivateEndpointMapOutputWithContext(ctx context.Context) VpcPrivateEndpointMapOutput {
	return o
}

func (o VpcPrivateEndpointMapOutput) MapIndex(k pulumi.StringInput) VpcPrivateEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPrivateEndpoint {
		return vs[0].(map[string]*VpcPrivateEndpoint)[vs[1].(string)]
	}).(VpcPrivateEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateEndpointInput)(nil)).Elem(), &VpcPrivateEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateEndpointArrayInput)(nil)).Elem(), VpcPrivateEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateEndpointMapInput)(nil)).Elem(), VpcPrivateEndpointMap{})
	pulumi.RegisterOutputType(VpcPrivateEndpointOutput{})
	pulumi.RegisterOutputType(VpcPrivateEndpointArrayOutput{})
	pulumi.RegisterOutputType(VpcPrivateEndpointMapOutput{})
}
