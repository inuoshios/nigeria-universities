# Nigeria Universities API

This api returns all the list of Universities in Nigeria. You can also get a specific univeristy by searching for it using that particualar university abbreviation. E.g. `{uniben, eksu}`.

## How it works.

- To get all universities
  Make a `GET` request to `{your host}/api/v1`

- To get a specific university
  Make a `GET` request to `{your host}/api/v1/{university abbreviation}`

## Want to contribute?

Make sure you have Go installed on your machine (well, not neccesary, but you understand).

- Clone this repository

```bash
git clone https://github.com/0xmlx/nigeria-universities
```

- Change your directory

```bash
cd nigeria-universities
```

- Run the program

```bash
go run ./cmd/api
```

Before you run the server, make sure your run `source {envname}` on your preferred terminal. You can also use `joho/godotenv` if you like. 🤷‍♂️
