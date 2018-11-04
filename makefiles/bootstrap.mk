bootstrap:
	@echo "\n\n> Creating and initializing campaign database"
	-PGPASSWORD=postgres docker exec -it adsmanagerservices_postgresdb_1 psql -h localhost -U postgres -p 5432 -c "CREATE DATABASE campaign"
	-PGPASSWORD=postgres docker exec -it adsmanagerservices_postgresdb_1 psql -h localhost -d campaign -U postgres -p 5432 \
		-c 'CREATE TABLE campaigns(id SERIAL PRIMARY KEY, name text, start_ts integer, end_ts integer, visits_goal integer, status text, places text ARRAY, ads text ARRAY);'