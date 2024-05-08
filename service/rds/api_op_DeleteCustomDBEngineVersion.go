// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a custom engine version. To run this command, make sure you meet the
// following prerequisites:
//
//   - The CEV must not be the default for RDS Custom. If it is, change the
//     default before running this command.
//
//   - The CEV must not be associated with an RDS Custom DB instance, RDS Custom
//     instance snapshot, or automated backup of your RDS Custom instance.
//
// Typically, deletion takes a few minutes.
//
// The MediaImport service that imports files from Amazon S3 to create CEVs isn't
// integrated with Amazon Web Services CloudTrail. If you turn on data logging for
// Amazon RDS in CloudTrail, calls to the DeleteCustomDbEngineVersion event aren't
// logged. However, you might see calls from the API gateway that accesses your
// Amazon S3 bucket. These calls originate from the MediaImport service for the
// DeleteCustomDbEngineVersion event.
//
// For more information, see [Deleting a CEV] in the Amazon RDS User Guide.
//
// [Deleting a CEV]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.delete
func (c *Client) DeleteCustomDBEngineVersion(ctx context.Context, params *DeleteCustomDBEngineVersionInput, optFns ...func(*Options)) (*DeleteCustomDBEngineVersionOutput, error) {
	if params == nil {
		params = &DeleteCustomDBEngineVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomDBEngineVersion", params, optFns, c.addOperationDeleteCustomDBEngineVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomDBEngineVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomDBEngineVersionInput struct {

	// The database engine. RDS Custom for Oracle supports the following values:
	//
	//   - custom-oracle-ee
	//
	//   - custom-oracle-ee-cdb
	//
	//   - custom-oracle-se2
	//
	//   - custom-oracle-se2-cdb
	//
	// This member is required.
	Engine *string

	// The custom engine version (CEV) for your DB instance. This option is required
	// for RDS Custom, but optional for Amazon RDS. The combination of Engine and
	// EngineVersion is unique per customer per Amazon Web Services Region.
	//
	// This member is required.
	EngineVersion *string

	noSmithyDocumentSerde
}

// This data type is used as a response element in the action
// DescribeDBEngineVersions .
type DeleteCustomDBEngineVersionOutput struct {

	// The creation time of the DB engine version.
	CreateTime *time.Time

	// JSON string that lists the installation files and parameters that RDS Custom
	// uses to create a custom engine version (CEV). RDS Custom applies the patches in
	// the order in which they're listed in the manifest. You can set the Oracle home,
	// Oracle base, and UNIX/Linux user and group using the installation parameters.
	// For more information, see [JSON fields in the CEV manifest]in the Amazon RDS User Guide.
	//
	// [JSON fields in the CEV manifest]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.preparing.html#custom-cev.preparing.manifest.fields
	CustomDBEngineVersionManifest *string

	// The description of the database engine.
	DBEngineDescription *string

	// A value that indicates the source media provider of the AMI based on the usage
	// operation. Applicable for RDS Custom for SQL Server.
	DBEngineMediaType *string

	// The ARN of the custom engine version.
	DBEngineVersionArn *string

	// The description of the database engine version.
	DBEngineVersionDescription *string

	// The name of the DB parameter group family for the database engine.
	DBParameterGroupFamily *string

	// The name of the Amazon S3 bucket that contains your database installation files.
	DatabaseInstallationFilesS3BucketName *string

	// The Amazon S3 directory that contains the database installation files. If not
	// specified, then no prefix is assumed.
	DatabaseInstallationFilesS3Prefix *string

	// The default character set for new instances of this engine version, if the
	// CharacterSetName parameter of the CreateDBInstance API isn't specified.
	DefaultCharacterSet *types.CharacterSet

	// The name of the database engine.
	Engine *string

	// The version number of the database engine.
	EngineVersion *string

	// The types of logs that the database engine has available for export to
	// CloudWatch Logs.
	ExportableLogTypes []string

	// The EC2 image
	Image *types.CustomDBEngineVersionAMI

	// The Amazon Web Services KMS key identifier for an encrypted CEV. This parameter
	// is required for RDS Custom, but optional for Amazon RDS.
	KMSKeyId *string

	// The major engine version of the CEV.
	MajorEngineVersion *string

	// The status of the DB engine version, either available or deprecated .
	Status *string

	// A list of the supported CA certificate identifiers.
	//
	// For more information, see [Using SSL/TLS to encrypt a connection to a DB instance] in the Amazon RDS User Guide and [Using SSL/TLS to encrypt a connection to a DB cluster] in the Amazon
	// Aurora User Guide.
	//
	// [Using SSL/TLS to encrypt a connection to a DB cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html
	// [Using SSL/TLS to encrypt a connection to a DB instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
	SupportedCACertificateIdentifiers []string

	// A list of the character sets supported by this engine for the CharacterSetName
	// parameter of the CreateDBInstance operation.
	SupportedCharacterSets []types.CharacterSet

	// A list of the supported DB engine modes.
	SupportedEngineModes []string

	// A list of features supported by the DB engine.
	//
	// The supported features vary by DB engine and DB engine version.
	//
	// To determine the supported features for a specific DB engine and DB engine
	// version using the CLI, use the following command:
	//
	//     aws rds describe-db-engine-versions --engine --engine-version
	//
	// For example, to determine the supported features for RDS for PostgreSQL version
	// 13.3 using the CLI, use the following command:
	//
	//     aws rds describe-db-engine-versions --engine postgres --engine-version 13.3
	//
	// The supported features are listed under SupportedFeatureNames in the output.
	SupportedFeatureNames []string

	// A list of the character sets supported by the Oracle DB engine for the
	// NcharCharacterSetName parameter of the CreateDBInstance operation.
	SupportedNcharCharacterSets []types.CharacterSet

	// A list of the time zones supported by this engine for the Timezone parameter of
	// the CreateDBInstance action.
	SupportedTimezones []types.Timezone

	// Indicates whether the engine version supports Babelfish for Aurora PostgreSQL.
	SupportsBabelfish *bool

	// Indicates whether the engine version supports rotating the server certificate
	// without rebooting the DB instance.
	SupportsCertificateRotationWithoutRestart *bool

	// Indicates whether you can use Aurora global databases with a specific DB engine
	// version.
	SupportsGlobalDatabases *bool

	// Indicates whether the DB engine version supports zero-ETL integrations with
	// Amazon Redshift.
	SupportsIntegrations *bool

	// Indicates whether the DB engine version supports Aurora Limitless Database.
	SupportsLimitlessDatabase *bool

	// Indicates whether the DB engine version supports forwarding write operations
	// from reader DB instances to the writer DB instance in the DB cluster. By
	// default, write operations aren't allowed on reader DB instances.
	//
	// Valid for: Aurora DB clusters only
	SupportsLocalWriteForwarding *bool

	// Indicates whether the engine version supports exporting the log types specified
	// by ExportableLogTypes to CloudWatch Logs.
	SupportsLogExportsToCloudwatchLogs *bool

	// Indicates whether you can use Aurora parallel query with a specific DB engine
	// version.
	SupportsParallelQuery *bool

	// Indicates whether the database engine version supports read replicas.
	SupportsReadReplica *bool

	// A list of tags. For more information, see [Tagging Amazon RDS Resources] in the Amazon RDS User Guide.
	//
	// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	TagList []types.Tag

	// A list of engine versions that this database engine version can be upgraded to.
	ValidUpgradeTarget []types.UpgradeTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomDBEngineVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCustomDBEngineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCustomDBEngineVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteCustomDBEngineVersion"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteCustomDBEngineVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomDBEngineVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteCustomDBEngineVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteCustomDBEngineVersion",
	}
}
