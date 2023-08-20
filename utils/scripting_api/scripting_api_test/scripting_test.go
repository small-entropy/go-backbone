package scripting_api

import (
	"fmt"
	"testing"

	ScriptingApi "github.com/small-entropy/go-backbone/utils/scripting_api"
	Interpreters "github.com/small-entropy/go-backbone/utils/scripting_api/interpreters"
)

func TestExec(t *testing.T) {

	// вызовы скриптов через os/exec должен делать вывод в stdout, вопрос открытый как перехватить
	// возврат функции/выражения без этого
	py_payload := ScriptingApi.NewPayload("print(sum([1,1]))", []string{"-c"})

	os_python := Interpreters.OsPython{}
	t.Logf("Call python subprocess with args: %v %v \n", py_payload.Flags, py_payload.Script)
	res := os_python.Exec(*py_payload)
	validate(t, "OsPython", res)

	go_python := Interpreters.GoPython{}
	t.Logf("Call klctl with args: %v %v \n", py_payload.Flags, py_payload.Script)
	res = go_python.Exec(*py_payload)
	validate(t, "GoPython", res)

	lua_payload := ScriptingApi.NewPayload("Hello from lua", []string{"print"})
	go_lua := Interpreters.GoLua{}
	t.Logf("Call go-lua with args: %v %v \n", lua_payload.Flags, lua_payload.Script)
	_ = go_lua.Exec(*lua_payload)

}

func validate(t *testing.T, executor_name string, res interface{}) {
	var res_int int
	var status string
	fmt.Sscan(res.(string), &res_int)

	if res_int == 2 {
		status = "ok"
	} else {
		status = "fail"
	}
	t.Logf("exec %v: %v", executor_name, status)
}
