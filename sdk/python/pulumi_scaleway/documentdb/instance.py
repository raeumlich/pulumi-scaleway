# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 engine: pulumi.Input[str],
                 node_type: pulumi.Input[str],
                 is_ha_cluster: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 telemetry_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 volume_size_in_gb: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] engine: Database Instance's engine version (e.g. `FerretDB-1`).
               
               > **Important:** Updates to `engine` will recreate the Database Instance.
        :param pulumi.Input[str] node_type: The type of database instance you want to create (e.g. `docdb-play2-pico`).
               
               > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
               interruption. Keep in mind that you cannot downgrade a Database Instance.
        :param pulumi.Input[bool] is_ha_cluster: Enable or disable high availability for the database instance.
        :param pulumi.Input[str] name: The name of the Database Instance.
        :param pulumi.Input[str] password: Password for the first user of the database instance.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Database
               Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database Instance should be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the Database Instance.
        :param pulumi.Input[bool] telemetry_enabled: Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
               
               > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        :param pulumi.Input[str] user_name: Identifier for the first user of the database instance.
               
               > **Important:** Updates to `user_name` will recreate the Database Instance.
        :param pulumi.Input[int] volume_size_in_gb: Volume size (in GB) when `volume_type` is set to `bssd`.
        :param pulumi.Input[str] volume_type: Type of volume where data are stored (`bssd` or `lssd`).
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "node_type", node_type)
        if is_ha_cluster is not None:
            pulumi.set(__self__, "is_ha_cluster", is_ha_cluster)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if telemetry_enabled is not None:
            pulumi.set(__self__, "telemetry_enabled", telemetry_enabled)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if volume_size_in_gb is not None:
            pulumi.set(__self__, "volume_size_in_gb", volume_size_in_gb)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        Database Instance's engine version (e.g. `FerretDB-1`).

        > **Important:** Updates to `engine` will recreate the Database Instance.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        """
        The type of database instance you want to create (e.g. `docdb-play2-pico`).

        > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        interruption. Keep in mind that you cannot downgrade a Database Instance.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="isHaCluster")
    def is_ha_cluster(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable high availability for the database instance.
        """
        return pulumi.get(self, "is_ha_cluster")

    @is_ha_cluster.setter
    def is_ha_cluster(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ha_cluster", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Database Instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the first user of the database instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the Database
        Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the Database Instance.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="telemetryEnabled")
    def telemetry_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).

        > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        """
        return pulumi.get(self, "telemetry_enabled")

    @telemetry_enabled.setter
    def telemetry_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "telemetry_enabled", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the first user of the database instance.

        > **Important:** Updates to `user_name` will recreate the Database Instance.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        Volume size (in GB) when `volume_type` is set to `bssd`.
        """
        return pulumi.get(self, "volume_size_in_gb")

    @volume_size_in_gb.setter
    def volume_size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size_in_gb", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of volume where data are stored (`bssd` or `lssd`).
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 engine: Optional[pulumi.Input[str]] = None,
                 is_ha_cluster: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 telemetry_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 volume_size_in_gb: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] engine: Database Instance's engine version (e.g. `FerretDB-1`).
               
               > **Important:** Updates to `engine` will recreate the Database Instance.
        :param pulumi.Input[bool] is_ha_cluster: Enable or disable high availability for the database instance.
        :param pulumi.Input[str] name: The name of the Database Instance.
        :param pulumi.Input[str] node_type: The type of database instance you want to create (e.g. `docdb-play2-pico`).
               
               > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
               interruption. Keep in mind that you cannot downgrade a Database Instance.
        :param pulumi.Input[str] password: Password for the first user of the database instance.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Database
               Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database Instance should be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the Database Instance.
        :param pulumi.Input[bool] telemetry_enabled: Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
               
               > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        :param pulumi.Input[str] user_name: Identifier for the first user of the database instance.
               
               > **Important:** Updates to `user_name` will recreate the Database Instance.
        :param pulumi.Input[int] volume_size_in_gb: Volume size (in GB) when `volume_type` is set to `bssd`.
        :param pulumi.Input[str] volume_type: Type of volume where data are stored (`bssd` or `lssd`).
        """
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if is_ha_cluster is not None:
            pulumi.set(__self__, "is_ha_cluster", is_ha_cluster)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if telemetry_enabled is not None:
            pulumi.set(__self__, "telemetry_enabled", telemetry_enabled)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if volume_size_in_gb is not None:
            pulumi.set(__self__, "volume_size_in_gb", volume_size_in_gb)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        Database Instance's engine version (e.g. `FerretDB-1`).

        > **Important:** Updates to `engine` will recreate the Database Instance.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="isHaCluster")
    def is_ha_cluster(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable high availability for the database instance.
        """
        return pulumi.get(self, "is_ha_cluster")

    @is_ha_cluster.setter
    def is_ha_cluster(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ha_cluster", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Database Instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of database instance you want to create (e.g. `docdb-play2-pico`).

        > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        interruption. Keep in mind that you cannot downgrade a Database Instance.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the first user of the database instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the Database
        Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the Database Instance.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="telemetryEnabled")
    def telemetry_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).

        > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        """
        return pulumi.get(self, "telemetry_enabled")

    @telemetry_enabled.setter
    def telemetry_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "telemetry_enabled", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the first user of the database instance.

        > **Important:** Updates to `user_name` will recreate the Database Instance.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        Volume size (in GB) when `volume_type` is set to `bssd`.
        """
        return pulumi.get(self, "volume_size_in_gb")

    @volume_size_in_gb.setter
    def volume_size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size_in_gb", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of volume where data are stored (`bssd` or `lssd`).
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 is_ha_cluster: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 telemetry_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 volume_size_in_gb: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Database Instances.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).

        ## Example Usage
        ### Example Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.documentdb.Instance("main",
            engine="FerretDB-1",
            node_type="docdb-play2-pico",
            password="thiZ_is_v&ry_s3cret",
            tags=[
                "terraform-test",
                "scaleway_documentdb_instance",
                "minimal",
            ],
            user_name="my_initial_user",
            volume_size_in_gb=20)
        ```

        ## Import

        Database Instance can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:documentdb/instance:Instance db fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine: Database Instance's engine version (e.g. `FerretDB-1`).
               
               > **Important:** Updates to `engine` will recreate the Database Instance.
        :param pulumi.Input[bool] is_ha_cluster: Enable or disable high availability for the database instance.
        :param pulumi.Input[str] name: The name of the Database Instance.
        :param pulumi.Input[str] node_type: The type of database instance you want to create (e.g. `docdb-play2-pico`).
               
               > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
               interruption. Keep in mind that you cannot downgrade a Database Instance.
        :param pulumi.Input[str] password: Password for the first user of the database instance.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Database
               Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database Instance should be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the Database Instance.
        :param pulumi.Input[bool] telemetry_enabled: Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
               
               > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        :param pulumi.Input[str] user_name: Identifier for the first user of the database instance.
               
               > **Important:** Updates to `user_name` will recreate the Database Instance.
        :param pulumi.Input[int] volume_size_in_gb: Volume size (in GB) when `volume_type` is set to `bssd`.
        :param pulumi.Input[str] volume_type: Type of volume where data are stored (`bssd` or `lssd`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Database Instances.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).

        ## Example Usage
        ### Example Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.documentdb.Instance("main",
            engine="FerretDB-1",
            node_type="docdb-play2-pico",
            password="thiZ_is_v&ry_s3cret",
            tags=[
                "terraform-test",
                "scaleway_documentdb_instance",
                "minimal",
            ],
            user_name="my_initial_user",
            volume_size_in_gb=20)
        ```

        ## Import

        Database Instance can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:documentdb/instance:Instance db fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 is_ha_cluster: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 telemetry_enabled: Optional[pulumi.Input[bool]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 volume_size_in_gb: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["is_ha_cluster"] = is_ha_cluster
            __props__.__dict__["name"] = name
            if node_type is None and not opts.urn:
                raise TypeError("Missing required property 'node_type'")
            __props__.__dict__["node_type"] = node_type
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["telemetry_enabled"] = telemetry_enabled
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["volume_size_in_gb"] = volume_size_in_gb
            __props__.__dict__["volume_type"] = volume_type
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Instance, __self__).__init__(
            'scaleway:documentdb/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            engine: Optional[pulumi.Input[str]] = None,
            is_ha_cluster: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_type: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            telemetry_enabled: Optional[pulumi.Input[bool]] = None,
            user_name: Optional[pulumi.Input[str]] = None,
            volume_size_in_gb: Optional[pulumi.Input[int]] = None,
            volume_type: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] engine: Database Instance's engine version (e.g. `FerretDB-1`).
               
               > **Important:** Updates to `engine` will recreate the Database Instance.
        :param pulumi.Input[bool] is_ha_cluster: Enable or disable high availability for the database instance.
        :param pulumi.Input[str] name: The name of the Database Instance.
        :param pulumi.Input[str] node_type: The type of database instance you want to create (e.g. `docdb-play2-pico`).
               
               > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
               interruption. Keep in mind that you cannot downgrade a Database Instance.
        :param pulumi.Input[str] password: Password for the first user of the database instance.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Database
               Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database Instance should be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the Database Instance.
        :param pulumi.Input[bool] telemetry_enabled: Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
               
               > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        :param pulumi.Input[str] user_name: Identifier for the first user of the database instance.
               
               > **Important:** Updates to `user_name` will recreate the Database Instance.
        :param pulumi.Input[int] volume_size_in_gb: Volume size (in GB) when `volume_type` is set to `bssd`.
        :param pulumi.Input[str] volume_type: Type of volume where data are stored (`bssd` or `lssd`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["engine"] = engine
        __props__.__dict__["is_ha_cluster"] = is_ha_cluster
        __props__.__dict__["name"] = name
        __props__.__dict__["node_type"] = node_type
        __props__.__dict__["password"] = password
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tags"] = tags
        __props__.__dict__["telemetry_enabled"] = telemetry_enabled
        __props__.__dict__["user_name"] = user_name
        __props__.__dict__["volume_size_in_gb"] = volume_size_in_gb
        __props__.__dict__["volume_type"] = volume_type
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        Database Instance's engine version (e.g. `FerretDB-1`).

        > **Important:** Updates to `engine` will recreate the Database Instance.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="isHaCluster")
    def is_ha_cluster(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable or disable high availability for the database instance.
        """
        return pulumi.get(self, "is_ha_cluster")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Database Instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Output[str]:
        """
        The type of database instance you want to create (e.g. `docdb-play2-pico`).

        > **Important:** Updates to `node_type` will upgrade the Database Instance to the desired `node_type` without any
        interruption. Keep in mind that you cannot downgrade a Database Instance.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for the first user of the database instance.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the Database
        Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region
        in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the Database Instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="telemetryEnabled")
    def telemetry_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).

        > **Important:** Updates to `is_ha_cluster` will recreate the Database Instance.
        """
        return pulumi.get(self, "telemetry_enabled")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier for the first user of the database instance.

        > **Important:** Updates to `user_name` will recreate the Database Instance.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> pulumi.Output[int]:
        """
        Volume size (in GB) when `volume_type` is set to `bssd`.
        """
        return pulumi.get(self, "volume_size_in_gb")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of volume where data are stored (`bssd` or `lssd`).
        """
        return pulumi.get(self, "volume_type")

