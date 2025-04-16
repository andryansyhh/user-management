CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(100) NOT NULL,
  role_id INT REFERENCES roles(id),
  last_access TIMESTAMP DEFAULT NOW()
);

CREATE TABLE role_rights (
  id SERIAL PRIMARY KEY,
  role_id INT REFERENCES roles(id),
  section VARCHAR(10), -- example: "be"
  route TEXT,
  r_create BOOLEAN DEFAULT FALSE,
  r_read BOOLEAN DEFAULT FALSE,
  r_update BOOLEAN DEFAULT FALSE,
  r_delete BOOLEAN DEFAULT FALSE
);

-- INDEXES
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_role_rights_role_id ON role_rights(role_id);
CREATE INDEX idx_role_rights_route_section ON role_rights(route, section);

-- Seed Admin Role/User/Access
INSERT INTO roles (id, name) VALUES (1, 'Admin');

INSERT INTO users (id, name, email, password, role_id)
VALUES (1, 'Administrator', 'admin@gmail.com', 'adminadmin', 1);

INSERT INTO role_rights (role_id, section, route, r_create, r_read, r_update, r_delete)
VALUES
  (1, 'be', '/users/user', TRUE, TRUE, TRUE, TRUE),
  (1, 'be', '/auth/login', TRUE, FALSE, FALSE, FALSE);
