use fast_writing;
insert into user values
('fast-writing-db', 'test1', 'test1@example.com', CURRENT_TIMESTAMP())
;

insert into writing_contents values
(1, 'fast-writing-db', 'title1', CURRENT_TIMESTAMP())
;

insert into quiz values
(1, 1, "私は誰ですか？", "Who am I?", CURRENT_TIMESTAMP()),
(2, 1, "今日は晴れています。", "It is fine today.", CURRENT_TIMESTAMP()),
(3, 1, "今朝ロンドンは雨が降っています。", "It is rainy in London in the morning today.", CURRENT_TIMESTAMP())
;