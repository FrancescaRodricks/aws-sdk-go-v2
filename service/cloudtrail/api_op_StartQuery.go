// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a CloudTrail Lake query. Use the QueryStatement parameter to provide
// your SQL query, enclosed in single quotation marks. Use the optional
// DeliveryS3Uri parameter to deliver the query results to an S3 bucket.
//
// StartQuery requires you specify either the QueryStatement parameter, or a
// QueryAlias and any QueryParameters . In the current release, the QueryAlias and
// QueryParameters parameters are used only for the queries that populate the
// CloudTrail Lake dashboards.
func (c *Client) StartQuery(ctx context.Context, params *StartQueryInput, optFns ...func(*Options)) (*StartQueryOutput, error) {
	if params == nil {
		params = &StartQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartQuery", params, optFns, c.addOperationStartQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQueryInput struct {

	//  The URI for the S3 bucket where CloudTrail delivers the query results.
	DeliveryS3Uri *string

	//  The alias that identifies a query template.
	QueryAlias *string

	//  The query parameters for the specified QueryAlias .
	QueryParameters []string

	// The SQL code of your query.
	QueryStatement *string

	noSmithyDocumentSerde
}

type StartQueryOutput struct {

	// The ID of the started query.
	QueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartQuery{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartQuery"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartQuery",
	}
}
