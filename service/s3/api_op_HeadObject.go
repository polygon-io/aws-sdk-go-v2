// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The HEAD operation retrieves metadata from an object without returning the
// object itself. This operation is useful if you're only interested in an object's
// metadata. To use HEAD, you must have READ access to the object. A HEAD request
// has the same options as a GET operation on an object. The response is identical
// to the GET response except that there is no response body. If you encrypt an
// object by using server-side encryption with customer-provided encryption keys
// (SSE-C) when you store the object in Amazon S3, then when you retrieve the
// metadata from the object, you must use the following headers:
//
// *
// x-amz-server-side-encryption-customer-algorithm
//
// *
// x-amz-server-side-encryption-customer-key
//
// *
// x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C,
// see Server-Side Encryption (Using Customer-Provided Encryption Keys)
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html).
// Encryption request headers, like x-amz-server-side-encryption, should not be
// sent for GET requests if your object uses server-side encryption with CMKs
// stored in AWS KMS (SSE-KMS) or server-side encryption with Amazon S3–managed
// encryption keys (SSE-S3). If your object does use these types of keys, you’ll
// get an HTTP 400 BadRequest error. Request headers are limited to 8 KB in size.
// For more information, see Common Request Headers
// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonRequestHeaders.html).
// Consider the following when using request headers:
//
// * Consideration 1 – If both
// of the If-Match and If-Unmodified-Since headers are present in the request as
// follows:
//
// * If-Match condition evaluates to true, and;
//
// * If-Unmodified-Since
// condition evaluates to false;
//
// Then Amazon S3 returns 200 OK and the data
// requested.
//
// * Consideration 2 – If both of the If-None-Match and
// If-Modified-Since headers are present in the request as follows:
//
// *
// If-None-Match condition evaluates to false, and;
//
// * If-Modified-Since condition
// evaluates to true;
//
// Then Amazon S3 returns the 304 Not Modified response
// code.
//
// For more information about conditional requests, see RFC 7232
// (https://tools.ietf.org/html/rfc7232). Permissions You need the s3:GetObject
// permission for this operation. For more information, see Specifying Permissions
// in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html). If
// the object you request does not exist, the error Amazon S3 returns depends on
// whether you also have the s3:ListBucket permission.
//
// * If you have the
// s3:ListBucket permission on the bucket, Amazon S3 returns an HTTP status code
// 404 ("no such key") error.
//
// * If you don’t have the s3:ListBucket permission,
// Amazon S3 returns an HTTP status code 403 ("access denied") error.
//
// The
// following operation is related to HeadObject:
//
// * GetObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
func (c *Client) HeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*Options)) (*HeadObjectOutput, error) {
	if params == nil {
		params = &HeadObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HeadObject", params, optFns, addOperationHeadObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HeadObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HeadObjectInput struct {

	// The name of the bucket containing the object. When using this API with an access
	// point, you must direct requests to the access point hostname. The access point
	// hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation with an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide. When using this API with
	// Amazon S3 on Outposts, you must direct requests to the S3 on Outposts hostname.
	// The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this operation using S3 on Outposts through the AWS SDKs, you provide the
	// Outposts bucket ARN in place of the bucket name. For more information about S3
	// on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// The object key.
	//
	// This member is required.
	Key *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Return the object only if its entity tag (ETag) is the same as the one
	// specified, otherwise return a 412 (precondition failed).
	IfMatch *string

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time

	// Part number of the object being read. This is a positive integer between 1 and
	// 10,000. Effectively performs a 'ranged' HEAD request for the part specified.
	// Useful querying about the size of the part and the number of parts in this
	// object.
	PartNumber *int32

	// Downloads the specified range bytes of an object. For more information about the
	// HTTP Range header, see
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35). Amazon S3
	// doesn't support retrieving multiple ranges of data per GET request.
	Range *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// VersionId used to reference a specific version of the object.
	VersionId *string
}

type HeadObjectOutput struct {

	// Indicates that a range of bytes was specified.
	AcceptRanges *string

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Size of the body in bytes.
	ContentLength *int64

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Last modified date of the object
	LastModified *time.Time

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP, you
	// can create metadata whose values are not legal HTTP headers.
	MissingMeta *int32

	// Specifies whether a legal hold is in effect for this object. This header is only
	// returned if the requester has the s3:GetObjectLegalHold permission. This header
	// is not returned if the specified version of this object has never had a legal
	// hold applied. For more information about S3 Object Lock, see Object Lock
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode, if any, that's in effect for this object. This header is
	// only returned if the requester has the s3:GetObjectRetention permission. For
	// more information about S3 Object Lock, see Object Lock
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockMode types.ObjectLockMode

	// The date and time when the Object Lock retention period expires. This header is
	// only returned if the requester has the s3:GetObjectRetention permission.
	ObjectLockRetainUntilDate *time.Time

	// The count of parts this object has.
	PartsCount *int32

	// Amazon S3 can return this header if your request involves a bucket that is
	// either a source or destination in a replication rule. In replication, you have a
	// source bucket on which you configure replication and destination bucket where
	// Amazon S3 stores object replicas. When you request an object (GetObject) or
	// object metadata (HeadObject) from these buckets, Amazon S3 will return the
	// x-amz-replication-status header in the response as follows:
	//
	// * If requesting an
	// object from the source bucket — Amazon S3 will return the
	// x-amz-replication-status header if the object in your request is eligible for
	// replication. For example, suppose that in your replication configuration, you
	// specify object prefix TaxDocs requesting Amazon S3 to replicate objects with key
	// prefix TaxDocs. Any objects you upload with this key name prefix, for example
	// TaxDocs/document1.pdf, are eligible for replication. For any object request with
	// this key name prefix, Amazon S3 will return the x-amz-replication-status header
	// with value PENDING, COMPLETED or FAILED indicating object replication status.
	//
	// *
	// If requesting an object from the destination bucket — Amazon S3 will return the
	// x-amz-replication-status header with value REPLICA if the object in your request
	// is a replica that Amazon S3 created.
	//
	// For more information, see Replication
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).
	ReplicationStatus types.ReplicationStatus

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If the object is an archived object (an object whose storage class is GLACIER),
	// the response includes this header if either the archive restoration is in
	// progress (see RestoreObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html) or an
	// archive copy is already restored. If an archive copy is already restored, the
	// header value indicates when Amazon S3 is scheduled to delete the object copy.
	// For example: x-amz-restore: ongoing-request="false", expiry-date="Fri, 23 Dec
	// 2012 00:00:00 GMT" If the object restoration is in progress, the header returns
	// the value ongoing-request="true". For more information about archiving objects,
	// see Transitioning Objects: General Considerations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-transition-general-considerations).
	Restore *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// If the object is stored using server-side encryption either with an AWS KMS
	// customer master key (CMK) or an Amazon S3-managed encryption key, the response
	// includes this header with the value of the server-side encryption algorithm used
	// when storing this object in Amazon S3 (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects. For more
	// information, see Storage Classes
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html).
	StorageClass types.StorageClass

	// Version of the object.
	VersionId *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHeadObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHeadObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHeadObject{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpHeadObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHeadObject(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addHeadObjectUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opHeadObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "HeadObject",
	}
}

// getHeadObjectBucketMember returns a pointer to string denoting a provided bucket
// member valueand a boolean indicating if the input has a modeled bucket name,
func getHeadObjectBucketMember(input interface{}) (*string, bool) {
	in := input.(*HeadObjectInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addHeadObjectUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getHeadObjectBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
