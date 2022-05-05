
# Note 

- [查看 runtime 時期的記憶體用量](./Example03.02/main.go)
- 對於 Bytes, Runes, Char, String 及 unicode
    - [逐字列印 string, 應該要 string(rune())](./Example03.09/main.go)
        - [如果只取 string() 會發生出錯(但你可能連怎麼死的都不知道XD)](./Example03.08/main.go)
        - [如果迴圈直接印只會得到 uint8](./Example03.07/main.go).
    - rune 概念範例
        - len("tony周") -> 7          (這是錯誤的)
        - len([]rune("tony周")) -> 5  (這才是正確的)
    - [看這個應該會有感覺](./Example03.10/main.go)
    - [日文字的 unicode 及 string](./Exercise03.05/main.go)
    - [重設密碼時, 字元合理性驗證的範例](./Exercise03.01/main.go)
- 數值型別
    - [int8 與 uint8 的範圍, 示範溢位繞回(wraparound)](./Exercise03.03/main.go)
    - [大數值的型別, 請使用 math](./Exercise03.04/main.go)


# type

## number

type   | Scope                    | Desc
------ | ------------------------ | -------------------------------
uint8  | 0 ~ 255  (1 bytes)       | 
uint16 | 0 ~ 65535                | 
uint32 | 0 ~ 2^32                 | 
uint64 | 0 ~ 2^64                 | 
int8   | -128 ~ 127               | 
int16  | -32768 ~ 32767           | 
int32  | -2^32 ~ 2^32-1           | 
int64  | -2^64 ~ 2^64-1           | 
uint   | uint32 for 32 bits PC    | 
int    | int64 for 64 bits PC     | 
byte   | 等同於 uint8              | 
rune   | 等同於 uint32 (1~4 bytes) | 可完整容納一個 UTF-8 字元 (Unicode編碼)


## string

可用 "string" 或是 `string`(重音符號) 來把字串匡起來

- 「"」裏頭可放置跳脫字元
- 「`」裏頭則代表純文字 (等同於 Bash 的 ')

string literal(字串常值)都是用 UTF-8 編碼

如果文字用 string 型別來儲存時, golang 會以 「byte 集合」來儲存所有字串

說穿了, string 實際上就是 `唯讀的 byte 切片`, a.k.a. 有些 UTF-8 字元, 會被拆開成多個位元組

因此為了安全處理字串, 不管編碼方式採用單一還是多位元組, 最好都把字串從 「byte 集合」轉換成 「rune 集合」
