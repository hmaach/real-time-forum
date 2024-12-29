-- Insert Users
INSERT INTO users (email, username, password) VALUES
('user1@example.com', 'User1', 'password1'),
('user2@example.com', 'User2', 'password2'),
('user3@example.com', 'User3', 'password3'),
('user4@example.com', 'User4', 'password4'),
('user5@example.com', 'User5', 'password5'),
('user6@example.com', 'User6', 'password6'),
('user7@example.com', 'User7', 'password7'),
('user8@example.com', 'User8', 'password8'),
('user9@example.com', 'User9', 'password9'),
('user10@example.com', 'User10', 'password10'),
('user11@example.com', 'User11', 'password11'),
('user12@example.com', 'User12', 'password12'),
('user13@example.com', 'User13', 'password13'),
('user14@example.com', 'User14', 'password14'),
('user15@example.com', 'User15', 'password15'),
('user16@example.com', 'User16', 'password16'),
('user17@example.com', 'User17', 'password17'),
('user18@example.com', 'User18', 'password18'),
('user19@example.com', 'User19', 'password19'),
('user20@example.com', 'User20', 'password20');

-- Insert Posts
INSERT INTO posts (user_id, title, content) VALUES
(1, 'Post 1 Title', 'Content of post 1'),
(2, 'Post 2 Title', 'Content of post 2'),
(3, 'Post 3 Title', 'Content of post 3'),
(4, 'Post 4 Title', 'Content of post 4'),
(5, 'Post 5 Title', 'Content of post 5'),
(6, 'Post 6 Title', 'Content of post 6'),
(7, 'Post 7 Title', 'Content of post 7'),
(8, 'Post 8 Title', 'Content of post 8'),
(9, 'Post 9 Title', 'Content of post 9'),
(10, 'Post 10 Title', 'Content of post 10'),
(11, 'Post 11 Title', 'Content of post 11'),
(12, 'Post 12 Title', 'Content of post 12'),
(13, 'Post 13 Title', 'Content of post 13'),
(14, 'Post 14 Title', 'Content of post 14'),
(15, 'Post 15 Title', 'Content of post 15'),
(16, 'Post 16 Title', 'Content of post 16'),
(17, 'Post 17 Title', 'Content of post 17'),
(18, 'Post 18 Title', 'Content of post 18'),
(19, 'Post 19 Title', 'Content of post 19'),
(20, 'Post 20 Title', 'Content of post 20'),
(1, 'Post 21 Title', 'Content of post 21'),
(2, 'Post 22 Title', 'Content of post 22'),
(3, 'Post 23 Title', 'Content of post 23'),
(4, 'Post 24 Title', 'Content of post 24'),
(5, 'Post 25 Title', 'Content of post 25'),
(6, 'Post 26 Title', 'Content of post 26'),
(7, 'Post 27 Title', 'Content of post 27'),
(8, 'Post 28 Title', 'Content of post 28'),
(9, 'Post 29 Title', 'Content of post 29'),
(10, 'Post 30 Title', 'Content of post 30'),
(11, 'Post 31 Title', 'Content of post 31'),
(12, 'Post 32 Title', 'Content of post 32'),
(13, 'Post 33 Title', 'Content of post 33'),
(14, 'Post 34 Title', 'Content of post 34'),
(15, 'Post 35 Title', 'Content of post 35');

-- Link Posts with Categories
INSERT INTO post_category (post_id, category_id) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5),
(6, 1), (7, 2), (8, 3), (9, 4), (10, 5),
(11, 1), (12, 2), (13, 3), (14, 4), (15, 5),
(16, 1), (17, 2), (18, 3), (19, 4), (20, 5),
(21, 1), (22, 2), (23, 3), (24, 4), (25, 5),
(26, 1), (27, 2), (28, 3), (29, 4), (30, 5),
(31, 1), (32, 2), (33, 3), (34, 4), (35, 5);

-- Insert Comments
INSERT INTO comments (user_id, post_id, content) VALUES
(1, 1, 'Comment 1 on Post 1'),
(2, 2, 'Comment 2 on Post 2'),
(3, 3, 'Comment 3 on Post 3'),
(4, 4, 'Comment 4 on Post 4'),
(5, 5, 'Comment 5 on Post 5'),
(6, 6, 'Comment 6 on Post 6'),
(7, 7, 'Comment 7 on Post 7'),
(8, 8, 'Comment 8 on Post 8'),
(9, 9, 'Comment 9 on Post 9'),
(10, 10, 'Comment 10 on Post 10'),
(11, 11, 'Comment 11 on Post 11'),
(12, 12, 'Comment 12 on Post 12'),
(13, 13, 'Comment 13 on Post 13'),
(14, 14, 'Comment 14 on Post 14'),
(1, 15, 'Comment 15 on Post 15');

-- Insert Post Reactions (Likes and Dislikes)
INSERT INTO post_reactions (user_id, post_id, reaction) VALUES
(1, 1, 'like'), (2, 1, 'dislike'), (3, 2, 'like'), (4, 2, 'dislike'),
(5, 3, 'like'), (6, 4, 'dislike'), (7, 5, 'like'), (8, 6, 'dislike'),
(9, 7, 'like'), (10, 8, 'dislike'), (11, 9, 'like'), (12, 10, 'dislike'),
(13, 11, 'like'), (14, 12, 'dislike'), (15, 13, 'like'), (16, 14, 'dislike');

-- Insert Comment Reactions (Likes and Dislikes)
INSERT INTO comment_reactions (user_id, comment_id, reaction) VALUES
(1, 1, 'like'), (2, 2, 'dislike'), (3, 3, 'like'), (4, 4, 'dislike'),
(5, 5, 'like'), (6, 6, 'dislike'), (7, 7, 'like'), (8, 8, 'dislike'),
(9, 9, 'like'), (10, 10, 'dislike'), (11, 11, 'like'), (12, 12, 'dislike'),
(13, 13, 'like'), (14, 14, 'dislike'), (15, 15, 'like');
