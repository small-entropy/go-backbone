package interpreters

import (
	"fmt"
	"os/exec"
	"strings"

	ScriptingApi "github.com/small-entropy/go-backbone/utils/scripting_api"
)

type OsPython struct {
}

func (interp OsPython) Exec(input ScriptingApi.Payload) any {
	cmd := exec.Command("python", strings.Join(input.Flags, " "), input.Script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}
