DOCKER_IMAGE ?= "colonelmo/talkative"

.PHONY: help clean build docker push

help:
	@echo "make <target> against one of the available <target>s:"
	@echo "	clean:  to delete application executable"
	@echo "	build:  to build the executable"
	@echo "	docker:	to create a docker image with tag \$DOCKER_IMAGE"
	@echo "	push:   to create a docker image and push it upstream"

clean:
	rm talk 1>/dev/null 2>/dev/null || exit 0

build: talk.go clean
	go get ./...
	go build -o talk

docker: build
	docker build -t $(DOCKER_IMAGE) .

push: docker
	docker push $(DOCKER_IMAGE)
