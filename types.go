package s3

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Config contains the available options for configuring a Client.
type Config struct {
	// S3 Access key ID
	AccessKey string
	// S3 Secret Access key
	SecretKey string
	// S3 region
	Region string
	// Endpoint is URL to the s3 service.
	Endpoint string
}

// Client provides an interface for interacting with the S3 API.
type Client struct {
	config      Config
	endpointURL string
	httpClient  *http.Client
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html#AmazonS3-CreateMultipartUpload-response-CreateMultipartUploadOutput
type InitiateMultipartUploadResult struct {
	XMLName  xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	UploadId string   `xml:"UploadId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompletedPart.html
type CompletedPart struct {
	XMLName           xml.Name `xml:"Part"`
	ChecksumCRC32     string   `xml:"ChecksumCRC32"`
	ChecksumCRC32C    string   `xml:"ChecksumCRC32C"`
	ChecksumCRC64NVME string   `xml:"ChecksumCRC64NVME"`
	ChecksumSHA1      string   `xml:"ChecksumSHA1"`
	ChecksumSHA256    string   `xml:"ChecksumSHA256"`
	ETag              string   `xml:"ETag"`
	PartNumber        int      `xml:"PartNumber"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompletedMultipartUpload.html
type CompletedMultipartUpload struct {
	XmlName xml.Name        `xml:"CompleteMultipartUpload"`
	Parts   []CompletedPart `xml:"Part"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html#API_ListBuckets_ResponseSyntax
type ListBucketsResponse struct {
	Buckets []BucketInfo `xml:"Buckets>Bucket"`
	Owner   Owner
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Bucket.html
type BucketInfo struct {
	Name         string
	CreationDate time.Time
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html#API_ListObjects_ResponseSyntax
type ListObjectsResponse struct {
	CommonPrefixes []CommonPrefix
	Contents       []ObjectInfo
	Delimiter      string
	EncodingType   string
	IsTruncated    bool
	Marker         string
	MaxKeys        int
	Name           string
	NextMarker     string
	Prefix         string
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CommonPrefix.html
type CommonPrefix struct {
	Prefix string `xml:"Prefix"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Object.html
type ObjectInfo struct {
	Key          string
	ETag         string
	Size         int
	LastModified time.Time
	StorageClass string
	Owner        Owner
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
type ErrorResponse struct {
	Code      string `xml:"Code"`
	Message   string `xml:"Message"`
	Resource  string `xml:"Resource"`
	RequestID string `xml:"RequestId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html#AmazonS3-ListMultipartUploads-response-ListMultipartUploadsOutput
type ListMultipartUploadsResult struct {
	XMLName            xml.Name          `xml:"ListMultipartUploadsResult"`
	Bucket             string            `xml:"Bucket"`
	KeyMarker          string            `xml:"KeyMarker"`
	UploadIdMarker     string            `xml:"UploadIdMarker"`
	NextKeyMarker      string            `xml:"NextKeyMarker"`
	Prefix             string            `xml:"Prefix"`
	Delimiter          string            `xml:"Delimiter"`
	NextUploadIdMarker string            `xml:"NextUploadIdMarker"`
	MaxUploads         int               `xml:"MaxUploads"`
	IsTruncated        bool              `xml:"IsTruncated"`
	Uploads            []MultipartUpload `xml:"Upload"`
	CommonPrefixes     []CommonPrefix    `xml:"CommonPrefixes>Prefix"`
	EncodingType       string            `xml:"EncodingType"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_MultipartUpload.html
type MultipartUpload struct {
	XMLName           xml.Name   `xml:"Upload"`
	ChecksumAlgorithm string     `xml:"ChecksumAlgorithm"`
	ChecksumType      string     `xml:"ChecksumType"`
	Initiated         string     `xml:"Initiated"`
	Initiator         *Initiator `xml:"Initiator"`
	Key               string     `xml:"Key"`
	Owner             *Initiator `xml:"Owner"`
	StorageClass      string     `xml:"StorageClass"`
	UploadId          string     `xml:"UploadId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Initiator.html
type Initiator struct {
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html#AmazonS3-ListParts-response-ListPartsOutput
type ListPartsResult struct {
	XMLName              xml.Name   `xml:"ListPartsResult"`
	Bucket               string     `xml:"Bucket"`
	Key                  string     `xml:"Key"`
	UploadId             string     `xml:"UploadId"`
	PartNumberMarker     int        `xml:"PartNumberMarker"`
	NextPartNumberMarker int        `xml:"NextPartNumberMarker"`
	MaxParts             int        `xml:"MaxParts"`
	IsTruncated          bool       `xml:"IsTruncated"`
	Parts                []Part     `xml:"Part"`
	Initiator            *Initiator `xml:"Initiator"`
	Owner                *Initiator `xml:"Owner"`
	StorageClass         string     `xml:"StorageClass"`
	ChecksumAlgorithm    string     `xml:"ChecksumAlgorithm"`
	ChecksumType         string     `xml:"ChecksumType"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Part.html
type Part struct {
	ChecksumCRC32     string `xml:"ChecksumCRC32"`
	ChecksumCRC32C    string `xml:"ChecksumCRC32C"`
	ChecksumCRC64NVME string `xml:"ChecksumCRC64NVME"`
	ChecksumSHA1      string `xml:"ChecksumSHA1"`
	ChecksumSHA256    string `xml:"ChecksumSHA256"`
	ETag              string `xml:"ETag"`
	LastModified      string `xml:"LastModified"`
	PartNumber        int    `xml:"PartNumber"`
	Size              int64  `xml:"Size"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ObjectPart.html
type ObjectPart struct {
	ChecksumCRC32     string `xml:"ChecksumCRC32"`
	ChecksumCRC32C    string `xml:"ChecksumCRC32C"`
	ChecksumCRC64NVME string `xml:"ChecksumCRC64NVME"`
	ChecksumSHA1      string `xml:"ChecksumSHA1"`
	ChecksumSHA256    string `xml:"ChecksumSHA256"`
	ETag              string `xml:"ETag"`
	LastModified      string `xml:"LastModified"`
	PartNumber        int    `xml:"PartNumber"`
	Size              int64  `xml:"Size"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html#AmazonS3-GetObjectAttributes-response-GetObjectAttributesOutput
type GetObjectAttributesResponse struct {
	XMLName              xml.Name                 `xml:"GetObjectAttributesResponse"`
	ETag                 string                   `xml:"ETag"`
	Checksum             Checksum                 `xml:"Checksum"`
	ObjectAttributeParts GetObjectAttributesParts `xml:"ObjectParts"`
	StorageClass         string                   `xml:"StorageClass"`
	ObjectSize           int64                    `xml:"ObjectSize"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Checksum.html
type Checksum struct {
	ChecksumCRC32     string `xml:"ChecksumCRC32"`
	ChecksumCRC32C    string `xml:"ChecksumCRC32C"`
	ChecksumCRC64NVME string `xml:"ChecksumCRC64NVME"`
	ChecksumSHA1      string `xml:"ChecksumSHA1"`
	ChecksumSHA256    string `xml:"ChecksumSHA256"`
	ChecksumType      string `xml:"ChecksumType"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributesParts.html
type GetObjectAttributesParts struct {
	IsTruncated          bool         `xml:"IsTruncated"`
	MaxParts             int          `xml:"MaxParts"`
	NextPartNumberMarker int          `xml:"NextPartNumberMarker"`
	PartNumberMarker     int          `xml:"PartNumberMarker"`
	Parts                []ObjectPart `xml:"Part"`
	PartsCount           int          `xml:"PartsCount"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Tagging.html
type Tagging struct {
	XMLName xml.Name `xml:"Tagging"`
	TagSet  TagSet   `xml:"TagSet"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html#AmazonS3-GetObjectTagging-response-TagSet
type TagSet struct {
	Tags []Tag `xml:"Tag"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Tag.html
type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListDirectoryBuckets.html#AmazonS3-ListDirectoryBuckets-response-ListDirectoryBucketsOutput
type ListAllMyDirectoryBucketsResult struct {
	XMLName           xml.Name `xml:"ListAllMyDirectoryBucketsResult"`
	Buckets           []Bucket `xml:"Buckets>Bucket"`
	ContinuationToken string   `xml:"ContinuationToken"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Bucket.html
type Bucket struct {
	BucketRegion string    `xml:"BucketRegion"`
	CreationDate time.Time `xml:"CreationDate"`
	Name         string    `xml:"Name"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_WebsiteConfiguration.html
type WebsiteConfiguration struct {
	XMLName               xml.Name               `xml:"WebsiteConfiguration"`
	RedirectAllRequestsTo *RedirectAllRequestsTo `xml:"RedirectAllRequestsTo"`
	IndexDocument         *IndexDocument         `xml:"IndexDocument"`
	ErrorDocument         *ErrorDocument         `xml:"ErrorDocument"`
	RoutingRules          []RoutingRule          `xml:"RoutingRules>RoutingRule"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_RedirectAllRequestsTo.html
type RedirectAllRequestsTo struct {
	HostName string `xml:"HostName"`
	Protocol string `xml:"Protocol"` // optional
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_IndexDocument.html
type IndexDocument struct {
	Suffix string `xml:"Suffix"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ErrorDocument.html
type ErrorDocument struct {
	Key string `xml:"Key"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_RoutingRule.html
type RoutingRule struct {
	Condition *Condition `xml:"Condition"`
	Redirect  Redirect   `xml:"Redirect"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Condition.html
type Condition struct {
	HttpErrorCodeReturnedEquals string `xml:"HttpErrorCodeReturnedEquals"`
	KeyPrefixEquals             string `xml:"KeyPrefixEquals"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Redirect.html
type Redirect struct {
	HostName             string `xml:"HostName"`
	HttpRedirectCode     string `xml:"HttpRedirectCode"`
	Protocol             string `xml:"Protocol"`
	ReplaceKeyPrefixWith string `xml:"ReplaceKeyPrefixWith"`
	ReplaceKeyWith       string `xml:"ReplaceKeyWith"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Delete.html
type Delete struct {
	XMLName xml.Name           `xml:"Delete"`
	Objects []ObjectIdentifier `xml:"Object"`
	Quiet   bool               `xml:"Quiet"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ObjectIdentifier.html
type ObjectIdentifier struct {
	ETag             string    `xml:"ETag"`
	Key              string    `xml:"Key"`
	LastModifiedTime time.Time `xml:"LastModifiedTime"`
	Size             int64     `xml:"Size"`
	VersionId        string    `xml:"VersionId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjects.html#AmazonS3-DeleteObjects-response-DeleteObjectsOutput
type DeleteResult struct {
	XMLName xml.Name        `xml:"DeleteResult"`
	Deleted []DeletedObject `xml:"Deleted"`
	Errors  []Error         `xml:"Error"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletedObject.html
type DeletedObject struct {
	DeleteMarker          bool   `xml:"DeleteMarker"`
	DeleteMarkerVersionId string `xml:"DeleteMarkerVersionId"`
	Key                   string `xml:"Key"`
	VersionId             string `xml:"VersionId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Error.html
type Error struct {
	Code      string `xml:"Code"`
	Key       string `xml:"Key"`
	Message   string `xml:"Message"`
	VersionId string `xml:"VersionId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectVersions.html#AmazonS3-ListObjectVersions-response-ListObjectVersionsOutput
type ListVersionsResult struct {
	XMLName             xml.Name            `xml:"ListVersionsResult"`
	IsTruncated         bool                `xml:"IsTruncated"`
	KeyMarker           string              `xml:"KeyMarker"`
	VersionIdMarker     string              `xml:"VersionIdMarker"`
	NextKeyMarker       string              `xml:"NextKeyMarker"`
	NextVersionIdMarker string              `xml:"NextVersionIdMarker"`
	Versions            []ObjectVersion     `xml:"Version"`
	DeleteMarkers       []DeleteMarkerEntry `xml:"DeleteMarker"`
	Name                string              `xml:"Name"`
	Prefix              string              `xml:"Prefix"`
	Delimiter           string              `xml:"Delimiter"`
	MaxKeys             int                 `xml:"MaxKeys"`
	CommonPrefixes      []CommonPrefix      `xml:"CommonPrefixes"`
	EncodingType        string              `xml:"EncodingType"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ObjectVersion.html
type ObjectVersion struct {
	ChecksumAlgorithm []string       `xml:"ChecksumAlgorithm"`
	ChecksumType      string         `xml:"ChecksumType"`
	ETag              string         `xml:"ETag"`
	IsLatest          bool           `xml:"IsLatest"`
	Key               string         `xml:"Key"`
	LastModified      time.Time      `xml:"LastModified"`
	Owner             *Owner         `xml:"Owner"`
	RestoreStatus     *RestoreStatus `xml:"RestoreStatus"`
	Size              int64          `xml:"Size"`
	StorageClass      string         `xml:"StorageClass"`
	VersionId         string         `xml:"VersionId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteMarkerEntry.html
type DeleteMarkerEntry struct {
	IsLatest     bool      `xml:"IsLatest"`
	Key          string    `xml:"Key"`
	LastModified time.Time `xml:"LastModified"`
	Owner        *Owner    `xml:"Owner"`
	VersionId    string    `xml:"VersionId"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Owner.html
type Owner struct {
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreStatus.html
type RestoreStatus struct {
	IsRestoreInProgress bool      `xml:"IsRestoreInProgress"`
	RestoreExpiryDate   time.Time `xml:"RestoreExpiryDate"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html#AmazonS3-GetBucketVersioning-response-GetBucketVersioningOutput
type VersioningConfiguration struct {
	XMLName   xml.Name `xml:"VersioningConfiguration"`
	Status    string   `xml:"Status"`
	MfaDelete string   `xml:"MfaDelete"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html#AmazonS3-GetObjectLockConfiguration-response-ObjectLockConfigurationhttps://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html#AmazonS3-GetObjectLockConfiguration-response-ObjectLockConfiguration
type ObjectLockConfiguration struct {
	XMLName           xml.Name        `xml:"ObjectLockConfiguration"`
	ObjectLockEnabled string          `xml:"ObjectLockEnabled"`
	Rule              *ObjectLockRule `xml:"Rule"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ObjectLockRule.html
type ObjectLockRule struct {
	DefaultRetention DefaultRetention `xml:"DefaultRetention"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_DefaultRetention.html
type DefaultRetention struct {
	Days  int    `xml:"Days"`
	Mode  string `xml:"Mode"`
	Years int    `xml:"Years"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectRetention.html#API_PutObjectRetention_RequestSyntax
type Retention struct {
	XMLName         xml.Name `xml:"Retention"`
	Mode            string   `xml:"Mode"`            // e.g., "GOVERNANCE" or "COMPLIANCE"
	RetainUntilDate string   `xml:"RetainUntilDate"` // ISO8601 timestamp, e.g., "2025-12-31T00:00:00Z"
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html#API_GetPublicAccessBlock_ResponseSyntax
type PublicAccessBlockConfiguration struct {
	XMLName               xml.Name `xml:"PublicAccessBlockConfiguration"`
	Xmlns                 string   `xml:"xmlns,attr"`
	BlockPublicAcls       bool     `xml:"BlockPublicAcls"`
	IgnorePublicAcls      bool     `xml:"IgnorePublicAcls"`
	BlockPublicPolicy     bool     `xml:"BlockPublicPolicy"`
	RestrictPublicBuckets bool     `xml:"RestrictPublicBuckets"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html#AmazonS3-GetBucketLogging-response-GetBucketLoggingOutput
type BucketLoggingStatus struct {
	XMLName        xml.Name        `xml:"BucketLoggingStatus"`
	Xmlns          string          `xml:"xmlns,attr"`
	LoggingEnabled *LoggingEnabled `xml:"LoggingEnabled"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_LoggingEnabled.html
type LoggingEnabled struct {
	TargetBucket          string                 `xml:"TargetBucket"`
	TargetPrefix          string                 `xml:"TargetPrefix"`
	TargetGrants          *TargetGrants          `xml:"TargetGrants"`
	TargetObjectKeyFormat *TargetObjectKeyFormat `xml:"TargetObjectKeyFormat"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_TargetGrant.html
type TargetGrants struct {
	Grants []Grant `xml:"Grant"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_TargetObjectKeyFormat.html
type TargetObjectKeyFormat struct {
	PartitionedPrefix *PartitionedPrefix `xml:"PartitionedPrefix"`
	SimplePrefix      *SimplePrefix      `xml:"SimplePrefix"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_PartitionedPrefix.html
type PartitionedPrefix struct {
	PartitionDateSource string `xml:"PartitionDateSource"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html#AmazonS3-GetObjectAcl-response-GetObjectAclOutput
type AccessControlPolicy struct {
	XMLName           xml.Name          `xml:"AccessControlPolicy"`
	Xmlns             string            `xml:"xmlns,attr"`
	Owner             Owner             `xml:"Owner"`
	AccessControlList AccessControlList `xml:"AccessControlList"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html#AmazonS3-GetObjectAcl-response-Grants
type AccessControlList struct {
	Grants []Grant `xml:"Grant"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html
type Grant struct {
	Grantee    Grantee `xml:"Grantee"`
	Permission string  `xml:"Permission"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_SimplePrefix.html
type SimplePrefix struct{}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grantee.html
type Grantee struct {
	XMLName      xml.Name `xml:"Grantee"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XsiType      string   `xml:"xsi:type,attr"`
	ID           string   `xml:"ID"`
	DisplayName  string   `xml:"DisplayName"`
	URI          string   `xml:"URI"`
	EmailAddress string   `xml:"EmailAddress"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html#AmazonS3-GetBucketLifecycle-response-GetBucketLifecycleOutput
type LifecycleConfiguration struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ LifecycleConfiguration"`
	Rules   []Rule   `xml:"Rule"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Rule.html
type Rule struct {
	ID                             string                          `xml:"ID"`
	Status                         string                          `xml:"Status"`
	Filter                         *LifecycleRuleFilter            `xml:"Filter"`
	Prefix                         string                          `xml:"Prefix"`
	Expiration                     *LifecycleExpiration            `xml:"Expiration"`
	Transitions                    []Transition                    `xml:"Transition"`
	NoncurrentVersionExpiration    *NoncurrentVersionExpiration    `xml:"NoncurrentVersionExpiration"`
	NoncurrentVersionTransitions   []NoncurrentVersionTransition   `xml:"NoncurrentVersionTransition"`
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload `xml:"AbortIncompleteMultipartUpload"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleRuleFilter.html
type LifecycleRuleFilter struct {
	And                   *LifecycleRuleAndOperator `xml:"And"`
	Prefix                string                    `xml:"Prefix"`
	Tag                   *Tag                      `xml:"Tag"`
	ObjectSizeGreaterThan int64                     `xml:"ObjectSizeGreaterThan"`
	ObjectSizeLessThan    int64                     `xml:"ObjectSizeLessThan"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleRuleAndOperator.html
type LifecycleRuleAndOperator struct {
	Prefix                string `xml:"Prefix"`
	Tags                  []Tag  `xml:"Tag"`
	ObjectSizeGreaterThan int64  `xml:"ObjectSizeGreaterThan"`
	ObjectSizeLessThan    int64  `xml:"ObjectSizeLessThan"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleExpiration.html
type LifecycleExpiration struct {
	Date                      string `xml:"Date"` // ISO8601 format expected
	Days                      int    `xml:"Days"`
	ExpiredObjectDeleteMarker bool   `xml:"ExpiredObjectDeleteMarker"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_Transition.html
type Transition struct {
	Date         string `xml:"Date"` // ISO8601 format
	Days         int    `xml:"Days"`
	StorageClass string `xml:"StorageClass"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_NoncurrentVersionExpiration.html
type NoncurrentVersionExpiration struct {
	NoncurrentDays          int `xml:"NoncurrentDays"`
	NewerNoncurrentVersions int `xml:"NewerNoncurrentVersions"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_NoncurrentVersionTransition.html
type NoncurrentVersionTransition struct {
	NoncurrentDays          int    `xml:"NoncurrentDays"`
	NewerNoncurrentVersions int    `xml:"NewerNoncurrentVersions"`
	StorageClass            string `xml:"StorageClass"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortIncompleteMultipartUpload.html
type AbortIncompleteMultipartUpload struct {
	DaysAfterInitiation int `xml:"DaysAfterInitiation"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_NotificationConfiguration.html
type NotificationConfiguration struct {
	XMLName                      xml.Name                      `xml:"NotificationConfiguration"`
	XMLNS                        string                        `xml:"xmlns,attr,omitempty"`
	TopicConfigurations          []TopicConfiguration          `xml:"TopicConfiguration,omitempty"`
	QueueConfigurations          []QueueConfiguration          `xml:"QueueConfiguration,omitempty"`
	CloudFunctionConfigurations  []LambdaFunctionConfiguration `xml:"CloudFunctionConfiguration,omitempty"`
	EventBridgeConfiguration     *EventBridgeConfiguration     `xml:"EventBridgeConfiguration,omitempty"`
	LambdaFunctionConfigurations []LambdaFunctionConfiguration `xml:"LambdaFunctionConfiguration,omitempty"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_TopicConfiguration.html
type TopicConfiguration struct {
	ID       string                           `xml:"Id"`
	TopicArn string                           `xml:"Topic"`
	Events   []string                         `xml:"Event"`
	Filter   *NotificationConfigurationFilter `xml:"Filter"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_QueueConfiguration.html
type QueueConfiguration struct {
	ID       string                           `xml:"Id"`
	QueueArn string                           `xml:"Queue"`
	Events   []string                         `xml:"Event"`
	Filter   *NotificationConfigurationFilter `xml:"Filter"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_EventBridgeConfiguration.html
type EventBridgeConfiguration struct{}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_NotificationConfigurationFilter.html
type NotificationConfigurationFilter struct {
	Key *S3KeyFilter `xml:"Key"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_S3KeyFilter.html
type S3KeyFilter struct {
	FilterRules []FilterRule `xml:"FilterRule"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_FilterRule.html
type FilterRule struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_LambdaFunctionConfiguration.html
type LambdaFunctionConfiguration struct {
	ID                string                           `xml:"Id"`
	LambdaFunctionArn string                           `xml:"LambdaFunctionArn"`
	Events            []string                         `xml:"Event"`
	Filter            *NotificationConfigurationFilter `xml:"Filter"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html#AmazonS3-GetBucketMetricsConfiguration-response-MetricsConfiguration
type MetricsConfiguration struct {
	XMLName xml.Name       `xml:"MetricsConfiguration"`
	Xmlns   string         `xml:"xmlns,attr"`
	Id      string         `xml:"Id"`
	Filter  *MetricsFilter `xml:"Filter"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_MetricsFilter.html
type MetricsFilter struct {
	AccessPointArn string              `xml:"AccessPointArn"`
	Prefix         string              `xml:"Prefix"`
	Tag            *Tag                `xml:"Tag"`
	And            *MetricsAndOperator `xml:"And"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_MetricsAndOperator.html
type MetricsAndOperator struct {
	AccessPointArn string `xml:"AccessPointArn"`
	Prefix         string `xml:"Prefix"`
	Tags           []Tag  `xml:"Tag"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html#AmazonS3-ListBucketMetricsConfigurations-response-ListBucketMetricsConfigurationsOutput
type ListMetricsConfigurationsResult struct {
	XMLName               xml.Name               `xml:"ListMetricsConfigurationsResult"`
	IsTruncated           bool                   `xml:"IsTruncated"`
	ContinuationToken     string                 `xml:"ContinuationToken"`
	NextContinuationToken string                 `xml:"NextContinuationToken"`
	MetricsConfigurations []MetricsConfiguration `xml:"MetricsConfiguration"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html#AmazonS3-GetObjectLegalHold-response-LegalHold
type LegalHold struct {
	XMLName xml.Name `xml:"LegalHold"`
	Status  string   `xml:"Status"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/userguide/security_iam_service-with-iam.html
type BucketPolicy struct {
	Version   string      `json:"Version"`
	Id        string      `json:"Id,omitempty"`
	Statement []Statement `json:"Statement"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/userguide/security_iam_service-with-iam.html
type Statement struct {
	Sid       string             `json:"Sid,omitempty"`
	Effect    string             `json:"Effect"`
	Principal json.RawMessage    `json:"Principal"` // supports object or string
	Action    interface{}        `json:"Action"`    // string or []string
	Resource  interface{}        `json:"Resource"`  // string or []string
	Condition StatementCondition `json:"Condition,omitempty"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/userguide/security_iam_service-with-iam.html
type StatementCondition map[string]map[string]interface{}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html#AmazonS3-GetBucketPolicyStatus-response-PolicyStatus
type PolicyStatus struct {
	XMLName  xml.Name `xml:"PolicyStatus"`
	IsPublic bool     `xml:"IsPublic"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html#AmazonS3-GetBucketMetadataTableConfiguration-response-GetBucketMetadataTableConfigurationResult
type GetBucketMetadataTableConfigurationResult struct {
	XMLName                    xml.Name                         `xml:"GetBucketMetadataTableConfigurationResult"`
	MetadataTableConfiguration MetadataTableConfigurationResult `xml:"MetadataTableConfigurationResult"`
	Status                     string                           `xml:"Status"`
	Error                      *ErrorDetails                    `xml:"Error"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_ErrorDetails.html
type ErrorDetails struct {
	ErrorCode    string `xml:"ErrorCode"`
	ErrorMessage string `xml:"ErrorMessage"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_MetadataTableConfigurationResult.html
type MetadataTableConfigurationResult struct {
	S3TablesDestination S3TablesDestinationResult `xml:"S3TablesDestinationResult"`
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/API_S3TablesDestinationResult.html
type S3TablesDestinationResult struct {
	TableArn       string `xml:"TableArn"`
	TableBucketArn string `xml:"TableBucketArn"`
	TableName      string `xml:"TableName"`
	TableNamespace string `xml:"TableNamespace"`
}

type PutObjectMetadata struct {
	ContentLength int64
}
