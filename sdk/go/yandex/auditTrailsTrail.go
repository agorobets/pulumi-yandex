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

// Allows management of [trail](https://yandex.cloud/docs/audit-trails/concepts/trail)
//
// ## Example Usage
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_1.tf" }}
//
// Trail delivering events to YDS and gathering such events:
//
// * Management events from the 'some-organization' organization
// * DNS data events from the 'some-organization' organization
// * Object Storage data events from the 'some-organization' organization
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_2.tf" }}
//
// Trail delivering events to Object Storage and gathering such events:
//
// * Management events from the 'home-folder' folder
// * Managed PostgreSQL data events from the 'home-folder' folder
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_3.tf" }}
//
// ## Migration from deprecated filter field
//
// In order to migrate from unsing `filter` to the `filteringPolicy`, you will have to:
//
// * Remove the `filter.event_filters.categories` blocks. With the introduction of `includedEvents`/`excludedEvents` you can configure filtering per each event type.
//
// * Replace the `filter.event_filters.path_filter` with the appropriate `resourceScope` blocks. You have to account that `resourceScope` does not support specifying relations between resources, so your configuration will simplify to only the actual resources, that will be monitored.
//
// # Before
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_4.tf" }}
//
// # After
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_5.tf" }}
//
// * Replace the `filter.path_filter` block with the `filtering_policy.management_events_filter`. New API states management events filtration in a more clear way. The resources, that were specified, must migrate into the `filtering_policy.management_events_filter.resource_scope`
//
// # Before
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_6.tf" }}
//
// # After
//
// {{ tffile "examples/audit_trails_trail/r_audit_trails_trail_7.tf" }}
//
// ## Import
//
// A trail can be imported using the `id` of the resource, e.g.
//
// bash
//
// ```sh
// $ pulumi import yandex:index/auditTrailsTrail:AuditTrailsTrail infosec-trail trail_id
// ```
type AuditTrailsTrail struct {
	pulumi.CustomResourceState

	// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
	DataStreamDestination AuditTrailsTrailDataStreamDestinationPtrOutput `pulumi:"dataStreamDestination"`
	// Description of the trail.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Structure describing event filtering process for the trail.
	//
	// Deprecated: Configure filteringPolicy instead. This attribute will be removed
	Filter          AuditTrailsTrailFilterPtrOutput          `pulumi:"filter"`
	FilteringPolicy AuditTrailsTrailFilteringPolicyPtrOutput `pulumi:"filteringPolicy"`
	// ID of the folder to which the trail belongs.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Labels defined by the user.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
	LoggingDestination AuditTrailsTrailLoggingDestinationPtrOutput `pulumi:"loggingDestination"`
	// Name of the trail.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// Status of this trail.
	Status pulumi.StringOutput `pulumi:"status"`
	// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
	StorageDestination AuditTrailsTrailStorageDestinationPtrOutput `pulumi:"storageDestination"`
	// ID of the trail resource.
	TrailId pulumi.StringOutput `pulumi:"trailId"`
}

