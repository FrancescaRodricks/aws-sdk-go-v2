// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the status of importing resources to an application version.
//
// If you get a 404 error with ResourceImportStatusNotFoundAppMetadataException ,
// you must call importResourcesToDraftAppVersion after creating the application
// and before calling describeDraftAppVersionResourcesImportStatus to obtain the
// status.
func (c *Client) DescribeDraftAppVersionResourcesImportStatus(ctx context.Context, params *DescribeDraftAppVersionResourcesImportStatusInput, optFns ...func(*Options)) (*DescribeDraftAppVersionResourcesImportStatusOutput, error) {
	if params == nil {
		params = &DescribeDraftAppVersionResourcesImportStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDraftAppVersionResourcesImportStatus", params, optFns, c.addOperationDescribeDraftAppVersionResourcesImportStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDraftAppVersionResourcesImportStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDraftAppVersionResourcesImportStatusInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference
	// guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	AppArn *string

	noSmithyDocumentSerde
}

type DescribeDraftAppVersionResourcesImportStatusOutput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference
	// guide.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Status of the action.
	//
	// This member is required.
	Status types.ResourceImportStatusType

	// The time when the status last changed.
	//
	// This member is required.
	StatusChangeTime *time.Time

	// The returned error message for the request.
	ErrorMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDraftAppVersionResourcesImportStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDraftAppVersionResourcesImportStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDraftAppVersionResourcesImportStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDraftAppVersionResourcesImportStatus"); err != nil {
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
	if err = addOpDescribeDraftAppVersionResourcesImportStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDraftAppVersionResourcesImportStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDraftAppVersionResourcesImportStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDraftAppVersionResourcesImportStatus",
	}
}
