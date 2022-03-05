use contract;
insert into category values
(1, '賃貸'),
(2, '駐車場');

insert into contract values
(UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), '賃貸契約書', 'description', 'note', '2020-01-01 00:00:00', '2020-01-01 00:00:00', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
(UUID_TO_BIN('11eae085-55e6-eb9f-a15d-0242ac110002'), '駐車場契約書', 'description', 'note', '2020-01-01 00:00:00', '2020-01-01 00:00:00', '2020-01-01 00:00:00', '2020-01-01 00:00:00');

insert into contract_category values
(UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), 1),
(UUID_TO_BIN('11eae085-55e6-eb9f-a15d-0242ac110002'), 2);

-- insert into editor values
-- (UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), '賃貸'),
-- (UUID_TO_BIN('11eae085-55e6-e2ca-a15d-0242ac110002'), '駐車場');