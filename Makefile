name := imanhodjaev/cameo
tag ?= $(shell git log -1 --pretty=format:%h)

image:
	@docker build . \
	-t $(name):latest \
	-t $(name):$(tag)

test:
	CGO_ENABLED=0 go test ./...

# Docker targets
compile:
	GOOS=linux CGO_ENABLED=0 go build -o /cameo/cameo
