## Repositories
---

データベースとのやり取りを担当する。Repository Port（インターフェース）の実装。
Use Case レイヤーから連携された Entity を ORM モデルに変換し、データベースとのやり取りを行う。反対に、Use Case レイヤーへの返却時は、ORM モデルを Entity に変換して返却する。


## Repositories Mapper
---

ORM モデルと Entity の変換を行う。
データベースエラーからアプリケーションエラーへの変換も行う。（詳細は Error Strategy を参照）
