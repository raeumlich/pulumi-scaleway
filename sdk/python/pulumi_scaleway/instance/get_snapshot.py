# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

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
    def __init__(__self__, created_at=None, id=None, imports=None, name=None, organization_id=None, project_id=None, size_in_gb=None, snapshot_id=None, tags=None, type=None, volume_id=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if imports and not isinstance(imports, list):
            raise TypeError("Expected argument 'imports' to be a list")
        pulumi.set(__self__, "imports", imports)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if size_in_gb and not isinstance(size_in_gb, int):
            raise TypeError("Expected argument 'size_in_gb' to be a int")
        pulumi.set(__self__, "size_in_gb", size_in_gb)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def imports(self) -> Sequence['outputs.GetSnapshotImportResult']:
        return pulumi.get(self, "imports")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> int:
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
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
            created_at=self.created_at,
            id=self.id,
            imports=self.imports,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            size_in_gb=self.size_in_gb,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            type=self.type,
            volume_id=self.volume_id,
            zone=self.zone)


def get_snapshot(name: Optional[str] = None,
                 project_id: Optional[str] = None,
                 snapshot_id: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotResult:
    """
    Gets information about an instance snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.instance.get_snapshot(name="my-snapshot-name")
    by_id = scaleway.instance.get_snapshot(snapshot_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The snapshot name.
           Only one of `name` and `snapshot_id` should be specified.
    :param str project_id: `project_id`) The ID of the project the snapshot is associated with.
    :param str snapshot_id: The snapshot id.
           Only one of `name` and `snapshot_id` should be specified.
    :param str zone: `zone`) The zone in which the snapshot exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['snapshotId'] = snapshot_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:instance/getSnapshot:getSnapshot', __args__, opts=opts, typ=GetSnapshotResult).value

    return AwaitableGetSnapshotResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        imports=pulumi.get(__ret__, 'imports'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        size_in_gb=pulumi.get(__ret__, 'size_in_gb'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_snapshot)
def get_snapshot_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        snapshot_id: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotResult]:
    """
    Gets information about an instance snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.instance.get_snapshot(name="my-snapshot-name")
    by_id = scaleway.instance.get_snapshot(snapshot_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The snapshot name.
           Only one of `name` and `snapshot_id` should be specified.
    :param str project_id: `project_id`) The ID of the project the snapshot is associated with.
    :param str snapshot_id: The snapshot id.
           Only one of `name` and `snapshot_id` should be specified.
    :param str zone: `zone`) The zone in which the snapshot exists.
    """
    ...
