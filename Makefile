all: build-docker run-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm
