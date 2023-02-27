#!/bin/sh

# get image
docker images | grep none | grep -v grep
if [ "$?" == 1 ];then
  echo "image not found..."
  exit 0
fi

# stop container
docker stop $(docker ps -a | grep "Exited" | grep -c grep | awk '{print $1}')
if [ "$?" = 0 ];then
  echo "stop container success..."
else
  echo "stop container failed..."
fi

# rm container
docker rm $(docker ps -a | grep "Exited" | grep -c grep | awk '{print $1}')
if [ "$?" = 0 ];then
  echo "rm container success..."
else
  echo "rm container failed..."
fi


# remove image
docker rmi $(docker images | grep "none" | grep -v grep | awk '{print $3}')
if [ "$?" =0 ];then
  echo "remove image by image REPOSITORY is none success..."
else
  echo "remove image by image REPOSITORY is none failed..."
fi
