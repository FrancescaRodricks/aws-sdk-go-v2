// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountRoleStatus string

// Enum values for AccountRoleStatus
const (
	AccountRoleStatusReady           AccountRoleStatus = "READY"
	AccountRoleStatusCreating        AccountRoleStatus = "CREATING"
	AccountRoleStatusPendingDeletion AccountRoleStatus = "PENDING_DELETION"
	AccountRoleStatusDeleting        AccountRoleStatus = "DELETING"
	AccountRoleStatusDeleted         AccountRoleStatus = "DELETED"
)

// Values returns all known values for AccountRoleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccountRoleStatus) Values() []AccountRoleStatus {
	return []AccountRoleStatus{
		"READY",
		"CREATING",
		"PENDING_DELETION",
		"DELETING",
		"DELETED",
	}
}

type CustomerPolicyScopeIdType string

// Enum values for CustomerPolicyScopeIdType
const (
	CustomerPolicyScopeIdTypeAccount CustomerPolicyScopeIdType = "ACCOUNT"
	CustomerPolicyScopeIdTypeOrgUnit CustomerPolicyScopeIdType = "ORG_UNIT"
)

// Values returns all known values for CustomerPolicyScopeIdType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomerPolicyScopeIdType) Values() []CustomerPolicyScopeIdType {
	return []CustomerPolicyScopeIdType{
		"ACCOUNT",
		"ORG_UNIT",
	}
}

type CustomerPolicyStatus string

// Enum values for CustomerPolicyStatus
const (
	CustomerPolicyStatusActive          CustomerPolicyStatus = "ACTIVE"
	CustomerPolicyStatusOutOfAdminScope CustomerPolicyStatus = "OUT_OF_ADMIN_SCOPE"
)

// Values returns all known values for CustomerPolicyStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CustomerPolicyStatus) Values() []CustomerPolicyStatus {
	return []CustomerPolicyStatus{
		"ACTIVE",
		"OUT_OF_ADMIN_SCOPE",
	}
}

type DependentServiceName string

// Enum values for DependentServiceName
const (
	DependentServiceNameAWSConfig              DependentServiceName = "AWSCONFIG"
	DependentServiceNameAwswaf                 DependentServiceName = "AWSWAF"
	DependentServiceNameAWSShieldAdvanced      DependentServiceName = "AWSSHIELD_ADVANCED"
	DependentServiceNameAWSVirtualPrivateCloud DependentServiceName = "AWSVPC"
)

// Values returns all known values for DependentServiceName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DependentServiceName) Values() []DependentServiceName {
	return []DependentServiceName{
		"AWSCONFIG",
		"AWSWAF",
		"AWSSHIELD_ADVANCED",
		"AWSVPC",
	}
}

type DestinationType string

// Enum values for DestinationType
const (
	DestinationTypeIpv4       DestinationType = "IPV4"
	DestinationTypeIpv6       DestinationType = "IPV6"
	DestinationTypePrefixList DestinationType = "PREFIX_LIST"
)

// Values returns all known values for DestinationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DestinationType) Values() []DestinationType {
	return []DestinationType{
		"IPV4",
		"IPV6",
		"PREFIX_LIST",
	}
}

type FailedItemReason string

// Enum values for FailedItemReason
const (
	FailedItemReasonNotValidArn          FailedItemReason = "NOT_VALID_ARN"
	FailedItemReasonNotValidPartition    FailedItemReason = "NOT_VALID_PARTITION"
	FailedItemReasonNotValidRegion       FailedItemReason = "NOT_VALID_REGION"
	FailedItemReasonNotValidService      FailedItemReason = "NOT_VALID_SERVICE"
	FailedItemReasonNotValidResourceType FailedItemReason = "NOT_VALID_RESOURCE_TYPE"
	FailedItemReasonNotValidAccountId    FailedItemReason = "NOT_VALID_ACCOUNT_ID"
)

