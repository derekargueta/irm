all: build-docker run-docker

build-docker:
	docker build -t irm-test .

run-docker:
	docker run -it irm-test


# -v localDIR:dockerDIR
# -v /{$PATH}/certs:/app/ -> directory for user generated certificate
 