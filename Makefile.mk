MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := build
PHONY: all build test push push_latest clean clean_all

M = $(shell printf "\033[34;1m▶\033[0m")

NAME	   := "quay.io/ahelal/cibully"
VERSION    := $(shell cat .VERSION)
BUILD_ARGS :=  --build-arg VERSION=${VERSION}


docker: dbuild dpush dpush_latest

dbuild:
	$(info $(M) Building ${NAME}:${VERSION}…)
	@#docker build ${BUILD_ARGS} --no-cache=True -t ${NAME}:${VERSION} -f Dockerfile .
	docker build ${BUILD_ARGS} -t ${NAME}:${VERSION} -f Dockerfile .

dtests:
	@if ! docker images $(NAME) | awk '{print $$2 }' | grep -q -F $(VERSION); then echo "$(NAME) version $(VERSION) is not yet built. Please run 'make build'"; false; fi
	@docker tag $(NAME):$(VERSION) $(NAME):dev
	@./test/run_tests.sh tests

dtests-debug:
	@if ! docker images $(NAME) | awk '{print $$2 }' | grep -q -F $(VERSION); then echo "$(NAME) version $(VERSION) is not yet built. Please run 'make build'"; false; fi
	@docker tag $(NAME):$(VERSION) $(NAME):dev
	@./test/run_tests.sh tests-debug

dpush:
	@echo ${VERSION}
	@if ! docker images $(NAME) | awk '{print $$2 }' | grep -q -F $(VERSION); then echo "$(NAME) version $(VERSION) is not yet built. Please run 'make build'"; false; fi
	$(info $(M) Pushing ${NAME}:${VERSION}…)
	@docker push "${NAME}:${VERSION}"

dpush_latest:
	$(info $(M) Linking latest and pushing ${NAME}:${VERSION}…)
	docker tag $(NAME):$(VERSION) $(NAME):latest
	docker push "${NAME}:latest"

dclean:
	@if docker images $(NAME) | awk '{print $$2 }' | grep -q -F $(VERSION); then echo "*** Cleaning ${NAME}:${VERSION} ***"; docker rmi "${NAME}:${VERSION}"; else echo "*** No image ${NAME}:${VERSION}  ***"; fi


dclean_all: clean
	$(info $(M) Cleaning all…)
	@#@if docker images $(NAME) | awk '{print $$2 }'; then echo "*** Cleaning ${NAME} ***"; docker rmi "${NAME}"; fi
	docker system prune -f
