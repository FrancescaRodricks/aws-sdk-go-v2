// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

type Application struct {

	// Application aggregated status.
	ApplicationAggregatedStatus *ApplicationAggregatedStatus

	// Application ID.
	ApplicationID *string

	// Application ARN.
	Arn *string

	// Application creation dateTime.
	CreationDateTime *string

	// Application description.
	Description *string

	// Application archival status.
	IsArchived *bool

	// Application last modified dateTime.
	LastModifiedDateTime *string

	// Application name.
	Name *string

	// Application tags.
	Tags map[string]string

	// Application wave ID.
	WaveID *string

	noSmithyDocumentSerde
}

// Application aggregated status.
type ApplicationAggregatedStatus struct {

	// Application aggregated status health status.
	HealthStatus ApplicationHealthStatus

	// Application aggregated status last update dateTime.
	LastUpdateDateTime *string

	// Application aggregated status progress status.
	ProgressStatus ApplicationProgressStatus

	// Application aggregated status total source servers amount.
	TotalSourceServers int64

	noSmithyDocumentSerde
}

// The request to change the source server migration lifecycle state.
type ChangeServerLifeCycleStateSourceServerLifecycle struct {

	// The request to change the source server migration lifecycle state.
	//
	// This member is required.
	State ChangeServerLifeCycleStateSourceServerLifecycleState

	noSmithyDocumentSerde
}

// Source server CPU information.
type CPU struct {

	// The number of CPU cores on the source server.
	Cores int64

	// The source server's CPU model name.
	ModelName *string

	noSmithyDocumentSerde
}

// Error in data replication.
type DataReplicationError struct {

	// Error in data replication.
	Error DataReplicationErrorString

	// Error in data replication.
	RawError *string

	noSmithyDocumentSerde
}

// Request data replication info.
type DataReplicationInfo struct {

	// Error in obtaining data replication info.
	DataReplicationError *DataReplicationError

	// Request to query whether data replication has been initiated.
	DataReplicationInitiation *DataReplicationInitiation

	// Request to query the data replication state.
	DataReplicationState DataReplicationState

	// Request to query the time when data replication will be complete.
	EtaDateTime *string

	// Request to query data replication lag duration.
	LagDuration *string

	// Request to query data replication last snapshot time.
	LastSnapshotDateTime *string

	// Request to query disks replicated.
	ReplicatedDisks []DataReplicationInfoReplicatedDisk

	noSmithyDocumentSerde
}

// Request to query disks replicated.
type DataReplicationInfoReplicatedDisk struct {

	// Request to query data replication backlog size in bytes.
	BackloggedStorageBytes int64

	// Request to query device name.
	DeviceName *string

	// Request to query amount of data replicated in bytes.
	ReplicatedStorageBytes int64

	// Request to query amount of data rescanned in bytes.
	RescannedStorageBytes int64

	// Request to query total amount of data replicated in bytes.
	TotalStorageBytes int64

	noSmithyDocumentSerde
}

// Data replication initiation.
type DataReplicationInitiation struct {

	// Request to query next data initiation date and time.
	NextAttemptDateTime *string

	// Request to query data initiation start date and time.
	StartDateTime *string

	// Request to query data initiation steps.
	Steps []DataReplicationInitiationStep

	noSmithyDocumentSerde
}

// Data replication initiation step.
type DataReplicationInitiationStep struct {

	// Request to query data initiation step name.
	Name DataReplicationInitiationStepName

	// Request to query data initiation status.
	Status DataReplicationInitiationStepStatus

	noSmithyDocumentSerde
}

// Request to describe Job log filters.
type DescribeJobsRequestFilters struct {

	// Request to describe Job log filters by date.
	FromDate *string

	// Request to describe Job log filters by job ID.
	JobIDs []string

	// Request to describe job log items by last date.
	ToDate *string

	noSmithyDocumentSerde
}

// Request to filter Source Servers list.
type DescribeSourceServersRequestFilters struct {

	// Request to filter Source Servers list by application IDs.
	ApplicationIDs []string

	// Request to filter Source Servers list by archived.
	IsArchived *bool

	// Request to filter Source Servers list by life cycle states.
	LifeCycleStates []LifeCycleState

	// Request to filter Source Servers list by replication type.
	ReplicationTypes []ReplicationType

	// Request to filter Source Servers list by Source Server ID.
	SourceServerIDs []string

	noSmithyDocumentSerde
}

