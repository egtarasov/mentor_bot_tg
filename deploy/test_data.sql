
insert into employees (name, surname, telegram_id, occupation_id)
values ('egor', 'tarasov', 1040655631, 1);

insert into todo_list (label, priority, employee_id, completed)
values ('Сходить в кафе', 1, 1, false),
       ('Покушать в офисе', 2, 1, false),
       ('Взять технику', 3, 1, true),
       ('Встретиться с коммандой', 4, 1, false),
       ('Найти рабочий стол', 5, 1, true);


insert into goal_tracks (track)
values ('default');

insert into goals (name, description, employee_id, track_id)
values ('Стать успещным', 'Достигнуть успеза', 1, 1);