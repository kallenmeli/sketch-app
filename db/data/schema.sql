create table drawings
(
    id         varchar(36) not null primary key,
    drawing    text        not null,
    created_at timestamp   not null
);