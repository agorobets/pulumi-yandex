# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['LoggingGroupArgs', 'LoggingGroup']

@pulumi.input_type
class LoggingGroupArgs:
    def __init__(__self__, *,
                 data_stream: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoggingGroup resource.
        :param pulumi.Input[str] description: A description for the Yandex Cloud Logging group.
        :param pulumi.Input[str] folder_id: ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        :param pulumi.Input[str] name: Name for the Yandex Cloud Logging group.
        :param pulumi.Input[str] retention_period: Log entries retention period for the Yandex Cloud Logging group.
        """
        if data_stream is not None:
            pulumi.set(__self__, "data_stream", data_stream)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)

    @property
    @pulumi.getter(name="dataStream")
    def data_stream(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_stream")

    @data_stream.setter
    def data_stream(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_stream", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[str]]:
        """
        Log entries retention period for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention_period", value)


@pulumi.input_type
class _LoggingGroupState:
    def __init__(__self__, *,
                 cloud_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 data_stream: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoggingGroup resources.
        :param pulumi.Input[str] cloud_id: ID of the cloud that the Yandex Cloud Logging group belong to.
        :param pulumi.Input[str] created_at: The Yandex Cloud Logging group creation timestamp.
        :param pulumi.Input[str] description: A description for the Yandex Cloud Logging group.
        :param pulumi.Input[str] folder_id: ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        :param pulumi.Input[str] name: Name for the Yandex Cloud Logging group.
        :param pulumi.Input[str] retention_period: Log entries retention period for the Yandex Cloud Logging group.
        :param pulumi.Input[str] status: The Yandex Cloud Logging group status.
        """
        if cloud_id is not None:
            pulumi.set(__self__, "cloud_id", cloud_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if data_stream is not None:
            pulumi.set(__self__, "data_stream", data_stream)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cloudId")
    def cloud_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the cloud that the Yandex Cloud Logging group belong to.
        """
        return pulumi.get(self, "cloud_id")

    @cloud_id.setter
    def cloud_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The Yandex Cloud Logging group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dataStream")
    def data_stream(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_stream")

    @data_stream.setter
    def data_stream(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_stream", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[str]]:
        """
        Log entries retention period for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention_period", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Yandex Cloud Logging group status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class LoggingGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_stream: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Yandex Cloud Logging group resource. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).

        ## Example Usage

        {{ tffile "examples/logging_group/r_logging_group_1.tf" }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the Yandex Cloud Logging group.
        :param pulumi.Input[str] folder_id: ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        :param pulumi.Input[str] name: Name for the Yandex Cloud Logging group.
        :param pulumi.Input[str] retention_period: Log entries retention period for the Yandex Cloud Logging group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LoggingGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Yandex Cloud Logging group resource. For more information, see [the official documentation](https://yandex.cloud/docs/logging/concepts/log-group).

        ## Example Usage

        {{ tffile "examples/logging_group/r_logging_group_1.tf" }}

        :param str resource_name: The name of the resource.
        :param LoggingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoggingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_stream: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoggingGroupArgs.__new__(LoggingGroupArgs)

            __props__.__dict__["data_stream"] = data_stream
            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["retention_period"] = retention_period
            __props__.__dict__["cloud_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
        super(LoggingGroup, __self__).__init__(
            'yandex:index/loggingGroup:LoggingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cloud_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            data_stream: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_period: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'LoggingGroup':
        """
        Get an existing LoggingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_id: ID of the cloud that the Yandex Cloud Logging group belong to.
        :param pulumi.Input[str] created_at: The Yandex Cloud Logging group creation timestamp.
        :param pulumi.Input[str] description: A description for the Yandex Cloud Logging group.
        :param pulumi.Input[str] folder_id: ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        :param pulumi.Input[str] name: Name for the Yandex Cloud Logging group.
        :param pulumi.Input[str] retention_period: Log entries retention period for the Yandex Cloud Logging group.
        :param pulumi.Input[str] status: The Yandex Cloud Logging group status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoggingGroupState.__new__(_LoggingGroupState)

        __props__.__dict__["cloud_id"] = cloud_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["data_stream"] = data_stream
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["retention_period"] = retention_period
        __props__.__dict__["status"] = status
        return LoggingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudId")
    def cloud_id(self) -> pulumi.Output[str]:
        """
        ID of the cloud that the Yandex Cloud Logging group belong to.
        """
        return pulumi.get(self, "cloud_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The Yandex Cloud Logging group creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dataStream")
    def data_stream(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "data_stream")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        ID of the folder that the Yandex Cloud Logging group belongs to. It will be deduced from provider configuration if not set explicitly.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> pulumi.Output[str]:
        """
        Log entries retention period for the Yandex Cloud Logging group.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The Yandex Cloud Logging group status.
        """
        return pulumi.get(self, "status")

