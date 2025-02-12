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

// Allows creation and management of a single binding within IAM policy for an existing Yandex KMS Asymmetric Encryption Key.
//
// > Roles controlled by `KmsAsymmetricEncryptionKeyIamBinding` should not be assigned using `KmsAsymmetricEncryptionKeyIamMember`.
//
// > When you delete `KmsAsymmetricEncryptionKeyIamBinding` resource, the roles can be deleted from other users within the folder as well. Be careful!
//
// ## Example Usage
//
// {{ tffile "examples/kms_asymmetric_encryption_key_iam_binding/r_kms_asymmetric_encryption_key_iam_binding_1.tf" }}
//
// ## Import
//
// IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `asymmetric_encryption_key_id` and role, e.g.
//
// ```sh
// $ pulumi import yandex:index/kmsAsymmetricEncryptionKeyIamBinding:KmsAsymmetricEncryptionKeyIamBinding viewer "asymmetric_encryption_key_id viewer"
// ```
type KmsAsymmetricEncryptionKeyIamBinding struct {
	pulumi.CustomResourceState

	// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringOutput `pulumi:"asymmetricEncryptionKeyId"`
	// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:group:federation:{federation_id}:users**: All users in federation.
	// * **system:group:organization:{organization_id}:users**: All users in organization.
	// * **system:allAuthenticatedUsers**: All authenticated users.
	// * **system:allUsers**: All users, including unauthenticated ones.
	//
	// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
	Role       pulumi.StringOutput `pulumi:"role"`
	SleepAfter pulumi.IntPtrOutput `pulumi:"sleepAfter"`
}

// NewKmsAsymmetricEncryptionKeyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewKmsAsymmetricEncryptionKeyIamBinding(ctx *pulumi.Context,
	name string, args *KmsAsymmetricEncryptionKeyIamBindingArgs, opts ...pulumi.ResourceOption) (*KmsAsymmetricEncryptionKeyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AsymmetricEncryptionKeyId == nil {
		return nil, errors.New("invalid value for required argument 'AsymmetricEncryptionKeyId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KmsAsymmetricEncryptionKeyIamBinding
	err := ctx.RegisterResource("yandex:index/kmsAsymmetricEncryptionKeyIamBinding:KmsAsymmetricEncryptionKeyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsAsymmetricEncryptionKeyIamBinding gets an existing KmsAsymmetricEncryptionKeyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsAsymmetricEncryptionKeyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsAsymmetricEncryptionKeyIamBindingState, opts ...pulumi.ResourceOption) (*KmsAsymmetricEncryptionKeyIamBinding, error) {
	var resource KmsAsymmetricEncryptionKeyIamBinding
	err := ctx.ReadResource("yandex:index/kmsAsymmetricEncryptionKeyIamBinding:KmsAsymmetricEncryptionKeyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsAsymmetricEncryptionKeyIamBinding resources.
type kmsAsymmetricEncryptionKeyIamBindingState struct {
	// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId *string `pulumi:"asymmetricEncryptionKeyId"`
	// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:group:federation:{federation_id}:users**: All users in federation.
	// * **system:group:organization:{organization_id}:users**: All users in organization.
	// * **system:allAuthenticatedUsers**: All authenticated users.
	// * **system:allUsers**: All users, including unauthenticated ones.
	//
	// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
	Members []string `pulumi:"members"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
	Role       *string `pulumi:"role"`
	SleepAfter *int    `pulumi:"sleepAfter"`
}

type KmsAsymmetricEncryptionKeyIamBindingState struct {
	// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:group:federation:{federation_id}:users**: All users in federation.
	// * **system:group:organization:{organization_id}:users**: All users in organization.
	// * **system:allAuthenticatedUsers**: All authenticated users.
	// * **system:allUsers**: All users, including unauthenticated ones.
	//
	// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayInput
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
	Role       pulumi.StringPtrInput
	SleepAfter pulumi.IntPtrInput
}

func (KmsAsymmetricEncryptionKeyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymmetricEncryptionKeyIamBindingState)(nil)).Elem()
}

type kmsAsymmetricEncryptionKeyIamBindingArgs struct {
	// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId string `pulumi:"asymmetricEncryptionKeyId"`
	// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:group:federation:{federation_id}:users**: All users in federation.
	// * **system:group:organization:{organization_id}:users**: All users in organization.
	// * **system:allAuthenticatedUsers**: All authenticated users.
	// * **system:allUsers**: All users, including unauthenticated ones.
	//
	// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
	Members []string `pulumi:"members"`
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
	Role       string `pulumi:"role"`
	SleepAfter *int   `pulumi:"sleepAfter"`
}

// The set of arguments for constructing a KmsAsymmetricEncryptionKeyIamBinding resource.
type KmsAsymmetricEncryptionKeyIamBindingArgs struct {
	// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyId pulumi.StringInput
	// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
	// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
	// * **serviceAccount:{service_account_id}**: A unique service account ID.
	// * **system:group:federation:{federation_id}:users**: All users in federation.
	// * **system:group:organization:{organization_id}:users**: All users in organization.
	// * **system:allAuthenticatedUsers**: All authenticated users.
	// * **system:allUsers**: All users, including unauthenticated ones.
	//
	// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
	Members pulumi.StringArrayInput
	// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
	Role       pulumi.StringInput
	SleepAfter pulumi.IntPtrInput
}

func (KmsAsymmetricEncryptionKeyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsAsymmetricEncryptionKeyIamBindingArgs)(nil)).Elem()
}

type KmsAsymmetricEncryptionKeyIamBindingInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamBindingOutput() KmsAsymmetricEncryptionKeyIamBindingOutput
	ToKmsAsymmetricEncryptionKeyIamBindingOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingOutput
}

func (*KmsAsymmetricEncryptionKeyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (i *KmsAsymmetricEncryptionKeyIamBinding) ToKmsAsymmetricEncryptionKeyIamBindingOutput() KmsAsymmetricEncryptionKeyIamBindingOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamBindingOutputWithContext(context.Background())
}

func (i *KmsAsymmetricEncryptionKeyIamBinding) ToKmsAsymmetricEncryptionKeyIamBindingOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamBindingOutput)
}

// KmsAsymmetricEncryptionKeyIamBindingArrayInput is an input type that accepts KmsAsymmetricEncryptionKeyIamBindingArray and KmsAsymmetricEncryptionKeyIamBindingArrayOutput values.
// You can construct a concrete instance of `KmsAsymmetricEncryptionKeyIamBindingArrayInput` via:
//
//	KmsAsymmetricEncryptionKeyIamBindingArray{ KmsAsymmetricEncryptionKeyIamBindingArgs{...} }
type KmsAsymmetricEncryptionKeyIamBindingArrayInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamBindingArrayOutput() KmsAsymmetricEncryptionKeyIamBindingArrayOutput
	ToKmsAsymmetricEncryptionKeyIamBindingArrayOutputWithContext(context.Context) KmsAsymmetricEncryptionKeyIamBindingArrayOutput
}

type KmsAsymmetricEncryptionKeyIamBindingArray []KmsAsymmetricEncryptionKeyIamBindingInput

func (KmsAsymmetricEncryptionKeyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (i KmsAsymmetricEncryptionKeyIamBindingArray) ToKmsAsymmetricEncryptionKeyIamBindingArrayOutput() KmsAsymmetricEncryptionKeyIamBindingArrayOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamBindingArrayOutputWithContext(context.Background())
}

func (i KmsAsymmetricEncryptionKeyIamBindingArray) ToKmsAsymmetricEncryptionKeyIamBindingArrayOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamBindingArrayOutput)
}

// KmsAsymmetricEncryptionKeyIamBindingMapInput is an input type that accepts KmsAsymmetricEncryptionKeyIamBindingMap and KmsAsymmetricEncryptionKeyIamBindingMapOutput values.
// You can construct a concrete instance of `KmsAsymmetricEncryptionKeyIamBindingMapInput` via:
//
//	KmsAsymmetricEncryptionKeyIamBindingMap{ "key": KmsAsymmetricEncryptionKeyIamBindingArgs{...} }
type KmsAsymmetricEncryptionKeyIamBindingMapInput interface {
	pulumi.Input

	ToKmsAsymmetricEncryptionKeyIamBindingMapOutput() KmsAsymmetricEncryptionKeyIamBindingMapOutput
	ToKmsAsymmetricEncryptionKeyIamBindingMapOutputWithContext(context.Context) KmsAsymmetricEncryptionKeyIamBindingMapOutput
}

type KmsAsymmetricEncryptionKeyIamBindingMap map[string]KmsAsymmetricEncryptionKeyIamBindingInput

func (KmsAsymmetricEncryptionKeyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (i KmsAsymmetricEncryptionKeyIamBindingMap) ToKmsAsymmetricEncryptionKeyIamBindingMapOutput() KmsAsymmetricEncryptionKeyIamBindingMapOutput {
	return i.ToKmsAsymmetricEncryptionKeyIamBindingMapOutputWithContext(context.Background())
}

func (i KmsAsymmetricEncryptionKeyIamBindingMap) ToKmsAsymmetricEncryptionKeyIamBindingMapOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsAsymmetricEncryptionKeyIamBindingMapOutput)
}

type KmsAsymmetricEncryptionKeyIamBindingOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamBindingOutput) ToKmsAsymmetricEncryptionKeyIamBindingOutput() KmsAsymmetricEncryptionKeyIamBindingOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamBindingOutput) ToKmsAsymmetricEncryptionKeyIamBindingOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingOutput {
	return o
}

// The [Yandex Key Management Service](https://cloud.yandex.com/docs/kms/) Asymmetric Encryption Key ID to apply a binding to.
func (o KmsAsymmetricEncryptionKeyIamBindingOutput) AsymmetricEncryptionKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamBinding) pulumi.StringOutput { return v.AsymmetricEncryptionKeyId }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`. Each entry can have one of the following values:
// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
// * **serviceAccount:{service_account_id}**: A unique service account ID.
// * **system:group:federation:{federation_id}:users**: All users in federation.
// * **system:group:organization:{organization_id}:users**: All users in organization.
// * **system:allAuthenticatedUsers**: All authenticated users.
// * **system:allUsers**: All users, including unauthenticated ones.
//
// Note: for more information about system groups, see the [documentation](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group).
func (o KmsAsymmetricEncryptionKeyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. See [roles](https://cloud.yandex.com/docs/kms/security/).
func (o KmsAsymmetricEncryptionKeyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o KmsAsymmetricEncryptionKeyIamBindingOutput) SleepAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KmsAsymmetricEncryptionKeyIamBinding) pulumi.IntPtrOutput { return v.SleepAfter }).(pulumi.IntPtrOutput)
}

type KmsAsymmetricEncryptionKeyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamBindingArrayOutput) ToKmsAsymmetricEncryptionKeyIamBindingArrayOutput() KmsAsymmetricEncryptionKeyIamBindingArrayOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamBindingArrayOutput) ToKmsAsymmetricEncryptionKeyIamBindingArrayOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingArrayOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamBindingArrayOutput) Index(i pulumi.IntInput) KmsAsymmetricEncryptionKeyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KmsAsymmetricEncryptionKeyIamBinding {
		return vs[0].([]*KmsAsymmetricEncryptionKeyIamBinding)[vs[1].(int)]
	}).(KmsAsymmetricEncryptionKeyIamBindingOutput)
}

