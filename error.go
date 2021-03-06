package vice

import (
	"github.com/jbowes/vice/skip"
)

// New returns an error that formats as the given text, and implements the
// behaviour described by the given Vice.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
func New(v Vice, text string) error { return skip.New(skip.Vice(v), 1, text) }

// Errorf returns an error that formats as the format specifier text, and
// implements the behaviour described by the given Vice.
//
// The returned error contains a Frame set to the caller's location and
// implements Formatter to show this information when printed with details.
func Errorf(v Vice, format string, a ...interface{}) error {
	return skip.Errorf(skip.Vice(v), 1, format, a...)
}

// Wrap returns an error wrapping err with the supplied text, and a frame
// from the caller's stack. The returned error implements the behaviour
// described by the given Vice. If err is nil, Wrap returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
func Wrap(err error, v Vice, text string) error {
	return skip.Wrap(err, skip.Vice(v), 1, text)
}

// Wrapf returns an error wrapping err with the supplied format specifier, and
// a frame from the caller's stack. The returned error implements the behaviour
// described by the given Vice. If err is nil, Wrap returns nil.
//
// The error returned implments the Unwrap method, for programatically
// extracting the error chain.
func Wrapf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Wrapf(err, skip.Vice(v), 1, format, a...)
}

// Seal returns an error wrapping err with the supplied text, and
// a frame from the caller's stack. The returned error implements the behaviour
// described by the given Vice. If err is nil, Wrap returns nil.
//
// The error returned does not implment the Unwrap method.
func Seal(err error, v Vice, text string) error {
	return skip.Seal(err, skip.Vice(v), 1, text)
}

// Sealf returns an error wrapping err with the supplied format specifier, and
// a frame from the caller's stack. The returned error implements the behaviour
// described by the given Vice. If err is nil, Wrap returns nil.
//
// The error returned does not implment the Unwrap method.
func Sealf(err error, v Vice, format string, a ...interface{}) error {
	return skip.Sealf(err, skip.Vice(v), 1, format, a...)
}
