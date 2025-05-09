// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resources/telemetry/v1/telemetry.proto

package telemetryv1

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

// Validate checks the field values on TelemetryLogsGroupResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TelemetryLogsGroupResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TelemetryLogsGroupResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TelemetryLogsGroupResourceMultiError, or nil if none found.
func (m *TelemetryLogsGroupResource) ValidateAll() error {
	return m.validate(true)
}

func (m *TelemetryLogsGroupResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResourceId()) > 23 {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "ResourceId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsGroupResource_ResourceId_Pattern.MatchString(m.GetResourceId()) {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "ResourceId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTelemetryLogsGroupId()) > 23 {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "TelemetryLogsGroupId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsGroupResource_TelemetryLogsGroupId_Pattern.MatchString(m.GetTelemetryLogsGroupId()) {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "TelemetryLogsGroupId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) > 20 {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "Name",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsGroupResource_Name_Pattern.MatchString(m.GetName()) {
		err := TelemetryLogsGroupResourceValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./: ]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CollectorKind

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if !_TelemetryLogsGroupResource_Groups_Pattern.MatchString(item) {
			err := TelemetryLogsGroupResourceValidationError{
				field:  fmt.Sprintf("Groups[%v]", idx),
				reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./:;=@?!#,<>*()\\\" ]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return TelemetryLogsGroupResourceMultiError(errors)
	}

	return nil
}

// TelemetryLogsGroupResourceMultiError is an error wrapping multiple
// validation errors returned by TelemetryLogsGroupResource.ValidateAll() if
// the designated constraints aren't met.
type TelemetryLogsGroupResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TelemetryLogsGroupResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TelemetryLogsGroupResourceMultiError) AllErrors() []error { return m }

// TelemetryLogsGroupResourceValidationError is the validation error returned
// by TelemetryLogsGroupResource.Validate if the designated constraints aren't met.
type TelemetryLogsGroupResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TelemetryLogsGroupResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TelemetryLogsGroupResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TelemetryLogsGroupResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TelemetryLogsGroupResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TelemetryLogsGroupResourceValidationError) ErrorName() string {
	return "TelemetryLogsGroupResourceValidationError"
}

// Error satisfies the builtin error interface
func (e TelemetryLogsGroupResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTelemetryLogsGroupResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TelemetryLogsGroupResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TelemetryLogsGroupResourceValidationError{}

var _TelemetryLogsGroupResource_ResourceId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")

var _TelemetryLogsGroupResource_TelemetryLogsGroupId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")

var _TelemetryLogsGroupResource_Name_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./: ]+$")

var _TelemetryLogsGroupResource_Groups_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./:;=@?!#,<>*()\" ]+$")

// Validate checks the field values on TelemetryMetricsGroupResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TelemetryMetricsGroupResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TelemetryMetricsGroupResource with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// TelemetryMetricsGroupResourceMultiError, or nil if none found.
func (m *TelemetryMetricsGroupResource) ValidateAll() error {
	return m.validate(true)
}

func (m *TelemetryMetricsGroupResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResourceId()) > 23 {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "ResourceId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsGroupResource_ResourceId_Pattern.MatchString(m.GetResourceId()) {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "ResourceId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTelemetryMetricsGroupId()) > 23 {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "TelemetryMetricsGroupId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsGroupResource_TelemetryMetricsGroupId_Pattern.MatchString(m.GetTelemetryMetricsGroupId()) {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "TelemetryMetricsGroupId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) > 20 {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "Name",
			reason: "value length must be at most 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsGroupResource_Name_Pattern.MatchString(m.GetName()) {
		err := TelemetryMetricsGroupResourceValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./: ]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CollectorKind

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if !_TelemetryMetricsGroupResource_Groups_Pattern.MatchString(item) {
			err := TelemetryMetricsGroupResourceValidationError{
				field:  fmt.Sprintf("Groups[%v]", idx),
				reason: "value does not match regex pattern \"^$|^[a-zA-Z-_0-9./:;=@?!#,<>*()\\\" ]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return TelemetryMetricsGroupResourceMultiError(errors)
	}

	return nil
}

// TelemetryMetricsGroupResourceMultiError is an error wrapping multiple
// validation errors returned by TelemetryMetricsGroupResource.ValidateAll()
// if the designated constraints aren't met.
type TelemetryMetricsGroupResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TelemetryMetricsGroupResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TelemetryMetricsGroupResourceMultiError) AllErrors() []error { return m }

