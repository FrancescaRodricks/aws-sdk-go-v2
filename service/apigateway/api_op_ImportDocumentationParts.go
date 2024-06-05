// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports documentation parts
func (c *Client) ImportDocumentationParts(ctx context.Context, params *ImportDocumentationPartsInput, optFns ...func(*Options)) (*ImportDocumentationPartsOutput, error) {
	if params == nil {
		params = &ImportDocumentationPartsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportDocumentationParts", params, optFns, c.addOperationImportDocumentationPartsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportDocumentationPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Import documentation parts from an external (e.g., OpenAPI) definition file.
type ImportDocumentationPartsInput struct {

	// Raw byte array representing the to-be-imported documentation parts. To import
	// from an OpenAPI file, this is a JSON object.
	//
	// This member is required.
	Body []byte

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A query parameter to specify whether to rollback the documentation importation (
	// true ) or not ( false ) when a warning is encountered. The default value is
	// false .
	FailOnWarnings bool

	// A query parameter to indicate whether to overwrite ( overwrite ) any existing
	// DocumentationParts definition or to merge ( merge ) the new definition into the
	// existing one. The default value is merge .
	Mode types.PutMode

	noSmithyDocumentSerde
}

// A collection of the imported DocumentationPart identifiers.
type ImportDocumentationPartsOutput struct {

	// A list of the returned documentation part identifiers.
	Ids []string

	// A list of warning messages reported during import of documentation parts.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportDocumentationPartsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportDocumentationParts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportDocumentationParts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportDocumentationParts"); err != nil {
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
	if err = addOpImportDocumentationPartsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportDocumentationParts(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opImportDocumentationParts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportDocumentationParts",
	}
}
