INSERT INTO questions(text, tag, display_order) values ('What are your favorite menus?', 'food', 1);		-- 1
INSERT INTO questions(text, tag, display_order) values ('What sports do you play?', 'sport', 2);			-- 2
INSERT INTO questions(text, tag, display_order) values ('What are your favorite movies?', 'movie', 3);		-- 3
INSERT INTO questions(text, tag, display_order) values ('What do music genres belong to you?', 'music', 4); -- 4
INSERT INTO questions(text, tag, display_order) values ('What are your favorite books?', 'book', 5);		-- 5
INSERT INTO questions(text, tag, display_order) values ('What type of trip is right for you?', 'travel', 6);-- 6

INSERT INTO categories (text, question_id, display_order) values ('Thai', 1, 1);		-- 1
INSERT INTO categories (text, question_id, display_order) values ('Japanese', 1, 2);	-- 2
INSERT INTO categories (text, question_id, display_order) values ('Chinese', 1, 3);	 	-- 3
INSERT INTO categories (text, question_id, display_order) values ('Western', 1, 4);		-- 4

INSERT INTO categories (text, question_id, display_order) values ('Extreme', 2, 1);		-- 5
INSERT INTO categories (text, question_id, display_order) values ('Motor', 2, 2); 		-- 6

INSERT INTO categories (text, question_id, display_order) values ('Scifi/Fantasy', 3, 1); 	-- 7
INSERT INTO categories (text, question_id, display_order) values ('Action', 3, 2); 			-- 8
INSERT INTO categories (text, question_id, display_order) values ('Romance', 3, 3); 		-- 9
INSERT INTO categories (text, question_id, display_order) values ('Horror', 3, 4); 			-- 10

INSERT INTO categories (text, question_id, display_order) values ('Pop', 4, 1); 			-- 11
INSERT INTO categories (text, question_id, display_order) values ('Rock', 4, 2); 			-- 12
INSERT INTO categories (text, question_id, display_order) values ('Country', 4, 3); 		-- 13
INSERT INTO categories (text, question_id, display_order) values ('Rap', 4, 4); 			-- 14

INSERT INTO categories (text, question_id, display_order) values ('Biography', 5, 1); 		-- 15
INSERT INTO categories (text, question_id, display_order) values ('Romance', 5, 2); 		-- 16
INSERT INTO categories (text, question_id, display_order) values ('Fantasy', 5, 3); 		-- 17
INSERT INTO categories (text, question_id, display_order) values ('Fiction', 5, 4); 		-- 18

INSERT INTO categories (text, question_id, display_order) values ('Team', 2, 3); 			-- 19
INSERT INTO categories (text, question_id, display_order) values ('Individual', 2, 4); 		-- 20

INSERT INTO categories (text, question_id, display_order) values ('Snow and ski', 6, 1); 			-- 21
INSERT INTO categories (text, question_id, display_order) values ('Cruises', 6, 2); 				-- 22
INSERT INTO categories (text, question_id, display_order) values ('Activity and adventure', 6, 3); 	-- 23

-- OPTIONS --

INSERT INTO options (text, category_id, display_order) values ('Pad Thai', 1, 1);		-- 1
INSERT INTO options (text, category_id, display_order) values ('Tom Yum Kung', 1, 2);	-- 2
INSERT INTO options (text, category_id, display_order) values ('Kao Pad', 1, 3);		-- 3
INSERT INTO options (text, category_id, display_order) values ('Sushi', 2, 1);			-- 4
INSERT INTO options (text, category_id, display_order) values ('Ramen', 2, 1);			-- 5
INSERT INTO options (text, category_id, display_order) values ('Sashimi', 2, 1);		-- 6
INSERT INTO options (text, category_id, display_order) values ('T-bone steak', 4, 1);	-- 7
INSERT INTO options (text, category_id, display_order) values ('Fish & chip', 4, 2);	-- 8
INSERT INTO options (text, category_id, display_order) values ('Irish stew', 4, 3);	 	-- 9

INSERT INTO options (text, category_id, display_order) values ('Caving', 5, 1);			-- 10
INSERT INTO options (text, category_id, display_order) values ('Kitesurfing', 5, 2);	-- 11
INSERT INTO options (text, category_id, display_order) values ('Motocross', 5, 3);		-- 12

INSERT INTO options (text, category_id, display_order) values ('Star wars', 7, 1);				-- 13
INSERT INTO options (text, category_id, display_order) values ('Harry Potter', 7, 2);			-- 14
INSERT INTO options (text, category_id, display_order) values ('The lord of the ring', 7, 3);	-- 15
INSERT INTO options (text, category_id, display_order) values ('Fast and Furious', 8, 1);		-- 16
INSERT INTO options (text, category_id, display_order) values ('Fury', 8, 2);					-- 17
INSERT INTO options (text, category_id, display_order) values ('John Wick', 8, 3);				-- 18
INSERT INTO options (text, category_id, display_order) values ('Begin again', 9, 1);			-- 19
INSERT INTO options (text, category_id, display_order) values ('Blended', 9, 2);				-- 20
INSERT INTO options (text, category_id, display_order) values ('American Pie', 9, 3);			-- 21
INSERT INTO options (text, category_id, display_order) values ('Seed of Chucky', 10, 1);		-- 22
INSERT INTO options (text, category_id, display_order) values ('Gremlins', 10, 2);				-- 23
INSERT INTO options (text, category_id, display_order) values ('Evil dead', 10, 3);				-- 24

