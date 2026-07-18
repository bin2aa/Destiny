-- ============================================================================
-- SCHEMA APP ĐOÁN MỆNH - v2 (đã tái cấu trúc)
-- ============================================================================

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ============================================================================
-- ENUM TYPES
-- ============================================================================

CREATE TYPE subscription_status AS ENUM ('active', 'expired', 'cancelled', 'pending');
CREATE TYPE payment_status AS ENUM ('pending', 'success', 'failed', 'refunded');
CREATE TYPE analysis_type AS ENUM ('bazi', 'numerology', 'mbti', 'big_five', 'element_scores');
CREATE TYPE report_status AS ENUM ('pending', 'success', 'failed');

-- ============================================================================
-- 1. USERS & AUTH
-- ============================================================================

CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(150) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    avatar          VARCHAR(500),
    language        VARCHAR(10) DEFAULT 'vi',
    timezone        VARCHAR(50) DEFAULT 'Asia/Ho_Chi_Minh',
    country         VARCHAR(2),
    premium_until   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ============================================================================
-- 2. BILLING
-- ============================================================================

CREATE TABLE plans (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code            VARCHAR(50) NOT NULL UNIQUE,
    name            VARCHAR(100) NOT NULL,
    price           NUMERIC(12,2) NOT NULL,
    currency        VARCHAR(3) NOT NULL DEFAULT 'VND',
    duration_days   INT NOT NULL,
    features        JSONB,
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE subscriptions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    plan_id         UUID NOT NULL REFERENCES plans(id),
    start_at        TIMESTAMPTZ NOT NULL,
    end_at          TIMESTAMPTZ NOT NULL,
    status          subscription_status NOT NULL DEFAULT 'pending',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_subscriptions_user ON subscriptions(user_id);

CREATE TABLE payments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    subscription_id UUID NOT NULL REFERENCES subscriptions(id),
    provider        VARCHAR(50) NOT NULL,
    amount          NUMERIC(12,2) NOT NULL,
    currency        VARCHAR(3) NOT NULL DEFAULT 'VND',
    status          payment_status NOT NULL DEFAULT 'pending',
    transaction_id  VARCHAR(255),
    paid_at         TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_payments_subscription ON payments(subscription_id);
CREATE INDEX idx_payments_transaction ON payments(transaction_id);

-- ============================================================================
-- 3. BIRTH PROFILE
-- ============================================================================

CREATE TABLE birth_profiles (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    full_name       VARCHAR(150) NOT NULL,
    gender          VARCHAR(20),
    birth_date      DATE NOT NULL,
    birth_time      TIME,
    timezone        VARCHAR(50) NOT NULL,
    latitude        NUMERIC(9,6),
    longitude       NUMERIC(9,6),
    city            VARCHAR(150),
    country         VARCHAR(2),
    birth_place     VARCHAR(255),
    is_unknown_time BOOLEAN NOT NULL DEFAULT false,
    note            TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_birth_profiles_user ON birth_profiles(user_id);

-- ============================================================================
-- 4. LOOKUP TABLES
-- ============================================================================

CREATE TABLE zodiac_signs (
    id          SMALLSERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    element     VARCHAR(20),
    modality    VARCHAR(20),
    symbol      VARCHAR(10),
    description TEXT
);

CREATE TABLE planets (
    id          SMALLSERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    symbol      VARCHAR(10),
    description TEXT
);

CREATE TABLE houses (
    id          SMALLSERIAL PRIMARY KEY,
    number      SMALLINT NOT NULL,
    name        VARCHAR(100),
    description TEXT
);

CREATE TABLE aspects (
    id     SMALLSERIAL PRIMARY KEY,
    name   VARCHAR(50) NOT NULL,
    angle  NUMERIC(5,2) NOT NULL,
    orb    NUMERIC(5,2) NOT NULL
);

CREATE TABLE chinese_zodiac (
    id          SMALLSERIAL PRIMARY KEY,
    animal      VARCHAR(50) NOT NULL,
    yin_yang    VARCHAR(10),
    element     VARCHAR(20),
    description TEXT
);

CREATE TABLE five_elements (
    id          SMALLSERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    description TEXT
);

-- ============================================================================
-- 5. TÂY PHƯƠNG CHIÊM TINH
-- ============================================================================

CREATE TABLE planet_positions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    planet_id       SMALLINT NOT NULL REFERENCES planets(id),
    sign_id         SMALLINT NOT NULL REFERENCES zodiac_signs(id),
    house           SMALLINT REFERENCES houses(id),
    degree          SMALLINT NOT NULL CHECK (degree BETWEEN 0 AND 29),
    minute          SMALLINT NOT NULL CHECK (minute BETWEEN 0 AND 59),
    retrograde      BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_planet_positions_profile ON planet_positions(birth_profile_id);

CREATE TABLE natal_aspects (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    planet_a_id      SMALLINT NOT NULL REFERENCES planets(id),
    planet_b_id      SMALLINT NOT NULL REFERENCES planets(id),
    aspect_id        SMALLINT NOT NULL REFERENCES aspects(id),
    orb              NUMERIC(5,2) NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_natal_aspects_profile ON natal_aspects(birth_profile_id);

-- ============================================================================
-- 6. ANALYSIS RESULTS
-- ============================================================================

CREATE TABLE analysis_results (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    analysis_type    analysis_type NOT NULL,
    data             JSONB NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (birth_profile_id, analysis_type)
);

CREATE INDEX idx_analysis_results_profile ON analysis_results(birth_profile_id);
CREATE INDEX idx_analysis_results_data ON analysis_results USING GIN (data);

-- ============================================================================
-- 7. TAROT
-- ============================================================================

CREATE TABLE tarot_readings (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    question         TEXT,
    spread           VARCHAR(50),
    result           JSONB NOT NULL,
    ai_summary       TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tarot_readings_profile ON tarot_readings(birth_profile_id);

-- ============================================================================
-- 8. AI CHAT & REPORTS
-- ============================================================================

CREATE TABLE ai_chat (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id          UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    birth_profile_id UUID REFERENCES birth_profiles(id) ON DELETE SET NULL,
    title            VARCHAR(255),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_ai_chat_user ON ai_chat(user_id);

CREATE TABLE ai_messages (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id    UUID NOT NULL REFERENCES ai_chat(id) ON DELETE CASCADE,
    role       VARCHAR(20) NOT NULL CHECK (role IN ('user','assistant','system')),
    content    TEXT NOT NULL,
    token      INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_ai_messages_chat ON ai_messages(chat_id);

CREATE TABLE report_types (
    code        VARCHAR(50) PRIMARY KEY,
    name        VARCHAR(150) NOT NULL,
    description TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE ai_reports (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    report_type      VARCHAR(50) NOT NULL REFERENCES report_types(code),
    prompt_version   VARCHAR(20),
    model            VARCHAR(50),
    content          TEXT,
    language         VARCHAR(10) DEFAULT 'vi',
    status           report_status NOT NULL DEFAULT 'pending',
    generated_by     VARCHAR(50),
    version          INT NOT NULL DEFAULT 1,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_ai_reports_profile ON ai_reports(birth_profile_id);
CREATE INDEX idx_ai_reports_type ON ai_reports(report_type);

-- ============================================================================
-- 9. KNOWLEDGE BASE (RAG)
-- ============================================================================

CREATE TABLE knowledge_base (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category          VARCHAR(100),
    title             VARCHAR(255) NOT NULL,
    content           TEXT NOT NULL,
    -- embedding       VECTOR(1536),  -- requires pgvector extension
    embedding_model   VARCHAR(100),
    language          VARCHAR(10) DEFAULT 'vi',
    source            VARCHAR(255),
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE prompt_templates (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR(100) NOT NULL,
    system_prompt  TEXT NOT NULL,
    version        INT NOT NULL DEFAULT 1,
    active         BOOLEAN NOT NULL DEFAULT true,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- ============================================================================
-- 10. DAILY HOROSCOPE
-- ============================================================================

CREATE TABLE daily_horoscope (
    id               UUID DEFAULT gen_random_uuid(),
    birth_profile_id UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    date             DATE NOT NULL,
    summary          TEXT,
    career           TEXT,
    love             TEXT,
    health           TEXT,
    finance          TEXT,
    lucky_color      VARCHAR(50),
    lucky_number     SMALLINT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (birth_profile_id, date),
    PRIMARY KEY (id, date)
) PARTITION BY RANGE (date);

-- Create initial partitions for current and next month
CREATE TABLE daily_horoscope_2026_07 PARTITION OF daily_horoscope
    FOR VALUES FROM ('2026-07-01') TO ('2026-08-01');
CREATE TABLE daily_horoscope_2026_08 PARTITION OF daily_horoscope
    FOR VALUES FROM ('2026-08-01') TO ('2026-09-01');

-- ============================================================================
-- 11. COMPATIBILITY
-- ============================================================================

CREATE TABLE compatibility (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_a_id     UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    profile_b_id     UUID NOT NULL REFERENCES birth_profiles(id) ON DELETE CASCADE,
    love_score       SMALLINT CHECK (love_score BETWEEN 0 AND 100),
    friend_score     SMALLINT CHECK (friend_score BETWEEN 0 AND 100),
    marriage_score   SMALLINT CHECK (marriage_score BETWEEN 0 AND 100),
    business_score   SMALLINT CHECK (business_score BETWEEN 0 AND 100),
    summary          TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    CHECK (profile_a_id < profile_b_id),
    UNIQUE (profile_a_id, profile_b_id)
);

-- ============================================================================
-- 12. GAMIFICATION & NOTIFICATIONS
-- ============================================================================

CREATE TABLE achievements (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title       VARCHAR(150) NOT NULL,
    description TEXT,
    icon        VARCHAR(255)
);

CREATE TABLE user_achievements (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id        UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    achievement_id UUID NOT NULL REFERENCES achievements(id) ON DELETE CASCADE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, achievement_id)
);

CREATE TABLE notifications (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title      VARCHAR(255) NOT NULL,
    body       TEXT,
    type       VARCHAR(50),
    read_at    TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_notifications_user_unread
    ON notifications(user_id) WHERE read_at IS NULL;