-- +goose Up
create table devices
(
    serial_number bigserial primary key
);

insert into devices (serial_number) values (4221);

create table meals 
(
    id bigserial primary key,
    user_id bigserial references devices(serial_number) not null,
    meal_date text not null,
    weight real not null,
    proteins real not null,
    fat real not null,
    carbo real not null,
    nutrition real not null
);

-- +goose Down
drop table devices;
drop table meals;