// Values returns all known values for FailedItemReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailedItemReason) Values() []FailedItemReason {
	return []FailedItemReason{
		"NOT_VALID_ARN",
		"NOT_VALID_PARTITION",
		"NOT_VALID_REGION",
		"NOT_VALID_SERVICE",
		"NOT_VALID_RESOURCE_TYPE",
		"NOT_VALID_ACCOUNT_ID",
	}
}

type FirewallDeploymentModel string

// Enum values for FirewallDeploymentModel
const (
	FirewallDeploymentModelCentralized FirewallDeploymentModel = "CENTRALIZED"
	FirewallDeploymentModelDistributed FirewallDeploymentModel = "DISTRIBUTED"
)

// Values returns all known values for FirewallDeploymentModel. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FirewallDeploymentModel) Values() []FirewallDeploymentModel {
	return []FirewallDeploymentModel{
		"CENTRALIZED",
		"DISTRIBUTED",
	}
}

type MarketplaceSubscriptionOnboardingStatus string

// Enum values for MarketplaceSubscriptionOnboardingStatus
const (
	MarketplaceSubscriptionOnboardingStatusNoSubscription MarketplaceSubscriptionOnboardingStatus = "NO_SUBSCRIPTION"
	MarketplaceSubscriptionOnboardingStatusNotComplete    MarketplaceSubscriptionOnboardingStatus = "NOT_COMPLETE"
	MarketplaceSubscriptionOnboardingStatusComplete       MarketplaceSubscriptionOnboardingStatus = "COMPLETE"
)

// Values returns all known values for MarketplaceSubscriptionOnboardingStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (MarketplaceSubscriptionOnboardingStatus) Values() []MarketplaceSubscriptionOnboardingStatus {
	return []MarketplaceSubscriptionOnboardingStatus{
		"NO_SUBSCRIPTION",
		"NOT_COMPLETE",
		"COMPLETE",
	}
}

type NetworkFirewallOverrideAction string

// Enum values for NetworkFirewallOverrideAction
const (
	NetworkFirewallOverrideActionDropToAlert NetworkFirewallOverrideAction = "DROP_TO_ALERT"
)

// Values returns all known values for NetworkFirewallOverrideAction. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (NetworkFirewallOverrideAction) Values() []NetworkFirewallOverrideAction {
	return []NetworkFirewallOverrideAction{
		"DROP_TO_ALERT",
	}
}

type OrganizationStatus string

// Enum values for OrganizationStatus
const (
	OrganizationStatusOnboarding          OrganizationStatus = "ONBOARDING"
	OrganizationStatusOnboardingComplete  OrganizationStatus = "ONBOARDING_COMPLETE"
	OrganizationStatusOffboarding         OrganizationStatus = "OFFBOARDING"
	OrganizationStatusOffboardingComplete OrganizationStatus = "OFFBOARDING_COMPLETE"
)

// Values returns all known values for OrganizationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationStatus) Values() []OrganizationStatus {
	return []OrganizationStatus{
		"ONBOARDING",
		"ONBOARDING_COMPLETE",
		"OFFBOARDING",
		"OFFBOARDING_COMPLETE",
	}
}

type PolicyComplianceStatusType string

// Enum values for PolicyComplianceStatusType
const (
	PolicyComplianceStatusTypeCompliant    PolicyComplianceStatusType = "COMPLIANT"
	PolicyComplianceStatusTypeNonCompliant PolicyComplianceStatusType = "NON_COMPLIANT"
)

// Values returns all known values for PolicyComplianceStatusType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PolicyComplianceStatusType) Values() []PolicyComplianceStatusType {
	return []PolicyComplianceStatusType{
		"COMPLIANT",
		"NON_COMPLIANT",
	}
}

type RemediationActionType string

// Enum values for RemediationActionType
const (
	RemediationActionTypeRemove RemediationActionType = "REMOVE"
	RemediationActionTypeModify RemediationActionType = "MODIFY"
)

// Values returns all known values for RemediationActionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RemediationActionType) Values() []RemediationActionType {
	return []RemediationActionType{
		"REMOVE",
		"MODIFY",
	}
}

type ResourceSetStatus string

