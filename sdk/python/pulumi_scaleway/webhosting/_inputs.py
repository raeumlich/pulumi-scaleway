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
    'WebHostingCpanelUrlArgs',
    'WebHostingOptionArgs',
]

@pulumi.input_type
class WebHostingCpanelUrlArgs:
    def __init__(__self__, *,
                 dashboard: Optional[pulumi.Input[str]] = None,
                 webmail: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] dashboard: The URL of the Dashboard.
        :param pulumi.Input[str] webmail: The URL of the Webmail interface.
        """
        if dashboard is not None:
            pulumi.set(__self__, "dashboard", dashboard)
        if webmail is not None:
            pulumi.set(__self__, "webmail", webmail)

    @property
    @pulumi.getter
    def dashboard(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the Dashboard.
        """
        return pulumi.get(self, "dashboard")

    @dashboard.setter
    def dashboard(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dashboard", value)

    @property
    @pulumi.getter
    def webmail(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the Webmail interface.
        """
        return pulumi.get(self, "webmail")

    @webmail.setter
    def webmail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webmail", value)


@pulumi.input_type
class WebHostingOptionArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The option ID.
        :param pulumi.Input[str] name: The option name.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The option ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The option name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


