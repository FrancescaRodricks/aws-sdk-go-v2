// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a logger definition version.
func (c *Client) GetLoggerDefinitionVersion(ctx context.Context, params *GetLoggerDefinitionVersionInput, optFns ...func(*Options)) (*GetLoggerDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetLoggerDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoggerDefinitionVersion", params, optFns, c.addOperationGetLoggerDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoggerDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggerDefinitionVersionInput struct {

	// The ID of the logger definition.
	//
	// This member is required.
	LoggerDefinitionId *string

	// The ID of the logger definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListLoggerDefinitionVersions'' requests. If the version is the last one
	// that was associated with a logger definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	LoggerDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLoggerDefinitionVersionOutput struct {

	// The ARN of the logger definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the logger definition version
	// was created.
	CreationTimestamp *string

	// Information about the logger definition version.
	Definition *types.LoggerDefinitionVersion

	// The ID of the logger definition version.
	Id *string

	// The version of the logger definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLoggerDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLoggerDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLoggerDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLoggerDefinitionVersion"); err != nil {
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
	if err = addOpGetLoggerDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLoggerDefinitionVersion",
	}
}
