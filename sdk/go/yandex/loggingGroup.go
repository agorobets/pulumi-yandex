// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Yandex Cloud Logging group resource. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).
//
// ## Example Usage
//
// {{ tffile "examples/logging_group/r_logging_group_1.tf" }}
type LoggingGroup struct {
	pulumi.CustomResourceState

	// ID of the cloud that the Yandex Cloud Logging group belong to.
	CloudId pulumi.StringOutput `pulumi:"cloudId"`
	// The Yandex Cloud Logging group creation timestamp.
	CreatedAt  pulumi.StringOutput    `pulumi:"createdAt"`
	DataStream pulumi.StringPtrOutput `pulumi:"dataStream"`
	// A description for the Yandex Cloud Logging group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name for the Yandex Cloud Logging group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod pulumi.StringOutput `pulumi:"retentionPeriod"`
	// The Yandex Cloud Logging group status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewLoggingGroup registers a new resource with the given unique name, arguments, and options.
func NewLoggingGroup(ctx *pulumi.Context,
	name string, args *LoggingGroupArgs, opts ...pulumi.ResourceOption) (*LoggingGroup, error) {
	if args == nil {
		args = &LoggingGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoggingGroup
	err := ctx.RegisterResource("yandex:index/loggingGroup:LoggingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingGroup gets an existing LoggingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingGroupState, opts ...pulumi.ResourceOption) (*LoggingGroup, error) {
	var resource LoggingGroup
	err := ctx.ReadResource("yandex:index/loggingGroup:LoggingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingGroup resources.
type loggingGroupState struct {
	// ID of the cloud that the Yandex Cloud Logging group belong to.
	CloudId *string `pulumi:"cloudId"`
	// The Yandex Cloud Logging group creation timestamp.
	CreatedAt  *string `pulumi:"createdAt"`
	DataStream *string `pulumi:"dataStream"`
	// A description for the Yandex Cloud Logging group.
	Description *string `pulumi:"description"`
	// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels map[string]string `pulumi:"labels"`
	// Name for the Yandex Cloud Logging group.
	Name *string `pulumi:"name"`
	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod *string `pulumi:"retentionPeriod"`
	// The Yandex Cloud Logging group status.
	Status *string `pulumi:"status"`
}

type LoggingGroupState struct {
	// ID of the cloud that the Yandex Cloud Logging group belong to.
	CloudId pulumi.StringPtrInput
	// The Yandex Cloud Logging group creation timestamp.
	CreatedAt  pulumi.StringPtrInput
	DataStream pulumi.StringPtrInput
	// A description for the Yandex Cloud Logging group.
	Description pulumi.StringPtrInput
	// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels pulumi.StringMapInput
	// Name for the Yandex Cloud Logging group.
	Name pulumi.StringPtrInput
	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod pulumi.StringPtrInput
	// The Yandex Cloud Logging group status.
	Status pulumi.StringPtrInput
}

func (LoggingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingGroupState)(nil)).Elem()
}

type loggingGroupArgs struct {
	DataStream *string `pulumi:"dataStream"`
	// A description for the Yandex Cloud Logging group.
	Description *string `pulumi:"description"`
	// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
	FolderId *string `pulumi:"folderId"`
	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels map[string]string `pulumi:"labels"`
	// Name for the Yandex Cloud Logging group.
	Name *string `pulumi:"name"`
	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod *string `pulumi:"retentionPeriod"`
}

// The set of arguments for constructing a LoggingGroup resource.
type LoggingGroupArgs struct {
	DataStream pulumi.StringPtrInput
	// A description for the Yandex Cloud Logging group.
	Description pulumi.StringPtrInput
	// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
	FolderId pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
	Labels pulumi.StringMapInput
	// Name for the Yandex Cloud Logging group.
	Name pulumi.StringPtrInput
	// Log entries retention period for the Yandex Cloud Logging group.
	RetentionPeriod pulumi.StringPtrInput
}

func (LoggingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingGroupArgs)(nil)).Elem()
}

type LoggingGroupInput interface {
	pulumi.Input

	ToLoggingGroupOutput() LoggingGroupOutput
	ToLoggingGroupOutputWithContext(ctx context.Context) LoggingGroupOutput
}

func (*LoggingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingGroup)(nil)).Elem()
}

func (i *LoggingGroup) ToLoggingGroupOutput() LoggingGroupOutput {
	return i.ToLoggingGroupOutputWithContext(context.Background())
}

func (i *LoggingGroup) ToLoggingGroupOutputWithContext(ctx context.Context) LoggingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingGroupOutput)
}

