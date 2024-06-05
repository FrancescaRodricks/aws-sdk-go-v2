// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This section contains the Amazon Managed Workflows for Apache Airflow (MWAA)
// API reference documentation to create an environment. For more information, see [Get started with Amazon Managed Workflows for Apache Airflow]
// .
//
// [Get started with Amazon Managed Workflows for Apache Airflow]: https://docs.aws.amazon.com/mwaa/latest/userguide/get-started.html
type CreateEnvironmentInput struct {

	// The relative path to the DAGs folder on your Amazon S3 bucket. For example, dags
	// . For more information, see [Adding or updating DAGs].
	//
	// [Adding or updating DAGs]: https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html
	//
	// This member is required.
	DagS3Path *string

	// The Amazon Resource Name (ARN) of the execution role for your environment. An
	// execution role is an Amazon Web Services Identity and Access Management (IAM)
	// role that grants MWAA permission to access Amazon Web Services services and
	// resources used by your environment. For example,
	// arn:aws:iam::123456789:role/my-execution-role . For more information, see [Amazon MWAA Execution role].
	//
	// [Amazon MWAA Execution role]: https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment .
	//
	// This member is required.
	Name *string

	// The VPC networking components used to secure and enable network traffic between
	// the Amazon Web Services resources for your environment. For more information,
	// see [About networking on Amazon MWAA].
	//
	// [About networking on Amazon MWAA]: https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html
	//
	// This member is required.
	NetworkConfiguration *types.NetworkConfiguration

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and
	// supporting files are stored. For example,
	// arn:aws:s3:::my-airflow-bucket-unique-name . For more information, see [Create an Amazon S3 bucket for Amazon MWAA].
	//
	// [Create an Amazon S3 bucket for Amazon MWAA]: https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html
	//
	// This member is required.
	SourceBucketArn *string

	// A list of key-value pairs containing the Apache Airflow configuration options
	// you want to attach to your environment. For more information, see [Apache Airflow configuration options].
	//
	// [Apache Airflow configuration options]: https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html
	AirflowConfigurationOptions map[string]string

	// The Apache Airflow version for your environment. If no value is specified, it
	// defaults to the latest version. For more information, see [Apache Airflow versions on Amazon Managed Workflows for Apache Airflow (MWAA)].
	//
	// Valid values: 1.10.12 , 2.0.2 , 2.2.2 , 2.4.3 , 2.5.1 , 2.6.3 , 2.7.2
	//     2.8.1
	//
	// [Apache Airflow versions on Amazon Managed Workflows for Apache Airflow (MWAA)]: https://docs.aws.amazon.com/mwaa/latest/userguide/airflow-versions.html
	AirflowVersion *string

	// Defines whether the VPC endpoints configured for the environment are created,
	// and managed, by the customer or by Amazon MWAA. If set to SERVICE , Amazon MWAA
	// will create and manage the required VPC endpoints in your VPC. If set to
	// CUSTOMER , you must create, and manage, the VPC endpoints for your VPC. If you
	// choose to create an environment in a shared VPC, you must set this value to
	// CUSTOMER . In a shared VPC deployment, the environment will remain in PENDING
	// status until you create the VPC endpoints. If you do not take action to create
	// the endpoints within 72 hours, the status will change to CREATE_FAILED . You can
	// delete the failed environment and create a new one.
	EndpointManagement types.EndpointManagement

	// The environment class type. Valid values: mw1.small , mw1.medium , mw1.large ,
	// mw1.xlarge , and mw1.2xlarge . For more information, see [Amazon MWAA environment class].
	//
	// [Amazon MWAA environment class]: https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html
	EnvironmentClass *string

	// The Amazon Web Services Key Management Service (KMS) key to encrypt the data in
	// your environment. You can use an Amazon Web Services owned CMK, or a Customer
	// managed CMK (advanced). For more information, see [Create an Amazon MWAA environment].
	//
	// [Create an Amazon MWAA environment]: https://docs.aws.amazon.com/mwaa/latest/userguide/create-environment.html
	KmsKey *string

	// Defines the Apache Airflow logs to send to CloudWatch Logs.
	LoggingConfiguration *types.LoggingConfigurationInput

	//  The maximum number of web servers that you want to run in your environment.
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you
	// specify for MaxWebservers when you interact with your Apache Airflow
	// environment using Apache Airflow REST API, or the Apache Airflow CLI. For
	// example, in scenarios where your workload requires network calls to the Apache
	// Airflow REST API with a high transaction-per-second (TPS) rate, Amazon MWAA will
	// increase the number of web servers up to the number set in MaxWebserers . As TPS
	// rates decrease Amazon MWAA disposes of the additional web servers, and scales
	// down to the number set in MinxWebserers .
	//
	// Valid values: Accepts between 2 and 5 . Defaults to 2 .
	MaxWebservers *int32

	// The maximum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. For example, 20 . When there are no more tasks running, and no
	// more in the queue, MWAA disposes of the extra workers leaving the one worker
	// that is included with your environment, or the number you specify in MinWorkers .
	MaxWorkers *int32

	//  The minimum number of web servers that you want to run in your environment.
	// Amazon MWAA scales the number of Apache Airflow web servers up to the number you
	// specify for MaxWebservers when you interact with your Apache Airflow
	// environment using Apache Airflow REST API, or the Apache Airflow CLI. As the
	// transaction-per-second rate, and the network load, decrease, Amazon MWAA
	// disposes of the additional web servers, and scales down to the number set in
	// MinxWebserers .
	//
	// Valid values: Accepts between 2 and 5 . Defaults to 2 .
	MinWebservers *int32

	// The minimum number of workers that you want to run in your environment. MWAA
	// scales the number of Apache Airflow workers up to the number you specify in the
	// MaxWorkers field. When there are no more tasks running, and no more in the
	// queue, MWAA disposes of the extra workers leaving the worker count you specify
	// in the MinWorkers field. For example, 2 .
	MinWorkers *int32

	// The version of the plugins.zip file on your Amazon S3 bucket. You must specify
	// a version each time a plugins.zip file is updated. For more information, see [How S3 Versioning works].
	//
	// [How S3 Versioning works]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html
	PluginsS3ObjectVersion *string

	// The relative path to the plugins.zip file on your Amazon S3 bucket. For
	// example, plugins.zip . If specified, then the plugins.zip version is required.
	// For more information, see [Installing custom plugins].
	//
	// [Installing custom plugins]: https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html
	PluginsS3Path *string

	// The version of the requirements.txt file on your Amazon S3 bucket. You must
	// specify a version each time a requirements.txt file is updated. For more
	// information, see [How S3 Versioning works].
	//
	// [How S3 Versioning works]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html
	RequirementsS3ObjectVersion *string

	// The relative path to the requirements.txt file on your Amazon S3 bucket. For
	// example, requirements.txt . If specified, then a version is required. For more
	// information, see [Installing Python dependencies].
	//
	// [Installing Python dependencies]: https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html
	RequirementsS3Path *string

	// The number of Apache Airflow schedulers to run in your environment. Valid
	// values:
	//
	//   - v2 - Accepts between 2 to 5 . Defaults to 2 .
	//
	//   - v1 - Accepts 1 .
	Schedulers *int32

	// The version of the startup shell script in your Amazon S3 bucket. You must
	// specify the [version ID]that Amazon S3 assigns to the file every time you update the
	// script.
	//
	// Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no
	// more than 1,024 bytes long. The following is an example:
	//
	//     3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo
	//
	// For more information, see [Using a startup script].
	//
	// [Using a startup script]: https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html
	// [version ID]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/versioning-workflows.html
	StartupScriptS3ObjectVersion *string

	// The relative path to the startup shell script in your Amazon S3 bucket. For
	// example, s3://mwaa-environment/startup.sh .
	//
	// Amazon MWAA runs the script as your environment starts, and before running the
	// Apache Airflow process. You can use this script to install dependencies, modify
	// Apache Airflow configuration options, and set environment variables. For more
	// information, see [Using a startup script].
	//
	// [Using a startup script]: https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html
	StartupScriptS3Path *string

	// The key-value tag pairs you want to associate to your environment. For example,
	// "Environment": "Staging" . For more information, see [Tagging Amazon Web Services resources].
	//
	// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags map[string]string

	// Defines the access mode for the Apache Airflow web server. For more
	// information, see [Apache Airflow access modes].
	//
	// [Apache Airflow access modes]: https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html
	WebserverAccessMode types.WebserverAccessMode

	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour
	// standard time to start weekly maintenance updates of your environment in the
	// following format: DAY:HH:MM . For example: TUE:03:30 . You can specify a start
	// time in 30 minute increments only.
	WeeklyMaintenanceWindowStart *string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) returned in the response for the environment.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironment"); err != nil {
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
	if err = addEndpointPrefix_opCreateEnvironmentMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateEnvironmentMiddleware struct {
}

func (*endpointPrefix_opCreateEnvironmentMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateEnvironmentMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateEnvironmentMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateEnvironmentMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironment",
	}
}
