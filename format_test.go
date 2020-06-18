package format

import (
	"testing"
)

func TestFormatFunctionality(t *testing.T) {

	formattedString := String(`test {key1} {key1} {key2} "{key3}" {missingKey}`, Items{"key1": 2, "key2": true, "key3": "123"})
	expectedResult := `test 2 2 true "123" {missingKey}`

	if formattedString != expectedResult {
		t.Errorf(`invalid message format. expected "%v" and got "%v"`, expectedResult, formattedString)
	}

}
