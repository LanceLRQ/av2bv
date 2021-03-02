# Bilibili video id converter

将B站的AV号转换成BV号
Convert bilibili video id from AV to BV

## Install
```shell
$ go get github.com/LanceLRQ/av2bv
```

Go version >= 1.11 and go module on.

## Usage

Encode （AV => BV)

```go
av2bv.Encode(170001)        // Avid is a uint64 number.
```

```go
bvid, err := av2bv.EncodeString("av170001")
if err != nil { 
    // Error
}
fmt.Println(bvid)
// BV17x411w7KC
```

Decode (BV => AV)

```go
av2bv.Decode([]byte("17x411w7KC"))        // Bvid is a byte array
```

```go
avid, err := av2bv.DecodeString("BV17x411w7KC")
if err != nil { 
    // Error
}
fmt.Println(avid)
// av170001
```

## Algorithm & Thanks

https://www.zhihu.com/question/381784377/answer/1099438784

## LICENSE

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE **(WTFPL)**