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

__all__ = ['RouteArgs', 'Route']

@pulumi.input_type
class RouteArgs:
    def __init__(__self__, *,
                 hub_id: pulumi.Input[str],
                 topic: pulumi.Input[str],
                 database: Optional[pulumi.Input['RouteDatabaseArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rest: Optional[pulumi.Input['RouteRestArgs']] = None,
                 s3: Optional[pulumi.Input['RouteS3Args']] = None):
        """
        The set of arguments for constructing a Route resource.
        :param pulumi.Input[str] hub_id: The hub ID to which the Route will be attached to.
        :param pulumi.Input[str] topic: The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        :param pulumi.Input['RouteDatabaseArgs'] database: Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] name: The name of the IoT Route you want to create (e.g. `my-route`).
        :param pulumi.Input[str] region: (Defaults to provider `region`) The region in which the Route is attached to.
        :param pulumi.Input['RouteRestArgs'] rest: Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        :param pulumi.Input['RouteS3Args'] s3: Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        """
        pulumi.set(__self__, "hub_id", hub_id)
        pulumi.set(__self__, "topic", topic)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rest is not None:
            pulumi.set(__self__, "rest", rest)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> pulumi.Input[str]:
        """
        The hub ID to which the Route will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @hub_id.setter
    def hub_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hub_id", value)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input['RouteDatabaseArgs']]:
        """
        Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input['RouteDatabaseArgs']]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Route you want to create (e.g. `my-route`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        (Defaults to provider `region`) The region in which the Route is attached to.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rest(self) -> Optional[pulumi.Input['RouteRestArgs']]:
        """
        Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "rest")

    @rest.setter
    def rest(self, value: Optional[pulumi.Input['RouteRestArgs']]):
        pulumi.set(self, "rest", value)

    @property
    @pulumi.getter
    def s3(self) -> Optional[pulumi.Input['RouteS3Args']]:
        """
        Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: Optional[pulumi.Input['RouteS3Args']]):
        pulumi.set(self, "s3", value)


@pulumi.input_type
class _RouteState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input['RouteDatabaseArgs']] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rest: Optional[pulumi.Input['RouteRestArgs']] = None,
                 s3: Optional[pulumi.Input['RouteS3Args']] = None,
                 topic: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Route resources.
        :param pulumi.Input[str] created_at: The date and time the Route was created.
        :param pulumi.Input['RouteDatabaseArgs'] database: Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] hub_id: The hub ID to which the Route will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Route you want to create (e.g. `my-route`).
        :param pulumi.Input[str] region: (Defaults to provider `region`) The region in which the Route is attached to.
        :param pulumi.Input['RouteRestArgs'] rest: Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        :param pulumi.Input['RouteS3Args'] s3: Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] topic: The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if hub_id is not None:
            pulumi.set(__self__, "hub_id", hub_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rest is not None:
            pulumi.set(__self__, "rest", rest)
        if s3 is not None:
            pulumi.set(__self__, "s3", s3)
        if topic is not None:
            pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Route was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input['RouteDatabaseArgs']]:
        """
        Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input['RouteDatabaseArgs']]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> Optional[pulumi.Input[str]]:
        """
        The hub ID to which the Route will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @hub_id.setter
    def hub_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Route you want to create (e.g. `my-route`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        (Defaults to provider `region`) The region in which the Route is attached to.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rest(self) -> Optional[pulumi.Input['RouteRestArgs']]:
        """
        Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "rest")

    @rest.setter
    def rest(self, value: Optional[pulumi.Input['RouteRestArgs']]):
        pulumi.set(self, "rest", value)

    @property
    @pulumi.getter
    def s3(self) -> Optional[pulumi.Input['RouteS3Args']]:
        """
        Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: Optional[pulumi.Input['RouteS3Args']]):
        pulumi.set(self, "s3", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        """
        The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)


class Route(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[pulumi.InputType['RouteDatabaseArgs']]] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rest: Optional[pulumi.Input[pulumi.InputType['RouteRestArgs']]] = None,
                 s3: Optional[pulumi.Input[pulumi.InputType['RouteS3Args']]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### S3 Route

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_hub = scaleway.iot.Hub("mainHub", product_plan="plan_shared")
        main_bucket = scaleway.objectstorage.Bucket("mainBucket", region="fr-par")
        main_route = scaleway.iot.Route("mainRoute",
            hub_id=main_hub.id,
            topic="#",
            s3=scaleway.iot.RouteS3Args(
                bucket_region=main_bucket.region,
                bucket_name=main_bucket.name,
                object_prefix="foo",
                strategy="per_topic",
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Rest Route

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_hub = scaleway.iot.Hub("mainHub", product_plan="plan_shared")
        main_route = scaleway.iot.Route("mainRoute",
            hub_id=main_hub.id,
            topic="#",
            rest=scaleway.iot.RouteRestArgs(
                verb="get",
                uri="http://scaleway.com",
                headers={
                    "X-awesome-header": "my-awesome-value",
                },
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        IoT Routes can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:iot/route:Route route01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RouteDatabaseArgs']] database: Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] hub_id: The hub ID to which the Route will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Route you want to create (e.g. `my-route`).
        :param pulumi.Input[str] region: (Defaults to provider `region`) The region in which the Route is attached to.
        :param pulumi.Input[pulumi.InputType['RouteRestArgs']] rest: Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        :param pulumi.Input[pulumi.InputType['RouteS3Args']] s3: Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] topic: The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### S3 Route

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_hub = scaleway.iot.Hub("mainHub", product_plan="plan_shared")
        main_bucket = scaleway.objectstorage.Bucket("mainBucket", region="fr-par")
        main_route = scaleway.iot.Route("mainRoute",
            hub_id=main_hub.id,
            topic="#",
            s3=scaleway.iot.RouteS3Args(
                bucket_region=main_bucket.region,
                bucket_name=main_bucket.name,
                object_prefix="foo",
                strategy="per_topic",
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Rest Route

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_hub = scaleway.iot.Hub("mainHub", product_plan="plan_shared")
        main_route = scaleway.iot.Route("mainRoute",
            hub_id=main_hub.id,
            topic="#",
            rest=scaleway.iot.RouteRestArgs(
                verb="get",
                uri="http://scaleway.com",
                headers={
                    "X-awesome-header": "my-awesome-value",
                },
            ))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        IoT Routes can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:iot/route:Route route01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param RouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[pulumi.InputType['RouteDatabaseArgs']]] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rest: Optional[pulumi.Input[pulumi.InputType['RouteRestArgs']]] = None,
                 s3: Optional[pulumi.Input[pulumi.InputType['RouteS3Args']]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteArgs.__new__(RouteArgs)

            __props__.__dict__["database"] = database
            if hub_id is None and not opts.urn:
                raise TypeError("Missing required property 'hub_id'")
            __props__.__dict__["hub_id"] = hub_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["rest"] = rest
            __props__.__dict__["s3"] = s3
            if topic is None and not opts.urn:
                raise TypeError("Missing required property 'topic'")
            __props__.__dict__["topic"] = topic
            __props__.__dict__["created_at"] = None
        super(Route, __self__).__init__(
            'scaleway:iot/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[pulumi.InputType['RouteDatabaseArgs']]] = None,
            hub_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rest: Optional[pulumi.Input[pulumi.InputType['RouteRestArgs']]] = None,
            s3: Optional[pulumi.Input[pulumi.InputType['RouteS3Args']]] = None,
            topic: Optional[pulumi.Input[str]] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the Route was created.
        :param pulumi.Input[pulumi.InputType['RouteDatabaseArgs']] database: Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] hub_id: The hub ID to which the Route will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Route you want to create (e.g. `my-route`).
        :param pulumi.Input[str] region: (Defaults to provider `region`) The region in which the Route is attached to.
        :param pulumi.Input[pulumi.InputType['RouteRestArgs']] rest: Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        :param pulumi.Input[pulumi.InputType['RouteS3Args']] s3: Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        :param pulumi.Input[str] topic: The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteState.__new__(_RouteState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["database"] = database
        __props__.__dict__["hub_id"] = hub_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["rest"] = rest
        __props__.__dict__["s3"] = s3
        __props__.__dict__["topic"] = topic
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the Route was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional['outputs.RouteDatabase']]:
        """
        Configuration block for the database routes. See  [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Database-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> pulumi.Output[str]:
        """
        The hub ID to which the Route will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IoT Route you want to create (e.g. `my-route`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        (Defaults to provider `region`) The region in which the Route is attached to.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rest(self) -> pulumi.Output[Optional['outputs.RouteRest']]:
        """
        Configuration block for the rest routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-REST-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "rest")

    @property
    @pulumi.getter
    def s3(self) -> pulumi.Output[Optional['outputs.RouteS3']]:
        """
        Configuration block for the S3 routes. See [product documentation](https://www.scaleway.com/en/docs/scaleway-iothub-route/#-Scaleway-Object-Storage-Route) for a better understanding of the parameters.
        """
        return pulumi.get(self, "s3")

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Output[str]:
        """
        The topic the Route subscribes to, wildcards allowed (e.g. `thelab/+/temperature/#`).
        """
        return pulumi.get(self, "topic")

