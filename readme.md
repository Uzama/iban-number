# IBAN NUMBER

A REST API backend application to validates whether a given IBAN number is valid or not.

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