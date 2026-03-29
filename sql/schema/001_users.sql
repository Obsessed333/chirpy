-- +goose Up
create table users(
    id uuid primary key not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    email text unique not null,
    hashed_password text not null default 'unset',
    is_chirpy_red bool not null default false
);

-- +goose Down
drop table users;

