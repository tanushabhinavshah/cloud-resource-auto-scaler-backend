-- =========================
-- USERS TABLE
-- =========================
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- SYSTEM METRICS TABLE
-- =========================
CREATE TABLE system_metrics (
    id BIGSERIAL PRIMARY KEY,
    cpu_usage_percent FLOAT NOT NULL,
    memory_usage_percent FLOAT NOT NULL,
    network_usage_kbps FLOAT NOT NULL,
    is_outlier BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Fast retrieval of latest metrics
CREATE INDEX idx_system_metrics_created_at_desc
ON system_metrics (created_at DESC);


-- =========================
-- THRESHOLD CONFIG TABLE
-- Scalable per-resource design
-- =========================
CREATE TABLE threshold_config (
    id BIGSERIAL PRIMARY KEY,
    resource_name VARCHAR(50) NOT NULL,
    avg_value FLOAT NOT NULL,
    upper_limit FLOAT NOT NULL,
    lower_limit FLOAT NOT NULL,
    mode VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index for fast latest-threshold lookup
CREATE INDEX idx_threshold_resource_time
ON threshold_config (resource_name, created_at DESC);

-- =========================
-- SERVERS TABLE
-- Tracks server lifecycle
-- =========================
CREATE TABLE servers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    terminated_at TIMESTAMP
);

-- =========================
-- SCALING EVENTS TABLE
-- Audit log for scaling decisions
-- =========================
CREATE TABLE scaling_events (
    id BIGSERIAL PRIMARY KEY,
    event_type VARCHAR(10) NOT NULL,
    trigger_resource VARCHAR(50) NOT NULL,
    cpu_usage FLOAT NOT NULL,
    memory_usage FLOAT NOT NULL,
    network_usage FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- OPTIONAL SAFETY CHECKS
-- =========================
-- Ensure valid scaling direction
ALTER TABLE scaling_events
ADD CONSTRAINT chk_scaling_direction
CHECK (event_type IN ('up', 'down'));

-- Ensure valid threshold mode
ALTER TABLE threshold_config
ADD CONSTRAINT chk_threshold_mode
CHECK (mode IN ('dynamic', 'manual'));
