-- migrate:up
create table if not exists users (
    id serial primary key,
    name varchar(30)
)

-- migrate:down

