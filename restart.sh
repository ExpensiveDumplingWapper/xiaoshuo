#!bin/bash

while :

do

if [ $(ps -ef | grep "./main" | grep -v "grep" | wc -l) -eq 1 ];then

kill $(ps -ef|grep "./main"|awk '{print $2}')

echo "kill"

nohup ./main > /tmp/`date "+%Y-%m-%d"`.web &

echo "restart"

else

echo "not found"

fi

sleep 2

done