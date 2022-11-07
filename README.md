# variable-fixer
## About
- terraformの `variables.tf` に対して下記の値を可能な限り自動的に作成します
  - `description`
  - `type`

## How to use
- 下記のようにコマンドを実行

```bash
./bin/fixer path/to/variables.tf
```

- `variables.tf` が修正され、元の状態が `variables.tf.bak` に退避されるのでdiffを確認
  - `type` が補完できなかったvariableに関しては、標準出力に出てくるので手動で埋める必要がある

    ```
    This valiable could not be filled type value : xxxxx
    ```

- `terraform plan` で想定通りの結果になることを確認
- `variables.tf.bak` を削除してcommit

## How to update
- コンパイル言語なので、コードを修正したらbuildが必要

```bash
# version確認
go version
go version go1.18.4 darwin/arm64

# buildして成果物を更新
go build -o bin/fixer .
```
