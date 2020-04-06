// Code generated by "mapstructure-to-hcl2 -type Config,CustomerEncryptionKey"; DO NOT EDIT.
package googlecompute

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName              *string                    `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType            *string                    `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug                  *bool                      `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce                  *bool                      `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError                *string                    `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars               map[string]string          `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars          []string                   `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Type                         *string                    `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect           *string                    `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                      *string                    `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                      *int                       `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername                  *string                    `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword                  *string                    `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName               *string                    `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName      *string                    `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys       *bool                      `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile            *string                    `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                       *bool                      `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                   *string                    `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth                 *bool                      `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding    *bool                      `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts         *int                       `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost               *string                    `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort               *int                       `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth          *bool                      `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername           *string                    `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword           *string                    `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionInteractive        *bool                      `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile     *string                    `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod        *string                    `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost                 *string                    `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort                 *int                       `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername             *string                    `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword             *string                    `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval         *string                    `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout          *string                    `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels             []string                   `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels              []string                   `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey                 []byte                     `mapstructure:"ssh_public_key" cty:"ssh_public_key"`
	SSHPrivateKey                []byte                     `mapstructure:"ssh_private_key" cty:"ssh_private_key"`
	WinRMUser                    *string                    `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword                *string                    `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                    *string                    `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                    *int                       `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout                 *string                    `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL                  *bool                      `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure                *bool                      `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM                 *bool                      `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	AccountFile                  *string                    `mapstructure:"account_file" required:"false" cty:"account_file"`
	ProjectId                    *string                    `mapstructure:"project_id" required:"true" cty:"project_id"`
	AcceleratorType              *string                    `mapstructure:"accelerator_type" required:"false" cty:"accelerator_type"`
	AcceleratorCount             *int64                     `mapstructure:"accelerator_count" required:"false" cty:"accelerator_count"`
	Address                      *string                    `mapstructure:"address" required:"false" cty:"address"`
	DisableDefaultServiceAccount *bool                      `mapstructure:"disable_default_service_account" required:"false" cty:"disable_default_service_account"`
	DiskName                     *string                    `mapstructure:"disk_name" required:"false" cty:"disk_name"`
	DiskSizeGb                   *int64                     `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	DiskType                     *string                    `mapstructure:"disk_type" required:"false" cty:"disk_type"`
	ImageName                    *string                    `mapstructure:"image_name" required:"false" cty:"image_name"`
	ImageDescription             *string                    `mapstructure:"image_description" required:"false" cty:"image_description"`
	ImageEncryptionKey           *FlatCustomerEncryptionKey `mapstructure:"image_encryption_key" required:"false" cty:"image_encryption_key"`
	ImageFamily                  *string                    `mapstructure:"image_family" required:"false" cty:"image_family"`
	ImageLabels                  map[string]string          `mapstructure:"image_labels" required:"false" cty:"image_labels"`
	ImageLicenses                []string                   `mapstructure:"image_licenses" required:"false" cty:"image_licenses"`
	InstanceName                 *string                    `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	Labels                       map[string]string          `mapstructure:"labels" required:"false" cty:"labels"`
	MachineType                  *string                    `mapstructure:"machine_type" required:"false" cty:"machine_type"`
	Metadata                     map[string]string          `mapstructure:"metadata" required:"false" cty:"metadata"`
	MetadataFiles                map[string]string          `mapstructure:"metadata_files" cty:"metadata_files"`
	MinCpuPlatform               *string                    `mapstructure:"min_cpu_platform" required:"false" cty:"min_cpu_platform"`
	Network                      *string                    `mapstructure:"network" required:"false" cty:"network"`
	NetworkProjectId             *string                    `mapstructure:"network_project_id" required:"false" cty:"network_project_id"`
	OmitExternalIP               *bool                      `mapstructure:"omit_external_ip" required:"false" cty:"omit_external_ip"`
	OnHostMaintenance            *string                    `mapstructure:"on_host_maintenance" required:"false" cty:"on_host_maintenance"`
	Preemptible                  *bool                      `mapstructure:"preemptible" required:"false" cty:"preemptible"`
	StateTimeout                 *string                    `mapstructure:"state_timeout" required:"false" cty:"state_timeout"`
	Region                       *string                    `mapstructure:"region" required:"false" cty:"region"`
	Scopes                       []string                   `mapstructure:"scopes" required:"false" cty:"scopes"`
	ServiceAccountEmail          *string                    `mapstructure:"service_account_email" required:"false" cty:"service_account_email"`
	SourceImage                  *string                    `mapstructure:"source_image" required:"true" cty:"source_image"`
	SourceImageFamily            *string                    `mapstructure:"source_image_family" required:"true" cty:"source_image_family"`
	SourceImageProjectId         []string                   `mapstructure:"source_image_project_id" required:"false" cty:"source_image_project_id"`
	StartupScriptFile            *string                    `mapstructure:"startup_script_file" required:"false" cty:"startup_script_file"`
	Subnetwork                   *string                    `mapstructure:"subnetwork" required:"false" cty:"subnetwork"`
	Tags                         []string                   `mapstructure:"tags" required:"false" cty:"tags"`
	UseInternalIP                *bool                      `mapstructure:"use_internal_ip" required:"false" cty:"use_internal_ip"`
	VaultGCPOauthEngine          *string                    `mapstructure:"vault_gcp_oauth_engine" cty:"vault_gcp_oauth_engine"`
	Zone                         *string                    `mapstructure:"zone" required:"true" cty:"zone"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":               &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":             &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                    &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                    &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                 &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":           &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":      &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                    &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":         &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                        &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                        &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                    &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                    &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":         &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":       &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":            &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                         &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                     &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                  &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":    &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":          &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":          &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":            &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":            &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":         &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":    &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":        &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                  &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                  &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":              &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":              &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":         &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":          &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":              &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":               &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                  &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                 &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                  &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                  &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                      &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                      &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                   &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                   &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                  &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                  &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"account_file":                    &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"project_id":                      &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"accelerator_type":                &hcldec.AttrSpec{Name: "accelerator_type", Type: cty.String, Required: false},
		"accelerator_count":               &hcldec.AttrSpec{Name: "accelerator_count", Type: cty.Number, Required: false},
		"address":                         &hcldec.AttrSpec{Name: "address", Type: cty.String, Required: false},
		"disable_default_service_account": &hcldec.AttrSpec{Name: "disable_default_service_account", Type: cty.Bool, Required: false},
		"disk_name":                       &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"disk_size":                       &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_type":                       &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"image_name":                      &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":               &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_encryption_key":            &hcldec.BlockSpec{TypeName: "image_encryption_key", Nested: hcldec.ObjectSpec((*FlatCustomerEncryptionKey)(nil).HCL2Spec())},
		"image_family":                    &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"image_labels":                    &hcldec.BlockAttrsSpec{TypeName: "image_labels", ElementType: cty.String, Required: false},
		"image_licenses":                  &hcldec.AttrSpec{Name: "image_licenses", Type: cty.List(cty.String), Required: false},
		"instance_name":                   &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"labels":                          &hcldec.BlockAttrsSpec{TypeName: "labels", ElementType: cty.String, Required: false},
		"machine_type":                    &hcldec.AttrSpec{Name: "machine_type", Type: cty.String, Required: false},
		"metadata":                        &hcldec.BlockAttrsSpec{TypeName: "metadata", ElementType: cty.String, Required: false},
		"metadata_files":                  &hcldec.BlockAttrsSpec{TypeName: "metadata_files", ElementType: cty.String, Required: false},
		"min_cpu_platform":                &hcldec.AttrSpec{Name: "min_cpu_platform", Type: cty.String, Required: false},
		"network":                         &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"network_project_id":              &hcldec.AttrSpec{Name: "network_project_id", Type: cty.String, Required: false},
		"omit_external_ip":                &hcldec.AttrSpec{Name: "omit_external_ip", Type: cty.Bool, Required: false},
		"on_host_maintenance":             &hcldec.AttrSpec{Name: "on_host_maintenance", Type: cty.String, Required: false},
		"preemptible":                     &hcldec.AttrSpec{Name: "preemptible", Type: cty.Bool, Required: false},
		"state_timeout":                   &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
		"region":                          &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"scopes":                          &hcldec.AttrSpec{Name: "scopes", Type: cty.List(cty.String), Required: false},
		"service_account_email":           &hcldec.AttrSpec{Name: "service_account_email", Type: cty.String, Required: false},
		"source_image":                    &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"source_image_family":             &hcldec.AttrSpec{Name: "source_image_family", Type: cty.String, Required: false},
		"source_image_project_id":         &hcldec.AttrSpec{Name: "source_image_project_id", Type: cty.List(cty.String), Required: false},
		"startup_script_file":             &hcldec.AttrSpec{Name: "startup_script_file", Type: cty.String, Required: false},
		"subnetwork":                      &hcldec.AttrSpec{Name: "subnetwork", Type: cty.String, Required: false},
		"tags":                            &hcldec.AttrSpec{Name: "tags", Type: cty.List(cty.String), Required: false},
		"use_internal_ip":                 &hcldec.AttrSpec{Name: "use_internal_ip", Type: cty.Bool, Required: false},
		"vault_gcp_oauth_engine":          &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
		"zone":                            &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
	}
	return s
}

// FlatCustomerEncryptionKey is an auto-generated flat version of CustomerEncryptionKey.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCustomerEncryptionKey struct {
	KmsKeyName *string `json:"kmsKeyName,omitempty" cty:"kms_key_name"`
	RawKey     *string `json:"rawKey,omitempty" cty:"raw_key"`
}

// FlatMapstructure returns a new FlatCustomerEncryptionKey.
// FlatCustomerEncryptionKey is an auto-generated flat version of CustomerEncryptionKey.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CustomerEncryptionKey) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCustomerEncryptionKey)
}

// HCL2Spec returns the hcl spec of a CustomerEncryptionKey.
// This spec is used by HCL to read the fields of CustomerEncryptionKey.
// The decoded values from this spec will then be applied to a FlatCustomerEncryptionKey.
func (*FlatCustomerEncryptionKey) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"kms_key_name": &hcldec.AttrSpec{Name: "kms_key_name", Type: cty.String, Required: false},
		"raw_key":      &hcldec.AttrSpec{Name: "raw_key", Type: cty.String, Required: false},
	}
	return s
}