// Enum values for ResourceSetStatus
const (
	ResourceSetStatusActive          ResourceSetStatus = "ACTIVE"
	ResourceSetStatusOutOfAdminScope ResourceSetStatus = "OUT_OF_ADMIN_SCOPE"
)

// Values returns all known values for ResourceSetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSetStatus) Values() []ResourceSetStatus {
	return []ResourceSetStatus{
		"ACTIVE",
		"OUT_OF_ADMIN_SCOPE",
	}
}

type RuleOrder string

// Enum values for RuleOrder
const (
	RuleOrderStrictOrder        RuleOrder = "STRICT_ORDER"
	RuleOrderDefaultActionOrder RuleOrder = "DEFAULT_ACTION_ORDER"
)

// Values returns all known values for RuleOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RuleOrder) Values() []RuleOrder {
	return []RuleOrder{
		"STRICT_ORDER",
		"DEFAULT_ACTION_ORDER",
	}
}

type SecurityServiceType string

// Enum values for SecurityServiceType
const (
	SecurityServiceTypeWaf                        SecurityServiceType = "WAF"
	SecurityServiceTypeWafv2                      SecurityServiceType = "WAFV2"
	SecurityServiceTypeShieldAdvanced             SecurityServiceType = "SHIELD_ADVANCED"
	SecurityServiceTypeSecurityGroupsCommon       SecurityServiceType = "SECURITY_GROUPS_COMMON"
	SecurityServiceTypeSecurityGroupsContentAudit SecurityServiceType = "SECURITY_GROUPS_CONTENT_AUDIT"
	SecurityServiceTypeSecurityGroupsUsageAudit   SecurityServiceType = "SECURITY_GROUPS_USAGE_AUDIT"
	SecurityServiceTypeNetworkFirewall            SecurityServiceType = "NETWORK_FIREWALL"
	SecurityServiceTypeDnsFirewall                SecurityServiceType = "DNS_FIREWALL"
	SecurityServiceTypeThirdPartyFirewall         SecurityServiceType = "THIRD_PARTY_FIREWALL"
	SecurityServiceTypeImportNetworkFirewall      SecurityServiceType = "IMPORT_NETWORK_FIREWALL"
)

// Values returns all known values for SecurityServiceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SecurityServiceType) Values() []SecurityServiceType {
	return []SecurityServiceType{
		"WAF",
		"WAFV2",
		"SHIELD_ADVANCED",
		"SECURITY_GROUPS_COMMON",
		"SECURITY_GROUPS_CONTENT_AUDIT",
		"SECURITY_GROUPS_USAGE_AUDIT",
		"NETWORK_FIREWALL",
		"DNS_FIREWALL",
		"THIRD_PARTY_FIREWALL",
		"IMPORT_NETWORK_FIREWALL",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeGateway                   TargetType = "GATEWAY"
	TargetTypeCarrierGateway            TargetType = "CARRIER_GATEWAY"
	TargetTypeInstance                  TargetType = "INSTANCE"
	TargetTypeLocalGateway              TargetType = "LOCAL_GATEWAY"
	TargetTypeNatGateway                TargetType = "NAT_GATEWAY"
	TargetTypeNetworkInterface          TargetType = "NETWORK_INTERFACE"
	TargetTypeVPCEndpoint               TargetType = "VPC_ENDPOINT"
	TargetTypeVPCPeeringConnection      TargetType = "VPC_PEERING_CONNECTION"
	TargetTypeEgressOnlyInternetGateway TargetType = "EGRESS_ONLY_INTERNET_GATEWAY"
	TargetTypeTransitGateway            TargetType = "TRANSIT_GATEWAY"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"GATEWAY",
		"CARRIER_GATEWAY",
		"INSTANCE",
		"LOCAL_GATEWAY",
		"NAT_GATEWAY",
		"NETWORK_INTERFACE",
		"VPC_ENDPOINT",
		"VPC_PEERING_CONNECTION",
		"EGRESS_ONLY_INTERNET_GATEWAY",
		"TRANSIT_GATEWAY",
	}
}

type ThirdPartyFirewall string

