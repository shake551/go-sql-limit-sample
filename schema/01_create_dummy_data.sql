INSERT INTO articles () VALUES ();
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;
INSERT INTO articles (id) SELECT 0 FROM articles;

UPDATE articles SET
    title = SUBSTRING(MD5(RAND()), 1, 10),
    content = SUBSTRING(MD5(RAND()), 1, 50),
    created_at = 1674090479 + id,
    updated_at = 1674090479 + id;
