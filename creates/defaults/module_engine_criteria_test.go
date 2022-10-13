package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
)

func TestModule_engineCriteria_Success(t *testing.T) {
	script := `
		-> $name;;
        -> $index;;
        -> $includeChannels;;
		<- $criteria;;

        // criteria app
        module @engineCriteria;;
		@engineCriteria $criteriaApp;;
		attach $name:$name $criteriaApp;;
		attach $index:$index $criteriaApp;;
        attach $includeChannels:$includeChannels $criteriaApp;;
        $criteria = execute $criteriaApp;;

	`

	name := "roger"
	index := uint(0)
	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{
		"name":            name,
		"index":           index,
		"includeChannels": true,
	}, []byte(script))

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["criteria"].(criterias.Criteria); ok {
		if ins.Name() != name {
			t.Errorf("the name was expected to be '%s', '%s' returned", name, ins.Name())
			return
		}

		return
	}

	t.Errorf("the criteria output was expected to contain a Criteria instance")
	return
}
