package fortio

import "testing"

var testCases = [][]string{
	{"ServiceName", "SERVICE_NAME"},
	{"name", "NAME"},
	{"Address1", "ADDRESS1"},
	{"primaryAddress", "PRIMARY_ADDRESS"},
}

func TestCamelCaseToUnderscore(t *testing.T) {
	for _, test := range testCases {
		out := camelCaseToUnderscore(test[0])
		if out != test[1] {
			t.Errorf("Expecting underscore version of camel case to be %s, but got %s", test[1], out)
		}
	}
}
