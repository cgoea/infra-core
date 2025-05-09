// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/provider/v1/provider.proto

package providerv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on ProviderResource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProviderResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProviderResource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProviderResourceMultiError, or nil if none found.
func (m *ProviderResource) ValidateAll() error {
	return m.validate(true)
}

func (m *ProviderResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResourceId()) > 17 {
		err := ProviderResourceValidationError{
			field:  "ResourceId",
			reason: "value length must be at most 17 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProviderResource_ResourceId_Pattern.MatchString(m.GetResourceId()) {
		err := ProviderResourceValidationError{
			field:  "ResourceId",
			reason: "value does not match regex pattern \"^provider-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ProviderKind

	// no validation rules for ProviderVendor

	if len(m.GetName()) > 40 {
		err := ProviderResourceValidationError{
			field:  "Name",
			reason: "value length must be at most 40 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProviderResource_Name_Pattern.MatchString(m.GetName()) {
		err := ProviderResourceValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./: ]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetApiEndpoint()) > 500 {
		err := ProviderResourceValidationError{
			field:  "ApiEndpoint",
			reason: "value length must be at most 500 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProviderResource_ApiEndpoint_Pattern.MatchString(m.GetApiEndpoint()) {
		err := ProviderResourceValidationError{
			field:  "ApiEndpoint",
			reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./:;=?@!#,<>*() ]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetApiCredentials() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) > 500 {
			err := ProviderResourceValidationError{
				field:  fmt.Sprintf("ApiCredentials[%v]", idx),
				reason: "value length must be at most 500 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if !_ProviderResource_ApiCredentials_Pattern.MatchString(item) {
			err := ProviderResourceValidationError{
				field:  fmt.Sprintf("ApiCredentials[%v]", idx),
				reason: "value does not match regex pattern \"^[^|]*$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(m.GetConfig()) > 2000 {
		err := ProviderResourceValidationError{
			field:  "Config",
			reason: "value length must be at most 2000 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProviderResource_Config_Pattern.MatchString(m.GetConfig()) {
		err := ProviderResourceValidationError{
			field:  "Config",
			reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./:;=?@!#,<>*() ]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetProviderId()) > 17 {
		err := ProviderResourceValidationError{
			field:  "ProviderId",
			reason: "value length must be at most 17 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProviderResource_ProviderId_Pattern.MatchString(m.GetProviderId()) {
		err := ProviderResourceValidationError{
			field:  "ProviderId",
			reason: "value does not match regex pattern \"^provider-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProviderResourceMultiError(errors)
	}

	return nil
}

// ProviderResourceMultiError is an error wrapping multiple validation errors
// returned by ProviderResource.ValidateAll() if the designated constraints
// aren't met.
type ProviderResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProviderResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProviderResourceMultiError) AllErrors() []error { return m }

// ProviderResourceValidationError is the validation error returned by
// ProviderResource.Validate if the designated constraints aren't met.
type ProviderResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderResourceValidationError) ErrorName() string { return "ProviderResourceValidationError" }

// Error satisfies the builtin error interface
func (e ProviderResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderResourceValidationError{}

var _ProviderResource_ResourceId_Pattern = regexp.MustCompile("^provider-[0-9a-f]{8}$")

var _ProviderResource_Name_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./: ]+$")

var _ProviderResource_ApiEndpoint_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./:;=?@!#,<>*() ]+$")

var _ProviderResource_ApiCredentials_Pattern = regexp.MustCompile("^[^|]*$")

var _ProviderResource_Config_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./:;=?@!#,<>*() ]+$")

var _ProviderResource_ProviderId_Pattern = regexp.MustCompile("^provider-[0-9a-f]{8}$")
