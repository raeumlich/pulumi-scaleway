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
from ._inputs import *

__all__ = ['ContainerTriggerArgs', 'ContainerTrigger']

@pulumi.input_type
class ContainerTriggerArgs:
    def __init__(__self__, *,
                 container_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['ContainerTriggerNatsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input['ContainerTriggerSqsArgs']] = None):
        """
        The set of arguments for constructing a ContainerTrigger resource.
        :param pulumi.Input[str] container_id: The ID of the container to create a trigger for
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input['ContainerTriggerNatsArgs'] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input['ContainerTriggerSqsArgs'] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        pulumi.set(__self__, "container_id", container_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Input[str]:
        """
        The ID of the container to create a trigger for
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['ContainerTriggerNatsArgs']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['ContainerTriggerNatsArgs']]):
        pulumi.set(self, "nats", value)

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

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['ContainerTriggerSqsArgs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['ContainerTriggerSqsArgs']]):
        pulumi.set(self, "sqs", value)


@pulumi.input_type
class _ContainerTriggerState:
    def __init__(__self__, *,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['ContainerTriggerNatsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input['ContainerTriggerSqsArgs']] = None):
        """
        Input properties used for looking up and filtering ContainerTrigger resources.
        :param pulumi.Input[str] container_id: The ID of the container to create a trigger for
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input['ContainerTriggerNatsArgs'] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input['ContainerTriggerSqsArgs'] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        if container_id is not None:
            pulumi.set(__self__, "container_id", container_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the container to create a trigger for
        """
        return pulumi.get(self, "container_id")

    @container_id.setter
    def container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['ContainerTriggerNatsArgs']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['ContainerTriggerNatsArgs']]):
        pulumi.set(self, "nats", value)

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

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['ContainerTriggerSqsArgs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['ContainerTriggerSqsArgs']]):
        pulumi.set(self, "sqs", value)


class ContainerTrigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerNatsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerSqsArgs']]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Container Triggers.
        For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers).

        ## Example Usage

        ### SQS

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.serverless.ContainerTrigger("main",
            container_id=scaleway_container["main"]["id"],
            sqs=scaleway.serverless.ContainerTriggerSqsArgs(
                project_id=scaleway_mnq_sqs["main"]["project_id"],
                queue="MyQueue",
                region=scaleway_mnq_sqs["main"]["region"],
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Nats

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.serverless.ContainerTrigger("main",
            container_id=scaleway_container["main"]["id"],
            nats=scaleway.serverless.ContainerTriggerNatsArgs(
                account_id=scaleway_mnq_nats_account["main"]["id"],
                subject="MySubject",
                region=scaleway_mnq_nats_account["main"]["region"],
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Container Triggers can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:serverless/containerTrigger:ContainerTrigger main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container to create a trigger for
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input[pulumi.InputType['ContainerTriggerNatsArgs']] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[pulumi.InputType['ContainerTriggerSqsArgs']] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerTriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Container Triggers.
        For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-triggers).

        ## Example Usage

        ### SQS

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.serverless.ContainerTrigger("main",
            container_id=scaleway_container["main"]["id"],
            sqs=scaleway.serverless.ContainerTriggerSqsArgs(
                project_id=scaleway_mnq_sqs["main"]["project_id"],
                queue="MyQueue",
                region=scaleway_mnq_sqs["main"]["region"],
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Nats

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main = scaleway.serverless.ContainerTrigger("main",
            container_id=scaleway_container["main"]["id"],
            nats=scaleway.serverless.ContainerTriggerNatsArgs(
                account_id=scaleway_mnq_nats_account["main"]["id"],
                subject="MySubject",
                region=scaleway_mnq_nats_account["main"]["region"],
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Container Triggers can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:serverless/containerTrigger:ContainerTrigger main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ContainerTriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerTriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerNatsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerSqsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerTriggerArgs.__new__(ContainerTriggerArgs)

            if container_id is None and not opts.urn:
                raise TypeError("Missing required property 'container_id'")
            __props__.__dict__["container_id"] = container_id
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["nats"] = nats
            __props__.__dict__["region"] = region
            __props__.__dict__["sqs"] = sqs
        super(ContainerTrigger, __self__).__init__(
            'scaleway:serverless/containerTrigger:ContainerTrigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nats: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerNatsArgs']]] = None,
            region: Optional[pulumi.Input[str]] = None,
            sqs: Optional[pulumi.Input[pulumi.InputType['ContainerTriggerSqsArgs']]] = None) -> 'ContainerTrigger':
        """
        Get an existing ContainerTrigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_id: The ID of the container to create a trigger for
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input[pulumi.InputType['ContainerTriggerNatsArgs']] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[pulumi.InputType['ContainerTriggerSqsArgs']] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerTriggerState.__new__(_ContainerTriggerState)

        __props__.__dict__["container_id"] = container_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["nats"] = nats
        __props__.__dict__["region"] = region
        __props__.__dict__["sqs"] = sqs
        return ContainerTrigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerId")
    def container_id(self) -> pulumi.Output[str]:
        """
        The ID of the container to create a trigger for
        """
        return pulumi.get(self, "container_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nats(self) -> pulumi.Output[Optional['outputs.ContainerTriggerNats']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def sqs(self) -> pulumi.Output[Optional['outputs.ContainerTriggerSqs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

