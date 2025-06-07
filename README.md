

# Go-S3-Wrapper

A wrapper module usable for communicating with a S3 compatible object storage via HTTP.
Http client for request execution can be configured manually so that it could be used in a e.g. WebAssembly context.

Prototype WASI-HTTP Interface: https://github.com/ydnar/wasi-http-go 


## Get Started


1. Add the module to your ``go.mod``

````go
require github.com/cedweber/go-s3-wrapper
````


2. Use the module within your logic

````go 
s3 "github.com/cedweber/go-s3-wrapper"
````


3. Use the module
````go
	cfg := s3.Config{
		Endpoint:  baseDomain,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
	}

	// Create a New S3 client.
	s3Client, err := s3.New(cfg)
	if err != nil {
		fmt.Printf("failed to create source client %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
````

Use the client to interact via REST with S3, e.g.

````go
	bucketName := "my-bucket"
	filePath := "my-file"
	ctx := context.Background()

	// Get http response from HEAD request
	resp, err := s3Client.HeadObject(ctx, bucketName, filePath)
	if err != nil {
		fmt.Printf("failed to get file info %\n", err)
	}
````

## Supported Operations

The following operations are supported:

- CreateBucket
- ListBuckets

##### Object Operations

- ListObjects
- ListObjectsV2
- ListObjectVersions
- HeadObject
- GetObject
- GetObjectPart
- PutObject
- PutObjectStream
- DeleteObject
- DeleteObjects

##### Multipart

- CreateMultipartUpload
- UploadPart
- CompleteMultipartUpload
- ListMultipartUploads
- AbortMultipartUpload
- ListParts

##### Object Tagging

- GetObjectTagging
- PutObjectTagging
- DeleteObjectTagging

##### Others

- GetObjectAttributes
- ListDirectoryBuckets

##### Bucket Website

- GetBucketWebsite
- PutBucketWebsite
- DeleteBucketWebsite

##### Bucket Versioning

- GetBucketVersioning
- PutBucketVersioning

##### Bucket Tagging

- GetBucketTagging
- PutBucketTagging
- DeleteBucketTagging

##### Object Lock Config

- PutObjectLockConfiguration
- GetObjectLockConfiguration

##### Object Retention

- GetObjectRetention
- PutObjectRetention

##### Object Access Control

- GetObjectAcl
- PutObjectAcl

##### Bucket Access Control

- GetBucketAcl
- PutBucketAcl

##### Bucket Logging

- GetBucketLogging
- PutBucketLogging

##### Public Access Block

- GetPublicAccessBlock
- PutPublicAccessBlock
- DeletePublicAccessBlock

##### Bucket Notification Configuration

- GetBucketNotificationConfiguration
- PutBucketNotificationConfiguration

##### Bucket Metrics

- GetBucketMetricsConfiguration
- ListBucketMetricsConfigurations
- PutBucketMetricsConfiguration
- DeleteBucketMetricsConfiguration

##### Object Legal Hold

- GetObjectLegalHold
- PutObjectLegalHold

##### Bucket Policy

- GetBucketPolicyStatus
- GetBucketPolicy
- PutBucketPolicy
- DeleteBucketPolicy

##### Bucket Lifecycle Configuration

- GetBucketLifecycleConfiguration
- PutBucketLifecycleConfiguration
- DeleteBucketLifecycle

##### Bucket Metadata Configuration

- GetBucketMetadataTableConfiguration
- CreateBucketMetadataTableConfiguration
- DeleteBucketMetadataTableConfiguration


## Contribution
If you think something is missing or wrong feel free to contribute.


## Credits
Thanks and credits to [Fermyon](https://www.fermyon.com) for the support and the initial code base