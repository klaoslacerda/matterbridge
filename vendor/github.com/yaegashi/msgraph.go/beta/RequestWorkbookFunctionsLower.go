// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLowerRequestBuilder struct{ BaseRequestBuilder }

// Lower action undocumented
func (b *WorkbookFunctionsRequestBuilder) Lower(reqObj *WorkbookFunctionsLowerRequestParameter) *WorkbookFunctionsLowerRequestBuilder {
	bb := &WorkbookFunctionsLowerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/lower"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLowerRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLowerRequestBuilder) Request() *WorkbookFunctionsLowerRequest {
	return &WorkbookFunctionsLowerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLowerRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}