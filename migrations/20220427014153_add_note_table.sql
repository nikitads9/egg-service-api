-- +goose Up
create table notes 
(
    id bigserial primary key,
    title text not null,
    content text not null
);

-- +goose Down
drop table notes;
