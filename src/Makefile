.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./dist/choose_stores ./dist/create_store ./dist/list_stores
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/choose_stores ./functions/choose_stores
	GOOS=linux GOARCH=amd64 go build -o ./dist/create_store ./functions/create_store
	GOOS=linux GOARCH=amd64 go build -o ./dist/list_stores ./functions/list_stores

package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket random-lunch --profile // 実際にはここに profile を書く

deploy:
	sam deploy --template-file output-template.yaml --stack-name random-lunch --capabilities CAPABILITY_IAM --profile // 実際にはここに profile を書く
