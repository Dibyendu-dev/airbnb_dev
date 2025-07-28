-- +goose Up
-- +goose StatementBegin
create table if not exists role_permissions (
    id serial primary key,
    role_id bigint unsigned not null,
    permission_id bigint unsigned not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp, 
    foreign key (role_id) references roles(id) on delete cascade,
    foreign key (permission_id) references permissions(id) on delete cascade

);
-- +goose StatementEnd

-- insert into role_permissions(roll_id,permission_id)
-- select 1,id from permissions; -- assume role_id 1 is 'admin',admin has all permissions

-- insert into role_permissions (role_id,permission_id)
-- select 2,id from permissions where name in ('user:read');
-- +goose Down
-- +goose StatementBegin
drop table if exists role_permissions;
-- +goose StatementEnd
