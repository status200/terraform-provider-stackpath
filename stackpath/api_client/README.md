# Go API client for api_client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.stackpath.com/](https://support.stackpath.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./api_client"
```

## Documentation for API Endpoints

All URIs are relative to *https://gateway.stackpath.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BucketsApi* | [**CreateBucket**](docs/BucketsApi.md#createbucket) | **Post** /storage/v1/stacks/{stack_id}/buckets | Create a bucket under a stack
*BucketsApi* | [**DeleteBucket**](docs/BucketsApi.md#deletebucket) | **Delete** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Delete a given bucket
*BucketsApi* | [**GetBucket**](docs/BucketsApi.md#getbucket) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Retrieve a bucket in the storage provider for a given stack
*BucketsApi* | [**GetBuckets**](docs/BucketsApi.md#getbuckets) | **Get** /storage/v1/stacks/{stack_id}/buckets | Retrieve all buckets in the storage provider for a given stack
*BucketsApi* | [**UpdateBucket**](docs/BucketsApi.md#updatebucket) | **Put** /storage/v1/stacks/{stack_id}/buckets/{bucket_id} | Updates the name of a bucket
*MetricsApi* | [**GetBucketMetrics**](docs/MetricsApi.md#getbucketmetrics) | **Get** /storage/v1/stacks/{stack_id}/buckets/{bucket_id}/metrics | Get all daily utilizations for specific bucket
*MetricsApi* | [**GetStackMetrics**](docs/MetricsApi.md#getstackmetrics) | **Get** /storage/v1/stacks/{stack_id}/metrics | Get all daily utilizations for all buckets on a stack
*UserCredentialsApi* | [**DeleteCredential**](docs/UserCredentialsApi.md#deletecredential) | **Delete** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/{access_key} | Delete provided storage access credentials for the given user
*UserCredentialsApi* | [**GenerateCredentials**](docs/UserCredentialsApi.md#generatecredentials) | **Post** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials/generate | Generate storage credentials for the given user
*UserCredentialsApi* | [**GetCredentials**](docs/UserCredentialsApi.md#getcredentials) | **Get** /storage/v1/stacks/{stack_id}/users/{user_id}/credentials | Get credentials for a given user.


## Documentation For Models

 - [ApiStatusDetail](docs/ApiStatusDetail.md)
 - [DataMatrix](docs/DataMatrix.md)
 - [DataMatrixResult](docs/DataMatrixResult.md)
 - [DataValue](docs/DataValue.md)
 - [DataVector](docs/DataVector.md)
 - [DataVectorResult](docs/DataVectorResult.md)
 - [GetCredentialsResponseCredential](docs/GetCredentialsResponseCredential.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2004Data](docs/InlineResponse2004Data.md)
 - [InlineResponse2004DataMatrix](docs/InlineResponse2004DataMatrix.md)
 - [InlineResponse2004DataMatrixResults](docs/InlineResponse2004DataMatrixResults.md)
 - [InlineResponse2004DataMatrixValues](docs/InlineResponse2004DataMatrixValues.md)
 - [InlineResponse2004DataVector](docs/InlineResponse2004DataVector.md)
 - [InlineResponse2004DataVectorResults](docs/InlineResponse2004DataVectorResults.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2005Credentials](docs/InlineResponse2005Credentials.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse200PageInfo](docs/InlineResponse200PageInfo.md)
 - [InlineResponse200Results](docs/InlineResponse200Results.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse401Details](docs/InlineResponse401Details.md)
 - [MetricsData](docs/MetricsData.md)
 - [PaginationPageInfo](docs/PaginationPageInfo.md)
 - [PaginationPageRequest](docs/PaginationPageRequest.md)
 - [PrometheusMetrics](docs/PrometheusMetrics.md)
 - [PrometheusMetricsStatus](docs/PrometheusMetricsStatus.md)
 - [StackpathRpcBadRequest](docs/StackpathRpcBadRequest.md)
 - [StackpathRpcBadRequestAllOf](docs/StackpathRpcBadRequestAllOf.md)
 - [StackpathRpcBadRequestAllOfFieldViolations](docs/StackpathRpcBadRequestAllOfFieldViolations.md)
 - [StackpathRpcBadRequestFieldViolation](docs/StackpathRpcBadRequestFieldViolation.md)
 - [StackpathRpcHelp](docs/StackpathRpcHelp.md)
 - [StackpathRpcHelpAllOf](docs/StackpathRpcHelpAllOf.md)
 - [StackpathRpcHelpAllOfLinks](docs/StackpathRpcHelpAllOfLinks.md)
 - [StackpathRpcHelpLink](docs/StackpathRpcHelpLink.md)
 - [StackpathRpcLocalizedMessage](docs/StackpathRpcLocalizedMessage.md)
 - [StackpathRpcLocalizedMessageAllOf](docs/StackpathRpcLocalizedMessageAllOf.md)
 - [StackpathRpcPreconditionFailure](docs/StackpathRpcPreconditionFailure.md)
 - [StackpathRpcPreconditionFailureAllOf](docs/StackpathRpcPreconditionFailureAllOf.md)
 - [StackpathRpcPreconditionFailureAllOfViolations](docs/StackpathRpcPreconditionFailureAllOfViolations.md)
 - [StackpathRpcPreconditionFailureViolation](docs/StackpathRpcPreconditionFailureViolation.md)
 - [StackpathRpcQuotaFailure](docs/StackpathRpcQuotaFailure.md)
 - [StackpathRpcQuotaFailureAllOf](docs/StackpathRpcQuotaFailureAllOf.md)
 - [StackpathRpcQuotaFailureAllOfViolations](docs/StackpathRpcQuotaFailureAllOfViolations.md)
 - [StackpathRpcQuotaFailureViolation](docs/StackpathRpcQuotaFailureViolation.md)
 - [StackpathRpcRequestInfo](docs/StackpathRpcRequestInfo.md)
 - [StackpathRpcRequestInfoAllOf](docs/StackpathRpcRequestInfoAllOf.md)
 - [StackpathRpcResourceInfo](docs/StackpathRpcResourceInfo.md)
 - [StackpathRpcResourceInfoAllOf](docs/StackpathRpcResourceInfoAllOf.md)
 - [StackpathRpcRetryInfo](docs/StackpathRpcRetryInfo.md)
 - [StackpathRpcRetryInfoAllOf](docs/StackpathRpcRetryInfoAllOf.md)
 - [StackpathapiStatus](docs/StackpathapiStatus.md)
 - [StorageBucket](docs/StorageBucket.md)
 - [StorageBucketVisibility](docs/StorageBucketVisibility.md)
 - [StorageCreateBucketRequest](docs/StorageCreateBucketRequest.md)
 - [StorageCreateBucketResponse](docs/StorageCreateBucketResponse.md)
 - [StorageGenerateCredentialsResponse](docs/StorageGenerateCredentialsResponse.md)
 - [StorageGetBucketResponse](docs/StorageGetBucketResponse.md)
 - [StorageGetBucketsResponse](docs/StorageGetBucketsResponse.md)
 - [StorageGetCredentialsResponse](docs/StorageGetCredentialsResponse.md)
 - [StorageUpdateBucketRequest](docs/StorageUpdateBucketRequest.md)
 - [StorageUpdateBucketResponse](docs/StorageUpdateBucketResponse.md)


## Documentation For Authorization



## oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author


