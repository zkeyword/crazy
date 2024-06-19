run:export MY_ENV_VAR:=prod
run_dev:export MY_ENV_VAR:=dev

build:
	go build -ldflags "-s -w" main.go

clean:
	rm -rf ./main

dev:
	export MY_ENV_VAR="dev" && fresh

run:
	nohup ./main crazy >> log 2>&1 &

run_dev:
	nohup ./main crazy >> log 2>&1 &

ps:
	ps aux | grep crazy