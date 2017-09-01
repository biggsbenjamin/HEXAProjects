/*
* MoveLegsSkill implements a skill
* that makes the HEXA to move its right
* arm up and down and left arm in a circle
 */

//edited version of hexa example skill

package examples

import (
	"math"
	"os"
	"strings"
	"strconv"

	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/skill"
	"mind/core/framework/log"
)

type MoveLegsSkill struct {
	skill.Base
	stop chan bool
}

func NewSkill() skill.Interface {
	return &MoveLegsSkill{
		stop: make(chan bool),
	}
}
//from example
func ready() {
	hexabody.Stand()
	hexabody.MoveHead(0.0, 1)
	hexabody.MoveLeg(2, hexabody.NewLegPosition(-100, 50.0, 70.0), 1)
	hexabody.MoveLeg(5, hexabody.NewLegPosition(100, 50.0, 70.0), 1)
	hexabody.MoveJoint(0, 1, 90, 1)
	hexabody.MoveJoint(0, 2, 45, 1)
	hexabody.MoveJoint(1, 1, 90, 1)
	hexabody.MoveJoint(1, 2, 45, 1)
}

//from example
func moveLegs(v float64) {
	hexabody.MoveJoint(0, 1, 45*math.Sin(v*math.Pi/180)+70, 1)
	hexabody.MoveJoint(0, 0, 35*math.Cos(v*math.Pi/180)+60, 1)
	hexabody.MoveJoint(1, 1, 45*math.Cos(v*math.Pi/180)+70, 1)
	hexabody.MoveJoint(1, 1, 45*math.Cos(v*math.Pi/180)+70, 1)
}

func legTest(data string) {
	//splits incoming string info into a slice (which is the equivalent of a c++ vector or python list)
	pos_slice := strings.Split(data, ",")
	//for loop iterates through each value in slice, "joint" var is the index
	for joint, angle := range pos_slice {
		//slice of strings converted to float
		ang, _ := strconv.ParseFloat(angle, 64)
		log.Info.Println(joint, ang)
		//moves individual hexa joint
		hexabody.MoveJoint(0, joint, ang, 1)
	}
}
//from example
func (d *MoveLegsSkill) play() {
	ready()
	v := 0.0
	for {
		select {
		case <-d.stop:
			return
		default:
			moveLegs(v)
			v += 1
		}
	}
}

//initialises the hexa motors
func (d *MoveLegsSkill) OnStart() {
	hexabody.Start()
}
//terminates the hexa motors
func (d *MoveLegsSkill) OnClose() {
	hexabody.Close()
}

func (d *MoveLegsSkill) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}
//function runs when strings are sent from a webpage
func (d *MoveLegsSkill) OnRecvString(data string) {
	log.Info.Println(data)
	
	if data == "start" {
		//starts thread that runs the play function as part of the d class
		go d.play()
	} else if data == "stop" {
		d.stop <- true
		hexabody.RelaxLegs()
	} else {
		log.Info.Println("starting")
		legTest(data)
		log.Info.Println("returned")
	}

}
