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
    'SNSCredentialsPermissions',
    'SQSCredentialsPermissions',
]

@pulumi.output_type
class SNSCredentialsPermissions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "canManage":
            suggest = "can_manage"
        elif key == "canPublish":
            suggest = "can_publish"
        elif key == "canReceive":
            suggest = "can_receive"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SNSCredentialsPermissions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SNSCredentialsPermissions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SNSCredentialsPermissions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 can_manage: Optional[bool] = None,
                 can_publish: Optional[bool] = None,
                 can_receive: Optional[bool] = None):
        """
        :param bool can_manage: . Defines if user can manage the associated resource(s).
        :param bool can_publish: . Defines if user can publish messages to the service.
        :param bool can_receive: . Defines if user can receive messages from the service.
        """
        if can_manage is not None:
            pulumi.set(__self__, "can_manage", can_manage)
        if can_publish is not None:
            pulumi.set(__self__, "can_publish", can_publish)
        if can_receive is not None:
            pulumi.set(__self__, "can_receive", can_receive)

    @property
    @pulumi.getter(name="canManage")
    def can_manage(self) -> Optional[bool]:
        """
        . Defines if user can manage the associated resource(s).
        """
        return pulumi.get(self, "can_manage")

    @property
    @pulumi.getter(name="canPublish")
    def can_publish(self) -> Optional[bool]:
        """
        . Defines if user can publish messages to the service.
        """
        return pulumi.get(self, "can_publish")

    @property
    @pulumi.getter(name="canReceive")
    def can_receive(self) -> Optional[bool]:
        """
        . Defines if user can receive messages from the service.
        """
        return pulumi.get(self, "can_receive")


@pulumi.output_type
class SQSCredentialsPermissions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "canManage":
            suggest = "can_manage"
        elif key == "canPublish":
            suggest = "can_publish"
        elif key == "canReceive":
            suggest = "can_receive"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SQSCredentialsPermissions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SQSCredentialsPermissions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SQSCredentialsPermissions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 can_manage: Optional[bool] = None,
                 can_publish: Optional[bool] = None,
                 can_receive: Optional[bool] = None):
        """
        :param bool can_manage: . Defines if user can manage the associated resource(s).
        :param bool can_publish: . Defines if user can publish messages to the service.
        :param bool can_receive: . Defines if user can receive messages from the service.
        """
        if can_manage is not None:
            pulumi.set(__self__, "can_manage", can_manage)
        if can_publish is not None:
            pulumi.set(__self__, "can_publish", can_publish)
        if can_receive is not None:
            pulumi.set(__self__, "can_receive", can_receive)

    @property
    @pulumi.getter(name="canManage")
    def can_manage(self) -> Optional[bool]:
        """
        . Defines if user can manage the associated resource(s).
        """
        return pulumi.get(self, "can_manage")

    @property
    @pulumi.getter(name="canPublish")
    def can_publish(self) -> Optional[bool]:
        """
        . Defines if user can publish messages to the service.
        """
        return pulumi.get(self, "can_publish")

    @property
    @pulumi.getter(name="canReceive")
    def can_receive(self) -> Optional[bool]:
        """
        . Defines if user can receive messages from the service.
        """
        return pulumi.get(self, "can_receive")


