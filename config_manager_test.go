package fortio

import "testing"

func TestCamelCaseToUnderscore(t *testing.T) {
    var testCases = [][]string{
        {"ServiceName", "SERVICE_NAME"},
        {"name", "NAME"},
        {"Address1", "ADDRESS1"},
        {"primaryAddress", "PRIMARY_ADDRESS"},
    }

	for _, test := range testCases {
		out := camelCaseToUnderscore(test[0])
		if out != test[1] {
			t.Errorf("Expecting underscore version of camel case to be %s, but got %s", test[1], out)
		}
	}
}

func TestCamelCaseToLowerFirst(t *testing.T) {

    var testCases = [][]string{
        {"Name", "name"},
        {"ServiceName", "serviceName"},
        {"AWSInfo", "awsInfo"},
        {"AWSInstanceName", "awsInstanceName"},
    }

    for _, test := range testCases {
        out := lowerFirst(test[0])
        if out != test[1] {
            t.Errorf("Expecting lowerFirst version of camel case to be %s, but got %s", test[1], out)
        }
    }
}

func TestRequired(t *testing.T) {

}
