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

__all__ = [
    'GetComputeGpuClusterResult',
    'AwaitableGetComputeGpuClusterResult',
    'get_compute_gpu_cluster',
    'get_compute_gpu_cluster_output',
]

@pulumi.output_type
class GetComputeGpuClusterResult:
    """
    A collection of values returned by getComputeGpuCluster.
    """
    def __init__(__self__, created_at=None, description=None, folder_id=None, gpu_cluster_id=None, id=None, interconnect_type=None, labels=None, name=None, status=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if folder_id and not isinstance(folder_id, str):
            raise TypeError("Expected argument 'folder_id' to be a str")
        pulumi.set(__self__, "folder_id", folder_id)
        if gpu_cluster_id and not isinstance(gpu_cluster_id, str):
            raise TypeError("Expected argument 'gpu_cluster_id' to be a str")
        pulumi.set(__self__, "gpu_cluster_id", gpu_cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interconnect_type and not isinstance(interconnect_type, str):
            raise TypeError("Expected argument 'interconnect_type' to be a str")
        pulumi.set(__self__, "interconnect_type", interconnect_type)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Creation timestamp.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional description of the GPU cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> str:
        """
        ID of the folder that the GPU cluster belongs to.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="gpuClusterId")
    def gpu_cluster_id(self) -> str:
        return pulumi.get(self, "gpu_cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interconnectType")
    def interconnect_type(self) -> str:
        """
        type of interconnect used between nodes in GPU cluster.
        """
        return pulumi.get(self, "interconnect_type")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        GPU cluster labels as `key:value` pairs. For details about the concept, see [documentation](https://cloud.yandex.com/docs/overview/concepts/services#labels).
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status of the GPU cluster.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        ID of the zone where the GPU cluster resides.
        """
        return pulumi.get(self, "zone")


class AwaitableGetComputeGpuClusterResult(GetComputeGpuClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeGpuClusterResult(
            created_at=self.created_at,
            description=self.description,
            folder_id=self.folder_id,
            gpu_cluster_id=self.gpu_cluster_id,
            id=self.id,
            interconnect_type=self.interconnect_type,
            labels=self.labels,
            name=self.name,
            status=self.status,
            zone=self.zone)


def get_compute_gpu_cluster(folder_id: Optional[str] = None,
                            gpu_cluster_id: Optional[str] = None,
                            name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeGpuClusterResult:
    """
    Get information about a Yandex Compute GPU cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/compute/concepts/gpu-cluster).

    ## Example Usage

    {{ tffile "examples/compute_gpu_cluster/d_compute_gpu_cluster_1.tf" }}


    :param str folder_id: ID of the folder that the GPU cluster belongs to.
    :param str gpu_cluster_id: ID of the GPU cluster.
    :param str name: Name of the GPU cluster.
           
           > One of `gpu_cluster_id` or `name` should be specified.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['gpuClusterId'] = gpu_cluster_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getComputeGpuCluster:getComputeGpuCluster', __args__, opts=opts, typ=GetComputeGpuClusterResult).value

    return AwaitableGetComputeGpuClusterResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        folder_id=pulumi.get(__ret__, 'folder_id'),
        gpu_cluster_id=pulumi.get(__ret__, 'gpu_cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        interconnect_type=pulumi.get(__ret__, 'interconnect_type'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        status=pulumi.get(__ret__, 'status'),
        zone=pulumi.get(__ret__, 'zone'))
def get_compute_gpu_cluster_output(folder_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   gpu_cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   name: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetComputeGpuClusterResult]:
    """
    Get information about a Yandex Compute GPU cluster. For more information, see [the official documentation](https://cloud.yandex.com/docs/compute/concepts/gpu-cluster).

    ## Example Usage

    {{ tffile "examples/compute_gpu_cluster/d_compute_gpu_cluster_1.tf" }}


    :param str folder_id: ID of the folder that the GPU cluster belongs to.
    :param str gpu_cluster_id: ID of the GPU cluster.
    :param str name: Name of the GPU cluster.
           
           > One of `gpu_cluster_id` or `name` should be specified.
    """
    __args__ = dict()
    __args__['folderId'] = folder_id
    __args__['gpuClusterId'] = gpu_cluster_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('yandex:index/getComputeGpuCluster:getComputeGpuCluster', __args__, opts=opts, typ=GetComputeGpuClusterResult)
    return __ret__.apply(lambda __response__: GetComputeGpuClusterResult(
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        folder_id=pulumi.get(__response__, 'folder_id'),
        gpu_cluster_id=pulumi.get(__response__, 'gpu_cluster_id'),
        id=pulumi.get(__response__, 'id'),
        interconnect_type=pulumi.get(__response__, 'interconnect_type'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        status=pulumi.get(__response__, 'status'),
        zone=pulumi.get(__response__, 'zone')))
