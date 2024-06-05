// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of data deletion jobs for a dataset group ordered by creation
// time, with the most recent first. When a dataset group is not specified, all the
// data deletion jobs associated with the account are listed. The response provides
// the properties for each job, including the Amazon Resource Name (ARN). For more
// information on data deletion jobs, see [Deleting users].
//
// [Deleting users]: https://docs.aws.amazon.com/personalize/latest/dg/delete-records.html
func (c *Client) ListDataDeletionJobs(ctx context.Context, params *ListDataDeletionJobsInput, optFns ...func(*Options)) (*ListDataDeletionJobsOutput, error) {
	if params == nil {
		params = &ListDataDeletionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataDeletionJobs", params, optFns, c.addOperationListDataDeletionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataDeletionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataDeletionJobsInput struct {

	// The Amazon Resource Name (ARN) of the dataset group to list data deletion jobs
	// for.
	DatasetGroupArn *string

	// The maximum number of data deletion jobs to return.
	MaxResults *int32

	// A token returned from the previous call to ListDataDeletionJobs for getting the
	// next set of jobs (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataDeletionJobsOutput struct {

	// The list of data deletion jobs.
	DataDeletionJobs []types.DataDeletionJobSummary

	// A token for getting the next set of data deletion jobs (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataDeletionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataDeletionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataDeletionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataDeletionJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataDeletionJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDataDeletionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataDeletionJobs",
	}
}
