-- Create todos table
CREATE TABLE todos (
    id VARCHAR(255) NOT NULL,
    todo VARCHAR(255) NOT NULL,
    done BOOLEAN DEFAULT false
);