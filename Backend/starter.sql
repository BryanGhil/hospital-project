create table roles (
    id bigserial primary key,
    role varchar not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create table users (
    id bigserial primary key,
    email varchar not null unique,
    password varchar not null,
    role_id bigint not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    foreign key (role_id) references roles(id)
);

insert into roles (role, created_at, updated_at) 
values ('admin', NOW(), NOW()),('user', NOW(), NOW());
