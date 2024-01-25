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
    'ClusterAutoUpgrade',
    'ClusterAutoscalerConfig',
    'ClusterKubeconfig',
    'ClusterOpenIdConnectConfig',
    'PoolNode',
    'PoolUpgradePolicy',
    'GetClusterAutoUpgradeResult',
    'GetClusterAutoscalerConfigResult',
    'GetClusterKubeconfigResult',
    'GetClusterOpenIdConnectConfigResult',
    'GetPoolNodeResult',
    'GetPoolUpgradePolicyResult',
]

@pulumi.output_type
class ClusterAutoUpgrade(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maintenanceWindowDay":
            suggest = "maintenance_window_day"
        elif key == "maintenanceWindowStartHour":
            suggest = "maintenance_window_start_hour"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterAutoUpgrade. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterAutoUpgrade.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterAutoUpgrade.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enable: bool,
                 maintenance_window_day: str,
                 maintenance_window_start_hour: int):
        """
        :param bool enable: Set to `true` to enable Kubernetes patch version auto upgrades.
               > **Important:** When enabling auto upgrades, the `version` field take a minor version like x.y (ie 1.18).
        :param str maintenance_window_day: The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
        :param int maintenance_window_start_hour: The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
        """
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "maintenance_window_day", maintenance_window_day)
        pulumi.set(__self__, "maintenance_window_start_hour", maintenance_window_start_hour)

    @property
    @pulumi.getter
    def enable(self) -> bool:
        """
        Set to `true` to enable Kubernetes patch version auto upgrades.
        > **Important:** When enabling auto upgrades, the `version` field take a minor version like x.y (ie 1.18).
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="maintenanceWindowDay")
    def maintenance_window_day(self) -> str:
        """
        The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
        """
        return pulumi.get(self, "maintenance_window_day")

    @property
    @pulumi.getter(name="maintenanceWindowStartHour")
    def maintenance_window_start_hour(self) -> int:
        """
        The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
        """
        return pulumi.get(self, "maintenance_window_start_hour")


@pulumi.output_type
class ClusterAutoscalerConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "balanceSimilarNodeGroups":
            suggest = "balance_similar_node_groups"
        elif key == "disableScaleDown":
            suggest = "disable_scale_down"
        elif key == "expendablePodsPriorityCutoff":
            suggest = "expendable_pods_priority_cutoff"
        elif key == "ignoreDaemonsetsUtilization":
            suggest = "ignore_daemonsets_utilization"
        elif key == "maxGracefulTerminationSec":
            suggest = "max_graceful_termination_sec"
        elif key == "scaleDownDelayAfterAdd":
            suggest = "scale_down_delay_after_add"
        elif key == "scaleDownUnneededTime":
            suggest = "scale_down_unneeded_time"
        elif key == "scaleDownUtilizationThreshold":
            suggest = "scale_down_utilization_threshold"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterAutoscalerConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterAutoscalerConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterAutoscalerConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 balance_similar_node_groups: Optional[bool] = None,
                 disable_scale_down: Optional[bool] = None,
                 estimator: Optional[str] = None,
                 expander: Optional[str] = None,
                 expendable_pods_priority_cutoff: Optional[int] = None,
                 ignore_daemonsets_utilization: Optional[bool] = None,
                 max_graceful_termination_sec: Optional[int] = None,
                 scale_down_delay_after_add: Optional[str] = None,
                 scale_down_unneeded_time: Optional[str] = None,
                 scale_down_utilization_threshold: Optional[float] = None):
        """
        :param bool balance_similar_node_groups: Detect similar node groups and balance the number of nodes between them.
        :param bool disable_scale_down: Disables the scale down feature of the autoscaler.
        :param str estimator: Type of resource estimator to be used in scale up.
        :param str expander: Type of node group expander to be used in scale up.
        :param int expendable_pods_priority_cutoff: Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        :param bool ignore_daemonsets_utilization: Ignore DaemonSet pods when calculating resource utilization for scaling down.
        :param int max_graceful_termination_sec: Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
        :param str scale_down_delay_after_add: How long after scale up that scale down evaluation resumes.
        :param str scale_down_unneeded_time: How long a node should be unneeded before it is eligible for scale down.
        :param float scale_down_utilization_threshold: Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
        """
        if balance_similar_node_groups is not None:
            pulumi.set(__self__, "balance_similar_node_groups", balance_similar_node_groups)
        if disable_scale_down is not None:
            pulumi.set(__self__, "disable_scale_down", disable_scale_down)
        if estimator is not None:
            pulumi.set(__self__, "estimator", estimator)
        if expander is not None:
            pulumi.set(__self__, "expander", expander)
        if expendable_pods_priority_cutoff is not None:
            pulumi.set(__self__, "expendable_pods_priority_cutoff", expendable_pods_priority_cutoff)
        if ignore_daemonsets_utilization is not None:
            pulumi.set(__self__, "ignore_daemonsets_utilization", ignore_daemonsets_utilization)
        if max_graceful_termination_sec is not None:
            pulumi.set(__self__, "max_graceful_termination_sec", max_graceful_termination_sec)
        if scale_down_delay_after_add is not None:
            pulumi.set(__self__, "scale_down_delay_after_add", scale_down_delay_after_add)
        if scale_down_unneeded_time is not None:
            pulumi.set(__self__, "scale_down_unneeded_time", scale_down_unneeded_time)
        if scale_down_utilization_threshold is not None:
            pulumi.set(__self__, "scale_down_utilization_threshold", scale_down_utilization_threshold)

    @property
    @pulumi.getter(name="balanceSimilarNodeGroups")
    def balance_similar_node_groups(self) -> Optional[bool]:
        """
        Detect similar node groups and balance the number of nodes between them.
        """
        return pulumi.get(self, "balance_similar_node_groups")

    @property
    @pulumi.getter(name="disableScaleDown")
    def disable_scale_down(self) -> Optional[bool]:
        """
        Disables the scale down feature of the autoscaler.
        """
        return pulumi.get(self, "disable_scale_down")

    @property
    @pulumi.getter
    def estimator(self) -> Optional[str]:
        """
        Type of resource estimator to be used in scale up.
        """
        return pulumi.get(self, "estimator")

    @property
    @pulumi.getter
    def expander(self) -> Optional[str]:
        """
        Type of node group expander to be used in scale up.
        """
        return pulumi.get(self, "expander")

    @property
    @pulumi.getter(name="expendablePodsPriorityCutoff")
    def expendable_pods_priority_cutoff(self) -> Optional[int]:
        """
        Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        """
        return pulumi.get(self, "expendable_pods_priority_cutoff")

    @property
    @pulumi.getter(name="ignoreDaemonsetsUtilization")
    def ignore_daemonsets_utilization(self) -> Optional[bool]:
        """
        Ignore DaemonSet pods when calculating resource utilization for scaling down.
        """
        return pulumi.get(self, "ignore_daemonsets_utilization")

    @property
    @pulumi.getter(name="maxGracefulTerminationSec")
    def max_graceful_termination_sec(self) -> Optional[int]:
        """
        Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
        """
        return pulumi.get(self, "max_graceful_termination_sec")

    @property
    @pulumi.getter(name="scaleDownDelayAfterAdd")
    def scale_down_delay_after_add(self) -> Optional[str]:
        """
        How long after scale up that scale down evaluation resumes.
        """
        return pulumi.get(self, "scale_down_delay_after_add")

    @property
    @pulumi.getter(name="scaleDownUnneededTime")
    def scale_down_unneeded_time(self) -> Optional[str]:
        """
        How long a node should be unneeded before it is eligible for scale down.
        """
        return pulumi.get(self, "scale_down_unneeded_time")

    @property
    @pulumi.getter(name="scaleDownUtilizationThreshold")
    def scale_down_utilization_threshold(self) -> Optional[float]:
        """
        Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
        """
        return pulumi.get(self, "scale_down_utilization_threshold")


@pulumi.output_type
class ClusterKubeconfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterCaCertificate":
            suggest = "cluster_ca_certificate"
        elif key == "configFile":
            suggest = "config_file"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterKubeconfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterKubeconfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterKubeconfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_ca_certificate: Optional[str] = None,
                 config_file: Optional[str] = None,
                 host: Optional[str] = None,
                 token: Optional[str] = None):
        """
        :param str cluster_ca_certificate: The CA certificate of the Kubernetes API server.
        :param str config_file: The raw kubeconfig file.
        :param str host: The URL of the Kubernetes API server.
        :param str token: The token to connect to the Kubernetes API server.
        """
        if cluster_ca_certificate is not None:
            pulumi.set(__self__, "cluster_ca_certificate", cluster_ca_certificate)
        if config_file is not None:
            pulumi.set(__self__, "config_file", config_file)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> Optional[str]:
        """
        The CA certificate of the Kubernetes API server.
        """
        return pulumi.get(self, "cluster_ca_certificate")

    @property
    @pulumi.getter(name="configFile")
    def config_file(self) -> Optional[str]:
        """
        The raw kubeconfig file.
        """
        return pulumi.get(self, "config_file")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        The URL of the Kubernetes API server.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        The token to connect to the Kubernetes API server.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class ClusterOpenIdConnectConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "issuerUrl":
            suggest = "issuer_url"
        elif key == "groupsClaims":
            suggest = "groups_claims"
        elif key == "groupsPrefix":
            suggest = "groups_prefix"
        elif key == "requiredClaims":
            suggest = "required_claims"
        elif key == "usernameClaim":
            suggest = "username_claim"
        elif key == "usernamePrefix":
            suggest = "username_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterOpenIdConnectConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterOpenIdConnectConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterOpenIdConnectConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 issuer_url: str,
                 groups_claims: Optional[Sequence[str]] = None,
                 groups_prefix: Optional[str] = None,
                 required_claims: Optional[Sequence[str]] = None,
                 username_claim: Optional[str] = None,
                 username_prefix: Optional[str] = None):
        """
        :param str client_id: A client id that all tokens must be issued for
        :param str issuer_url: URL of the provider which allows the API server to discover public signing keys
        :param Sequence[str] groups_claims: JWT claim to use as the user's group
        :param str groups_prefix: Prefix prepended to group claims
        :param Sequence[str] required_claims: Multiple key=value pairs that describes a required claim in the ID Token
        :param str username_claim: JWT claim to use as the user name
        :param str username_prefix: Prefix prepended to username
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "issuer_url", issuer_url)
        if groups_claims is not None:
            pulumi.set(__self__, "groups_claims", groups_claims)
        if groups_prefix is not None:
            pulumi.set(__self__, "groups_prefix", groups_prefix)
        if required_claims is not None:
            pulumi.set(__self__, "required_claims", required_claims)
        if username_claim is not None:
            pulumi.set(__self__, "username_claim", username_claim)
        if username_prefix is not None:
            pulumi.set(__self__, "username_prefix", username_prefix)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        A client id that all tokens must be issued for
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> str:
        """
        URL of the provider which allows the API server to discover public signing keys
        """
        return pulumi.get(self, "issuer_url")

    @property
    @pulumi.getter(name="groupsClaims")
    def groups_claims(self) -> Optional[Sequence[str]]:
        """
        JWT claim to use as the user's group
        """
        return pulumi.get(self, "groups_claims")

    @property
    @pulumi.getter(name="groupsPrefix")
    def groups_prefix(self) -> Optional[str]:
        """
        Prefix prepended to group claims
        """
        return pulumi.get(self, "groups_prefix")

    @property
    @pulumi.getter(name="requiredClaims")
    def required_claims(self) -> Optional[Sequence[str]]:
        """
        Multiple key=value pairs that describes a required claim in the ID Token
        """
        return pulumi.get(self, "required_claims")

    @property
    @pulumi.getter(name="usernameClaim")
    def username_claim(self) -> Optional[str]:
        """
        JWT claim to use as the user name
        """
        return pulumi.get(self, "username_claim")

    @property
    @pulumi.getter(name="usernamePrefix")
    def username_prefix(self) -> Optional[str]:
        """
        Prefix prepended to username
        """
        return pulumi.get(self, "username_prefix")


