INSERT INTO accesses (id, name) VALUES
(1, 'create'),
(2, 'read'),
(3, 'update'),
(4, 'delete');

INSERT INTO roles (id, name) VALUES
(1, 'admin'),
(2, 'member'),
(3, 'manager'),
(4, 'guest');

INSERT INTO role_accesses (id, role_id, access_id) VALUES
(1, 1, 1), -- admin: create
(2, 1, 2), -- admin: read
(3, 1, 3), -- admin: update
(4, 1, 4), -- admin: delete
(5, 2, 1), -- member: create
(6, 2, 2), -- member: read
(7, 2, 3), -- member: update
(8, 2, 4), -- member: delete
(9, 3, 1), -- manager: create
(10, 3, 2), -- manager: read
(11, 3, 3), -- manager: update
(12, 3, 4), -- manager: delete
(13, 4, 2); -- guest: read

INSERT INTO users (id, name, email) VALUES
(1, 'admin', 'admin@admin.admin');

INSERT INTO user_roles (user_id, role_id) VALUES (1, 1);
INSERT INTO user_roles (user_id, role_id) VALUES (1, 2);