// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the current traffic distribution for a given traffic distribution
// group.
func (c *Client) GetTrafficDistribution(ctx context.Context, params *GetTrafficDistributionInput, optFns ...func(*Options)) (*GetTrafficDistributionOutput, error) {
	if params == nil {
		params = &GetTrafficDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTrafficDistribution", params, optFns, c.addOperationGetTrafficDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTrafficDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTrafficDistributionInput struct {

	// The identifier of the traffic distribution group. This can be the ID or the ARN
	// if the API is being called in the Region where the traffic distribution group
	// was created. The ARN must be provided if the call is from the replicated Region.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetTrafficDistributionOutput struct {

	// The distribution of agents between the instance and its replica(s).
	AgentConfig *types.AgentConfig

	// The Amazon Resource Name (ARN) of the traffic distribution group.
	Arn *string

	// The identifier of the traffic distribution group. This can be the ID or the ARN
	// if the API is being called in the Region where the traffic distribution group
	// was created. The ARN must be provided if the call is from the replicated Region.
	Id *string

	// The distribution that determines which Amazon Web Services Regions should be
	// used to sign in agents in to both the instance and its replica(s).
	SignInConfig *types.SignInConfig

	// The distribution of traffic between the instance and its replicas.
	TelephonyConfig *types.TelephonyConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTrafficDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTrafficDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTrafficDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTrafficDistribution"); err != nil {
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
	if err = addOpGetTrafficDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrafficDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTrafficDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTrafficDistribution",
	}
}
