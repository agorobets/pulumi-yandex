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

__all__ = ['DnsZoneArgs', 'DnsZone']

@pulumi.input_type
class DnsZoneArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 public: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DnsZone resource.
        :param pulumi.Input[str] zone: The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        :param pulumi.Input[bool] deletion_protection: Flag that protects the dns zone from accidental deletion.
        :param pulumi.Input[str] description: Description of the DNS zone.
        :param pulumi.Input[str] folder_id: ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the DNS zone.
        :param pulumi.Input[str] name: User assigned name of a specific resource. Must be unique within the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_networks: For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        :param pulumi.Input[bool] public: The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
               Cloud resources.
        """
        pulumi.set(__self__, "zone", zone)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_networks is not None:
            pulumi.set(__self__, "private_networks", private_networks)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag that protects the dns zone from accidental deletion.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the DNS zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the DNS zone.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User assigned name of a specific resource. Must be unique within the folder.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        """
        return pulumi.get(self, "private_networks")

    @private_networks.setter
    def private_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "private_networks", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
        Cloud resources.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)


@pulumi.input_type
class _DnsZoneState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DnsZone resources.
        :param pulumi.Input[str] created_at: The DNS zone creation timestamp.
        :param pulumi.Input[bool] deletion_protection: Flag that protects the dns zone from accidental deletion.
        :param pulumi.Input[str] description: Description of the DNS zone.
        :param pulumi.Input[str] folder_id: ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the DNS zone.
        :param pulumi.Input[str] name: User assigned name of a specific resource. Must be unique within the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_networks: For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        :param pulumi.Input[bool] public: The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
               Cloud resources.
        :param pulumi.Input[str] zone: The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_networks is not None:
            pulumi.set(__self__, "private_networks", private_networks)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS zone creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag that protects the dns zone from accidental deletion.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the DNS zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the DNS zone.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User assigned name of a specific resource. Must be unique within the folder.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        """
        return pulumi.get(self, "private_networks")

    @private_networks.setter
    def private_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "private_networks", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
        Cloud resources.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class DnsZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        {{ .Description }}

        ## Example Usage

        {{ tffile "examples/dns_zone/r_dns_zone_1.tf" }}

        {{ .SchemaMarkdown }}

        ## Import

        {{ codefile "shell" "examples/dns_zone/import.sh" }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deletion_protection: Flag that protects the dns zone from accidental deletion.
        :param pulumi.Input[str] description: Description of the DNS zone.
        :param pulumi.Input[str] folder_id: ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the DNS zone.
        :param pulumi.Input[str] name: User assigned name of a specific resource. Must be unique within the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_networks: For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        :param pulumi.Input[bool] public: The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
               Cloud resources.
        :param pulumi.Input[str] zone: The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnsZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        {{ .Description }}

        ## Example Usage

        {{ tffile "examples/dns_zone/r_dns_zone_1.tf" }}

        {{ .SchemaMarkdown }}

        ## Import

        {{ codefile "shell" "examples/dns_zone/import.sh" }}

        :param str resource_name: The name of the resource.
        :param DnsZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnsZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DnsZoneArgs.__new__(DnsZoneArgs)

            __props__.__dict__["deletion_protection"] = deletion_protection
            __props__.__dict__["description"] = description
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["private_networks"] = private_networks
            __props__.__dict__["public"] = public
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
        super(DnsZone, __self__).__init__(
            'yandex:index/dnsZone:DnsZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            public: Optional[pulumi.Input[bool]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'DnsZone':
        """
        Get an existing DnsZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The DNS zone creation timestamp.
        :param pulumi.Input[bool] deletion_protection: Flag that protects the dns zone from accidental deletion.
        :param pulumi.Input[str] description: Description of the DNS zone.
        :param pulumi.Input[str] folder_id: ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the DNS zone.
        :param pulumi.Input[str] name: User assigned name of a specific resource. Must be unique within the folder.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_networks: For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        :param pulumi.Input[bool] public: The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
               Cloud resources.
        :param pulumi.Input[str] zone: The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnsZoneState.__new__(_DnsZoneState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["deletion_protection"] = deletion_protection
        __props__.__dict__["description"] = description
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["private_networks"] = private_networks
        __props__.__dict__["public"] = public
        __props__.__dict__["zone"] = zone
        return DnsZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The DNS zone creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag that protects the dns zone from accidental deletion.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the DNS zone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the DNS zone.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User assigned name of a specific resource. Must be unique within the folder.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
        """
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[bool]:
        """
        The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private
        Cloud resources.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The DNS name of this zone, e.g. "example.com.". Must ends with dot.
        """
        return pulumi.get(self, "zone")

