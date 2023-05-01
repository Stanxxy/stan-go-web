# RESTful Util Module API document  

## 1 Get State Interface 

### 1.1 Interface Description    

- Get state from database

### 1.2 Address  

`{apiAddress}/api/utils/get-state`  

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

**Special Note**: 
1. at the moment we only support states in the US. In the future we may consider support more countries. 

### 1.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Return available states",  // 提示信息
    "data": [
      "New York",
      "New Jersey",
      "Texas"
    ]  // 返回内容
}
```

### 1.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 2 Get City Interface 

### 2.1 Interface Description    

- Get City from database

### 2.2 Address  

`{apiAddress}/api/utils/get-cities`  

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
| state   | Yes   | string | 1 < length < 50 | name of the state |

**Special Note**: 
1. at the moment we only support states in the US. In the future we may consider support more countries. 

### 2.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Return available states",  // 提示信息
    "data": [
      "New York",
      "Long Island",
      "Rochester",
      "Albaney"
    ]  // 返回内容
}
```

### 2.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  


## 3 Get Checkcode

### 3.1 Interface Description    

- Get Checkcode to prevent bot from register。
- Could be outsourced to google api. (To be investigated)

### 3.2 Address  

`{apiAddress}/api/utils/get-cities`  

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
| state   | Yes   | string | 1 < length < 50 | name of the state |

**Special Note**: 
1. at the moment we only support states in the US. In the future we may consider support more countries. 

### 3.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Return available states",  // 提示信息
    "data": [
      "New York",
      "Long Island",
      "Rochester",
      "Albaney"
    ]  // 返回内容
}
```

### 3.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  



## 4 Get Waiver

### 4.1 Interface Description    

- Get City from database

### 4.2 Address  

`{apiAddress}/api/utils/get-cities`  

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
| state   | Yes   | string | 1 < length < 50 | name of the state |

**Special Note**: 
1. at the moment we only support states in the US. In the future we may consider support more countries. 

### 4.5 Sample Response

```json
{
    "code": 200,  // 状态码
    "msg": "Return available states",  // 提示信息
    "data": [
      "New York",
      "Long Island",
      "Rochester",
      "Albaney"
    ]  // 返回内容
}
```

### 4.6 Miscellaneous  

For more return state please check return state table  

[Return State Table](URL/for/api/responseCode/table)  