type KmsAsymmetricEncryptionKeyIamBindingMapOutput struct{ *pulumi.OutputState }

func (KmsAsymmetricEncryptionKeyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KmsAsymmetricEncryptionKeyIamBinding)(nil)).Elem()
}

func (o KmsAsymmetricEncryptionKeyIamBindingMapOutput) ToKmsAsymmetricEncryptionKeyIamBindingMapOutput() KmsAsymmetricEncryptionKeyIamBindingMapOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamBindingMapOutput) ToKmsAsymmetricEncryptionKeyIamBindingMapOutputWithContext(ctx context.Context) KmsAsymmetricEncryptionKeyIamBindingMapOutput {
	return o
}

func (o KmsAsymmetricEncryptionKeyIamBindingMapOutput) MapIndex(k pulumi.StringInput) KmsAsymmetricEncryptionKeyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KmsAsymmetricEncryptionKeyIamBinding {
		return vs[0].(map[string]*KmsAsymmetricEncryptionKeyIamBinding)[vs[1].(string)]
	}).(KmsAsymmetricEncryptionKeyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamBindingInput)(nil)).Elem(), &KmsAsymmetricEncryptionKeyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamBindingArrayInput)(nil)).Elem(), KmsAsymmetricEncryptionKeyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KmsAsymmetricEncryptionKeyIamBindingMapInput)(nil)).Elem(), KmsAsymmetricEncryptionKeyIamBindingMap{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamBindingOutput{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(KmsAsymmetricEncryptionKeyIamBindingMapOutput{})
}
