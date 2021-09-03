#广告信息平台

```
docker run --name wzb-mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:8
docker run --name wzb-redis -p 6379:6379 -d redis:6
redis-cli -h 8.136.119.124 -p 6379


tail -f /var/log/ad_platform_info/2021-08-01_17.log
```

部署hadoop集群
1. 创建网络

`docker network create --driver bridge --subnet=192.168.0.0/16 --gateway=192.168.1.1 mynet`
2. 创建数据卷

`docker volume create hadoop-data`
3. 查看数据卷挂载路径

```
dcoker volume inspect hadoop-data
ln -s /var/lib/docker/volumes/hadoop-data/_data/ /modules
```
4. 根据dockerfile创建镜像

`docker build -t hadoop:v1 .`

dockerfile文件如下
```dockerfile
FROM centos:8
ENV JAVA_HOME=/modules/java
ENV HADOOP_HOME=/modules/hadoop
ENV PATH=${PATH}:${JAVA_HOME}/bin:${HADOOP_HOME}/bin:${HADOOP_HOME}/sbin
ENV TZ=Asia/Shanghai
RUN yum install -y net-tools \
&& yum install -y openssh-server \
&& yum install -y openssh-clients \
&& yum install -y passwd \
&& echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
```
创建环境变量文件
cd /modules
touch env-hadoop.sh
```shell
#!/bin/bash
export JAVA_HOME=/modules/java
export HADOOP_HOME=/modules/hadoop
export PATH=$PATH:$JAVA_HOME/bin:$HADOOP_HOME/bin:$HADOOP_HOME/sbin
export TZ=Asia/Shanghai
```

5. 部署hadoop集群


```shell
#部署数据卷容器: 
docker run -itd --name hadoop-data-container --privileged -p 30000:22 -h hadoop.com  --network=mynet --ip 192.168.100.100 -v hadoop-data:/modules hadoop:v2 /usr/sbin/init
#部署hadoop01:
docker run -itd --name hadoop01 --privileged -p 30001:22 -h hadoop01.com  --network=mynet --ip 192.168.100.101  --volumes-from hadoop-data-container hadoop:v2 /usr/sbin/init
#部署hadoop02: 
docker run -itd --name hadoop02 --privileged -p 30002:22 -h hadoop02.com  --network=mynet --ip 192.168.100.102 --volumes-from hadoop-data-container hadoop:v2 /usr/sbin/init
#部署hadoop03:
docker run -itd --name hadoop03 --privileged -p 30003:22 -h hadoop03.com  --network=mynet --ip 192.168.100.103 --volumes-from hadoop-data-container hadoop:v2 /usr/sbin/init
```
每次部署完后，进入容器设置root密码，并重启sshd服务
```shell
docker exec -it hadoop01 /bin/bash
passwd
systemctl restart sshd.service
cp /modules/env-hadoop.sh /etc/profile.d
source /etc/profile
```
ssh 连接失败时：
```shell
# 删除掉所有的
vi ~/.ssh/known_hosts         
```

hadoop的使用
```shell
hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-3.3.1.jar wordcount wcinput/ wcoutput
```

```shell
ssh-keygen -t rsa
ssh-copy-id hadoop01
```

