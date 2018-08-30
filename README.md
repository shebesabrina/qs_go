# Quantified Self Back-end Golang

## Initial Setup

1. Clone this repository and rename the repository to `qs_rails` in one command

  ```shell
  https://shebesabrina.github.io/quantified-self-fe/index.html
  ```
2. Change directory into the `qs_rails` directory
  ```
  cd qs_go
  ```

3. Install the dependencies

  ```shell
  go build
  ```

3. Run test suite

  ```shell
    go test -v
  ```

## Running the Server Locally

To see your code in action locally, you need to fire up a development server. Use the command:

```shell
go run $(ls -1 *.go | grep -v _test.go)
```

Once the server is running, visit API endpoints in your browser:

* `http://localhost:8080/` to run your application. Enpoints are available in the * [Project Spec](https://github.com/turingschool/backend-curriculum-site/blob/gh-pages/module4/projects/quantified-self/quantified-self.md)

## Deployed
* Back end is deployed here: https://qs-go-mine.herokuapp.com/api/v1/foods

* To see the [front end] visit: (https://shebesabrina.github.io/quantified-self-fe/index.html)

