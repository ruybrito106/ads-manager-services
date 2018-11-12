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

	@echo "\n\n> Creating and initializing place database"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -U postgres -p 5432 -c "CREATE DATABASE place"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -d place -U postgres -p 5432 \
		-c 'CREATE TABLE places(id text PRIMARY KEY, name text, lat double precision, lng double precision);'

	@echo "\n\n> Creating and initializing ad database"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -U postgres -p 5432 -c "CREATE DATABASE ad"
	-PGPASSWORD=postgres docker exec -it ads-manager-services_postgresdb_1 psql -h localhost -d ad -U postgres -p 5432 \
		-c 'CREATE TABLE ads(id text PRIMARY KEY, title text, description text, icon_url text, user_id text);'