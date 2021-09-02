// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v4alpha/tls.proto

package envoy_extensions_transport_sockets_tls_v4alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on UpstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpstreamTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamTlsContextValidationError{
				field:  "CommonTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetSni()) > 255 {
		return UpstreamTlsContextValidationError{
			field:  "Sni",
			reason: "value length must be at most 255 bytes",
		}
	}

	// no validation rules for AllowRenegotiation

	if v, ok := interface{}(m.GetMaxSessionKeys()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamTlsContextValidationError{
				field:  "MaxSessionKeys",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AppendDownstreamPortToSni

	return nil
}

// UpstreamTlsContextValidationError is the validation error returned by
// UpstreamTlsContext.Validate if the designated constraints aren't met.
type UpstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamTlsContextValidationError) ErrorName() string {
	return "UpstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamTlsContextValidationError{}

// Validate checks the field values on DownstreamTlsContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DownstreamTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "CommonTlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequireClientCertificate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "RequireClientCertificate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequireSni()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownstreamTlsContextValidationError{
				field:  "RequireSni",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetSessionTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			return DownstreamTlsContextValidationError{
				field:  "SessionTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		lt := time.Duration(4294967296*time.Second + 0*time.Nanosecond)
		gte := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur < gte || dur >= lt {
			return DownstreamTlsContextValidationError{
				field:  "SessionTimeout",
				reason: "value must be inside range [0s, 1193046h28m16s)",
			}
		}

	}

	if _, ok := DownstreamTlsContext_OcspStaplePolicy_name[int32(m.GetOcspStaplePolicy())]; !ok {
		return DownstreamTlsContextValidationError{
			field:  "OcspStaplePolicy",
			reason: "value must be one of the defined enum values",
		}
	}

	switch m.SessionTicketKeysType.(type) {

	case *DownstreamTlsContext_SessionTicketKeys:

		if v, ok := interface{}(m.GetSessionTicketKeys()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "SessionTicketKeys",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DownstreamTlsContext_SessionTicketKeysSdsSecretConfig:

		if v, ok := interface{}(m.GetSessionTicketKeysSdsSecretConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamTlsContextValidationError{
					field:  "SessionTicketKeysSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DownstreamTlsContext_DisableStatelessSessionResumption:
		// no validation rules for DisableStatelessSessionResumption

	}

	return nil
}

// DownstreamTlsContextValidationError is the validation error returned by
// DownstreamTlsContext.Validate if the designated constraints aren't met.
type DownstreamTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownstreamTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownstreamTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownstreamTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownstreamTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownstreamTlsContextValidationError) ErrorName() string {
	return "DownstreamTlsContextValidationError"
}

// Error satisfies the builtin error interface
func (e DownstreamTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownstreamTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownstreamTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownstreamTlsContextValidationError{}

// Validate checks the field values on CommonTlsContext with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CommonTlsContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTlsParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "TlsParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTlsCertificates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  fmt.Sprintf("TlsCertificates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetTlsCertificateSdsSecretConfigs()) > 2 {
		return CommonTlsContextValidationError{
			field:  "TlsCertificateSdsSecretConfigs",
			reason: "value must contain no more than 2 item(s)",
		}
	}

	for idx, item := range m.GetTlsCertificateSdsSecretConfigs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  fmt.Sprintf("TlsCertificateSdsSecretConfigs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetTlsCertificateProviderInstance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "TlsCertificateProviderInstance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedTlsCertificateCertificateProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "HiddenEnvoyDeprecatedTlsCertificateCertificateProvider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedTlsCertificateCertificateProviderInstance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "HiddenEnvoyDeprecatedTlsCertificateCertificateProviderInstance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCustomHandshaker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContextValidationError{
				field:  "CustomHandshaker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.ValidationContextType.(type) {

	case *CommonTlsContext_ValidationContext:

		if v, ok := interface{}(m.GetValidationContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "ValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_ValidationContextSdsSecretConfig:

		if v, ok := interface{}(m.GetValidationContextSdsSecretConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "ValidationContextSdsSecretConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_CombinedValidationContext:

		if v, ok := interface{}(m.GetCombinedValidationContext()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "CombinedValidationContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_HiddenEnvoyDeprecatedValidationContextCertificateProvider:

		if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedValidationContextCertificateProvider()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "HiddenEnvoyDeprecatedValidationContextCertificateProvider",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CommonTlsContext_HiddenEnvoyDeprecatedValidationContextCertificateProviderInstance:

		if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedValidationContextCertificateProviderInstance()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContextValidationError{
					field:  "HiddenEnvoyDeprecatedValidationContextCertificateProviderInstance",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CommonTlsContextValidationError is the validation error returned by
// CommonTlsContext.Validate if the designated constraints aren't met.
type CommonTlsContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonTlsContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonTlsContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonTlsContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContextValidationError) ErrorName() string { return "CommonTlsContextValidationError" }

// Error satisfies the builtin error interface
func (e CommonTlsContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContextValidationError{}

// Validate checks the field values on CommonTlsContext_CertificateProvider
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CommonTlsContext_CertificateProvider) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CommonTlsContext_CertificateProviderValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	switch m.Config.(type) {

	case *CommonTlsContext_CertificateProvider_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonTlsContext_CertificateProviderValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return CommonTlsContext_CertificateProviderValidationError{
			field:  "Config",
			reason: "value is required",
		}

	}

	return nil
}

// CommonTlsContext_CertificateProviderValidationError is the validation error
// returned by CommonTlsContext_CertificateProvider.Validate if the designated
// constraints aren't met.
type CommonTlsContext_CertificateProviderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContext_CertificateProviderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonTlsContext_CertificateProviderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonTlsContext_CertificateProviderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonTlsContext_CertificateProviderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContext_CertificateProviderValidationError) ErrorName() string {
	return "CommonTlsContext_CertificateProviderValidationError"
}

// Error satisfies the builtin error interface
func (e CommonTlsContext_CertificateProviderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext_CertificateProvider.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContext_CertificateProviderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContext_CertificateProviderValidationError{}

// Validate checks the field values on
// CommonTlsContext_CertificateProviderInstance with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CommonTlsContext_CertificateProviderInstance) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for InstanceName

	// no validation rules for CertificateName

	return nil
}

// CommonTlsContext_CertificateProviderInstanceValidationError is the
// validation error returned by
// CommonTlsContext_CertificateProviderInstance.Validate if the designated
// constraints aren't met.
type CommonTlsContext_CertificateProviderInstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContext_CertificateProviderInstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonTlsContext_CertificateProviderInstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonTlsContext_CertificateProviderInstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonTlsContext_CertificateProviderInstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonTlsContext_CertificateProviderInstanceValidationError) ErrorName() string {
	return "CommonTlsContext_CertificateProviderInstanceValidationError"
}

// Error satisfies the builtin error interface
func (e CommonTlsContext_CertificateProviderInstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext_CertificateProviderInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContext_CertificateProviderInstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContext_CertificateProviderInstanceValidationError{}

// Validate checks the field values on
// CommonTlsContext_CombinedCertificateValidationContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CommonTlsContext_CombinedCertificateValidationContext) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDefaultValidationContext() == nil {
		return CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "DefaultValidationContext",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDefaultValidationContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "DefaultValidationContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetValidationContextSdsSecretConfig() == nil {
		return CommonTlsContext_CombinedCertificateValidationContextValidationError{
			field:  "ValidationContextSdsSecretConfig",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetValidationContextSdsSecretConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "ValidationContextSdsSecretConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedValidationContextCertificateProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "HiddenEnvoyDeprecatedValidationContextCertificateProvider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedValidationContextCertificateProviderInstance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonTlsContext_CombinedCertificateValidationContextValidationError{
				field:  "HiddenEnvoyDeprecatedValidationContextCertificateProviderInstance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CommonTlsContext_CombinedCertificateValidationContextValidationError is the
// validation error returned by
// CommonTlsContext_CombinedCertificateValidationContext.Validate if the
// designated constraints aren't met.
type CommonTlsContext_CombinedCertificateValidationContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) ErrorName() string {
	return "CommonTlsContext_CombinedCertificateValidationContextValidationError"
}

// Error satisfies the builtin error interface
func (e CommonTlsContext_CombinedCertificateValidationContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonTlsContext_CombinedCertificateValidationContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonTlsContext_CombinedCertificateValidationContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonTlsContext_CombinedCertificateValidationContextValidationError{}
