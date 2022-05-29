# IBAN NUMBER

A REST API backend application to validates whether a given IBAN number is valid or not.

## Run Buid Test with makefile
- ```make build```: build the binary according to source OS (linux or macos)
- ```make run```: run the built binary
- ```make test```: run the unit test and display the coverage
- ```make clean```: clean the built files

## Build and run with docker
- run ```docker compose up```

## API Reference

```http
  GET /validate?iban_number=GB33BUKB20201555555555
```

| Parameter | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `iban_number` | `GB33BUKB20201555555555` | **required**. value cannot be empty |

Response 

- 200
```json
    {
        "isValid":true
    }
```

- 400 | 422
```
error message
```