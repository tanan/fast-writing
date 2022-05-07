use fast_writing;
insert into user values
('fast-writing-db', 'test1', 'test1@example.com', CURRENT_TIMESTAMP())
;

insert into writing_contents values
(1, 'fast-writing-db', 'title1', 'これは説明です。', 'public', 'Advertisement', CURRENT_TIMESTAMP(), '[{"question" : "私は誰ですか？", "answer" : "Who am I?"}, {"question" : "今朝ロンドンは雨が降っています。", "answer" : "It is rainy in London this morning."}]')
;
