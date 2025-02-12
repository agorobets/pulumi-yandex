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
from . import outputs
from ._inputs import *

__all__ = [
    'GetMdbKafkaClusterResult',
    'AwaitableGetMdbKafkaClusterResult',
    'get_mdb_kafka_cluster',
    'get_mdb_kafka_cluster_output',
]

@pulumi.output_type
class GetMdbKafkaClusterResult:
    """
    A collection of values returned by getMdbKafkaCluster.
    """
    def __init__(__self__, accesses=None, cluster_id=None, config=None, created_at=None, deletion_protection=None, description=None, disk_size_autoscalings=None, environment=None, folder_id=None, health=None, host_group_ids=None, hosts=None, id=None, labels=None, maintenance_windows=None, name=None, network_id=None, security_group_ids=None, status=None, subnet_ids=None, topics=None, users=None):
        if accesses and not isinstance(accesses, list):
            raise TypeError("Expected argument 'accesses' to be a list")
        pulumi.set(__self__, "accesses", accesses)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if deletion_protection and not isinstance(deletion_protection, bool):
            raise TypeError("Expected argument 'deletion_protection' to be a bool")
        pulumi.set(__self__, "deletion_protection", deletion_protection)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_size_autoscalings and not isinstance(disk_size_autoscalings, list):
            raise TypeError("Expected argument 'disk_size_autoscalings' to be a list")
        pulumi.set(__self__, "disk_size_autoscalings", disk_size_autoscalings)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if health and not isinstance(health, str):
            raise TypeError("Expected argument 'health' to be a str")
        pulumi.set(__self__, "health", health)
        if host_group_ids and not isinstance(host_group_ids, list):
            raise TypeError("Expected argument 'host_group_ids' to be a list")
        pulumi.set(__self__, "host_group_ids", host_group_ids)
        if hosts and not isinstance(hosts, list):
            raise TypeError("Expected argument 'hosts' to be a list")
        pulumi.set(__self__, "hosts", hosts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if maintenance_windows and not isinstance(maintenance_windows, list):
            raise TypeError("Expected argument 'maintenance_windows' to be a list")
        pulumi.set(__self__, "maintenance_windows", maintenance_windows)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def accesses(self) -> Sequence['outputs.GetMdbKafkaClusterAccessResult']:
        """
        (Optional) Access policy to the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "accesses")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def config(self) -> Optional['outputs.GetMdbKafkaClusterConfigResult']:
        """
        Configuration of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Creation timestamp of the key.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> bool:
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the Kafka cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskSizeAutoscalings")
    def disk_size_autoscalings(self) -> Sequence['outputs.GetMdbKafkaClusterDiskSizeAutoscalingResult']:
        """
        Disk autoscaling settings of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "disk_size_autoscalings")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        Deployment environment of the Kafka cluster.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> str:
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def health(self) -> str:
        """
        Health of the host.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter(name="hostGroupIds")
    def host_group_ids(self) -> Sequence[str]:
        """
        A list of IDs of the host groups hosting VMs of the cluster.
        """
        return pulumi.get(self, "host_group_ids")

    @property
    @pulumi.getter
    def hosts(self) -> Sequence['outputs.GetMdbKafkaClusterHostResult']:
        """
        A host of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        A set of key/value label pairs to assign to the Kafka cluster.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maintenanceWindows")
    def maintenance_windows(self) -> Sequence['outputs.GetMdbKafkaClusterMaintenanceWindowResult']:
        """
        Maintenance window settings of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "maintenance_windows")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The fully qualified domain name of the host.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        ID of the network, to which the Kafka cluster belongs.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Sequence[str]:
        """
        A list of security groups IDs of the Kafka cluster.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def topics(self) -> Optional[Sequence['outputs.GetMdbKafkaClusterTopicResult']]:
        """
        A topic of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence['outputs.GetMdbKafkaClusterUserResult']]:
        """
        A user of the Kafka cluster. The structure is documented below.
        """
        return pulumi.get(self, "users")


class AwaitableGetMdbKafkaClusterResult(GetMdbKafkaClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMdbKafkaClusterResult(
            accesses=self.accesses,
            cluster_id=self.cluster_id,
            config=self.config,
            created_at=self.created_at,
            deletion_protection=self.deletion_protection,
            description=self.description,
            disk_size_autoscalings=self.disk_size_autoscalings,
            environment=self.environment,
            folder_id=self.folder_id,
            health=self.health,
            host_group_ids=self.host_group_ids,
            hosts=self.hosts,
            id=self.id,
            labels=self.labels,
            maintenance_windows=self.maintenance_windows,
            name=self.name,
            network_id=self.network_id,
            security_group_ids=self.security_group_ids,
            status=self.status,
            subnet_ids=self.subnet_ids,
            topics=self.topics,
            users=self.users)


def get_mdb_kafka_cluster(cluster_id: Optional[str] = None,
                          config: Optional[Union['GetMdbKafkaClusterConfigArgs', 'GetMdbKafkaClusterConfigArgsDict']] = None,
                          deletion_protection: Optional[bool] = None,
                          folder_id: Optional[str] = None,
                          name: Optional[str] = None,
                          subnet_ids: Optional[Sequence[str]] = None,
                          topics: Optional[Sequence[Union['GetMdbKafkaClusterTopicArgs', 'GetMdbKafkaClusterTopicArgsDict']]] = None,
                          users: Optional[Sequence[Union['GetMdbKafkaClusterUserArgs', 'GetMdbKafkaClusterUserArgsDict']]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMdbKafkaClusterResult:
    """
    Get information about a Yandex Managed Kafka cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).

    ## Example Usage

    {{ tffile "examples/mdb_kafka_cluster/d_mdb_kafka_cluster_1.tf" }}


    :param str cluster_id: The ID of the Kafka cluster.
    :param Union['GetMdbKafkaClusterConfigArgs', 'GetMdbKafkaClusterConfigArgsDict'] config: Configuration of the Kafka cluster. The structure is documented below.
    :param str folder_id: The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
    :param str name: The name of the Kafka cluster.
           
           > Either `cluster_id` or `name` should be specified.
    :param Sequence[Union['GetMdbKafkaClusterTopicArgs', 'GetMdbKafkaClusterTopicArgsDict']] topics: A topic of the Kafka cluster. The structure is documented below.
    :param Sequence[Union['GetMdbKafkaClusterUserArgs', 'GetMdbKafkaClusterUserArgsDict']] users: A user of the Kafka cluster. The structure is documented below.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['config'] = config
    __args__['deletionProtection'] = deletion_protection
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['subnetIds'] = subnet_ids
    __args__['topics'] = topics
    __args__['users'] = users
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getMdbKafkaCluster:getMdbKafkaCluster', __args__, opts=opts, typ=GetMdbKafkaClusterResult).value

    return AwaitableGetMdbKafkaClusterResult(
        accesses=pulumi.get(__ret__, 'accesses'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        config=pulumi.get(__ret__, 'config'),
        created_at=pulumi.get(__ret__, 'created_at'),
        deletion_protection=pulumi.get(__ret__, 'deletion_protection'),
        description=pulumi.get(__ret__, 'description'),
        disk_size_autoscalings=pulumi.get(__ret__, 'disk_size_autoscalings'),
        environment=pulumi.get(__ret__, 'environment'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        health=pulumi.get(__ret__, 'health'),
        host_group_ids=pulumi.get(__ret__, 'host_group_ids'),
        hosts=pulumi.get(__ret__, 'hosts'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        maintenance_windows=pulumi.get(__ret__, 'maintenance_windows'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        status=pulumi.get(__ret__, 'status'),
        subnet_ids=pulumi.get(__ret__, 'subnet_ids'),
        topics=pulumi.get(__ret__, 'topics'),
        users=pulumi.get(__ret__, 'users'))
def get_mdb_kafka_cluster_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 config: Optional[pulumi.Input[Optional[Union['GetMdbKafkaClusterConfigArgs', 'GetMdbKafkaClusterConfigArgsDict']]]] = None,
                                 deletion_protection: Optional[pulumi.Input[Optional[bool]]] = None,
                                 folder_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 subnet_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 topics: Optional[pulumi.Input[Optional[Sequence[Union['GetMdbKafkaClusterTopicArgs', 'GetMdbKafkaClusterTopicArgsDict']]]]] = None,
                                 users: Optional[pulumi.Input[Optional[Sequence[Union['GetMdbKafkaClusterUserArgs', 'GetMdbKafkaClusterUserArgsDict']]]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMdbKafkaClusterResult]:
    """
    Get information about a Yandex Managed Kafka cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-kafka/concepts).

    ## Example Usage

    {{ tffile "examples/mdb_kafka_cluster/d_mdb_kafka_cluster_1.tf" }}


    :param str cluster_id: The ID of the Kafka cluster.
    :param Union['GetMdbKafkaClusterConfigArgs', 'GetMdbKafkaClusterConfigArgsDict'] config: Configuration of the Kafka cluster. The structure is documented below.
    :param str folder_id: The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
    :param str name: The name of the Kafka cluster.
           
           > Either `cluster_id` or `name` should be specified.
    :param Sequence[Union['GetMdbKafkaClusterTopicArgs', 'GetMdbKafkaClusterTopicArgsDict']] topics: A topic of the Kafka cluster. The structure is documented below.
    :param Sequence[Union['GetMdbKafkaClusterUserArgs', 'GetMdbKafkaClusterUserArgsDict']] users: A user of the Kafka cluster. The structure is documented below.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['config'] = config
    __args__['deletionProtection'] = deletion_protection
    __args__['folderId'] = folder_id
    __args__['name'] = name
    __args__['subnetIds'] = subnet_ids
    __args__['topics'] = topics
    __args__['users'] = users
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getMdbKafkaCluster:getMdbKafkaCluster', __args__, opts=opts, typ=GetMdbKafkaClusterResult)
    return __ret__.apply(lambda __response__: GetMdbKafkaClusterResult(
        accesses=pulumi.get(__response__, 'accesses'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        config=pulumi.get(__response__, 'config'),
        created_at=pulumi.get(__response__, 'created_at'),
        deletion_protection=pulumi.get(__response__, 'deletion_protection'),
        description=pulumi.get(__response__, 'description'),
        disk_size_autoscalings=pulumi.get(__response__, 'disk_size_autoscalings'),
        environment=pulumi.get(__response__, 'environment'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        health=pulumi.get(__response__, 'health'),
        host_group_ids=pulumi.get(__response__, 'host_group_ids'),
        hosts=pulumi.get(__response__, 'hosts'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        maintenance_windows=pulumi.get(__response__, 'maintenance_windows'),
        name=pulumi.get(__response__, 'name'),
        network_id=pulumi.get(__response__, 'network_id'),
        security_group_ids=pulumi.get(__response__, 'security_group_ids'),
        status=pulumi.get(__response__, 'status'),
        subnet_ids=pulumi.get(__response__, 'subnet_ids'),
        topics=pulumi.get(__response__, 'topics'),
        users=pulumi.get(__response__, 'users')))
