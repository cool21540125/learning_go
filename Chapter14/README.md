
# Note

- 發送 http.Request 以後, 如果發生錯誤(即使 nil 沒東西), 應檢查 `resp.StatusCode` 是否為 2xx 或 3xx, 才能確保請求成功
- 發送完請求, 務必釋放 http.Request 佔用的連線資源 `defer resp.Body.Close()`
- Client 上傳 MIME 的實作範例 [MIME Server](./Exercise14.04/server/server.go) 及 [MIME Client](./Exercise14.04/client/main.go)
- HTTP Custom Header 的範例 [Client Get Custom Header](./Exercise14.05/client/main.go) 及 [Server Set Custom Header](./Exercise14.05/server/server.go)

# Signature

```go
// http.Get
func Get(url string) (resp *Response, err error)

// http.Post
func Post(url, contentType string, body io.Reader) (resp *Response, err error)

// http.NewRequest - 
func NewRequest(method, url string, body io.Reader) (*Request, error)
func (c *Client) Do(req *Request) (*Response, error)

// 
```