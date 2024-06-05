// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with most of the source's original
// configuration, including the default security group and DB parameter group. By
// default, the new DB instance is created as a Single-AZ deployment, except when
// the instance is a SQL Server instance that has an option group associated with
// mirroring. In this case, the instance becomes a Multi-AZ deployment, not a
// Single-AZ deployment.
//
// If you want to replace your original DB instance with the new, restored DB
// instance, then rename your original DB instance before you call the
// RestoreDBInstanceFromDBSnapshot operation. RDS doesn't allow two DB instances
// with the same name. After you have renamed your original DB instance with a
// different identifier, then you can pass the original name of the DB instance as
// the DBInstanceIdentifier in the call to the RestoreDBInstanceFromDBSnapshot
// operation. The result is that you replace the original DB instance with the DB
// instance created from the snapshot.
//
// If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
// must be the ARN of the shared DB snapshot.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot .
func (c *Client) RestoreDBInstanceFromDBSnapshot(ctx context.Context, params *RestoreDBInstanceFromDBSnapshotInput, optFns ...func(*Options)) (*RestoreDBInstanceFromDBSnapshotOutput, error) {
	if params == nil {
		params = &RestoreDBInstanceFromDBSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBInstanceFromDBSnapshot", params, optFns, c.addOperationRestoreDBInstanceFromDBSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBInstanceFromDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBInstanceFromDBSnapshotInput struct {

	// The name of the DB instance to create from the DB snapshot. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 numbers, letters, or hyphens.
	//
	//   - First character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-snapshot-id
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The amount of storage (in gibibytes) to allocate initially for the DB instance.
	// Follow the allocation rules specified in CreateDBInstance.
	//
	// Be sure to allocate enough storage for your new DB instance so that the restore
	// operation can succeed. You can also allocate additional storage for future
	// growth.
	AllocatedStorage *int32

	// Specifies whether to automatically apply minor version upgrades to the DB
	// instance during the maintenance window.
	//
	// If you restore an RDS Custom DB instance, you must disable this parameter.
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the DB instance will be created.
	//
	// Default: A random, system-chosen Availability Zone.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	//
	// Example: us-east-1a
	AvailabilityZone *string

	// Specifies where automated backups and manual snapshots are stored for the
	// restored DB instance.
	//
	// Possible values are outposts (Amazon Web Services Outposts) and region (Amazon
	// Web Services Region). The default is region .
	//
	// For more information, see [Working with Amazon RDS on Amazon Web Services Outposts] in the Amazon RDS User Guide.
	//
	// [Working with Amazon RDS on Amazon Web Services Outposts]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html
	BackupTarget *string

	// The CA certificate identifier to use for the DB instance's server certificate.
	//
	// This setting doesn't apply to RDS Custom DB instances.
	//
	// For more information, see [Using SSL/TLS to encrypt a connection to a DB instance] in the Amazon RDS User Guide and [Using SSL/TLS to encrypt a connection to a DB cluster] in the Amazon
	// Aurora User Guide.
	//
	// [Using SSL/TLS to encrypt a connection to a DB cluster]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html
	// [Using SSL/TLS to encrypt a connection to a DB instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
	CACertificateIdentifier *string

	// Specifies whether to copy all tags from the restored DB instance to snapshots
	// of the DB instance.
	//
	// In most cases, tags aren't copied by default. However, when you restore a DB
	// instance from a DB snapshot, RDS checks whether you specify new tags. If yes,
	// the new tags are added to the restored DB instance. If there are no new tags,
	// RDS looks for the tags from the source DB instance for the DB snapshot, and then
	// adds those tags to the restored DB instance.
	//
	// For more information, see [Copying tags to DB instance snapshots] in the Amazon RDS User Guide.
	//
	// [Copying tags to DB instance snapshots]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.CopyTags
	CopyTagsToSnapshot *bool

	// The instance profile associated with the underlying Amazon EC2 instance of an
	// RDS Custom DB instance. The instance profile must meet the following
	// requirements:
	//
	//   - The profile must exist in your account.
	//
	//   - The profile must have an IAM role that Amazon EC2 has permissions to assume.
	//
	//   - The instance profile name and the associated IAM role name must start with
	//   the prefix AWSRDSCustom .
	//
	// For the list of permissions required for the IAM role, see [Configure IAM and your VPC] in the Amazon RDS
	// User Guide.
	//
	// This setting is required for RDS Custom.
	//
	// [Configure IAM and your VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc
	CustomIamInstanceProfile *string

	// The identifier for the Multi-AZ DB cluster snapshot to restore from.
	//
	// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing Multi-AZ DB cluster snapshot.
	//
	//   - Can't be specified when DBSnapshotIdentifier is specified.
	//
	//   - Must be specified when DBSnapshotIdentifier isn't specified.
	//
	//   - If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the
	//   DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.
	//
	//   - Can't be the identifier of an Aurora DB cluster snapshot.
	//
	// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
	DBClusterSnapshotIdentifier *string

	// The compute and memory capacity of the Amazon RDS DB instance, for example
	// db.m4.large. Not all DB instance classes are available in all Amazon Web
	// Services Regions, or for all database engines. For the full list of DB instance
	// classes, and availability for your engine, see [DB Instance Class]in the Amazon RDS User Guide.
	//
	// Default: The same DBInstanceClass as the original DB instance.
	//
	// [DB Instance Class]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html
	DBInstanceClass *string

	// The name of the database for the restored DB instance.
	//
	// This parameter only applies to RDS for Oracle and RDS for SQL Server DB
	// instances. It doesn't apply to the other engines or to RDS Custom DB instances.
	DBName *string

	// The name of the DB parameter group to associate with this DB instance.
	//
	// If you don't specify a value for DBParameterGroupName , then RDS uses the
	// default DBParameterGroup for the specified DB engine.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Constraints:
	//
	//   - If supplied, must match the name of an existing DB parameter group.
	//
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//
	//   - First character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string

	// The identifier for the DB snapshot to restore from.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DB snapshot.
	//
	//   - Can't be specified when DBClusterSnapshotIdentifier is specified.
	//
	//   - Must be specified when DBClusterSnapshotIdentifier isn't specified.
	//
	//   - If you are restoring from a shared manual DB snapshot, the
	//   DBSnapshotIdentifier must be the ARN of the shared DB snapshot.
	DBSnapshotIdentifier *string

	// The name of the DB subnet group to use for the new instance.
	//
	// Constraints:
	//
	//   - If supplied, must match the name of an existing DB subnet group.
	//
	// Example: mydbsubnetgroup
	DBSubnetGroupName *string

	// Specifies whether to enable a dedicated log volume (DLV) for the DB instance.
	DedicatedLogVolume *bool

	// Specifies whether to enable deletion protection for the DB instance. The
	// database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. For more information, see [Deleting a DB Instance].
	//
	// [Deleting a DB Instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html
	DeletionProtection *bool

	// The Active Directory directory ID to restore the DB instance in. The domain/
	// must be created prior to this operation. Currently, you can create only Db2,
	// MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances in an Active
	// Directory Domain.
	//
	// For more information, see [Kerberos Authentication] in the Amazon RDS User Guide.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// [Kerberos Authentication]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html
	Domain *string

	// The ARN for the Secrets Manager secret with the credentials for the user
	// joining the domain.
	//
	// Constraints:
	//
	//   - Can't be longer than 64 characters.
	//
	// Example:
	// arn:aws:secretsmanager:region:account-number:secret:myselfmanagedADtestsecret-123456
	DomainAuthSecretArn *string

	// The IPv4 DNS IP addresses of your primary and secondary Active Directory domain
	// controllers.
	//
	// Constraints:
	//
	//   - Two IP addresses must be provided. If there isn't a secondary domain
	//   controller, use the IP address of the primary domain controller for both entries
	//   in the list.
	//
	// Example: 123.124.125.126,234.235.236.237
	DomainDnsIps []string

	// The fully qualified domain name (FQDN) of an Active Directory domain.
	//
	// Constraints:
	//
	//   - Can't be longer than 64 characters.
	//
	// Example: mymanagedADtest.mymanagedAD.mydomain
	DomainFqdn *string

	// The name of the IAM role to use when making API calls to the Directory Service.
	//
	// This setting doesn't apply to RDS Custom DB instances.
	DomainIAMRoleName *string

	// The Active Directory organizational unit for your DB instance to join.
	//
	// Constraints:
	//
	//   - Must be in the distinguished name format.
	//
	//   - Can't be longer than 64 characters.
	//
	// Example: OU=mymanagedADtestOU,DC=mymanagedADtest,DC=mymanagedAD,DC=mydomain
	DomainOu *string

	// The list of logs for the restored DB instance to export to CloudWatch Logs. The
	// values in the list depend on the DB engine. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs]in the
	// Amazon RDS User Guide.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// [Publishing Database Logs to Amazon CloudWatch Logs]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch
	EnableCloudwatchLogsExports []string

	// Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on
	// Outposts DB instance.
	//
	// A CoIP provides local or external connectivity to resources in your Outpost
	// subnets through your on-premises network. For some use cases, a CoIP can provide
	// lower latency for connections to the DB instance from outside of its virtual
	// private cloud (VPC) on your local network.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// For more information about RDS on Outposts, see [Working with Amazon RDS on Amazon Web Services Outposts] in the Amazon RDS User Guide.
	//
	// For more information about CoIPs, see [Customer-owned IP addresses] in the Amazon Web Services Outposts User
	// Guide.
	//
	// [Working with Amazon RDS on Amazon Web Services Outposts]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html
	// [Customer-owned IP addresses]: https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing
	EnableCustomerOwnedIp *bool

	// Specifies whether to enable mapping of Amazon Web Services Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	//
	// For more information about IAM database authentication, see [IAM Database Authentication for MySQL and PostgreSQL] in the Amazon RDS
	// User Guide.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// [IAM Database Authentication for MySQL and PostgreSQL]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html
	EnableIAMDatabaseAuthentication *bool

	// The database engine to use for the new instance.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source. For example, you
	// can restore a MariaDB 10.1 DB instance from a MySQL 5.6 snapshot.
	//
	// Valid Values:
	//
	//   - db2-ae
	//
	//   - db2-se
	//
	//   - mariadb
	//
	//   - mysql
	//
	//   - oracle-ee
	//
	//   - oracle-ee-cdb
	//
	//   - oracle-se2
	//
	//   - oracle-se2-cdb
	//
	//   - postgres
	//
	//   - sqlserver-ee
	//
	//   - sqlserver-se
	//
	//   - sqlserver-ex
	//
	//   - sqlserver-web
	Engine *string

	// The life cycle type for this DB instance.
	//
	// By default, this value is set to open-source-rds-extended-support , which
	// enrolls your DB instance into Amazon RDS Extended Support. At the end of
	// standard support, you can avoid charges for Extended Support by setting the
	// value to open-source-rds-extended-support-disabled . In this case, RDS
	// automatically upgrades your restored DB instance to a higher engine version, if
	// the major engine version is past its end of standard support date.
	//
	// You can use this setting to enroll your DB instance into Amazon RDS Extended
	// Support. With RDS Extended Support, you can run the selected major engine
	// version on your DB instance past the end of standard support for that engine
	// version. For more information, see [Using Amazon RDS Extended Support]in the Amazon RDS User Guide.
	//
	// This setting applies only to RDS for MySQL and RDS for PostgreSQL. For Amazon
	// Aurora DB instances, the life cycle type is managed by the DB cluster.
	//
	// Valid Values: open-source-rds-extended-support |
	// open-source-rds-extended-support-disabled
	//
	// Default: open-source-rds-extended-support
	//
	// [Using Amazon RDS Extended Support]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html
	EngineLifecycleSupport *string

	// Specifies the amount of provisioned IOPS for the DB instance, expressed in I/O
	// operations per second. If this parameter isn't specified, the IOPS value is
	// taken from the backup. If this parameter is set to 0, the new instance is
	// converted to a non-PIOPS instance. The conversion takes additional time, though
	// your DB instance is available for connections before the conversion starts.
	//
	// The provisioned IOPS value must follow the requirements for your database
	// engine. For more information, see [Amazon RDS Provisioned IOPS storage]in the Amazon RDS User Guide.
	//
	// Constraints: Must be an integer greater than 1000.
	//
	// [Amazon RDS Provisioned IOPS storage]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS
	Iops *int32

	// License model information for the restored DB instance.
	//
	// License models for RDS for Db2 require additional configuration. The Bring Your
	// Own License (BYOL) model requires a custom parameter group. The Db2 license
	// through Amazon Web Services Marketplace model requires an Amazon Web Services
	// Marketplace subscription. For more information, see [RDS for Db2 licensing options]in the Amazon RDS User
	// Guide.
	//
	// This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.
	//
	// Valid Values:
	//
	//   - RDS for Db2 - bring-your-own-license | marketplace-license
	//
	//   - RDS for MariaDB - general-public-license
	//
	//   - RDS for Microsoft SQL Server - license-included
	//
	//   - RDS for MySQL - general-public-license
	//
	//   - RDS for Oracle - bring-your-own-license | license-included
	//
	//   - RDS for PostgreSQL - postgresql-license
	//
	// Default: Same as the source.
	//
	// [RDS for Db2 licensing options]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html
	LicenseModel *string

	// Specifies whether the DB instance is a Multi-AZ deployment.
	//
	// This setting doesn't apply to RDS Custom.
	//
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	MultiAZ *bool

	// The network type of the DB instance.
	//
	// Valid Values:
	//
	//   - IPV4
	//
	//   - DUAL
	//
	// The network type is determined by the DBSubnetGroup specified for the DB
	// instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and
	// the IPv6 protocols ( DUAL ).
	//
	// For more information, see [Working with a DB instance in a VPC] in the Amazon RDS User Guide.
	//
	// [Working with a DB instance in a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html
	NetworkType *string

	// The name of the option group to be used for the restored DB instance.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance after it is associated with a DB instance.
	//
	// This setting doesn't apply to RDS Custom.
	OptionGroupName *string

	// The port number on which the database accepts connections.
	//
	// Default: The same port as the original DB instance
	//
	// Constraints: Value must be 1150-65535
	Port *int32

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	//
	// This setting doesn't apply to RDS Custom.
	ProcessorFeatures []types.ProcessorFeature

	// Specifies whether the DB instance is publicly accessible.
	//
	// When the DB instance is publicly accessible, its Domain Name System (DNS)
	// endpoint resolves to the private IP address from within the DB instance's
	// virtual private cloud (VPC). It resolves to the public IP address from outside
	// of the DB instance's VPC. Access to the DB instance is ultimately controlled by
	// the security group it uses. That public access is not permitted if the security
	// group assigned to the DB instance doesn't permit it.
	//
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address.
	//
	// For more information, see CreateDBInstance.
	PubliclyAccessible *bool

	// Specifies the storage throughput value for the DB instance.
	//
	// This setting doesn't apply to RDS Custom or Amazon Aurora.
	StorageThroughput *int32

	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid Values: gp2 | gp3 | io1 | io2 | standard
	//
	// If you specify io1 , io2 , or gp3 , you must also include a value for the Iops
	// parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string

	// A list of tags. For more information, see [Tagging Amazon RDS Resources] in the Amazon RDS User Guide.
	//
	// [Tagging Amazon RDS Resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	Tags []types.Tag

	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	//
	// This setting doesn't apply to RDS Custom.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	//
	// This setting doesn't apply to RDS Custom.
	TdeCredentialPassword *string

	// Specifies whether the DB instance class of the DB instance uses its default
	// processor features.
	//
	// This setting doesn't apply to RDS Custom.
	UseDefaultProcessorFeatures *bool

	// A list of EC2 VPC security groups to associate with this DB instance.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBInstanceFromDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the operations CreateDBInstance
	// , CreateDBInstanceReadReplica , DeleteDBInstance , DescribeDBInstances ,
	// ModifyDBInstance , PromoteReadReplica , RebootDBInstance ,
	// RestoreDBInstanceFromDBSnapshot , RestoreDBInstanceFromS3 ,
	// RestoreDBInstanceToPointInTime , StartDBInstance , and StopDBInstance .
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBInstanceFromDBSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBInstanceFromDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBInstanceFromDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreDBInstanceFromDBSnapshot"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addTimeOffsetDeserializer(stack, c); err != nil {
		return err
	}
	if err = addOpRestoreDBInstanceFromDBSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBInstanceFromDBSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBInstanceFromDBSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreDBInstanceFromDBSnapshot",
	}
}
