CREATE TABLE campaigns(
    id SERIAL PRIMARY KEY,
    name text,
    start_ts integer,
    end_ts integer,
    visits_goal integer,
    status text,
    places text ARRAY,
    ads text ARRAY
);