// The disk identifier.
type Disk struct {

	// The amount of storage on the disk in bytes.
	Bytes int64

	// The disk or device name.
	DeviceName *string

	noSmithyDocumentSerde
}

// Error details.
type ErrorDetails struct {

	// Error details code.
	Code *string

	// Error details message.
	Message *string

	// Error details resourceId.
	ResourceId *string

	// Error details resourceType.
	ResourceType *string

	noSmithyDocumentSerde
}

// Identification hints.
type IdentificationHints struct {

	// AWS Instance ID identification hint.
	AwsInstanceID *string

	// FQDN address identification hint.
	Fqdn *string

	// Hostname identification hint.
	Hostname *string

	// vCenter VM path identification hint.
	VmPath *string

	// vmWare UUID identification hint.
	VmWareUuid *string

	noSmithyDocumentSerde
}

// Job.
type Job struct {

	// Job ID.
	//
	// This member is required.
	JobID *string

	// the ARN of the specific Job.
	Arn *string

	// Job creation time.
	CreationDateTime *string

	// Job end time.
	EndDateTime *string

	// Job initiated by field.
	InitiatedBy InitiatedBy

	// Servers participating in a specific Job.
	ParticipatingServers []ParticipatingServer

	// Job status.
	Status JobStatus

	// Tags associated with specific Job.
	Tags map[string]string

	// Job type.
	Type JobType

	noSmithyDocumentSerde
}

// Job log.
type JobLog struct {

	// Job log event.
	Event JobLogEvent

	// Job event data
	EventData *JobLogEventData

	// Job log event date and time.
	LogDateTime *string

	noSmithyDocumentSerde
}

// Job log data
type JobLogEventData struct {

	// Job Event conversion Server ID.
	ConversionServerID *string

	// Job error.
	RawError *string

	// Job Event Source Server ID.
	SourceServerID *string

	// Job Event Target instance ID.
	TargetInstanceID *string

	noSmithyDocumentSerde
}

// Launch Status of the Job Post Launch Actions.
type JobPostLaunchActionsLaunchStatus struct {

	// AWS Systems Manager Document's execution ID of the of the Job Post Launch
	// Actions.
	ExecutionID *string

	// AWS Systems Manager Document's execution status.
	ExecutionStatus PostLaunchActionExecutionStatus

	// AWS Systems Manager Document's failure reason.
	FailureReason *string

	// AWS Systems Manager's Document of the of the Job Post Launch Actions.
	SsmDocument *SsmDocument

	// AWS Systems Manager Document type.
	SsmDocumentType SsmDocumentType

	noSmithyDocumentSerde
}

type LaunchConfigurationTemplate struct {

	// ID of the Launch Configuration Template.
	//
	// This member is required.
	LaunchConfigurationTemplateID *string

	// ARN of the Launch Configuration Template.
	Arn *string

	// Associate public Ip address.
	AssociatePublicIpAddress *bool

	// Launch configuration template boot mode.
	BootMode BootMode

	// Copy private Ip.
	CopyPrivateIp *bool

	// Copy tags.
	CopyTags *bool

	// EC2 launch template ID.
	Ec2LaunchTemplateID *string

	// Enable map auto tagging.
	EnableMapAutoTagging *bool

	// Large volume config.
	LargeVolumeConf *LaunchTemplateDiskConf

	// Launch disposition.
	LaunchDisposition LaunchDisposition

	// Configure Licensing.
	Licensing *Licensing

	// Launch configuration template map auto tagging MPE ID.
	MapAutoTaggingMpeID *string

	// Post Launch Actions of the Launch Configuration Template.
	PostLaunchActions *PostLaunchActions

	// Small volume config.
	SmallVolumeConf *LaunchTemplateDiskConf

	// Small volume maximum size.
	SmallVolumeMaxSize int64

	// Tags of the Launch Configuration Template.
	Tags map[string]string

	// Target instance type right-sizing method.
	TargetInstanceTypeRightSizingMethod TargetInstanceTypeRightSizingMethod

	noSmithyDocumentSerde
}

