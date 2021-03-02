package av2bv

import (
    "fmt"
    "strconv"
    "strings"
)

var (
	tableBytes = []byte{
        'f','Z','o','d','R','9','X','Q','D','S','U','m','2','1','y','C','k','r','6', 'z',
        'B','q','i','v','e','Y','a','h','8','b','t','4','x','s','W','p','H','n', 'J','E',
        '7','j','L','5','V','G','3','g','u','M','T','K','N','P','A','w','c', 'F',
    }
    tableReverse = []uint8 {
        0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,
        0,0,0,0,0,0,0,0,13,12,46,31,43,18,40,28,5,0,0,0,0,0,0,0,54,20,15,8,39,57,45,36,0,
        38,51,42,49,52,0,53,7,4,9,50,10,44,34,6,25,1,0,0,0,0,0,0,26,29,56,3,24,0,47,27,22,
        41,16,0,11,37,2,35,21,17,33,30,48,23,55,32,14,19,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
        0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,     // (8 bit)
    }
    s = []int{ 9,8,1,6,2,4 }
    sPow = []uint64 { 1,58,3364,195112,11316496,656356768 }
)

const (
    xor uint64 = 177451812
    add uint64 = 8728348608
    base uint64 = 58
)

// Convert AvId number to BvId Bytes
// 将整数的AvId翻译成BvId字节码
func Encode(avid uint64) []byte {
    x := (avid ^ xor) + add
    r := []byte {'1', 0, 0, '4', 0, '1' ,0 ,'7', 0, 0 }
    for i := 0; i < 6; i++ {
        r[s[i]] = tableBytes[(x / sPow[i]) % base]
    }
    return r
}

// Convert AvId string to BvId string
// 解析AvId字符串并转换成BvId字符串
func EncodeString(avid string) (string, error) {
    if avid == "" {
        return "", fmt.Errorf("avid is empty")
    }
    if strings.ToLower(avid[:2]) == "av" {
        // if avid start with "av"
        avid = avid[2:]
    } else if avid[0] < 0x30 || avid[0] > 0x39 {
        return "", fmt.Errorf("avid syntax error")
    }
    avNum, err := strconv.ParseUint(avid, 10, 64)
    if err != nil {
        return "", fmt.Errorf("avid syntax error")
    }
    return fmt.Sprintf("BV%s", string(Encode(avNum))), nil
}

// Convert BvId bytes to AvId number
// 将BvId字节码翻译成整数的AvId
func Decode(bvid []byte) uint64 {
    r := uint64(0)
    for i := 0; i < 6; i++ {
        r += uint64(tableReverse[bvid[s[i]]]) * sPow[i]
    }
    return (r - add) ^ xor
}

// Convert BvId string to AvId string
// 解析BvId字符串并转换成AvId字符串
func DecodeString (bvid string) (string, error) {
    if bvid == "" {
        return "", fmt.Errorf("bvid is empty")
    }
    if len(bvid) <= 2 {
        return "", fmt.Errorf("avid syntax error")
    }
    if strings.ToLower(bvid[:2]) == "bv" {
        // if bvid start with "bv"
        bvid = bvid[2:]
    }
    if len(bvid) != 10 {
        return "", fmt.Errorf("avid syntax error")
    }
    avid := Decode([]byte(bvid))
    return fmt.Sprintf("av%d", avid), nil
}