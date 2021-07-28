all: build-docker cron-docker

build-docker:
	docker build -t irm .

run-docker:
	docker run -it irm

cron-docker:
	docker run -it -e WORKERS=20 -v /root/irm/cmd/analyze/results:/app/cmd/analyze/results/ -v /root/SSH_KEY/id_ed25519:/app/id_ed25519 irm
