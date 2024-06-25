
create table  if not exists terminals(
                                         id uuid primary key  default gen_random_uuid(),
                                         station_id uuid not null  references stations(id),
                                         created_at timestamp default current_timestamp,
                                         updated_at timestamp default current_timestamp,
                                         deleted_at bigint default 0
)