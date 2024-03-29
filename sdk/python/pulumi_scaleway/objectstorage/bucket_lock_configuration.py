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

__all__ = ['BucketLockConfigurationArgs', 'BucketLockConfiguration']

@pulumi.input_type
class BucketLockConfigurationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 rule: pulumi.Input['BucketLockConfigurationRuleArgs'],
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketLockConfiguration resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input['BucketLockConfigurationRuleArgs'] rule: Specifies the Object Lock rule for the specified object.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "rule", rule)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Input['BucketLockConfigurationRuleArgs']:
        """
        Specifies the Object Lock rule for the specified object.
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: pulumi.Input['BucketLockConfigurationRuleArgs']):
        pulumi.set(self, "rule", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BucketLockConfigurationState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input['BucketLockConfigurationRuleArgs']] = None):
        """
        Input properties used for looking up and filtering BucketLockConfiguration resources.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input['BucketLockConfigurationRuleArgs'] rule: Specifies the Object Lock rule for the specified object.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rule is not None:
            pulumi.set(__self__, "rule", rule)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rule(self) -> Optional[pulumi.Input['BucketLockConfigurationRuleArgs']]:
        """
        Specifies the Object Lock rule for the specified object.
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: Optional[pulumi.Input['BucketLockConfigurationRuleArgs']]):
        pulumi.set(self, "rule", value)


class BucketLockConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['BucketLockConfigurationRuleArgs']]] = None,
                 __props__=None):
        """
        Provides an Object bucket lock configuration resource.
        For more information, see [Setting up object lock](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/).

        ## Example Usage

        ### Configure an Object Lock for a new bucket

        Please note that `object_lock_enabled` must be set to `true` before configuring the lock.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_bucket = scaleway.objectstorage.Bucket("mainBucket",
            acl="public-read",
            object_lock_enabled=True)
        main_bucket_lock_configuration = scaleway.objectstorage.BucketLockConfiguration("mainBucketLockConfiguration",
            bucket=main_bucket.name,
            rule=scaleway.objectstorage.BucketLockConfigurationRuleArgs(
                default_retention=scaleway.objectstorage.BucketLockConfigurationRuleDefaultRetentionArgs(
                    mode="GOVERNANCE",
                    days=1,
                ),
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Configure an Object Lock for an existing bucket

        You should [contact Scaleway support](https://console.scaleway.com/support/tickets/create) to enable object lock on an existing bucket.

        ## Import

        Bucket lock configurations can be imported using the `{region}/{bucketName}` identifier, e.g.

        bash

        ```sh
        $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[pulumi.InputType['BucketLockConfigurationRuleArgs']] rule: Specifies the Object Lock rule for the specified object.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketLockConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Object bucket lock configuration resource.
        For more information, see [Setting up object lock](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/).

        ## Example Usage

        ### Configure an Object Lock for a new bucket

        Please note that `object_lock_enabled` must be set to `true` before configuring the lock.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_bucket = scaleway.objectstorage.Bucket("mainBucket",
            acl="public-read",
            object_lock_enabled=True)
        main_bucket_lock_configuration = scaleway.objectstorage.BucketLockConfiguration("mainBucketLockConfiguration",
            bucket=main_bucket.name,
            rule=scaleway.objectstorage.BucketLockConfigurationRuleArgs(
                default_retention=scaleway.objectstorage.BucketLockConfigurationRuleDefaultRetentionArgs(
                    mode="GOVERNANCE",
                    days=1,
                ),
            ))
        ```
        <!--End PulumiCodeChooser -->

        ### Configure an Object Lock for an existing bucket

        You should [contact Scaleway support](https://console.scaleway.com/support/tickets/create) to enable object lock on an existing bucket.

        ## Import

        Bucket lock configurations can be imported using the `{region}/{bucketName}` identifier, e.g.

        bash

        ```sh
        $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param BucketLockConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketLockConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['BucketLockConfigurationRuleArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketLockConfigurationArgs.__new__(BucketLockConfigurationArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if rule is None and not opts.urn:
                raise TypeError("Missing required property 'rule'")
            __props__.__dict__["rule"] = rule
        super(BucketLockConfiguration, __self__).__init__(
            'scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rule: Optional[pulumi.Input[pulumi.InputType['BucketLockConfigurationRuleArgs']]] = None) -> 'BucketLockConfiguration':
        """
        Get an existing BucketLockConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[pulumi.InputType['BucketLockConfigurationRuleArgs']] rule: Specifies the Object Lock rule for the specified object.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketLockConfigurationState.__new__(_BucketLockConfigurationState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["rule"] = rule
        return BucketLockConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Output['outputs.BucketLockConfigurationRule']:
        """
        Specifies the Object Lock rule for the specified object.
        """
        return pulumi.get(self, "rule")

