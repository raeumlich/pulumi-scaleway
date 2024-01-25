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
    'CockpitEndpoint',
    'TokenScopes',
    'GetCockpitEndpointResult',
]

@pulumi.output_type
class CockpitEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alertmanagerUrl":
            suggest = "alertmanager_url"
        elif key == "grafanaUrl":
            suggest = "grafana_url"
        elif key == "logsUrl":
            suggest = "logs_url"
        elif key == "metricsUrl":
            suggest = "metrics_url"
        elif key == "tracesUrl":
            suggest = "traces_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CockpitEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CockpitEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CockpitEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alertmanager_url: Optional[str] = None,
                 grafana_url: Optional[str] = None,
                 logs_url: Optional[str] = None,
                 metrics_url: Optional[str] = None,
                 traces_url: Optional[str] = None):
        """
        :param str alertmanager_url: The alertmanager URL.
        :param str grafana_url: The grafana URL.
        :param str logs_url: The logs URL.
        :param str metrics_url: The metrics URL.
        :param str traces_url: The traces URL.
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
    def alertmanager_url(self) -> Optional[str]:
        """
        The alertmanager URL.
        """
        return pulumi.get(self, "alertmanager_url")

    @property
    @pulumi.getter(name="grafanaUrl")
    def grafana_url(self) -> Optional[str]:
        """
        The grafana URL.
        """
        return pulumi.get(self, "grafana_url")

    @property
    @pulumi.getter(name="logsUrl")
    def logs_url(self) -> Optional[str]:
        """
        The logs URL.
        """
        return pulumi.get(self, "logs_url")

    @property
    @pulumi.getter(name="metricsUrl")
    def metrics_url(self) -> Optional[str]:
        """
        The metrics URL.
        """
        return pulumi.get(self, "metrics_url")

    @property
    @pulumi.getter(name="tracesUrl")
    def traces_url(self) -> Optional[str]:
        """
        The traces URL.
        """
        return pulumi.get(self, "traces_url")


@pulumi.output_type
class TokenScopes(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "queryLogs":
            suggest = "query_logs"
        elif key == "queryMetrics":
            suggest = "query_metrics"
        elif key == "queryTraces":
            suggest = "query_traces"
        elif key == "setupAlerts":
            suggest = "setup_alerts"
        elif key == "setupLogsRules":
            suggest = "setup_logs_rules"
        elif key == "setupMetricsRules":
            suggest = "setup_metrics_rules"
        elif key == "writeLogs":
            suggest = "write_logs"
        elif key == "writeMetrics":
            suggest = "write_metrics"
        elif key == "writeTraces":
            suggest = "write_traces"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenScopes. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenScopes.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenScopes.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 query_logs: Optional[bool] = None,
                 query_metrics: Optional[bool] = None,
                 query_traces: Optional[bool] = None,
                 setup_alerts: Optional[bool] = None,
                 setup_logs_rules: Optional[bool] = None,
                 setup_metrics_rules: Optional[bool] = None,
                 write_logs: Optional[bool] = None,
                 write_metrics: Optional[bool] = None,
                 write_traces: Optional[bool] = None):
        """
        :param bool query_logs: Query logs.
        :param bool query_metrics: Query metrics.
        :param bool query_traces: Query traces.
        :param bool setup_alerts: Setup alerts.
        :param bool setup_logs_rules: Setup logs rules.
        :param bool setup_metrics_rules: Setup metrics rules.
        :param bool write_logs: Write logs.
        :param bool write_metrics: Write metrics.
        :param bool write_traces: Write traces.
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
    def query_logs(self) -> Optional[bool]:
        """
        Query logs.
        """
        return pulumi.get(self, "query_logs")

    @property
    @pulumi.getter(name="queryMetrics")
    def query_metrics(self) -> Optional[bool]:
        """
        Query metrics.
        """
        return pulumi.get(self, "query_metrics")

    @property
    @pulumi.getter(name="queryTraces")
    def query_traces(self) -> Optional[bool]:
        """
        Query traces.
        """
        return pulumi.get(self, "query_traces")

    @property
    @pulumi.getter(name="setupAlerts")
    def setup_alerts(self) -> Optional[bool]:
        """
        Setup alerts.
        """
        return pulumi.get(self, "setup_alerts")

    @property
    @pulumi.getter(name="setupLogsRules")
    def setup_logs_rules(self) -> Optional[bool]:
        """
        Setup logs rules.
        """
        return pulumi.get(self, "setup_logs_rules")

    @property
    @pulumi.getter(name="setupMetricsRules")
    def setup_metrics_rules(self) -> Optional[bool]:
        """
        Setup metrics rules.
        """
        return pulumi.get(self, "setup_metrics_rules")

    @property
    @pulumi.getter(name="writeLogs")
    def write_logs(self) -> Optional[bool]:
        """
        Write logs.
        """
        return pulumi.get(self, "write_logs")

    @property
    @pulumi.getter(name="writeMetrics")
    def write_metrics(self) -> Optional[bool]:
        """
        Write metrics.
        """
        return pulumi.get(self, "write_metrics")

    @property
    @pulumi.getter(name="writeTraces")
    def write_traces(self) -> Optional[bool]:
        """
        Write traces.
        """
        return pulumi.get(self, "write_traces")


@pulumi.output_type
class GetCockpitEndpointResult(dict):
    def __init__(__self__, *,
                 alertmanager_url: str,
                 grafana_url: str,
                 logs_url: str,
                 metrics_url: str,
                 traces_url: str):
        """
        :param str alertmanager_url: The alertmanager URL
        :param str grafana_url: The grafana URL
        :param str logs_url: The logs URL
        :param str metrics_url: The metrics URL
        """
        pulumi.set(__self__, "alertmanager_url", alertmanager_url)
        pulumi.set(__self__, "grafana_url", grafana_url)
        pulumi.set(__self__, "logs_url", logs_url)
        pulumi.set(__self__, "metrics_url", metrics_url)
        pulumi.set(__self__, "traces_url", traces_url)

    @property
    @pulumi.getter(name="alertmanagerUrl")
    def alertmanager_url(self) -> str:
        """
        The alertmanager URL
        """
        return pulumi.get(self, "alertmanager_url")

    @property
    @pulumi.getter(name="grafanaUrl")
    def grafana_url(self) -> str:
        """
        The grafana URL
        """
        return pulumi.get(self, "grafana_url")

    @property
    @pulumi.getter(name="logsUrl")
    def logs_url(self) -> str:
        """
        The logs URL
        """
        return pulumi.get(self, "logs_url")

    @property
    @pulumi.getter(name="metricsUrl")
    def metrics_url(self) -> str:
        """
        The metrics URL
        """
        return pulumi.get(self, "metrics_url")

    @property
    @pulumi.getter(name="tracesUrl")
    def traces_url(self) -> str:
        return pulumi.get(self, "traces_url")


