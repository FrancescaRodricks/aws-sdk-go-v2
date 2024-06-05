// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a simulation application.
func (c *Client) UpdateSimulationApplication(ctx context.Context, params *UpdateSimulationApplicationInput, optFns ...func(*Options)) (*UpdateSimulationApplicationOutput, error) {
	if params == nil {
		params = &UpdateSimulationApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSimulationApplication", params, optFns, c.addOperationUpdateSimulationApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSimulationApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSimulationApplicationInput struct {

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string

	// Information about the robot software suite (ROS distribution).
	//
	// This member is required.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	//
	// This member is required.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The revision id for the robot application.
	CurrentRevisionId *string

	// The object that contains the Docker image URI for your simulation application.
	Environment *types.Environment

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The sources of the simulation application.
	Sources []types.SourceConfig

	noSmithyDocumentSerde
}

type UpdateSimulationApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the updated simulation application.
	Arn *string

	// The object that contains the Docker image URI used for your simulation
	// application.
	Environment *types.Environment

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The revision id of the simulation application.
	RevisionId *string

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The sources of the simulation application.
	Sources []types.Source

	// The version of the robot application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSimulationApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSimulationApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSimulationApplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSimulationApplication"); err != nil {
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
	if err = addOpUpdateSimulationApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSimulationApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSimulationApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSimulationApplication",
	}
}
