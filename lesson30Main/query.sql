create table  if not exists  customers

(

    id serial primary key ,
    username varchar(50) not null unique ,
    email varchar(100) not null ,
    password varchar(100) not null
                                     );
select *from customers;
create table  products(
    id serial primary key ,
    name varchar(100) not null ,
    description text,price numeric(10,2) not null ,
    stock_quantity int not null
                      );
-- insert into  customers(username, email, password) values ('Azamat','Azamatov','123456')
insert into  products(name, description, price, stock_quantity) values ('telefon','alo',23.5,45)