-- CREATE TRIGGERS

CREATE OR REPLACE FUNCTION update_comments_count()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'DELETE' THEN
        UPDATE threads t
        SET num_comments = (
            SELECT COUNT(*)
            FROM comments c
            WHERE c.thread_id = t.id
        )
        WHERE t.id = OLD.thread_id;
        RETURN OLD;

    ELSIF TG_OP = 'INSERT' THEN
        UPDATE threads t
        SET num_comments = (
            SELECT COUNT(*)
            FROM comments c
            WHERE c.thread_id = t.id
        )
        WHERE t.id = NEW.thread_id;
        RETURN NEW;
    END IF;
END;
$$
    LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER on_comment
    AFTER INSERT OR DELETE ON comments
    FOR EACH ROW
EXECUTE FUNCTION update_comments_count();
