create table message
(
    id   integer generated always as identity,
    data text default ''::text not null
);

alter table message
    owner to postgres;

