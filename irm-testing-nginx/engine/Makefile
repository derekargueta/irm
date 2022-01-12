all: build-docker run-docker

build-docker:
	docker build -t engine .

run-docker:
	docker run -it --rm -p 8081:80 engine


# -v localDIR:dockerDIR
# -v /{$PATH}/certs:/app/ -> directory for user generated certificate
 
