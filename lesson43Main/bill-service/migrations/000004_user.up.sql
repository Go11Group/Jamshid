create table  if not exists transactions(
                                            id uuid primary key  default gen_random_uuid(),
                                            card_id uuid not null  references cards(id),
                                            amount bigint,
                                            terminal_id uuid not null  references terminals(id) default null,
                                            transaction_type TransactionType,
                                            created_at timestamp default current_timestamp,
                                            updated_at timestamp default current_timestamp,
                                            deleted_at bigint default 0


);
select u.id,u.name, u.phone, c.id,c.number,t.amount from cards c inner join   transactions t on c.id=t.card_id inner join  users u on c.user_id=u.id where u.id=$1 and u.deleted_at=0

;
select s.id,s.name from stations  s inner join  terminals t on s.id=t.station_id where s.id=$1 and s.deleted_at=0 and t.deleted_at=0;
select  sum(amount)  as totaly_summ from  transactions where transaction_type='deposit' or transaction_type='credit'
;
-- select *from cards  c inner join  transactions t on c.id = t.card_id  where c.id=$1 and c.deleted_at=0 and t.deleted_at=0;
-- select  *from  transactions  inner join  terminals  on transactions.terminal_id = terminals.id where terminals.id=$1 and terminals.deleted_at=0 and transactions.deleted_at=0;
-- select  c.id,c.number,t.id,t.transaction_type,s.name from  cards c inner join   transactions t on c.id=t.card_id inner join terminals  on t.terminal_id = terminals.id inner join stations s on terminals.station_id = s.id where c.id=$1 and c.deleted_at=0 and t.deleted_at=0 and terminals.deleted_at=0 and s.deleted_at=0
;
