package automation

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// NodeReportsClient is the automation Client
type NodeReportsClient struct {
	ManagementClient
}

// NewNodeReportsClient creates an instance of the NodeReportsClient client.
func NewNodeReportsClient(subscriptionID string) NodeReportsClient {
	return NewNodeReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNodeReportsClientWithBaseURI creates an instance of the NodeReportsClient client.
func NewNodeReportsClientWithBaseURI(baseURI string, subscriptionID string) NodeReportsClient {
	return NodeReportsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the Dsc node report data by node id and report id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. nodeID is the
// Dsc node id. reportID is the report id.
func (client NodeReportsClient) Get(resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result DscNodeReport, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.NodeReportsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, nodeID, reportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NodeReportsClient) GetPreparer(resourceGroupName string, automationAccountName string, nodeID string, reportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"reportId":              autorest.Encode("path", reportID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) GetResponder(resp *http.Response) (result DscNodeReport, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetContent retrieve the Dsc node reports by node id and report id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. nodeID is the
// Dsc node id. reportID is the report id.
func (client NodeReportsClient) GetContent(resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result ReadCloser, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.NodeReportsClient", "GetContent")
	}

	req, err := client.GetContentPreparer(resourceGroupName, automationAccountName, nodeID, reportID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "GetContent", resp, "Failure responding to request")
	}

	return
}

// GetContentPreparer prepares the GetContent request.
func (client NodeReportsClient) GetContentPreparer(resourceGroupName string, automationAccountName string, nodeID string, reportID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"reportId":              autorest.Encode("path", reportID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}/content", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetContentSender sends the GetContent request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) GetContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetContentResponder handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) GetContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByNode retrieve the Dsc node report list by node id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. nodeID is the
// parameters supplied to the list operation. filter is the filter to apply on the operation.
func (client NodeReportsClient) ListByNode(resourceGroupName string, automationAccountName string, nodeID string, filter string) (result DscNodeReportListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.NodeReportsClient", "ListByNode")
	}

	req, err := client.ListByNodePreparer(resourceGroupName, automationAccountName, nodeID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByNodeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure sending request")
		return
	}

	result, err = client.ListByNodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure responding to request")
	}

	return
}

// ListByNodePreparer prepares the ListByNode request.
func (client NodeReportsClient) ListByNodePreparer(resourceGroupName string, automationAccountName string, nodeID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByNodeSender sends the ListByNode request. The method will close the
// http.Response Body if it receives an error.
func (client NodeReportsClient) ListByNodeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByNodeResponder handles the response to the ListByNode request. The method always
// closes the http.Response Body.
func (client NodeReportsClient) ListByNodeResponder(resp *http.Response) (result DscNodeReportListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByNodeNextResults retrieves the next set of results, if any.
func (client NodeReportsClient) ListByNodeNextResults(lastResults DscNodeReportListResult) (result DscNodeReportListResult, err error) {
	req, err := lastResults.DscNodeReportListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByNodeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure sending next results request")
	}

	result, err = client.ListByNodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.NodeReportsClient", "ListByNode", resp, "Failure responding to next results request")
	}

	return
}

// ListByNodeComplete gets all elements from the list without paging.
func (client NodeReportsClient) ListByNodeComplete(resourceGroupName string, automationAccountName string, nodeID string, filter string, cancel <-chan struct{}) (<-chan DscNodeReport, <-chan error) {
	resultChan := make(chan DscNodeReport)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByNode(resourceGroupName, automationAccountName, nodeID, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByNodeNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}