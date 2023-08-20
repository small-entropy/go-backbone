package interpreters

import (
	lua "github.com/Shopify/go-lua"
	ScriptingApi "github.com/small-entropy/go-backbone/utils/scripting_api"
)

type GoLua struct {
}

func (interp GoLua) Exec(input ScriptingApi.Payload) any {
	l := lua.NewState()
	lua.OpenLibraries(l)
	l.Global("print")
	l.PushString(input.Script)
	l.Call(1, 0)

	return 0
}
