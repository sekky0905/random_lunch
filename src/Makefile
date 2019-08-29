.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./dist/choose_shops ./dist/create_shop ./dist/list_shops
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/choose_shops ./functions/choose_shops
	GOOS=linux GOARCH=amd64 go build -o ./dist/create_shop ./functions/create_shop
	GOOS=linux GOARCH=amd64 go build -o ./dist/list_shops ./functions/list_shops

package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket random-lunch --profile // 実際にはここに profile を書く

deploy:
	sam deploy --template-file output-template.yaml --stack-name random-lunch --capabilities CAPABILITY_IAM --profile // 実際にはここに profile を書く
