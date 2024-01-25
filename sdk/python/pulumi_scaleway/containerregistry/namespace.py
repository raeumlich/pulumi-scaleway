# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NamespaceArgs', 'Namespace']

@pulumi.input_type
class NamespaceArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Namespace resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] is_public: Whether the images stored in the namespace should be downloadable publicly (docker pull).
        :param pulumi.Input[str] name: The unique name of the namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the images stored in the namespace should be downloadable publicly (docker pull).
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _NamespaceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Namespace resources.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[str] endpoint: Endpoint reachable by Docker.
        :param pulumi.Input[bool] is_public: Whether the images stored in the namespace should be downloadable publicly (docker pull).
        :param pulumi.Input[str] name: The unique name of the namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] organization_id: The organization ID the namespace is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint reachable by Docker.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the images stored in the namespace should be downloadable publicly (docker pull).
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the namespace is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Namespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Container Registry.
        For more information see [the documentation](https://developers.scaleway.com/en/products/registry/api/).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.containerregistry.Namespace("main",
            description="Main container registry",
            is_public=False)
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:containerregistry/namespace:Namespace main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[bool] is_public: Whether the images stored in the namespace should be downloadable publicly (docker pull).
        :param pulumi.Input[str] name: The unique name of the namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NamespaceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Container Registry.
        For more information see [the documentation](https://developers.scaleway.com/en/products/registry/api/).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.containerregistry.Namespace("main",
            description="Main container registry",
            is_public=False)
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:containerregistry/namespace:Namespace main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param NamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceArgs.__new__(NamespaceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["is_public"] = is_public
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["organization_id"] = None
        super(Namespace, __self__).__init__(
            'scaleway:containerregistry/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the namespace.
        :param pulumi.Input[str] endpoint: Endpoint reachable by Docker.
        :param pulumi.Input[bool] is_public: Whether the images stored in the namespace should be downloadable publicly (docker pull).
        :param pulumi.Input[str] name: The unique name of the namespace.
               
               > **Important** Updates to `name` will recreate the namespace.
        :param pulumi.Input[str] organization_id: The organization ID the namespace is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the namespace is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceState.__new__(_NamespaceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["is_public"] = is_public
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return Namespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the namespace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint reachable by Docker.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the images stored in the namespace should be downloadable publicly (docker pull).
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the namespace.

        > **Important** Updates to `name` will recreate the namespace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the namespace is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the namespace is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

