# RESTful Cart Module API document

## 1 Update Food Quantity Interface

### 1.1 Interface Description    

- Get business list

### 1.2 Address  

`{apiAddress}/api/cart/update-food-quantity`

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
| uid_g   | Yes | string | 1 < length < 20 | guest id to find cart |
| uid_h   | Yes | string | 1 < length < 20 | host id to find cart |
| fid  | Yes | int | 0 < value | a bit controls whether to only show opening business |
| quantity  | Yes | int | 1 < value < 1000 | quantity of the food |


**Special Note**:
1. We use uid_g to mark the userid of a guest and a uid_h to mark the userid of a host.
2. fid is the food id in the host.
3. quantity should be a valid number. the validation logic should be hosted in frontend code.
​    

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Update food quantity succeed",  // 提示信息
    "data": null
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 2 Add Item To Cart Interface

### 2.1 Interface Description    

- Add a food into cart 

### 2.2 Address  

`{apiAddress}/api/cart/add-food`

### 2.3 Request Type  

**POST**  

### 2.4 Request Parameters  

#### 2.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 2.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| uid_g   | Yes | string | 1 < length < 20 | guest id to find cart |
| uid_h   | Yes | string | 1 < length < 20 | host id to find cart |
| fid  | Yes | int | 0 < value | a bit controls whether to only show opening business |
| quantity  | Yes | int | 1 < value < 1000 | quantity of the food |


**Special Note**:
1. We use uid_g to mark the userid of a guest and a uid_h to mark the userid of a host.
2. fid is the food id in the host.
3. quantity should be a valid number. the validation logic should be hosted in frontend code.
4. if quantity drop to zero, we update in backend to remove the item from the cart. Next time when cart content is retrieved, the removed food won't be shown up.
​    

### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Add food to cart succeed",  // 提示信息
    "data": null
}
```

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Remove Item from Cart Interface

### 3.1 Interface Description    

-  Remove item from cart

### 3.2 Address  

`{apiAddress}/api/cart/remove-food`

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
| uid_g   | Yes | string | 1 < length < 20 | guest id to find cart |
| uid_h   | Yes | string | 1 < length < 20 | host id to find cart |
| fid  | Yes | int | 0 < value | a bit controls whether to only show opening business |


**Special Note**:
1. We use uid_g to mark the userid of a guest and a uid_h to mark the userid of a host.
2. fid is the food id in the host.
​    

### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Remove food to cart succeed",  // 提示信息
    "data": null
}
```

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  




## 4 Load Cart Interface

### 4.1 Interface Description    

-  Load food in the cart from backend

### 4.2 Address  

`{apiAddress}/api/cart/load`

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
| uid_g   | Yes | string | 1 < length < 20 | guest id to find cart |
| uid_h   | Yes | string | 1 < length < 20 | host id to find cart |

**Special Note**:
1. We use uid_g to mark the userid of a guest and a uid_h to mark the userid of a host.

### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Got food to cart succeed",  // 提示信息
    "data": [
      {
        "fid": 1,
        "food_name": "apple pie",
        "quantity": 2,
        "price": 13.99,
        "icon": "<picurl>"
      },
      {
        "fid": 2,
        "food_name": "shrimp roll",
        "quantity": 1,
        "price": 12.99,
        "icon": "<picurl>"
      }
    ]
}
```

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  