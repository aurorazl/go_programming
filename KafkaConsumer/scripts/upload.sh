set -e
scp -r ../KafkaConsumer root@10.0.0.16:/home/hadoop/chenzenglong/
# ssh root@10.0.0.16 "cd /home/hadoop/chenzenglong/KafkaConsumer && sh scripts/build.sh"