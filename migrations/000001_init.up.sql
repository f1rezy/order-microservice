DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists public.orders
(
    id       uuid default uuid_generate_v4() not null
        constraint orders_pk
            primary key,
    item     text                    not null,
    quantity integer                 not null
);