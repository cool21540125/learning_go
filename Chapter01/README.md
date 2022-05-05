
# Note

- pointer 唯一的用途, 就是用來取得 value (省記憶體)
    - `*v` 拿到 v 地址的值
    - `&v` 拿到 v 值的地址
    - [取得指標](./Exercise01.13/main.go)
    - [從指標取得值](./Exercise01.14/main.go)
    - [func 裡頭使用指標](./Exercise01.15/main.go)
    - [func 裡頭使用指標](./Activity01.02/main.go)
- enums 是一種用來定義 一系列常數 的方式 (通常是整數) 
    - 通常與 `iota` 一起使用

## template language (格式化樣板語言)

Sign  | Note
----- | ----------------
%v    | 任何值 (如果不在意 the type of the value, 就用這個)
%+v   | 印出值 && 額外資訊. ex: 結構型別的欄位名稱
%#v   | 等同於  %+v 加上型別名稱 
%T    | 印出 the type of the value
%t    | 印出 bool
%s    | 字串
%%    | 印出百分比
%f    | 浮點數
%e    | 帶有科學符號的浮點數
%b    | 2進位 
%d    | 10進位
%x    | 16進位


## zero values

zero value | type
---------- | ----------------------------------------------
false      | bool
0          | number
""         | string
nil        | pointer, func, interface, slice, channel, map