// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DomainJoinConfiguration is navigation property
func (b *ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder) DomainJoinConfiguration() *WindowsDomainJoinConfigurationRequestBuilder {
	bb := &WindowsDomainJoinConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/domainJoinConfiguration"
	return bb
}