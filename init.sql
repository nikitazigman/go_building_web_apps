CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    page_guid VARCHAR(256) UNIQUE NOT NULL DEFAULT '',
    page_title VARCHAR(256) DEFAULT NULL, 
    page_content TEXT, 
    page_date TIMESTAMP NOT NULL DEFAULT NOW()
);
INSERT INTO pages (page_guid, page_title, page_content, page_date)
VALUES ('hello-world', 'hello, world', 'I am so glad you dound this page', CURRENT_TIMESTAMP);
