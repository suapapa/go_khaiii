# khaiii-api: khaiii API server

![khaiii_logo](_asset/khaiii_logo_256.webp)

## Build

Update docs:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

Build container image:
```sh
make build_image IMAGE_TAG=khaiii-api:latest
```

Run the API server locally:
```sh
docker run -it --rm -p 8082:8080 khaiii-api:latest
```

Client example:
```sh
cd examples/api_client
go run main.go
```

Example output:
```
orig_text: 사랑은 모든것을 덮어주고 모든것을 믿으며 모든것을 바라고 모든것을 견디어냅니다
word_chunks:
- word: 사랑은
  begin: 0
  len: 9
  morphs:
  - lex: 사랑
    tag: NNG
  - lex: 은
    tag: JX
- word: 모든것을
  begin: 10
  len: 12
  morphs:
  - lex: 모든
    tag: MM
  - lex: 것
    tag: NNB
  - lex: 을
    tag: JKO
# ...
```