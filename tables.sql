CREATE TABLE IF NOT EXISTS Sessions (
    id INT PRIMARY KEY,
    epoch_start INT NOT NULL,
    epoch_end INT NOT NULL,
    timezone TEXT NOT NULL DEFAULT "Etc/UTC"
);

CREATE TABLE IF NOT EXISTS Labels (
    id INT PRIMARY KEY,
    label TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS LabeledSessions (
    session_id INT NOT NULL,
    label_id INT NOT NULL,
    PRIMARY KEY (session_id, label_id),
    FOREIGN KEY session_id REFERENCES sessions (id),
    ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY label_id REFERENCES labels (id),
    ON DELETE CASCADE ON UPDATE NO ACTION
);
