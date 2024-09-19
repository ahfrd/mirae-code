-- +goose Up
-- +goose StatementBegin
create table mirae.stock
(
    id        int auto_increment
        primary key,
    name      varchar(255) not null,
    code      varchar(64)  not null,
    price     float        null,
    frequency float        null,
    volume    varchar(64)  null,
    constraint stock_fk
        unique (code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    drop TABLE mirae.stock
-- +goose StatementEnd
