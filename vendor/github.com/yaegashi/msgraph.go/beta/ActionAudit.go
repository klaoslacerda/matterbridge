// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DirectoryAudits returns request builder for DirectoryAudit collection
func (b *AuditLogRootRequestBuilder) DirectoryAudits() *AuditLogRootDirectoryAuditsCollectionRequestBuilder {
	bb := &AuditLogRootDirectoryAuditsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryAudits"
	return bb
}

// AuditLogRootDirectoryAuditsCollectionRequestBuilder is request builder for DirectoryAudit collection
type AuditLogRootDirectoryAuditsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryAudit collection
func (b *AuditLogRootDirectoryAuditsCollectionRequestBuilder) Request() *AuditLogRootDirectoryAuditsCollectionRequest {
	return &AuditLogRootDirectoryAuditsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryAudit item
func (b *AuditLogRootDirectoryAuditsCollectionRequestBuilder) ID(id string) *DirectoryAuditRequestBuilder {
	bb := &DirectoryAuditRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootDirectoryAuditsCollectionRequest is request for DirectoryAudit collection
type AuditLogRootDirectoryAuditsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryAudit, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryAudit
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DirectoryAudit
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DirectoryAudit collection, max N pages
func (r *AuditLogRootDirectoryAuditsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryAudit, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Get(ctx context.Context) ([]DirectoryAudit, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryAudit) (resObj *DirectoryAudit, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DirectoryProvisioning returns request builder for ProvisioningObjectSummary collection
func (b *AuditLogRootRequestBuilder) DirectoryProvisioning() *AuditLogRootDirectoryProvisioningCollectionRequestBuilder {
	bb := &AuditLogRootDirectoryProvisioningCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryProvisioning"
	return bb
}

// AuditLogRootDirectoryProvisioningCollectionRequestBuilder is request builder for ProvisioningObjectSummary collection
type AuditLogRootDirectoryProvisioningCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProvisioningObjectSummary collection
func (b *AuditLogRootDirectoryProvisioningCollectionRequestBuilder) Request() *AuditLogRootDirectoryProvisioningCollectionRequest {
	return &AuditLogRootDirectoryProvisioningCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProvisioningObjectSummary item
func (b *AuditLogRootDirectoryProvisioningCollectionRequestBuilder) ID(id string) *ProvisioningObjectSummaryRequestBuilder {
	bb := &ProvisioningObjectSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootDirectoryProvisioningCollectionRequest is request for ProvisioningObjectSummary collection
type AuditLogRootDirectoryProvisioningCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ProvisioningObjectSummary collection
func (r *AuditLogRootDirectoryProvisioningCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ProvisioningObjectSummary, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ProvisioningObjectSummary
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ProvisioningObjectSummary
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ProvisioningObjectSummary collection, max N pages
func (r *AuditLogRootDirectoryProvisioningCollectionRequest) GetN(ctx context.Context, n int) ([]ProvisioningObjectSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ProvisioningObjectSummary collection
func (r *AuditLogRootDirectoryProvisioningCollectionRequest) Get(ctx context.Context) ([]ProvisioningObjectSummary, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ProvisioningObjectSummary collection
func (r *AuditLogRootDirectoryProvisioningCollectionRequest) Add(ctx context.Context, reqObj *ProvisioningObjectSummary) (resObj *ProvisioningObjectSummary, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Provisioning returns request builder for ProvisioningObjectSummary collection
func (b *AuditLogRootRequestBuilder) Provisioning() *AuditLogRootProvisioningCollectionRequestBuilder {
	bb := &AuditLogRootProvisioningCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/provisioning"
	return bb
}

// AuditLogRootProvisioningCollectionRequestBuilder is request builder for ProvisioningObjectSummary collection
type AuditLogRootProvisioningCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProvisioningObjectSummary collection
func (b *AuditLogRootProvisioningCollectionRequestBuilder) Request() *AuditLogRootProvisioningCollectionRequest {
	return &AuditLogRootProvisioningCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProvisioningObjectSummary item
func (b *AuditLogRootProvisioningCollectionRequestBuilder) ID(id string) *ProvisioningObjectSummaryRequestBuilder {
	bb := &ProvisioningObjectSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootProvisioningCollectionRequest is request for ProvisioningObjectSummary collection
type AuditLogRootProvisioningCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ProvisioningObjectSummary collection
func (r *AuditLogRootProvisioningCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ProvisioningObjectSummary, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ProvisioningObjectSummary
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ProvisioningObjectSummary
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ProvisioningObjectSummary collection, max N pages
func (r *AuditLogRootProvisioningCollectionRequest) GetN(ctx context.Context, n int) ([]ProvisioningObjectSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ProvisioningObjectSummary collection
func (r *AuditLogRootProvisioningCollectionRequest) Get(ctx context.Context) ([]ProvisioningObjectSummary, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ProvisioningObjectSummary collection
func (r *AuditLogRootProvisioningCollectionRequest) Add(ctx context.Context, reqObj *ProvisioningObjectSummary) (resObj *ProvisioningObjectSummary, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RestrictedSignIns returns request builder for RestrictedSignIn collection
func (b *AuditLogRootRequestBuilder) RestrictedSignIns() *AuditLogRootRestrictedSignInsCollectionRequestBuilder {
	bb := &AuditLogRootRestrictedSignInsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/restrictedSignIns"
	return bb
}

// AuditLogRootRestrictedSignInsCollectionRequestBuilder is request builder for RestrictedSignIn collection
type AuditLogRootRestrictedSignInsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RestrictedSignIn collection
func (b *AuditLogRootRestrictedSignInsCollectionRequestBuilder) Request() *AuditLogRootRestrictedSignInsCollectionRequest {
	return &AuditLogRootRestrictedSignInsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RestrictedSignIn item
func (b *AuditLogRootRestrictedSignInsCollectionRequestBuilder) ID(id string) *RestrictedSignInRequestBuilder {
	bb := &RestrictedSignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootRestrictedSignInsCollectionRequest is request for RestrictedSignIn collection
type AuditLogRootRestrictedSignInsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RestrictedSignIn, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []RestrictedSignIn
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []RestrictedSignIn
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for RestrictedSignIn collection, max N pages
func (r *AuditLogRootRestrictedSignInsCollectionRequest) GetN(ctx context.Context, n int) ([]RestrictedSignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Get(ctx context.Context) ([]RestrictedSignIn, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Add(ctx context.Context, reqObj *RestrictedSignIn) (resObj *RestrictedSignIn, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SignIns returns request builder for SignIn collection
func (b *AuditLogRootRequestBuilder) SignIns() *AuditLogRootSignInsCollectionRequestBuilder {
	bb := &AuditLogRootSignInsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/signIns"
	return bb
}

// AuditLogRootSignInsCollectionRequestBuilder is request builder for SignIn collection
type AuditLogRootSignInsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SignIn collection
func (b *AuditLogRootSignInsCollectionRequestBuilder) Request() *AuditLogRootSignInsCollectionRequest {
	return &AuditLogRootSignInsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SignIn item
func (b *AuditLogRootSignInsCollectionRequestBuilder) ID(id string) *SignInRequestBuilder {
	bb := &SignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootSignInsCollectionRequest is request for SignIn collection
type AuditLogRootSignInsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SignIn, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SignIn
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []SignIn
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for SignIn collection, max N pages
func (r *AuditLogRootSignInsCollectionRequest) GetN(ctx context.Context, n int) ([]SignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Get(ctx context.Context) ([]SignIn, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Add(ctx context.Context, reqObj *SignIn) (resObj *SignIn, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
