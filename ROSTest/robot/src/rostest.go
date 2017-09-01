// ROS Skill implements a skill that shows how to publish Strings to an ROS topic via rosserial.
//subscriber and publisher don't work at the same time
package example

// #include "rostest.h"
import "C"

import (
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"time"
)

//export stringFromC
func stringFromC() {
	hi := C.GoString(C.subArray)
	log.Info.Println(hi)
}

const (
	rosMasterIP = "192.168.0.234" // ROS_MASTER_IP need to be modified manually
	rosSubTopic  = "subTopic"
	rosPubTopic  = "pubTopic"
	hello = "HelloWorld"
)



type rostest struct {
	skill.Base
	StringSubscriber *C.StringSubscriber
	StringPublisher *C.StringPublisher
	stop chan bool
}

func NewSkill() skill.Interface {
	return &rostest{
		StringSubscriber: C.NewStringSubscriber(C.CString(rosMasterIP), C.CString(rosSubTopic)),
	StringPublisher: C.NewStringPublisher(C.CString(rosMasterIP), C.CString(rosPubTopic)),
		stop: make(chan bool),
	}
}





func (d *rostest) pubStrings() {
	go d.subStrings()
	time.Sleep(20000 * time.Millisecond)
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		C.PublishString(
			d.StringPublisher,
			C.CString(hello),
			C.int(len(hello)),
		)
		log.Info.Println("Sent String with length:", len(hello))
		time.Sleep(1000 * time.Millisecond)
	}
}

func (d *rostest) subStrings(){
	for {
		C.SubscribeString(d.StringSubscriber)

		//log.Info.Println("spun")
		time.Sleep(1000 * time.Millisecond)
	}
}


func (d *rostest) OnStart() {
	d.pubStrings()

}

func (d *rostest) OnClose() {
	d.stop <- true
	C.DeleteStringPublisher(d.StringPublisher)
	C.DeleteStringSubscriber(d.StringSubscriber)
}
