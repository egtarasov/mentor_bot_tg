
truncate table questions, tasks, pictures, goals, goal_tracks, todo_list, employees, materials, commands;


INSERT INTO commands (id, name, action_id, parent_id)
values
    (1, '/start', 2, default),
    (2, 'Льготы и бонусы', 2, 1),
    (3, 'Питание', 1, 2),
    (4, 'Спорт', 1, 2),
    (5, 'Медецина', 1, 2),
    (6, 'Офис', 1, 2),
    (7, 'Правила и процедуры', 1, 1),
    (8, 'О Компании', 2, 1),
    (9, 'Культура и этика', 1, 8),
    (10, 'Ценности компании', 1, 8),
    (11, 'Карта офиса', 1, 1),
    (12, 'FAQ', 1, 1),
    (13, 'Задачи', 2, 1),
    (14, 'Чек-лист', 2, 13),
    (15, 'Показать чек-лист', 3, 14),
    (16, 'Отметить задачу в чек-листе', 3, 14),
    (17, 'Цели', 3, 13),
    (18, 'В меню', 3, default),
    (19, 'Задать вопрос', 3, 1),
    (20, 'Показать задачи', 3, 13),
    (21, 'Полезные матерьялы для меня', 3, 13),
    (22, 'Календарь', 3, 1);
INSERT INTO materials (message, command_id)
VALUES
    ('Привет, меня зовут Ботик!

Я буду твоим помощником на протяжении всего периода адаптации.
Немного важныx моментов по работе со мной:
– Ниже — главное меню, в которое всегда можно вернуться при помощи кнопки «В меню».
- Все частые вопросы собраны в FAQ, обязательно туда загляни.
- Если появился какой-то вопрос, то можешь задать его с помощью "/ask"', 1),
    ('В нашей компании огромное количество льгот и бонусов для сотрудников, Выбери о чем бы ты хотел узнать подробнее', 2),
    ('Для всех сотрулников ежедневно с 11 до 18 работает кафе на первом этаже. Можешь взять там все, что захочешь.', 3),
    ('В зданиях офиса есть спортивные залы и кабинеты для йоги. Помимо этого, наша копания предоставляет скидки от партнеров.
Подробнее обо все на https://best-company.ru/sport', 4),
    ('Для всех отрудников доступен ДМС. О том, как его оформить и что она дает, ты можешь прочитать здесь https://best-company.ru/medecine', 5),
    ('Для всех сотрудников, проживающих в Москве, круглосуточно доступен офис на ул. Пушкина, д. 5. В офисе всегда присутсвуют
свободные столы, а также мониторы, клвиатуры и другая переферия. Подробности на https://best-company.ru/office', 6),
    ('В нашей компании принят ряд правил, о которых должен знать каждый сотрудник. Смотри:
1. [Устав](https://best-company.ru/ustav.pdf)
2. [Об NDA](https://best-company.ru/nda.pdf)

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
    ('*Чатые вопросы*

Q: *Где можно поесть в офисе?*
A: На первом этаже находится кафе, работающее с 11 до 18.

Q: *Где находится вход в офис?*
A: Офис расположен на ул. Пушкина, д. 5. Вход с торца здания.

Q: *У меня не работает техника, что делать?*
A: В офисе ежедневно работает комманда специалистов, которые помогут тебе со всем разобраться.

Q: *Как у вас дела?*
A: Хорошо, спасибо :)
', 12),
    ('Здесь ты можешь найти всю информацию о своих задачах', 13),
    ('Тут ты сможешь посмотреть на свой чек-лист, а также изменить его в случае необходимости', 14);

insert into pictures (path, command_id)
values ('./images/start.jpg', 1),
       ('./images/cafe.jpg', 3),
       ('./images/sport.jpg', 4),
       ('./images/medicine.jpg', 5),
       ('./images/rules.jpg', 7),
       ('./images/ethic.jpg', 9),
       ('./images/values.jpg', 10);

insert into employees (id, name, surname, telegram_id, occupation_id)
values (1, 'egor', 'tarasov', 1040655631, 1);

insert into todo_list (label, priority, employee_id, completed)
values ('Сходить в кафе', 1, 1, false),
       ('Покушать в офисе', 2, 1, false),
       ('Взять технику', 3, 1, true),
       ('Встретиться с коммандой', 4, 1, false),
       ('Найти рабочий стол', 5, 1, true);


insert into goal_tracks (id, track)
values (1, 'default');

insert into goals (name, description, employee_id, track_id)
values ('Стать успещным', 'Достигнуть успеза', 1, 1);

insert into tasks (name, description, story_points, employee_id)
values ('A', 'Task A', 2, 1),('B', 'Task B', 20, 1),('C', 'Task C', 4, 1);
