// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/agorobets/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of [Yandex Cloud Storage Bucket](https://yandex.cloud/docs/storage/concepts/bucket).
//
// > By default, for authentication, you need to use [IAM token](https://yandex.cloud/docs/iam/concepts/authorization/iam-token) with the necessary permissions.
//
// > Alternatively, you can provide [static access keys](https://yandex.cloud/docs/iam/concepts/authorization/access-key) (Access and Secret). To generate these keys, you will need a Service Account with the appropriate permissions.
//
// > For extended API usage, such as setting the `maxSize`, `folderId`, `anonymousAccessFlags`, `defaultStorageClass`, and `https` parameters for a bucket,
// only the default authorization method will be used. This means the `IAM` token from the `provider` block will be applied.
// This can be confusing in cases where a separate service account is used for managing buckets because, in such scenarios,
// buckets may be accessed by two different accounts, each with potentially different permissions for the buckets.
//
// > In case you are using IAM token from UserAccount, you are needed to explicitly specify folderId in the resource,
// as it cannot be identified from such type of account. In case you are using IAM token from ServiceAccount or static access keys,
// folderId does not need to be specified unless you want to create the resource in a different folder than the account folder.
//
// ## Example Usage
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_1.tf" }}
//
// ### Simple Private Bucket With Static Access Keys
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_2.tf" }}
//
// ### Static Website Hosting
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_3.tf" }}
//
// ### Using ACL policy grants
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_4.tf" }}
//
// ### Using CORS
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_5.tf" }}
//
// ### Using versioning
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_6.tf" }}
//
// ### Using Object Lock Configuration
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_11.tf" }}
//
// ### Bucket Tagging
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_12.tf" }}
//
// ### Bucket Max Size
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_13.tf" }}
//
// ### Bucket Folder Id
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_14.tf" }}
//
// ### Bucket Anonymous Access Flags
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_15.tf" }}
//
// ### Bucket HTTPS Certificate
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_16.tf" }}
//
// ### Bucket Default Storage Class
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_17.tf" }}
//
// ### All settings example
//
// {{ tffile "examples/storage_bucket/r_storage_bucket_18.tf" }}
//
// ## Import
//
// Storage bucket can be imported using the `bucket`, e.g.
//
// ```sh
// $ pulumi import yandex:index/storageBucket:StorageBucket bucket bucket-name
// ```
// ~> Terraform will import this resource with `force_destroy` set to `false` in state. If you've set it to `true` in config, run `pulumi up` to update the value set in state. If you delete this resource before updating the value, objects in the bucket will not be destroyed.
type StorageBucket struct {
	pulumi.CustomResourceState

	// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	//
	// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
	Acl                  pulumi.StringPtrOutput                  `pulumi:"acl"`
	AnonymousAccessFlags StorageBucketAnonymousAccessFlagsOutput `pulumi:"anonymousAccessFlags"`
	Bucket               pulumi.StringOutput                     `pulumi:"bucket"`
	// The bucket domain name.
	BucketDomainName pulumi.StringOutput `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
	CorsRules           StorageBucketCorsRuleArrayOutput `pulumi:"corsRules"`
	DefaultStorageClass pulumi.StringOutput              `pulumi:"defaultStorageClass"`
	FolderId            pulumi.StringOutput              `pulumi:"folderId"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayOutput `pulumi:"grants"`
	Https  StorageBucketHttpsPtrOutput   `pulumi:"https"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayOutput `pulumi:"loggings"`
	MaxSize  pulumi.IntPtrOutput             `pulumi:"maxSize"`
	// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
	ObjectLockConfiguration StorageBucketObjectLockConfigurationPtrOutput `pulumi:"objectLockConfiguration"`
	Policy                  pulumi.StringPtrOutput                        `pulumi:"policy"`
	// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	SecretKey                         pulumi.StringPtrOutput                                  `pulumi:"secretKey"`
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrOutput `pulumi:"serverSideEncryptionConfiguration"`
	Tags                              pulumi.StringMapOutput                                  `pulumi:"tags"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	//
	// > To manage `versioning` argument, service account with `storage.admin` role should be used.
	Versioning StorageBucketVersioningOutput `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrOutput `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringOutput `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringOutput `pulumi:"websiteEndpoint"`
}

// NewStorageBucket registers a new resource with the given unique name, arguments, and options.
func NewStorageBucket(ctx *pulumi.Context,
	name string, args *StorageBucketArgs, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	if args == nil {
		args = &StorageBucketArgs{}
	}

	if args.SecretKey != nil {
		args.SecretKey = pulumi.ToSecret(args.SecretKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StorageBucket
	err := ctx.RegisterResource("yandex:index/storageBucket:StorageBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucket gets an existing StorageBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketState, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	var resource StorageBucket
	err := ctx.ReadResource("yandex:index/storageBucket:StorageBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucket resources.
type storageBucketState struct {
	// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	AccessKey *string `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	//
	// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
	Acl                  *string                            `pulumi:"acl"`
	AnonymousAccessFlags *StorageBucketAnonymousAccessFlags `pulumi:"anonymousAccessFlags"`
	Bucket               *string                            `pulumi:"bucket"`
	// The bucket domain name.
	BucketDomainName *string `pulumi:"bucketDomainName"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
	CorsRules           []StorageBucketCorsRule `pulumi:"corsRules"`
	DefaultStorageClass *string                 `pulumi:"defaultStorageClass"`
	FolderId            *string                 `pulumi:"folderId"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants []StorageBucketGrant `pulumi:"grants"`
	Https  *StorageBucketHttps  `pulumi:"https"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules []StorageBucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings []StorageBucketLogging `pulumi:"loggings"`
	MaxSize  *int                   `pulumi:"maxSize"`
	// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
	ObjectLockConfiguration *StorageBucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	Policy                  *string                               `pulumi:"policy"`
	// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	SecretKey                         *string                                         `pulumi:"secretKey"`
	ServerSideEncryptionConfiguration *StorageBucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	Tags                              map[string]string                               `pulumi:"tags"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	//
	// > To manage `versioning` argument, service account with `storage.admin` role should be used.
	Versioning *StorageBucketVersioning `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website *StorageBucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

type StorageBucketState struct {
	// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	AccessKey pulumi.StringPtrInput
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	//
	// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
	Acl                  pulumi.StringPtrInput
	AnonymousAccessFlags StorageBucketAnonymousAccessFlagsPtrInput
	Bucket               pulumi.StringPtrInput
	// The bucket domain name.
	BucketDomainName pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
	CorsRules           StorageBucketCorsRuleArrayInput
	DefaultStorageClass pulumi.StringPtrInput
	FolderId            pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayInput
	Https  StorageBucketHttpsPtrInput
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayInput
	MaxSize  pulumi.IntPtrInput
	// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
	ObjectLockConfiguration StorageBucketObjectLockConfigurationPtrInput
	Policy                  pulumi.StringPtrInput
	// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	SecretKey                         pulumi.StringPtrInput
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrInput
	Tags                              pulumi.StringMapInput
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	//
	// > To manage `versioning` argument, service account with `storage.admin` role should be used.
	Versioning StorageBucketVersioningPtrInput
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (StorageBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketState)(nil)).Elem()
}

type storageBucketArgs struct {
	// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	AccessKey *string `pulumi:"accessKey"`
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	//
	// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
	Acl                  *string                            `pulumi:"acl"`
	AnonymousAccessFlags *StorageBucketAnonymousAccessFlags `pulumi:"anonymousAccessFlags"`
	Bucket               *string                            `pulumi:"bucket"`
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
	CorsRules           []StorageBucketCorsRule `pulumi:"corsRules"`
	DefaultStorageClass *string                 `pulumi:"defaultStorageClass"`
	FolderId            *string                 `pulumi:"folderId"`
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants []StorageBucketGrant `pulumi:"grants"`
	Https  *StorageBucketHttps  `pulumi:"https"`
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules []StorageBucketLifecycleRule `pulumi:"lifecycleRules"`
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings []StorageBucketLogging `pulumi:"loggings"`
	MaxSize  *int                   `pulumi:"maxSize"`
	// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
	ObjectLockConfiguration *StorageBucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	Policy                  *string                               `pulumi:"policy"`
	// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	SecretKey                         *string                                         `pulumi:"secretKey"`
	ServerSideEncryptionConfiguration *StorageBucketServerSideEncryptionConfiguration `pulumi:"serverSideEncryptionConfiguration"`
	Tags                              map[string]string                               `pulumi:"tags"`
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	//
	// > To manage `versioning` argument, service account with `storage.admin` role should be used.
	Versioning *StorageBucketVersioning `pulumi:"versioning"`
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website *StorageBucketWebsite `pulumi:"website"`
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain *string `pulumi:"websiteDomain"`
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `pulumi:"websiteEndpoint"`
}

// The set of arguments for constructing a StorageBucket resource.
type StorageBucketArgs struct {
	// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	AccessKey pulumi.StringPtrInput
	// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	//
	// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
	Acl                  pulumi.StringPtrInput
	AnonymousAccessFlags StorageBucketAnonymousAccessFlagsPtrInput
	Bucket               pulumi.StringPtrInput
	// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
	CorsRules           StorageBucketCorsRuleArrayInput
	DefaultStorageClass pulumi.StringPtrInput
	FolderId            pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grants StorageBucketGrantArrayInput
	Https  StorageBucketHttpsPtrInput
	// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRules StorageBucketLifecycleRuleArrayInput
	// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Loggings StorageBucketLoggingArrayInput
	MaxSize  pulumi.IntPtrInput
	// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
	ObjectLockConfiguration StorageBucketObjectLockConfigurationPtrInput
	Policy                  pulumi.StringPtrInput
	// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
	SecretKey                         pulumi.StringPtrInput
	ServerSideEncryptionConfiguration StorageBucketServerSideEncryptionConfigurationPtrInput
	Tags                              pulumi.StringMapInput
	// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	//
	// > To manage `versioning` argument, service account with `storage.admin` role should be used.
	Versioning StorageBucketVersioningPtrInput
	// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website StorageBucketWebsitePtrInput
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain pulumi.StringPtrInput
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint pulumi.StringPtrInput
}

func (StorageBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketArgs)(nil)).Elem()
}

type StorageBucketInput interface {
	pulumi.Input

	ToStorageBucketOutput() StorageBucketOutput
	ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput
}

func (*StorageBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil)).Elem()
}

func (i *StorageBucket) ToStorageBucketOutput() StorageBucketOutput {
	return i.ToStorageBucketOutputWithContext(context.Background())
}

func (i *StorageBucket) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketOutput)
}

// StorageBucketArrayInput is an input type that accepts StorageBucketArray and StorageBucketArrayOutput values.
// You can construct a concrete instance of `StorageBucketArrayInput` via:
//
//	StorageBucketArray{ StorageBucketArgs{...} }
type StorageBucketArrayInput interface {
	pulumi.Input

	ToStorageBucketArrayOutput() StorageBucketArrayOutput
	ToStorageBucketArrayOutputWithContext(context.Context) StorageBucketArrayOutput
}

type StorageBucketArray []StorageBucketInput

func (StorageBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketArray) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return i.ToStorageBucketArrayOutputWithContext(context.Background())
}

func (i StorageBucketArray) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketArrayOutput)
}

// StorageBucketMapInput is an input type that accepts StorageBucketMap and StorageBucketMapOutput values.
// You can construct a concrete instance of `StorageBucketMapInput` via:
//
//	StorageBucketMap{ "key": StorageBucketArgs{...} }
type StorageBucketMapInput interface {
	pulumi.Input

	ToStorageBucketMapOutput() StorageBucketMapOutput
	ToStorageBucketMapOutputWithContext(context.Context) StorageBucketMapOutput
}

type StorageBucketMap map[string]StorageBucketInput

func (StorageBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketMap) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return i.ToStorageBucketMapOutputWithContext(context.Background())
}

func (i StorageBucketMap) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketMapOutput)
}

type StorageBucketOutput struct{ *pulumi.OutputState }

func (StorageBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil)).Elem()
}

func (o StorageBucketOutput) ToStorageBucketOutput() StorageBucketOutput {
	return o
}

func (o StorageBucketOutput) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return o
}

// The access key to use when applying changes. This value can also be provided as `storageAccessKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
func (o StorageBucketOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
//
// > To change ACL after creation, service account with `storage.admin` role should be used, though this role is not necessary to create a bucket with any ACL.
func (o StorageBucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

func (o StorageBucketOutput) AnonymousAccessFlags() StorageBucketAnonymousAccessFlagsOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketAnonymousAccessFlagsOutput { return v.AnonymousAccessFlags }).(StorageBucketAnonymousAccessFlagsOutput)
}

func (o StorageBucketOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The bucket domain name.
func (o StorageBucketOutput) BucketDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.BucketDomainName }).(pulumi.StringOutput)
}

// Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
func (o StorageBucketOutput) BucketPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.BucketPrefix }).(pulumi.StringPtrOutput)
}

// A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/concepts/cors) (documented below).
func (o StorageBucketOutput) CorsRules() StorageBucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketCorsRuleArrayOutput { return v.CorsRules }).(StorageBucketCorsRuleArrayOutput)
}

func (o StorageBucketOutput) DefaultStorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.DefaultStorageClass }).(pulumi.StringOutput)
}

func (o StorageBucketOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
func (o StorageBucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
func (o StorageBucketOutput) Grants() StorageBucketGrantArrayOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketGrantArrayOutput { return v.Grants }).(StorageBucketGrantArrayOutput)
}

func (o StorageBucketOutput) Https() StorageBucketHttpsPtrOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketHttpsPtrOutput { return v.Https }).(StorageBucketHttpsPtrOutput)
}

// A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
func (o StorageBucketOutput) LifecycleRules() StorageBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(StorageBucketLifecycleRuleArrayOutput)
}

// A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
func (o StorageBucketOutput) Loggings() StorageBucketLoggingArrayOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketLoggingArrayOutput { return v.Loggings }).(StorageBucketLoggingArrayOutput)
}

func (o StorageBucketOutput) MaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.IntPtrOutput { return v.MaxSize }).(pulumi.IntPtrOutput)
}

