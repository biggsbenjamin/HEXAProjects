#!/usr/bin/env python

import rospy
from std_msgs.msg import Float64, Float64MultiArray, MultiArrayLayout, MultiArrayDimension
import random 
import numpy as np
import csv

#setting up a data array using the standard ros messages
poseArr =Float64MultiArray()
#adding dimensions to the array
poseArr.layout.dim.append(MultiArrayDimension())
poseArr.layout.dim.append(MultiArrayDimension())
#labelling the dimensions
poseArr.layout.dim[0].label = "angle"
poseArr.layout.dim[1].label = "millis"
#size of the dimensions
poseArr.layout.dim[0].size = 19
poseArr.layout.dim[1].size = 19
#not required by this code but throws error if removed
poseArr.layout.dim[0].stride = 2
poseArr.layout.dim[1].stride = 19
poseArr.layout.data_offset = 0
#sets all values to 90 degrees
poseArr.data = [90.0]*38


def arrayPub():
	#prepare to read in the joint angles and movement time from a csv file
	with open('catkin_ws/src/beginner_tutorials/scripts/simple_gait2.csv', 'rb') as arrayFile:
		arrayData = csv.reader(arrayFile, delimiter=',', quotechar='"')
		#prepare the publisher with topic name, message type 
		pub = rospy.Publisher('floater', Float64MultiArray, queue_size=10)
		rospy.init_node('talker', anonymous=True)

		rate = rospy.Rate(1) #hz
		#skips the title line of the csv file
		arrayData.next()
		#each row contains the values for one array which contains all joint positions and movement times for one pose
		for row in arrayData:
			i = 0
			#error handling
			if not rospy.is_shutdown():
				#iterate through each cell of the csv and assign it to the data of the pose array
				for val in row:
					if val != "":
						poseArr.data[i] = float(val)
					i+=1
				pub.publish(poseArr)
				print poseArr.data
				#sleeps so that loop occurs at the rospy rate 
				rate.sleep()
			else:
				break

if __name__ == '__main__':

	try:
		arrayPub()
#error handling
	except rospy.ROSInterruptException:
	        pass
