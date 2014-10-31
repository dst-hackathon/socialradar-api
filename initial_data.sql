INSERT INTO questions(text, display_order) values ('What are your favorite menus?', 1);
INSERT INTO questions(text, display_order) values ('What sports do you play?', 2);

INSERT INTO categories (text, question_id, display_order) values ('Thai', 1, 1);
INSERT INTO categories (text, question_id, display_order) values ('Japanese', 1, 2);
INSERT INTO categories (text, question_id, display_order) values ('Chinese', 1, 3);
INSERT INTO categories (text, question_id, display_order) values ('Western', 1, 4);

INSERT INTO categories (text, question_id, display_order) values ('Team', 2, 1);
INSERT INTO categories (text, question_id, display_order) values ('Individual', 2, 2);

INSERT INTO options (text, category_id, display_order) values ('Pad Thai', 1, 1);
INSERT INTO options (text, category_id, display_order) values ('Tom Yum Kung', 1, 2);
INSERT INTO options (text, category_id, display_order) values ('Kao Pad', 1, 3);
INSERT INTO options (text, category_id, display_order) values ('Sushi', 2, 1);
INSERT INTO options (text, category_id, display_order) values ('Ramen', 2, 1);
INSERT INTO options (text, category_id, display_order) values ('Sashimi', 2, 1);
INSERT INTO options (text, category_id, display_order) values ('T-bone steak', 4, 1);
INSERT INTO options (text, category_id, display_order) values ('Fish & chip', 4, 2);
INSERT INTO options (text, category_id, display_order) values ('Irish stew', 4, 3);

INSERT INTO options (text, category_id, display_order) values ('Football', 5, 1);
INSERT INTO options (text, category_id, display_order) values ('Basketball', 5, 2);
INSERT INTO options (text, category_id, display_order) values ('Bowling', 5, 3);

INSERT INTO users(email) values ('pongsanti.tanvejsilp@gmail.com')