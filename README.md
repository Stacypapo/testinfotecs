Приложение, реализующее систему обработки транзакции.
Сервер имеет три метода:
Send, имеющий эндпоинт POST /api/send;
GetLast, имеющий эндпоинт GET /api/transactions?count=N;
GetBalance, имеющий эндпоинт GET /api/wallet/{address}/balance.
![image](https://github.com/user-attachments/assets/c18c665e-dad7-4faa-bac6-5ec62db67ec4)
Использовался фреймворк Gin и БД postgresql.
Имеется dockerfile для сборки контейнера с приложением.
Для запуска приложения потребуется заменить адрес БД в main
dsn := "user=postgres password=12345 dbname=test host=XXXXXXX port=5432 sslmode=disable"
