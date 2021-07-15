all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
	docker run -it -e WORKERS=20 -v /Users/Macintosh_HD/Documents/irm/cmd/analyze/results:/app/cmd/analyze/results/ irm
