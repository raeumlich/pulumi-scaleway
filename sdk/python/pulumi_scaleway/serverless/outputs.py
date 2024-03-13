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
    'ContainerTriggerNats',
    'ContainerTriggerSqs',
    'FunctionTriggerNats',
    'FunctionTriggerSqs',
    'JobDefinitionCron',
]

@pulumi.output_type
class ContainerTriggerNats(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerTriggerNats. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerTriggerNats.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerTriggerNats.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 subject: str,
                 account_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None):
        """
        :param str subject: The subject to listen to
        :param str account_id: ID of the mnq nats account.
        :param str project_id: ID of the project that contain the mnq nats account, defaults to provider's project
        :param str region: `region`). The region in which the namespace should be created.
        """
        pulumi.set(__self__, "subject", subject)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def subject(self) -> str:
        """
        The subject to listen to
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        """
        ID of the mnq nats account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        ID of the project that contain the mnq nats account, defaults to provider's project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class ContainerTriggerSqs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "namespaceId":
            suggest = "namespace_id"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerTriggerSqs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerTriggerSqs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerTriggerSqs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 queue: str,
                 namespace_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None):
        """
        :param str queue: Name of the queue
        :param str namespace_id: ID of the mnq namespace. Deprecated.
        :param str project_id: ID of the project that contain the mnq nats account, defaults to provider's project
        :param str region: `region`). The region in which the namespace should be created.
        """
        pulumi.set(__self__, "queue", queue)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def queue(self) -> str:
        """
        Name of the queue
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[str]:
        """
        ID of the mnq namespace. Deprecated.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        ID of the project that contain the mnq nats account, defaults to provider's project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class FunctionTriggerNats(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FunctionTriggerNats. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FunctionTriggerNats.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FunctionTriggerNats.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 subject: str,
                 account_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None):
        """
        :param str subject: The subject to listen to
        :param str account_id: ID of the mnq nats account.
        :param str project_id: ID of the project that contain the mnq nats account, defaults to provider's project
        :param str region: `region`). The region in which the namespace should be created.
        """
        pulumi.set(__self__, "subject", subject)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def subject(self) -> str:
        """
        The subject to listen to
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[str]:
        """
        ID of the mnq nats account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        ID of the project that contain the mnq nats account, defaults to provider's project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class FunctionTriggerSqs(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "namespaceId":
            suggest = "namespace_id"
        elif key == "projectId":
            suggest = "project_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FunctionTriggerSqs. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FunctionTriggerSqs.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FunctionTriggerSqs.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 queue: str,
                 namespace_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None):
        """
        :param str queue: Name of the queue
        :param str namespace_id: ID of the mnq namespace. Deprecated.
        :param str project_id: ID of the project that contain the mnq nats account, defaults to provider's project
        :param str region: `region`). The region in which the namespace should be created.
        """
        pulumi.set(__self__, "queue", queue)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def queue(self) -> str:
        """
        Name of the queue
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[str]:
        """
        ID of the mnq namespace. Deprecated.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        ID of the project that contain the mnq nats account, defaults to provider's project
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")


@pulumi.output_type
class JobDefinitionCron(dict):
    def __init__(__self__, *,
                 schedule: str,
                 timezone: str):
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def schedule(self) -> str:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def timezone(self) -> str:
        return pulumi.get(self, "timezone")


