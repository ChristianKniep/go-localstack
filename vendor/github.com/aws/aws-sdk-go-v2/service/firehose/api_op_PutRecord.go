// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Writes a single data record into an Amazon Firehose delivery stream. To write
// multiple data records into a delivery stream, use PutRecordBatch. Applications using these
// operations are referred to as producers.
//
// By default, each delivery stream can take in up to 2,000 transactions per
// second, 5,000 records per second, or 5 MB per second. If you use PutRecordand PutRecordBatch, the
// limits are an aggregate across these two operations for each delivery stream.
// For more information about limits and how to request an increase, see [Amazon Firehose Limits].
//
// Firehose accumulates and publishes a particular metric for a customer account
// in one minute intervals. It is possible that the bursts of incoming
// bytes/records ingested to a delivery stream last only for a few seconds. Due to
// this, the actual spikes in the traffic might not be fully visible in the
// customer's 1 minute CloudWatch metrics.
//
// You must specify the name of the delivery stream and the data record when using PutRecord
// . The data record consists of a data blob that can be up to 1,000 KiB in size,
// and any kind of data. For example, it can be a segment from a log file,
// geographic location data, website clickstream data, and so on.
//
// Firehose buffers records before delivering them to the destination. To
// disambiguate the data blobs at the destination, a common solution is to use
// delimiters in the data, such as a newline ( \n ) or some other character unique
// within the data. This allows the consumer application to parse individual data
// items when reading the data from the destination.
//
// The PutRecord operation returns a RecordId , which is a unique string assigned
// to each record. Producer applications can use this ID for purposes such as
// auditability and investigation.
//
// If the PutRecord operation throws a ServiceUnavailableException , the API is
// automatically reinvoked (retried) 3 times. If the exception persists, it is
// possible that the throughput limits have been exceeded for the delivery stream.
//
// Re-invoking the Put API operations (for example, PutRecord and PutRecordBatch)
// can result in data duplicates. For larger data assets, allow for a longer time
// out before retrying Put API operations.
//
// Data records sent to Firehose are stored for 24 hours from the time they are
// added to a delivery stream as it tries to send the records to the destination.
// If the destination is unreachable for more than 24 hours, the data is no longer
// available.
//
// Don't concatenate two or more base64 strings to form the data fields of your
// records. Instead, concatenate the raw data, then perform base64 encoding.
//
// [Amazon Firehose Limits]: https://docs.aws.amazon.com/firehose/latest/dev/limits.html
func (c *Client) PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) {
	if params == nil {
		params = &PutRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecord", params, optFns, c.addOperationPutRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecordInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// The record.
	//
	// This member is required.
	Record *types.Record

	noSmithyDocumentSerde
}

type PutRecordOutput struct {

	// The ID of the record.
	//
	// This member is required.
	RecordId *string

	// Indicates whether server-side encryption (SSE) was enabled during this
	// operation.
	Encrypted *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecord{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRecord"); err != nil {
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
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpPutRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecord(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRecord",
	}
}