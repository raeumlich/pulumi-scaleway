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
                 type: pulumi.Input[str],
                 from_snapshot_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Volume resource.
        :param pulumi.Input[str] type: The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        :param pulumi.Input[str] from_snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        pulumi.set(__self__, "type", type)
        if from_snapshot_id is not None:
            pulumi.set(__self__, "from_snapshot_id", from_snapshot_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="fromSnapshotId")
    def from_snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "from_snapshot_id")

    @from_snapshot_id.setter
    def from_snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_snapshot_id", value)

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
        The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_gb", value)

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
                 from_snapshot_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Volume resources.
        :param pulumi.Input[str] from_snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the volume is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[str] server_id: The id of the associated server.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] type: The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        if from_snapshot_id is not None:
            pulumi.set(__self__, "from_snapshot_id", from_snapshot_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="fromSnapshotId")
    def from_snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "from_snapshot_id")

    @from_snapshot_id.setter
    def from_snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_snapshot_id", value)

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
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the volume is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

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
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the associated server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_gb", value)

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
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
                 from_snapshot_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Instance Volumes.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        server_volume = scaleway.instance.Volume("serverVolume",
            size_in_gb=20,
            type="l_ssd")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        volumes can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:instance/volume:Volume server_volume fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] from_snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] type: The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Instance Volumes.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        server_volume = scaleway.instance.Volume("serverVolume",
            size_in_gb=20,
            type="l_ssd")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        volumes can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:instance/volume:Volume server_volume fr-par-1/11111111-1111-1111-1111-111111111111
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
                 from_snapshot_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeArgs.__new__(VolumeArgs)

            __props__.__dict__["from_snapshot_id"] = from_snapshot_id
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["size_in_gb"] = size_in_gb
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["zone"] = zone
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["server_id"] = None
        super(Volume, __self__).__init__(
            'scaleway:instance/volume:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            from_snapshot_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            size_in_gb: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Volume':
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] from_snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[str] name: The name of the volume. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the volume is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the volume is associated with.
        :param pulumi.Input[str] server_id: The id of the associated server.
        :param pulumi.Input[int] size_in_gb: The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[str] type: The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        :param pulumi.Input[str] zone: `zone`) The zone in which the volume should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeState.__new__(_VolumeState)

        __props__.__dict__["from_snapshot_id"] = from_snapshot_id
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["size_in_gb"] = size_in_gb
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["zone"] = zone
        return Volume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fromSnapshotId")
    def from_snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "from_snapshot_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the volume. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the volume is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The id of the associated server.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> pulumi.Output[Optional[int]]:
        """
        The size of the volume. Only one of `size_in_gb` and `from_snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the volume. The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD), `scratch` (Local Scratch SSD).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

