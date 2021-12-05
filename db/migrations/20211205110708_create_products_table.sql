-- migrate:up
create table if not exists products (
    id serial primary key,
    name varchar(100)
)
-- migrate:down
