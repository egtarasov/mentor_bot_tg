user_id=1
tasks='{"employee_id" : $(user_id),\
"todos" : ["Взять технику для работы", "Встретиться с коммандой", "Найти рабочий стол", "Получить доступы", "Пройти [онбординг для новых сотрудников](\"https://best-compony.ru/onboarding\")",\
"Познакомится с системой Jira","Прочитать полезные матерьялы для своей специальности","Сделать первую задачу","Сходить на встречу с HR","Сходить на встречу с ментором", "Записаться в группы по интересу"],\
"tasks": [\
{"name":"Написать ручку для API", "description":"Нужно сделать ручку в соответсвии с требованиями", "story_points":8, "deadline" : "2023-05-21 16:00"},\
{"name":"Баг с длинными письмами", "description":"Исправить баг с max количесвтом символов в сообщение на почту", "story_points":2, "deadline" : "2023-05-21 16:00"},\
{"name":"Проанализировать датасет с клиентами", "description":"Построить пару статистик по датасету", "story_points":4, "deadline" : "2023-05-21 16:00"}]}'


upload_tasks:
	curl -X POST -v "http://localhost:80/add/tasks" -d $(tasks)


deploy:
	docker-compose up -d


# Сервр
#curl -X GET -v "http://localhost:80/commands"
#
#curl -X PUT -v "http://localhost:80/commands" -d '{"command_id" : 12, "message" : "Частые вопросы и ответы"}'
#
#curl -X GET -v "http://localhost:80/questions"
#
#curl -X POST -v "http://localhost:80/questions" -d '{"question_id" : 4, "answer" : "Все ультра биг"}'
#
#curl -X POST -v -F photo=@./images/start.jpg -F message="hello my friend"  "http://localhost:80/send"