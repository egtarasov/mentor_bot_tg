Что умеет аминка:
1. Посмотреть на вопросоы +
2. Ответить на вопрос +
3. Изменить содержание сообщения, которое выводит шаблонная комманда +
4. Показать комманды 1 и 2 типа (id, action_id, message, comand_name) +
5. Добавить новую коммманду с привязкой к старой. Тип может быть только 1 или 2. 
Должно содержаться название комманды, строка с матерьялом, тип.


curl -X GET -v "http://localhost:80/commands"

curl -X PUT -v "http://localhost:80/commands" -d '{"command_id" : 12, "message" : "Частые вопросы и ответы"}'

curl -X GET -v "http://localhost:80/questions"

curl -X POST -v "http://localhost:80/questions" -d '{"question_id" : 4, "answer" : "Все ультра биг"}'

curl -X POST -v -F photo=@./images/start.jpg -F message="hello my friend"  "http://localhost:80/send"
