NAMESPACE = rightlag
REPOSITORY = slate
NAME = slate
VERSION ?= latest

.PHONY: build shell

build:
	docker build -t $(NAMESPACE)/$(REPOSITORY):$(VERSION) .

shell:
	docker run -it --rm --name $(NAME) $(NAMESPACE)/$(REPOSITORY):$(VERSION)

default: build
