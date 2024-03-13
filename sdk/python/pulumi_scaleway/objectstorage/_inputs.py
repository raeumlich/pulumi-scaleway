# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'BucketACLAccessControlPolicyArgs',
    'BucketACLAccessControlPolicyGrantArgs',
    'BucketACLAccessControlPolicyGrantGranteeArgs',
    'BucketACLAccessControlPolicyOwnerArgs',
    'BucketCorsRuleArgs',
    'BucketLifecycleRuleArgs',
    'BucketLifecycleRuleExpirationArgs',
    'BucketLifecycleRuleTransitionArgs',
    'BucketLockConfigurationRuleArgs',
    'BucketLockConfigurationRuleDefaultRetentionArgs',
    'BucketVersioningArgs',
    'BucketWebsiteConfigurationErrorDocumentArgs',
    'BucketWebsiteConfigurationIndexDocumentArgs',
]

@pulumi.input_type
class BucketACLAccessControlPolicyArgs:
    def __init__(__self__, *,
                 owner: pulumi.Input['BucketACLAccessControlPolicyOwnerArgs'],
                 grants: Optional[pulumi.Input[Sequence[pulumi.Input['BucketACLAccessControlPolicyGrantArgs']]]] = None):
        """
        :param pulumi.Input['BucketACLAccessControlPolicyOwnerArgs'] owner: Configuration block of the bucket project owner's display organization ID.
        """
        pulumi.set(__self__, "owner", owner)
        if grants is not None:
            pulumi.set(__self__, "grants", grants)

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Input['BucketACLAccessControlPolicyOwnerArgs']:
        """
        Configuration block of the bucket project owner's display organization ID.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: pulumi.Input['BucketACLAccessControlPolicyOwnerArgs']):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def grants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketACLAccessControlPolicyGrantArgs']]]]:
        return pulumi.get(self, "grants")

    @grants.setter
    def grants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketACLAccessControlPolicyGrantArgs']]]]):
        pulumi.set(self, "grants", value)


@pulumi.input_type
class BucketACLAccessControlPolicyGrantArgs:
    def __init__(__self__, *,
                 permission: pulumi.Input[str],
                 grantee: Optional[pulumi.Input['BucketACLAccessControlPolicyGrantGranteeArgs']] = None):
        """
        :param pulumi.Input[str] permission: Logging permissions assigned to the grantee for the bucket.
        :param pulumi.Input['BucketACLAccessControlPolicyGrantGranteeArgs'] grantee: Configuration block for the project being granted permissions.
        """
        pulumi.set(__self__, "permission", permission)
        if grantee is not None:
            pulumi.set(__self__, "grantee", grantee)

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Input[str]:
        """
        Logging permissions assigned to the grantee for the bucket.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def grantee(self) -> Optional[pulumi.Input['BucketACLAccessControlPolicyGrantGranteeArgs']]:
        """
        Configuration block for the project being granted permissions.
        """
        return pulumi.get(self, "grantee")

    @grantee.setter
    def grantee(self, value: Optional[pulumi.Input['BucketACLAccessControlPolicyGrantGranteeArgs']]):
        pulumi.set(self, "grantee", value)