// A configuration of [object lock management](https://yandex.cloud/docs/storage/concepts/object-lock) (documented below).
func (o StorageBucketOutput) ObjectLockConfiguration() StorageBucketObjectLockConfigurationPtrOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketObjectLockConfigurationPtrOutput { return v.ObjectLockConfiguration }).(StorageBucketObjectLockConfigurationPtrOutput)
}

func (o StorageBucketOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// The secret key to use when applying changes. This value can also be provided as `storageSecretKey` specified in provider config (explicitly or within `sharedCredentialsFile`) is used.
func (o StorageBucketOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.SecretKey }).(pulumi.StringPtrOutput)
}

func (o StorageBucketOutput) ServerSideEncryptionConfiguration() StorageBucketServerSideEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketServerSideEncryptionConfigurationPtrOutput {
		return v.ServerSideEncryptionConfiguration
	}).(StorageBucketServerSideEncryptionConfigurationPtrOutput)
}

func (o StorageBucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
//
// > To manage `versioning` argument, service account with `storage.admin` role should be used.
func (o StorageBucketOutput) Versioning() StorageBucketVersioningOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketVersioningOutput { return v.Versioning }).(StorageBucketVersioningOutput)
}

// A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
func (o StorageBucketOutput) Website() StorageBucketWebsitePtrOutput {
	return o.ApplyT(func(v *StorageBucket) StorageBucketWebsitePtrOutput { return v.Website }).(StorageBucketWebsitePtrOutput)
}

// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
func (o StorageBucketOutput) WebsiteDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.WebsiteDomain }).(pulumi.StringOutput)
}

// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
func (o StorageBucketOutput) WebsiteEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.WebsiteEndpoint }).(pulumi.StringOutput)
}

type StorageBucketArrayOutput struct{ *pulumi.OutputState }

func (StorageBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBucket)(nil)).Elem()
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) Index(i pulumi.IntInput) StorageBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageBucket {
		return vs[0].([]*StorageBucket)[vs[1].(int)]
	}).(StorageBucketOutput)
}

type StorageBucketMapOutput struct{ *pulumi.OutputState }

func (StorageBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBucket)(nil)).Elem()
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) MapIndex(k pulumi.StringInput) StorageBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageBucket {
		return vs[0].(map[string]*StorageBucket)[vs[1].(string)]
	}).(StorageBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketInput)(nil)).Elem(), &StorageBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketArrayInput)(nil)).Elem(), StorageBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketMapInput)(nil)).Elem(), StorageBucketMap{})
	pulumi.RegisterOutputType(StorageBucketOutput{})
	pulumi.RegisterOutputType(StorageBucketArrayOutput{})
	pulumi.RegisterOutputType(StorageBucketMapOutput{})
}
