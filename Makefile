#!make
include configs/dev.env

SONAR_PATH = ~/sonar/bin

migration-%:
	@migrate create -ext sql -dir database/migrations $*

migrate-dev:
	migrate -path ./database/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up

sonar:
	echo $(SONAR_PATH)
	 ${SONAR_PATH}/sonar-scanner \
	  -Dsonar.host.url=http://localhost:9000 \
	  -Dsonar.token=${SONAR_TOKEN}