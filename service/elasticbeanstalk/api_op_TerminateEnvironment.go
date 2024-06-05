// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Terminates the specified environment.
func (c *Client) TerminateEnvironment(ctx context.Context, params *TerminateEnvironmentInput, optFns ...func(*Options)) (*TerminateEnvironmentOutput, error) {
	if params == nil {
		params = &TerminateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateEnvironment", params, optFns, c.addOperationTerminateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to terminate an environment.
type TerminateEnvironmentInput struct {

	// The ID of the environment to terminate.
	//
	// Condition: You must specify either this or an EnvironmentName, or both. If you
	// do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentId *string

	// The name of the environment to terminate.
	//
	// Condition: You must specify either this or an EnvironmentId, or both. If you do
	// not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentName *string

	// Terminates the target environment even if another environment in the same group
	// is dependent on it.
	ForceTerminate *bool

	// Indicates whether the associated AWS resources should shut down when the
	// environment is terminated:
	//
	//   - true : The specified environment as well as the associated AWS resources,
	//   such as Auto Scaling group and LoadBalancer, are terminated.
	//
	//   - false : AWS Elastic Beanstalk resource management is removed from the
	//   environment, but the AWS resources continue to operate.
	//
	// For more information, see the [AWS Elastic Beanstalk User Guide.]
	//
	// Default: true
	//
	// Valid Values: true | false
	//
	// [AWS Elastic Beanstalk User Guide.]: https://docs.aws.amazon.com/elasticbeanstalk/latest/ug/
	TerminateResources *bool

	noSmithyDocumentSerde
}

// Describes the properties of an environment.
type TerminateEnvironmentOutput struct {

	// Indicates if there is an in-progress environment configuration update or
	// application version deployment that you can cancel.
	//
	// true: There is an update in progress.
	//
	// false: There are no updates currently in progress.
	AbortableOperationInProgress *bool

	// The name of the application associated with this environment.
	ApplicationName *string

	// The URL to the CNAME for this environment.
	CNAME *string

	// The creation date for this environment.
	DateCreated *time.Time

	// The last modified date for this environment.
	DateUpdated *time.Time

	// Describes this environment.
	Description *string

	// For load-balanced, autoscaling environments, the URL to the LoadBalancer. For
	// single-instance environments, the IP address of the instance.
	EndpointURL *string

	// The environment's Amazon Resource Name (ARN), which can be used in other API
	// requests that require an ARN.
	EnvironmentArn *string

	// The ID of this environment.
	EnvironmentId *string

	// A list of links to other environments in the same group.
	EnvironmentLinks []types.EnvironmentLink

	// The name of this environment.
	EnvironmentName *string

	// Describes the health status of the environment. AWS Elastic Beanstalk indicates
	// the failure levels for a running environment:
	//
	//   - Red : Indicates the environment is not responsive. Occurs when three or more
	//   consecutive failures occur for an environment.
	//
	//   - Yellow : Indicates that something is wrong. Occurs when two consecutive
	//   failures occur for an environment.
	//
	//   - Green : Indicates the environment is healthy and fully functional.
	//
	//   - Grey : Default health for a new environment. The environment is not fully
	//   launched and health checks have not started or health checks are suspended
	//   during an UpdateEnvironment or RestartEnvironment request.
	//
	// Default: Grey
	Health types.EnvironmentHealth

	// Returns the health status of the application running in your environment. For
	// more information, see [Health Colors and Statuses].
	//
	// [Health Colors and Statuses]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html
	HealthStatus types.EnvironmentHealthStatus

	// The Amazon Resource Name (ARN) of the environment's operations role. For more
	// information, see [Operations roles]in the AWS Elastic Beanstalk Developer Guide.
	//
	// [Operations roles]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html
	OperationsRole *string

	// The ARN of the platform version.
	PlatformArn *string

	// The description of the AWS resources used by this environment.
	Resources *types.EnvironmentResourcesDescription

	//  The name of the SolutionStack deployed with this environment.
	SolutionStackName *string

	// The current operational status of the environment:
	//
	//   - Launching : Environment is in the process of initial deployment.
	//
	//   - Updating : Environment is in the process of updating its configuration
	//   settings or application version.
	//
	//   - Ready : Environment is available to have an action performed on it, such as
	//   update or terminate.
	//
	//   - Terminating : Environment is in the shut-down process.
	//
	//   - Terminated : Environment is not running.
	Status types.EnvironmentStatus

	// The name of the configuration template used to originally launch this
	// environment.
	TemplateName *string

	// Describes the current tier of this environment.
	Tier *types.EnvironmentTier

	// The application version deployed in this environment.
	VersionLabel *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTerminateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTerminateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTerminateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TerminateEnvironment"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTerminateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TerminateEnvironment",
	}
}
