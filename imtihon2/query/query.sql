

create  table  if not exists companies (
    id uuid primary key default gen_random_uuid(),
    name varchar(250),
    location varchar (250),
    workers int ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0
);
create  table if not exists interviews (
    id uuid primary key default gen_random_uuid(),
    user_id uuid  not null references users(id),
    vacancy_id  uuid not null references  vacancies(id),
    recruiter_id uuid  not null references  recruiters(id),
    interview_date time,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0

);
create  type    Gender as  enum('female','male');
create  table if not exists recruiters(
   id uuid primary key default gen_random_uuid(),
    name varchar(250),
    email varchar(250),
    phone_number varchar(250),
    birthday varchar,
    gender Gender,
    company_id uuid not null references companies(id),
   created_at timestamp default current_timestamp,
   updated_at timestamp default current_timestamp,
   deleted_at bigint default 0
);

create  table  if not exists users(
     id uuid primary key default gen_random_uuid(),
     name varchar(250),
     email varchar(250),
     phone_number varchar(250),
     birthday vacancies,
     gender Gender,
     created_at timestamp default current_timestamp,
     updated_at timestamp default current_timestamp,
     deleted_at bigint default 0

);
create  table  if not exists resumes(
     id uuid primary key default gen_random_uuid(),
    position varchar(250),
    experience int ,
    description text,
    user_id uuid not null references users(id),
     created_at timestamp default current_timestamp,
     updated_at timestamp default current_timestamp,
     deleted_at bigint default 0
);
create  table  if not exists vacancies(
    id uuid primary key default gen_random_uuid(),
    name varchar(200),
    min_experience int,
    company_id uuid  not null references companies(id),
    description text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0

)


