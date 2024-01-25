// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scaleway

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	// Replace this provider with the provider you are bridging.
	scaleway "github.com/scaleway/terraform-provider-scaleway/v2/scaleway"

	"github.com/raeumlich/pulumi-scaleway/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "scaleway"
	// modules:
	mainMod              = "index" // the scaleway module
	accountMod           = "account"
	appleSiliconMod      = "applesilicon"
	billingMod           = "billing"
	blockStorageMod      = "blockstorage"
	cockpitMod           = "cockpit"
	containerRegistryMod = "containerregistry"
	dnsMod               = "dns"
	documentDBMod        = "documentdb"
	elasticMetalMod      = "elasticmetal"
	iamMod               = "iam"
	instanceMod          = "instance"
	iotMod               = "iot"
	ipamMod              = "ipam"
	kubernetesMod        = "kubernetes"
	loadBalancerMod      = "loadbalancer"
	mnqMod               = "mnq"
	objectStorageMod     = "objectstorage"
	rdbMod               = "rdb"
	redisMod             = "redis"
	secretMod            = "secret"
	serverlessMod        = "serverless"
	temMod               = "tem"
	vpcMod               = "vpc"
	webhostingMod        = "webhosting"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-scaleway/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           shimv2.NewProvider(scaleway.Provider(scaleway.DefaultProviderConfig())()),
		Name:        "scaleway",
		DisplayName: "Scaleway",
		Publisher:   "RÄUMLICH plus GmbH",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL:                 "",
		PluginDownloadURL:       "https://github.com/raeumlich/pulumi-scaleway/releases/",
		Description:             "A Pulumi package for creating and managing scaleway cloud resources.",
		Keywords:                []string{"pulumi", "scaleway", "category/cloud"},
		License:                 "Apache-2.0",
		Homepage:                "https://www.pulumi.com",
		Repository:              "https://github.com/raeumlich/pulumi-scaleway",
		GitHubOrg:               "scaleway",
		TFProviderModuleVersion: "v2",
		TFProviderLicense:       tfbridge.SetProviderLicense(tfbridge.MPL20LicenseType),
		MetadataInfo:            tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_ACCESS_KEY"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_SECRET_KEY"},
				},
			},
			"project_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_PROJECT_ID"},
				},
			},
			"organization_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ORGANIZATION_ID"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_REGION"},
				},
			},
			"zone": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"SCW_DEFAULT_ZONE"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },

			// account module
			"scaleway_account_project": {
				Tok:  tfbridge.MakeResource(mainPkg, accountMod, "Project"),
				Docs: &tfbridge.DocInfo{Source: "account_project.md"},
			},
			"scaleway_account_ssh_key": {
				Tok:  tfbridge.MakeResource(mainPkg, accountMod, "SSHKey"),
				Docs: &tfbridge.DocInfo{Source: "account_ssh_key.md"},
			},

			// applesilicon module
			"scaleway_apple_silicon_server": {
				Tok:  tfbridge.MakeResource(mainPkg, appleSiliconMod, "Server"),
				Docs: &tfbridge.DocInfo{Source: "apple_silicon.md"},
			},

			// blockstorage module
			"scaleway_block_snapshot": {
				Tok:  tfbridge.MakeResource(mainPkg, blockStorageMod, "Snapshot"),
				Docs: &tfbridge.DocInfo{Source: "block_snapshot.md"},
			},
			"scaleway_block_volume": {
				Tok:  tfbridge.MakeResource(mainPkg, blockStorageMod, "Volume"),
				Docs: &tfbridge.DocInfo{Source: "block_volume.md"},
			},

			// cockpit module
			"scaleway_cockpit": {
				Tok:  tfbridge.MakeResource(mainPkg, cockpitMod, "Cockpit"),
				Docs: &tfbridge.DocInfo{Source: "cockpit.md"},
			},
			"scaleway_cockpit_grafana_user": {
				Tok:  tfbridge.MakeResource(mainPkg, cockpitMod, "GrafanaUser"),
				Docs: &tfbridge.DocInfo{Source: "cockpit_grafana_user.md"},
			},
			"scaleway_cockpit_token": {
				Tok:  tfbridge.MakeResource(mainPkg, cockpitMod, "Token"),
				Docs: &tfbridge.DocInfo{Source: "cockpit_token.md"},
			},

			// containerregistry module
			"scaleway_registry_namespace": {
				Tok:  tfbridge.MakeResource(mainPkg, containerRegistryMod, "Namespace"),
				Docs: &tfbridge.DocInfo{Source: "registry_namespace.md"},
			},

			// dns module
			"scaleway_domain_record": {
				Tok:  tfbridge.MakeResource(mainPkg, dnsMod, "Record"),
				Docs: &tfbridge.DocInfo{Source: "domain_record.md"},
			},
			"scaleway_domain_zone": {
				Tok:  tfbridge.MakeResource(mainPkg, dnsMod, "Zone"),
				Docs: &tfbridge.DocInfo{Source: "domain_zone.md"},
			},

			// documentdb module
			"scaleway_documentdb_database": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "Database"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_database.md"},
			},
			"scaleway_documentdb_instance": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "Instance"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_instance.md"},
			},
			"scaleway_documentdb_private_network_endpoint": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "PrivateNetworkEndpoint"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_private_network_endpoint.md"},
			},
			"scaleway_documentdb_privilege": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "Privilege"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_privilege.md"},
			},
			"scaleway_documentdb_read_replica": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "ReadReplica"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_read_replica.md"},
			},
			"scaleway_documentdb_user": {
				Tok:  tfbridge.MakeResource(mainPkg, documentDBMod, "User"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_user.md"},
			},

			// elasticmetal module
			"scaleway_baremetal_server": {
				Tok:  tfbridge.MakeResource(mainPkg, elasticMetalMod, "BareMetalServer"),
				Docs: &tfbridge.DocInfo{Source: "baremetal_server.md"},
			},
			"scaleway_flexible_ip": {
				Tok:  tfbridge.MakeResource(mainPkg, elasticMetalMod, "FlexibleIP"),
				Docs: &tfbridge.DocInfo{Source: "flexible_ip.md"},
			},
			"scaleway_flexible_ip_mac_address": {
				Tok:  tfbridge.MakeResource(mainPkg, elasticMetalMod, "FlexibleIPMACAddress"),
				Docs: &tfbridge.DocInfo{Source: "flexible_ip_mac_address.md"},
			},

			// iam module
			"scaleway_iam_api_key": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "APIKey"),
				Docs: &tfbridge.DocInfo{Source: "iam_api_key.md"},
			},
			"scaleway_iam_application": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "Application"),
				Docs: &tfbridge.DocInfo{Source: "iam_application.md"},
			},
			"scaleway_iam_group": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "Group"),
				Docs: &tfbridge.DocInfo{Source: "iam_group.md"},
			},
			"scaleway_iam_group_membership": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "GroupMembership"),
				Docs: &tfbridge.DocInfo{Source: "iam_group_membership.md"},
			},
			"scaleway_iam_policy": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "Policy"),
				Docs: &tfbridge.DocInfo{Source: "iam_policy.md"},
			},
			"scaleway_iam_ssh_key": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "SSHKey"),
				Docs: &tfbridge.DocInfo{Source: "iam_ssh_key.md"},
			},
			"scaleway_iam_user": {
				Tok:  tfbridge.MakeResource(mainPkg, iamMod, "User"),
				Docs: &tfbridge.DocInfo{Source: "iam_user.md"},
			},

			// instance module
			"scaleway_instance_image": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "Image"),
				Docs: &tfbridge.DocInfo{Source: "instance_image.md"},
			},
			"scaleway_instance_ip": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "IP"),
				Docs: &tfbridge.DocInfo{Source: "instance_ip.md"},
			},
			"scaleway_instance_ip_reverse_dns": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "IPReverseDNS"),
				Docs: &tfbridge.DocInfo{Source: "instance_ip_reverse_dns.md"},
			},
			"scaleway_instance_placement_group": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "PlacementGroup"),
				Docs: &tfbridge.DocInfo{Source: "instance_placement_group.md"},
			},
			"scaleway_instance_private_nic": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "PrivateNIC"),
				Docs: &tfbridge.DocInfo{Source: "instance_private_nic.md"},
			},
			"scaleway_instance_security_group": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "SecurityGroup"),
				Docs: &tfbridge.DocInfo{Source: "instance_security_group.md"},
			},
			"scaleway_instance_security_group_rules": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "SecurityGroupRules"),
				Docs: &tfbridge.DocInfo{Source: "instance_security_group_rules.md"},
			},
			"scaleway_instance_server": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "Server"),
				Docs: &tfbridge.DocInfo{Source: "instance_server.md"},
			},
			"scaleway_instance_snapshot": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "Snapshot"),
				Docs: &tfbridge.DocInfo{Source: "instance_snapshot.md"},
			},
			"scaleway_instance_user_data": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "UserData"),
				Docs: &tfbridge.DocInfo{Source: "instance_user_data.md"},
			},
			"scaleway_instance_volume": {
				Tok:  tfbridge.MakeResource(mainPkg, instanceMod, "Volume"),
				Docs: &tfbridge.DocInfo{Source: "instance_volume.md"},
			},

			// iot module
			"scaleway_iot_device": {
				Tok:  tfbridge.MakeResource(mainPkg, iotMod, "Device"),
				Docs: &tfbridge.DocInfo{Source: "iot_device.md"},
			},
			"scaleway_iot_hub": {
				Tok:  tfbridge.MakeResource(mainPkg, iotMod, "Hub"),
				Docs: &tfbridge.DocInfo{Source: "iot_hub.md"},
			},
			"scaleway_iot_network": {
				Tok:  tfbridge.MakeResource(mainPkg, iotMod, "Network"),
				Docs: &tfbridge.DocInfo{Source: "iot_network.md"},
			},
			"scaleway_iot_route": {
				Tok:  tfbridge.MakeResource(mainPkg, iotMod, "Route"),
				Docs: &tfbridge.DocInfo{Source: "iot_route.md"},
			},

			// ipam module
			"scaleway_ipam_ip": {
				Tok:  tfbridge.MakeResource(mainPkg, ipamMod, "IP"),
				Docs: &tfbridge.DocInfo{Source: "ipam_ip.md"},
			},

			// kubernetes module
			"scaleway_k8s_cluster": {
				Tok:  tfbridge.MakeResource(mainPkg, kubernetesMod, "Cluster"),
				Docs: &tfbridge.DocInfo{Source: "k8s_cluster.md"},
			},
			"scaleway_k8s_pool": {
				Tok:  tfbridge.MakeResource(mainPkg, kubernetesMod, "Pool"),
				Docs: &tfbridge.DocInfo{Source: "k8s_pool.md"},
			},

			// loadbalancer module
			"scaleway_lb": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "LoadBalancer"),
				Docs: &tfbridge.DocInfo{Source: "lb.md"},
			},
			"scaleway_lb_acl": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "ACL"),
				Docs: &tfbridge.DocInfo{Source: "lb_acl.md"},
			},
			"scaleway_lb_backend": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "Backend"),
				Docs: &tfbridge.DocInfo{Source: "lb_backend.md"},
			},
			"scaleway_lb_certificate": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "Certficate"),
				Docs: &tfbridge.DocInfo{Source: "lb_certificate.md"},
			},
			"scaleway_lb_frontend": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "Frontend"),
				Docs: &tfbridge.DocInfo{Source: "lb_frontend.md"},
			},
			"scaleway_lb_ip": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "IP"),
				Docs: &tfbridge.DocInfo{Source: "lb_ip.md"},
			},
			"scaleway_lb_route": {
				Tok:  tfbridge.MakeResource(mainPkg, loadBalancerMod, "Route"),
				Docs: &tfbridge.DocInfo{Source: "lb_route.md"},
			},

			// mnq module
			"scaleway_mnq_nats_account": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "NATSAccount"),
				Docs: &tfbridge.DocInfo{Source: "mnq_nats_account.md"},
			},
			"scaleway_mnq_nats_credentials": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "NATSCredentials"),
				Docs: &tfbridge.DocInfo{Source: "mnq_nats_credentials.md"},
			},
			"scaleway_mnq_sns": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SNS"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sns.md"},
			},
			"scaleway_mnq_sns_credentials": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SNSCredentials"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sns_credentials.md"},
			},
			"scaleway_mnq_sns_topic": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SNSTopic"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sns_topic.md"},
			},
			"scaleway_mnq_sns_topic_subscription": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SNSTopicSubscription"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sns_topic_subscription.md"},
			},
			"scaleway_mnq_sqs": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SQS"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sqs.md"},
			},
			"scaleway_mnq_sqs_credentials": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SQSCredentials"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sqs_credentials.md"},
			},
			"scaleway_mnq_sqs_queue": {
				Tok:  tfbridge.MakeResource(mainPkg, mnqMod, "SQSQueue"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sqs_queue.md"},
			},

			// objectstorage module
			"scaleway_object": {
				// Cannot be named 'Object' due to TypeScript reserved word.
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "Item"),
				Docs: &tfbridge.DocInfo{Source: "object.md"},
			},
			"scaleway_object_bucket": {
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "Bucket"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket.md"},
			},
			"scaleway_object_bucket_acl": {
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "BucketACL"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket_acl.md"},
			},
			"scaleway_object_bucket_lock_configuration": {
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "BucketLockConfiguration"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket_lock_configuration.md"},
			},
			"scaleway_object_bucket_policy": {
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "BucketPolicy"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket_policy.md"},
			},
			"scaleway_object_bucket_website_configuration": {
				Tok:  tfbridge.MakeResource(mainPkg, objectStorageMod, "BucketWebsiteConfiguration"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket_website_configuration.md"},
			},

			// rdb module
			"scaleway_rdb_acl": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "ACL"),
				Docs: &tfbridge.DocInfo{Source: "rdb_acl.md"},
			},
			"scaleway_rdb_database": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "Database"),
				Docs: &tfbridge.DocInfo{Source: "rdb_database.md"},
			},
			"scaleway_rdb_database_backup": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "DatabaseBackup"),
				Docs: &tfbridge.DocInfo{Source: "rdb_database_backup.md"},
			},
			"scaleway_rdb_instance": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "Instance"),
				Docs: &tfbridge.DocInfo{Source: "rdb_instance.md"},
			},
			"scaleway_rdb_privilege": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "Privilege"),
				Docs: &tfbridge.DocInfo{Source: "rdb_privilege.md"},
			},
			"scaleway_rdb_read_replica": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "ReadReplica"),
				Docs: &tfbridge.DocInfo{Source: "rdb_read_replica.md"},
			},
			"scaleway_rdb_user": {
				Tok:  tfbridge.MakeResource(mainPkg, rdbMod, "User"),
				Docs: &tfbridge.DocInfo{Source: "rdb_user.md"},
			},

			// redis module
			"scaleway_redis_cluster": {
				Tok:  tfbridge.MakeResource(mainPkg, redisMod, "Cluster"),
				Docs: &tfbridge.DocInfo{Source: "redis_cluster.md"},
			},

			// secret module
			"scaleway_secret": {
				Tok:  tfbridge.MakeResource(mainPkg, secretMod, "Secret"),
				Docs: &tfbridge.DocInfo{Source: "secret.md"},
			},
			"scaleway_secret_version": {
				Tok:  tfbridge.MakeResource(mainPkg, secretMod, "Version"),
				Docs: &tfbridge.DocInfo{Source: "secret_version.md"},
			},

			// serverless module
			"scaleway_container": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "Container"),
				Docs: &tfbridge.DocInfo{Source: "container.md"},
			},
			"scaleway_container_cron": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "ContainerCron"),
				Docs: &tfbridge.DocInfo{Source: "container_cron.md"},
			},
			"scaleway_container_domain": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "ContainerDomain"),
				Docs: &tfbridge.DocInfo{Source: "container_domain.md"},
			},
			"scaleway_container_namespace": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "ContainerNamespace"),
				Docs: &tfbridge.DocInfo{Source: "container_namespace.md"},
			},
			"scaleway_container_token": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "ContainerToken"),
				Docs: &tfbridge.DocInfo{Source: "container_token.md"},
			},
			"scaleway_container_trigger": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "ContainerTrigger"),
				Docs: &tfbridge.DocInfo{Source: "container_trigger.md"},
			},
			"scaleway_function": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "Function"),
				Docs: &tfbridge.DocInfo{Source: "function.md"},
			},
			"scaleway_function_cron": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "FunctionCron"),
				Docs: &tfbridge.DocInfo{Source: "function_cron.md"},
			},
			"scaleway_function_domain": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "FunctionDomain"),
				Docs: &tfbridge.DocInfo{Source: "function_domain.md"},
			},
			"scaleway_function_namespace": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "FunctionNamespace"),
				Docs: &tfbridge.DocInfo{Source: "function_namespace.md"},
			},
			"scaleway_function_token": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "FunctionToken"),
				Docs: &tfbridge.DocInfo{Source: "function_token.md"},
			},
			"scaleway_function_trigger": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "FunctionTrigger"),
				Docs: &tfbridge.DocInfo{Source: "function_trigger.md"},
			},
			"scaleway_job_definition": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "JobDefinition"),
				Docs: &tfbridge.DocInfo{Source: "job_definition.md"},
			},
			"scaleway_sdb_sql_database": {
				Tok:  tfbridge.MakeResource(mainPkg, serverlessMod, "SDBDatabase"),
				Docs: &tfbridge.DocInfo{Source: "sdb_sql_database.md"},
			},

			// tem module
			"scaleway_tem_domain": {
				Tok:  tfbridge.MakeResource(mainPkg, temMod, "Domain"),
				Docs: &tfbridge.DocInfo{Source: "tem_domain.md"},
			},

			// vpc module
			"scaleway_vpc": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "VPC"),
				Docs: &tfbridge.DocInfo{Source: "vpc.md"},
			},
			"scaleway_vpc_gateway_network": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "GatewayNetwork"),
				Docs: &tfbridge.DocInfo{Source: "vpc_gateway_network.md"},
			},
			"scaleway_vpc_private_network": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PrivateNetwork"),
				Docs: &tfbridge.DocInfo{Source: "vpc_private_network.md"},
			},
			"scaleway_vpc_public_gateway": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGateway"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway.md"},
			},
			"scaleway_vpc_public_gateway_dhcp": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGatewayDHCP"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_dhcp.md"},
			},
			"scaleway_vpc_public_gateway_dhcp_reservation": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGatewayDHCPReservation"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_dhcp_reservation.md"},
			},
			"scaleway_vpc_public_gateway_ip": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGatewayIP"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_ip.md"},
			},
			"scaleway_vpc_public_gateway_ip_reverse_dns": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGatewayIPReverseDNS"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_ip_reverse_dns.md"},
			},
			"scaleway_vpc_public_gateway_pat_rule": {
				Tok:  tfbridge.MakeResource(mainPkg, vpcMod, "PublicGatewayPATRule"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_pat_rule.md"},
			},

			// webhosting module
			"scaleway_webhosting": {
				Tok:  tfbridge.MakeResource(mainPkg, webhostingMod, "WebHosting"),
				Docs: &tfbridge.DocInfo{Source: "webhosting.md"},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// account module
			"scaleway_account_project": {
				Tok:  tfbridge.MakeDataSource(mainPkg, accountMod, "getProject"),
				Docs: &tfbridge.DocInfo{Source: "account_project.md"},
			},
			"scaleway_account_ssh_key": {
				Tok:  tfbridge.MakeDataSource(mainPkg, accountMod, "getSSHKey"),
				Docs: &tfbridge.DocInfo{Source: "account_ssh_key.md"},
			},
			"scaleway_availability_zones": {
				Tok:  tfbridge.MakeDataSource(mainPkg, accountMod, "getAvailabilityZones"),
				Docs: &tfbridge.DocInfo{Source: "availability_zones.md"},
			},

			// billing module
			"scaleway_billing_consumptions": {
				Tok:  tfbridge.MakeDataSource(mainPkg, billingMod, "getConsumptions"),
				Docs: &tfbridge.DocInfo{Source: "billing_consumption.md"},
			},
			"scaleway_billing_invoices": {
				Tok:  tfbridge.MakeDataSource(mainPkg, billingMod, "getInvoices"),
				Docs: &tfbridge.DocInfo{Source: "billing_invoices.md"},
			},

			// blockstorage module
			"scaleway_block_snapshot": {
				Tok:  tfbridge.MakeDataSource(mainPkg, blockStorageMod, "getSnapshot"),
				Docs: &tfbridge.DocInfo{Source: "bloc_snapshot.md"}, // typo in the TF provider
			},
			"scaleway_block_volume": {
				Tok:  tfbridge.MakeDataSource(mainPkg, blockStorageMod, "getVolume"),
				Docs: &tfbridge.DocInfo{Source: "block_volume.md"},
			},

			// cockpit module
			"scaleway_cockpit": {
				Tok:  tfbridge.MakeDataSource(mainPkg, cockpitMod, "getCockpit"),
				Docs: &tfbridge.DocInfo{Source: "cockpit.md"},
			},
			"scaleway_cockpit_plan": {
				Tok:  tfbridge.MakeDataSource(mainPkg, cockpitMod, "getPlan"),
				Docs: &tfbridge.DocInfo{Source: "cockpit_plan.md"},
			},

			// containerregistry module
			"scaleway_registry_image": {
				Tok:  tfbridge.MakeDataSource(mainPkg, containerRegistryMod, "getImage"),
				Docs: &tfbridge.DocInfo{Source: "registry_image.md"},
			},
			"scaleway_registry_namespace": {
				Tok:  tfbridge.MakeDataSource(mainPkg, containerRegistryMod, "getNamespace"),
				Docs: &tfbridge.DocInfo{Source: "registry_namespace.md"},
			},

			// dns module
			"scaleway_domain_record": {
				Tok:  tfbridge.MakeDataSource(mainPkg, dnsMod, "getRecord"),
				Docs: &tfbridge.DocInfo{Source: "domain_record.md"},
			},
			"scaleway_domain_zone": {
				Tok:  tfbridge.MakeDataSource(mainPkg, dnsMod, "getZone"),
				Docs: &tfbridge.DocInfo{Source: "domain_zone.md"},
			},

			// documentdb module
			"scaleway_documentdb_database": {
				Tok:  tfbridge.MakeDataSource(mainPkg, documentDBMod, "getDatabase"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_database.md"},
			},
			"scaleway_documentdb_instance": {
				Tok:  tfbridge.MakeDataSource(mainPkg, documentDBMod, "getInstance"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_instance.md"},
			},
			"scaleway_documentdb_load_balancer_endpoint": {
				Tok:  tfbridge.MakeDataSource(mainPkg, documentDBMod, "getLoadBalancerEndpoint"),
				Docs: &tfbridge.DocInfo{Source: "documentdb_load_balancer_endpoint.md"},
			},

			// elasticmetal module
			"scaleway_baremetal_offer": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getBareMetalOffer"),
				Docs: &tfbridge.DocInfo{Source: "baremetal_offer.md"},
			},
			"scaleway_baremetal_option": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getBareMetalOption"),
				Docs: &tfbridge.DocInfo{Source: "baremetal_option.md"},
			},
			"scaleway_baremetal_os": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getBareMetalOS"),
				Docs: &tfbridge.DocInfo{Source: "baremetal_os.md"},
			},
			"scaleway_baremetal_server": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getBareMetalServer"),
				Docs: &tfbridge.DocInfo{Source: "baremetal_server.md"},
			},
			"scaleway_flexible_ip": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getFlexibleIP"),
				Docs: &tfbridge.DocInfo{Source: "flexible_ip.md"},
			},
			"scaleway_flexible_ips": {
				Tok:  tfbridge.MakeDataSource(mainPkg, elasticMetalMod, "getFlexibleIPs"),
				Docs: &tfbridge.DocInfo{Source: "flexible_ips.md"},
			},

			// iam module
			"scaleway_iam_application": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iamMod, "getApplication"),
				Docs: &tfbridge.DocInfo{Source: "iam_application.md"},
			},
			"scaleway_iam_group": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iamMod, "getGroup"),
				Docs: &tfbridge.DocInfo{Source: "iam_group.md"},
			},
			"scaleway_iam_ssh_key": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iamMod, "getSSHKey"),
				Docs: &tfbridge.DocInfo{Source: "iam_ssh_key.md"},
			},
			"scaleway_iam_user": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iamMod, "getUser"),
				Docs: &tfbridge.DocInfo{Source: "iam_user.md"},
			},

			// instance module
			"scaleway_instance_image": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getImage"),
				Docs: &tfbridge.DocInfo{Source: "instance_image.md"},
			},
			"scaleway_instance_ip": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getIP"),
				Docs: &tfbridge.DocInfo{Source: "instance_ip.md"},
			},
			"scaleway_instance_placement_group": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getPlacementGroup"),
				Docs: &tfbridge.DocInfo{Source: "instance_placement_group.md"},
			},
			"scaleway_instance_private_nic": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getPrivateNIC"),
				Docs: &tfbridge.DocInfo{Source: "instance_private_nic.md"},
			},
			"scaleway_instance_security_group": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getSecurityGroup"),
				Docs: &tfbridge.DocInfo{Source: "instance_security_group.md"},
			},
			"scaleway_instance_server": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getServer"),
				Docs: &tfbridge.DocInfo{Source: "instance_server.md"},
			},
			"scaleway_instance_servers": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getServers"),
				Docs: &tfbridge.DocInfo{Source: "instance_servers.md"},
			},
			"scaleway_instance_snapshot": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getSnapshot"),
				Docs: &tfbridge.DocInfo{Source: "instance_snapshot.md"},
			},
			"scaleway_instance_volume": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getVolume"),
				Docs: &tfbridge.DocInfo{Source: "instance_volume.md"},
			},
			"scaleway_marketplace_image": {
				Tok:  tfbridge.MakeDataSource(mainPkg, instanceMod, "getMarketplaceImage"),
				Docs: &tfbridge.DocInfo{Source: "marketplace_image.md"},
			},

			// iot module
			"scaleway_iot_device": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iotMod, "getDevice"),
				Docs: &tfbridge.DocInfo{Source: "iot_device.md"},
			},
			"scaleway_iot_hub": {
				Tok:  tfbridge.MakeDataSource(mainPkg, iotMod, "getHub"),
				Docs: &tfbridge.DocInfo{Source: "iot_hub.md"},
			},

			// ipam module
			"scaleway_ipam_ip": {
				Tok:  tfbridge.MakeDataSource(mainPkg, ipamMod, "getIP"),
				Docs: &tfbridge.DocInfo{Source: "ipam_ip.md"},
			},
			"scaleway_ipam_ips": {
				Tok:  tfbridge.MakeDataSource(mainPkg, ipamMod, "getIPs"),
				Docs: &tfbridge.DocInfo{Source: "ipam_ips.md"},
			},

			// kubernetes module
			"scaleway_k8s_cluster": {
				Tok:  tfbridge.MakeDataSource(mainPkg, kubernetesMod, "getCluster"),
				Docs: &tfbridge.DocInfo{Source: "k8s_cluster.md"},
			},
			"scaleway_k8s_pool": {
				Tok:  tfbridge.MakeDataSource(mainPkg, kubernetesMod, "getPool"),
				Docs: &tfbridge.DocInfo{Source: "k8s_pool.md"},
			},
			"scaleway_k8s_version": {
				Tok:  tfbridge.MakeDataSource(mainPkg, kubernetesMod, "getVersion"),
				Docs: &tfbridge.DocInfo{Source: "k8s_version.md"},
			},

			// loadbalancer module
			"scaleway_lb": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getLoadBalancer"),
				Docs: &tfbridge.DocInfo{Source: "lb.md"},
			},
			"scaleway_lb_acls": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getACLs"),
				Docs: &tfbridge.DocInfo{Source: "lb_acls.md"},
			},
			"scaleway_lb_backend": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getBackend"),
				Docs: &tfbridge.DocInfo{Source: "lb_backend.md"},
			},
			"scaleway_lb_backends": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getBackends"),
				Docs: &tfbridge.DocInfo{Source: "lb_backends.md"},
			},
			"scaleway_lb_certificate": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getCertificate"),
				Docs: &tfbridge.DocInfo{Source: "lb_certificate.md"},
			},
			"scaleway_lb_frontend": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getFrontend"),
				Docs: &tfbridge.DocInfo{Source: "lb_frontend.md"},
			},
			"scaleway_lb_frontends": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getFrontends"),
				Docs: &tfbridge.DocInfo{Source: "lb_frontends.md"},
			},
			"scaleway_lb_ip": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getIP"),
				Docs: &tfbridge.DocInfo{Source: "lb_ip.md"},
			},
			"scaleway_lb_ips": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getIPs"),
				Docs: &tfbridge.DocInfo{Source: "lb_ips.md"},
			},
			"scaleway_lb_route": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getRoute"),
				Docs: &tfbridge.DocInfo{Source: "lb_route.md"},
			},
			"scaleway_lb_routes": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getRoutes"),
				Docs: &tfbridge.DocInfo{Source: "lb_routes.md"},
			},
			"scaleway_lbs": {
				Tok:  tfbridge.MakeDataSource(mainPkg, loadBalancerMod, "getLoadBalancers"),
				Docs: &tfbridge.DocInfo{Source: "lbs.md"},
			},

			// mnq module
			"scaleway_mnq_sqs": {
				Tok:  tfbridge.MakeDataSource(mainPkg, mnqMod, "getSQS"),
				Docs: &tfbridge.DocInfo{Source: "mnq_sqs.md"},
			},

			// objectstorage module
			"scaleway_object_bucket": {
				Tok:  tfbridge.MakeDataSource(mainPkg, objectStorageMod, "getBucket"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket.md"},
			},
			"scaleway_object_bucket_policy": {
				Tok:  tfbridge.MakeDataSource(mainPkg, objectStorageMod, "getBucketPolicy"),
				Docs: &tfbridge.DocInfo{Source: "object_bucket_policy.md"},
			},

			// rdb module
			"scaleway_rdb_acl": {
				Tok:  tfbridge.MakeDataSource(mainPkg, rdbMod, "getACL"),
				Docs: &tfbridge.DocInfo{Source: "rdb_acl.md"},
			},
			"scaleway_rdb_database": {
				Tok:  tfbridge.MakeDataSource(mainPkg, rdbMod, "getDatabase"),
				Docs: &tfbridge.DocInfo{Source: "rdb_database.md"},
			},
			"scaleway_rdb_database_backup": {
				Tok:  tfbridge.MakeDataSource(mainPkg, rdbMod, "getDatabaseBackup"),
				Docs: &tfbridge.DocInfo{Source: "rdb_database_backup.md"},
			},
			"scaleway_rdb_instance": {
				Tok:  tfbridge.MakeDataSource(mainPkg, rdbMod, "getInstance"),
				Docs: &tfbridge.DocInfo{Source: "rdb_instance.md"},
			},
			"scaleway_rdb_privilege": {
				Tok:  tfbridge.MakeDataSource(mainPkg, rdbMod, "getPrivilege"),
				Docs: &tfbridge.DocInfo{Source: "rdb_privilege.md"},
			},

			// redis module
			"scaleway_redis_cluster": {
				Tok:  tfbridge.MakeDataSource(mainPkg, redisMod, "getCluster"),
				Docs: &tfbridge.DocInfo{Source: "redis_cluster.md"},
			},

			// secret module
			"scaleway_secret": {
				Tok:  tfbridge.MakeDataSource(mainPkg, secretMod, "getSecret"),
				Docs: &tfbridge.DocInfo{Source: "secret.md"},
			},
			"scaleway_secret_version": {
				Tok:  tfbridge.MakeDataSource(mainPkg, secretMod, "getVersion"),
				Docs: &tfbridge.DocInfo{Source: "secret_version.md"},
			},

			// serverless module
			"scaleway_container": {
				Tok:  tfbridge.MakeDataSource(mainPkg, serverlessMod, "getContainer"),
				Docs: &tfbridge.DocInfo{Source: "container.md"},
			},
			"scaleway_container_namespace": {
				Tok:  tfbridge.MakeDataSource(mainPkg, serverlessMod, "getContainerNamespace"),
				Docs: &tfbridge.DocInfo{Source: "container_namespace.md"},
			},
			"scaleway_function": {
				Tok:  tfbridge.MakeDataSource(mainPkg, serverlessMod, "getFunction"),
				Docs: &tfbridge.DocInfo{Source: "function.md"},
			},
			"scaleway_function_namespace": {
				Tok:  tfbridge.MakeDataSource(mainPkg, serverlessMod, "getFunctionNamespace"),
				Docs: &tfbridge.DocInfo{Source: "function_namespace.md"},
			},

			// tem module
			"scaleway_tem_domain": {
				Tok:  tfbridge.MakeDataSource(mainPkg, temMod, "getDomain"),
				Docs: &tfbridge.DocInfo{Source: "tem_domain.md"},
			},

			// vpc module
			"scaleway_vpc": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getVPC"),
				Docs: &tfbridge.DocInfo{Source: "vpc.md"},
			},
			"scaleway_vpc_gateway_network": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getGatewayNetwork"),
				Docs: &tfbridge.DocInfo{Source: "vpc_gateway_network.md"},
			},
			"scaleway_vpc_private_network": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPrivateNetwork"),
				Docs: &tfbridge.DocInfo{Source: "vpc_private_network.md"},
			},
			"scaleway_vpc_public_gateway": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPublicGateway"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway.md"},
			},
			"scaleway_vpc_public_gateway_dhcp": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPublicGatewayDHCP"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_dhcp.md"},
			},
			"scaleway_vpc_public_gateway_dhcp_reservation": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPublicGatewayDHCPReservation"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_dhcp_reservation.md"},
			},
			"scaleway_vpc_public_gateway_ip": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPublicGatewayIP"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_ip.md"},
			},
			"scaleway_vpc_public_gateway_pat_rule": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getPublicGatewayPATRule"),
				Docs: &tfbridge.DocInfo{Source: "vpc_public_gateway_pat_rule.md"},
			},
			"scaleway_vpcs": {
				Tok:  tfbridge.MakeDataSource(mainPkg, vpcMod, "getVPCs"),
				Docs: &tfbridge.DocInfo{Source: "vpcs.md"},
			},

			// webhosting module
			"scaleway_webhosting": {
				Tok:  tfbridge.MakeDataSource(mainPkg, webhostingMod, "getWebHosting"),
				Docs: &tfbridge.DocInfo{Source: "webhosting.md"},
			},
			"scaleway_webhosting_offer": {
				Tok:  tfbridge.MakeDataSource(mainPkg, webhostingMod, "getOffer"),
				Docs: &tfbridge.DocInfo{Source: "webhosting_offer.md"},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/raeumlich/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("scaleway_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