// NewAuditTrailsTrail registers a new resource with the given unique name, arguments, and options.
func NewAuditTrailsTrail(ctx *pulumi.Context,
	name string, args *AuditTrailsTrailArgs, opts ...pulumi.ResourceOption) (*AuditTrailsTrail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuditTrailsTrail
	err := ctx.RegisterResource("yandex:index/auditTrailsTrail:AuditTrailsTrail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditTrailsTrail gets an existing AuditTrailsTrail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditTrailsTrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditTrailsTrailState, opts ...pulumi.ResourceOption) (*AuditTrailsTrail, error) {
	var resource AuditTrailsTrail
	err := ctx.ReadResource("yandex:index/auditTrailsTrail:AuditTrailsTrail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditTrailsTrail resources.
type auditTrailsTrailState struct {
	// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
	DataStreamDestination *AuditTrailsTrailDataStreamDestination `pulumi:"dataStreamDestination"`
	// Description of the trail.
	Description *string `pulumi:"description"`
	// Structure describing event filtering process for the trail.
	//
	// Deprecated: Configure filteringPolicy instead. This attribute will be removed
	Filter          *AuditTrailsTrailFilter          `pulumi:"filter"`
	FilteringPolicy *AuditTrailsTrailFilteringPolicy `pulumi:"filteringPolicy"`
	// ID of the folder to which the trail belongs.
	FolderId *string `pulumi:"folderId"`
	// Labels defined by the user.
	Labels map[string]string `pulumi:"labels"`
	// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
	LoggingDestination *AuditTrailsTrailLoggingDestination `pulumi:"loggingDestination"`
	// Name of the trail.
	Name *string `pulumi:"name"`
	// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// Status of this trail.
	Status *string `pulumi:"status"`
	// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
	StorageDestination *AuditTrailsTrailStorageDestination `pulumi:"storageDestination"`
	// ID of the trail resource.
	TrailId *string `pulumi:"trailId"`
}

type AuditTrailsTrailState struct {
	// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
	DataStreamDestination AuditTrailsTrailDataStreamDestinationPtrInput
	// Description of the trail.
	Description pulumi.StringPtrInput
	// Structure describing event filtering process for the trail.
	//
	// Deprecated: Configure filteringPolicy instead. This attribute will be removed
	Filter          AuditTrailsTrailFilterPtrInput
	FilteringPolicy AuditTrailsTrailFilteringPolicyPtrInput
	// ID of the folder to which the trail belongs.
	FolderId pulumi.StringPtrInput
	// Labels defined by the user.
	Labels pulumi.StringMapInput
	// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
	LoggingDestination AuditTrailsTrailLoggingDestinationPtrInput
	// Name of the trail.
	Name pulumi.StringPtrInput
	// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
	ServiceAccountId pulumi.StringPtrInput
	// Status of this trail.
	Status pulumi.StringPtrInput
	// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
	StorageDestination AuditTrailsTrailStorageDestinationPtrInput
	// ID of the trail resource.
	TrailId pulumi.StringPtrInput
}

func (AuditTrailsTrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditTrailsTrailState)(nil)).Elem()
}

type auditTrailsTrailArgs struct {
	// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
	DataStreamDestination *AuditTrailsTrailDataStreamDestination `pulumi:"dataStreamDestination"`
	// Description of the trail.
	Description *string `pulumi:"description"`
	// Structure describing event filtering process for the trail.
	//
	// Deprecated: Configure filteringPolicy instead. This attribute will be removed
	Filter          *AuditTrailsTrailFilter          `pulumi:"filter"`
	FilteringPolicy *AuditTrailsTrailFilteringPolicy `pulumi:"filteringPolicy"`
	// ID of the folder to which the trail belongs.
	FolderId string `pulumi:"folderId"`
	// Labels defined by the user.
	Labels map[string]string `pulumi:"labels"`
	// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
	LoggingDestination *AuditTrailsTrailLoggingDestination `pulumi:"loggingDestination"`
	// Name of the trail.
	Name *string `pulumi:"name"`
	// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
	StorageDestination *AuditTrailsTrailStorageDestination `pulumi:"storageDestination"`
}

// The set of arguments for constructing a AuditTrailsTrail resource.
type AuditTrailsTrailArgs struct {
	// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
	DataStreamDestination AuditTrailsTrailDataStreamDestinationPtrInput
	// Description of the trail.
	Description pulumi.StringPtrInput
	// Structure describing event filtering process for the trail.
	//
	// Deprecated: Configure filteringPolicy instead. This attribute will be removed
	Filter          AuditTrailsTrailFilterPtrInput
	FilteringPolicy AuditTrailsTrailFilteringPolicyPtrInput
	// ID of the folder to which the trail belongs.
	FolderId pulumi.StringInput
	// Labels defined by the user.
	Labels pulumi.StringMapInput
	// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
	LoggingDestination AuditTrailsTrailLoggingDestinationPtrInput
	// Name of the trail.
	Name pulumi.StringPtrInput
	// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
	ServiceAccountId pulumi.StringInput
	// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
	StorageDestination AuditTrailsTrailStorageDestinationPtrInput
}

func (AuditTrailsTrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditTrailsTrailArgs)(nil)).Elem()
}

