#!/usr/bin/env python


import rospy
from std_msgs.msg import Float64, Float64MultiArray, MultiArrayLayout, MultiArrayDimension
import random 
import numpy as np
import csv


testArr =Float64MultiArray()
testArr.layout.dim.append(MultiArrayDimension())
testArr.layout.dim.append(MultiArrayDimension())
testArr.layout.dim[0].label = "angle"
testArr.layout.dim[1].label = "millis"

testArr.layout.dim[0].size = 19
testArr.layout.dim[1].size = 19

testArr.layout.dim[0].stride = 2
testArr.layout.dim[1].stride = 19
testArr.layout.data_offset = 0
testArr.data = [0]*38

dstride0 = testArr.layout.dim[0].stride
dstride1 = testArr.layout.dim[1].stride
offset = testArr.layout.data_offset

oscillateAng = [70.0,110.0]

def talker():
	with open('simple_4_leg_gait.csv', 'rb') as arrayFile:
		arrayData = csv.reader(arrayFile, delimiter=',', quotechar='"')
		pub = rospy.Publisher('floater', Float64MultiArray, queue_size=10)
		rospy.init_node('talker', anonymous=True)
		rate = rospy.Rate(0.25) #hz
		tmpmat = np.zeros((19,2))
		for row in arrayData:
			print row
			if not rospy.is_shutdown():
					for i in range(19):
						for j in range(2):
							testArr.data[offset + i*dstride0 + j] = 90.0
							tmpmat[i,j] = 90.0

					pub.publish(testArr)
					print tmpmat,"\r\n"
					rate.sleep()
			else:
				break

if __name__ == '__main__':

	try:
		talker()
	except rospy.ROSInterruptException:
	        pass
