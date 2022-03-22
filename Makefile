all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
	docker run -it -v /root/.ssh:/home/.ssh/ irm /bin/bash
# -v localDIR:dockerDIR
# -v 
 