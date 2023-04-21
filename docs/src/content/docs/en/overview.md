---
title: "Overview"
description: "General Overview of accessing Data Dojo API"
---

## Requests

### Base URL

The base address of the Web API is `https://datadojo.app/api/`

### HTTP Utilization

All resources can be accessed by standard HTTP requests in UTF-8 format.
(_While technically the API follows CRUD principles, GET is the only
HTTP verb utilized for accessing the API_)

## Responses

### Status Codes

The Data Dojo API utilizes the following HTTP status codes:
| Status Code | Description
|-------------- | --------------
| 200 | OK - request has succeeded. Will return a readable result
| 400 | Bad Request - the request was invalid or cannot be otherwise served
| 404 | Not Found - the URI requested is invalid or the resource requested does not exist
| 500 | Internal Server Error - something went wrong on the Data Dojo server side

### Error Responses

The Data Dojo API utilizes the following error response format (400 and 500 code responses):
| Field | Description
|-------------- | --------------
| error | Type of error encountered.
| error_description | Detailed description of cause of error.

#### Example Error Response

```
$ curl https://datadojo.app/api/games/9999999
{
    "error": "not_found",
    "error_description": "Could not find game"
}

```

### Time Stamps

All responses with a time related field will be returned in [ISO_8601](https://en.wikipedia.org/wiki/ISO_8601) zero offset format.
(_YYYY-MM-DDTHH:MM:SSZ_)
