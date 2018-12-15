# aws localstack sample

## setup

```console
$ aws configure --profile localstack                                                                                                                          [~/src/github.com/wasanx25/bass]
AWS Access Key ID [None]: dummy
AWS Secret Access Key [None]: dummy
Default region name [None]: ap-northeast-1
Default output format [None]: json
$ docker-compose up -d
$ aws sqs create-queue \
    --queue-name 'TEST' \
    --region ap-northeast-1 \
    --endpoint-url http://localhost:4576 \
    --profile localstack
```

### sample

```
$ aws sqs send-message \
    --queue-url http://localhost:4576/queue/TEST \
    --message-body 'TESTtest' \
    --endpoint-url http://localhost:4576 \
    --profile localstack
    
$ aws  sqs receive-message \
    --queue-url http://localhost:4576/queue/TEST \
    --endpoint-url http://localhost:4576 \
    --profile localstack
```