// TelemetryMetricsGroupResourceValidationError is the validation error
// returned by TelemetryMetricsGroupResource.Validate if the designated
// constraints aren't met.
type TelemetryMetricsGroupResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TelemetryMetricsGroupResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TelemetryMetricsGroupResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TelemetryMetricsGroupResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TelemetryMetricsGroupResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TelemetryMetricsGroupResourceValidationError) ErrorName() string {
	return "TelemetryMetricsGroupResourceValidationError"
}

// Error satisfies the builtin error interface
func (e TelemetryMetricsGroupResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTelemetryMetricsGroupResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TelemetryMetricsGroupResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TelemetryMetricsGroupResourceValidationError{}

var _TelemetryMetricsGroupResource_ResourceId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")

var _TelemetryMetricsGroupResource_TelemetryMetricsGroupId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")

var _TelemetryMetricsGroupResource_Name_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./: ]+$")

var _TelemetryMetricsGroupResource_Groups_Pattern = regexp.MustCompile("^$|^[a-zA-Z-_0-9./:;=@?!#,<>*()\" ]+$")

// Validate checks the field values on TelemetryLogsProfileResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TelemetryLogsProfileResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TelemetryLogsProfileResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TelemetryLogsProfileResourceMultiError, or nil if none found.
func (m *TelemetryLogsProfileResource) ValidateAll() error {
	return m.validate(true)
}

func (m *TelemetryLogsProfileResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResourceId()) > 25 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "ResourceId",
			reason: "value length must be at most 25 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_ResourceId_Pattern.MatchString(m.GetResourceId()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "ResourceId",
			reason: "value does not match regex pattern \"^telemetryprofile-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetProfileId()) > 25 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "ProfileId",
			reason: "value length must be at most 25 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_ProfileId_Pattern.MatchString(m.GetProfileId()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "ProfileId",
			reason: "value does not match regex pattern \"^telemetryprofile-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetInstance()) > 13 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetInstance",
			reason: "value length must be at most 13 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_TargetInstance_Pattern.MatchString(m.GetTargetInstance()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetInstance",
			reason: "value does not match regex pattern \"^inst-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetSite()) > 13 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetSite",
			reason: "value length must be at most 13 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_TargetSite_Pattern.MatchString(m.GetTargetSite()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetSite",
			reason: "value does not match regex pattern \"^site-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetRegion()) > 15 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetRegion",
			reason: "value length must be at most 15 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_TargetRegion_Pattern.MatchString(m.GetTargetRegion()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "TargetRegion",
			reason: "value does not match regex pattern \"^region-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for LogLevel

	if len(m.GetLogsGroupId()) > 23 {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "LogsGroupId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryLogsProfileResource_LogsGroupId_Pattern.MatchString(m.GetLogsGroupId()) {
		err := TelemetryLogsProfileResourceValidationError{
			field:  "LogsGroupId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetLogsGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TelemetryLogsProfileResourceValidationError{
					field:  "LogsGroup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TelemetryLogsProfileResourceValidationError{
					field:  "LogsGroup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLogsGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TelemetryLogsProfileResourceValidationError{
				field:  "LogsGroup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TelemetryLogsProfileResourceMultiError(errors)
	}

	return nil
}

// TelemetryLogsProfileResourceMultiError is an error wrapping multiple
// validation errors returned by TelemetryLogsProfileResource.ValidateAll() if
// the designated constraints aren't met.
type TelemetryLogsProfileResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TelemetryLogsProfileResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TelemetryLogsProfileResourceMultiError) AllErrors() []error { return m }

// TelemetryLogsProfileResourceValidationError is the validation error returned
// by TelemetryLogsProfileResource.Validate if the designated constraints
// aren't met.
type TelemetryLogsProfileResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TelemetryLogsProfileResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TelemetryLogsProfileResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TelemetryLogsProfileResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TelemetryLogsProfileResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TelemetryLogsProfileResourceValidationError) ErrorName() string {
	return "TelemetryLogsProfileResourceValidationError"
}

