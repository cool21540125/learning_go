
- Marshal 與 UnMarshal
    - 要從 json 的角度來看, Marshal 是組織成 json, UnMarshal 是拆解 json(變成 golang struct)
    - [json Unmarshal -> struct 的基本範例](./Example11.03/main.go)
    - [使用 MarshalIndent 來漂亮排版](./Example11.06/main.go)
- *Marshal 與 Encode* 對應 *UnMarshal 與 Decode*, 常被拿出來討論
    - Marshal => string 與此相對 Encode => stream
    - json bytes 拆解成 struct, 可看[這個 Decode 用法](./Example11.07/main.go)
    - 可參考[這篇](https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling)
- golang 專用的 json 編碼: `encoding/gob`
    - [struct => bytes 範例](./Exercise11.04/main.go)
