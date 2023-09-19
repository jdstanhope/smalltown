
create table if not exists photos
(
    id         bigserial primary key,
    created_at timestamp(0) with time zone not null default now(),
    name       text                        not null,
    version    integer                     not null default 1
);