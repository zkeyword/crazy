#!/bin/bash
remote_workdir=/data/deploy/target/pre/api
local_target=target/pre_api
remote_host=root@127.0.0.1
current_app=pre_api
new_app=pre_api.new
startShell=api_pre_start.sh
config=pre_api.json
log=pre-api-nohup.out
echo "release pre"
GOOS=linux GOARCH=amd64 go build -o $local_target cmd/api/main.go

echo ">>> scp"
scp  setting/$config $remote_host:$remote_workdir/$config
scp  script/$startShell $remote_host:$remote_workdir/$startShell
scp  $local_target $remote_host:$remote_workdir/$new_app

echo ">>> remote operate"
echo ">>> kill exist process"
ssh $remote_host "cd $remote_workdir && fuser -k $current_app"
echo ">>> rename new process"
ssh $remote_host "cd $remote_workdir && mv $new_app $current_app"
echo ">>> start new process"
ssh $remote_host "cd $remote_workdir && sh $startShell  >> $log 2>&1 &"
echo "done"

