.PHONY: *

dev:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose up --build -d
	docker-compose logs -f

prod:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose -f docker-compose-prod.yml up --build -d
	docker-compose logs -f

backup-db:
	docker exec -it house-finder_house_finder_db_1 pg_dump -h house_finder_db -U root -d house_finder -f /root/house_finder.backup -F c -v --data-only
	docker cp house-finder_house_finder_db_1:/root/house_finder.backup .
