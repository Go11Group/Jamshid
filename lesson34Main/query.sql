-- drop  table  students;
-- create  table if not exists students (
--     id uuid   primary key  default gen_random_uuid(),
--     first_name varchar,last_name varchar,
--     age int ,
--     email  varchar not null,
--     phone_number varchar not null,
--     created_at timestamp,
--     updated_at timestamp ,
--     deleted_at timestamp
--                                      );
--
--

--

-- CREATE TABLE IF NOT EXISTS students_groups_reference
-- (
--                                                        id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
--                                                        student_id UUID ,
--                                                        group_id UUID,
--                                                        FOREIGN KEY (student_id) REFERENCES students(id) on delete cascade,
--                                                        FOREIGN KEY (group_id) REFERENCES groups(id)  on delete  cascade ,
--                                                        created_at timestamp,
--                                                        updated_at timestamp ,
--                                                        deleted_at timestamp
-- );


-- drop table faculties;
-- create  table  faculties(
--     id uuid primary key default gen_random_uuid(),
--     faculty_name varchar,
--     group_count int ,
--     decant varchar,
--     faculty_email varchar,
--     faculty_phone varchar not null  ,
--     created_at timestamp,
--     updated_at timestamp ,
--     deleted_at timestamp
-- );
-- create  table  if not exists  groups_faculties_references(
--     id uuid primary key  default gen_random_uuid(),
--     groups_id  uuid  ,
--     faculty_id uuid ,
--     foreign key (groups_id) references groups(id) on delete  cascade ,
--     foreign key (faculty_id) references  faculties(id) on delete  cascade
-- );


create  table faculties_universities_reference(
    id uuid primary key  default gen_random_uuid(),
    faculty_id uuid ,
    universities uuid,
    foreign key (faculty_id) references faculties(id) on delete  cascade ,
    foreign key (universities) references  universities(id) on delete cascade

);
--
-- create  table  if not exists  universities
-- (
--     id uuid primary key ,university_name varchar,
--     rector varchar,university_email varchar,
--     university_phone varchar not null  check ( '+9989' ),
--     created_at timestamp,
--     updated_at timestamp ,
--     deleted_at timestamp
                                          );