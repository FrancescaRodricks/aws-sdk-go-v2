// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssetType string

// Enum values for AssetType
const (
	AssetTypeS3Snapshot        AssetType = "S3_SNAPSHOT"
	AssetTypeRedshiftDataShare AssetType = "REDSHIFT_DATA_SHARE"
	AssetTypeApiGatewayApi     AssetType = "API_GATEWAY_API"
)

// Values returns all known values for AssetType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AssetType) Values() []AssetType {
	return []AssetType{
		"S3_SNAPSHOT",
		"REDSHIFT_DATA_SHARE",
		"API_GATEWAY_API",
	}
}

type Code string

// Enum values for Code
const (
	CodeAccessDeniedException         Code = "ACCESS_DENIED_EXCEPTION"
	CodeInternalServerException       Code = "INTERNAL_SERVER_EXCEPTION"
	CodeMalwareDetected               Code = "MALWARE_DETECTED"
	CodeResourceNotFoundException     Code = "RESOURCE_NOT_FOUND_EXCEPTION"
	CodeServiceQuotaExceededException Code = "SERVICE_QUOTA_EXCEEDED_EXCEPTION"
	CodeValidationException           Code = "VALIDATION_EXCEPTION"
	CodeMalwareScanEncryptedFile      Code = "MALWARE_SCAN_ENCRYPTED_FILE"
)

// Values returns all known values for Code. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Code) Values() []Code {
	return []Code{
		"ACCESS_DENIED_EXCEPTION",
		"INTERNAL_SERVER_EXCEPTION",
		"MALWARE_DETECTED",
		"RESOURCE_NOT_FOUND_EXCEPTION",
		"SERVICE_QUOTA_EXCEEDED_EXCEPTION",
		"VALIDATION_EXCEPTION",
		"MALWARE_SCAN_ENCRYPTED_FILE",
	}
}

type ExceptionCause string

// Enum values for ExceptionCause
const (
	ExceptionCauseInsufficientS3BucketPolicy ExceptionCause = "InsufficientS3BucketPolicy"
	ExceptionCauseS3AccessDenied             ExceptionCause = "S3AccessDenied"
)

// Values returns all known values for ExceptionCause. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExceptionCause) Values() []ExceptionCause {
	return []ExceptionCause{
		"InsufficientS3BucketPolicy",
		"S3AccessDenied",
	}
}

type JobErrorLimitName string

// Enum values for JobErrorLimitName
const (
	JobErrorLimitNameAssetsPerRevision                        JobErrorLimitName = "Assets per revision"
	JobErrorLimitNameAssetSizeInGb                            JobErrorLimitName = "Asset size in GB"
	JobErrorLimitNameAmazonRedshiftDatashareAssetsPerRevision JobErrorLimitName = "Amazon Redshift datashare assets per revision"
)

// Values returns all known values for JobErrorLimitName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobErrorLimitName) Values() []JobErrorLimitName {
	return []JobErrorLimitName{
		"Assets per revision",
		"Asset size in GB",
		"Amazon Redshift datashare assets per revision",
	}
}

type JobErrorResourceTypes string

// Enum values for JobErrorResourceTypes
const (
	JobErrorResourceTypesRevision JobErrorResourceTypes = "REVISION"
	JobErrorResourceTypesAsset    JobErrorResourceTypes = "ASSET"
	JobErrorResourceTypesDataSet  JobErrorResourceTypes = "DATA_SET"
)

// Values returns all known values for JobErrorResourceTypes. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobErrorResourceTypes) Values() []JobErrorResourceTypes {
	return []JobErrorResourceTypes{
		"REVISION",
		"ASSET",
		"DATA_SET",
	}
}

type LimitName string

// Enum values for LimitName
const (
	LimitNameProductsPerAccount                                                 LimitName = "Products per account"
	LimitNameDataSetsPerAccount                                                 LimitName = "Data sets per account"
	LimitNameDataSetsPerProduct                                                 LimitName = "Data sets per product"
	LimitNameRevisionsPerDataSet                                                LimitName = "Revisions per data set"
	LimitNameAssetsPerRevision                                                  LimitName = "Assets per revision"
	LimitNameAssetsPerImportJobFromAmazonS3                                     LimitName = "Assets per import job from Amazon S3"
	LimitNameAssetPerExportJobFromAmazonS3                                      LimitName = "Asset per export job from Amazon S3"
	LimitNameAssetSizeInGb                                                      LimitName = "Asset size in GB"
	LimitNameConcurrentInProgressJobsToExportAssetsToAmazonS3                   LimitName = "Concurrent in progress jobs to export assets to Amazon S3"
	LimitNameConcurrentInProgressJobsToExportAssetsToASignedUrl                 LimitName = "Concurrent in progress jobs to export assets to a signed URL"
	LimitNameConcurrentInProgressJobsToImportAssetsFromAmazonS3                 LimitName = "Concurrent in progress jobs to import assets from Amazon S3"
	LimitNameConcurrentInProgressJobsToImportAssetsFromASignedUrl               LimitName = "Concurrent in progress jobs to import assets from a signed URL"
	LimitNameConcurrentInProgressJobsToExportRevisionsToAmazonS3                LimitName = "Concurrent in progress jobs to export revisions to Amazon S3"
	LimitNameEventActionsPerAccount                                             LimitName = "Event actions per account"
	LimitNameAutoExportEventActionsPerDataSet                                   LimitName = "Auto export event actions per data set"
	LimitNameAmazonRedshiftDatashareAssetsPerImportJobFromRedshift              LimitName = "Amazon Redshift datashare assets per import job from Redshift"
	LimitNameConcurrentInProgressJobsToImportAssetsFromAmazonRedshiftDatashares LimitName = "Concurrent in progress jobs to import assets from Amazon Redshift datashares"
	LimitNameRevisionsPerAmazonRedshiftDatashareDataSet                         LimitName = "Revisions per Amazon Redshift datashare data set"
	LimitNameAmazonRedshiftDatashareAssetsPerRevision                           LimitName = "Amazon Redshift datashare assets per revision"
	LimitNameConcurrentInProgressJobsToImportAssetsFromAnApiGatewayApi          LimitName = "Concurrent in progress jobs to import assets from an API Gateway API"
	LimitNameAmazonApiGatewayApiAssetsPerRevision                               LimitName = "Amazon API Gateway API assets per revision"
	LimitNameRevisionsPerAmazonApiGatewayApiDataSet                             LimitName = "Revisions per Amazon API Gateway API data set"
)

