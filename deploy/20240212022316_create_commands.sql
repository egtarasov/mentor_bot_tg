
truncate table questions, tasks, pictures, goals, goal_tracks, todo_list, employees, materials, commands;


INSERT INTO commands (id, name, action_id, parent_id)
values
    (1, '/start', 2, default),
    (2, 'Бенефиты', 2, 1),
    (3, 'Питание', 1, 2),
    (4, 'Спорт', 1, 2),
    (5, 'Медецина', 1, 2),
    (6, 'Офис', 1, 2),
    (7, 'Правила и процедуры', 1, 1),
    (8, 'О Компании', 2, 1),
    (9, 'Культура и этика', 1, 8),
    (10, 'Ценности компании', 1, 8),
    (11, 'Карта офиса', 1, 1),
    (12, 'FAQ', 2, 1),
    (13, 'Задачи', 2, 1),
    (14, 'Чек-лист', 2, 13),
    (15, 'Показать чек-лист', 3, 14),
    (16, 'Отметить задачу в чек-листе', 3, 14),
    (17, 'Цели', 3, 13),
    (18, 'В меню', 3, default),
    (19, 'Задать вопрос', 3, 1),
    (20, 'Показать задачи', 3, 13),
    (21, 'Полезные матерьялы для меня', 3, 13),
    (22, 'Календарь', 3, 1),
    (23, 'Щрифты и брендбуки', 1, 8);

