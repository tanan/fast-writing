CREATE TABLE `push_notifications` (
   `id` SERIAL PRIMARY KEY,
   `user_id` varchar(32) NOT NULL,
   `created_at` datetime default CURRENT_TIMESTAMP,
   `updated_at` datetime default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
   FOREIGN KEY fk_user(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;