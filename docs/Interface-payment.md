# RESTful Payment Module API document

## 1 Get Payment Method Status Interface

### 1.1 Interface Description    

- Get all possible pament method of an order

### 1.2 Address  

`{apiAddress}/api/payment/get-payment-method`

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
| uid   | Yes | string | 1 < length < 20 | user id to find the payment info |
​    

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Update food quantity succeed",  // 提示信息
    "data": {
            "uid": "2312f12dab003e0e",
            "zelle": "<zelle-account-name>",
            "venmo": "<venmo-account-name>",
            "paypal": "<paypal-account-name>",
        }
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 2 Confirm Payment Interface

### 2.1 Interface Description    

- Confirm payment from the guest. Request when payment was finished from the guest side.

### 2.2 Address  

`{apiAddress}/api/payment/confirm-payment`

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
| oid   | Yes | string | 1 < length < 64 | order id to find the order |
| uid   | Yes | string | 1 < length < 20 | user id to get the order |


### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "payment confirmed",  // 提示信息
    "data": null
}
```

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Receive Payment Interface

### 3.1 Interface Description    

- Receive payment from guest. Request after payment is received from the host side.

### 3.2 Address  

`{apiAddress}/api/payment/receive-payment`

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
​    

### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "payment received",  // 提示信息
    "data": null
}
```

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 4 Dispute Payment Interface

### 4.1 Interface Description    

-  Dispute when problem happened during payment.

### 4.2 Address  

`{apiAddress}/api/payment/dispyte-payment`

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
| oid   | Yes | string | 1 < length < 64 | order id to find the order |
| uid   | Yes | string | 1 < length < 20 | user id to get the order |
| context | Yes | string | 1 < length < 1024 | content for the issue |
| attachments  | Yes | list[files] | 1 < length 10 | A list of supportive attachment | 

**Special Note**:
1. We use rich text in the frontend to collect the contents user submit
2. Content and attachments will be submitted at the sametime when user click submit.
3. Backend needs to store the two type of infos for a replay when user start a chat with customer service

 ### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "ticket submitted",  // 提示信息
    "data": null
}
```

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 5 Set Payment Method Interface

### 5.1 Interface Description    

- Set up the payment method one used to receive transaction from guest. Could be null if one just wanted to be a guest

### 5.2 Address  

`{apiAddress}/api/payment/set-payment-method`

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
| uid   | Yes | string | 1 < length < 20 | guest id to find cart |
| venmo_account   | Yes | string | 1 < length < 20 | venmo account to receive trasctiion |
| zelle_account   | Yes | string | 1 < length < 20 | zelle account to receive trasctiion |
| paypal_account | Yes | string | 1 < length < 20 | paypal account to receive trasactions |

 ### 5.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Update payment method succeed",  // 提示信息
    "data": null
}
```

### 5.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  