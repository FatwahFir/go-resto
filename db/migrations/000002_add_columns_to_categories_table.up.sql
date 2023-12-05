ALTER TABLE categories
    ADD COLUMN created_at timestamp not null default current_timestamp,
    ADD COLUMN updated_at timestamp not null default current_timestamp on update current_timestamp;