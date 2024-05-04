main:
	go build -ldflags "-s -w" main.go

clean:
	rm -rf ./main

dev:
	export MY_ENV_VAR="xxxx" && fresh