all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
<<<<<<< HEAD
	docker run -it -e WORKERS=20 -v /Users/Macintosh_HD/Documents/irm/cmd/analyze/results:/app/cmd/analyze/results/ irm
=======
	docker run -it --mount type=bind,source="$(pwd)"/cmd/analyze/results/results.csv, target=/root/irm-data/ -e WORKERS=20 irm
>>>>>>> 587a0407039d13c2e07efc5107b1263dd7e72ae2