// LoggingGroupArrayInput is an input type that accepts LoggingGroupArray and LoggingGroupArrayOutput values.
// You can construct a concrete instance of `LoggingGroupArrayInput` via:
//
//	LoggingGroupArray{ LoggingGroupArgs{...} }
type LoggingGroupArrayInput interface {
	pulumi.Input

	ToLoggingGroupArrayOutput() LoggingGroupArrayOutput
	ToLoggingGroupArrayOutputWithContext(context.Context) LoggingGroupArrayOutput
}

type LoggingGroupArray []LoggingGroupInput

func (LoggingGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoggingGroup)(nil)).Elem()
}

func (i LoggingGroupArray) ToLoggingGroupArrayOutput() LoggingGroupArrayOutput {
	return i.ToLoggingGroupArrayOutputWithContext(context.Background())
}

func (i LoggingGroupArray) ToLoggingGroupArrayOutputWithContext(ctx context.Context) LoggingGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingGroupArrayOutput)
}

// LoggingGroupMapInput is an input type that accepts LoggingGroupMap and LoggingGroupMapOutput values.
// You can construct a concrete instance of `LoggingGroupMapInput` via:
//
//	LoggingGroupMap{ "key": LoggingGroupArgs{...} }
type LoggingGroupMapInput interface {
	pulumi.Input

	ToLoggingGroupMapOutput() LoggingGroupMapOutput
	ToLoggingGroupMapOutputWithContext(context.Context) LoggingGroupMapOutput
}

type LoggingGroupMap map[string]LoggingGroupInput

func (LoggingGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoggingGroup)(nil)).Elem()
}

func (i LoggingGroupMap) ToLoggingGroupMapOutput() LoggingGroupMapOutput {
	return i.ToLoggingGroupMapOutputWithContext(context.Background())
}

func (i LoggingGroupMap) ToLoggingGroupMapOutputWithContext(ctx context.Context) LoggingGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingGroupMapOutput)
}

type LoggingGroupOutput struct{ *pulumi.OutputState }

func (LoggingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingGroup)(nil)).Elem()
}

func (o LoggingGroupOutput) ToLoggingGroupOutput() LoggingGroupOutput {
	return o
}

func (o LoggingGroupOutput) ToLoggingGroupOutputWithContext(ctx context.Context) LoggingGroupOutput {
	return o
}

// ID of the cloud that the Yandex Cloud Logging group belong to.
func (o LoggingGroupOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.CloudId }).(pulumi.StringOutput)
}

// The Yandex Cloud Logging group creation timestamp.
func (o LoggingGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LoggingGroupOutput) DataStream() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringPtrOutput { return v.DataStream }).(pulumi.StringPtrOutput)
}

// A description for the Yandex Cloud Logging group.
func (o LoggingGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
func (o LoggingGroupOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the Yandex Cloud Logging group.
func (o LoggingGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name for the Yandex Cloud Logging group.
func (o LoggingGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Log entries retention period for the Yandex Cloud Logging group.
func (o LoggingGroupOutput) RetentionPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.RetentionPeriod }).(pulumi.StringOutput)
}

// The Yandex Cloud Logging group status.
func (o LoggingGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type LoggingGroupArrayOutput struct{ *pulumi.OutputState }

func (LoggingGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoggingGroup)(nil)).Elem()
}

func (o LoggingGroupArrayOutput) ToLoggingGroupArrayOutput() LoggingGroupArrayOutput {
	return o
}

func (o LoggingGroupArrayOutput) ToLoggingGroupArrayOutputWithContext(ctx context.Context) LoggingGroupArrayOutput {
	return o
}

func (o LoggingGroupArrayOutput) Index(i pulumi.IntInput) LoggingGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoggingGroup {
		return vs[0].([]*LoggingGroup)[vs[1].(int)]
	}).(LoggingGroupOutput)
}

type LoggingGroupMapOutput struct{ *pulumi.OutputState }

func (LoggingGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoggingGroup)(nil)).Elem()
}

func (o LoggingGroupMapOutput) ToLoggingGroupMapOutput() LoggingGroupMapOutput {
	return o
}

func (o LoggingGroupMapOutput) ToLoggingGroupMapOutputWithContext(ctx context.Context) LoggingGroupMapOutput {
	return o
}

func (o LoggingGroupMapOutput) MapIndex(k pulumi.StringInput) LoggingGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoggingGroup {
		return vs[0].(map[string]*LoggingGroup)[vs[1].(string)]
	}).(LoggingGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingGroupInput)(nil)).Elem(), &LoggingGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingGroupArrayInput)(nil)).Elem(), LoggingGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingGroupMapInput)(nil)).Elem(), LoggingGroupMap{})
	pulumi.RegisterOutputType(LoggingGroupOutput{})
	pulumi.RegisterOutputType(LoggingGroupArrayOutput{})
	pulumi.RegisterOutputType(LoggingGroupMapOutput{})
}
