#!/bin/bash

reso_addr='registry.cn-hangzhou.aliyuncs.com/im-zero/user-rpc-dev'
tag='latest'

comtainer_name='im-zero-user-rpc-test'

docker stop ${comtainer_name}

docker rm ${comtainer_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}

docker run -p 10000:8080 --name=${comtainer_name} -d ${reso_addr}:${tag}