@pulumi.input_type
class BucketACLAccessControlPolicyGrantGranteeArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The `region`,`bucket` and `acl` separated by (`/`).
        :param pulumi.Input[str] type: Type of grantee. Valid values: `CanonicalUser`
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "type", type)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The `region`,`bucket` and `acl` separated by (`/`).
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of grantee. Valid values: `CanonicalUser`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class BucketACLAccessControlPolicyOwnerArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The `region`,`bucket` and `acl` separated by (`/`).
        :param pulumi.Input[str] display_name: The project ID of the grantee.
        """
        pulumi.set(__self__, "id", id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The `region`,`bucket` and `acl` separated by (`/`).
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID of the grantee.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class BucketCorsRuleArgs:
    def __init__(__self__, *,
                 allowed_methods: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_origins: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allowed_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expose_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_age_seconds: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_methods: Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_origins: Specifies which origins are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_headers: Specifies which headers are allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] expose_headers: Specifies expose header in the response.
        :param pulumi.Input[int] max_age_seconds: Specifies time in seconds that browser can cache the response for a preflight request.
        """
        pulumi.set(__self__, "allowed_methods", allowed_methods)
        pulumi.set(__self__, "allowed_origins", allowed_origins)
        if allowed_headers is not None:
            pulumi.set(__self__, "allowed_headers", allowed_headers)
        if expose_headers is not None:
            pulumi.set(__self__, "expose_headers", expose_headers)
        if max_age_seconds is not None:
            pulumi.set(__self__, "max_age_seconds", max_age_seconds)

    @property
    @pulumi.getter(name="allowedMethods")
    def allowed_methods(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
        """
        return pulumi.get(self, "allowed_methods")

    @allowed_methods.setter
    def allowed_methods(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_methods", value)

    @property
    @pulumi.getter(name="allowedOrigins")
    def allowed_origins(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies which origins are allowed.
        """
        return pulumi.get(self, "allowed_origins")

    @allowed_origins.setter
    def allowed_origins(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_origins", value)

    @property
    @pulumi.getter(name="allowedHeaders")
    def allowed_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies which headers are allowed.
        """
        return pulumi.get(self, "allowed_headers")

    @allowed_headers.setter
    def allowed_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_headers", value)

    @property
    @pulumi.getter(name="exposeHeaders")
    def expose_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies expose header in the response.
        """
        return pulumi.get(self, "expose_headers")

    @expose_headers.setter
    def expose_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "expose_headers", value)

    @property
    @pulumi.getter(name="maxAgeSeconds")
    def max_age_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies time in seconds that browser can cache the response for a preflight request.
        """
        return pulumi.get(self, "max_age_seconds")

    @max_age_seconds.setter
    def max_age_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age_seconds", value)


@pulumi.input_type
class BucketLifecycleRuleArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 abort_incomplete_multipart_upload_days: Optional[pulumi.Input[int]] = None,
                 expiration: Optional[pulumi.Input['BucketLifecycleRuleExpirationArgs']] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]] = None):
        """
        :param pulumi.Input[bool] enabled: The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
        :param pulumi.Input[int] abort_incomplete_multipart_upload_days: Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
               
               * > **Important:** It's not recommended using `prefix` for `AbortIncompleteMultipartUpload` as any incomplete multipart upload will be billed
        :param pulumi.Input['BucketLifecycleRuleExpirationArgs'] expiration: Specifies a period in the object's expire (documented below).
        :param pulumi.Input[str] id: Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        :param pulumi.Input[str] prefix: Object key prefix identifying one or more objects to which the rule applies.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies object tags key and value.
        :param pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]] transitions: Specifies a period in the object's transitions (documented below).
               
               At least one of `abort_incomplete_multipart_upload_days`, `expiration`, `transition` must be specified.
        """
        pulumi.set(__self__, "enabled", enabled)
        if abort_incomplete_multipart_upload_days is not None:
            pulumi.set(__self__, "abort_incomplete_multipart_upload_days", abort_incomplete_multipart_upload_days)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="abortIncompleteMultipartUploadDays")
    def abort_incomplete_multipart_upload_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.

        * > **Important:** It's not recommended using `prefix` for `AbortIncompleteMultipartUpload` as any incomplete multipart upload will be billed
        """
        return pulumi.get(self, "abort_incomplete_multipart_upload_days")

    @abort_incomplete_multipart_upload_days.setter
    def abort_incomplete_multipart_upload_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "abort_incomplete_multipart_upload_days", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input['BucketLifecycleRuleExpirationArgs']]:
        """
        Specifies a period in the object's expire (documented below).
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input['BucketLifecycleRuleExpirationArgs']]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the rule. Must be less than or equal to 255 characters in length.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Object key prefix identifying one or more objects to which the rule applies.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies object tags key and value.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]]:
        """
        Specifies a period in the object's transitions (documented below).

        At least one of `abort_incomplete_multipart_upload_days`, `expiration`, `transition` must be specified.
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketLifecycleRuleTransitionArgs']]]]):
        pulumi.set(self, "transitions", value)


