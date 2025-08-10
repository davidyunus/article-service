CREATE TABLE authors (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE articles (
  id TEXT PRIMARY KEY,
  author_id TEXT NOT NULL REFERENCES authors(id),
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_articles_created_at ON articles(created_at DESC);
CREATE INDEX idx_authors_name ON authors(name);