## Use Case
---

ビジネスロジック（ソフトウェアで何が出来るかの表現）を担当する。Use Case Port（インターフェース）の実装。
Use Case フォーマットから Entity への変換を行い、Repository に渡す。Repository との連携は、Repository Port を介して行う。

Use Case フォーマットと Entity の変換は、Use Case Mapper を参照する


## Use Case Mapper
---

Use Case フォーマット と Entity の変換を行う。
依存方向を下位レイヤー → 上位レイヤーに限定する必要があるため、Repository Port は （Gateway 側ではなく）Use Case 側に定義。