@pulumi.input_type
class BucketLifecycleRuleExpirationArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[int]):
        """
        :param pulumi.Input[int] days: Specifies the number of days after object creation when the specific rule action takes effect.
               
               > **Important:**  If versioning is enabled, this rule only deletes the current version of an object.
        """
        pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[int]:
        """
        Specifies the number of days after object creation when the specific rule action takes effect.

        > **Important:**  If versioning is enabled, this rule only deletes the current version of an object.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[int]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class BucketLifecycleRuleTransitionArgs:
    def __init__(__self__, *,
                 storage_class: pulumi.Input[str],
                 days: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] storage_class: Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.
               
               > **Important:**  `ONEZONE_IA` is only available in `fr-par` region. The storage class `GLACIER` is not available in `pl-waw` region.
        :param pulumi.Input[int] days: Specifies the number of days after object creation when the specific rule action takes effect.
        """
        pulumi.set(__self__, "storage_class", storage_class)
        if days is not None:
            pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Input[str]:
        """
        Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA`  to which you want the object to transition.

        > **Important:**  `ONEZONE_IA` is only available in `fr-par` region. The storage class `GLACIER` is not available in `pl-waw` region.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_class", value)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days after object creation when the specific rule action takes effect.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "days", value)


@pulumi.input_type
class BucketLockConfigurationRuleArgs:
    def __init__(__self__, *,
                 default_retention: pulumi.Input['BucketLockConfigurationRuleDefaultRetentionArgs']):
        """
        :param pulumi.Input['BucketLockConfigurationRuleDefaultRetentionArgs'] default_retention: The default retention for the lock.
        """
        pulumi.set(__self__, "default_retention", default_retention)

    @property
    @pulumi.getter(name="defaultRetention")
    def default_retention(self) -> pulumi.Input['BucketLockConfigurationRuleDefaultRetentionArgs']:
        """
        The default retention for the lock.
        """
        return pulumi.get(self, "default_retention")

    @default_retention.setter
    def default_retention(self, value: pulumi.Input['BucketLockConfigurationRuleDefaultRetentionArgs']):
        pulumi.set(self, "default_retention", value)


@pulumi.input_type
class BucketLockConfigurationRuleDefaultRetentionArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str],
                 days: Optional[pulumi.Input[int]] = None,
                 years: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] mode: The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are `GOVERNANCE` or `COMPLIANCE`. To learn more about the difference between these modes, see [Object Lock retention modes](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/#retention-modes).
        :param pulumi.Input[int] days: The number of days that you want to specify for the default retention period.
        :param pulumi.Input[int] years: The number of years that you want to specify for the default retention period.
        """
        pulumi.set(__self__, "mode", mode)
        if days is not None:
            pulumi.set(__self__, "days", days)
        if years is not None:
            pulumi.set(__self__, "years", years)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are `GOVERNANCE` or `COMPLIANCE`. To learn more about the difference between these modes, see [Object Lock retention modes](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/#retention-modes).
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days that you want to specify for the default retention period.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter
    def years(self) -> Optional[pulumi.Input[int]]:
        """
        The number of years that you want to specify for the default retention period.
        """
        return pulumi.get(self, "years")

    @years.setter
    def years(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "years", value)


@pulumi.input_type
class BucketVersioningArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class BucketWebsiteConfigurationErrorDocumentArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The object key name to use when a 4XX class error occurs.
        """
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The object key name to use when a 4XX class error occurs.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class BucketWebsiteConfigurationIndexDocumentArgs:
    def __init__(__self__, *,
                 suffix: pulumi.Input[str]):
        """
        :param pulumi.Input[str] suffix: A suffix that is appended to a request that is for a directory on the website endpoint.
               
               > **Important:** The suffix must not be empty and must not include a slash character. The routing is not supported.
        """
        pulumi.set(__self__, "suffix", suffix)

    @property
    @pulumi.getter
    def suffix(self) -> pulumi.Input[str]:
        """
        A suffix that is appended to a request that is for a directory on the website endpoint.

        > **Important:** The suffix must not be empty and must not include a slash character. The routing is not supported.
        """
        return pulumi.get(self, "suffix")

    @suffix.setter
    def suffix(self, value: pulumi.Input[str]):
        pulumi.set(self, "suffix", value)


