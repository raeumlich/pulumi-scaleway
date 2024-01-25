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

__all__ = ['BucketACLArgs', 'BucketACL']

@pulumi.input_type
class BucketACLArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 access_control_policy: Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 expected_bucket_owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketACL resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input['BucketACLAccessControlPolicyArgs'] access_control_policy: A configuration block that sets the ACL permissions for an object per grantee documented below.
        :param pulumi.Input[str] acl: The canned ACL you want to apply to the bucket.
        :param pulumi.Input[str] expected_bucket_owner: The project ID of the expected bucket owner.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        pulumi.set(__self__, "bucket", bucket)
        if access_control_policy is not None:
            pulumi.set(__self__, "access_control_policy", access_control_policy)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if expected_bucket_owner is not None:
            pulumi.set(__self__, "expected_bucket_owner", expected_bucket_owner)
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
    @pulumi.getter(name="accessControlPolicy")
    def access_control_policy(self) -> Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']]:
        """
        A configuration block that sets the ACL permissions for an object per grantee documented below.
        """
        return pulumi.get(self, "access_control_policy")

    @access_control_policy.setter
    def access_control_policy(self, value: Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']]):
        pulumi.set(self, "access_control_policy", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input[str]]:
        """
        The canned ACL you want to apply to the bucket.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter(name="expectedBucketOwner")
    def expected_bucket_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID of the expected bucket owner.
        """
        return pulumi.get(self, "expected_bucket_owner")

    @expected_bucket_owner.setter
    def expected_bucket_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_bucket_owner", value)

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
        The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BucketACLState:
    def __init__(__self__, *,
                 access_control_policy: Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 expected_bucket_owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BucketACL resources.
        :param pulumi.Input['BucketACLAccessControlPolicyArgs'] access_control_policy: A configuration block that sets the ACL permissions for an object per grantee documented below.
        :param pulumi.Input[str] acl: The canned ACL you want to apply to the bucket.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] expected_bucket_owner: The project ID of the expected bucket owner.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        if access_control_policy is not None:
            pulumi.set(__self__, "access_control_policy", access_control_policy)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if expected_bucket_owner is not None:
            pulumi.set(__self__, "expected_bucket_owner", expected_bucket_owner)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="accessControlPolicy")
    def access_control_policy(self) -> Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']]:
        """
        A configuration block that sets the ACL permissions for an object per grantee documented below.
        """
        return pulumi.get(self, "access_control_policy")

    @access_control_policy.setter
    def access_control_policy(self, value: Optional[pulumi.Input['BucketACLAccessControlPolicyArgs']]):
        pulumi.set(self, "access_control_policy", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input[str]]:
        """
        The canned ACL you want to apply to the bucket.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acl", value)

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
    @pulumi.getter(name="expectedBucketOwner")
    def expected_bucket_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID of the expected bucket owner.
        """
        return pulumi.get(self, "expected_bucket_owner")

    @expected_bucket_owner.setter
    def expected_bucket_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_bucket_owner", value)

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
        The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class BucketACL(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_policy: Optional[pulumi.Input[pulumi.InputType['BucketACLAccessControlPolicyArgs']]] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 expected_bucket_owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        some_bucket = scaleway.objectstorage.Bucket("someBucket")
        main = scaleway.objectstorage.BucketACL("main",
            bucket=scaleway_object_bucket["main"]["id"],
            acl="private")
        ```
        ### With Grants

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_bucket = scaleway.objectstorage.Bucket("mainBucket")
        main_bucket_acl = scaleway.objectstorage.BucketACL("mainBucketACL",
            bucket=main_bucket.id,
            access_control_policy=scaleway.objectstorage.BucketACLAccessControlPolicyArgs(
                grants=[
                    scaleway.objectstorage.BucketACLAccessControlPolicyGrantArgs(
                        grantee=scaleway.objectstorage.BucketACLAccessControlPolicyGrantGranteeArgs(
                            id="<project-id>:<project-id>",
                            type="CanonicalUser",
                        ),
                        permission="FULL_CONTROL",
                    ),
                    scaleway.objectstorage.BucketACLAccessControlPolicyGrantArgs(
                        grantee=scaleway.objectstorage.BucketACLAccessControlPolicyGrantGranteeArgs(
                            id="<project-id>",
                            type="CanonicalUser",
                        ),
                        permission="WRITE",
                    ),
                ],
                owner=scaleway.objectstorage.BucketACLAccessControlPolicyOwnerArgs(
                    id="<project-id>",
                ),
            ))
        ```
        ## The ACL

        Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl)

        ## The Access Control policy

        The `access_control_policy` configuration block supports the following arguments:

        * `grant` - (Required) Set of grant configuration blocks documented below.
        * `owner` - (Required) Configuration block of the bucket owner's display name and ID documented below.

        ## The Grant

        The `grant` configuration block supports the following arguments:

        * `grantee` - (Required) Configuration block for the project being granted permissions documented below.
        * `permission` - (Required) Logging permissions assigned to the grantee for the bucket.

        ## The permission

        The following list shows each access policy permissions supported.

        `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`

        For more information about ACL permissions in the S3 bucket, see [ACL permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html).

        ## The owner

        The `owner` configuration block supports the following arguments:

        * `id` - (Required) The ID of the project owner.
        * `display_name` - (Optional) The display name of the owner.

        ## the grantee

        The `grantee` configuration block supports the following arguments:

        * `id` - (Required) The canonical user ID of the grantee.
        * `type` - (Required) Type of grantee. Valid values: CanonicalUser.

        ## Import

        Bucket ACLs can be imported using the `{region}/{bucketName}/{acl}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketACL:BucketACL some_bucket fr-par/some-bucket/private
        ```

         ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketACL:BucketACL some_bucket fr-par/some-bucket/private@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BucketACLAccessControlPolicyArgs']] access_control_policy: A configuration block that sets the ACL permissions for an object per grantee documented below.
        :param pulumi.Input[str] acl: The canned ACL you want to apply to the bucket.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] expected_bucket_owner: The project ID of the expected bucket owner.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketACLArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        some_bucket = scaleway.objectstorage.Bucket("someBucket")
        main = scaleway.objectstorage.BucketACL("main",
            bucket=scaleway_object_bucket["main"]["id"],
            acl="private")
        ```
        ### With Grants

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        main_bucket = scaleway.objectstorage.Bucket("mainBucket")
        main_bucket_acl = scaleway.objectstorage.BucketACL("mainBucketACL",
            bucket=main_bucket.id,
            access_control_policy=scaleway.objectstorage.BucketACLAccessControlPolicyArgs(
                grants=[
                    scaleway.objectstorage.BucketACLAccessControlPolicyGrantArgs(
                        grantee=scaleway.objectstorage.BucketACLAccessControlPolicyGrantGranteeArgs(
                            id="<project-id>:<project-id>",
                            type="CanonicalUser",
                        ),
                        permission="FULL_CONTROL",
                    ),
                    scaleway.objectstorage.BucketACLAccessControlPolicyGrantArgs(
                        grantee=scaleway.objectstorage.BucketACLAccessControlPolicyGrantGranteeArgs(
                            id="<project-id>",
                            type="CanonicalUser",
                        ),
                        permission="WRITE",
                    ),
                ],
                owner=scaleway.objectstorage.BucketACLAccessControlPolicyOwnerArgs(
                    id="<project-id>",
                ),
            ))
        ```
        ## The ACL

        Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl)

        ## The Access Control policy

        The `access_control_policy` configuration block supports the following arguments:

        * `grant` - (Required) Set of grant configuration blocks documented below.
        * `owner` - (Required) Configuration block of the bucket owner's display name and ID documented below.

        ## The Grant

        The `grant` configuration block supports the following arguments:

        * `grantee` - (Required) Configuration block for the project being granted permissions documented below.
        * `permission` - (Required) Logging permissions assigned to the grantee for the bucket.

        ## The permission

        The following list shows each access policy permissions supported.

        `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`

        For more information about ACL permissions in the S3 bucket, see [ACL permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html).

        ## The owner

        The `owner` configuration block supports the following arguments:

        * `id` - (Required) The ID of the project owner.
        * `display_name` - (Optional) The display name of the owner.

        ## the grantee

        The `grantee` configuration block supports the following arguments:

        * `id` - (Required) The canonical user ID of the grantee.
        * `type` - (Required) Type of grantee. Valid values: CanonicalUser.

        ## Import

        Bucket ACLs can be imported using the `{region}/{bucketName}/{acl}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketACL:BucketACL some_bucket fr-par/some-bucket/private
        ```

         ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketACL:BucketACL some_bucket fr-par/some-bucket/private@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param BucketACLArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketACLArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_policy: Optional[pulumi.Input[pulumi.InputType['BucketACLAccessControlPolicyArgs']]] = None,
                 acl: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 expected_bucket_owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketACLArgs.__new__(BucketACLArgs)

            __props__.__dict__["access_control_policy"] = access_control_policy
            __props__.__dict__["acl"] = acl
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["expected_bucket_owner"] = expected_bucket_owner
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
        super(BucketACL, __self__).__init__(
            'scaleway:objectstorage/bucketACL:BucketACL',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_control_policy: Optional[pulumi.Input[pulumi.InputType['BucketACLAccessControlPolicyArgs']]] = None,
            acl: Optional[pulumi.Input[str]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            expected_bucket_owner: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'BucketACL':
        """
        Get an existing BucketACL resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BucketACLAccessControlPolicyArgs']] access_control_policy: A configuration block that sets the ACL permissions for an object per grantee documented below.
        :param pulumi.Input[str] acl: The canned ACL you want to apply to the bucket.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] expected_bucket_owner: The project ID of the expected bucket owner.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketACLState.__new__(_BucketACLState)

        __props__.__dict__["access_control_policy"] = access_control_policy
        __props__.__dict__["acl"] = acl
        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["expected_bucket_owner"] = expected_bucket_owner
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return BucketACL(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessControlPolicy")
    def access_control_policy(self) -> pulumi.Output['outputs.BucketACLAccessControlPolicy']:
        """
        A configuration block that sets the ACL permissions for an object per grantee documented below.
        """
        return pulumi.get(self, "access_control_policy")

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output[Optional[str]]:
        """
        The canned ACL you want to apply to the bucket.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="expectedBucketOwner")
    def expected_bucket_owner(self) -> pulumi.Output[Optional[str]]:
        """
        The project ID of the expected bucket owner.
        """
        return pulumi.get(self, "expected_bucket_owner")

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
        The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
        """
        return pulumi.get(self, "region")

