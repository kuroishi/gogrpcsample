gogrpcsample
===============

referneces
----------
作ってわかる！ はじめてのgRPC
https://zenn.dev/hsaki/books/golang-grpc-starting


grpcurl
--------

サーバー内に実装されているサービス一覧の確認

% grpcurl -plaintext localhost:8080 list

あるサービスのメソッド一覧の確認

% grpcurl -plaintext localhost:8080 list myapp.GreetingService

実際の呼出し

$ grpcurl -plaintext -d '{"name": "foobarbaz"}' localhost:8080 myapp.GreetingService.Hello
