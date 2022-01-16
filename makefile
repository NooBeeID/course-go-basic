

run-db:
	docker run -d --name db-bootcamp \
	-p 9990:5432 \
	-e POSTGRES_USER=noobee \
	-e POSTGRES_PASSWORD=noobee \
	-e POSTGRES_DB=bootcamp \
	--restart always \
	postgres:alpine