// Enum values for ThirdPartyFirewall
const (
	ThirdPartyFirewallPaloAltoNetworksCloudNgfw    ThirdPartyFirewall = "PALO_ALTO_NETWORKS_CLOUD_NGFW"
	ThirdPartyFirewallFortigateCloudNativeFirewall ThirdPartyFirewall = "FORTIGATE_CLOUD_NATIVE_FIREWALL"
)

// Values returns all known values for ThirdPartyFirewall. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThirdPartyFirewall) Values() []ThirdPartyFirewall {
	return []ThirdPartyFirewall{
		"PALO_ALTO_NETWORKS_CLOUD_NGFW",
		"FORTIGATE_CLOUD_NATIVE_FIREWALL",
	}
}

type ThirdPartyFirewallAssociationStatus string

// Enum values for ThirdPartyFirewallAssociationStatus
const (
	ThirdPartyFirewallAssociationStatusOnboarding       ThirdPartyFirewallAssociationStatus = "ONBOARDING"
	ThirdPartyFirewallAssociationStatusOnboardComplete  ThirdPartyFirewallAssociationStatus = "ONBOARD_COMPLETE"
	ThirdPartyFirewallAssociationStatusOffboarding      ThirdPartyFirewallAssociationStatus = "OFFBOARDING"
	ThirdPartyFirewallAssociationStatusOffboardComplete ThirdPartyFirewallAssociationStatus = "OFFBOARD_COMPLETE"
	ThirdPartyFirewallAssociationStatusNotExist         ThirdPartyFirewallAssociationStatus = "NOT_EXIST"
)

// Values returns all known values for ThirdPartyFirewallAssociationStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ThirdPartyFirewallAssociationStatus) Values() []ThirdPartyFirewallAssociationStatus {
	return []ThirdPartyFirewallAssociationStatus{
		"ONBOARDING",
		"ONBOARD_COMPLETE",
		"OFFBOARDING",
		"OFFBOARD_COMPLETE",
		"NOT_EXIST",
	}
}

type ViolationReason string

// Enum values for ViolationReason
const (
	ViolationReasonWebAclMissingRuleGroup                  ViolationReason = "WEB_ACL_MISSING_RULE_GROUP"
	ViolationReasonResourceMissingWebAcl                   ViolationReason = "RESOURCE_MISSING_WEB_ACL"
	ViolationReasonResourceIncorrectWebAcl                 ViolationReason = "RESOURCE_INCORRECT_WEB_ACL"
	ViolationReasonResourceMissingShieldProtection         ViolationReason = "RESOURCE_MISSING_SHIELD_PROTECTION"
	ViolationReasonResourceMissingWebaclOrShieldProtection ViolationReason = "RESOURCE_MISSING_WEB_ACL_OR_SHIELD_PROTECTION"
	ViolationReasonResourceMissingSecurityGroup            ViolationReason = "RESOURCE_MISSING_SECURITY_GROUP"
	ViolationReasonResourceViolatesAuditSecurityGroup      ViolationReason = "RESOURCE_VIOLATES_AUDIT_SECURITY_GROUP"
	ViolationReasonSecurityGroupUnused                     ViolationReason = "SECURITY_GROUP_UNUSED"
	ViolationReasonSecurityGroupRedundant                  ViolationReason = "SECURITY_GROUP_REDUNDANT"
	ViolationReasonFMSCreatedSecurityGroupEdited           ViolationReason = "FMS_CREATED_SECURITY_GROUP_EDITED"
	ViolationReasonMissingFirewall                         ViolationReason = "MISSING_FIREWALL"
	ViolationReasonMissingFirewallSubnetInAZ               ViolationReason = "MISSING_FIREWALL_SUBNET_IN_AZ"
	ViolationReasonMissingExpectedRouteTable               ViolationReason = "MISSING_EXPECTED_ROUTE_TABLE"
	ViolationReasonNetworkFirewallPolicyModified           ViolationReason = "NETWORK_FIREWALL_POLICY_MODIFIED"
	ViolationReasonFirewallSubnetIsOutOfScope              ViolationReason = "FIREWALL_SUBNET_IS_OUT_OF_SCOPE"
	ViolationReasonInternetGatewayMissingExpectedRoute     ViolationReason = "INTERNET_GATEWAY_MISSING_EXPECTED_ROUTE"
	ViolationReasonFirewallSubnetMissingExpectedRoute      ViolationReason = "FIREWALL_SUBNET_MISSING_EXPECTED_ROUTE"
	ViolationReasonUnexpectedFirewallRoutes                ViolationReason = "UNEXPECTED_FIREWALL_ROUTES"
	ViolationReasonUnexpectedTargetGatewayRoutes           ViolationReason = "UNEXPECTED_TARGET_GATEWAY_ROUTES"
	ViolationReasonTrafficInspectionCrossesAZBoundary      ViolationReason = "TRAFFIC_INSPECTION_CROSSES_AZ_BOUNDARY"
	ViolationReasonInvalidRouteConfiguration               ViolationReason = "INVALID_ROUTE_CONFIGURATION"
	ViolationReasonMissingTargetGateway                    ViolationReason = "MISSING_TARGET_GATEWAY"
	ViolationReasonInternetTrafficNotInspected             ViolationReason = "INTERNET_TRAFFIC_NOT_INSPECTED"
	ViolationReasonBlackHoleRouteDetected                  ViolationReason = "BLACK_HOLE_ROUTE_DETECTED"
	ViolationReasonBlackHoleRouteDetectedInFirewallSubnet  ViolationReason = "BLACK_HOLE_ROUTE_DETECTED_IN_FIREWALL_SUBNET"
	ViolationReasonResourceMissingDnsFirewall              ViolationReason = "RESOURCE_MISSING_DNS_FIREWALL"
	ViolationReasonRouteHasOutOfScopeEndpoint              ViolationReason = "ROUTE_HAS_OUT_OF_SCOPE_ENDPOINT"
	ViolationReasonFirewallSubnetMissingVPCEndpoint        ViolationReason = "FIREWALL_SUBNET_MISSING_VPCE_ENDPOINT"
)

