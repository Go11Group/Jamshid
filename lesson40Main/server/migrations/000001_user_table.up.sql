create table  users (
    id uuid primary key  default gen_random_uuid(),
    name varchar,
    age int ,
    email varchar
    )