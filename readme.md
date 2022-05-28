# IBAN NUMBER

A REST API backend application to validates whether an IBAN number is valid or not.

## API Reference

```http
  GET /validate
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `iban_number` | `GB33BUKB20201555555555` | **required**. it cannot be empty |

- Response 
```json
    {
        "isValid":true
    }
```