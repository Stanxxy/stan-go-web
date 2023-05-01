# RESTful Food Module API document



## 1 Get Food From Business Interface 

### 1.1 Interface Description    

- Get the food provide by a business

### 1.2 Address  

`{apiAddress}/api/business/get-businesses-food`

### 1.3 Request Type  

**POST**  

### 1.4 Request Parameters  

#### 1.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 1.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| uid_h   | Yes | string | 1 < length < 20 | user id to locate a host |
| startNum | Yes | int | 1 < value | start index of the food in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the food in the sorted list |


**Special Note**:
1. We could add uid as an input here, once we could do personalized recommendation
2. We cache all the in-day food in Cassandra

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Get business list successful",  // 提示信息
    "data": [
        {
            "uid": "2312f12dab003e0e",
            "fid": 1,
            "food_name": "foodiePathPie",
            "food_ingradients": ["apple", "power","eggs"],
            "food_notes": "Clients may get allergic against blablabla",
            "food_order_cut_time": "20230501T23:59:59",
            "number": 1,
            "pic": "<picurl>"
        },
        {
            "uid": "2312f12dab003e0e",
            "fid": 2,
            "food_name": "bestieShopBowl",
            "food_ingradients": ["white rice", "cucumber","mushroom"],
            "food_notes": "Much carbon, may cause food coma",
            "food_order_cut_time": "202300501T21:59:59",
            "number": 3,
            "pic": "<picurl>"
        }
    ] 
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 1 Get Food From Keyword Interface 

### 1.1 Interface Description    

- Get the food by using a keyword to do search

### 1.2 Address  

`{apiAddress}/api/business/get-food-from-keyword`

### 1.3 Request Type  

**POST**  

### 1.4 Request Parameters  

#### 1.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 1.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| keyword   | Yes | string | 1 < length < 20 | the keyword to search for food |
| startNum | Yes | int | 1 < value | start index of the food in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the food in the sorted list |


**Special Note**:
1. We could add uid as an input here, once we could do personalized recommendation
2. We cache all the in-day food in Cassandra

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Get business list successful",  // 提示信息
    "data": [
        {
            "uid": "2312f12dab003e0e",
            "fid": 1,
            "food_name": "foodiePathPie",
            "food_ingradients": ["apple", "power","eggs"],
            "food_notes": "Clients may get allergic against blablabla",
            "food_order_cut_time": "20230501T23:59:59",
            "number": 1,
            "pic": "<picurl>"
        },
        {
            "uid": "2312f12dab003e0e",
            "fid": 2,
            "food_name": "bestieShopBowl",
            "food_ingradients": ["white rice", "cucumber","mushroom"],
            "food_notes": "Much carbon, may cause food coma",
            "food_order_cut_time": "202300501T21:59:59",
            "number": 3,
            "pic": "<picurl>"
        }
    ] 
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  