INSERT INTO materials (message, command_id)
VALUES
    ('Привет, меня зовут Ботик!

Я буду твоим помощником на протяжении всего периода адаптации.
Немного важныx моментов по работе со мной:
– Ниже — главное меню, в которое всегда можно вернуться при помощи кнопки «В меню».
- Все частые вопросы собраны в FAQ, обязательно туда загляни.
 - Если появился какой-то вопрос, то можешь задать его с помощью "/ask"', 1),
    ('В нашей компании огромное количество бенефитов для сотрудников! Выбери о чем бы ты хотел узнать подробнее', 2),
    ('Для всех сотрудников ежедневно с 11 до 18 работает кафе на первом этаже. Там ты можешь взять все, что захочешь, но только один раз :)', 3),
    ('В зданиях офиса есть спортивные залы и кабинеты для йоги. Помимо этого, наша копания предоставляет скидки от партнеров.
Подробнее обо все на https://best-company.ru/sport', 4),
    ('Для всех отрудников доступен ДМС. О том, как его оформить и что она дает, ты можешь прочитать здесь https://best-company.ru/medecine', 5),
    ('Для всех сотрудников, проживающих в Москве, круглосуточно доступен офис на ул. Пушкина, д. 5. В офисе всегда присутсвуют
свободные столы, а также мониторы, клвиатуры и другая переферия. Подробности на https://best-company.ru/office', 6),
    ('Часто сотрудники сталкиваются с различными проблемами. Чтобы тебе проще было разобраться, мы собрали самые частые проблемы в одном месте, смотри:

1. [Оформить болничный](https://best-company.ru/sick)

2. [О командировке](https://best-company.ru/bussines_trip)

3. [Подписи докумнетов](https://best-company.ru/docs)

4. [Об NDA](https://best-company.ru/nda)

Кроме того, как новый сотрудник, тебе предстоит много всего сделать в первые дни. Чтобы ты ни о чем не забыл, мы собрали для тебя все процедуры, которые тебе предстоит пройти:

1. Получить доступ к ресурсам.

2. Взять рабочую технику.

3. Познакомится с коммандой.

4. Получть электронную подпись (если оформлял при устрйостве на работу).

За поробностями обращайся к своему руководителю и HR-отделу!', 7),
    ('В этом разделе собрана информация о компании: наш этикет, образ жизнь, принципы и другое!', 8),
    ('Как и везде, каждый сотрудник обязуется соблюдать деловой этикет. Это значит, что он:

1. Не разглашает персональные данные и конфиденциальную информацию.

2. Не применяет оценночные суждения в отношении других людей.

3. Уважительно относится к своим коллегам.

Помимо этого, за долгое время в нашей компании сформировалось своя собственная культура, которую мы стараемся популяризовать.

1. Мы все стримимся развиваться и создавать все больше иноваций.

2. Общение - самое главное в работе. Мы одна больша комманда, которая решает все проблемы вместе.

3. Семья - это главное.', 9),

    ('За долги годы у нас появился ряд ценностей, которыми мы дорожим, и которых мы придерживаемся.

Во-первых, это наша команда. Мы считаем, что только вместе мы можем достичь наших целей. Так что дружба и взаимная поддержка - вот что нас объединяет.

Кроме того, мы ценим инновации. Новые идеи и подходы помогают нам развиваться и оставаться на плаву в быстро меняющемся мире.

И, конечно, гибкость. Мы готовы приспосабливаться к любым изменениям и искать решения даже в самых неожиданных ситуациях.

Вот такие ценности у нас в компании. Они помогают нам быть успешными и эффективными!', 10),
    ('Карты офиса, к сожалению, у нас нету. Но ты всегда сможешь разобраться на месте :)', 11),
    ('Чатые вопросы', 12),
    ('Здесь ты можешь найти всю информацию о своих задачах', 13),
    ('Тут ты сможешь посмотреть на свой чек-лист, а также изменить его в случае необходимости', 14),
    ('Шрифты и брендбук - ключевые элементы идентичности нашей компании! Они помогают создать единую и запоминающуюся визуальную атмосферу для всех наших продуктов.

Подробнее про все это ты можешь прочитать [здесь](https://best-compony.ru/style)
', 23);

insert into pictures (path, command_id)
values ('./images/start.jpg', 1),
       ('./images/cafe.jpg', 3),
       ('./images/sport.jpg', 4),
       ('./images/medicine.jpg', 5),
       ('./images/rules.jpg', 7),
       ('./images/ethic.jpg', 9),
       ('./images/values.jpg', 10);

insert into commands (id, name, action_id, parent_id)
values (24, 'Начало работы' ,1, 12),
       (25, 'Работа' ,1, 12),
       (26, 'Активности вне работы' ,1, 12),
       (27, 'Прочее' ,1, 12);

insert into materials (message, command_id)
values
('Q: *Как получить доступы?*
A: Все доступы тебе должен выдать руководитель, обращайся к нему, если что-то не работает.

Q: *Что делать, если выдали не ту технику?*
A: Всеми техническиими вопросами занимается комманда _Service_. Они находятся в офисе на 2-ом этаже.

Q: *Не могу найти нужную переговорку?*
A: Каждая переговорка имеет следующий стиль названия "Номер\_Этажа. Название". Поднимись на нужный этаж и посмотри на диаграмму.', 24),

('Q: *Как работать с Jira*
A: У нас есть мини-найд на [сайте](https://best-compony.ru/working_flow). Там ты сможешь со всем разобраться.

Q: *Не работает zoom на ноутбуке, что делать?*
A: Ознакомся с инструкций на [сайте](https://best-compony.ru/meetings). Если ничего не поможет, то попроси разобраться помочь коллег.

Q: *Не справляюсь с задачами, что делать?*
A: Не переживай, если что-то не получается на старте. Главное, дай знать об этом руководителю. С ним вы гарантированно сможете найти хорошее решение.',25),

('Q: *Есть ли какие-то тусовки в компании?*
A: Конечно, огромное количество сотрудников ежеднедельно играют в настольные игры, занимаются спортом, танцуют и многое другое. Про все ты можешь узнать [здесь](https://best-compony.ru/activities)

Q: *Как зарегесрироваться на мероприятие?*
A: Если ты принял приглашение, то делать больше ничего не нужно.',26),
('Q: *Где можно поесть в офисе?*
A: На первом этаже находится кафе, работающее с 11 до 18.

Q: *Где находится вход в офис?*
A: Офис расположен на ул. Пушкина, д. 5. Вход с торца здания.

Q: *Как получить абонимент в спортзал?*
A: В зале всегда работают наше сотруднк. Просто попроси их помочь.',27);

insert into commands(id, name, action_id, parent_id, is_admin)
values (28, '/admin', 2, default, true),
       (29, 'Вопросы сотрудников', 2, 28, true),
       (30, 'Изменить FAQ', 3, 28, true),
       (31, 'Неотвеченные вопросы', 3, 29, true),
       (32, 'Ответить на вопрос', 3, 29, true);

insert into materials (message, command_id)
values ('Выбери что ты хочешь сделать', 28),
       ('Вопросы сотрудников', 29);


insert into employees (id, name, surname, telegram_id, occupation_id, is_admin, grade)
values (1, 'egor', 'tarasov', 1040655631, 1, true, 17);

insert into todo_list (label, priority, employee_id, completed)
values ('Взять технику для работы', 1, 1, false),
       ('Встретиться с коммандой', 2, 1, false),
       ('Найти рабочий стол', 3, 1, false),
       ('Получить доступы', 4, 1, false),
       ('Пройти [онбординг для новых сотрудников]("https://best-compony.ru/onboarding")', 5, 1, false),
       ('Познакомится с системой Jira', 6, 1, false),
       ('Прочитать полезные матерьялы для своей специальности', 7, 1, false),
       ('Сделать первую задачу', 8, 1, false),
       ('Сходить на встречу с HR', 9, 1, false),
       ('Сходить на встречу с ментором', 10, 1, false),
       ('Записаться в группы по интересу', 11, 1, false);

insert into goal_tracks (id, track)
values (1, 'default');

insert into tasks (name, description, story_points, employee_id)
values ('A', 'Task A', 2, 1),('B', 'Task B', 20, 1),('C', 'Task C', 4, 1);
