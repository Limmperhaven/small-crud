CREATE TABLE IF NOT EXISTS records
(
    uuid         varchar(64)  not null unique primary key,
    first_name   varchar(255) not null,
    last_name    varchar(255) not null,
    mobile_phone varchar(255) not null,
    home_phone   varchar(255)
);