type AuditTrailsTrailInput interface {
	pulumi.Input

	ToAuditTrailsTrailOutput() AuditTrailsTrailOutput
	ToAuditTrailsTrailOutputWithContext(ctx context.Context) AuditTrailsTrailOutput
}

func (*AuditTrailsTrail) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditTrailsTrail)(nil)).Elem()
}

func (i *AuditTrailsTrail) ToAuditTrailsTrailOutput() AuditTrailsTrailOutput {
	return i.ToAuditTrailsTrailOutputWithContext(context.Background())
}

func (i *AuditTrailsTrail) ToAuditTrailsTrailOutputWithContext(ctx context.Context) AuditTrailsTrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailsTrailOutput)
}

// AuditTrailsTrailArrayInput is an input type that accepts AuditTrailsTrailArray and AuditTrailsTrailArrayOutput values.
// You can construct a concrete instance of `AuditTrailsTrailArrayInput` via:
//
//	AuditTrailsTrailArray{ AuditTrailsTrailArgs{...} }
type AuditTrailsTrailArrayInput interface {
	pulumi.Input

	ToAuditTrailsTrailArrayOutput() AuditTrailsTrailArrayOutput
	ToAuditTrailsTrailArrayOutputWithContext(context.Context) AuditTrailsTrailArrayOutput
}

type AuditTrailsTrailArray []AuditTrailsTrailInput

func (AuditTrailsTrailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditTrailsTrail)(nil)).Elem()
}

func (i AuditTrailsTrailArray) ToAuditTrailsTrailArrayOutput() AuditTrailsTrailArrayOutput {
	return i.ToAuditTrailsTrailArrayOutputWithContext(context.Background())
}

func (i AuditTrailsTrailArray) ToAuditTrailsTrailArrayOutputWithContext(ctx context.Context) AuditTrailsTrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailsTrailArrayOutput)
}

// AuditTrailsTrailMapInput is an input type that accepts AuditTrailsTrailMap and AuditTrailsTrailMapOutput values.
// You can construct a concrete instance of `AuditTrailsTrailMapInput` via:
//
//	AuditTrailsTrailMap{ "key": AuditTrailsTrailArgs{...} }
type AuditTrailsTrailMapInput interface {
	pulumi.Input

	ToAuditTrailsTrailMapOutput() AuditTrailsTrailMapOutput
	ToAuditTrailsTrailMapOutputWithContext(context.Context) AuditTrailsTrailMapOutput
}

type AuditTrailsTrailMap map[string]AuditTrailsTrailInput

func (AuditTrailsTrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditTrailsTrail)(nil)).Elem()
}

func (i AuditTrailsTrailMap) ToAuditTrailsTrailMapOutput() AuditTrailsTrailMapOutput {
	return i.ToAuditTrailsTrailMapOutputWithContext(context.Background())
}

func (i AuditTrailsTrailMap) ToAuditTrailsTrailMapOutputWithContext(ctx context.Context) AuditTrailsTrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditTrailsTrailMapOutput)
}

type AuditTrailsTrailOutput struct{ *pulumi.OutputState }

func (AuditTrailsTrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditTrailsTrail)(nil)).Elem()
}

func (o AuditTrailsTrailOutput) ToAuditTrailsTrailOutput() AuditTrailsTrailOutput {
	return o
}

func (o AuditTrailsTrailOutput) ToAuditTrailsTrailOutputWithContext(ctx context.Context) AuditTrailsTrailOutput {
	return o
}

// Structure describing destination data stream of the trail. Mutually exclusive with `loggingDestination` and `storageDestination`.
func (o AuditTrailsTrailOutput) DataStreamDestination() AuditTrailsTrailDataStreamDestinationPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) AuditTrailsTrailDataStreamDestinationPtrOutput {
		return v.DataStreamDestination
	}).(AuditTrailsTrailDataStreamDestinationPtrOutput)
}

