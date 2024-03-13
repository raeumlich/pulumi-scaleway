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
    'GetImageResult',
    'AwaitableGetImageResult',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class GetImageResult:
    """
    A collection of values returned by getImage.
    """
    def __init__(__self__, additional_volume_ids=None, architecture=None, creation_date=None, default_bootscript_id=None, from_server_id=None, id=None, image_id=None, latest=None, modification_date=None, name=None, organization_id=None, project_id=None, public=None, root_volume_id=None, state=None, zone=None):
        if additional_volume_ids and not isinstance(additional_volume_ids, list):
            raise TypeError("Expected argument 'additional_volume_ids' to be a list")
        pulumi.set(__self__, "additional_volume_ids", additional_volume_ids)
        if architecture and not isinstance(architecture, str):
            raise TypeError("Expected argument 'architecture' to be a str")
        pulumi.set(__self__, "architecture", architecture)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if default_bootscript_id and not isinstance(default_bootscript_id, str):
            raise TypeError("Expected argument 'default_bootscript_id' to be a str")
        pulumi.set(__self__, "default_bootscript_id", default_bootscript_id)
        if from_server_id and not isinstance(from_server_id, str):
            raise TypeError("Expected argument 'from_server_id' to be a str")
        pulumi.set(__self__, "from_server_id", from_server_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if latest and not isinstance(latest, bool):
            raise TypeError("Expected argument 'latest' to be a bool")
        pulumi.set(__self__, "latest", latest)
        if modification_date and not isinstance(modification_date, str):
            raise TypeError("Expected argument 'modification_date' to be a str")
        pulumi.set(__self__, "modification_date", modification_date)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public and not isinstance(public, bool):
            raise TypeError("Expected argument 'public' to be a bool")
        pulumi.set(__self__, "public", public)
        if root_volume_id and not isinstance(root_volume_id, str):
            raise TypeError("Expected argument 'root_volume_id' to be a str")
        pulumi.set(__self__, "root_volume_id", root_volume_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="additionalVolumeIds")
    def additional_volume_ids(self) -> Sequence[str]:
        """
        IDs of the additional volumes in this image.
        """
        return pulumi.get(self, "additional_volume_ids")

    @property
    @pulumi.getter
    def architecture(self) -> Optional[str]:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Date of the image creation.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="defaultBootscriptId")
    def default_bootscript_id(self) -> str:
        """
        ID of the default bootscript for this image.
        """
        return pulumi.get(self, "default_bootscript_id")

    @property
    @pulumi.getter(name="fromServerId")
    def from_server_id(self) -> str:
        """
        ID of the server the image if based from.
        """
        return pulumi.get(self, "from_server_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter
    def latest(self) -> Optional[bool]:
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> str:
        """
        Date of image latest update.
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The ID of the organization the image is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the project the image is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def public(self) -> bool:
        """
        Set to `true` if the image is public.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter(name="rootVolumeId")
    def root_volume_id(self) -> str:
        """
        ID of the root volume in this image.
        """
        return pulumi.get(self, "root_volume_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the image. Possible values are: `available`, `creating` or `error`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetImageResult(GetImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageResult(
            additional_volume_ids=self.additional_volume_ids,
            architecture=self.architecture,
            creation_date=self.creation_date,
            default_bootscript_id=self.default_bootscript_id,
            from_server_id=self.from_server_id,
            id=self.id,
            image_id=self.image_id,
            latest=self.latest,
            modification_date=self.modification_date,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            public=self.public,
            root_volume_id=self.root_volume_id,
            state=self.state,
            zone=self.zone)


def get_image(architecture: Optional[str] = None,
              image_id: Optional[str] = None,
              latest: Optional[bool] = None,
              name: Optional[str] = None,
              project_id: Optional[str] = None,
              zone: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageResult:
    """
    Gets information about an instance image.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_image = scaleway.instance.get_image(image_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
    :param str image_id: The image id. Only one of `name` and `image_id` should be specified.
    :param bool latest: Use the latest image ID.
    :param str name: The image name. Only one of `name` and `image_id` should be specified.
    :param str project_id: The ID of the project the image is associated with.
    :param str zone: `zone`) The zone in which the image exists.
    """
    __args__ = dict()
    __args__['architecture'] = architecture
    __args__['imageId'] = image_id
    __args__['latest'] = latest
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:instance/getImage:getImage', __args__, opts=opts, typ=GetImageResult).value

    return AwaitableGetImageResult(
        additional_volume_ids=pulumi.get(__ret__, 'additional_volume_ids'),
        architecture=pulumi.get(__ret__, 'architecture'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        default_bootscript_id=pulumi.get(__ret__, 'default_bootscript_id'),
        from_server_id=pulumi.get(__ret__, 'from_server_id'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        latest=pulumi.get(__ret__, 'latest'),
        modification_date=pulumi.get(__ret__, 'modification_date'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public=pulumi.get(__ret__, 'public'),
        root_volume_id=pulumi.get(__ret__, 'root_volume_id'),
        state=pulumi.get(__ret__, 'state'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_image)
def get_image_output(architecture: Optional[pulumi.Input[Optional[str]]] = None,
                     image_id: Optional[pulumi.Input[Optional[str]]] = None,
                     latest: Optional[pulumi.Input[Optional[bool]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     project_id: Optional[pulumi.Input[Optional[str]]] = None,
                     zone: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageResult]:
    """
    Gets information about an instance image.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_image = scaleway.instance.get_image(image_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
    :param str image_id: The image id. Only one of `name` and `image_id` should be specified.
    :param bool latest: Use the latest image ID.
    :param str name: The image name. Only one of `name` and `image_id` should be specified.
    :param str project_id: The ID of the project the image is associated with.
    :param str zone: `zone`) The zone in which the image exists.
    """
    ...
