rm ${JOB_NAME}.yml || true
rm change_image.sh || true
wget --no-cookie --no-check-certificate https://raw.githubusercontent.com/brynelee/fsdemo-usercenter/master/fsdemo-common-tools/docker/change_image.sh
wget --no-cookie --no-check-certificate https://raw.githubusercontent.com/brynelee/fsdemo-usercenter/master/fsdemo-common-tools/kubernetes/${JOB_NAME}.yml
kubectl delete -f ${JOB_NAME}".yml" || true
sleep 15
docker rmi registry.cn-hangzhou.aliyuncs.com/xdorg1/${JOB_NAME} || true
docker rmi xdorg1/${JOB_NAME} || true
docker build -t xdorg1/fsdemo-priceservice -f Dockerfile.cn .
chmod +x change_image.sh
sh change_image.sh ${JOB_NAME}
kubectl apply -f ${JOB_NAME}".yml"