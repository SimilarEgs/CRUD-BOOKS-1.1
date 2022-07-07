
## How to use

- Clone this repository
- Create `.env` file in a root folder:
```PowerShell
~/CRUD-BOOKS-1.1 touch .env
```
- Paste next inside created .env file:
```.env
DB_HOST=127.0.0.1
DB_DRIVER=mysql 
DB_USER="fill with ur DB user"
DB_PASSWORD="fill with ur DB password"
DB_NAME="fill with ur DB name"
DB_PORT=3306 
```
- Run code by typing in console:
```PowerShell
~/CRUD-BOOKS go run main.go
```

## API
___create book___

![CREATE-BOOK](https://user-images.githubusercontent.com/90198202/173189850-3b3d682d-6908-4b64-abbc-8f30a96f9f92.jpg)

___get book by ID___

![GET-BOOK_BY_ID](https://user-images.githubusercontent.com/90198202/173189925-4e162d94-9668-42de-8d43-c2a0965ca2e4.jpg)

___get all books___

![GET-BOOKS](https://user-images.githubusercontent.com/90198202/173190132-1f3a709d-709b-470a-b157-e1594325e513.jpg)

___update book by ID___

![UPDATE-BOOK_BY_ID](https://user-images.githubusercontent.com/90198202/173190311-f6ed44d7-b3cc-4085-ab75-2fc49c1820fa.jpg)

___delete book by ID___

![DELETE-BOOK_BY_ID](https://user-images.githubusercontent.com/90198202/173190408-cc6f2847-7c89-458e-8c8b-cea6bab90bfb.jpg)
![DELETE-BOOK_BY_ID_AFFTER](https://user-images.githubusercontent.com/90198202/173190465-70dca06b-3f8f-43d8-a9db-f65a3e4598a8.jpg)
![DELETED_AT](https://user-images.githubusercontent.com/90198202/173190509-a139a1e3-66b2-411f-9698-f45486b95549.jpg)
