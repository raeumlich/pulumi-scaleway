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
    'CockpitEndpointArgs',
    'TokenScopesArgs',
]

@pulumi.input_type
class CockpitEndpointArgs:
    def __init__(__self__, *,
                 alertmanager_url: Optional[pulumi.Input[str]] = None,
                 grafana_url: Optional[pulumi.Input[str]] = None,
                 logs_url: Optional[pulumi.Input[str]] = None,
                 metrics_url: Optional[pulumi.Input[str]] = None,
                 traces_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] alertmanager_url: The alertmanager URL.
        :param pulumi.Input[str] grafana_url: The grafana URL.
        :param pulumi.Input[str] logs_url: The logs URL.
        :param pulumi.Input[str] metrics_url: The metrics URL.
        :param pulumi.Input[str] traces_url: The traces URL.
        """
        if alertmanager_url is not None:
            pulumi.set(__self__, "alertmanager_url", alertmanager_url)
        if grafana_url is not None:
            pulumi.set(__self__, "grafana_url", grafana_url)
        if logs_url is not None:
            pulumi.set(__self__, "logs_url", logs_url)
        if metrics_url is not None:
            pulumi.set(__self__, "metrics_url", metrics_url)
        if traces_url is not None:
            pulumi.set(__self__, "traces_url", traces_url)

    @property
    @pulumi.getter(name="alertmanagerUrl")
    def alertmanager_url(self) -> Optional[pulumi.Input[str]]:
        """
        The alertmanager URL.
        """
        return pulumi.get(self, "alertmanager_url")

    @alertmanager_url.setter
    def alertmanager_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alertmanager_url", value)

    @property
    @pulumi.getter(name="grafanaUrl")
    def grafana_url(self) -> Optional[pulumi.Input[str]]:
        """
        The grafana URL.
        """
        return pulumi.get(self, "grafana_url")

    @grafana_url.setter
    def grafana_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grafana_url", value)

    @property
    @pulumi.getter(name="logsUrl")
    def logs_url(self) -> Optional[pulumi.Input[str]]:
        """
        The logs URL.
        """
        return pulumi.get(self, "logs_url")

    @logs_url.setter
    def logs_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logs_url", value)

    @property
    @pulumi.getter(name="metricsUrl")
    def metrics_url(self) -> Optional[pulumi.Input[str]]:
        """
        The metrics URL.
        """
        return pulumi.get(self, "metrics_url")

    @metrics_url.setter
    def metrics_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metrics_url", value)

    @property
    @pulumi.getter(name="tracesUrl")
    def traces_url(self) -> Optional[pulumi.Input[str]]:
        """
        The traces URL.
        """
        return pulumi.get(self, "traces_url")

    @traces_url.setter
    def traces_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traces_url", value)


@pulumi.input_type
class TokenScopesArgs:
    def __init__(__self__, *,
                 query_logs: Optional[pulumi.Input[bool]] = None,
                 query_metrics: Optional[pulumi.Input[bool]] = None,
                 query_traces: Optional[pulumi.Input[bool]] = None,
                 setup_alerts: Optional[pulumi.Input[bool]] = None,
                 setup_logs_rules: Optional[pulumi.Input[bool]] = None,
                 setup_metrics_rules: Optional[pulumi.Input[bool]] = None,
                 write_logs: Optional[pulumi.Input[bool]] = None,
                 write_metrics: Optional[pulumi.Input[bool]] = None,
                 write_traces: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] query_logs: Query logs.
        :param pulumi.Input[bool] query_metrics: Query metrics.
        :param pulumi.Input[bool] query_traces: Query traces.
        :param pulumi.Input[bool] setup_alerts: Setup alerts.
        :param pulumi.Input[bool] setup_logs_rules: Setup logs rules.
        :param pulumi.Input[bool] setup_metrics_rules: Setup metrics rules.
        :param pulumi.Input[bool] write_logs: Write logs.
        :param pulumi.Input[bool] write_metrics: Write metrics.
        :param pulumi.Input[bool] write_traces: Write traces.
        """
        if query_logs is not None:
            pulumi.set(__self__, "query_logs", query_logs)
        if query_metrics is not None:
            pulumi.set(__self__, "query_metrics", query_metrics)
        if query_traces is not None:
            pulumi.set(__self__, "query_traces", query_traces)
        if setup_alerts is not None:
            pulumi.set(__self__, "setup_alerts", setup_alerts)
        if setup_logs_rules is not None:
            pulumi.set(__self__, "setup_logs_rules", setup_logs_rules)
        if setup_metrics_rules is not None:
            pulumi.set(__self__, "setup_metrics_rules", setup_metrics_rules)
        if write_logs is not None:
            pulumi.set(__self__, "write_logs", write_logs)
        if write_metrics is not None:
            pulumi.set(__self__, "write_metrics", write_metrics)
        if write_traces is not None:
            pulumi.set(__self__, "write_traces", write_traces)

    @property
    @pulumi.getter(name="queryLogs")
    def query_logs(self) -> Optional[pulumi.Input[bool]]:
        """
        Query logs.
        """
        return pulumi.get(self, "query_logs")

    @query_logs.setter
    def query_logs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "query_logs", value)

    @property
    @pulumi.getter(name="queryMetrics")
    def query_metrics(self) -> Optional[pulumi.Input[bool]]:
        """
        Query metrics.
        """
        return pulumi.get(self, "query_metrics")

    @query_metrics.setter
    def query_metrics(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "query_metrics", value)

    @property
    @pulumi.getter(name="queryTraces")
    def query_traces(self) -> Optional[pulumi.Input[bool]]:
        """
        Query traces.
        """
        return pulumi.get(self, "query_traces")

    @query_traces.setter
    def query_traces(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "query_traces", value)

    @property
    @pulumi.getter(name="setupAlerts")
    def setup_alerts(self) -> Optional[pulumi.Input[bool]]:
        """
        Setup alerts.
        """
        return pulumi.get(self, "setup_alerts")

    @setup_alerts.setter
    def setup_alerts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "setup_alerts", value)

    @property
    @pulumi.getter(name="setupLogsRules")
    def setup_logs_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Setup logs rules.
        """
        return pulumi.get(self, "setup_logs_rules")

    @setup_logs_rules.setter
    def setup_logs_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "setup_logs_rules", value)

    @property
    @pulumi.getter(name="setupMetricsRules")
    def setup_metrics_rules(self) -> Optional[pulumi.Input[bool]]:
        """
        Setup metrics rules.
        """
        return pulumi.get(self, "setup_metrics_rules")

    @setup_metrics_rules.setter
    def setup_metrics_rules(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "setup_metrics_rules", value)

    @property
    @pulumi.getter(name="writeLogs")
    def write_logs(self) -> Optional[pulumi.Input[bool]]:
        """
        Write logs.
        """
        return pulumi.get(self, "write_logs")

    @write_logs.setter
    def write_logs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write_logs", value)

    @property
    @pulumi.getter(name="writeMetrics")
    def write_metrics(self) -> Optional[pulumi.Input[bool]]:
        """
        Write metrics.
        """
        return pulumi.get(self, "write_metrics")

    @write_metrics.setter
    def write_metrics(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write_metrics", value)

    @property
    @pulumi.getter(name="writeTraces")
    def write_traces(self) -> Optional[pulumi.Input[bool]]:
        """
        Write traces.
        """
        return pulumi.get(self, "write_traces")

    @write_traces.setter
    def write_traces(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write_traces", value)

