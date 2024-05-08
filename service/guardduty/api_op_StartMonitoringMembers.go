// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Turns on GuardDuty monitoring of the specified member accounts. Use this
// operation to restart monitoring of accounts that you stopped monitoring with the
// [StopMonitoringMembers]operation.
//
// [StopMonitoringMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_StopMonitoringMembers.html
func (c *Client) StartMonitoringMembers(ctx context.Context, params *StartMonitoringMembersInput, optFns ...func(*Options)) (*StartMonitoringMembersOutput, error) {
	if params == nil {
		params = &StartMonitoringMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMonitoringMembers", params, optFns, c.addOperationStartMonitoringMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMonitoringMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMonitoringMembersInput struct {

	// A list of account IDs of the GuardDuty member accounts to start monitoring.
	//
	// This member is required.
	AccountIds []string

	// The unique ID of the detector of the GuardDuty administrator account associated
	// with the member accounts to monitor.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type StartMonitoringMembersOutput struct {

	// A list of objects that contain the unprocessed account and a result string that
	// explains why it was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMonitoringMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMonitoringMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMonitoringMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMonitoringMembers"); err != nil {
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
	if err = addOpStartMonitoringMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMonitoringMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMonitoringMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMonitoringMembers",
	}
}
