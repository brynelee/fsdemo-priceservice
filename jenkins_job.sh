#!/bin/sh

# clean the tool files before proceeding
rm $1.yml || true
rm change_image.sh || true

# get the change_image.sh and yaml file
git clone https://github.com/brynelee/fsdemo-usercenter.git
cp "${WORKSPACE}/fsdemo-usercenter/fsdemo-common-tools/docker/change_image.sh" .
cp "${WORKSPACE}/fsdemo-usercenter/fsdemo-common-tools/kubernetes/"$1".yml" .
rm -rf "${WORKSPACE}/fsdemo-usercenter"
#wget --no-cookie --no-check-certificate https://raw.githubusercontent.com/brynelee/fsdemo-usercenter/master/fsdemo-common-tools/docker/change_image.sh
#wget --no-cookie --no-check-certificate https://raw.githubusercontent.com/brynelee/fsdemo-usercenter/master/fsdemo-common-tools/kubernetes/$1.yml
kubectl delete -f $1".yml" || true
sleep 15
docker rmi registry.cn-hangzhou.aliyuncs.com/xdorg1/$1 || true
docker rmi xdorg1/$1 || true
docker build -t xdorg1/$1 -f Dockerfile.cn .
chmod +x change_image.sh
sh change_image.sh $1
kubectl apply -f $1".yml"