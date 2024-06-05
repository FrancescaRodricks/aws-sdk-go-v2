// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new job definition.
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, c.addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// The JobCommand that runs this job.
	//
	// This member is required.
	Command *types.JobCommand

	// The name you assign to this job definition. It must be unique in your account.
	//
	// This member is required.
	Name *string

	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	//
	// This member is required.
	Role *string

	// This parameter is deprecated. Use MaxCapacity instead.
	//
	// The number of Glue data processing units (DPUs) to allocate to this Job. You
	// can allocate a minimum of 2 DPUs; the default is 10. A DPU is a relative measure
	// of processing power that consists of 4 vCPUs of compute capacity and 16 GB of
	// memory. For more information, see the [Glue pricing page].
	//
	// [Glue pricing page]: https://aws.amazon.com/glue/pricing/
	//
	// Deprecated: This property is deprecated, use MaxCapacity instead.
	AllocatedCapacity int32

	// The representation of a directed acyclic graph on which both the Glue Studio
	// visual component and Glue Studio code generation is based.
	CodeGenConfigurationNodes map[string]types.CodeGenConfigurationNode

	// The connections used for this job.
	Connections *types.ConnectionsList

	// The default arguments for every run of this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, as
	// well as arguments that Glue itself consumes.
	//
	// Job arguments may be logged. Do not pass plaintext secrets as arguments.
	// Retrieve secrets from a Glue Connection, Secrets Manager or other secret
	// management mechanism if you intend to keep them within the Job.
	//
	// For information about how to specify and consume your own Job arguments, see
	// the [Calling Glue APIs in Python]topic in the developer guide.
	//
	// For information about the arguments you can provide to this field when
	// configuring Spark jobs, see the [Special Parameters Used by Glue]topic in the developer guide.
	//
	// For information about the arguments you can provide to this field when
	// configuring Ray jobs, see [Using job parameters in Ray jobs]in the developer guide.
	//
	// [Using job parameters in Ray jobs]: https://docs.aws.amazon.com/glue/latest/dg/author-job-ray-job-parameters.html
	// [Calling Glue APIs in Python]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html
	// [Special Parameters Used by Glue]: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	DefaultArguments map[string]string

	// Description of the job being defined.
	Description *string

	// Indicates whether the job is run with a standard or flexible execution class.
	// The standard execution-class is ideal for time-sensitive workloads that require
	// fast job startup and dedicated resources.
	//
	// The flexible execution class is appropriate for time-insensitive jobs whose
	// start and completion times may vary.
	//
	// Only jobs with Glue version 3.0 and above and command type glueetl will be
	// allowed to set ExecutionClass to FLEX . The flexible execution class is
	// available for Spark jobs.
	ExecutionClass types.ExecutionClass

	// An ExecutionProperty specifying the maximum number of concurrent runs allowed
	// for this job.
	ExecutionProperty *types.ExecutionProperty

	// In Spark jobs, GlueVersion determines the versions of Apache Spark and Python
	// that Glue available in a job. The Python version indicates the version supported
	// for jobs of type Spark.
	//
	// Ray jobs should set GlueVersion to 4.0 or greater. However, the versions of
	// Ray, Python and additional libraries available in your Ray job are determined by
	// the Runtime parameter of the Job command.
	//
	// For more information about the available Glue versions and corresponding Spark
	// and Python versions, see [Glue version]in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	//
	// [Glue version]: https://docs.aws.amazon.com/glue/latest/dg/add-job.html
	GlueVersion *string

	// A mode that describes how a job was created. Valid values are:
	//
	//   - SCRIPT - The job was created using the Glue Studio script editor.
	//
	//   - VISUAL - The job was created using the Glue Studio visual editor.
	//
	//   - NOTEBOOK - The job was created using an interactive sessions notebook.
	//
	// When the JobMode field is missing or null, SCRIPT is assigned as the default
	// value.
	JobMode types.JobMode

	// This field is reserved for future use.
	LogUri *string

	// This field specifies a day of the week and hour for a maintenance window for
	// streaming jobs. Glue periodically performs maintenance activities. During these
	// maintenance windows, Glue will need to restart your streaming jobs.
	//
	// Glue will restart the job within 3 hours of the specified maintenance window.
	// For instance, if you set up the maintenance window for Monday at 10:00AM GMT,
	// your jobs will be restarted between 10:00AM GMT to 1:00PM GMT.
	MaintenanceWindow *string

	// For Glue version 1.0 or earlier jobs, using the standard worker type, the
	// number of Glue data processing units (DPUs) that can be allocated when this job
	// runs. A DPU is a relative measure of processing power that consists of 4 vCPUs
	// of compute capacity and 16 GB of memory. For more information, see the [Glue pricing page].
	//
	// For Glue version 2.0+ jobs, you cannot specify a Maximum capacity . Instead, you
	// should specify a Worker type and the Number of workers .
	//
	// Do not set MaxCapacity if using WorkerType and NumberOfWorkers .
	//
	// The value that can be allocated for MaxCapacity depends on whether you are
	// running a Python shell job, an Apache Spark ETL job, or an Apache Spark
	// streaming ETL job:
	//
	//   - When you specify a Python shell job ( JobCommand.Name ="pythonshell"), you
	//   can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	//   - When you specify an Apache Spark ETL job ( JobCommand.Name ="glueetl") or
	//   Apache Spark streaming ETL job ( JobCommand.Name ="gluestreaming"), you can
	//   allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a
	//   fractional DPU allocation.
	//
	// [Glue pricing page]: https://aws.amazon.com/glue/pricing/
	MaxCapacity *float64

	// The maximum number of times to retry this job if it fails.
	MaxRetries int32

	// Arguments for this job that are not overridden when providing job arguments in
	// a job run, specified as name-value pairs.
	NonOverridableArguments map[string]string

	// Specifies configuration properties of a job notification.
	NotificationProperty *types.NotificationProperty

	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	NumberOfWorkers *int32

	// The name of the SecurityConfiguration structure to be used with this job.
	SecurityConfiguration *string

	// The details for a source control configuration for a job, allowing
	// synchronization of job artifacts to or from a remote repository.
	SourceControlDetails *types.SourceControlDetails

	// The tags to use with this job. You may use tags to limit access to the job. For
	// more information about tags in Glue, see [Amazon Web Services Tags in Glue]in the developer guide.
	//
	// [Amazon Web Services Tags in Glue]: https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html
	Tags map[string]string

	// The job timeout in minutes. This is the maximum time that a job run can consume
	// resources before it is terminated and enters TIMEOUT status. The default is
	// 2,880 minutes (48 hours).
	Timeout *int32

	// The type of predefined worker that is allocated when a job runs. Accepts a
	// value of G.1X, G.2X, G.4X, G.8X or G.025X for Spark jobs. Accepts the value Z.2X
	// for Ray jobs.
	//
	//   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPUs, 16 GB of
	//   memory) with 84GB disk (approximately 34GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//
	//   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPUs, 32 GB of
	//   memory) with 128GB disk (approximately 77GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//
	//   - For the G.4X worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB of
	//   memory) with 256GB disk (approximately 235GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs in the following
	//   Amazon Web Services Regions: US East (Ohio), US East (N. Virginia), US West
	//   (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo),
	//   Canada (Central), Europe (Frankfurt), Europe (Ireland), and Europe (Stockholm).
	//
	//   - For the G.8X worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB of
	//   memory) with 512GB disk (approximately 487GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs, in the same Amazon
	//   Web Services Regions as supported for the G.4X worker type.
	//
	//   - For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPUs, 4 GB of
	//   memory) with 84GB disk (approximately 34GB free), and provides 1 executor per
	//   worker. We recommend this worker type for low volume streaming jobs. This worker
	//   type is only available for Glue version 3.0 streaming jobs.
	//
	//   - For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPUs, 64 GB of
	//   memory) with 128 GB disk (approximately 120GB free), and provides up to 8 Ray
	//   workers based on the autoscaler.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type CreateJobOutput struct {

	// The unique name that was provided for this job definition.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateJob"); err != nil {
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
	if err = addOpCreateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateJob",
	}
}
