# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketPolicyArgs', 'BucketPolicy']

@pulumi.input_type
class BucketPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 policy: pulumi.Input[str],
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketPolicy resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "policy", policy)
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
    def policy(self) -> pulumi.Input[str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

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
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BucketPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BucketPolicy resources.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

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
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

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
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class BucketPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway object storage bucket policy.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/).

        ## Example Usage
        ### Example Usage with an IAM user

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        user = scaleway.iam.get_user(email="user@scaleway.com")
        policy_policy = scaleway.iam.Policy("policyPolicy",
            user_id=user.id,
            rules=[scaleway.iam.PolicyRuleArgs(
                project_ids=[default.id],
                permission_set_names=["ObjectStorageFullAccess"],
            )])
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy_bucket_policy = scaleway.objectstorage.BucketPolicy("policyBucketPolicy",
            bucket=bucket.name,
            policy=pulumi.Output.all(bucket.name, bucket.name).apply(lambda bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": ["s3:*"],
                    "Principal": {
                        "SCW": f"user_id:{user.id}",
                    },
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```
        ### Example with an IAM application
        ### Creating a bucket and delegating read access to an application

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # IAM configuration
        reading_app = scaleway.iam.Application("reading-app")
        policy_policy = scaleway.iam.Policy("policyPolicy",
            application_id=reading_app.id,
            rules=[scaleway.iam.PolicyRuleArgs(
                project_ids=[default.id],
                permission_set_names=["ObjectStorageBucketsRead"],
            )])
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy_bucket_policy = scaleway.objectstorage.BucketPolicy("policyBucketPolicy",
            bucket=bucket.id,
            policy=pulumi.Output.all(reading_app.id, bucket.name, bucket.name).apply(lambda id, bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Statement": [{
                    "Sid": "Delegate read access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": f"application_id:{id}",
                    },
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObject",
                    ],
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```
        ### Reading the bucket with the application

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        reading_app = scaleway.iam.get_application(name="reading-app")
        reading_api_key = scaleway.iam.APIKey("reading-api-key", application_id=reading_app.id)
        reading_profile = scaleway.Provider("reading-profile",
            access_key=reading_api_key.access_key,
            secret_key=reading_api_key.secret_key)
        bucket = scaleway.objectstorage.get_bucket(name="some-unique-name")
        ```
        ### Example with AWS provider

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy = aws.iam.get_policy_document_output(version="2012-10-17",
            statements=[aws.iam.GetPolicyDocumentStatementArgs(
                sid="Delegate access",
                effect="Allow",
                principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                    type="SCW",
                    identifiers=[f"project_id:{default.id}"],
                )],
                actions=["s3:ListBucket"],
                resources=[
                    bucket.name,
                    bucket.name.apply(lambda name: f"{name}/*"),
                ],
            )])
        main = scaleway.objectstorage.BucketPolicy("main",
            bucket=bucket.id,
            policy=policy.json)
        ```
        ### Example with deprecated version 2012-10-17

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket", region="fr-par")
        policy = scaleway.objectstorage.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.all(bucket.name, bucket.name).apply(lambda bucketName, bucketName1: json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObjectTagging",
                    ],
                    "Principal": {
                        "SCW": f"project_id:{default.id}",
                    },
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```

        **NB:** To configure the AWS provider with Scaleway credentials, please visit this [tutorial](https://www.scaleway.com/en/docs/storage/object/api-cli/object-storage-aws-cli/).

        ## Import

        Bucket policies can be imported using the `{region}/{bucketName}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket
        ```

         ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway object storage bucket policy.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/).

        ## Example Usage
        ### Example Usage with an IAM user

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        user = scaleway.iam.get_user(email="user@scaleway.com")
        policy_policy = scaleway.iam.Policy("policyPolicy",
            user_id=user.id,
            rules=[scaleway.iam.PolicyRuleArgs(
                project_ids=[default.id],
                permission_set_names=["ObjectStorageFullAccess"],
            )])
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy_bucket_policy = scaleway.objectstorage.BucketPolicy("policyBucketPolicy",
            bucket=bucket.name,
            policy=pulumi.Output.all(bucket.name, bucket.name).apply(lambda bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": ["s3:*"],
                    "Principal": {
                        "SCW": f"user_id:{user.id}",
                    },
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```
        ### Example with an IAM application
        ### Creating a bucket and delegating read access to an application

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # IAM configuration
        reading_app = scaleway.iam.Application("reading-app")
        policy_policy = scaleway.iam.Policy("policyPolicy",
            application_id=reading_app.id,
            rules=[scaleway.iam.PolicyRuleArgs(
                project_ids=[default.id],
                permission_set_names=["ObjectStorageBucketsRead"],
            )])
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy_bucket_policy = scaleway.objectstorage.BucketPolicy("policyBucketPolicy",
            bucket=bucket.id,
            policy=pulumi.Output.all(reading_app.id, bucket.name, bucket.name).apply(lambda id, bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Statement": [{
                    "Sid": "Delegate read access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": f"application_id:{id}",
                    },
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObject",
                    ],
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```
        ### Reading the bucket with the application

        ```python
        import pulumi
        import pulumi_scaleway as scaleway

        reading_app = scaleway.iam.get_application(name="reading-app")
        reading_api_key = scaleway.iam.APIKey("reading-api-key", application_id=reading_app.id)
        reading_profile = scaleway.Provider("reading-profile",
            access_key=reading_api_key.access_key,
            secret_key=reading_api_key.secret_key)
        bucket = scaleway.objectstorage.get_bucket(name="some-unique-name")
        ```
        ### Example with AWS provider

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket")
        policy = aws.iam.get_policy_document_output(version="2012-10-17",
            statements=[aws.iam.GetPolicyDocumentStatementArgs(
                sid="Delegate access",
                effect="Allow",
                principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                    type="SCW",
                    identifiers=[f"project_id:{default.id}"],
                )],
                actions=["s3:ListBucket"],
                resources=[
                    bucket.name,
                    bucket.name.apply(lambda name: f"{name}/*"),
                ],
            )])
        main = scaleway.objectstorage.BucketPolicy("main",
            bucket=bucket.id,
            policy=policy.json)
        ```
        ### Example with deprecated version 2012-10-17

        ```python
        import pulumi
        import json
        import pulumi_scaleway as scaleway

        default = scaleway.account.get_project(name="default")
        # Object storage configuration
        bucket = scaleway.objectstorage.Bucket("bucket", region="fr-par")
        policy = scaleway.objectstorage.BucketPolicy("policy",
            bucket=bucket.name,
            policy=pulumi.Output.all(bucket.name, bucket.name).apply(lambda bucketName, bucketName1: json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": [
                        "s3:ListBucket",
                        "s3:GetObjectTagging",
                    ],
                    "Principal": {
                        "SCW": f"project_id:{default.id}",
                    },
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```

        **NB:** To configure the AWS provider with Scaleway credentials, please visit this [tutorial](https://www.scaleway.com/en/docs/storage/object/api-cli/object-storage-aws-cli/).

        ## Import

        Bucket policies can be imported using the `{region}/{bucketName}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket
        ```

         ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project. If you are using a project different from the default one, you have to specify the project ID at the end of the import command. bash

        ```sh
         $ pulumi import scaleway:objectstorage/bucketPolicy:BucketPolicy some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param BucketPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketPolicyArgs.__new__(BucketPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
        super(BucketPolicy, __self__).__init__(
            'scaleway:objectstorage/bucketPolicy:BucketPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'BucketPolicy':
        """
        Get an existing BucketPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The bucket's name or regional ID.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketPolicyState.__new__(_BucketPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["policy"] = policy
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return BucketPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The bucket's name or regional ID.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

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
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

