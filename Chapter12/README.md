
# Note

- flag
    - [flag 的基本用法](./Example12.01/main.go)
        - 運行方式: `go run . --value 10`
    - [印出所有 flag 說明](./Example12.02/main.go)
- signal
    - Signal 是指 OS 發送給 process 的非同步通知. process 收到時, 應設法處理它
        - [signal 範例](./Exercise12.01/main.go)
    - Signal
        - `syscall.SIGINT` 為 `kill -15`
        - `syscall.SIGTERM` 為 `kill -9`
- os file
    - 檔案操作, 記得都需要 `defer f.Close()`
    - [建立檔案](./Example12.03/main.go) (已存在則覆蓋)
    - [開啟檔案](./Example12.07/main.go)
    - [刪除檔案](./Practice12.01/main.go)
    - 此方式可直接 [寫入到要建立的檔案](./Example12.04/main.go)
        - golang 1.15 以前, 需使用 `ioutil.WriteFile()`
    - [取得檔案摘要 & 判斷檔案是否存在](./Example12.05/main.go)
    - 一次讀取整個檔案
        - [os.ReadFile(string)](./Example12.06/main.go)
        - [os.ReadAll(io.Reader)](./Example12.07/main.go)
    - [較強大的開啟建立檔案的功能](./Exercise12.02/main.go)
    - [檔案備份](./Exercise12.02/main.go)
    - [用 log 套件將日誌寫入到檔案](./Practice12.02/main.go)


## Signature

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

`OpenFile` 的 flat 如下(使用 「|」串聯)

- syscall.O_RDONLY : r
- syscall.O_WRONLY : w
- syscall.O_RDWR   : rw
- syscall.O_APPEND : append to file
- syscall.O_CREATE : create if not exist
- syscall.O_EXCL   : 搭配 O_CREATE, 確保檔案不存在
- syscall.O_SYNC   : I/O 同步模式 (等待儲存裝置寫入完成)
- syscall.O_TRUNC  : 開啟檔案時清空內容
