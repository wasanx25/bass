# aws localstack sample

## setup

```console
$ aws configure --profile localstack                                                                                                                          [~/src/github.com/wasanx25/bass]
AWS Access Key ID [None]: dummy
AWS Secret Access Key [None]: dummy
Default region name [None]: ap-northeast-1
Default output format [None]: text
$ docker-compose up -d
$ aws sqs create-queue \
    --queue-name 'TEST' \
    --region ap-northeast-1 \
    --endpoint-url http://localhost:4579 \
    --profile localstack
```