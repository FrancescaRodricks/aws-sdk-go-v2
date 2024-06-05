// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets metadata information about the test execution.
func (c *Client) DescribeTestExecution(ctx context.Context, params *DescribeTestExecutionInput, optFns ...func(*Options)) (*DescribeTestExecutionOutput, error) {
	if params == nil {
		params = &DescribeTestExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTestExecution", params, optFns, c.addOperationDescribeTestExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTestExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTestExecutionInput struct {

	// The execution Id of the test set execution.
	//
	// This member is required.
	TestExecutionId *string

	noSmithyDocumentSerde
}

type DescribeTestExecutionOutput struct {

	// Indicates whether we use streaming or non-streaming APIs are used for the test
	// set execution. For streaming, StartConversation Amazon Lex Runtime API is used.
	// Whereas for non-streaming, RecognizeUtterance and RecognizeText Amazon Lex
	// Runtime API is used.
	ApiMode types.TestExecutionApiMode

	// The execution creation date and time for the test set execution.
	CreationDateTime *time.Time

	// Reasons for the failure of the test set execution.
	FailureReasons []string

	// The date and time of the last update for the execution.
	LastUpdatedDateTime *time.Time

	// The target bot for the test set execution details.
	Target *types.TestExecutionTarget

	// The execution Id for the test set execution.
	TestExecutionId *string

	// Indicates whether test set is audio or text.
	TestExecutionModality types.TestExecutionModality

	// The test execution status for the test execution.
	TestExecutionStatus types.TestExecutionStatus

	// The test set Id for the test set execution.
	TestSetId *string

	// The test set name of the test set execution.
	TestSetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTestExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTestExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTestExecution"); err != nil {
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
	if err = addOpDescribeTestExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTestExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTestExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTestExecution",
	}
}
