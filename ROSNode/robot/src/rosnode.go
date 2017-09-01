package ROSNode

import (
	"mind/core/framework/skill"
)

type ROSNode struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &ROSNode{}
}

func (d *ROSNode) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d *ROSNode) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d *ROSNode) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d *ROSNode) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d *ROSNode) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *ROSNode) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
