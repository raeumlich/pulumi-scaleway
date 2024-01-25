# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ContainerTokenArgs', 'ContainerToken']

@pulumi.input_type
class ContainerTokenArgs:
    def __init__(__self__, *,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ContainerToken resource.
        :param pulumi.Input[str] container_id: The ID of the container.
               
               > Only one of `namespace_id` or `container_id` must be set.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] namespace_id: The ID of the container namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        """
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container.

        > Only one of `namespace_id` or `container_id` must be set.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ContainerTokenState:
    def __init__(__self__, *,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContainerToken resources.
        :param pulumi.Input[str] container_id: The ID of the container.
               
               > Only one of `namespace_id` or `container_id` must be set.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] namespace_id: The ID of the container namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        :param pulumi.Input[str] token: The token.
        """
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container.

        > Only one of `namespace_id` or `container_id` must be set.
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class ContainerToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Container Token.
        For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#tokens-26b085).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_container_namespace = scaleway.serverless.ContainerNamespace("mainContainerNamespace")
        main_container = scaleway.serverless.Container("mainContainer", namespace_id=main_container_namespace.id)
        # Namespace Token
        namespace = scaleway.serverless.ContainerToken("namespace",
            namespace_id=main_container_namespace.id,
            expires_at="2022-10-18T11:35:15+02:00")
        # Container Token
        container = scaleway.serverless.ContainerToken("container", container_id=main_container.id)
        ```

        ## Import

        Tokens can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:serverless/containerToken:ContainerToken main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container.
               
               > Only one of `namespace_id` or `container_id` must be set.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] namespace_id: The ID of the container namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContainerTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Container Token.
        For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#tokens-26b085).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_container_namespace = scaleway.serverless.ContainerNamespace("mainContainerNamespace")
        main_container = scaleway.serverless.Container("mainContainer", namespace_id=main_container_namespace.id)
        # Namespace Token
        namespace = scaleway.serverless.ContainerToken("namespace",
            namespace_id=main_container_namespace.id,
            expires_at="2022-10-18T11:35:15+02:00")
        # Container Token
        container = scaleway.serverless.ContainerToken("container", container_id=main_container.id)
        ```

        ## Import

        Tokens can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:serverless/containerToken:ContainerToken main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ContainerTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerTokenArgs.__new__(ContainerTokenArgs)

            __props__.__dict__["container_id"] = container_id
            __props__.__dict__["description"] = description
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["region"] = region
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ContainerToken, __self__).__init__(
            'scaleway:serverless/containerToken:ContainerToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'ContainerToken':
        """
        Get an existing ContainerToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container.
               
               > Only one of `namespace_id` or `container_id` must be set.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] namespace_id: The ID of the container namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        :param pulumi.Input[str] token: The token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerTokenState.__new__(_ContainerTokenState)

        __props__.__dict__["container_id"] = container_id
        __props__.__dict__["description"] = description
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["region"] = region
        __props__.__dict__["token"] = token
        return ContainerToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the container.

        > Only one of `namespace_id` or `container_id` must be set.
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the container namespace.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token.
        """
        return pulumi.get(self, "token")

