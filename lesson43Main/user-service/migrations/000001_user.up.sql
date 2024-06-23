create table if not exists users(
    id uuid primary key default gen_random_uuid(),
    name varchar,
    phone varchar,
    age int check ( age>0 ),
    created_at timestamp default current_timestamp,
    deleted_at bigint default 0,
    updated_at timestamp default current_timestamp
)
;
select u.name,sum(t.amount) from users  u inner join  cards   c on u.id=c.user_id inner join transactions t on  c.id=t.card_id where t.transaction_type='deposit' or t.transaction_type='credit' and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and u.id=$1 group by u.name
;
select u.name,c.number from users u inner join cards c on u.id=c.user_id;
select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='deposit';
select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='credit';