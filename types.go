package fortio

import (
    "strings"
    "time"
)

// StringList is of type []string
type StringList []string

// String returns the string value of a StringList
func (sl *StringList) String() string {
    return strings.Join(*sl, ",")
}

// Set sets the value of the StringList given the argument from cmdline
func (sl *StringList) Set(s string) error {
    if s == "" {
        return nil
    }
    st := &StringList{}
    *sl = *st
    for _, v := range strings.Split(strings.TrimPrefix(strings.TrimSuffix(s, "\""), "\""), ",") {
        *sl = append(*sl, strings.TrimSpace(v))
    }
    return nil
}

// Type returns type of StringList
func (sl *StringList) Type() string {
    return "fortio.StringList"
}

type StringParsable interface {
    ParseString(string) error
}

// Duration is wrapper on time.Duration to support via
// new config structure for auto wiring
type Duration struct {
    time.Duration
}

// String return string version of Duration, same as time.Duration
func (d *Duration) String() string { return d.Duration.String() }

// Set will parse given string as time.Duration and sets to Duration
func (d *Duration) Set(s string) error {
    v, err := time.ParseDuration(s)
    *d = Duration{v}
    return err
}

// Type returns type name
func (d *Duration) Type() string {
    return "fortio.Duration"
}