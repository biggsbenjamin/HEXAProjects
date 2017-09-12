#include <string>
#include <ros.h>
#include <embedded_linux_hardware.h>
#include <std_msgs/String.h>
#include "_cgo_export.h"



void messageCb(const std_msgs::String& received_msg){
	//send data to go side
	std::string data = received_msg.data;
	subArray = &data[0];
	stringFromC();
		
}

class StringNodeWrapper {
	private:
    	ros::NodeHandle nh;
	ros::Subscriber<std_msgs::String> sub;
	std_msgs::String str_msg;
	ros::Publisher publisher;

	public:
        StringNodeWrapper(char *ip, char *subtopic, char *pubtopic): sub(subtopic, messageCb), publisher(pubtopic, &str_msg){
		nh.initNode(ip);
		nh.advertise(publisher);
	        nh.subscribe(sub);
		
        }
        void subscribeString() {
        	nh.spinOnce();
        }
        void publishString(char* data, int len) {
            	str_msg.data = (char*)data;
            	publisher.publish(&str_msg);
		nh.spinOnce();
        }
};

std::string test = "null";

char *subArray = &test[0];

/* Wrapper functions */

extern "C" StringNode* NewStringNode(char *ip, char *subtopic, char *pubtopic) {
    StringNode* str_node = new StringNode();
    str_node->wrapper = new StringNodeWrapper(ip, subtopic, pubtopic);
    return str_node;
}

extern "C" void SubscribeString(StringNode* str_node) {
    ((StringNodeWrapper *)str_node->wrapper)->subscribeString();
}

extern "C" void DeleteStringNode(StringNode* str_node) {
    StringNodeWrapper* wrapper = (StringNodeWrapper *)(str_node->wrapper);
    delete wrapper;
    delete str_node;
}

extern "C" void PublishString(StringNode* str_node, char *data, int len) {
    ((StringNodeWrapper *)str_node->wrapper)->publishString(data, len);
}

