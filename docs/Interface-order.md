# RESTful Order Module API document

## 1 Get Order Status Interface

### 1.1 Interface Description    

- Get the status of an order

### 1.2 Address  

`{apiAddress}/api/order/get-order-status`

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
| oid   | Yes | string | 1 < length < 64 | order id to find the order |
| uid   | Yes | string | 1 < length < 20 | user id to get the order |
| isHost  | Yes | int | {0,1} | a bit controls whether this is a request from host |


**Special Note**:
1. We set up uid here as a purpose to store logs in the backend, as well as counter malicious get requests
2. Here we combine the order getter of pickup order and deliver order. For deliver order guest has to provide deliver address. 
3. In terms of workflow, pickup order has a special step named "waiting for batch filling" after order was placed. Only after the number of order was over a threshold can the order being valid.
​    

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Update food quantity succeed",  // 提示信息
    "data": {
            "uid": "2312f12dab003e0e",
            "oid": "<SHA256 number>",
            "total_price": 13.99,
            "detailed_price": {
              "fried toufu": 5.99, 
              "vegies": 6.00,
              "rice": 2.00
            },
            "payment_options": ["Zelle", "Venmo", "cash"],
            "create_time": "20230501T23:59:59",
            "last_time": "20230502T07:23:32",
            "shipping_method": "pickup",
            "is_group": 1, // is this order a group order
            "state_history": [
              {"order created": "<datetime>"},
              {"waiting for payment": "<datetime>"},
              {"order paid": "<datetime>"},
              {"shipped": "<datetime>"},
              ]
            "host_info": {
              "host_name": "Stan Liu",
              "host_address": "125 Lexington Ave."
            },
            "guest_info": {
              "guest_name": "Bob Rophy",
              "guest_address": null,
            },
            "notes_to_host": "The toufu should be less spicy"
        }
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 2 Create Order Interface

### 2.1 Interface Description    

- Create a new order from cart

### 2.2 Address  

`{apiAddress}/api/order/create-order`

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
| uid_g   | Yes | string | 1 < length < 20 | guest id to bind with the order |
| uid_h   | Yes | string | 1 < length < 20 | host id to bind with the order |
| is_group  | Yes | int | {0,1} | a bit indicates if this is a group order |


**Special Note**:
1. We use uid_g to mark the userid of a guest and a uid_h to mark the userid of a host.​    

### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "order creation succeed",  // 提示信息
    "data": null
}
```

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Cancel Order Interface

### 3.1 Interface Description    

- Cancel pending order

### 3.2 Address  

`{apiAddress}/api/order/cancel-order`

### 3.3 Request Type  

**PUT**  

### 3.4 Request Parameters  

#### 3.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 3.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| oid   | Yes | string | 1 < length < 64 | order id to find the order |
| uid   | Yes | string | 1 < length < 20 | user id to get the order |
| isHost  | Yes | int | {0,1} | a bit controls whether this is a request from host |


**Special Note**:
1. We set up uid here as a purpose to store logs in the backend, as well as counter malicious requests
2. Cancel order could be done from either side.
​    

### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Order cancel succeed",  // 提示信息
    "data": null
}
```

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 4 Retrieve Orders Interface

### 4.1 Interface Description    

-  Load all history orders in the onto the front page from backend

### 4.2 Address  

`{apiAddress}/api/order/retrieve-orders`

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
| uid   | Yes | string | 1 < length < 20 | guest id to find cart |
| time_start   | Yes | string | 1 < length < 20 | start time of time range of orders |
| time_end   | Yes | string | 1 < length < 20 | end time of time range of orders |
| startNum | Yes | int | 1 < value | start index of the orders in the sorted list |
| quantity   | Yes | int | startNum < value | end index of the orders in the sorted list |

**Special Note**:
1. We use the start_time and end_time in request refer to the time when order is created
2. 

 ### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Got food to cart succeed",  // 提示信息
    "data": [
        {
            "uid": "2312f12dab003e0e",
            "oid": "<SHA256 number>",
            "total_price": 15.89,
            "payment_method": "Zelle",
            "create_time": "20230423T23:59:59",
            "last_time": "20230424T12:26:32",
            "shipping_method": "delivery",
            "state": "delivered",
            "host_info": {
              "host_name": "Berney Bane"
            },
            "guest_info": {
              "guest_name": "Bob Rophy"
            },
            "notes_to_host": "No salt on potato chips."
        }
    ]
}
```

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 5 Update Order State Interface

### 5.1 Interface Description    

-  Load all history orders in the onto the front page from backend

### 5.2 Address  

`{apiAddress}/api/order/retrieve-orders`

### 5.3 Request Type  

**POST**  

### 5.4 Request Parameters  

#### 5.4.1 Header Parameters  

| Key       | Must | Type/Value      | Note         |
| ------------ | ---- | ---------------- | ------------ |
| Content-Type | Yes   | application/json | Request parameter type |

#### 5.4.2 Body Parameters  

| Key    | Must | Type   | Limit        | Note     |
| --------- | ---- | ------ | --------------- | -------- |
| oid   | Yes | string | 1 < length < 64 | order id to find the order |
| uid   | Yes | string | 1 < length < 20 | user id to get the order |
| isHost  | Yes | int | {0,1} | a bit controls whether this is a request from host |
| newState | Yes | string | 1 < length < 50 | new order state |

**Special Note**:
1. new state should be one of the five:
```
1. waiting for payment
2. waiting for grouping
3. confirmed
4. start to deliver
5. delivered
6. canceled
```

 ### 5.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "order state updated!",  // 提示信息
    "data": null
}
```

### 5.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  