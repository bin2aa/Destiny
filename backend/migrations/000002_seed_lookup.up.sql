-- ============================================================================
-- SEED DATA: Lookup tables
-- ============================================================================

-- Zodiac Signs (12 cung hoàng đạo)
INSERT INTO zodiac_signs (name, element, modality, symbol, description) VALUES
    ('Bạch Dương',   'Lửa',   'Thống lĩnh', '♈', 'Cung hoàng đạo đầu tiên, đại diện cho sự khởi đầu và năng lượng'),
    ('Kim Ngưu',     'Đất',   'Cố định',    '♉', 'Cung của sự ổn định, kiên nhẫn và đáng tin cậy'),
    ('Song Tử',      'Khí',   'Biến đổi',   '♊', 'Cung của giao tiếp, linh hoạt và đa năng'),
    ('Cự Giải',      'Nước',  'Thống lĩnh', '♋', 'Cung của cảm xúc, gia đình và bản năng'),
    ('Sư Tử',        'Lửa',   'Cố định',    '♌', 'Cung của sự sáng tạo, lãnh đạo và hào phóng'),
    ('Xử Nữ',        'Đất',   'Biến đổi',   '♍', 'Cung của sự tỉ mỉ, phân tích và phục vụ'),
    ('Thiên Bình',   'Khí',   'Thống lĩnh', '♎', 'Cung của cân bằng, hài hòa và công lý'),
    ('Bọ Cạp',       'Nước',  'Cố định',    '♏', 'Cung của bí ẩn, đam mê và biến đổi'),
    ('Nhân Mã',      'Lửa',   'Biến đổi',   '♐', 'Cung của phiêu lưu, lạc quan và triết học'),
    ('Ma Kết',       'Đất',   'Thống lĩnh', '♑', 'Cung của kỷ luật, trách nhiệm và tham vọng'),
    ('Bảo Bình',     'Khí',   'Cố định',    '♒', 'Cung của đổi mới, nhân đạo và độc lập'),
    ('Song Ngư',     'Nước',  'Biến đổi',   '♓', 'Cung của trực giác, nghệ thuật và lòng trắc ẩn');

-- Planets
INSERT INTO planets (name, symbol, description) VALUES
    ('Mặt Trời', '☉', 'Đại diện cho bản ngã, ý chí và sức sống'),
    ('Mặt Trăng', '☽', 'Đại diện cho cảm xúc, tiềm thức và thói quen'),
    ('Sao Thủy', '☿', 'Đại diện cho giao tiếp, tư duy và di chuyển'),
    ('Sao Kim', '♀', 'Đại diện cho tình yêu, sắc đẹp và giá trị'),
    ('Sao Hỏa', '♂', 'Đại diện cho năng lượng, hành động và ham muốn'),
    ('Sao Mộc', '♃', 'Đại diện cho may mắn, mở rộng và triết học'),
    ('Sao Thổ', '♄', 'Đại diện cho kỷ luật, trách nhiệm và bài học'),
    ('Sao Thiên Vương', '♅', 'Đại diện cho đổi mới, nổi loạn và bất ngờ'),
    ('Sao Hải Vương', '♆', 'Đại diện cho giấc mơ, ảo giác và tâm linh'),
    ('Sao Diêm Vương', '♇', 'Đại diện cho biến đổi, sức mạnh và tái sinh');

-- Houses
INSERT INTO houses (number, name, description) VALUES
    (1,  'Nhà 1 - Bản Ngã',      'Nhà của bản thân, ngoại hình và tính cách'),
    (2,  'Nhà 2 - Tài Chính',    'Nhà của tài sản, giá trị và sở hữu'),
    (3,  'Nhà 3 - Giao Tiếp',     'Nhà của giao tiếp, học tập và anh chị em'),
    (4,  'Nhà 4 - Gia Đình',     'Nhà của gia đình, cội nguồn và cảm xúc'),
    (5,  'Nhà 5 - Sáng Tạo',     'Nhà của sáng tạo, tình yêu và giải trí'),
    (6,  'Nhà 6 - Sức Khỏe',     'Nhà của sức khỏe, công việc và phục vụ'),
    (7,  'Nhà 7 - Hôn Nhân',     'Nhà của đối tác, hôn nhân và quan hệ'),
    (8,  'Nhà 8 - Biến Đổi',     'Nhà của biến đổi, tài sản chung và tâm linh'),
    (9,  'Nhà 9 - Triết Học',    'Nhà của triết học, du lịch và giáo dục cao'),
    (10, 'Nhà 10 - Sự Nghiệp',   'Nhà của sự nghiệp, danh vọng và mục tiêu'),
    (11, 'Nhà 11 - Xã Hội',      'Nhà của bạn bè, nhóm xã hội và ước mơ'),
    (12, 'Nhà 12 - Tâm Linh',    'Nhà của tiềm thức, tâm linh và ẩn dật');

