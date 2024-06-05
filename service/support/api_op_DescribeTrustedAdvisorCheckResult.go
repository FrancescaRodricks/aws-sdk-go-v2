// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the results of the Trusted Advisor check that has the specified check
// ID. You can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// The response contains a TrustedAdvisorCheckResult object, which contains these three objects:
//
// # TrustedAdvisorCategorySpecificSummary
//
// # TrustedAdvisorResourceDetail
//
// # TrustedAdvisorResourcesSummary
//
// In addition, the response contains these fields:
//
//   - status - The alert status of the check can be ok (green), warning (yellow),
//     error (red), or not_available .
//
//   - timestamp - The time of the last refresh of the check.
//
//   - checkId - The unique identifier for the check.
//
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func (c *Client) DescribeTrustedAdvisorCheckResult(ctx context.Context, params *DescribeTrustedAdvisorCheckResultInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckResultOutput, error) {
	if params == nil {
		params = &DescribeTrustedAdvisorCheckResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrustedAdvisorCheckResult", params, optFns, c.addOperationDescribeTrustedAdvisorCheckResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustedAdvisorCheckResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorCheckResultInput struct {

	// The unique identifier for the Trusted Advisor check.
	//
	// This member is required.
	CheckId *string

	// The ISO 639-1 code for the language that you want your check results to appear
	// in.
	//
	// The Amazon Web Services Support API currently supports the following languages
	// for Trusted Advisor:
	//
	//   - Chinese, Simplified - zh
	//
	//   - Chinese, Traditional - zh_TW
	//
	//   - English - en
	//
	//   - French - fr
	//
	//   - German - de
	//
	//   - Indonesian - id
	//
	//   - Italian - it
	//
	//   - Japanese - ja
	//
	//   - Korean - ko
	//
	//   - Portuguese, Brazilian - pt_BR
	//
	//   - Spanish - es
	Language *string

	noSmithyDocumentSerde
}

// The result of the Trusted Advisor check returned by the DescribeTrustedAdvisorCheckResult operation.
type DescribeTrustedAdvisorCheckResultOutput struct {

	// The detailed results of the Trusted Advisor check.
	Result *types.TrustedAdvisorCheckResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrustedAdvisorCheckResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorCheckResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTrustedAdvisorCheckResult"); err != nil {
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
	if err = addOpDescribeTrustedAdvisorCheckResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckResult(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTrustedAdvisorCheckResult",
	}
}
