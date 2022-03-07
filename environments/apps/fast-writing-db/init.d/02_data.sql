use fast_writing;
insert into user values
(UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), 'test1', 'test1@example.com', CURRENT_TIMESTAMP())
;

insert into writing_contents values
(1, UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), 'title1', CURRENT_TIMESTAMP())
;