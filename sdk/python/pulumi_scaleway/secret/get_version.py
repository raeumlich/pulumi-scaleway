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
    def __init__(__self__, created_at=None, data=None, description=None, id=None, project_id=None, region=None, revision=None, secret_id=None, secret_name=None, status=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if data and not isinstance(data, str):
            raise TypeError("Expected argument 'data' to be a str")
        pulumi.set(__self__, "data", data)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if revision and not isinstance(revision, str):
            raise TypeError("Expected argument 'revision' to be a str")
        pulumi.set(__self__, "revision", revision)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if secret_name and not isinstance(secret_name, str):
            raise TypeError("Expected argument 'secret_name' to be a str")
        pulumi.set(__self__, "secret_name", secret_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date and time of secret version's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def data(self) -> str:
        """
        The data payload of the secret version. more on the data section
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (Optional) Description of the secret version (e.g. `my-new-description`).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def revision(self) -> Optional[str]:
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[str]:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[str]:
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the Secret Version.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date and time of secret version's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetVersionResult(GetVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVersionResult(
            created_at=self.created_at,
            data=self.data,
            description=self.description,
            id=self.id,
            project_id=self.project_id,
            region=self.region,
            revision=self.revision,
            secret_id=self.secret_id,
            secret_name=self.secret_name,
            status=self.status,
            updated_at=self.updated_at)


def get_version(project_id: Optional[str] = None,
                region: Optional[str] = None,
                revision: Optional[str] = None,
                secret_id: Optional[str] = None,
                secret_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVersionResult:
    """
    Gets information about Scaleway a Secret Version.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).

    ## Examples

    ### Basic

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main_secret = scaleway.secret.Secret("mainSecret", description="barr")
    main_version = scaleway.secret.Version("mainVersion",
        description="your description",
        secret_id=main_secret.id,
        data="your_secret")
    data_by_secret_id = scaleway.secret.get_version_output(secret_id=main_secret.id,
        revision="1")
    data_by_secret_name = scaleway.secret.get_version_output(secret_name=main_secret.name,
        revision="1")
    pulumi.export("scalewaySecretAccessPayload", data_by_secret_name.data)
    pulumi.export("scalewaySecretAccessPayloadById", data_by_secret_id.data)
    ```
    <!--End PulumiCodeChooser -->

    ## Data

    Note: This Data Source give you **access** to the secret payload encoded en base64.

    Be aware that this is a sensitive attribute. For more information,
    see Sensitive Data in State.

    > **Important:**  This property is sensitive and will not be displayed in the plan.


    :param str project_id: The ID of the project the Secret version is associated with.
    :param str region: `region`) The region
           in which the resource exists.
    :param str revision: The revision for this Secret Version.
    :param str secret_id: The Secret ID associated wit the secret version.
           Only one of `secret_id` and `secret_name` should be specified.
    :param str secret_name: The Name of Secret associated wit the secret version.
           Only one of `secret_id` and `secret_name` should be specified.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['revision'] = revision
    __args__['secretId'] = secret_id
    __args__['secretName'] = secret_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:secret/getVersion:getVersion', __args__, opts=opts, typ=GetVersionResult).value

    return AwaitableGetVersionResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        data=pulumi.get(__ret__, 'data'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        revision=pulumi.get(__ret__, 'revision'),
        secret_id=pulumi.get(__ret__, 'secret_id'),
        secret_name=pulumi.get(__ret__, 'secret_name'),
        status=pulumi.get(__ret__, 'status'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_version)
def get_version_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       revision: Optional[pulumi.Input[Optional[str]]] = None,
                       secret_id: Optional[pulumi.Input[Optional[str]]] = None,
                       secret_name: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVersionResult]:
    """
    Gets information about Scaleway a Secret Version.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/#secret-versions-079501).

    ## Examples

    ### Basic

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main_secret = scaleway.secret.Secret("mainSecret", description="barr")
    main_version = scaleway.secret.Version("mainVersion",
        description="your description",
        secret_id=main_secret.id,
        data="your_secret")
    data_by_secret_id = scaleway.secret.get_version_output(secret_id=main_secret.id,
        revision="1")
    data_by_secret_name = scaleway.secret.get_version_output(secret_name=main_secret.name,
        revision="1")
    pulumi.export("scalewaySecretAccessPayload", data_by_secret_name.data)
    pulumi.export("scalewaySecretAccessPayloadById", data_by_secret_id.data)
    ```
    <!--End PulumiCodeChooser -->

    ## Data

    Note: This Data Source give you **access** to the secret payload encoded en base64.

    Be aware that this is a sensitive attribute. For more information,
    see Sensitive Data in State.

    > **Important:**  This property is sensitive and will not be displayed in the plan.


    :param str project_id: The ID of the project the Secret version is associated with.
    :param str region: `region`) The region
           in which the resource exists.
    :param str revision: The revision for this Secret Version.
    :param str secret_id: The Secret ID associated wit the secret version.
           Only one of `secret_id` and `secret_name` should be specified.
    :param str secret_name: The Name of Secret associated wit the secret version.
           Only one of `secret_id` and `secret_name` should be specified.
    """
    ...
