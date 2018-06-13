serve:
	go run main.go config.go util.go middleware.go uploader.go slackbot.go slackrtm.go

build:
	CGO_ENABLE=0 go build -o cdrobot main.go config.go util.go middleware.go uploader.go slackbot.go slackrtm.go

lint:
	golint -set_exit_status ./...

test:
	go test -race -timeout 2m -coverprofile=coverprofile.out ./...

