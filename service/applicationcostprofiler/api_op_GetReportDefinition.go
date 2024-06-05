// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationcostprofiler

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the definition of a report already configured in AWS Application Cost
// Profiler.
func (c *Client) GetReportDefinition(ctx context.Context, params *GetReportDefinitionInput, optFns ...func(*Options)) (*GetReportDefinitionOutput, error) {
	if params == nil {
		params = &GetReportDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReportDefinition", params, optFns, c.addOperationGetReportDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReportDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReportDefinitionInput struct {

	// ID of the report to retrieve.
	//
	// This member is required.
	ReportId *string

	noSmithyDocumentSerde
}

type GetReportDefinitionOutput struct {

	// Timestamp (milliseconds) when this report definition was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Amazon Simple Storage Service (Amazon S3) location where the report is uploaded.
	//
	// This member is required.
	DestinationS3Location *types.S3Location

	// Format of the generated report.
	//
	// This member is required.
	Format types.Format

	// Timestamp (milliseconds) when this report definition was last updated.
	//
	// This member is required.
	LastUpdated *time.Time

	// Description of the report.
	//
	// This member is required.
	ReportDescription *string

	// Cadence used to generate the report.
	//
	// This member is required.
	ReportFrequency types.ReportFrequency

	// ID of the report retrieved.
	//
	// This member is required.
	ReportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReportDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReportDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReportDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReportDefinition"); err != nil {
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
	if err = addOpGetReportDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReportDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReportDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReportDefinition",
	}
}
