// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a report generator.
//
// After you make changes to a report generator, it starts generating new reports
// within 60 minutes of being updated.
func (c *Client) UpdateLicenseManagerReportGenerator(ctx context.Context, params *UpdateLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*UpdateLicenseManagerReportGeneratorOutput, error) {
	if params == nil {
		params = &UpdateLicenseManagerReportGeneratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLicenseManagerReportGenerator", params, optFns, c.addOperationUpdateLicenseManagerReportGeneratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLicenseManagerReportGeneratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLicenseManagerReportGeneratorInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// Amazon Resource Name (ARN) of the report generator to update.
	//
	// This member is required.
	LicenseManagerReportGeneratorArn *string

	// The report context.
	//
	// This member is required.
	ReportContext *types.ReportContext

	// Frequency by which reports are generated.
	//
	// This member is required.
	ReportFrequency *types.ReportFrequency

	// Name of the report generator.
	//
	// This member is required.
	ReportGeneratorName *string

	// Type of reports to generate. The following report types are supported:
	//
	//   - License configuration report - Reports the number and details of consumed
	//   licenses for a license configuration.
	//
	//   - Resource report - Reports the tracked licenses and resource consumption for
	//   a license configuration.
	//
	// This member is required.
	Type []types.ReportType

	// Description of the report generator.
	Description *string

	noSmithyDocumentSerde
}

type UpdateLicenseManagerReportGeneratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLicenseManagerReportGeneratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLicenseManagerReportGenerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLicenseManagerReportGenerator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLicenseManagerReportGenerator"); err != nil {
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
	if err = addOpUpdateLicenseManagerReportGeneratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLicenseManagerReportGenerator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLicenseManagerReportGenerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLicenseManagerReportGenerator",
	}
}
