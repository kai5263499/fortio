package fortio

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCmdLineConfigLoader(t *testing.T) {
	c := &Conf{}
	cmdLineLoader := NewCmdLineConfigLoader()
	cm := NewConfigManager("fortio-test", "My Fortio test", cmdLineLoader)
	err := cm.load(c, false)
	if err != nil {
		t.Errorf("Config loading not supposed to fail - %s", err.Error())
	}
	j, err := c.DumpJSON()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Loaded config: %s\n", j)

	if c.Name != "my name" {
		t.Errorf("Name is not loaded correctly")
	}
	if c.Int8 != 11 {
		t.Errorf("Int8 is not loaded correctly")
	}
	if c.Int32 != 12 {
		t.Errorf("Int32 is not loaded correctly")
	}
	if c.Int64 != 30 {
		t.Errorf("Int64 is not loaded correctly")
	}
	if c.Number != -10 {
		t.Errorf("Number is not loaded correctly")
	}
	if c.UInt != 15 {
		t.Errorf("UInt is not loaded correctly")
	}
	if c.UInt64 != 10 {
		t.Errorf("UInt64 is not loaded correctly")
	}
	if c.Float != 9.99 {
		t.Errorf("Float is not loaded correctly")
	}
	if !c.Bool {
		t.Errorf("Bool is not loaded correctly")
	}
	if c.Duration.Duration != time.Duration(100)*time.Millisecond {
		t.Errorf("Duration is not loaded correctly - %v", c.Duration.Duration)
	}
	if c.Array == nil || len(c.Array) == 0 {
		t.Errorf("StringList is not loaded correctly")
	}

	for i, v := range c.Array {
		if (i == 0 && v != "foo") || (i == 1 && v != "bar") || (i == 2 && v != "baz") {
			t.Errorf("StringList values not matched")
		}
	}
}

type Conf struct {
	Name     string     `config:";default=my name;usage=Give me a name"`
	Number   int        `config:";default=-10;usage=Give me a int"`
	Int8     int8       `config:";default=11;usage=Give me a int8"`
	Int32    int32      `config:";default=12;usage=Give me a int32"`
	Int64    int64      `config:";default=30;usage=Give me a int64"`
	UInt     uint       `config:";default=15;usage=Give me a uint"`
	UInt8    uint8      `config:";default=15;usage=Give me a uint8"`
	UInt16   uint16     `config:";default=16;usage=Give me a uint16"`
	UInt32   uint32     `config:";default=17;usage=Give me a uint32"`
	UInt64   uint64     `config:";default=10;usage=Give me a uint"`
	Float32  float32    `config:";default=8.99;usage=Give me a float32"`
	Float    float64    `config:";default=9.99;usage=Give me a float"`
	Bool     bool       `config:";default=true;usage=Give me a bool"`
	Array    StringList `config:";default=foo,bar,baz;usage=Give me a string array"`
	Duration Duration   `config:";default=100ms;usage=Give me some duration"`
}

func (c *Conf) Validate() error {
	return nil
}

func (c *Conf) DumpJSON() (string, error) {
	b, err := json.MarshalIndent(c, "", "  ")
	return string(b), err
}

func (c *Conf) DumpYAML() (string, error) {
	return "", nil
}
