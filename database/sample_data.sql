INSERT INTO users VALUES ('testuser', 'password');

INSERT INTO threads VALUES ('b0764ee4-95c2-461f-b7c3-18c49fecb9d9', 'testtitle', 'testbody', 'testuser');

INSERT INTO comments VALUES ('c41c7a90-4188-49bf-924f-0a6331f43bc5', 'testbody', 'testuser', 'b0764ee4-95c2-461f-b7c3-18c49fecb9d9');

SELECT * FROM threads

DELETE FROM comments WHERE id = 'c41c7a90-4188-49bf-924f-0a6331f43bc5'

DELETE FROM users WHERE username = 'testuser'
