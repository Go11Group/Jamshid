create table  if not exists stations(
    id uuid primary key  default gen_random_uuid(),
    name varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0
)