INSERT INTO options (text, category_id, display_order) values ('Madonna', 11, 1);				-- 25
INSERT INTO options (text, category_id, display_order) values ('Britney Spears', 11, 2);		-- 26
INSERT INTO options (text, category_id, display_order) values ('Michael Jackson', 11, 3);		-- 27
INSERT INTO options (text, category_id, display_order) values ('Metallica', 12, 1);					-- 28
INSERT INTO options (text, category_id, display_order) values ('Coldplay', 12, 2);					-- 29
INSERT INTO options (text, category_id, display_order) values ('Red hot chilli peppers', 12, 3);	-- 30
INSERT INTO options (text, category_id, display_order) values ('Taylor Swift', 13, 1);			-- 31
INSERT INTO options (text, category_id, display_order) values ('Adam Ryan', 13, 2);				-- 32
INSERT INTO options (text, category_id, display_order) values ('Carrie Underwood', 13, 3);		-- 33
INSERT INTO options (text, category_id, display_order) values ('Pitbull', 14, 1);				-- 34
INSERT INTO options (text, category_id, display_order) values ('Nicki Minaj', 14, 2);			-- 35
INSERT INTO options (text, category_id, display_order) values ('Eminem', 14, 3);				-- 36

INSERT INTO options (text, category_id, display_order) values ('John Lahr''s ''compulsively readable'' biography', 15, 1);				-- 37
INSERT INTO options (text, category_id, display_order) values ('The kind of primary source that historians drool over', 15, 2);		-- 
INSERT INTO options (text, category_id, display_order) values ('Taking Command', 15, 3);											-- 
INSERT INTO options (text, category_id, display_order) values ('Dear Thief', 16, 1);												-- 40
INSERT INTO options (text, category_id, display_order) values ('Curses and Smoke', 16, 2);											-- 
INSERT INTO options (text, category_id, display_order) values ('Year of the Rat', 16, 3);											-- 
INSERT INTO options (text, category_id, display_order) values ('Tribute', 17, 1);													-- 43
INSERT INTO options (text, category_id, display_order) values ('Birth of a new world', 17, 2);										-- 
INSERT INTO options (text, category_id, display_order) values ('We could be heroes', 17, 3);										-- 
INSERT INTO options (text, category_id, display_order) values ('Prince Lestat', 18, 1);												-- 46
INSERT INTO options (text, category_id, display_order) values ('Amnesia', 18, 2);													-- 
INSERT INTO options (text, category_id, display_order) values ('Horror stories edited', 18, 3);										-- 

INSERT INTO options (text, category_id, display_order) values ('Auto racing', 6, 1);			-- 49
INSERT INTO options (text, category_id, display_order) values ('Motor rallying', 6, 2);			-- 50
INSERT INTO options (text, category_id, display_order) values ('Hovercraft racing', 6, 3);		-- 51

INSERT INTO options (text, category_id, display_order) values ('Football', 19, 1);			-- 52
INSERT INTO options (text, category_id, display_order) values ('Basketball', 19, 2);		-- 
INSERT INTO options (text, category_id, display_order) values ('Field hockey', 19, 3);		-- 

INSERT INTO options (text, category_id, display_order) values ('Golf', 20, 1);				-- 55
INSERT INTO options (text, category_id, display_order) values ('Boxing', 20, 2);			-- 
INSERT INTO options (text, category_id, display_order) values ('Fishing', 20, 3);			-- 

INSERT INTO options (text, category_id, display_order) values ('Zermatt', 21, 1);			-- 58
INSERT INTO options (text, category_id, display_order) values ('Jackson Hole', 21, 2);		-- 
INSERT INTO options (text, category_id, display_order) values ('Meribel', 21, 3);			-- 

INSERT INTO options (text, category_id, display_order) values ('Douro', 22, 1);			-- 61
INSERT INTO options (text, category_id, display_order) values ('Rhine', 22, 2);			-- 
INSERT INTO options (text, category_id, display_order) values ('Nile', 22, 3);			-- 

INSERT INTO options (text, category_id, display_order) values ('El Salvador', 23, 1);	-- 64
INSERT INTO options (text, category_id, display_order) values ('Namibia', 23, 2);		-- 
INSERT INTO options (text, category_id, display_order) values ('Cinque Terre', 23, 3);	-- 

INSERT INTO users(email) values ('pongsanti.tanvejsilp@gmail.com');
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
