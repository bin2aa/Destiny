-- Drop indexes first
DROP INDEX IF EXISTS idx_notifications_user_unread;
DROP INDEX IF EXISTS idx_ai_reports_type;
DROP INDEX IF EXISTS idx_ai_reports_profile;
DROP INDEX IF EXISTS idx_ai_messages_chat;
DROP INDEX IF EXISTS idx_ai_chat_user;
DROP INDEX IF EXISTS idx_tarot_readings_profile;
DROP INDEX IF EXISTS idx_analysis_results_data;
DROP INDEX IF EXISTS idx_analysis_results_profile;
DROP INDEX IF EXISTS idx_natal_aspects_profile;
DROP INDEX IF EXISTS idx_planet_positions_profile;
DROP INDEX IF EXISTS idx_birth_profiles_user;
DROP INDEX IF EXISTS idx_payments_transaction;
DROP INDEX IF EXISTS idx_payments_subscription;
DROP INDEX IF EXISTS idx_subscriptions_user;

-- Drop tables (order matters for FK constraints)
DROP TABLE IF EXISTS daily_horoscope CASCADE;
DROP TABLE IF EXISTS daily_horoscope_2026_08 CASCADE;
DROP TABLE IF EXISTS daily_horoscope_2026_07 CASCADE;
DROP TABLE IF EXISTS compatibility CASCADE;
DROP TABLE IF EXISTS user_achievements CASCADE;
DROP TABLE IF EXISTS achievements CASCADE;
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS ai_messages CASCADE;
DROP TABLE IF EXISTS ai_chat CASCADE;
DROP TABLE IF EXISTS ai_reports CASCADE;
DROP TABLE IF EXISTS report_types CASCADE;
DROP TABLE IF EXISTS knowledge_base CASCADE;
DROP TABLE IF EXISTS prompt_templates CASCADE;
DROP TABLE IF EXISTS tarot_readings CASCADE;
DROP TABLE IF EXISTS analysis_results CASCADE;
DROP TABLE IF EXISTS natal_aspects CASCADE;
DROP TABLE IF EXISTS planet_positions CASCADE;
DROP TABLE IF EXISTS five_elements CASCADE;
DROP TABLE IF EXISTS chinese_zodiac CASCADE;
DROP TABLE IF EXISTS aspects CASCADE;
DROP TABLE IF EXISTS houses CASCADE;
DROP TABLE IF EXISTS planets CASCADE;
DROP TABLE IF EXISTS zodiac_signs CASCADE;
DROP TABLE IF EXISTS payments CASCADE;
DROP TABLE IF EXISTS subscriptions CASCADE;
DROP TABLE IF EXISTS plans CASCADE;
DROP TABLE IF EXISTS birth_profiles CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Drop ENUM types
DROP TYPE IF EXISTS report_status;
DROP TYPE IF EXISTS analysis_type;
DROP TYPE IF EXISTS payment_status;
DROP TYPE IF EXISTS subscription_status;