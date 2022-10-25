package defaults

import (
	"fmt"
	"io/ioutil"
	"testing"

	engines "github.com/steve-care-software/webx/applications"
	"github.com/steve-care-software/webx/creates/disks/defaults"
)

const bitrateForTests = 1024
const basePathForTests = "./test_files"
const delimiterForTests = "."
const extensionForTests = "identity"

func TestRodan_Success(t *testing.T) {
	script := `
        $value = 45;
	`

	svmInstructions, err := ioutil.ReadFile("./rodan.svm")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	output, _, err := engines.NewApplication(defaults.NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests)).ParseThenInterpret(map[string]interface{}{
		"script": script,
	}, []byte(svmInstructions))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", output)
}
