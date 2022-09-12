CREATE TABLE IF NOT EXISTS user (
    id  varchar(36) PRIMARY KEY,
    email  varchar(100) NOT NULL UNIQUE,
    first_name varchar(200) NOT NULL,
    last_name varchar(100) NOT NULL,
    password TEXT NOT NULL,
    is_active BOOLEAN DEFAULT 0,
    is_staff BOOLEAN DEFAULT 0,
    is_superuser BOOLEAN DEFAULT 0,
    created_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6)
);