create table  Book(id serial primary key ,name varchar not null ,age int ,author_name varchar not null );
insert into Book (id, name, author_name, age) values (1, 'Gecko, bent-toed', 'Dina', 18);
insert into Book (id, name, author_name, age) values (2, 'Goliath heron', 'Curtice', 2);
insert into Book (id, name, author_name, age) values (3, 'Quoll, eastern', 'Harlene', 61);
insert into Book (id, name, author_name, age) values (4, 'Secretary bird', 'Sheridan', 35);
insert into Book (id, name, author_name, age) values (5, 'Lourie, grey', 'Cole', 24);
insert into Book (id, name, author_name, age) values (6, 'Bleu, red-cheeked cordon', 'Arlan', 71);
insert into Book (id, name, author_name, age) values (7, 'Buffalo, asian water', 'Hobey', 82);
insert into Book (id, name, author_name, age) values (8, 'Deer, white-tailed', 'Horatia', 6);
insert into Book (id, name, author_name, age) values (9, 'Water legaan', 'Danielle', 49);
insert into Book (id, name, author_name, age) values (10, 'Black curlew', 'Hodge', 25);
insert into Book (id, name, author_name, age) values (11, 'Water legaan', 'Lorrin', 100);
insert into Book (id, name, author_name, age) values (12, 'South American sea lion', 'Ervin', 19);
insert into Book (id, name, author_name, age) values (13, 'Yellow-necked spurfowl', 'Vale', 46);
insert into Book (id, name, author_name, age) values (14, 'Jackal, asiatic', 'Caroline', 96);
insert into Book (id, name, author_name, age) values (15, 'Stork, marabou', 'Cacilie', 9);
insert into Book (id, name, author_name, age) values (16, 'Ibis, sacred', 'Andie', 4);
insert into Book (id, name, author_name, age) values (17, 'Brown antechinus', 'Julianna', 91);
insert into Book (id, name, author_name, age) values (18, 'Goose, snow', 'Heidie', 100);
insert into Book (id, name, author_name, age) values (19, 'Brush-tailed bettong', 'Linea', 82);
insert into Book (id, name, author_name, age) values (20, 'African wild dog', 'Merrick', 52);
select  * from  Book;

create table Author  (id serial primary key ,name varchar not null,foreign key (id) references  Book(id));
insert into Author (id, name) values (1, 'Craggy');
insert into Author (id, name) values (2, 'Tarrance');
insert into Author (id, name) values (3, 'Correy');
insert into Author (id, name) values (4, 'Eadith');
insert into Author (id, name) values (5, 'Glennie');
insert into Author (id, name) values (6, 'Derrek');
insert into Author (id, name) values (7, 'Emalee');
insert into Author (id, name) values (8, 'Gallagher');
insert into Author (id, name) values (9, 'Phil');
insert into Author (id, name) values (10, 'Stephine');
insert into Author (id, name) values (11, 'Obediah');
insert into Author (id, name) values (12, 'Con');
insert into Author (id, name) values (13, 'Clarance');
insert into Author (id, name) values (14, 'Janine');
insert into Author (id, name) values (15, 'Cosimo');
insert into Author (id, name) values (16, 'Obadiah');
insert into Author (id, name) values (17, 'Ase');
insert into Author (id, name) values (18, 'Shurwood');
insert into Author (id, name) values (19, 'Jania');
insert into Author (id, name) values (20, 'Emmye');


select  * from  Author;



