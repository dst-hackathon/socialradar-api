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
INSERT INTO users(email) values ('prawit@gmail.com');
INSERT INTO users(email) values ('prakit@gmail.com');
INSERT INTO users(email) values ('opas@gmail.com');
INSERT INTO users(email) values ('atip@gmail.com');

INSERT INTO public.users_options (id,user_id,option_id) VALUES (
7,1,1);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
8,1,2);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
9,1,10);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
10,2,1);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
11,2,2);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
12,2,7);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
13,2,11);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
14,3,8);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
15,3,9);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
16,4,5);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
17,4,10);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
18,4,12);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
19,5,10);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
20,5,12);
INSERT INTO public.users_options (id,user_id,option_id) VALUES (
21,5,5);

INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
8,1,2);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
9,1,1);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
10,1,5);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
11,2,1);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
12,2,4);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
13,2,5);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
14,3,3);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
15,3,4);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
16,3,6);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
17,4,3);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
18,4,2);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
19,4,5);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
20,5,5);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
21,5,3);
INSERT INTO public.users_categories (id,user_id,category_id) VALUES (
22,5,2);
