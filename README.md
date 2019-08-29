# random_lunch

お店を自動で選んでくれる Slack コマンド


## アーキテクチャ

- なんちゃってレイヤードアーキテクチャ
- 簡素化のためにDI等は行わない

```bash
├── src                
   ├── functions      <-- controler相当
   └── applications    
   └── infra           
   | └── repository    
   └── models         
```

## Makefile

- ビルド
    - `make build`
- デプロイ用にパッケージ化
    - `make package`
- デプロイ
    - `make deploy`  
    

- 実際に使用するときには、以下の箇所に実際のprofileを入れる
	
```
package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket random-lunch --profile // 実際にはここに profile を書く

deploy:
	sam deploy --template-file output-template.yaml --stack-name random-lunch --capabilities CAPABILITY_IAM --profile // 実際にはここに profile を書く
```

## Slack 上でのコマンドの扱い

- お店を選んでもらう
    - `/choose_lunch_shops`
- お店を新規に登録する
    - `/create_lunch_shops お店の名前 url メモ(任意入力)` 
    - コマンド 、お店の名前 、url 、メモ それぞれの間は半角空文字で区切る
- お店の一覧を取得する
    - `/list_lunch_shops`
 
