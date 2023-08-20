package interpreters

import (
	"strings"

	Python "github.com/kluctl/go-embed-python/python"
	ScriptingApi "github.com/small-entropy/go-backbone/utils/scripting_api"
)

type GoPython struct {
}

func (interp GoPython) Exec(input ScriptingApi.Payload) any {
	ep, err := Python.NewEmbeddedPython("example")
	if err != nil {
		panic(err)
	}

	cmd := ep.PythonCmd(strings.Join(input.Flags, " "), input.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(out)
}
