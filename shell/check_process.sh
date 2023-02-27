#!/bin/sh

process_name=$1
workdir=/root/workdir

functionc check_status(){
  ps -ef | grep "${process_name}" | grep -v grep
  status=$?
  if [ ${status} = 0 ];then
    echo "${process_name} running..."
  else
    cd ${workdir}
    nohup ${process_name} >> /dev/null 2>&1 &
    if [ $? = 0 ];then
      echo "${process_name} restart running success."
    else
      echo "{process_name} restart failed"
    fi
  fi
}

check_status
