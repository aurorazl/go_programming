#!/bin/sh

pwd
ls -la mau_shell

if [[ $# -ge 1 ]]; then
    date=${1}
else
    date=`date +%Y-%m-%d -d "-1 day"`
fi

check_exist()
{
    echo "check if path: "$1" exist..."
    while [[ 1 == 1 ]]
    do
        /usr/local/service/hadoop/bin/hadoop fs -test -e $1/_SUCCESS
        if [[ $? == 0 ]]
        then
            echo "Path "$1" exist, continue"
            break
        fi
        echo "Path "$1" does not exist, wait"
        sleep 60
    done
}#!/bin/sh

if [[ $# -ge 1 ]]; then
    date=${1}
else
    date=`date +%Y-%m-%d -d "-1 day"`
fi

check_exist()
{
    echo "check if path: "$1" exist..."
    while [[ 1 == 1 ]]
    do
        /usr/local/service/hadoop/bin/hadoop fs -test -e $1/_SUCCESS
        if [[ $? == 0 ]]
        then
            echo "Path "$1" exist, continue"
            break
        fi
        echo "Path "$1" does not exist, wait"
        sleep 60
    done
}

slidingMauPath="/user/hadoop/bigdata/zgstat/sliding_mau_attribute_from_clickHouse/dt=${date}"
echo "slidingMauPath: ${slidingMauPath}"
check_exist ${slidingMauPath}

/usr/local/service/hive/bin/hive -e "alter table sliding_mau_attribute add if not exists partition(dt='$date') location 'hdfs://HDFS14992/user/hadoop/bigdata/zgstat/sliding_mau_attribute_from_clickHouse/dt=$date';"

sleep 30

/root/zero/WaterDrop/waterdrop-dist-2.0.4-2.11.8/bin/start-waterdrop-spark.sh --master yarn --deploy-mode client --config mau_shell/export_sliding_mau_attribute.conf

/root/zero/WaterDrop/waterdrop-dist-2.0.4-2.11.8/bin/start-waterdrop-spark.sh --master yarn --deploy-mode client --config mau_shell/export_sliding_mau_attribute_new.conf
