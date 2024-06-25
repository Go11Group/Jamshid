create table  if not exists cards(
                                     id uuid primary key  default gen_random_uuid(),
    number int default 0,
    user_id uuid  not null references users(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0

    )