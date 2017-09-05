package fortio

import "strings"

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
    return "config.StringList"
}

type StringParsable interface {
    ParseString(string) error
}
