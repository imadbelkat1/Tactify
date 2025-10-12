CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    team_id INT NOT NULL,
    team_name VARCHAR(255) NOT NULL,
    country VARCHAR(100),
    primary_color VARCHAR(7),
    secondary_color VARCHAR(7),
    UNIQUE(team_id, country)
)