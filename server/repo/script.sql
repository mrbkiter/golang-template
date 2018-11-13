drop table if exists candidate;
create table candidate (
	id varchar(64) primary key,
	first_name varchar(64),
	last_name varchar(64),
	deleted_at timestamp,
	created_at timestamp default now(),
	updated_at timestamp
);

insert into candidate(id, first_name, last_name) values ('1234', 'FN', 'LN');