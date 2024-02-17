CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table users (
    id UUID primary key, 
    username varchar(256) unique not null, 
    email varchar(256) unique not null, 
    password_hash varchar(256) unique not null
);

create table problems (
    id serial primary key, 
    title varchar(256) not null,
    description varchar(256) not null, 
    is_done BOOLEAN not null, 
    created_by UUID not null, 
    Foreign Key (created_by) REFERENCES users(id)
);