package fortio

import (
	"encoding/json"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
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

// StringParsable interface can be implemented by any struct that needs to be
// loaded as string as is and not go through its fields recursively
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
	return d.ParseString(s)
}

func (d *Duration) ParseString(s string) error {
	v, err := time.ParseDuration(s)
	*d = Duration{v}
	return err
}

// Type returns type name
func (d *Duration) Type() string {
	return "fortio.Duration"
}

// MapObject implements StringParsable and read int a JSON object into
// map[string]interface{}
type MapObject struct {
	Mapping map[string]interface{}
}

// ParseString exists to satisfy the StringParsable interface
func (jm *MapObject) ParseString(s string) error {
	jsonMap := &MapObject{}
	mapping := map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &mapping)
	if err != nil {
		err := yaml.Unmarshal([]byte(s), &mapping)
		if err != nil {
			return err
		}
	}
	*jm = *jsonMap
	jm.Mapping = mapping
	return nil
}

// String returns JSON string of the MapObject
func (jm *MapObject) String() string {
	b, err := json.Marshal(jm.Mapping)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// Set will set the JSON/YAML string to MapObject
func (jm *MapObject) Set(s string) error {
	return jm.ParseString(s)

}

// Type returns type of the object
func (jm *MapObject) Type() string {
	return "fortio.MapObject"
}
