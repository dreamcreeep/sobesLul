-- Предложи оптимальный индекс для blogpostsid и напиши --  уникальный индекс по запросу 

CREATE TABLE blogposttranslations (
    id SERIAL PRIMARY KEY,
    blogpostid INTEGER NOT NULL REFERENCES blog_posts (id) ON DELETE CASCADE,
    language_code VARCHAR(5) NOT NULL, -- e.g., 'en', 'es', 'fr'
    title TEXT,
    description TEXT,
    -- Составной уникальный индекс
    UNIQUE (blogpostid, language_code)
);


--  мой индекс
CREATE INDEX idxblogposttranslationspost ON blogposttranslations(blogpostid);
