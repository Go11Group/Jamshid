-- name,surname,phone_number,email,address,password
-- create  table students(id serial primary key, name varchar,surname varchar,phoneNumber varchar,email varchar,address varchar,password varchar);



create  index  name_index on  students(name);
create  index  name_index_hash  on  students using  hash(surname);

drop  index  name_index;
drop  index  name_index_hash;
-- bunda error beradi chunki phone number yunik emas
create  unique index  name_index on students(phone_number);
create  index  name_index on students(surname,name);

explain (analyse )
select *from students where  name='Tatyana';

explain (analyse )
select  *from  students where  surname='Jones';


explain (analyse )
select  *from  students where  surname='Jones' and name='Tatyana';
