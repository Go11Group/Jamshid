create  table  if not exists  users(
    id uuid primary key default gen_random_uuid(),
    first_name varchar,
    last_name varchar,
    age int,
    email varchar,
    phone varchar,

    created_at time default localtimestamp,
    updated_at time  default localtimestamp,
    deleted_at time default localtimestamp


);

drop  table  users;
drop  table  problems;
drop  table solved_problems;



create  table  if not exists problems (
    id uuid primary key default gen_random_uuid(),
    description varchar,
    type varchar,
    created_at time default localtimestamp,
    updated_at time  default localtimestamp ,
    deleted_at time default localtimestamp

                                      );

create  table  if not exists  solved_problems(
    id uuid primary key default gen_random_uuid(),
    solution varchar,
    problem_solution uuid  not null references problems(id) on  delete  cascade ,
    user_solution uuid not null references   users(id) on delete cascade ,
    created_at time default localtimestamp,
    updated_at time  default localtimestamp,
    deleted_at time default localtimestamp
                                            );


-- insert into  users(first_name, last_name, age, email, phone) values ('Jamshid','Hatamov',12,'hatamov','12345')
-- delete  from users where  id='cb40fadb-3a7c-47f5-b353-155171bdb25a';
-- insert into  solved_problems(solution, problem_solution, user_solution) values ('mukalads','644102d3-61f0-4041-97b0-18c09472cd05','8a1bc6f7-6de8-4ed6-8462-c2ca23fc4f4e')