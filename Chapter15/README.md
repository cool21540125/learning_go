
# Note

- http Server 實作時有 2 個容易搞混的地方:
    - `http.Handle(param1, param2)`
        - param1 為 請求的 location
        - param2 為 實作了 *ServeHTTP(w http.ResponseWriter, r *http.Request)* 方法的結構
        - [運行一個最簡單範例的 http Server](./Exercise15.01/main.go)
    - `http.HandleFunc(param1, param2)`
        - 此 func 參數為 *(string, func(w http.ResponseWriter, r *http.Request))*
            - param1 為 請求的 location
            - param2 也可以是 anonymous func
        - 但需要注意的是, 這個會把路由設定到 http 套件的 **DefaultServeMux** (這東西就是 http 預設的 ServeMux 結構)
            - [須留意, 明明訪問不同 location, 但結果卻不符合直覺](./Exercise15.02/main.go)
- [Server 取得 Query String](./Exercise15.03/main.go)
- [塞資料到 html/template (template 寫死)](./Exercise15.04/main.go)
    - 由於上述的方式太蠢了(template 字串方式寫死), 請使用下列方式:
        - [指定檔案路徑, 提供單一靜態資源](./Exercise15.05/main.go)
        - [指定路徑, 提供多個靜態資源](./Exercise15.06/main.go)
        - [指定 template 檔案路徑, 產生動態網頁 的方式](./Example15.01/main.go)
    - [接收表單 Post 請求](./Exercise15.07/main.go)
- [Restful API Server 回傳 Server 時間的範例](./Exercise15.08/main.go)
- 關於本章節的一些額外參考:
    - [golang http StatusCode](https://pkg.go.dev/net/http#pkg-constants)
    - [http/template 語法控制](https://pkg.go.dev/text/template#hdr-Actions)
