# With thanks to Joseph Redmon, https://pjreddie.com/

if test -f "./yolov3.weights"; then
  echo "neural network weights loaded"
else
  echo "downloading neural network weights"
  wget https://pjreddie.com/media/files/yolov3.weights
  wget https://github.com/pjreddie/darknet/blob/master/cfg/yolov3.cfg?raw=true -O "./yolov3.cfg"
  wget https://github.com/pjreddie/darknet/blob/master/data/coco.names?raw=true -O "./coco.names"
fi