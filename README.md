# data-interface-for-salesforce-price-bulk-get

## 概要
data-interface-for-salesforce-price-bulk-get は、salesforce の価格オブジェクト取得に必要なデータの整形、および作成時に salesforce から返ってきた response の MySQL への格納をバルク単位で行うマイクロサービスです。

## 動作環境
data-interface-for-salesforce-price-bulk-get は、aion-coreのプラットフォーム上での動作を前提としています。  
使用する際は、事前に下記の通りAIONの動作環境を用意してください。     

・OS: Linux OS   
・CPU: ARM/AMD/Intel   
・Kubernetes   
・AION のリソース   

## セットアップ
以下のコマンドを実行して、docker imageを作成してください。
```
$ cd /path/to/data-interface-for-salesforce-price-bulk-get
$ make docker-build
```

## 起動方法
以下のコマンドを実行して、podを立ち上げてください。
```
$ cd /path/to/data-interface-for-salesforce-price-bulk-get
$ kubectl apply -f data-interface-for-salesforce-price-bulk-get.yaml
```


## kanban との通信
### kanban(ui-backend) から受信するデータ
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| method | get |
| object | PriceRelatedList |
| connection_type | request |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "PriceRelatedList"
"connection_type": "request"
```

### kanban に送信するデータ
kanban に送信する metadata は下記の情報を含みます。

| key | type | description |
| --- | --- | --- |
| method | string | get |
| object | string | PriceMasterList/PriceRecordList |
| connection_key | string | price_bulk_get |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "PriceMasterList"
"connection_key": "price_bulk_get"
```

## kanban(salesforce-api-kube) から受信するデータ
kanban からの受信可能データは下記の形式です

### PriceMaster/PriceRecord
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| key | 文字列 "PriceMasterList"/"PriceRecordList" を指定 |
| content | PriceMaster/PriceRecord の詳細情報を含む JSON 配列 |
| connection_type | response |

具体例:
```example
# metadata (map[string]interface{}) の中身

"key": "PriceMasterList"
"content": "[{xxxxxxxxxxx}]"
"connection_type": "response"
```
