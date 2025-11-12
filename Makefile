






migrateup:
	migrate -path internal/db/migration  -database "postgresql://root:root@localhost:5432/bank-database?sslmode=disable" -verbose up

createdb: 
		sudo docker exec -it bank-container createdb --username=root bank-database 
postgres:
	  sudo docker run --name bank-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
start:
		sudo docker start bank-container 
stop:
		sudo docker stop bank-container 
dropdb:
	sudo docker exec -it bank-container dropdb --username=root bank-database 
.PHONY: createdb dropdb postgres migrateup
