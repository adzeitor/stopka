package main

import (
	"fmt"
	"syscall/js"

	"github.com/adzeitor/stopka"
)

func main() {
	machine := stopka.New()
	js.Global().Set("stopka", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		machine.Eval(args[0].String())
		output := fmt.Sprint(machine.Stack())
		if machine.IsHalted() {
			output += "( exception:  " + machine.Err.Error() + ")"
		}
		return output
	}))
	select {} // Code must not finish
}
