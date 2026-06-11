SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS post_tags;
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS users;

SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(150) NOT NULL,
  active TINYINT(1) NOT NULL DEFAULT 1,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uniq_users_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE user_profiles (
  user_id INT NOT NULL,
  bio TEXT,
  avatar_url VARCHAR(255) DEFAULT NULL,
  website VARCHAR(255) DEFAULT NULL,
  birthday DATE DEFAULT NULL,
  updated_at TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (user_id),
  CONSTRAINT fk_user_profiles_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE posts (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  title VARCHAR(200) NOT NULL,
  body TEXT NOT NULL,
  published TINYINT(1) NOT NULL DEFAULT 0,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY idx_posts_user_id (user_id),
  CONSTRAINT fk_posts_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE tags (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uniq_tags_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE post_tags (
  post_id INT NOT NULL,
  tag_id INT NOT NULL,
  PRIMARY KEY (post_id, tag_id),
  KEY idx_post_tags_tag_id (tag_id),
  CONSTRAINT fk_post_tags_post_id
    FOREIGN KEY (post_id) REFERENCES posts (id)
    ON DELETE CASCADE,
  CONSTRAINT fk_post_tags_tag_id
    FOREIGN KEY (tag_id) REFERENCES tags (id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE comments (
  id INT NOT NULL AUTO_INCREMENT,
  post_id INT NOT NULL,
  user_id INT NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY idx_comments_post_id (post_id),
  KEY idx_comments_user_id (user_id),
  CONSTRAINT fk_comments_post_id
    FOREIGN KEY (post_id) REFERENCES posts (id)
    ON DELETE CASCADE,
  CONSTRAINT fk_comments_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE events (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  type VARCHAR(100) NOT NULL,
  payload JSON DEFAULT NULL,
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY idx_events_user_id_created_at (user_id, created_at),
  CONSTRAINT fk_events_user_id
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (id, name, email, active, created_at) VALUES
(1, 'John Doe', 'john@example.com', 1, '2026-05-07 11:15:21'),
(2, 'Jane Smith', 'jane@example.com', 1, '2026-05-07 11:15:21');

INSERT INTO user_profiles (user_id, bio, avatar_url, website, birthday, updated_at) VALUES
(1, 'Backend developer and database enthusiast.', 'https://example.com/avatars/john.png', 'https://john.example.com', '1990-04-12', '2026-05-07 11:20:00'),
(2, NULL, NULL, 'https://jane.example.com', NULL, NULL);

INSERT INTO posts (id, user_id, title, body, published, created_at) VALUES
(1, 1, 'Hello MySQL', 'First demo post.', 1, '2026-05-07 11:15:21'),
(2, 2, 'Another Post', 'Second demo post.', 0, '2026-05-07 11:15:21');

INSERT INTO tags (id, name, created_at) VALUES
(1, 'mysql', '2026-05-07 11:15:21'),
(2, 'golang', '2026-05-07 11:15:21'),
(3, 'mapper', '2026-05-07 11:15:21');

INSERT INTO post_tags (post_id, tag_id) VALUES
(1, 1),
(1, 2),
(2, 3);

INSERT INTO comments (id, post_id, user_id, body, created_at) VALUES
(1, 1, 2, 'Nice post!', '2026-05-07 11:25:00'),
(2, 2, 1, 'Thanks for sharing.', '2026-05-07 11:30:00');

INSERT INTO events (id, user_id, type, payload, created_at) VALUES
(1, 1, 'user.login', JSON_OBJECT('ip', '127.0.0.1'), '2026-05-07 11:15:21'),
(2, 2, 'post.created', JSON_OBJECT('post_id', 2), '2026-05-07 11:15:21');
