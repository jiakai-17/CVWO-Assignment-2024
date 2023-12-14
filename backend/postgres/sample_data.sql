INSERT INTO users VALUES ('testuser', 'password');

INSERT INTO threads (id, title, body, creator) VALUES ('11223344-4444-4444-4444-000000000001', 'Hello World!', 'This is my first thread', 'testuser');

INSERT INTO tags VALUES ('my-first-post');

INSERT INTO thread_tags VALUES ('11223344-4444-4444-4444-000000000001', 'my-first-post');

INSERT INTO comments (body, creator, thread_id) VALUES ('Good first post!', 'testuser', '11223344-4444-4444-4444-000000000001');

INSERT INTO threads (id, title, body, creator) VALUES ('11223344-4444-4444-4444-000000000002', 'Need help with my homework!', 'How do I get this as the answer?', 'testuser');

INSERT INTO tags VALUES ('homework');

INSERT INTO thread_tags VALUES ('11223344-4444-4444-4444-000000000002', 'homework');

SELECT * FROM threads
SELECT * FROM comments

INSERT INTO threads (title, body, creator) VALUES ('How can one ', 'This is my second thread', 'testuser');

SELECT * FROM threads

INSERT INTO comments (body, creator, thread_id) VALUES ('Good first post!', 'testuser', 'd51becba-a6bc-4139-9d33-d61f55b018cb');

SELECT * FROM threads
SELECT * FROM comments


INSERT INTO comments VALUES ('c41c7a90-4188-49bf-924f-0a6331f43bc5', 'testbody', 'testuser', 'b0764ee4-95c2-461f-b7c3-18c49fecb9d9');

SELECT * FROM threads

DELETE FROM comments WHERE id = '5eec3239-3d33-4f7e-811e-4b055cb59c8e'

DELETE FROM users WHERE username = 'testuser'