// Values returns all known values for LimitName. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LimitName) Values() []LimitName {
	return []LimitName{
		"Products per account",
		"Data sets per account",
		"Data sets per product",
		"Revisions per data set",
		"Assets per revision",
		"Assets per import job from Amazon S3",
		"Asset per export job from Amazon S3",
		"Asset size in GB",
		"Concurrent in progress jobs to export assets to Amazon S3",
		"Concurrent in progress jobs to export assets to a signed URL",
		"Concurrent in progress jobs to import assets from Amazon S3",
		"Concurrent in progress jobs to import assets from a signed URL",
		"Concurrent in progress jobs to export revisions to Amazon S3",
		"Event actions per account",
		"Auto export event actions per data set",
		"Amazon Redshift datashare assets per import job from Redshift",
		"Concurrent in progress jobs to import assets from Amazon Redshift datashares",
		"Revisions per Amazon Redshift datashare data set",
		"Amazon Redshift datashare assets per revision",
		"Concurrent in progress jobs to import assets from an API Gateway API",
		"Amazon API Gateway API assets per revision",
		"Revisions per Amazon API Gateway API data set",
	}
}

type Origin string

// Enum values for Origin
const (
	OriginOwned    Origin = "OWNED"
	OriginEntitled Origin = "ENTITLED"
)

// Values returns all known values for Origin. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Origin) Values() []Origin {
	return []Origin{
		"OWNED",
		"ENTITLED",
	}
}

type ProtocolType string

// Enum values for ProtocolType
const (
	ProtocolTypeRest ProtocolType = "REST"
)

// Values returns all known values for ProtocolType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ProtocolType) Values() []ProtocolType {
	return []ProtocolType{
		"REST",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeDataSet     ResourceType = "DATA_SET"
	ResourceTypeRevision    ResourceType = "REVISION"
	ResourceTypeAsset       ResourceType = "ASSET"
	ResourceTypeJob         ResourceType = "JOB"
	ResourceTypeEventAction ResourceType = "EVENT_ACTION"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"DATA_SET",
		"REVISION",
		"ASSET",
		"JOB",
		"EVENT_ACTION",
	}
}

type ServerSideEncryptionTypes string

// Enum values for ServerSideEncryptionTypes
const (
	ServerSideEncryptionTypesAwsKms ServerSideEncryptionTypes = "aws:kms"
	ServerSideEncryptionTypesAes256 ServerSideEncryptionTypes = "AES256"
)

// Values returns all known values for ServerSideEncryptionTypes. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServerSideEncryptionTypes) Values() []ServerSideEncryptionTypes {
	return []ServerSideEncryptionTypes{
		"aws:kms",
		"AES256",
	}
}

type State string

// Enum values for State
const (
	StateWaiting    State = "WAITING"
	StateInProgress State = "IN_PROGRESS"
	StateError      State = "ERROR"
	StateCompleted  State = "COMPLETED"
	StateCancelled  State = "CANCELLED"
	StateTimedOut   State = "TIMED_OUT"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"WAITING",
		"IN_PROGRESS",
		"ERROR",
		"COMPLETED",
		"CANCELLED",
		"TIMED_OUT",
	}
}

type Type string

// Enum values for Type
const (
	TypeImportAssetsFromS3                 Type = "IMPORT_ASSETS_FROM_S3"
	TypeImportAssetFromSignedUrl           Type = "IMPORT_ASSET_FROM_SIGNED_URL"
	TypeExportAssetsToS3                   Type = "EXPORT_ASSETS_TO_S3"
	TypeExportAssetToSignedUrl             Type = "EXPORT_ASSET_TO_SIGNED_URL"
	TypeExportRevisionsToS3                Type = "EXPORT_REVISIONS_TO_S3"
	TypeImportAssetsFromRedshiftDataShares Type = "IMPORT_ASSETS_FROM_REDSHIFT_DATA_SHARES"
	TypeImportAssetFromApiGatewayApi       Type = "IMPORT_ASSET_FROM_API_GATEWAY_API"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"IMPORT_ASSETS_FROM_S3",
		"IMPORT_ASSET_FROM_SIGNED_URL",
		"EXPORT_ASSETS_TO_S3",
		"EXPORT_ASSET_TO_SIGNED_URL",
		"EXPORT_REVISIONS_TO_S3",
		"IMPORT_ASSETS_FROM_REDSHIFT_DATA_SHARES",
		"IMPORT_ASSET_FROM_API_GATEWAY_API",
	}
}