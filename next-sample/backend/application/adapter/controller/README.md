## Contr0ller
---

Web API とのやり取りを担当する。
リクエストデータを Use Case で扱うフォーマットに変換の上、Use Case レイヤーに連携する。
レスポンスでは、Use Case レイヤーの出力データを Web API フォーマットに変換して返却する。
Use Case レイヤーとの連携は、Use Case Port （インターフェース）を介して行う。

## Controller Mapper
---

Web API フォーマットと Use Case フォーマットの変換を行う。
アプリケーションエラーから HTTP エラーへの変換も行う。（詳細は Error Strategy を参照）