// Error satisfies the builtin error interface
func (e TelemetryLogsProfileResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTelemetryLogsProfileResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TelemetryLogsProfileResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TelemetryLogsProfileResourceValidationError{}

var _TelemetryLogsProfileResource_ResourceId_Pattern = regexp.MustCompile("^telemetryprofile-[0-9a-f]{8}$")

var _TelemetryLogsProfileResource_ProfileId_Pattern = regexp.MustCompile("^telemetryprofile-[0-9a-f]{8}$")

var _TelemetryLogsProfileResource_TargetInstance_Pattern = regexp.MustCompile("^inst-[0-9a-f]{8}$")

var _TelemetryLogsProfileResource_TargetSite_Pattern = regexp.MustCompile("^site-[0-9a-f]{8}$")

var _TelemetryLogsProfileResource_TargetRegion_Pattern = regexp.MustCompile("^region-[0-9a-f]{8}$")

var _TelemetryLogsProfileResource_LogsGroupId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")

// Validate checks the field values on TelemetryMetricsProfileResource with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TelemetryMetricsProfileResource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TelemetryMetricsProfileResource with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// TelemetryMetricsProfileResourceMultiError, or nil if none found.
func (m *TelemetryMetricsProfileResource) ValidateAll() error {
	return m.validate(true)
}

func (m *TelemetryMetricsProfileResource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResourceId()) > 25 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "ResourceId",
			reason: "value length must be at most 25 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_ResourceId_Pattern.MatchString(m.GetResourceId()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "ResourceId",
			reason: "value does not match regex pattern \"^telemetryprofile-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetProfileId()) > 25 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "ProfileId",
			reason: "value length must be at most 25 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_ProfileId_Pattern.MatchString(m.GetProfileId()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "ProfileId",
			reason: "value does not match regex pattern \"^telemetryprofile-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetInstance()) > 13 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetInstance",
			reason: "value length must be at most 13 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_TargetInstance_Pattern.MatchString(m.GetTargetInstance()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetInstance",
			reason: "value does not match regex pattern \"^inst-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetSite()) > 13 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetSite",
			reason: "value length must be at most 13 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_TargetSite_Pattern.MatchString(m.GetTargetSite()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetSite",
			reason: "value does not match regex pattern \"^site-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTargetRegion()) > 15 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetRegion",
			reason: "value length must be at most 15 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_TargetRegion_Pattern.MatchString(m.GetTargetRegion()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "TargetRegion",
			reason: "value does not match regex pattern \"^region-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for MetricsInterval

	if len(m.GetMetricsGroupId()) > 23 {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "MetricsGroupId",
			reason: "value length must be at most 23 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_TelemetryMetricsProfileResource_MetricsGroupId_Pattern.MatchString(m.GetMetricsGroupId()) {
		err := TelemetryMetricsProfileResourceValidationError{
			field:  "MetricsGroupId",
			reason: "value does not match regex pattern \"^telemetrygroup-[0-9a-f]{8}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetricsGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TelemetryMetricsProfileResourceValidationError{
					field:  "MetricsGroup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TelemetryMetricsProfileResourceValidationError{
					field:  "MetricsGroup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetricsGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TelemetryMetricsProfileResourceValidationError{
				field:  "MetricsGroup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TelemetryMetricsProfileResourceMultiError(errors)
	}

	return nil
}

// TelemetryMetricsProfileResourceMultiError is an error wrapping multiple
// validation errors returned by TelemetryMetricsProfileResource.ValidateAll()
// if the designated constraints aren't met.
type TelemetryMetricsProfileResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TelemetryMetricsProfileResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TelemetryMetricsProfileResourceMultiError) AllErrors() []error { return m }

// TelemetryMetricsProfileResourceValidationError is the validation error
// returned by TelemetryMetricsProfileResource.Validate if the designated
// constraints aren't met.
type TelemetryMetricsProfileResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TelemetryMetricsProfileResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TelemetryMetricsProfileResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TelemetryMetricsProfileResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TelemetryMetricsProfileResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TelemetryMetricsProfileResourceValidationError) ErrorName() string {
	return "TelemetryMetricsProfileResourceValidationError"
}

// Error satisfies the builtin error interface
func (e TelemetryMetricsProfileResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTelemetryMetricsProfileResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TelemetryMetricsProfileResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TelemetryMetricsProfileResourceValidationError{}

var _TelemetryMetricsProfileResource_ResourceId_Pattern = regexp.MustCompile("^telemetryprofile-[0-9a-f]{8}$")

var _TelemetryMetricsProfileResource_ProfileId_Pattern = regexp.MustCompile("^telemetryprofile-[0-9a-f]{8}$")

var _TelemetryMetricsProfileResource_TargetInstance_Pattern = regexp.MustCompile("^inst-[0-9a-f]{8}$")

var _TelemetryMetricsProfileResource_TargetSite_Pattern = regexp.MustCompile("^site-[0-9a-f]{8}$")

var _TelemetryMetricsProfileResource_TargetRegion_Pattern = regexp.MustCompile("^region-[0-9a-f]{8}$")

var _TelemetryMetricsProfileResource_MetricsGroupId_Pattern = regexp.MustCompile("^telemetrygroup-[0-9a-f]{8}$")
