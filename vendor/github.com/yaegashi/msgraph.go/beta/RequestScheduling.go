// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SchedulingGroupRequestBuilder is request builder for SchedulingGroup
type SchedulingGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns SchedulingGroupRequest
func (b *SchedulingGroupRequestBuilder) Request() *SchedulingGroupRequest {
	return &SchedulingGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SchedulingGroupRequest is request for SchedulingGroup
type SchedulingGroupRequest struct{ BaseRequest }

// Get performs GET request for SchedulingGroup
func (r *SchedulingGroupRequest) Get(ctx context.Context) (resObj *SchedulingGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SchedulingGroup
func (r *SchedulingGroupRequest) Update(ctx context.Context, reqObj *SchedulingGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SchedulingGroup
func (r *SchedulingGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
