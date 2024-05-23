// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MacManagedAppProtectionRequestBuilder is request builder for MacManagedAppProtection
type MacManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacManagedAppProtectionRequest
func (b *MacManagedAppProtectionRequestBuilder) Request() *MacManagedAppProtectionRequest {
	return &MacManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacManagedAppProtectionRequest is request for MacManagedAppProtection
type MacManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Get(ctx context.Context) (resObj *MacManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Update(ctx context.Context, reqObj *MacManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacManagedAppProtection
func (r *MacManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSCertificateProfileBaseRequestBuilder is request builder for MacOSCertificateProfileBase
type MacOSCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSCertificateProfileBaseRequest
func (b *MacOSCertificateProfileBaseRequestBuilder) Request() *MacOSCertificateProfileBaseRequest {
	return &MacOSCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSCertificateProfileBaseRequest is request for MacOSCertificateProfileBase
type MacOSCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *MacOSCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *MacOSCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSCertificateProfileBase
func (r *MacOSCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSDeviceFeaturesConfigurationRequestBuilder is request builder for MacOSDeviceFeaturesConfiguration
type MacOSDeviceFeaturesConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSDeviceFeaturesConfigurationRequest
func (b *MacOSDeviceFeaturesConfigurationRequestBuilder) Request() *MacOSDeviceFeaturesConfigurationRequest {
	return &MacOSDeviceFeaturesConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSDeviceFeaturesConfigurationRequest is request for MacOSDeviceFeaturesConfiguration
type MacOSDeviceFeaturesConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSDeviceFeaturesConfiguration
func (r *MacOSDeviceFeaturesConfigurationRequest) Get(ctx context.Context) (resObj *MacOSDeviceFeaturesConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSDeviceFeaturesConfiguration
func (r *MacOSDeviceFeaturesConfigurationRequest) Update(ctx context.Context, reqObj *MacOSDeviceFeaturesConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSDeviceFeaturesConfiguration
func (r *MacOSDeviceFeaturesConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSEnterpriseWiFiConfigurationRequestBuilder is request builder for MacOSEnterpriseWiFiConfiguration
type MacOSEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSEnterpriseWiFiConfigurationRequest
func (b *MacOSEnterpriseWiFiConfigurationRequestBuilder) Request() *MacOSEnterpriseWiFiConfigurationRequest {
	return &MacOSEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSEnterpriseWiFiConfigurationRequest is request for MacOSEnterpriseWiFiConfiguration
type MacOSEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *MacOSEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *MacOSEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSEnterpriseWiFiConfiguration
func (r *MacOSEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSImportedPFXCertificateProfileRequestBuilder is request builder for MacOSImportedPFXCertificateProfile
type MacOSImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSImportedPFXCertificateProfileRequest
func (b *MacOSImportedPFXCertificateProfileRequestBuilder) Request() *MacOSImportedPFXCertificateProfileRequest {
	return &MacOSImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSImportedPFXCertificateProfileRequest is request for MacOSImportedPFXCertificateProfile
type MacOSImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Get(ctx context.Context) (resObj *MacOSImportedPFXCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Update(ctx context.Context, reqObj *MacOSImportedPFXCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSPkcsCertificateProfileRequestBuilder is request builder for MacOSPkcsCertificateProfile
type MacOSPkcsCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSPkcsCertificateProfileRequest
func (b *MacOSPkcsCertificateProfileRequestBuilder) Request() *MacOSPkcsCertificateProfileRequest {
	return &MacOSPkcsCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSPkcsCertificateProfileRequest is request for MacOSPkcsCertificateProfile
type MacOSPkcsCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Get(ctx context.Context) (resObj *MacOSPkcsCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Update(ctx context.Context, reqObj *MacOSPkcsCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSScepCertificateProfileRequestBuilder is request builder for MacOSScepCertificateProfile
type MacOSScepCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSScepCertificateProfileRequest
func (b *MacOSScepCertificateProfileRequestBuilder) Request() *MacOSScepCertificateProfileRequest {
	return &MacOSScepCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSScepCertificateProfileRequest is request for MacOSScepCertificateProfile
type MacOSScepCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for MacOSScepCertificateProfile
func (r *MacOSScepCertificateProfileRequest) Get(ctx context.Context) (resObj *MacOSScepCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSScepCertificateProfile
func (r *MacOSScepCertificateProfileRequest) Update(ctx context.Context, reqObj *MacOSScepCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSScepCertificateProfile
func (r *MacOSScepCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSTrustedRootCertificateRequestBuilder is request builder for MacOSTrustedRootCertificate
type MacOSTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSTrustedRootCertificateRequest
func (b *MacOSTrustedRootCertificateRequestBuilder) Request() *MacOSTrustedRootCertificateRequest {
	return &MacOSTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSTrustedRootCertificateRequest is request for MacOSTrustedRootCertificate
type MacOSTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *MacOSTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *MacOSTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSVpnConfigurationRequestBuilder is request builder for MacOSVpnConfiguration
type MacOSVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSVpnConfigurationRequest
func (b *MacOSVpnConfigurationRequestBuilder) Request() *MacOSVpnConfigurationRequest {
	return &MacOSVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSVpnConfigurationRequest is request for MacOSVpnConfiguration
type MacOSVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Get(ctx context.Context) (resObj *MacOSVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Update(ctx context.Context, reqObj *MacOSVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOSWiredNetworkConfigurationRequestBuilder is request builder for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSWiredNetworkConfigurationRequest
func (b *MacOSWiredNetworkConfigurationRequestBuilder) Request() *MacOSWiredNetworkConfigurationRequest {
	return &MacOSWiredNetworkConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSWiredNetworkConfigurationRequest is request for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Get(ctx context.Context) (resObj *MacOSWiredNetworkConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Update(ctx context.Context, reqObj *MacOSWiredNetworkConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOsVPPAppRequestBuilder is request builder for MacOsVPPApp
type MacOsVPPAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOsVPPAppRequest
func (b *MacOsVPPAppRequestBuilder) Request() *MacOsVPPAppRequest {
	return &MacOsVPPAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOsVPPAppRequest is request for MacOsVPPApp
type MacOsVPPAppRequest struct{ BaseRequest }

// Get performs GET request for MacOsVPPApp
func (r *MacOsVPPAppRequest) Get(ctx context.Context) (resObj *MacOsVPPApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOsVPPApp
func (r *MacOsVPPAppRequest) Update(ctx context.Context, reqObj *MacOsVPPApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOsVPPApp
func (r *MacOsVPPAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MacOsVPPAppAssignedLicenseRequestBuilder is request builder for MacOsVPPAppAssignedLicense
type MacOsVPPAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOsVPPAppAssignedLicenseRequest
func (b *MacOsVPPAppAssignedLicenseRequestBuilder) Request() *MacOsVPPAppAssignedLicenseRequest {
	return &MacOsVPPAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOsVPPAppAssignedLicenseRequest is request for MacOsVPPAppAssignedLicense
type MacOsVPPAppAssignedLicenseRequest struct{ BaseRequest }

// Get performs GET request for MacOsVPPAppAssignedLicense
func (r *MacOsVPPAppAssignedLicenseRequest) Get(ctx context.Context) (resObj *MacOsVPPAppAssignedLicense, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOsVPPAppAssignedLicense
func (r *MacOsVPPAppAssignedLicenseRequest) Update(ctx context.Context, reqObj *MacOsVPPAppAssignedLicense) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOsVPPAppAssignedLicense
func (r *MacOsVPPAppAssignedLicenseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
