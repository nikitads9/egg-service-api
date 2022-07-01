-- +goose Up
create table meals 
(
    id bigserial primary key,
    user_id bigserial not null,
    meal_date text not null,
    weight real not null,
    proteins real,
    fat real,
    carbo real,
    nutrition real
);

-- +goose Down
drop table meals;
