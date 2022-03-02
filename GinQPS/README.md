
## 压测介绍：
go get -u github.com/gin-gonic/gin

数据库单表百万量级，200线程 * 500轮 的随机主键id查表压测
![image-20220302212918808](https://raw.githubusercontent.com/aurorazl/markdown/main/image-20220302212918808.png)
平均耗时为113毫秒，最高为1.7秒, mysql配置为12G + 4核，QPS峰值为3.5K左右

##使用的gorm框架报错：
    错误1：
        Too many connections
        原因：
            mysql没有配置连接数上限max_connections，默认为152。gorm pool的上限应该小于这个。
            Threads_connected由max_connections决定。
        解决： 
            gorm设置连接池上限、idle时间、lifetime
    错误2：
        dial tcp 10.0.0.15:3306: connect: can't assign requested address
        原因：
            gorm每次新建本地socket，本机ulimit为10240，socket多数处于TIME_WAIT状态，压测连接数超过文件描述符上限（未超port上限）
            netstat -tn |grep 3306 | wc -l
        解决：
            服务器ulimit -n 50000调高上限后情况减缓。
            gorm 1.20版本后支持连接池，默认没有上限，可以设置连接上限。

## gin请求监控:  
    go get github.com/zsais/go-gin-prometheus  
    prometheus配置抓取
    其他方法：gorm plugin：Prometheus
## mysql监控:  
    grafana + prometheus + mysqld-exporter

## 非主键字段查找：
![image-20220302212040896](https://raw.githubusercontent.com/aurorazl/markdown/main/image-20220302212040896.png)
平均耗时为20秒，最低耗时0.4秒

## 该字段加索引后：
ALTER TABLE qps.user ADD INDEX (NAME) 14秒左右完成
![image-20220302212535994](https://raw.githubusercontent.com/aurorazl/markdown/main/image-20220302212535994.png)
耗时大幅下降，最高也不过1秒。

## 走redis查询主键：
![image-20220302232705848](https://raw.githubusercontent.com/aurorazl/markdown/main/image-20220302232705848.png)
qps明显上升，局限于机器的性能