// Values returns all known values for ViolationReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ViolationReason) Values() []ViolationReason {
	return []ViolationReason{
		"WEB_ACL_MISSING_RULE_GROUP",
		"RESOURCE_MISSING_WEB_ACL",
		"RESOURCE_INCORRECT_WEB_ACL",
		"RESOURCE_MISSING_SHIELD_PROTECTION",
		"RESOURCE_MISSING_WEB_ACL_OR_SHIELD_PROTECTION",
		"RESOURCE_MISSING_SECURITY_GROUP",
		"RESOURCE_VIOLATES_AUDIT_SECURITY_GROUP",
		"SECURITY_GROUP_UNUSED",
		"SECURITY_GROUP_REDUNDANT",
		"FMS_CREATED_SECURITY_GROUP_EDITED",
		"MISSING_FIREWALL",
		"MISSING_FIREWALL_SUBNET_IN_AZ",
		"MISSING_EXPECTED_ROUTE_TABLE",
		"NETWORK_FIREWALL_POLICY_MODIFIED",
		"FIREWALL_SUBNET_IS_OUT_OF_SCOPE",
		"INTERNET_GATEWAY_MISSING_EXPECTED_ROUTE",
		"FIREWALL_SUBNET_MISSING_EXPECTED_ROUTE",
		"UNEXPECTED_FIREWALL_ROUTES",
		"UNEXPECTED_TARGET_GATEWAY_ROUTES",
		"TRAFFIC_INSPECTION_CROSSES_AZ_BOUNDARY",
		"INVALID_ROUTE_CONFIGURATION",
		"MISSING_TARGET_GATEWAY",
		"INTERNET_TRAFFIC_NOT_INSPECTED",
		"BLACK_HOLE_ROUTE_DETECTED",
		"BLACK_HOLE_ROUTE_DETECTED_IN_FIREWALL_SUBNET",
		"RESOURCE_MISSING_DNS_FIREWALL",
		"ROUTE_HAS_OUT_OF_SCOPE_ENDPOINT",
		"FIREWALL_SUBNET_MISSING_VPCE_ENDPOINT",
	}
}
