package scripting_api

type ScriptExecutor interface {
	Exec(input Payload) any
}

type Payload struct {
	Script string
	Flags  []string
}

func NewPayload(script string, flags []string) *Payload {
	return &Payload{
		Script: script,
		Flags:  flags,
	}
}
