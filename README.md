# Contents ads_storage_api

- [Contents ads_storage_api](#contents-ads_storage_api)
  - [Prepare](#prepare)
  - [Cmd](#cmd)
  - [Docker-compose](#docker-compose)
  - [API Documentation](#api-documentation)
    - [`GET` api/ads/all](#get-apiadsall)
    - [`GET` api/ads/get](#get-apiadsget)
    - [`POST` api/ads/create](#post-apiadscreate)
  - [Program architecture scheme](#program-architecture-scheme)
  - [Database scheme](#database-scheme)

## Prepare

Create .env file in root directory and add following values:

``` dotenv
SERVER_PORT=8080
MARIA_USER=root
MARIA_PASSWORD=secret
MARIA_HOST=db_domain
MARIA_PORT=3306
MARIA_DB=database
```

## CMD

Run

``` sh
go run app/main.go
```

Run tests

``` sh
go test -v
```

**[⬆ Back to Top](#contents-ads_storage_api)**

## Docker-compose

``` sh
docker-compose up -d --build
```

**[⬆ Back to Top](#contents-ads_storage_api)**

## API Documentation

### `GET` api/ads/all

**Метод получения списка объявлений.** На одной странице присутствует 10 объявлений. Возможна сортировка по цене (возрастание/убывание) и по дате создания (возрастание/убывание). Поля в ответе: название объявления, ссылка на главное фото (первое в списке), цена.

#### Parameters

| Name             | Description                                                          |
|------------------|----------------------------------------------------------------------|
| page `*required` | Номер страницы (номер объявления + 10 штук)                          |
| sort_by          | По какому критерию сортировать `price`/`data` (по умолчанию `price`) |
| order            | 0 - по убыванию, 1 - по возрастанию (по умолчанию 1)                 |

#### Example

Valid

``` sh
curl http://localhost:8080/api/ayds/all?page=0&sort_by=price&order=1
```

Invalid

``` sh
curl http://localhost:8080/api/ayds/all?page=0&sort_by=price&order=2
```

#### Response

<table>
<tr>
<th> Status </th> 
<th> Response </th>
</tr>
<tr>
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

``` json
{
    "ok": true,
    "response": [
        {
            "id": 1,
            "title": "title0",
            "main_image": "ref0",
            "price": 0
        },
        {
            "id": 2,
            "title": "title1",
            "main_image": "ref3",
            "price": 1
        },
        {
            "id": 3,
            "title": "title2",
            "main_image": "ref5",
            "price": 2.13
        }
    ]
}
```

</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса<br/>
<b>Example Value:</b>

``` json
{
    "ok": false,
    "error_code": 400,
    "description": "Invalid order param"
}
```

</td>
</tr>
</table>

**[⬆ Back to Top](#contents-ads_storage_api)**

### `GET` api/ads/get

**Метод получения конкретного объявления.** Принимает существующий ID объявления. Возвращает ID брони.

#### Parameters

| Name           | Description                            |
|----------------|----------------------------------------|
| id `*required` | Cуществующий ID объявления             |
| fields         | Опциональные поля: description, images |

#### Example

``` sh
curl http://localhost:8080/api/ads/get?id=1&fields=description
```

#### Response

<table>
<tr>
<th> Status </th>
<th> Response </th>
</tr>
<tr>
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

``` json
{
    "ok": true,
    "response": {
        "id": 1,
        "title": "title0",
        "description": "desc0",
        "price": 0
    }
}
```

</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса<br/>
<b>Example Value:</b>

``` json
{
    "status": false,
    "description": "ID not Found"
}
```

</td>
</tr>
</table>

**[⬆ Back to Top](#contents-ads_storage_api)**

### `POST` api/ads/create

**Метод создания объявления.** Принимает на вход название, описание, несколько ссылок на фотографии(до трех ссылок), цена. Возвращает ID созданного объявления и код результата (ошибка или успех).

#### Parameters

| Name               | Description         |
|--------------------|---------------------|
| title `*required`  | Название объявления |
| description        | Описание объявления |
| images `*required` | Ссылки на картинки  |
| price `*required`  | Цена                |

#### Example

``` sh
curl -X POST -d "title=title1" -d "images=ref1,ref2" -d "price=100.01" http://localhost:8080/api/ads/create
```

#### Response

<table>
<tr>
<th> Status </th>
<th> Response </th>
</tr>
<tr>
<td style="text-align:center"> 200 </td>
<td>
Запрос успешно обработан<br/>
<b>Example Value:</b>

``` json
{
    "ok": true,
    "response": 1
}
```

</td>
</tr>
<tr>
<td style="text-align:center"> 400 </td>
<td>
Ошибка запроса<br/>
<b>Example Value:</b>

``` json
{
    "ok": false,
    "error_code": 400,
    "description": "Title not Found"
}
```

</td>
</tr>
</table>

**[⬆ Back to Top](#contents-ads_storage_api)**

## Program architecture scheme

Классы, отмеченные символами \<I>, — interface; отмеченные символами \<DS> — data structure. Простые стрелки
соответствуют отношениям использования. Стрелки с треугольным наконечником соответствуют отношениям реализации
[<img src="https://live.staticflickr.com/65535/50982424283_4481a7b085_b.jpg" width=900>](https://live.staticflickr.com/65535/50982424283_4481a7b085_b.jpg)
**[⬆ Back to Top](#contents-ads_storage_api)**

## Database scheme

[<img src="https://live.staticflickr.com/65535/50983253337_ac6a5eccb7_b.jpg" width=900>](https://live.staticflickr.com/65535/50983253337_ac6a5eccb7_b.jpg)

**[⬆ Back to Top](#contents-ads_storage_api)**
