[![CircleCI](https://circleci.com/gh/victorsteven/manage-jwt.svg?style=svg)](https://circleci.com/gh/victorsteven/manage-jwt)


#### Clone

- Clone this project to your local machine `https://github.com/victorsteven/manage-jwt.git`

#### Heroku App

Application was deployed to Heroku. Use public URL [https://manage-jwt.herokuapp.com](https://manage-jwt.herokuapp.com) with API endpoints.

#### Setup

- Add your database details in the .env file.

- Running Application
  > Run the command below
  ```shell
  $ go run main.go
  ```
- Use `http://localhost:8888` as base url for endpoints

## API Endpoints

| METHOD | DESCRIPTION                             | ENDPOINTS                 |
| ------ | --------------------------------------- | ------------------------- |
| POST   | Register/Signup                         | `/user`                   |
| POST   | Login                                   | `/login`             |
| POST   | Create a todo                           | `/todo`                   |
| POST   | Logout                                  | `/logout`                 |

## Tests

- Run test for all endpoints
  > run the command below(ensure that your test details is setup in the .env file)
  ```shell
  $ go test ./...
  ```

## Author

- [Steven Victor](https://twitter.com/stevensunflash)
