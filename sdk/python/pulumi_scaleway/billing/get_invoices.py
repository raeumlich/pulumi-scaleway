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
    'GetInvoicesResult',
    'AwaitableGetInvoicesResult',
    'get_invoices',
    'get_invoices_output',
]

@pulumi.output_type
class GetInvoicesResult:
    """
    A collection of values returned by getInvoices.
    """
    def __init__(__self__, id=None, invoice_type=None, invoices=None, organization_id=None, started_after=None, started_before=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invoice_type and not isinstance(invoice_type, str):
            raise TypeError("Expected argument 'invoice_type' to be a str")
        pulumi.set(__self__, "invoice_type", invoice_type)
        if invoices and not isinstance(invoices, list):
            raise TypeError("Expected argument 'invoices' to be a list")
        pulumi.set(__self__, "invoices", invoices)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if started_after and not isinstance(started_after, str):
            raise TypeError("Expected argument 'started_after' to be a str")
        pulumi.set(__self__, "started_after", started_after)
        if started_before and not isinstance(started_before, str):
            raise TypeError("Expected argument 'started_before' to be a str")
        pulumi.set(__self__, "started_before", started_before)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invoiceType")
    def invoice_type(self) -> Optional[str]:
        """
        The type of invoice.
        """
        return pulumi.get(self, "invoice_type")

    @property
    @pulumi.getter
    def invoices(self) -> Sequence['outputs.GetInvoicesInvoiceResult']:
        """
        List of found invoices
        """
        return pulumi.get(self, "invoices")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="startedAfter")
    def started_after(self) -> Optional[str]:
        return pulumi.get(self, "started_after")

    @property
    @pulumi.getter(name="startedBefore")
    def started_before(self) -> Optional[str]:
        return pulumi.get(self, "started_before")


class AwaitableGetInvoicesResult(GetInvoicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInvoicesResult(
            id=self.id,
            invoice_type=self.invoice_type,
            invoices=self.invoices,
            organization_id=self.organization_id,
            started_after=self.started_after,
            started_before=self.started_before)


def get_invoices(invoice_type: Optional[str] = None,
                 started_after: Optional[str] = None,
                 started_before: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInvoicesResult:
    """
    Gets information about your Invoices.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_invoices = scaleway.billing.get_invoices(invoice_type="periodic")
    ```


    :param str invoice_type: Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
    :param str started_after: Invoices with a start date that are greater or equal to `started_after` are listed (RFC 3339 format).
    :param str started_before: Invoices with a start date that precedes `started_before` are listed (RFC 3339 format).
    """
    __args__ = dict()
    __args__['invoiceType'] = invoice_type
    __args__['startedAfter'] = started_after
    __args__['startedBefore'] = started_before
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:billing/getInvoices:getInvoices', __args__, opts=opts, typ=GetInvoicesResult).value

    return AwaitableGetInvoicesResult(
        id=pulumi.get(__ret__, 'id'),
        invoice_type=pulumi.get(__ret__, 'invoice_type'),
        invoices=pulumi.get(__ret__, 'invoices'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        started_after=pulumi.get(__ret__, 'started_after'),
        started_before=pulumi.get(__ret__, 'started_before'))


@_utilities.lift_output_func(get_invoices)
def get_invoices_output(invoice_type: Optional[pulumi.Input[Optional[str]]] = None,
                        started_after: Optional[pulumi.Input[Optional[str]]] = None,
                        started_before: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInvoicesResult]:
    """
    Gets information about your Invoices.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_invoices = scaleway.billing.get_invoices(invoice_type="periodic")
    ```


    :param str invoice_type: Invoices with the given type are listed. Valid values are `periodic` and `purchase`.
    :param str started_after: Invoices with a start date that are greater or equal to `started_after` are listed (RFC 3339 format).
    :param str started_before: Invoices with a start date that precedes `started_before` are listed (RFC 3339 format).
    """
    ...
