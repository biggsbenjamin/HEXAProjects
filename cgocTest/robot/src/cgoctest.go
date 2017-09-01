package cgocTest

// #include "cgoctest.h"
import "C"
// ^ required to use c++ functions in go
import (
	"mind/core/framework/skill"
	"mind/core/framework/log"
)

//export intFromGo
func intFromGo(hello int) {
	//int type is easy to convert as it can be directly translated across
	log.Info.Println(hello)
}

//export stringFromGo
func stringFromGo() {
	hi := C.GoString(C.hello)
	//see the header file (.h) for explanation
	log.Info.Println(hi)
}


type cgocTest struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &cgocTest{}
}

func (d *cgocTest) OnStart() {
	// Use this method to do something when this skill is starting.
	//running c++ code
	C.main2()
}