@pulumi.output_type
class PoolNode(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "publicIp":
            suggest = "public_ip"
        elif key == "publicIpV6":
            suggest = "public_ip_v6"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PoolNode. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PoolNode.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PoolNode.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: Optional[str] = None,
                 public_ip: Optional[str] = None,
                 public_ip_v6: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str name: The name for the pool.
               > **Important:** Updates to this field will recreate a new resource.
        :param str public_ip: The public IPv4.
        :param str public_ip_v6: The public IPv6.
        :param str status: The status of the node.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_ip is not None:
            pulumi.set(__self__, "public_ip", public_ip)
        if public_ip_v6 is not None:
            pulumi.set(__self__, "public_ip_v6", public_ip_v6)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name for the pool.
        > **Important:** Updates to this field will recreate a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> Optional[str]:
        """
        The public IPv4.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIpV6")
    def public_ip_v6(self) -> Optional[str]:
        """
        The public IPv6.
        """
        return pulumi.get(self, "public_ip_v6")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the node.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class PoolUpgradePolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxSurge":
            suggest = "max_surge"
        elif key == "maxUnavailable":
            suggest = "max_unavailable"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PoolUpgradePolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PoolUpgradePolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PoolUpgradePolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_surge: Optional[int] = None,
                 max_unavailable: Optional[int] = None):
        """
        :param int max_surge: The maximum number of nodes to be created during the upgrade
        :param int max_unavailable: The maximum number of nodes that can be not ready at the same time
        """
        if max_surge is not None:
            pulumi.set(__self__, "max_surge", max_surge)
        if max_unavailable is not None:
            pulumi.set(__self__, "max_unavailable", max_unavailable)

    @property
    @pulumi.getter(name="maxSurge")
    def max_surge(self) -> Optional[int]:
        """
        The maximum number of nodes to be created during the upgrade
        """
        return pulumi.get(self, "max_surge")

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> Optional[int]:
        """
        The maximum number of nodes that can be not ready at the same time
        """
        return pulumi.get(self, "max_unavailable")


@pulumi.output_type
class GetClusterAutoUpgradeResult(dict):
    def __init__(__self__, *,
                 enable: bool,
                 maintenance_window_day: str,
                 maintenance_window_start_hour: int):
        """
        :param bool enable: True if Kubernetes patch version auto upgrades is enabled.
        :param str maintenance_window_day: The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
        :param int maintenance_window_start_hour: The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
        """
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "maintenance_window_day", maintenance_window_day)
        pulumi.set(__self__, "maintenance_window_start_hour", maintenance_window_start_hour)

    @property
    @pulumi.getter
    def enable(self) -> bool:
        """
        True if Kubernetes patch version auto upgrades is enabled.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="maintenanceWindowDay")
    def maintenance_window_day(self) -> str:
        """
        The day of the auto upgrade maintenance window (`monday` to `sunday`, or `any`).
        """
        return pulumi.get(self, "maintenance_window_day")

    @property
    @pulumi.getter(name="maintenanceWindowStartHour")
    def maintenance_window_start_hour(self) -> int:
        """
        The start hour (UTC) of the 2-hour auto upgrade maintenance window (0 to 23).
        """
        return pulumi.get(self, "maintenance_window_start_hour")


