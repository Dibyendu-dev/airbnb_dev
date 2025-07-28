-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- seeders 
-- insert into permissions (name,description,resourse,action) values
-- ('user:read', 'permission to read user data', 'user','read'),
-- ('user:write', 'permission to write user data', 'user','write'),
-- ('user:delete', 'permission to delete user data', 'user','delete'),
-- ('role:read', 'permission to read role data', 'role','read'),
-- ('role:write', 'permission to write role data', 'role','write'),
-- ('role:delete', 'permission to delete role data', 'role','delete'),
-- ('permission:read', 'permission to read permissions', 'permission','read'),
-- ('permission:write', 'permission to write permissions', 'permission','write'),
-- ('permission:delete', 'permission to delete permissions', 'permission','delete');


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
