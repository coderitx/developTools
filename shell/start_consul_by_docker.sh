#!/bin/sh

function start_consul(){
  docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul:latest agent -server -bootstrap-expect 2 -ui -bind=0.0.0.0 -client=0.0.0.0
  consul1_addr=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul1)
  for((i=1;i<=10;i++));
  do
    port=$((8500+$i))
    docker run --name consul${i} -d -p ${port}:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join ${consul1_addr}
  done
}

function check_consul(){
  docker ps | grep "consul" | grep -v grep
  docker exec -it consul1 "consul members"
}

start_consul
check_consul
