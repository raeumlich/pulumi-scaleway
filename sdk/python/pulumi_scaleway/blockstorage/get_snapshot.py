# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSnapshotResult',
    'AwaitableGetSnapshotResult',
    'get_snapshot',
    'get_snapshot_output',
]

@pulumi.output_type
class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, id=None, name=None, project_id=None, snapshot_id=None, tags=None, volume_id=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[str]:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            volume_id=self.volume_id,
            zone=self.zone)


def get_snapshot(name: Optional[str] = None,
                 project_id: Optional[str] = None,
                 snapshot_id: Optional[str] = None,
                 volume_id: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotResult:
    """
    Gets information about a Block Snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_snapshot = scaleway.blockstorage.get_snapshot(snapshot_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The name of the snapshot. Only one of `name` and `snapshot_id` should be specified.
    :param str project_id: The ID of the project the snapshot is associated with.
    :param str snapshot_id: The ID of the snapshot. Only one of `name` and `snapshot_id` should be specified.
    :param str volume_id: The ID of the volume from which the snapshot has been created.
    :param str zone: `zone`) The zone in which the snapshot exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['snapshotId'] = snapshot_id
    __args__['volumeId'] = volume_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:blockstorage/getSnapshot:getSnapshot', __args__, opts=opts, typ=GetSnapshotResult).value

    return AwaitableGetSnapshotResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_snapshot)
def get_snapshot_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        snapshot_id: Optional[pulumi.Input[Optional[str]]] = None,
                        volume_id: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotResult]:
    """
    Gets information about a Block Snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_snapshot = scaleway.blockstorage.get_snapshot(snapshot_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The name of the snapshot. Only one of `name` and `snapshot_id` should be specified.
    :param str project_id: The ID of the project the snapshot is associated with.
    :param str snapshot_id: The ID of the snapshot. Only one of `name` and `snapshot_id` should be specified.
    :param str volume_id: The ID of the volume from which the snapshot has been created.
    :param str zone: `zone`) The zone in which the snapshot exists.
    """
    ...
