// ROS Skill implements a skill that shows how to publish Strings to an ROS topic via rosserial.
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
	rosMasterIP = "192.168.0.102" // ROS_MASTER_IP need to be modified manually
	rosPubTopic  = "chatter1"
	hello = "HelloWorld"
)



type rostest struct {
	skill.Base
	//StringPublisher *C.StringPublisher
	StringSubscriber *C.StringSubscriber
	stop           chan bool
}

func NewSkill() skill.Interface {
	return &rostest{
		//StringPublisher: C.NewStringPublisher(C.CString(rosMasterIP), C.CString(rosPubTopic)),

		StringSubscriber: C.NewStringSubscriber(C.CString(rosMasterIP), C.CString(rosSubTopic)),
		stop:           make(chan bool),
	}
}





func (d *rostest) subPubStrings() {
	for {
		select {
		case <-d.stop:
			break
		default:
		}
		/* C.PublishString(
			d.StringPublisher,
			C.CString(hello),
			C.int(len(hello)),
		)
		log.Info.Println("Sent String with length:", len(hello))
		*/
		C.SubscribeString(d.StringSubscriber)

		//log.Info.Println("spun")
		time.Sleep(1000 * time.Millisecond)
	}
}


func (d *rostest) OnStart() {
	d.subPubStrings()
	log.Info.Println("starting subscriber")
}

func (d *rostest) OnClose() {
	d.stop <- true
	//C.DeleteStringPublisher(d.StringPublisher)
	C.DeleteStringSubscriber(d.StringSubscriber)
}
