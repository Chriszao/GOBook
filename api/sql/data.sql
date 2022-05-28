INSERT INTO user (name, nick, email, password) 
VALUES 
("User 1", "user1", "user1@gmail.com", "$2a$10$1ouRkT2InkHDQgrm2xhJDOPP1n5SoxE2xekB4hbj9eIHx5/cbbCla"),
("User 2", "user2", "user2@gmail.com", "$2a$10$1ouRkT2InkHDQgrm2xhJDOPP1n5SoxE2xekB4hbj9eIHx5/cbbCla"),
("User 3", "user3", "user3@gmail.com", "$2a$10$1ouRkT2InkHDQgrm2xhJDOPP1n5SoxE2xekB4hbj9eIHx5/cbbCla");

INSERT INTO follower (userId, followerId)
VALUES
(1, 2),
(3, 1),
(1, 3);
