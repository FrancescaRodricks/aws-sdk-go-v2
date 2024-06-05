// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the detector version's description. You can update the metadata for any
// detector version ( DRAFT, ACTIVE, or INACTIVE ).
func (c *Client) UpdateDetectorVersionMetadata(ctx context.Context, params *UpdateDetectorVersionMetadataInput, optFns ...func(*Options)) (*UpdateDetectorVersionMetadataOutput, error) {
	if params == nil {
		params = &UpdateDetectorVersionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDetectorVersionMetadata", params, optFns, c.addOperationUpdateDetectorVersionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDetectorVersionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorVersionMetadataInput struct {

	// The description.
	//
	// This member is required.
	Description *string

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	noSmithyDocumentSerde
}

type UpdateDetectorVersionMetadataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDetectorVersionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDetectorVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDetectorVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDetectorVersionMetadata"); err != nil {
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
	if err = addOpUpdateDetectorVersionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetectorVersionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDetectorVersionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDetectorVersionMetadata",
	}
}