// Launched instance.
type LaunchedInstance struct {

	// Launched instance EC2 ID.
	Ec2InstanceID *string

	// Launched instance first boot.
	FirstBoot FirstBoot

	// Launched instance Job ID.
	JobID *string

	noSmithyDocumentSerde
}

// Launch template disk configuration.
type LaunchTemplateDiskConf struct {

	// Launch template disk iops configuration.
	Iops int64

	// Launch template disk throughput configuration.
	Throughput int64

	// Launch template disk volume type configuration.
	VolumeType VolumeType

	noSmithyDocumentSerde
}

// Configure Licensing.
type Licensing struct {

	// Configure BYOL OS licensing.
	OsByol *bool

	noSmithyDocumentSerde
}

// Lifecycle.
type LifeCycle struct {

	// Lifecycle added to service data and time.
	AddedToServiceDateTime *string

	// Lifecycle elapsed time and duration.
	ElapsedReplicationDuration *string

	// Lifecycle replication initiation date and time.
	FirstByteDateTime *string

	// Lifecycle last Cutover.
	LastCutover *LifeCycleLastCutover

	// Lifecycle last seen date and time.
	LastSeenByServiceDateTime *string

	// Lifecycle last Test.
	LastTest *LifeCycleLastTest

	// Lifecycle state.
	State LifeCycleState

	noSmithyDocumentSerde
}

// Lifecycle last Cutover .
type LifeCycleLastCutover struct {

	// Lifecycle Cutover finalized date and time.
	Finalized *LifeCycleLastCutoverFinalized

	// Lifecycle last Cutover initiated.
	Initiated *LifeCycleLastCutoverInitiated

	// Lifecycle last Cutover reverted.
	Reverted *LifeCycleLastCutoverReverted

	noSmithyDocumentSerde
}

// Lifecycle Cutover finalized
type LifeCycleLastCutoverFinalized struct {

	// Lifecycle Cutover finalized date and time.
	ApiCallDateTime *string

	noSmithyDocumentSerde
}

// Lifecycle last Cutover initiated.
type LifeCycleLastCutoverInitiated struct {

	//
	ApiCallDateTime *string

	// Lifecycle last Cutover initiated by Job ID.
	JobID *string

	noSmithyDocumentSerde
}

// Lifecycle last Cutover reverted.
type LifeCycleLastCutoverReverted struct {

	// Lifecycle last Cutover reverted API call date time.
	ApiCallDateTime *string

	noSmithyDocumentSerde
}

// Lifecycle last Test.
type LifeCycleLastTest struct {

	// Lifecycle last Test finalized.
	Finalized *LifeCycleLastTestFinalized

	// Lifecycle last Test initiated.
	Initiated *LifeCycleLastTestInitiated

	// Lifecycle last Test reverted.
	Reverted *LifeCycleLastTestReverted

	noSmithyDocumentSerde
}

// Lifecycle last Test finalized.
type LifeCycleLastTestFinalized struct {

	// Lifecycle Test failed API call date and time.
	ApiCallDateTime *string

	noSmithyDocumentSerde
}

// Lifecycle last Test initiated.
type LifeCycleLastTestInitiated struct {

	// Lifecycle last Test initiated API call date and time.
	ApiCallDateTime *string

	// Lifecycle last Test initiated Job ID.
	JobID *string

	noSmithyDocumentSerde
}

// Lifecycle last Test reverted.
type LifeCycleLastTestReverted struct {

	// Lifecycle last Test reverted API call date and time.
	ApiCallDateTime *string

	noSmithyDocumentSerde
}

// Applications list filters.
type ListApplicationsRequestFilters struct {

	// Filter applications list by application ID.
	ApplicationIDs []string

	// Filter applications list by archival status.
	IsArchived *bool

	// Filter applications list by wave ID.
	WaveIDs []string

	noSmithyDocumentSerde
}

// Waves list filters.
type ListWavesRequestFilters struct {

	// Filter waves list by archival status.
	IsArchived *bool

	// Filter waves list by wave ID.
	WaveIDs []string

	noSmithyDocumentSerde
}

// Network interface.
type NetworkInterface struct {

	// Network interface IPs.
	Ips []string

	// Network interface primary IP.
	IsPrimary *bool

	// Network interface Mac address.
	MacAddress *string

	noSmithyDocumentSerde
}

