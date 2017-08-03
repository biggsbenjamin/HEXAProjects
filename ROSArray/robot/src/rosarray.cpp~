#include <string>
#include <ros.h>
#include <embedded_linux_hardware.h>
#include <std_msgs/Float64.h>
#include <std_msgs/Float64MultiArray.h>
#include <std_msgs/MultiArrayDimension.h>
#include "_cgo_export.h"

//subscriber callback function
void messageCb(const std_msgs::Float64MultiArray& msg){
	//send data to go side
	for (int i=0;i<19;i++){
		for(int j=0;j<2;j++){
			floatFromC(msg.data[
			(msg.layout.data_offset) + 
			(i*msg.layout.dim[0].stride) + j],
			i, j );
		}
	}
	printPose();
}

class FloatSubscriberWrapper {
    private:

        ros::NodeHandle nh;
        ros::Subscriber<std_msgs::Float64MultiArray> sub;
    public:
        FloatSubscriberWrapper(char *ip, char *topic): sub(topic, messageCb) {
            nh.initNode(ip);
            nh.subscribe(sub);
        }
        void subscribeFloat() {
            nh.spinOnce();
        }
};

std::string test = "null";

char *subArray = &test[0];

/* Wrapper functions */

extern "C" FloatSubscriber* NewFloatSubscriber(char *ip, char *topic) {
    FloatSubscriber* flo_sub = new FloatSubscriber();
    flo_sub->wrapper = new FloatSubscriberWrapper(ip, topic);
    return flo_sub;
}

extern "C" void SubscribeFloat(FloatSubscriber* flo_sub) {
    ((FloatSubscriberWrapper *)flo_sub->wrapper)->subscribeFloat();
}

extern "C" void DeleteFloatSubscriber(FloatSubscriber* flo_sub) {
    FloatSubscriberWrapper* wrapper = (FloatSubscriberWrapper *)(flo_sub->wrapper);
    delete wrapper;
    delete flo_sub;
}
