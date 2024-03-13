# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VolumeArgs', 'Volume']

@pulumi.input_type
class VolumeArgs:
    def __init__(__self__, *,
                 iops: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Volume resource.
        :param pulumi.Input[int] iops: The maximum IO/s expected, must match available options.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        pulumi.set(__self__, "iops", iops)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Input[int]:
        """
        The maximum IO/s expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: pulumi.Input[int]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the volume. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_gb", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _VolumeState:
    def __init__(__self__, *,
                 iops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Volume resources.
        :param pulumi.Input[int] iops: The maximum IO/s expected, must match available options.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum IO/s expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the volume. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_gb", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Volume(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Block Volumes.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        block_volume = scaleway.blockstorage.Volume("blockVolume",
            iops=5000,
            size_in_gb=20)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Block Volumes can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:blockstorage/volume:Volume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] iops: The maximum IO/s expected, must match available options.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Block Volumes.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        block_volume = scaleway.blockstorage.Volume("blockVolume",
            iops=5000,
            size_in_gb=20)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Block Volumes can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:blockstorage/volume:Volume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param VolumeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeArgs.__new__(VolumeArgs)

            if iops is None and not opts.urn:
                raise TypeError("Missing required property 'iops'")
            __props__.__dict__["iops"] = iops
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["size_in_gb"] = size_in_gb
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
        super(Volume, __self__).__init__(
            'scaleway:blockstorage/volume:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            iops: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            size_in_gb: Optional[pulumi.Input[int]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Volume':
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] iops: The maximum IO/s expected, must match available options.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeState.__new__(_VolumeState)

        __props__.__dict__["iops"] = iops
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["size_in_gb"] = size_in_gb
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return Volume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Output[int]:
        """
        The maximum IO/s expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the volume. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> pulumi.Output[int]:
        """
        The size of the volume. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

