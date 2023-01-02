# Nigeria Universities API

This API returns all the list of Universities in Nigeria. You can also get a specific univeristy by searching for it using that particualar university abbreviation. E.g. `{uniben, eksu}`.

## How it works.

- To get all universities
  Make a `GET` request to `{your host}/api/v1`

- To get a specific university
  Make a `GET` request to `{your host}/api/v1/{university abbreviation}`

## Want to contribute?

- Clone this repository

```bash
git clone https://github.com/inuoshios/nigeria-universities
```

- Change your directory

```bash
cd nigeria-universities
```

- Run the program

```bash
make run
```

or you could run this command instead

```bash
go run main.go
```