// Description of the trail.
func (o AuditTrailsTrailOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Structure describing event filtering process for the trail.
//
// Deprecated: Configure filteringPolicy instead. This attribute will be removed
func (o AuditTrailsTrailOutput) Filter() AuditTrailsTrailFilterPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) AuditTrailsTrailFilterPtrOutput { return v.Filter }).(AuditTrailsTrailFilterPtrOutput)
}

func (o AuditTrailsTrailOutput) FilteringPolicy() AuditTrailsTrailFilteringPolicyPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) AuditTrailsTrailFilteringPolicyPtrOutput { return v.FilteringPolicy }).(AuditTrailsTrailFilteringPolicyPtrOutput)
}

// ID of the folder to which the trail belongs.
func (o AuditTrailsTrailOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Labels defined by the user.
func (o AuditTrailsTrailOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Structure describing destination log group of the trail. Mutually exclusive with `storageDestination` and `dataStreamDestination`.
func (o AuditTrailsTrailOutput) LoggingDestination() AuditTrailsTrailLoggingDestinationPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) AuditTrailsTrailLoggingDestinationPtrOutput { return v.LoggingDestination }).(AuditTrailsTrailLoggingDestinationPtrOutput)
}

// Name of the trail.
func (o AuditTrailsTrailOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the [IAM service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) that is used by the trail.
func (o AuditTrailsTrailOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// Status of this trail.
func (o AuditTrailsTrailOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Structure describing destination bucket of the trail. Mutually exclusive with `loggingDestination` and `dataStreamDestination`.
func (o AuditTrailsTrailOutput) StorageDestination() AuditTrailsTrailStorageDestinationPtrOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) AuditTrailsTrailStorageDestinationPtrOutput { return v.StorageDestination }).(AuditTrailsTrailStorageDestinationPtrOutput)
}

// ID of the trail resource.
func (o AuditTrailsTrailOutput) TrailId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditTrailsTrail) pulumi.StringOutput { return v.TrailId }).(pulumi.StringOutput)
}

type AuditTrailsTrailArrayOutput struct{ *pulumi.OutputState }

func (AuditTrailsTrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditTrailsTrail)(nil)).Elem()
}

func (o AuditTrailsTrailArrayOutput) ToAuditTrailsTrailArrayOutput() AuditTrailsTrailArrayOutput {
	return o
}

func (o AuditTrailsTrailArrayOutput) ToAuditTrailsTrailArrayOutputWithContext(ctx context.Context) AuditTrailsTrailArrayOutput {
	return o
}

func (o AuditTrailsTrailArrayOutput) Index(i pulumi.IntInput) AuditTrailsTrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuditTrailsTrail {
		return vs[0].([]*AuditTrailsTrail)[vs[1].(int)]
	}).(AuditTrailsTrailOutput)
}

type AuditTrailsTrailMapOutput struct{ *pulumi.OutputState }

func (AuditTrailsTrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditTrailsTrail)(nil)).Elem()
}

func (o AuditTrailsTrailMapOutput) ToAuditTrailsTrailMapOutput() AuditTrailsTrailMapOutput {
	return o
}

func (o AuditTrailsTrailMapOutput) ToAuditTrailsTrailMapOutputWithContext(ctx context.Context) AuditTrailsTrailMapOutput {
	return o
}

func (o AuditTrailsTrailMapOutput) MapIndex(k pulumi.StringInput) AuditTrailsTrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuditTrailsTrail {
		return vs[0].(map[string]*AuditTrailsTrail)[vs[1].(string)]
	}).(AuditTrailsTrailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailsTrailInput)(nil)).Elem(), &AuditTrailsTrail{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailsTrailArrayInput)(nil)).Elem(), AuditTrailsTrailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditTrailsTrailMapInput)(nil)).Elem(), AuditTrailsTrailMap{})
	pulumi.RegisterOutputType(AuditTrailsTrailOutput{})
	pulumi.RegisterOutputType(AuditTrailsTrailArrayOutput{})
	pulumi.RegisterOutputType(AuditTrailsTrailMapOutput{})
}
