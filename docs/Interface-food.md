# RESTful Food Module API document

## 1 Get Food From Business Interface 

### 1.1 Interface Description    

- Get the food provide by a business

### 1.2 Address  

`{apiAddress}/api/business/get-businesses-food`

### 1.3 Request Type  

**GET**  

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



## 2 Get Food From Keyword Interface 

### 2.1 Interface Description    

- Get the food by using a keyword to do search

### 2.2 Address  

`{apiAddress}/api/business/get-food-from-keyword`

### 2.3 Request Type  

**GET**  

### 2.4 Request Parameters  

#### 2.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 2.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| keyword   | Yes | string | 1 < length < 20 | the keyword to search for food |
| startNum | Yes | int | 1 < value | start index of the food in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the food in the sorted list |


**Special Note**:
1. We could add uid as an input here, once we could do personalized recommendation
2. We cache all the in-day food in Cassandra

### 2.5 Sample Response

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

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Add Food to Business Interface 

### 3.1 Interface Description    

- Get the food provide by a business

### 3.2 Address  

`{apiAddress}/api/business/Add-business-food`

### 3.3 Request Type  

**POST**  

### 3.4 Request Parameters  

#### 3.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 3.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| uid_h   | Yes | string | 1 < length < 20 | user id to locate a host |
| food_data | Yes | food_card | N/A | The information used for providing food card |


**Special Note**:
1. Food card is shown as in the getter function
```
{
    "uid": "2312f12dab003e0e",
    "food_name": "foodiePathPie",
    "food_ingradients": ["apple", "power","eggs"],
    "food_notes": "Clients may get allergic against blablabla",
    "food_order_cut_time": "20230501T23:59:59",
    "number": 1,
    "pic": ["<picurl1>","<picurl2>","<picurl3>"]
}
```
2. We allow food to be copied. As a result, we could first allow user to retrieve food data from history food poster and refill the form in frontend. One key effect of this feature is that frontend is responsible for checking datetime and number, and backend is responsible for priventing injection.

### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Add food to business sucessfully",  // 提示信息
    "data": null
}
```

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 4 Update Food in Business Interface 

### 4.1 Interface Description    

- Update the food provide by a business

### 4.2 Address  

`{apiAddress}/api/business/update-business-food`

### 4.3 Request Type  

**POST**  

### 4.4 Request Parameters  

#### 4.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 4.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| uid_h   | Yes | string | 1 < length < 20 | user id to locate a host |
| food_data | Yes | food_card | N/A | The information used for providing food card |


**Special Note**:
1. Food card is shown as in the getter function
```
{
    "uid": "2312f12dab003e0e",
    "food_name": "foodiePathPie",
    "food_ingradients": ["apple", "power","eggs"],
    "food_notes": "Clients may get allergic against blablabla",
    "food_order_cut_time": "20230501T23:59:59",
    "number": 1,
    "pic": ["<picurl1>","<picurl2>","<picurl3>"]
}
```

### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Add food to business sucessfully",  // 提示信息
    "data": null
}
```

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  




##  5 Take off Food from Business Shelves Interface 

### 5.1 Interface Description    

- Take off the food from business

### 5.2 Address  

`{apiAddress}/api/business/take-off-food-from-business`

### 5.3 Request Type  

**POST**  

### 5.4 Request Parameters  

### 5.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

### 5.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| uid_h   | Yes | string | 1 < length < 20 | user id to locate a host |
| food_fid_list | Yes | list[int] | 1 < length | The fid used for selecting food card |


### 5.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Remove food from business sucessfully",  // 提示信息
    "data": null
}
```

### 5.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  