@pulumi.output_type
class GetClusterAutoscalerConfigResult(dict):
    def __init__(__self__, *,
                 balance_similar_node_groups: bool,
                 disable_scale_down: bool,
                 estimator: str,
                 expander: str,
                 expendable_pods_priority_cutoff: int,
                 ignore_daemonsets_utilization: bool,
                 max_graceful_termination_sec: int,
                 scale_down_delay_after_add: str,
                 scale_down_unneeded_time: str,
                 scale_down_utilization_threshold: float):
        """
        :param bool balance_similar_node_groups: True if detecting similar node groups and balance the number of nodes between them is enabled.
        :param bool disable_scale_down: True if the scale down feature of the autoscaler is disabled.
        :param str estimator: The type of resource estimator used in scale up.
        :param str expander: The type of node group expander be used in scale up.
        :param int expendable_pods_priority_cutoff: Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        :param bool ignore_daemonsets_utilization: True if ignoring DaemonSet pods when calculating resource utilization for scaling down is enabled.
        :param str scale_down_delay_after_add: The duration after scale up that scale down evaluation resumes.
        :param str scale_down_unneeded_time: The duration a node should be unneeded before it is eligible for scale down.
        """
        pulumi.set(__self__, "balance_similar_node_groups", balance_similar_node_groups)
        pulumi.set(__self__, "disable_scale_down", disable_scale_down)
        pulumi.set(__self__, "estimator", estimator)
        pulumi.set(__self__, "expander", expander)
        pulumi.set(__self__, "expendable_pods_priority_cutoff", expendable_pods_priority_cutoff)
        pulumi.set(__self__, "ignore_daemonsets_utilization", ignore_daemonsets_utilization)
        pulumi.set(__self__, "max_graceful_termination_sec", max_graceful_termination_sec)
        pulumi.set(__self__, "scale_down_delay_after_add", scale_down_delay_after_add)
        pulumi.set(__self__, "scale_down_unneeded_time", scale_down_unneeded_time)
        pulumi.set(__self__, "scale_down_utilization_threshold", scale_down_utilization_threshold)

    @property
    @pulumi.getter(name="balanceSimilarNodeGroups")
    def balance_similar_node_groups(self) -> bool:
        """
        True if detecting similar node groups and balance the number of nodes between them is enabled.
        """
        return pulumi.get(self, "balance_similar_node_groups")

    @property
    @pulumi.getter(name="disableScaleDown")
    def disable_scale_down(self) -> bool:
        """
        True if the scale down feature of the autoscaler is disabled.
        """
        return pulumi.get(self, "disable_scale_down")

    @property
    @pulumi.getter
    def estimator(self) -> str:
        """
        The type of resource estimator used in scale up.
        """
        return pulumi.get(self, "estimator")

    @property
    @pulumi.getter
    def expander(self) -> str:
        """
        The type of node group expander be used in scale up.
        """
        return pulumi.get(self, "expander")

    @property
    @pulumi.getter(name="expendablePodsPriorityCutoff")
    def expendable_pods_priority_cutoff(self) -> int:
        """
        Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        """
        return pulumi.get(self, "expendable_pods_priority_cutoff")

    @property
    @pulumi.getter(name="ignoreDaemonsetsUtilization")
    def ignore_daemonsets_utilization(self) -> bool:
        """
        True if ignoring DaemonSet pods when calculating resource utilization for scaling down is enabled.
        """
        return pulumi.get(self, "ignore_daemonsets_utilization")

    @property
    @pulumi.getter(name="maxGracefulTerminationSec")
    def max_graceful_termination_sec(self) -> int:
        return pulumi.get(self, "max_graceful_termination_sec")

    @property
    @pulumi.getter(name="scaleDownDelayAfterAdd")
    def scale_down_delay_after_add(self) -> str:
        """
        The duration after scale up that scale down evaluation resumes.
        """
        return pulumi.get(self, "scale_down_delay_after_add")

    @property
    @pulumi.getter(name="scaleDownUnneededTime")
    def scale_down_unneeded_time(self) -> str:
        """
        The duration a node should be unneeded before it is eligible for scale down.
        """
        return pulumi.get(self, "scale_down_unneeded_time")

    @property
    @pulumi.getter(name="scaleDownUtilizationThreshold")
    def scale_down_utilization_threshold(self) -> float:
        return pulumi.get(self, "scale_down_utilization_threshold")


