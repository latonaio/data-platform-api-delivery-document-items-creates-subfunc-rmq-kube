# data-platform-api-delivery-document-items-creates-subfunc-rmq-kube
data-platform-api-delivery-document-items-creates-subfunc-rmq-kube は、データ連携基盤において、入出荷APIサービスのヘッダ登録補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-delivery-document-items-creates-subfunc-rmq-kube の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_DELIVERY_DOCUMENT_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-delivery-document-items-creates-subfunc-rmq-kube が対応する データ は、次のものです。

* DeliveryDocumentHeader（入出荷 - ヘッダデータ）
* DeliveryDocumentHeaderPartner（入出荷 - ヘッダ取引先データ）
* DeliveryDocumentHeaderPartnerPlant（入出荷 - ヘッダ取引先プラントデータ）
* DeliveryDocumentHeaderPartnerContact（入出荷 - ヘッダ取引先コンタクトデータ）

## Output
data-platform-api-delivery-document-items-creates-subfunc-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
{
  "connection_key": "request",
  "result": true,
  "redis_key": "abcdefg",
  "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
  "api_status_code": 200,
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "business_partner": 201,
  "service_label": "DELIVERY_DOCUMENT",
  "message": {
    "Header": [
      {
        "DeliveryDocument": 80000006,
        "Buyer": 101,
        "Seller": 201,
        "ReferenceDocument": 100,
        "ReferenceDocumentItem": null,
        "OrderID": 100,
        "OrderItem": null,
        "ContractType": null,
        "OrderVaridityStartDate": null,
        "OrderValidityEndDate": null,
        "InvoiceScheduleStartDate": null,
        "InvoiceScheduleEndDate": null,
        "IssuingPlantTimeZone": "",
        "ReceivingPlantTimeZone": "",
        "DocumentDate": "2022-11-25",
        "PlannedGoodsIssueDate": "",
        "PlannedGoodsIssueTime": "",
        "PlannedGoodsReceiptDate": "",
        "PlannedGoodsReceiptTime": "",
        "BillingDocumentDate": "",
        "CompleteDeliveryIsDefined": null,
        "OverallDeliveryStatus": "",
        "CreationDate": "",
        "CreationTime": "",
        "IssuingBlockReason": null,
        "ReceivingBlockReason": null,
        "GoodsIssueOrReceiptSlipNumber": "",
        "HeaderBillingStatus": "",
        "HeaderBillingConfStatus": "",
        "HeaderBillingBlockReason": null,
        "HeaderGrossWeight": null,
        "HeaderNetWeight": null,
        "HeaderVolume": null,
        "HeaderVolumeUnit": "",
        "HeaderWeightUnit": "",
        "Incoterms": "CIF",
        "IsExportImportDelivery": null,
        "LastChangeDate": "",
        "IssuingPlantBusinessPartner": "",
        "IssuingPlant": "",
        "ReceivingPlant": "",
        "ReceivingPlantBusinessPartner": "",
        "DeliverToParty": null,
        "DeliverFromParty": null,
        "TransactionCurrency": "JPY",
        "OverallDelivReltdBillgStatus": ""
      }
    ],
    "HeaderPartner": null,
    "HeaderPartnerPlant": null
  },
  "api_schema": "DPFMDeliveryDocumentCreates",
  "accepter": [
    "All"
  ],
  "deleted": false,
  "sql_update_result": null,
  "sql_update_error": "",
  "subfunc_result": true,
  "subfunc_error": "",
  "exconf_result": null,
  "exconf_error": "",
  "api_processing_result": null,
  "api_processing_error": ""
}
```