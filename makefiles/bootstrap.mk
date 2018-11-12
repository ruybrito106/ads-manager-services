bootstrap:
	@echo "\n\n> Creating and initializing campaign database"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -U postgres -p 5432 -c "CREATE DATABASE campaign"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -d campaign -U postgres -p 5432 \
		-c 'CREATE TABLE campaigns(id SERIAL PRIMARY KEY, name text, start_ts integer, end_ts integer, visits_goal integer, status text, places text ARRAY, ads text ARRAY);'

	@echo "\n\n> Creating and initializing balance database"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -U postgres -p 5432 -c "CREATE DATABASE balance"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -d balance -U postgres -p 5432 \
		-c 'CREATE TABLE balances(user_id TEXT PRIMARY KEY, amount integer);'
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -d balance -U postgres -p 5432 \
		-c "INSERT INTO balances(user_id, amount) VALUES('superuser', 999999);"