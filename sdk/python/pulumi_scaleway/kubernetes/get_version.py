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
    'GetVersionResult',
    'AwaitableGetVersionResult',
    'get_version',
    'get_version_output',
]

@pulumi.output_type
class GetVersionResult:
    """
    A collection of values returned by getVersion.
    """
    def __init__(__self__, available_cnis=None, available_container_runtimes=None, available_feature_gates=None, id=None, name=None, region=None):
        if available_cnis and not isinstance(available_cnis, list):
            raise TypeError("Expected argument 'available_cnis' to be a list")
        pulumi.set(__self__, "available_cnis", available_cnis)
        if available_container_runtimes and not isinstance(available_container_runtimes, list):
            raise TypeError("Expected argument 'available_container_runtimes' to be a list")
        pulumi.set(__self__, "available_container_runtimes", available_container_runtimes)
        if available_feature_gates and not isinstance(available_feature_gates, list):
            raise TypeError("Expected argument 'available_feature_gates' to be a list")
        pulumi.set(__self__, "available_feature_gates", available_feature_gates)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="availableCnis")
    def available_cnis(self) -> Sequence[str]:
        """
        The list of supported Container Network Interface (CNI) plugins for this version.
        """
        return pulumi.get(self, "available_cnis")

    @property
    @pulumi.getter(name="availableContainerRuntimes")
    def available_container_runtimes(self) -> Sequence[str]:
        """
        The list of supported container runtimes for this version.
        """
        return pulumi.get(self, "available_container_runtimes")

    @property
    @pulumi.getter(name="availableFeatureGates")
    def available_feature_gates(self) -> Sequence[str]:
        """
        The list of supported feature gates for this version.
        """
        return pulumi.get(self, "available_feature_gates")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            available_cnis=self.available_cnis,
            available_container_runtimes=self.available_container_runtimes,
            available_feature_gates=self.available_feature_gates,
            id=self.id,
            name=self.name,
            region=self.region)


def get_version(name: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Gets information about a Kubernetes version.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/k8s/api).

    You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.

    ## Example Usage

    ### Use the latest version

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    latest = scaleway.kubernetes.get_version(name="latest")
    ```
    <!--End PulumiCodeChooser -->

    ### Use a specific version

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.kubernetes.get_version(name="1.26.0")
    ```
    <!--End PulumiCodeChooser -->


    :param str name: The name of the Kubernetes version.
    :param str region: `region`) The region in which the version exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:kubernetes/getVersion:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        available_cnis=pulumi.get(__ret__, 'available_cnis'),
        available_container_runtimes=pulumi.get(__ret__, 'available_container_runtimes'),
        available_feature_gates=pulumi.get(__ret__, 'available_feature_gates'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'))


@_utilities.lift_output_func(get_version)
def get_version_output(name: Optional[pulumi.Input[str]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Gets information about a Kubernetes version.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/k8s/api).

    You can also use the [scaleway-cli](https://github.com/scaleway/scaleway-cli) with `scw k8s version list` to list all available versions.

    ## Example Usage

    ### Use the latest version

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    latest = scaleway.kubernetes.get_version(name="latest")
    ```
    <!--End PulumiCodeChooser -->

    ### Use a specific version

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_name = scaleway.kubernetes.get_version(name="1.26.0")
    ```
    <!--End PulumiCodeChooser -->


    :param str name: The name of the Kubernetes version.
    :param str region: `region`) The region in which the version exists.
    """
    ...