@pulumi.output_type
class GetClusterKubeconfigResult(dict):
    def __init__(__self__, *,
                 cluster_ca_certificate: str,
                 config_file: str,
                 host: str,
                 token: str):
        """
        :param str cluster_ca_certificate: The CA certificate of the Kubernetes API server.
        :param str config_file: The raw kubeconfig file.
        :param str host: The URL of the Kubernetes API server.
        :param str token: The token to connect to the Kubernetes API server.
        """
        pulumi.set(__self__, "cluster_ca_certificate", cluster_ca_certificate)
        pulumi.set(__self__, "config_file", config_file)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> str:
        """
        The CA certificate of the Kubernetes API server.
        """
        return pulumi.get(self, "cluster_ca_certificate")

    @property
    @pulumi.getter(name="configFile")
    def config_file(self) -> str:
        """
        The raw kubeconfig file.
        """
        return pulumi.get(self, "config_file")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The URL of the Kubernetes API server.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The token to connect to the Kubernetes API server.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class GetClusterOpenIdConnectConfigResult(dict):
    def __init__(__self__, *,
                 client_id: str,
                 groups_claims: Sequence[str],
                 groups_prefix: str,
                 issuer_url: str,
                 required_claims: Sequence[str],
                 username_claim: str,
                 username_prefix: str):
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "groups_claims", groups_claims)
        pulumi.set(__self__, "groups_prefix", groups_prefix)
        pulumi.set(__self__, "issuer_url", issuer_url)
        pulumi.set(__self__, "required_claims", required_claims)
        pulumi.set(__self__, "username_claim", username_claim)
        pulumi.set(__self__, "username_prefix", username_prefix)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="groupsClaims")
    def groups_claims(self) -> Sequence[str]:
        return pulumi.get(self, "groups_claims")

    @property
    @pulumi.getter(name="groupsPrefix")
    def groups_prefix(self) -> str:
        return pulumi.get(self, "groups_prefix")

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> str:
        return pulumi.get(self, "issuer_url")

    @property
    @pulumi.getter(name="requiredClaims")
    def required_claims(self) -> Sequence[str]:
        return pulumi.get(self, "required_claims")

    @property
    @pulumi.getter(name="usernameClaim")
    def username_claim(self) -> str:
        return pulumi.get(self, "username_claim")

    @property
    @pulumi.getter(name="usernamePrefix")
    def username_prefix(self) -> str:
        return pulumi.get(self, "username_prefix")


@pulumi.output_type
class GetPoolNodeResult(dict):
    def __init__(__self__, *,
                 name: str,
                 public_ip: str,
                 public_ip_v6: str,
                 status: str):
        """
        :param str name: The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
        :param str public_ip: The public IPv4.
        :param str public_ip_v6: The public IPv6.
        :param str status: The status of the node.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "public_ip", public_ip)
        pulumi.set(__self__, "public_ip_v6", public_ip_v6)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        The public IPv4.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIpV6")
    def public_ip_v6(self) -> str:
        """
        The public IPv6.
        """
        return pulumi.get(self, "public_ip_v6")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the node.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetPoolUpgradePolicyResult(dict):
    def __init__(__self__, *,
                 max_surge: int,
                 max_unavailable: int):
        pulumi.set(__self__, "max_surge", max_surge)
        pulumi.set(__self__, "max_unavailable", max_unavailable)

    @property
    @pulumi.getter(name="maxSurge")
    def max_surge(self) -> int:
        return pulumi.get(self, "max_surge")

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> int:
        return pulumi.get(self, "max_unavailable")

