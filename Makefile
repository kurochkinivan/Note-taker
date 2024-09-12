rm: 
	docker-compose stop \
	&& docker-compose rm \
	&& docker image rm nt-backend:local \
	&& sudo rm -rf data/ 

up: 
	docker build -t nt-backend:local . \
	&& docker-compose -f docker-compose.yaml up --detach --force-recreate

up-db:
	docker-compose up -d postgresql

rm-db:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf data/ 
