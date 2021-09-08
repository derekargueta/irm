all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
	docker run -it -v /home/henry/.ssh:/home/.ssh/ irm
# -v localDIR:dockerDIR
 