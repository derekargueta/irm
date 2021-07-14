all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
	docker run -it --mount type=bind,source="$(pwd)"/cmd/analyze/results/results.csv, target=/root/irm-data/ -e WORKERS=20 irm
