# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResponseBodyServiceResponseBodyServiceGetResponseBody**](ResponseBodyServiceApi.md#ResponseBodyServiceResponseBodyServiceGetResponseBody) | **Get** /responsebody/{data} | 
[**ResponseBodyServiceResponseBodyServiceGetResponseBodyStream**](ResponseBodyServiceApi.md#ResponseBodyServiceResponseBodyServiceGetResponseBodyStream) | **Get** /responsebody/stream/{data} | 
[**ResponseBodyServiceResponseBodyServiceListResponseBodies**](ResponseBodyServiceApi.md#ResponseBodyServiceResponseBodyServiceListResponseBodies) | **Get** /responsebodies/{data} | 
[**ResponseBodyServiceResponseBodyServiceListResponseStrings**](ResponseBodyServiceApi.md#ResponseBodyServiceResponseBodyServiceListResponseStrings) | **Get** /responsestrings/{data} | 

# **ResponseBodyServiceResponseBodyServiceGetResponseBody**
> ExamplepbResponseBodyOutResponse ResponseBodyServiceResponseBodyServiceGetResponseBody(ctx, data)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 

### Return type

[**ExamplepbResponseBodyOutResponse**](examplepbResponseBodyOutResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResponseBodyServiceResponseBodyServiceGetResponseBodyStream**
> StreamResultOfExamplepbResponseBodyOut ResponseBodyServiceResponseBodyServiceGetResponseBodyStream(ctx, data)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 

### Return type

[**StreamResultOfExamplepbResponseBodyOut**](Stream result of examplepbResponseBodyOut.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResponseBodyServiceResponseBodyServiceListResponseBodies**
> []ExamplepbRepeatedResponseBodyOutResponse ResponseBodyServiceResponseBodyServiceListResponseBodies(ctx, data)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 

### Return type

[**[]ExamplepbRepeatedResponseBodyOutResponse**](examplepbRepeatedResponseBodyOutResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResponseBodyServiceResponseBodyServiceListResponseStrings**
> []string ResponseBodyServiceResponseBodyServiceListResponseStrings(ctx, data)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

