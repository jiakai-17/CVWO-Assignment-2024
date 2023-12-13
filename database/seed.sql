INSERT INTO users VALUES ('testuser', 'password');

INSERT INTO threads (id, title, body, creator) VALUES ('11223344-4444-4444-4444-000000000001', 'Hello World!', 'This is my first thread', 'testuser');

INSERT INTO tags VALUES ('my-first-post');

INSERT INTO thread_tags VALUES ('11223344-4444-4444-4444-000000000001', 'my-first-post');

INSERT INTO comments (body, creator, thread_id) VALUES ('Good first post!', 'testuser', '11223344-4444-4444-4444-000000000001');

INSERT INTO threads (id, title, body, creator) VALUES ('11223344-4444-4444-4444-000000000002', 'Need help with my homework!', 'How do I get this as the answer?', 'testuser');

INSERT INTO tags VALUES ('homework');

INSERT INTO thread_tags VALUES ('11223344-4444-4444-4444-000000000002', 'homework');
