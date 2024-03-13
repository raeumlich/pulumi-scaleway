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

__all__ = [
    'GetWebHostingResult',
    'AwaitableGetWebHostingResult',
    'get_web_hosting',
    'get_web_hosting_output',
]

@pulumi.output_type
class GetWebHostingResult:
    """
    A collection of values returned by getWebHosting.
    """
    def __init__(__self__, cpanel_urls=None, created_at=None, dns_status=None, domain=None, email=None, id=None, offer_id=None, offer_name=None, option_ids=None, options=None, organization_id=None, platform_hostname=None, platform_number=None, project_id=None, region=None, status=None, tags=None, updated_at=None, username=None, webhosting_id=None):
        if cpanel_urls and not isinstance(cpanel_urls, list):
            raise TypeError("Expected argument 'cpanel_urls' to be a list")
        pulumi.set(__self__, "cpanel_urls", cpanel_urls)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dns_status and not isinstance(dns_status, str):
            raise TypeError("Expected argument 'dns_status' to be a str")
        pulumi.set(__self__, "dns_status", dns_status)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if offer_id and not isinstance(offer_id, str):
            raise TypeError("Expected argument 'offer_id' to be a str")
        pulumi.set(__self__, "offer_id", offer_id)
        if offer_name and not isinstance(offer_name, str):
            raise TypeError("Expected argument 'offer_name' to be a str")
        pulumi.set(__self__, "offer_name", offer_name)
        if option_ids and not isinstance(option_ids, list):
            raise TypeError("Expected argument 'option_ids' to be a list")
        pulumi.set(__self__, "option_ids", option_ids)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if platform_hostname and not isinstance(platform_hostname, str):
            raise TypeError("Expected argument 'platform_hostname' to be a str")
        pulumi.set(__self__, "platform_hostname", platform_hostname)
        if platform_number and not isinstance(platform_number, int):
            raise TypeError("Expected argument 'platform_number' to be a int")
        pulumi.set(__self__, "platform_number", platform_number)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if webhosting_id and not isinstance(webhosting_id, str):
            raise TypeError("Expected argument 'webhosting_id' to be a str")
        pulumi.set(__self__, "webhosting_id", webhosting_id)

    @property
    @pulumi.getter(name="cpanelUrls")
    def cpanel_urls(self) -> Sequence['outputs.GetWebHostingCpanelUrlResult']:
        return pulumi.get(self, "cpanel_urls")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dnsStatus")
    def dns_status(self) -> str:
        return pulumi.get(self, "dns_status")

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="offerId")
    def offer_id(self) -> str:
        return pulumi.get(self, "offer_id")

    @property
    @pulumi.getter(name="offerName")
    def offer_name(self) -> str:
        return pulumi.get(self, "offer_name")

    @property
    @pulumi.getter(name="optionIds")
    def option_ids(self) -> Sequence[str]:
        return pulumi.get(self, "option_ids")

    @property
    @pulumi.getter
    def options(self) -> Sequence['outputs.GetWebHostingOptionResult']:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="platformHostname")
    def platform_hostname(self) -> str:
        return pulumi.get(self, "platform_hostname")

    @property
    @pulumi.getter(name="platformNumber")
    def platform_number(self) -> int:
        return pulumi.get(self, "platform_number")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="webhostingId")
    def webhosting_id(self) -> Optional[str]:
        return pulumi.get(self, "webhosting_id")


class AwaitableGetWebHostingResult(GetWebHostingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebHostingResult(
            cpanel_urls=self.cpanel_urls,
            created_at=self.created_at,
            dns_status=self.dns_status,
            domain=self.domain,
            email=self.email,
            id=self.id,
            offer_id=self.offer_id,
            offer_name=self.offer_name,
            option_ids=self.option_ids,
            options=self.options,
            organization_id=self.organization_id,
            platform_hostname=self.platform_hostname,
            platform_number=self.platform_number,
            project_id=self.project_id,
            region=self.region,
            status=self.status,
            tags=self.tags,
            updated_at=self.updated_at,
            username=self.username,
            webhosting_id=self.webhosting_id)


def get_web_hosting(domain: Optional[str] = None,
                    organization_id: Optional[str] = None,
                    project_id: Optional[str] = None,
                    webhosting_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebHostingResult:
    """
    Gets information about a webhosting.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_domain = scaleway.webhosting.get_web_hosting(domain="foobar.com")
    by_id = scaleway.webhosting.get_web_hosting(webhosting_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str domain: The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
    :param str organization_id: The ID of the organization the hosting is associated with.
    :param str project_id: `project_id`) The ID of the project the hosting is associated with.
    :param str webhosting_id: The hosting id. Only one of `domain` and `webhosting_id` should be specified.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['organizationId'] = organization_id
    __args__['projectId'] = project_id
    __args__['webhostingId'] = webhosting_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:webhosting/getWebHosting:getWebHosting', __args__, opts=opts, typ=GetWebHostingResult).value

    return AwaitableGetWebHostingResult(
        cpanel_urls=pulumi.get(__ret__, 'cpanel_urls'),
        created_at=pulumi.get(__ret__, 'created_at'),
        dns_status=pulumi.get(__ret__, 'dns_status'),
        domain=pulumi.get(__ret__, 'domain'),
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        offer_id=pulumi.get(__ret__, 'offer_id'),
        offer_name=pulumi.get(__ret__, 'offer_name'),
        option_ids=pulumi.get(__ret__, 'option_ids'),
        options=pulumi.get(__ret__, 'options'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        platform_hostname=pulumi.get(__ret__, 'platform_hostname'),
        platform_number=pulumi.get(__ret__, 'platform_number'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        username=pulumi.get(__ret__, 'username'),
        webhosting_id=pulumi.get(__ret__, 'webhosting_id'))


@_utilities.lift_output_func(get_web_hosting)
def get_web_hosting_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                           organization_id: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[Optional[str]]] = None,
                           webhosting_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWebHostingResult]:
    """
    Gets information about a webhosting.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    by_domain = scaleway.webhosting.get_web_hosting(domain="foobar.com")
    by_id = scaleway.webhosting.get_web_hosting(webhosting_id="11111111-1111-1111-1111-111111111111")
    ```
    <!--End PulumiCodeChooser -->


    :param str domain: The hosting domain name. Only one of `domain` and `webhosting_id` should be specified.
    :param str organization_id: The ID of the organization the hosting is associated with.
    :param str project_id: `project_id`) The ID of the project the hosting is associated with.
    :param str webhosting_id: The hosting id. Only one of `domain` and `webhosting_id` should be specified.
    """
    ...
