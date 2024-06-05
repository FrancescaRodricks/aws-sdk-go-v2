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

// Validates the supplied schema. This call has no side effects, it simply
// validates using the supplied schema using DataFormat as the format. Since it
// does not take a schema set name, no compatibility checks are performed.
func (c *Client) CheckSchemaVersionValidity(ctx context.Context, params *CheckSchemaVersionValidityInput, optFns ...func(*Options)) (*CheckSchemaVersionValidityOutput, error) {
	if params == nil {
		params = &CheckSchemaVersionValidityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckSchemaVersionValidity", params, optFns, c.addOperationCheckSchemaVersionValidityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckSchemaVersionValidityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckSchemaVersionValidityInput struct {

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	//
	// This member is required.
	DataFormat types.DataFormat

	// The definition of the schema that has to be validated.
	//
	// This member is required.
	SchemaDefinition *string

	noSmithyDocumentSerde
}

type CheckSchemaVersionValidityOutput struct {

	// A validation failure error message.
	Error *string

	// Return true, if the schema is valid and false otherwise.
	Valid bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCheckSchemaVersionValidityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCheckSchemaVersionValidity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckSchemaVersionValidity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CheckSchemaVersionValidity"); err != nil {
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
	if err = addOpCheckSchemaVersionValidityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckSchemaVersionValidity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckSchemaVersionValidity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CheckSchemaVersionValidity",
	}
}