// Operating System.
type OS struct {

	// OS full string.
	FullString *string

	noSmithyDocumentSerde
}

// Server participating in Job.
type ParticipatingServer struct {

	// Participating server Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	// Participating server launch status.
	LaunchStatus LaunchStatus

	// Participating server's launched ec2 instance ID.
	LaunchedEc2InstanceID *string

	// Participating server's Post Launch Actions Status.
	PostLaunchActionsStatus *PostLaunchActionsStatus

	noSmithyDocumentSerde
}

// Post Launch Actions to executed on the Test or Cutover instance.
type PostLaunchActions struct {

	// AWS Systems Manager Command's CloudWatch log group name.
	CloudWatchLogGroupName *string

	// Deployment type in which AWS Systems Manager Documents will be executed.
	Deployment PostLaunchActionsDeploymentType

	// AWS Systems Manager Command's logs S3 log bucket.
	S3LogBucket *string

	// AWS Systems Manager Command's logs S3 output key prefix.
	S3OutputKeyPrefix *string

	// AWS Systems Manager Documents.
	SsmDocuments []SsmDocument

	noSmithyDocumentSerde
}

// Status of the Post Launch Actions running on the Test or Cutover instance.
type PostLaunchActionsStatus struct {

	// List of Post Launch Action status.
	PostLaunchActionsLaunchStatusList []JobPostLaunchActionsLaunchStatus

	// Time where the AWS Systems Manager was detected as running on the Test or
	// Cutover instance.
	SsmAgentDiscoveryDatetime *string

	noSmithyDocumentSerde
}

// Replication Configuration replicated disk.
type ReplicationConfigurationReplicatedDisk struct {

	// Replication Configuration replicated disk device name.
	DeviceName *string

	// Replication Configuration replicated disk IOPs.
	Iops int64

	// Replication Configuration replicated disk boot disk.
	IsBootDisk *bool

	// Replication Configuration replicated disk staging disk type.
	StagingDiskType ReplicationConfigurationReplicatedDiskStagingDiskType

	// Replication Configuration replicated disk throughput.
	Throughput int64

	noSmithyDocumentSerde
}

type ReplicationConfigurationTemplate struct {

	// Replication Configuration template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// Replication Configuration template ARN.
	Arn *string

	// Replication Configuration template associate default Application Migration
	// Service Security group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration template bandwidth throttling.
	BandwidthThrottling int64

	// Replication Configuration template create Public IP.
	CreatePublicIP *bool

	// Replication Configuration template data plane routing.
	DataPlaneRouting ReplicationConfigurationDataPlaneRouting

	// Replication Configuration template use default large Staging Disk type.
	DefaultLargeStagingDiskType ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration template EBS encryption.
	EbsEncryption ReplicationConfigurationEbsEncryption

	// Replication Configuration template EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration template server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration template server Security Groups IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration template Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration template Staging Area Tags.
	StagingAreaTags map[string]string

	// Replication Configuration template Tags.
	Tags map[string]string

	// Replication Configuration template use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	noSmithyDocumentSerde
}

// Source server properties.
type SourceProperties struct {

	// Source Server CPUs.
	Cpus []CPU

	// Source Server disks.
	Disks []Disk

	// Source server identification hints.
	IdentificationHints *IdentificationHints

	// Source server last update date and time.
	LastUpdatedDateTime *string

	// Source server network interfaces.
	NetworkInterfaces []NetworkInterface

	// Source server OS.
	Os *OS

	// Source server RAM in bytes.
	RamBytes int64

	// Source server recommended instance type.
	RecommendedInstanceType *string

	noSmithyDocumentSerde
}

type SourceServer struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source server data replication info.
	DataReplicationInfo *DataReplicationInfo

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *LifeCycle

	// Source server replication type.
	ReplicationType ReplicationType

	// Source server properties.
	SourceProperties *SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server vCenter client id.
	VcenterClientID *string

	noSmithyDocumentSerde
}

type SourceServerActionDocument struct {

	// Source server post migration custom action ID.
	ActionID *string

	// Source server post migration custom action name.
	ActionName *string

	// Source server post migration custom action active status.
	Active *bool

	// Source server post migration custom action document identifier.
	DocumentIdentifier *string

	// Source server post migration custom action document version.
	DocumentVersion *string

	// Source server post migration custom action must succeed for cutover.
	MustSucceedForCutover *bool

	// Source server post migration custom action order.
	Order int32

	// Source server post migration custom action parameters.
	Parameters map[string][]SsmParameterStoreParameter

	// Source server post migration custom action timeout in seconds.
	TimeoutSeconds int32

	noSmithyDocumentSerde
}

