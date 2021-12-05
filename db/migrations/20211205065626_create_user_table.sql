-- migrate:up
create table if not exists users (
    id serial primary key,
    phone varchar(30)
)
-- migrate:down
