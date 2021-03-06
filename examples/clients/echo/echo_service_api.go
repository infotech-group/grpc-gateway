/*
 * Echo Service
 *
 * Echo Service API consists of a single service which returns a message.
 *
 * OpenAPI spec version: version not set
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package echo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type EchoServiceApi struct {
	Configuration *Configuration
}

func NewEchoServiceApi() *EchoServiceApi {
	configuration := NewConfiguration()
	return &EchoServiceApi{
		Configuration: configuration,
	}
}

func NewEchoServiceApiWithBasePath(basePath string) *EchoServiceApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &EchoServiceApi{
		Configuration: configuration,
	}
}

/**
 * Echo method receives a simple message and returns it.
 * The message posted as the id parameter will also be returned.
 *
 * @param id Id represents the message identifier.
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) Echo(id string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Echo method receives a simple message and returns it.
 * The message posted as the id parameter will also be returned.
 *
 * @param id Id represents the message identifier.
 * @param num
 * @param lineNum
 * @param lang
 * @param statusProgress
 * @param statusNote
 * @param en
 * @param noProgress
 * @param noNote
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) Echo2(id string, num string, lineNum string, lang string, statusProgress string, statusNote string, en string, noProgress string, noNote string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo/{id}/{num}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"num"+"}", fmt.Sprintf("%v", num), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("line_num", a.Configuration.APIClient.ParameterToString(lineNum, ""))
	localVarQueryParams.Add("lang", a.Configuration.APIClient.ParameterToString(lang, ""))
	localVarQueryParams.Add("status.progress", a.Configuration.APIClient.ParameterToString(statusProgress, ""))
	localVarQueryParams.Add("status.note", a.Configuration.APIClient.ParameterToString(statusNote, ""))
	localVarQueryParams.Add("en", a.Configuration.APIClient.ParameterToString(en, ""))
	localVarQueryParams.Add("no.progress", a.Configuration.APIClient.ParameterToString(noProgress, ""))
	localVarQueryParams.Add("no.note", a.Configuration.APIClient.ParameterToString(noNote, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo2", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Echo method receives a simple message and returns it.
 * The message posted as the id parameter will also be returned.
 *
 * @param id Id represents the message identifier.
 * @param num
 * @param lang
 * @param lineNum
 * @param statusProgress
 * @param statusNote
 * @param en
 * @param noProgress
 * @param noNote
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) Echo3(id string, num string, lang string, lineNum string, statusProgress string, statusNote string, en string, noProgress string, noNote string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo/{id}/{num}/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"num"+"}", fmt.Sprintf("%v", num), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", fmt.Sprintf("%v", lang), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("line_num", a.Configuration.APIClient.ParameterToString(lineNum, ""))
	localVarQueryParams.Add("status.progress", a.Configuration.APIClient.ParameterToString(statusProgress, ""))
	localVarQueryParams.Add("status.note", a.Configuration.APIClient.ParameterToString(statusNote, ""))
	localVarQueryParams.Add("en", a.Configuration.APIClient.ParameterToString(en, ""))
	localVarQueryParams.Add("no.progress", a.Configuration.APIClient.ParameterToString(noProgress, ""))
	localVarQueryParams.Add("no.note", a.Configuration.APIClient.ParameterToString(noNote, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo3", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Echo method receives a simple message and returns it.
 * The message posted as the id parameter will also be returned.
 *
 * @param id Id represents the message identifier.
 * @param lineNum
 * @param statusNote
 * @param num
 * @param lang
 * @param statusProgress
 * @param en
 * @param noProgress
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) Echo4(id string, lineNum string, statusNote string, num string, lang string, statusProgress string, en string, noProgress string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo1/{id}/{line_num}/{status.note}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"line_num"+"}", fmt.Sprintf("%v", lineNum), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"status.note"+"}", fmt.Sprintf("%v", statusNote), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("num", a.Configuration.APIClient.ParameterToString(num, ""))
	localVarQueryParams.Add("lang", a.Configuration.APIClient.ParameterToString(lang, ""))
	localVarQueryParams.Add("status.progress", a.Configuration.APIClient.ParameterToString(statusProgress, ""))
	localVarQueryParams.Add("en", a.Configuration.APIClient.ParameterToString(en, ""))
	localVarQueryParams.Add("no.progress", a.Configuration.APIClient.ParameterToString(noProgress, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo4", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Echo method receives a simple message and returns it.
 * The message posted as the id parameter will also be returned.
 *
 * @param noNote
 * @param id Id represents the message identifier.
 * @param num
 * @param lineNum
 * @param lang
 * @param statusProgress
 * @param en
 * @param noProgress
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) Echo5(noNote string, id string, num string, lineNum string, lang string, statusProgress string, en string, noProgress string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo2/{no.note}"
	localVarPath = strings.Replace(localVarPath, "{"+"no.note"+"}", fmt.Sprintf("%v", noNote), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("id", a.Configuration.APIClient.ParameterToString(id, ""))
	localVarQueryParams.Add("num", a.Configuration.APIClient.ParameterToString(num, ""))
	localVarQueryParams.Add("line_num", a.Configuration.APIClient.ParameterToString(lineNum, ""))
	localVarQueryParams.Add("lang", a.Configuration.APIClient.ParameterToString(lang, ""))
	localVarQueryParams.Add("status.progress", a.Configuration.APIClient.ParameterToString(statusProgress, ""))
	localVarQueryParams.Add("en", a.Configuration.APIClient.ParameterToString(en, ""))
	localVarQueryParams.Add("no.progress", a.Configuration.APIClient.ParameterToString(noProgress, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "Echo5", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * EchoBody method receives a simple message and returns it.
 *
 * @param body
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) EchoBody(body ExamplepbSimpleMessage) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo_body"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "EchoBody", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * EchoDelete method receives a simple message and returns it.
 *
 * @param id Id represents the message identifier.
 * @param num
 * @param lineNum
 * @param lang
 * @param statusProgress
 * @param statusNote
 * @param en
 * @param noProgress
 * @param noNote
 * @return *ExamplepbSimpleMessage
 */
func (a EchoServiceApi) EchoDelete(id string, num string, lineNum string, lang string, statusProgress string, statusNote string, en string, noProgress string, noNote string) (*ExamplepbSimpleMessage, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/example/echo_delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("id", a.Configuration.APIClient.ParameterToString(id, ""))
	localVarQueryParams.Add("num", a.Configuration.APIClient.ParameterToString(num, ""))
	localVarQueryParams.Add("line_num", a.Configuration.APIClient.ParameterToString(lineNum, ""))
	localVarQueryParams.Add("lang", a.Configuration.APIClient.ParameterToString(lang, ""))
	localVarQueryParams.Add("status.progress", a.Configuration.APIClient.ParameterToString(statusProgress, ""))
	localVarQueryParams.Add("status.note", a.Configuration.APIClient.ParameterToString(statusNote, ""))
	localVarQueryParams.Add("en", a.Configuration.APIClient.ParameterToString(en, ""))
	localVarQueryParams.Add("no.progress", a.Configuration.APIClient.ParameterToString(noProgress, ""))
	localVarQueryParams.Add("no.note", a.Configuration.APIClient.ParameterToString(noNote, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ExamplepbSimpleMessage)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "EchoDelete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