// Source server post migration custom action filters.
type SourceServerActionsRequestFilters struct {

	// Action IDs to filter source server post migration custom actions by.
	ActionIDs []string

	noSmithyDocumentSerde
}

// AWS Systems Manager Document.
type SsmDocument struct {

	// User-friendly name for the AWS Systems Manager Document.
	//
	// This member is required.
	ActionName *string

	// AWS Systems Manager Document name or full ARN.
	//
	// This member is required.
	SsmDocumentName *string

	// If true, Cutover will not be enabled if the document has failed.
	MustSucceedForCutover *bool

	// AWS Systems Manager Document parameters.
	Parameters map[string][]SsmParameterStoreParameter

	// AWS Systems Manager Document timeout seconds.
	TimeoutSeconds int32

	noSmithyDocumentSerde
}

// AWS Systems Manager Parameter Store parameter.
type SsmParameterStoreParameter struct {

	// AWS Systems Manager Parameter Store parameter name.
	//
	// This member is required.
	ParameterName *string

	// AWS Systems Manager Parameter Store parameter type.
	//
	// This member is required.
	ParameterType SsmParameterStoreParameterType

	noSmithyDocumentSerde
}

type TemplateActionDocument struct {

	// Template post migration custom action ID.
	ActionID *string

	// Template post migration custom action name.
	ActionName *string

	// Template post migration custom action active status.
	Active *bool

	// Template post migration custom action document identifier.
	DocumentIdentifier *string

	// Template post migration custom action document version.
	DocumentVersion *string

	// Template post migration custom action must succeed for cutover.
	MustSucceedForCutover *bool

	// Operating system eligible for this template post migration custom action.
	OperatingSystem *string

	// Template post migration custom action order.
	Order int32

	// Template post migration custom action parameters.
	Parameters map[string][]SsmParameterStoreParameter

	// Template post migration custom action timeout in seconds.
	TimeoutSeconds int32

	noSmithyDocumentSerde
}

// Template post migration custom action filters.
type TemplateActionsRequestFilters struct {

	// Action IDs to filter template post migration custom actions by.
	ActionIDs []string

	noSmithyDocumentSerde
}

// Validate exception field.
type ValidationExceptionField struct {

	// Validate exception field message.
	Message *string

	// Validate exception field name.
	Name *string

	noSmithyDocumentSerde
}

// vCenter client.
type VcenterClient struct {

	// Arn of vCenter client.
	Arn *string

	// Datacenter name of vCenter client.
	DatacenterName *string

	// Hostname of vCenter client .
	Hostname *string

	// Last seen time of vCenter client.
	LastSeenDatetime *string

	// Tags for Source Server of vCenter client.
	SourceServerTags map[string]string

	// Tags for vCenter client.
	Tags map[string]string

	// ID of vCenter client.
	VcenterClientID *string

	// Vcenter UUID of vCenter client.
	VcenterUUID *string

	noSmithyDocumentSerde
}

type Wave struct {

	// Wave ARN.
	Arn *string

	// Wave creation dateTime.
	CreationDateTime *string

	// Wave description.
	Description *string

	// Wave archival status.
	IsArchived *bool

	// Wave last modified dateTime.
	LastModifiedDateTime *string

	// Wave name.
	Name *string

	// Wave tags.
	Tags map[string]string

	// Wave aggregated status.
	WaveAggregatedStatus *WaveAggregatedStatus

	// Wave ID.
	WaveID *string

	noSmithyDocumentSerde
}

// Wave aggregated status.
type WaveAggregatedStatus struct {

	// Wave aggregated status health status.
	HealthStatus WaveHealthStatus

	// Wave aggregated status last update dateTime.
	LastUpdateDateTime *string

	// Wave aggregated status progress status.
	ProgressStatus WaveProgressStatus

	// DateTime marking when the first source server in the wave started replication.
	ReplicationStartedDateTime *string

	// Wave aggregated status total applications amount.
	TotalApplications int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
