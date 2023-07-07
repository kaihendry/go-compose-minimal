up:
	docker-compose up --build

reallyclean:
	docker-compose down -v --rmi all --remove-orphans
