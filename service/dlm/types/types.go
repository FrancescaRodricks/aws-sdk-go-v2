// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// [Event-based policies only] Specifies an action for an event-based policy.
type Action struct {

	// The rule for copying shared snapshots across Regions.
	//
	// This member is required.
	CrossRegionCopy []CrossRegionCopyAction

	// A descriptive name for the action.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

//	[Custom snapshot policies only] Specifies information about the archive
//
// storage tier retention period.
type ArchiveRetainRule struct {

	// Information about retention period in the Amazon EBS Snapshots Archive. For
	// more information, see [Archive Amazon EBS snapshots].
	//
	// [Archive Amazon EBS snapshots]: https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/snapshot-archive.html
	//
	// This member is required.
	RetentionArchiveTier *RetentionArchiveTier

	noSmithyDocumentSerde
}

//	[Custom snapshot policies only] Specifies a snapshot archiving rule for a
//
// schedule.
type ArchiveRule struct {

	// Information about the retention period for the snapshot archiving rule.
	//
	// This member is required.
	RetainRule *ArchiveRetainRule

	noSmithyDocumentSerde
}

//	[Custom snapshot and AMI policies only] Specifies when the policy should
//
// create snapshots or AMIs.
//
//   - You must specify either CronExpression, or Interval, IntervalUnit, and
//     Times.
//
//   - If you need to specify an [ArchiveRule]for the schedule, then you must specify a
//     creation frequency of at least 28 days.
//
// [ArchiveRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_ArchiveRule.html
type CreateRule struct {

	// The schedule, as a Cron expression. The schedule interval must be between 1
	// hour and 1 year. For more information, see [Cron expressions]in the Amazon CloudWatch User Guide.
	//
	// [Cron expressions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
	CronExpression *string

	// The interval between snapshots. The supported values are 1, 2, 3, 4, 6, 8, 12,
	// and 24.
	Interval *int32

	// The interval unit.
	IntervalUnit IntervalUnitValues

	//  [Custom snapshot policies only] Specifies the destination for snapshots
	// created by the policy. To create snapshots in the same Region as the source
	// resource, specify CLOUD . To create snapshots on the same Outpost as the source
	// resource, specify OUTPOST_LOCAL . If you omit this parameter, CLOUD is used by
	// default.
	//
	// If the policy targets resources in an Amazon Web Services Region, then you must
	// create snapshots in the same Region as the source resource. If the policy
	// targets resources on an Outpost, then you can create snapshots on the same
	// Outpost as the source resource, or in the Region of that Outpost.
	Location LocationValues

	//  [Custom snapshot policies that target instances only] Specifies pre and/or
	// post scripts for a snapshot lifecycle policy that targets instances. This is
	// useful for creating application-consistent snapshots, or for performing specific
	// administrative tasks before or after Amazon Data Lifecycle Manager initiates
	// snapshot creation.
	//
	// For more information, see [Automating application-consistent snapshots with pre and post scripts].
	//
	// [Automating application-consistent snapshots with pre and post scripts]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/automate-app-consistent-backups.html
	Scripts []Script

	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If
	// you do not specify a time, Amazon Data Lifecycle Manager selects a time within
	// the next 24 hours.
	Times []string

	noSmithyDocumentSerde
}

//	[Event-based policies only] Specifies a cross-Region copy action for
//
// event-based policies.
//
// To specify a cross-Region copy rule for snapshot and AMI policies, use [CrossRegionCopyRule].
//
// [CrossRegionCopyRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_CrossRegionCopyRule.html
type CrossRegionCopyAction struct {

	// The encryption settings for the copied snapshot.
	//
	// This member is required.
	EncryptionConfiguration *EncryptionConfiguration

	// The target Region.
	//
	// This member is required.
	Target *string

	// Specifies a retention rule for cross-Region snapshot copies created by snapshot
	// or event-based policies, or cross-Region AMI copies created by AMI policies.
	// After the retention period expires, the cross-Region copy is deleted.
	RetainRule *CrossRegionCopyRetainRule

	noSmithyDocumentSerde
}

//	[Custom AMI policies only] Specifies an AMI deprecation rule for cross-Region
//
// AMI copies created by an AMI policy.
type CrossRegionCopyDeprecateRule struct {

	// The period after which to deprecate the cross-Region AMI copies. The period
	// must be less than or equal to the cross-Region AMI copy retention period, and it
	// can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or
	// 3650 days.
	Interval *int32

	// The unit of time in which to measure the Interval. For example, to deprecate a
	// cross-Region AMI copy after 3 months, specify Interval=3 and IntervalUnit=MONTHS
	// .
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

// Specifies a retention rule for cross-Region snapshot copies created by snapshot
// or event-based policies, or cross-Region AMI copies created by AMI policies.
// After the retention period expires, the cross-Region copy is deleted.
type CrossRegionCopyRetainRule struct {

	// The amount of time to retain a cross-Region snapshot or AMI copy. The maximum
	// is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int32

	// The unit of time for time-based retention. For example, to retain a
	// cross-Region copy for 3 months, specify Interval=3 and IntervalUnit=MONTHS .
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

//	[Custom snapshot and AMI policies only] Specifies a cross-Region copy rule for
//
// a snapshot and AMI policies.
//
// To specify a cross-Region copy action for event-based polices, use [CrossRegionCopyAction].
//
// [CrossRegionCopyAction]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_CrossRegionCopyAction.html
type CrossRegionCopyRule struct {

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not
	// enabled, enable encryption using this parameter. Copies of encrypted snapshots
	// are encrypted, even if this parameter is false or if encryption by default is
	// not enabled.
	//
	// This member is required.
	Encrypted *bool

	// The Amazon Resource Name (ARN) of the KMS key to use for EBS encryption. If
	// this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string

	// Indicates whether to copy all user-defined tags from the source snapshot or AMI
	// to the cross-Region copy.
	CopyTags *bool

	//  [Custom AMI policies only] The AMI deprecation rule for cross-Region AMI
	// copies created by the rule.
	DeprecateRule *CrossRegionCopyDeprecateRule

	// The retention rule that indicates how long the cross-Region snapshot or AMI
	// copies are to be retained in the destination Region.
	RetainRule *CrossRegionCopyRetainRule

	// Use this parameter for snapshot policies only. For AMI policies, use
	// TargetRegion instead.
	//
	// [Custom snapshot policies only] The target Region or the Amazon Resource Name
	// (ARN) of the target Outpost for the snapshot copies.
	Target *string

	// Use this parameter for AMI policies only. For snapshot policies, use Target
	// instead. For snapshot policies created before the Target parameter was
	// introduced, this parameter indicates the target Region for snapshot copies.
	//
	// [Custom AMI policies only] The target Region or the Amazon Resource Name (ARN)
	// of the target Outpost for the snapshot copies.
	TargetRegion *string

	noSmithyDocumentSerde
}

//	[Default policies only] Specifies a destination Region for cross-Region copy
//
// actions.
type CrossRegionCopyTarget struct {

	// The target Region, for example us-east-1 .
	TargetRegion *string

	noSmithyDocumentSerde
}

//	[Custom AMI policies only] Specifies an AMI deprecation rule for AMIs created
//
// by an AMI lifecycle policy.
//
// For age-based schedules, you must specify Interval and IntervalUnit. For
// count-based schedules, you must specify Count.
type DeprecateRule struct {

	// If the schedule has a count-based retention rule, this parameter specifies the
	// number of oldest AMIs to deprecate. The count must be less than or equal to the
	// schedule's retention count, and it can't be greater than 1000.
	Count *int32

	// If the schedule has an age-based retention rule, this parameter specifies the
	// period after which to deprecate AMIs created by the schedule. The period must be
	// less than or equal to the schedule's retention period, and it can't be greater
	// than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *int32

	// The unit of time in which to measure the Interval.
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

//	[Event-based policies only] Specifies the encryption settings for cross-Region
//
// snapshot copies created by event-based policies.
type EncryptionConfiguration struct {

	// To encrypt a copy of an unencrypted snapshot when encryption by default is not
	// enabled, enable encryption using this parameter. Copies of encrypted snapshots
	// are encrypted, even if this parameter is false or when encryption by default is
	// not enabled.
	//
	// This member is required.
	Encrypted *bool

	// The Amazon Resource Name (ARN) of the KMS key to use for EBS encryption. If
	// this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string

	noSmithyDocumentSerde
}

//	[Event-based policies only] Specifies an event that activates an event-based
//
// policy.
type EventParameters struct {

	// The snapshot description that can trigger the policy. The description pattern
	// is specified using a regular expression. The policy runs only if a snapshot with
	// a description that matches the specified pattern is shared with your account.
	//
	// For example, specifying ^.*Created for policy: policy-1234567890abcdef0.*$
	// configures the policy to run only if snapshots created by policy
	// policy-1234567890abcdef0 are shared with your account.
	//
	// This member is required.
	DescriptionRegex *string

	// The type of event. Currently, only snapshot sharing events are supported.
	//
	// This member is required.
	EventType EventTypeValues

	// The IDs of the Amazon Web Services accounts that can trigger policy by sharing
	// snapshots with your account. The policy only runs if one of the specified Amazon
	// Web Services accounts shares a snapshot with your account.
	//
	// This member is required.
	SnapshotOwner []string

	noSmithyDocumentSerde
}

//	[Event-based policies only] Specifies an event that activates an event-based
//
// policy.
type EventSource struct {

	// The source of the event. Currently only managed CloudWatch Events rules are
	// supported.
	//
	// This member is required.
	Type EventSourceValues

	// Information about the event.
	Parameters *EventParameters

	noSmithyDocumentSerde
}

//	[Default policies only] Specifies exclusion parameters for volumes or
//
// instances for which you do not want to create snapshots or AMIs. The policy will
// not create snapshots or AMIs for target resources that match any of the
// specified exclusion parameters.
type Exclusions struct {

	//  [Default policies for EBS snapshots only] Indicates whether to exclude volumes
	// that are attached to instances as the boot volume. If you exclude boot volumes,
	// only volumes attached as data (non-boot) volumes will be backed up by the
	// policy. To exclude boot volumes, specify true .
	ExcludeBootVolumes *bool

	//  [Default policies for EBS-backed AMIs only] Specifies whether to exclude
	// volumes that have specific tags.
	ExcludeTags []Tag

	//  [Default policies for EBS snapshots only] Specifies the volume types to
	// exclude. Volumes of the specified types will not be targeted by the policy.
	ExcludeVolumeTypes []string

	noSmithyDocumentSerde
}

//	[Custom snapshot policies only] Specifies a rule for enabling fast snapshot
//
// restore for snapshots created by snapshot policies. You can enable fast snapshot
// restore based on either a count or a time interval.
type FastRestoreRule struct {

	// The Availability Zones in which to enable fast snapshot restore.
	//
	// This member is required.
	AvailabilityZones []string

	// The number of snapshots to be enabled with fast snapshot restore.
	Count *int32

	// The amount of time to enable fast snapshot restore. The maximum is 100 years.
	// This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int32

	// The unit of time for enabling fast snapshot restore.
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

//	[Custom policies only] Detailed information about a snapshot, AMI, or
//
// event-based lifecycle policy.
type LifecyclePolicy struct {

	// The local date and time when the lifecycle policy was created.
	DateCreated *time.Time

	// The local date and time when the lifecycle policy was last modified.
	DateModified *time.Time

	//  [Default policies only] The type of default policy. Values include:
	//
	//   - VOLUME - Default policy for EBS snapshots
	//
	//   - INSTANCE - Default policy for EBS-backed AMIs
	DefaultPolicy *bool

	// The description of the lifecycle policy.
	Description *string

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string

	// The configuration of the lifecycle policy
	PolicyDetails *PolicyDetails

	// The identifier of the lifecycle policy.
	PolicyId *string

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues

	// The description of the status.
	StatusMessage *string

	// The tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Summary information about a lifecycle policy.
type LifecyclePolicySummary struct {

	//  [Default policies only] The type of default policy. Values include:
	//
	//   - VOLUME - Default policy for EBS snapshots
	//
	//   - INSTANCE - Default policy for EBS-backed AMIs
	DefaultPolicy *bool

	// The description of the lifecycle policy.
	Description *string

	// The identifier of the lifecycle policy.
	PolicyId *string

	// The type of policy. EBS_SNAPSHOT_MANAGEMENT indicates that the policy manages
	// the lifecycle of Amazon EBS snapshots. IMAGE_MANAGEMENT indicates that the
	// policy manages the lifecycle of EBS-backed AMIs. EVENT_BASED_POLICY indicates
	// that the policy automates cross-account snapshot copies for snapshots that are
	// shared with your account.
	PolicyType PolicyTypeValues

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues

	// The tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

//	[Custom snapshot and AMI policies only] Specifies optional parameters for
//
// snapshot and AMI policies. The set of valid parameters depends on the
// combination of policy type and target resource type.
//
// If you choose to exclude boot volumes and you specify tags that consequently
// exclude all of the additional data volumes attached to an instance, then Amazon
// Data Lifecycle Manager will not create any snapshots for the affected instance,
// and it will emit a SnapshotsCreateFailed Amazon CloudWatch metric. For more
// information, see [Monitor your policies using Amazon CloudWatch].
//
// [Monitor your policies using Amazon CloudWatch]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitor-dlm-cw-metrics.html
type Parameters struct {

	//  [Custom snapshot policies that target instances only] Indicates whether to
	// exclude the root volume from multi-volume snapshot sets. The default is false .
	// If you specify true , then the root volumes attached to targeted instances will
	// be excluded from the multi-volume snapshot sets created by the policy.
	ExcludeBootVolume *bool

	//  [Custom snapshot policies that target instances only] The tags used to
	// identify data (non-root) volumes to exclude from multi-volume snapshot sets.
	//
	// If you create a snapshot lifecycle policy that targets instances and you
	// specify tags for this parameter, then data volumes with the specified tags that
	// are attached to targeted instances will be excluded from the multi-volume
	// snapshot sets created by the policy.
	ExcludeDataVolumeTags []Tag

	//  [Custom AMI policies only] Indicates whether targeted instances are rebooted
	// when the lifecycle policy runs. true indicates that targeted instances are not
	// rebooted when the policy runs. false indicates that target instances are
	// rebooted when the policy runs. The default is true (instances are not rebooted).
	NoReboot *bool

	noSmithyDocumentSerde
}

// Specifies the configuration of a lifecycle policy.
type PolicyDetails struct {

	//  [Event-based policies only] The actions to be performed when the event-based
	// policy is activated. You can specify only one action per policy.
	Actions []Action

	//  [Default policies only] Indicates whether the policy should copy tags from the
	// source resource to the snapshot or AMI. If you do not specify a value, the
	// default is false .
	//
	// Default: false
	CopyTags *bool

	//  [Default policies only] Specifies how often the policy should run and create
	// snapshots or AMIs. The creation frequency can range from 1 to 7 days. If you do
	// not specify a value, the default is 1.
	//
	// Default: 1
	CreateInterval *int32

	//  [Default policies only] Specifies destination Regions for snapshot or AMI
	// copies. You can specify up to 3 destination Regions. If you do not want to
	// create cross-Region copies, omit this parameter.
	CrossRegionCopyTargets []CrossRegionCopyTarget

	//  [Event-based policies only] The event that activates the event-based policy.
	EventSource *EventSource

	//  [Default policies only] Specifies exclusion parameters for volumes or
	// instances for which you do not want to create snapshots or AMIs. The policy will
	// not create snapshots or AMIs for target resources that match any of the
	// specified exclusion parameters.
	Exclusions *Exclusions

	//  [Default policies only] Defines the snapshot or AMI retention behavior for the
	// policy if the source volume or instance is deleted, or if the policy enters the
	// error, disabled, or deleted state.
	//
	// By default (ExtendDeletion=false):
	//
	//   - If a source resource is deleted, Amazon Data Lifecycle Manager will
	//   continue to delete previously created snapshots or AMIs, up to but not including
	//   the last one, based on the specified retention period. If you want Amazon Data
	//   Lifecycle Manager to delete all snapshots or AMIs, including the last one,
	//   specify true .
	//
	//   - If a policy enters the error, disabled, or deleted state, Amazon Data
	//   Lifecycle Manager stops deleting snapshots and AMIs. If you want Amazon Data
	//   Lifecycle Manager to continue deleting snapshots or AMIs, including the last
	//   one, if the policy enters one of these states, specify true .
	//
	// If you enable extended deletion (ExtendDeletion=true), you override both
	// default behaviors simultaneously.
	//
	// If you do not specify a value, the default is false .
	//
	// Default: false
	ExtendDeletion *bool

	//  [Custom snapshot and AMI policies only] A set of optional parameters for
	// snapshot and AMI lifecycle policies.
	//
	// If you are modifying a policy that was created or previously modified using the
	// Amazon Data Lifecycle Manager console, then you must include this parameter and
	// specify either the default values or the new values that you require. You can't
	// omit this parameter or set its values to null.
	Parameters *Parameters

	// The type of policy to create. Specify one of the following:
	//
	//   - SIMPLIFIED To create a default policy.
	//
	//   - STANDARD To create a custom policy.
	PolicyLanguage PolicyLanguageValues

	//  [Custom policies only] The valid target resource types and actions a policy
	// can manage. Specify EBS_SNAPSHOT_MANAGEMENT to create a lifecycle policy that
	// manages the lifecycle of Amazon EBS snapshots. Specify IMAGE_MANAGEMENT to
	// create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify
	// EVENT_BASED_POLICY  to create an event-based policy that performs specific
	// actions when a defined event occurs in your Amazon Web Services account.
	//
	// The default is EBS_SNAPSHOT_MANAGEMENT .
	PolicyType PolicyTypeValues

	//  [Custom snapshot and AMI policies only] The location of the resources to
	// backup. If the source resources are located in an Amazon Web Services Region,
	// specify CLOUD . If the source resources are located on an Outpost in your
	// account, specify OUTPOST .
	//
	// If you specify OUTPOST , Amazon Data Lifecycle Manager backs up all resources of
	// the specified type with matching target tags across all of the Outposts in your
	// account.
	ResourceLocations []ResourceLocationValues

	//  [Default policies only] Specify the type of default policy to create.
	//
	//   - To create a default policy for EBS snapshots, that creates snapshots of all
	//   volumes in the Region that do not have recent backups, specify VOLUME .
	//
	//   - To create a default policy for EBS-backed AMIs, that creates EBS-backed
	//   AMIs from all instances in the Region that do not have recent backups, specify
	//   INSTANCE .
	ResourceType ResourceTypeValues

	//  [Custom snapshot policies only] The target resource type for snapshot and AMI
	// lifecycle policies. Use VOLUME to create snapshots of individual volumes or use
	// INSTANCE to create multi-volume snapshots from the volumes for an instance.
	ResourceTypes []ResourceTypeValues

	//  [Default policies only] Specifies how long the policy should retain snapshots
	// or AMIs before deleting them. The retention period can range from 2 to 14 days,
	// but it must be greater than the creation frequency to ensure that the policy
	// retains at least 1 snapshot or AMI at any given time. If you do not specify a
	// value, the default is 7.
	//
	// Default: 7
	RetainInterval *int32

	//  [Custom snapshot and AMI policies only] The schedules of policy-defined
	// actions for snapshot and AMI lifecycle policies. A policy can have up to four
	// schedules—one mandatory schedule and up to three optional schedules.
	Schedules []Schedule

	//  [Custom snapshot and AMI policies only] The single tag that identifies
	// targeted resources for this policy.
	TargetTags []Tag

	noSmithyDocumentSerde
}

//	[Custom snapshot and AMI policies only] Specifies a retention rule for
//
// snapshots created by snapshot policies, or for AMIs created by AMI policies.
//
// For snapshot policies that have an [ArchiveRule], this retention rule applies to standard
// tier retention. When the retention threshold is met, snapshots are moved from
// the standard to the archive tier.
//
// For snapshot policies that do not have an ArchiveRule, snapshots are
// permanently deleted when this retention threshold is met.
//
// You can retain snapshots based on either a count or a time interval.
//
//   - Count-based retention
//
// You must specify Count. If you specify an [ArchiveRule]for the schedule, then you can
//
//	specify a retention count of 0 to archive snapshots immediately after
//	creation. If you specify a [FastRestoreRule], [ShareRule], or a [CrossRegionCopyRule], then you must specify a retention count
//	of 1 or more.
//
//	- Age-based retention
//
// You must specify Interval and IntervalUnit. If you specify an [ArchiveRule]for the schedule,
//
//	then you can specify a retention interval of 0 days to archive snapshots
//	immediately after creation. If you specify a [FastRestoreRule], [ShareRule], or a [CrossRegionCopyRule], then you must specify
//	a retention interval of 1 day or more.
//
// [ArchiveRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_ArchiveRule.html
//
// [CrossRegionCopyRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_CrossRegionCopyRule.html
// [FastRestoreRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_FastRestoreRule.html
// [ShareRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_ShareRule.html
type RetainRule struct {

	// The number of snapshots to retain for each volume, up to a maximum of 1000. For
	// example if you want to retain a maximum of three snapshots, specify 3 . When the
	// fourth snapshot is created, the oldest retained snapshot is deleted, or it is
	// moved to the archive tier if you have specified an [ArchiveRule].
	//
	// [ArchiveRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_ArchiveRule.html
	Count *int32

	// The amount of time to retain each snapshot. The maximum is 100 years. This is
	// equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *int32

	// The unit of time for time-based retention. For example, to retain snapshots for
	// 3 months, specify Interval=3 and IntervalUnit=MONTHS . Once the snapshot has
	// been retained for 3 months, it is deleted, or it is moved to the archive tier if
	// you have specified an [ArchiveRule].
	//
	// [ArchiveRule]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_ArchiveRule.html
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

//	[Custom snapshot policies only] Describes the retention rule for archived
//
// snapshots. Once the archive retention threshold is met, the snapshots are
// permanently deleted from the archive tier.
//
// The archive retention rule must retain snapshots in the archive tier for a
// minimum of 90 days.
//
// For count-based schedules, you must specify Count. For age-based schedules, you
// must specify Interval and IntervalUnit.
//
// For more information about using snapshot archiving, see [Considerations for snapshot lifecycle policies].
//
// [Considerations for snapshot lifecycle policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-ami-policy.html#dlm-archive
type RetentionArchiveTier struct {

	// The maximum number of snapshots to retain in the archive storage tier for each
	// volume. The count must ensure that each snapshot remains in the archive tier for
	// at least 90 days. For example, if the schedule creates snapshots every 30 days,
	// you must specify a count of 3 or more to ensure that each snapshot is archived
	// for at least 90 days.
	Count *int32

	// Specifies the period of time to retain snapshots in the archive tier. After
	// this period expires, the snapshot is permanently deleted.
	Interval *int32

	// The unit of time in which to measure the Interval. For example, to retain a
	// snapshots in the archive tier for 6 months, specify Interval=6 and
	// IntervalUnit=MONTHS .
	IntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

//	[Custom snapshot and AMI policies only] Specifies a schedule for a snapshot or
//
// AMI lifecycle policy.
type Schedule struct {

	//  [Custom snapshot policies that target volumes only] The snapshot archiving
	// rule for the schedule. When you specify an archiving rule, snapshots are
	// automatically moved from the standard tier to the archive tier once the
	// schedule's retention threshold is met. Snapshots are then retained in the
	// archive tier for the archive retention period that you specify.
	//
	// For more information about using snapshot archiving, see [Considerations for snapshot lifecycle policies].
	//
	// [Considerations for snapshot lifecycle policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-ami-policy.html#dlm-archive
	ArchiveRule *ArchiveRule

	// Copy all user-defined tags on a source volume to snapshots of the volume
	// created by this policy.
	CopyTags *bool

	// The creation rule.
	CreateRule *CreateRule

	// Specifies a rule for copying snapshots or AMIs across regions.
	//
	// You can't specify cross-Region copy rules for policies that create snapshots on
	// an Outpost. If the policy creates snapshots in a Region, then snapshots can be
	// copied to up to three Regions or Outposts.
	CrossRegionCopyRules []CrossRegionCopyRule

	//  [Custom AMI policies only] The AMI deprecation rule for the schedule.
	DeprecateRule *DeprecateRule

	//  [Custom snapshot policies only] The rule for enabling fast snapshot restore.
	FastRestoreRule *FastRestoreRule

	// The name of the schedule.
	Name *string

	// The retention rule for snapshots or AMIs created by the policy.
	RetainRule *RetainRule

	//  [Custom snapshot policies only] The rule for sharing snapshots with other
	// Amazon Web Services accounts.
	ShareRules []ShareRule

	// The tags to apply to policy-created resources. These user-defined tags are in
	// addition to the Amazon Web Services-added lifecycle tags.
	TagsToAdd []Tag

	//  [AMI policies and snapshot policies that target instances only] A collection
	// of key/value pairs with values determined dynamically when the policy is
	// executed. Keys may be any valid Amazon EC2 tag key. Values must be in one of the
	// two following formats: $(instance-id) or $(timestamp) . Variable tags are only
	// valid for EBS Snapshot Management – Instance policies.
	VariableTags []Tag

	noSmithyDocumentSerde
}

//	[Custom snapshot policies that target instances only] Information about pre
//
// and/or post scripts for a snapshot lifecycle policy that targets instances. For
// more information, see [Automating application-consistent snapshots with pre and post scripts].
//
// [Automating application-consistent snapshots with pre and post scripts]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/automate-app-consistent-backups.html
type Script struct {

	// The SSM document that includes the pre and/or post scripts to run.
	//
	//   - If you are automating VSS backups, specify AWS_VSS_BACKUP . In this case,
	//   Amazon Data Lifecycle Manager automatically uses the AWSEC2-CreateVssSnapshot
	//   SSM document.
	//
	//   - If you are automating application-consistent snapshots for SAP HANA
	//   workloads, specify AWSSystemsManagerSAP-CreateDLMSnapshotForSAPHANA .
	//
	//   - If you are using a custom SSM document that you own, specify either the
	//   name or ARN of the SSM document. If you are using a custom SSM document that is
	//   shared with you, specify the ARN of the SSM document.
	//
	// This member is required.
	ExecutionHandler *string

	// Indicates whether Amazon Data Lifecycle Manager should default to
	// crash-consistent snapshots if the pre script fails.
	//
	//   - To default to crash consistent snapshot if the pre script fails, specify
	//   true .
	//
	//   - To skip the instance for snapshot creation if the pre script fails, specify
	//   false .
	//
	// This parameter is supported only if you run a pre script. If you run a post
	// script only, omit this parameter.
	//
	// Default: true
	ExecuteOperationOnScriptFailure *bool

	// Indicates the service used to execute the pre and/or post scripts.
	//
	//   - If you are using custom SSM documents or automating application-consistent
	//   snapshots of SAP HANA workloads, specify AWS_SYSTEMS_MANAGER .
	//
	//   - If you are automating VSS Backups, omit this parameter.
	//
	// Default: AWS_SYSTEMS_MANAGER
	ExecutionHandlerService ExecutionHandlerServiceValues

	// Specifies a timeout period, in seconds, after which Amazon Data Lifecycle
	// Manager fails the script run attempt if it has not completed. If a script does
	// not complete within its timeout period, Amazon Data Lifecycle Manager fails the
	// attempt. The timeout period applies to the pre and post scripts individually.
	//
	// If you are automating VSS Backups, omit this parameter.
	//
	// Default: 10
	ExecutionTimeout *int32

	// Specifies the number of times Amazon Data Lifecycle Manager should retry
	// scripts that fail.
	//
	//   - If the pre script fails, Amazon Data Lifecycle Manager retries the entire
	//   snapshot creation process, including running the pre and post scripts.
	//
	//   - If the post script fails, Amazon Data Lifecycle Manager retries the post
	//   script only; in this case, the pre script will have completed and the snapshot
	//   might have been created.
	//
	// If you do not want Amazon Data Lifecycle Manager to retry failed scripts,
	// specify 0 .
	//
	// Default: 0
	MaximumRetryCount *int32

	// Indicate which scripts Amazon Data Lifecycle Manager should run on target
	// instances. Pre scripts run before Amazon Data Lifecycle Manager initiates
	// snapshot creation. Post scripts run after Amazon Data Lifecycle Manager
	// initiates snapshot creation.
	//
	//   - To run a pre script only, specify PRE . In this case, Amazon Data Lifecycle
	//   Manager calls the SSM document with the pre-script parameter before initiating
	//   snapshot creation.
	//
	//   - To run a post script only, specify POST . In this case, Amazon Data
	//   Lifecycle Manager calls the SSM document with the post-script parameter after
	//   initiating snapshot creation.
	//
	//   - To run both pre and post scripts, specify both PRE and POST . In this case,
	//   Amazon Data Lifecycle Manager calls the SSM document with the pre-script
	//   parameter before initiating snapshot creation, and then it calls the SSM
	//   document again with the post-script parameter after initiating snapshot
	//   creation.
	//
	// If you are automating VSS Backups, omit this parameter.
	//
	// Default: PRE and POST
	Stages []StageValues

	noSmithyDocumentSerde
}

//	[Custom snapshot policies only] Specifies a rule for sharing snapshots across
//
// Amazon Web Services accounts.
type ShareRule struct {

	// The IDs of the Amazon Web Services accounts with which to share the snapshots.
	//
	// This member is required.
	TargetAccounts []string

	// The period after which snapshots that are shared with other Amazon Web Services
	// accounts are automatically unshared.
	UnshareInterval *int32

	// The unit of time for the automatic unsharing interval.
	UnshareIntervalUnit RetentionIntervalUnitValues

	noSmithyDocumentSerde
}

// Specifies a tag for a resource.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
