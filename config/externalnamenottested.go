/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{

	// amp
	//
	// The prometheus alert manager definition can be imported using the workspace identifier
	"aws_prometheus_alert_manager_definition": config.ParameterAsIdentifier("workspace_id"),
	// The prometheus rule group namespace can be imported using the arn
	"aws_prometheus_rule_group_namespace": config.IdentifierFromProvider,
	// AMP Workspaces can be imported using the identifier
	"aws_prometheus_workspace": config.IdentifierFromProvider,

	// amplify
	//
	// Amplify domain association can be imported using app_id and domain_name: d2ypk4k47z8u6/example.com
	"aws_amplify_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .parameters.app_id }}/{{ .external_name }}"),

	// appconfig

	// AppConfig Applications can be imported using their application ID,
	"aws_appconfig_application": config.IdentifierFromProvider,
	// AppConfig Configuration Profiles can be imported by using the configuration profile ID and application ID separated by a colon (:)
	"aws_appconfig_configuration_profile": config.IdentifierFromProvider,
	// AppConfig Deployments can be imported by using the application ID, environment ID, and deployment number separated by a slash (/)
	"aws_appconfig_deployment": config.IdentifierFromProvider,
	// AppConfig Deployment Strategies can be imported by using their deployment strategy ID
	"aws_appconfig_deployment_strategy": config.IdentifierFromProvider,
	// AppConfig Environments can be imported by using the environment ID and application ID separated by a colon (:)
	"aws_appconfig_environment": config.IdentifierFromProvider,
	// AppConfig Hosted Configuration Versions can be imported by using the application ID, configuration profile ID, and version number separated by a slash (/)
	"aws_appconfig_hosted_configuration_version": config.IdentifierFromProvider,

	// appflow
	//
	// AppFlow flows can be imported using the arn
	"aws_appflow_flow": config.IdentifierFromProvider,

	// appintegrations
	//
	// Amazon AppIntegrations Event Integrations can be imported using the name
	"aws_appintegrations_event_integration": config.NameAsIdentifier,

	// apprunner
	//
	// App Runner Custom Domain Associations can be imported by using the domain_name and service_arn separated by a comma (,)
	"aws_apprunner_custom_domain_association": config.TemplatedStringAsIdentifier("domain_name", "{{ .external_name }},{{ .parameters.service_arn }}"),

	// appsync
	//
	// aws_appsync_api_cache can be imported using the AppSync API ID
	"aws_appsync_api_cache": config.IdentifierFromProvider,
	// aws_appsync_api_key can be imported using the AppSync API ID and key separated by :
	"aws_appsync_api_key": config.IdentifierFromProvider,
	// aws_appsync_datasource can be imported with their api_id, a hyphen, and name
	"aws_appsync_datasource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}-{{ .external_name }}"),
	// aws_appsync_domain_name can be imported using the AppSync domain name
	"aws_appsync_domain_name": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_domain_name_api_association can be imported using the AppSync domain name
	"aws_appsync_domain_name_api_association": config.ParameterAsIdentifier("domain_name"),
	// aws_appsync_function can be imported using the AppSync API ID and Function ID separated by -
	"aws_appsync_function": config.IdentifierFromProvider,
	// AppSync GraphQL API can be imported using the GraphQL API ID
	"aws_appsync_graphql_api": config.IdentifierFromProvider,
	// aws_appsync_resolver can be imported with their api_id, a hyphen, type, a hypen and field
	"aws_appsync_resolver": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}-{{ .parameters.type }}-{{ .parameters.field }}"),

	// autoscaling
	//
	// aws_autoscaling_group_tag can be imported by using the ASG name and key, separated by a comma (,)
	"aws_autoscaling_group_tag": config.TemplatedStringAsIdentifier("autoscaling_group_name", "{{ .external_name }},{{ .parameters.tag.key }}"),
	// AutoScaling Lifecycle Hooks can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_lifecycle_hook": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// No import
	"aws_autoscaling_notification": config.IdentifierFromProvider,
	// AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by /
	"aws_autoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// AutoScaling ScheduledAction can be imported using the auto-scaling-group-name and scheduled-action-name: auto-scaling-group-name/scheduled-action-name
	"aws_autoscaling_schedule": config.TemplatedStringAsIdentifier("scheduled_action_name", "{{ .parameters.autoscaling_group_name }}/{{ .external_name }}"),
	// Launch configurations can be imported using the name
	"aws_launch_configuration": config.NameAsIdentifier,

	// autoscalingplans
	//
	// Auto Scaling scaling plans can be imported using the name
	"aws_autoscalingplans_scaling_plan": config.NameAsIdentifier,

	// batch
	//
	// AWS Batch compute can be imported using the compute_environment_name
	"aws_batch_compute_environment": config.ParameterAsIdentifier("compute_environment_name"),
	// Batch Job Definition can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-definition/sample
	"aws_batch_job_definition": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-definition/{{ .external_name }}"),
	// Batch Job Queue can be imported using the arn: arn:aws:batch:us-east-1:123456789012:job-queue/sample
	"aws_batch_job_queue": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:job-queue/{{ .external_name }}"),
	// Batch Scheduling Policy can be imported using the arn: arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
	"aws_batch_scheduling_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:scheduling-policy/{{ .external_name }}"),

	// budgets
	//
	// Budgets can be imported using AccountID:BudgetName
	"aws_budgets_budget": config.TemplatedStringAsIdentifier("name", "{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// Budgets can be imported using AccountID:ActionID:BudgetName
	"aws_budgets_budget_action": config.IdentifierFromProvider,

	// ce
	//
	// aws_ce_cost_category can be imported using the id
	"aws_ce_cost_category": config.IdentifierFromProvider,

	// chime
	//
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector_group": config.NameAsIdentifier,
	// Chime Voice Connector Logging can be imported using the voice_connector_id
	"aws_chime_voice_connector_logging": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Origination can be imported using the voice_connector_id
	"aws_chime_voice_connector_origination": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Streaming can be imported using the voice_connector_id
	"aws_chime_voice_connector_streaming": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Termination can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination": config.ParameterAsIdentifier("voice_connector_id"),
	// Chime Voice Connector Termination Credentials can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination_credentials": config.ParameterAsIdentifier("voice_connector_id"),

	// cloud9
	//
	// No import
	"aws_cloud9_environment_ec2": config.IdentifierFromProvider,
	// Cloud9 environment membership can be imported using the environment-id#user-arn
	"aws_cloud9_environment_membership": config.TemplatedStringAsIdentifier("", "{{ .parameters.environment_id }}#{{ .parameters.user_arn }}"),

	// cloudcontrol
	//
	// No import
	"aws_cloudcontrolapi_resource": config.IdentifierFromProvider,

	// cloudformation
	//
	// Cloudformation Stacks can be imported using the name
	"aws_cloudformation_stack": config.NameAsIdentifier,
	// CloudFormation StackSets can be imported using the name
	"aws_cloudformation_stack_set": config.NameAsIdentifier,
	//
	"aws_cloudformation_stack_set_instance": config.IdentifierFromProvider,
	// aws_cloudformation_type can be imported with their type version Amazon Resource Name (ARN)
	"aws_cloudformation_type": config.IdentifierFromProvider,

	// cloudhsmv2
	//
	// CloudHSM v2 Clusters can be imported using the cluster id
	"aws_cloudhsm_v2_cluster": config.IdentifierFromProvider,
	// HSM modules can be imported using their HSM ID
	"aws_cloudhsm_v2_hsm": config.IdentifierFromProvider,

	// cloudtrail
	//
	// Cloudtrails can be imported using the name
	"aws_cloudtrail": config.NameAsIdentifier,
	// Event data stores can be imported using their arn
	"aws_cloudtrail_event_data_store": config.IdentifierFromProvider,

	// cloudwatchlogs
	//
	// CloudWatch Logs destinations can be imported using the name
	"aws_cloudwatch_log_destination": config.NameAsIdentifier,
	// CloudWatch Logs destination policies can be imported using the destination_name
	"aws_cloudwatch_log_destination_policy": config.ParameterAsIdentifier("destination_name"),
	// CloudWatch Logs subscription filter can be imported using the log group name and subscription filter name separated by |
	"aws_cloudwatch_log_subscription_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}|{{ .external_name }}"),

	// codeartifact
	//
	// CodeArtifact Domain can be imported using the CodeArtifact Domain arn
	"aws_codeartifact_domain": config.IdentifierFromProvider,
	// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN
	"aws_codeartifact_domain_permissions_policy": config.IdentifierFromProvider,
	// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository": config.IdentifierFromProvider,
	// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN
	"aws_codeartifact_repository_permissions_policy": config.IdentifierFromProvider,

	// codebuild
	//
	// CodeBuild Project can be imported using the name
	"aws_codebuild_project": config.NameAsIdentifier,
	// CodeBuild Report Group can be imported using the CodeBuild Report Group arn: arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
	"aws_codebuild_report_group": config.TemplatedStringAsIdentifier("name", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:report-group/{{ .external_name }}"),
	// CodeBuild Resource Policy can be imported using the CodeBuild Resource Policy arn
	"aws_codebuild_resource_policy": config.IdentifierFromProvider,
	// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn: arn:aws:codebuild:us-west-2:123456789:token:github
	"aws_codebuild_source_credential": config.TemplatedStringAsIdentifier("", "arn:aws:codebuild:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:token:{{ .parameters.token }}"),
	// CodeBuild Webhooks can be imported using the CodeBuild Project name
	"aws_codebuild_webhook": config.ParameterAsIdentifier("project_name"),

	// cognitoidp
	//
	// Cognito User Groups can be imported using the user_pool_id/name attributes concatenated
	"aws_cognito_user_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// No import
	"aws_cognito_user_in_group": config.IdentifierFromProvider,

	// configservice
	//
	// Config aggregate authorizations can be imported using account_id:region
	"aws_config_aggregate_authorization": config.TemplatedStringAsIdentifier("", "{{ .parameters.account_id }}:{{ .parameters.region }}"),
	// Config Organization Conformance Packs can be imported using the name
	"aws_config_organization_conformance_pack": config.NameAsIdentifier,
	// Config Organization Custom Rules can be imported using the name
	"aws_config_organization_custom_rule": config.NameAsIdentifier,
	// Config Organization Managed Rules can be imported using the name
	"aws_config_organization_managed_rule": config.NameAsIdentifier,

	// connect
	//
	// Amazon Connect User Hierarchy Groups can be imported using the instance_id and hierarchy_group_id separated by a colon (:)
	"aws_connect_user_hierarchy_group": config.IdentifierFromProvider,

	// cur
	//
	// Report Definitions can be imported using the report_name
	"aws_cur_report_definition": config.ParameterAsIdentifier("report_name"),

	// dataexchange
	//
	// DataExchange DataSets can be imported by their arn
	"aws_dataexchange_data_set": config.IdentifierFromProvider,
	// DataExchange Revisions can be imported by their data-set-id:revision-id
	"aws_dataexchange_revision": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline can be imported by using the id (Pipeline ID)
	"aws_datapipeline_pipeline": config.IdentifierFromProvider,
	// aws_datapipeline_pipeline_definition can be imported using the id
	"aws_datapipeline_pipeline_definition": config.IdentifierFromProvider,

	// datasync
	//
	// aws_datasync_agent can be imported by using the DataSync Agent Amazon Resource Name (ARN)
	"aws_datasync_agent": config.IdentifierFromProvider,
	// aws_datasync_location_efs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_efs": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_lustre_file_system can be imported by using the DataSync-ARN#FSx-Lustre-ARN
	"aws_datasync_location_fsx_lustre_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_openzfs_file_system can be imported by using the DataSync-ARN#FSx-openzfs-ARN
	"aws_datasync_location_fsx_openzfs_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_fsx_windows_file_system can be imported by using the DataSync-ARN#FSx-Windows-ARN
	"aws_datasync_location_fsx_windows_file_system": config.IdentifierFromProvider,
	// aws_datasync_location_hdfs can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_hdfs": config.IdentifierFromProvider,
	// aws_datasync_location_nfs can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_nfs": config.IdentifierFromProvider,
	// aws_datasync_location_s3 can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_location_s3": config.IdentifierFromProvider,
	// aws_datasync_location_smb can be imported by using the Amazon Resource Name (ARN)
	"aws_datasync_location_smb": config.IdentifierFromProvider,
	// aws_datasync_task can be imported by using the DataSync Task Amazon Resource Name (ARN)
	"aws_datasync_task": config.IdentifierFromProvider,

	// detective
	//
	// aws_detective_graph can be imported using the ARN
	"aws_detective_graph": config.IdentifierFromProvider,
	// aws_detective_invitation_accepter can be imported using the graph ARN
	"aws_detective_invitation_accepter": config.IdentifierFromProvider,
	// aws_detective_member can be imported using the ARN of the graph followed by the account ID of the member account
	"aws_detective_member": config.IdentifierFromProvider,

	// devicefarm
	//
	// DeviceFarm Device Pools can be imported by their arn
	"aws_devicefarm_device_pool": config.IdentifierFromProvider,
	// DeviceFarm Instance Profiles can be imported by their arn
	"aws_devicefarm_instance_profile": config.IdentifierFromProvider,
	// DeviceFarm Network Profiles can be imported by their arn
	"aws_devicefarm_network_profile": config.IdentifierFromProvider,
	// DeviceFarm Projects can be imported by their arn
	"aws_devicefarm_project": config.IdentifierFromProvider,
	// DeviceFarm Test Grid Projects can be imported by their arn
	"aws_devicefarm_test_grid_project": config.IdentifierFromProvider,
	// DeviceFarm Uploads can be imported by their arn
	"aws_devicefarm_upload": config.IdentifierFromProvider,

	// directconnect
	//
	// No import
	"aws_dx_bgp_peer": config.IdentifierFromProvider,
	// Direct Connect connections can be imported using the connection id
	"aws_dx_connection": config.IdentifierFromProvider,
	// No import
	"aws_dx_connection_association": config.IdentifierFromProvider,
	// No import
	"aws_dx_connection_confirmation": config.IdentifierFromProvider,
	// Direct Connect Gateways can be imported using the gateway id
	"aws_dx_gateway": config.IdentifierFromProvider,
	// Direct Connect gateway associations can be imported using dx_gateway_id together with associated_gateway_id
	// TODO: associated_gateway_id parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_dx_gateway_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.dx_gateway_id }}/{{ .parameters.associated_gateway_id }}"),
	//
	"aws_dx_gateway_association_proposal": config.IdentifierFromProvider,
	// No import
	"aws_dx_hosted_connection": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted private virtual interfaces can be imported using the vif id
	"aws_dx_hosted_private_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted public virtual interfaces can be imported using the vif id
	"aws_dx_hosted_public_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect hosted transit virtual interfaces can be imported using the vif id
	"aws_dx_hosted_transit_virtual_interface_accepter": config.ParameterAsIdentifier("virtual_interface_id"),
	// Direct Connect LAGs can be imported using the lag id
	"aws_dx_lag": config.IdentifierFromProvider,
	// Direct Connect private virtual interfaces can be imported using the vif id
	"aws_dx_private_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect public virtual interfaces can be imported using the vif id
	"aws_dx_public_virtual_interface": config.IdentifierFromProvider,
	// Direct Connect transit virtual interfaces can be imported using the vif id
	"aws_dx_transit_virtual_interface": config.IdentifierFromProvider,

	// dlm
	//
	// DLM lifecycle policies can be imported by their policy ID
	"aws_dlm_lifecycle_policy": config.IdentifierFromProvider,

	// dms
	//
	// Certificates can be imported using the certificate_id
	"aws_dms_certificate": config.ParameterAsIdentifier("certificate_id"),
	// Endpoints can be imported using the endpoint_id
	"aws_dms_endpoint": config.ParameterAsIdentifier("endpoint_id"),
	// Event subscriptions can be imported using the name
	"aws_dms_event_subscription": config.NameAsIdentifier,
	// Replication instances can be imported using the replication_instance_id
	"aws_dms_replication_instance": config.ParameterAsIdentifier("replication_instance_id"),
	// Replication subnet groups can be imported using the replication_subnet_group_id
	"aws_dms_replication_subnet_group": config.ParameterAsIdentifier("replication_subnet_group_id"),
	// Replication tasks can be imported using the replication_task_id
	"aws_dms_replication_task": config.ParameterAsIdentifier("replication_task_id"),

	// docdb
	//
	// DocumentDB Cluster Parameter Groups can be imported using the name
	"aws_docdb_cluster_parameter_group": config.NameAsIdentifier,
	// aws_docdb_cluster_snapshot can be imported by using the cluster snapshot identifier
	"aws_docdb_cluster_snapshot": config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	// DocDB Event Subscriptions can be imported using the name
	"aws_docdb_event_subscription": config.NameAsIdentifier,

	// ds
	//
	// Conditional forwarders can be imported using the directory id and remote_domain_name: d-1234567890:example.com
	"aws_directory_service_conditional_forwarder": config.TemplatedStringAsIdentifier("", "{{ .parameters.directory_id }}:{{ .parameters.remote_domain_name }}"),
	// DirectoryService directories can be imported using the directory id
	"aws_directory_service_directory": config.IdentifierFromProvider,
	// Directory Service Log Subscriptions can be imported using the directory id
	"aws_directory_service_log_subscription": config.ParameterAsIdentifier("directory_id"),

	// dynamodb
	//
	// aws_dynamodb_tag can be imported by using the DynamoDB resource identifier and key, separated by a comma (,)
	"aws_dynamodb_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_arn }},{{ .parameters.key }}"),

	// ec2
	//
	// aws_ami can be imported using the ID of the AMI
	"aws_ami": config.IdentifierFromProvider,
	// No import
	"aws_ami_copy": config.IdentifierFromProvider,
	// No import
	"aws_ami_from_instance": config.IdentifierFromProvider,
	// AMI Launch Permissions can be imported using [ACCOUNT-ID|GROUP-NAME|ORGANIZATION-ARN|ORGANIZATIONAL-UNIT-ARN]/IMAGE-ID
	"aws_ami_launch_permission": config.IdentifierFromProvider,
	// Customer Gateways can be imported using the id
	"aws_customer_gateway": config.IdentifierFromProvider,
	// Default Network ACLs can be imported using the id
	"aws_default_network_acl": config.IdentifierFromProvider,
	// Security Groups can be imported using the security group id
	"aws_default_security_group": config.IdentifierFromProvider,
	//
	"aws_ec2_client_vpn_authorization_rule": config.IdentifierFromProvider,
	// AWS Client VPN endpoints can be imported using the id value found via aws ec2 describe-client-vpn-endpoints
	"aws_ec2_client_vpn_endpoint": config.IdentifierFromProvider,
	// AWS Client VPN network associations can be imported using the endpoint ID and the association ID. Values are separated by a ,
	"aws_ec2_client_vpn_network_association": config.IdentifierFromProvider,
	// AWS Client VPN routes can be imported using the endpoint ID, target subnet ID, and destination CIDR block. All values are separated by a ,
	"aws_ec2_client_vpn_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.client_vpn_endpoint_id }},{{ .parameters.target_vpc_subnet_id }},{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_fleet can be imported by using the Fleet identifier
	"aws_ec2_fleet": config.IdentifierFromProvider,
	// aws_ec2_local_gateway_route can be imported by using the EC2 Local Gateway Route Table identifier and destination CIDR block separated by underscores (_)
	"aws_ec2_local_gateway_route": config.TemplatedStringAsIdentifier("", "{{ .parameters.local_gateway_route_table_id }}_{{ .parameters.destination_cidr_block }}"),
	// aws_ec2_local_gateway_route_table_vpc_association can be imported by using the Local Gateway Route Table VPC Association identifier
	"aws_ec2_local_gateway_route_table_vpc_association": config.IdentifierFromProvider,
	// aws_ec2_tag can be imported by using the EC2 resource identifier and key, separated by a comma (,)
	"aws_ec2_tag": config.TemplatedStringAsIdentifier("", "{{ .parameters.resource_id }}_{{ .parameters.key }}"),
	// Traffic mirror sessions can be imported using the id
	"aws_ec2_traffic_mirror_session": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_traffic_mirror_target": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_connect_peer can be imported by using the EC2 Transit Gateway Connect Peer identifier
	"aws_ec2_transit_gateway_connect_peer": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment_accepter can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment_accepter": config.IdentifierFromProvider,
	// Internet Gateway Attachments can be imported using the id
	"aws_internet_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_acl_association": config.IdentifierFromProvider,
	// No import
	"aws_snapshot_create_volume_permission": config.IdentifierFromProvider,
	// Spot Fleet Requests can be imported using id
	"aws_spot_fleet_request": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using ID of the connection, which is the VPC Endpoint Service ID and VPC Endpoint ID separated by underscore (_)
	"aws_vpc_endpoint_connection_accepter": config.TemplatedStringAsIdentifier("", "{{ .parameters.vpc_endpoint_service_id }}_{{ .parameters.vpc_endpoint_id }}"),
	// VPC Endpoint Policies can be imported using the id
	"aws_vpc_endpoint_policy": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_security_group_association": config.IdentifierFromProvider,
	// IPAMs can be imported using the ipam id
	"aws_vpc_ipam": config.IdentifierFromProvider,
	// IPAMs can be imported using the delegate account id
	"aws_vpc_ipam_organization_admin_account": config.ParameterAsIdentifier("delegated_admin_account_id"),
	// IPAMs can be imported using the ipam pool id
	"aws_vpc_ipam_pool": config.IdentifierFromProvider,
	// IPAMs can be imported using the <cidr>_<ipam-pool-id>
	"aws_vpc_ipam_pool_cidr": config.IdentifierFromProvider,
	// IPAMs can be imported using the allocation id
	"aws_vpc_ipam_pool_cidr_allocation": config.IdentifierFromProvider,
	// No import
	"aws_vpc_ipam_preview_next_cidr": config.IdentifierFromProvider,
	// IPAMs can be imported using the scope_id
	"aws_vpc_ipam_scope": config.IdentifierFromProvider,
	// aws_vpc_ipv6_cidr_block_association can be imported by using the VPC CIDR Association ID
	"aws_vpc_ipv6_cidr_block_association": config.IdentifierFromProvider,
	// VPC Peering Connection Accepters can be imported by using the Peering Connection ID
	"aws_vpc_peering_connection_accepter": config.ParameterAsIdentifier("vpc_peering_connection_id"),
	// VPC Peering Connection Options can be imported using the vpc peering id
	"aws_vpc_peering_connection_options": config.ParameterAsIdentifier("vpc_peering_connection_id"),
	// VPN Connections can be imported using the vpn connection id
	"aws_vpn_connection": config.IdentifierFromProvider,
	// No import
	"aws_vpn_connection_route": config.IdentifierFromProvider,
	// VPN Gateways can be imported using the vpn gateway id
	"aws_vpn_gateway": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_attachment": config.IdentifierFromProvider,
	// No import
	"aws_vpn_gateway_route_propagation": config.IdentifierFromProvider,
}