-- Aspects
INSERT INTO aspects (name, angle, orb) VALUES
    ('Góc Hợp (Trine)',     120, 8),
    ('Góc Xung (Opposition)', 180, 8),
    ('Góc Vuông (Square)',   90,  7),
    ('Góc Lục Hợp (Sextile)', 60, 6),
    ('Góc Bán Hợp (Semi-Sextile)', 30, 3),
    ('Góc Bán Vuông (Semi-Square)', 45, 3),
    ('Góc Hợp Nhất (Conjunction)', 0,  10);

-- Chinese Zodiac (12 con giáp)
INSERT INTO chinese_zodiac (animal, yin_yang, element, description) VALUES
    ('Tý',   'Dương', 'Thuỷ', 'Chuột - Thông minh, nhanh nhẹn, tháo vát'),
    ('Sửu',  'Âm',   'Thổ',  'Trâu - Chăm chỉ, đáng tin cậy, kiên nhẫn'),
    ('Dần',  'Dương', 'Mộc',  'Hổ - Dũng cảm, tự tin, quyết đoán'),
    ('Mão',  'Âm',   'Mộc',  'Mèo - Nhẹ nhàng, tinh tế, nghệ thuật'),
    ('Thìn', 'Dương', 'Thổ',  'Rồng - Mạnh mẽ, uy quyền, may mắn'),
    ('Tỵ',   'Âm',   'Hoả',  'Rắn - Khôn ngoan, bí ẩn, sâu sắc'),
    ('Ngọ',  'Dương', 'Hoả',  'Ngựa - Tự do, nhiệt huyết, nhanh nhẹn'),
    ('Mùi',  'Âm',   'Thổ',  'Dê - Hiền lành, nghệ thuật, sáng tạo'),
    ('Thân', 'Dương', 'Kim',  'Khỉ - Thông minh, linh hoạt, hài hước'),
    ('Dậu',  'Âm',   'Kim',  'Gà - Tỉ mỉ, trung thực, đúng giờ'),
    ('Tuất', 'Dương', 'Thổ',  'Chó - Trung thành, đáng tin, bảo vệ'),
    ('Hợi',  'Âm',   'Thuỷ', 'Lợn - Tốt bụng, hào phóng, thoải mái');

-- Five Elements (Ngũ hành)
INSERT INTO five_elements (name, description) VALUES
    ('Kim',  'Kim loại - Đại diện cho sự cứng cáp, quyết đoán, sắc bén'),
    ('Mộc',  'Gỗ - Đại diện cho sự phát triển, linh hoạt, sinh sôi'),
    ('Thuỷ', 'Nước - Đại diện cho sự mềm dẻo, thông thái, giao tiếp'),
    ('Hoả',  'Lửa - Đại diện cho nhiệt huyết, năng lượng, đam mê'),
    ('Thổ',  'Đất - Đại diện cho sự ổn định, nuôi dưỡng, trung dung');

-- Report Types
INSERT INTO report_types (code, name, description) VALUES
    ('western_natal',   'Lá Số Tây Phương',     'Báo cáo chiêm tinh học phương Tây dựa trên ngày sinh'),
    ('bazi',            'Tứ Trụ (Bát Tự)',       'Báo cáo Tứ Trụ - Bát Tự Hà Lạc'),
    ('numerology',      'Thần Số Học',           'Báo cáo Thần số học Pythagoras'),
    ('compatibility',   'Hợp Mệnh',              'Báo cáo tương hợp giữa hai người'),
    ('tarot',           'Bài Tarot',             'Báo cáo luận giải bài Tarot'),
    ('daily_horoscope', 'Vận Trình Ngày',         'Báo cáo vận trình chiêm tinh theo ngày'),
    ('yearly_forecast', 'Vận Trình Năm',          'Báo cáo vận trình chiêm tinh theo năm');

-- Plans
INSERT INTO plans (code, name, price, currency, duration_days, features, is_active) VALUES
    ('premium_monthly',  'Premium Tháng',   99000,  'VND', 30,  '["Xem lá số chi tiết", "Báo cáo Tứ Trụ", "Tarot", "Chat AI không giới hạn"]'::jsonb, true),
    ('premium_yearly',   'Premium Năm',     499000, 'VND', 365, '["Xem lá số chi tiết", "Báo cáo Tứ Trụ", "Tarot", "Chat AI không giới hạn", "Vận trình ngày"]'::jsonb, true),
    ('premium_lifetime', 'Premium Trọn Đời', 1999000, 'VND', 36500, '["Tất cả tính năng", "Ưu tiên AI", "Hỗ trợ ưu tiên"]'::jsonb, true);