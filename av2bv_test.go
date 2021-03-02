package av2bv

import (
    "math/rand"
    "testing"
    "time"
)

func TestAv170001 (t *testing.T) {
    if string(Encode(170001)) != "17x411w7KC" {
        t.Errorf("encode error")
    }
}

func TestAv455017605 (t *testing.T) {
    if string(Encode(455017605)) != "1Q541167Qg" {
        t.Errorf("encode error")
    }
}

func TestAv882584971 (t *testing.T) {
    if string(Encode(882584971)) != "1mK4y1C7Bz" {
        t.Errorf("encode error")
    }
}

func TestBV17x411w7KC (t *testing.T) {
    if Decode([]byte("17x411w7KC")) != uint64(170001) {
        t.Errorf("decode error")
    }
}

func TestBV1Q541167Qg (t *testing.T) {
    if Decode([]byte("1Q541167Qg")) != uint64(455017605) {
        t.Errorf("decode error")
    }
}

func TestBV1mK4y1C7Bz(t *testing.T) {
    if Decode([]byte("1mK4y1C7Bz")) != uint64(882584971) {
        t.Errorf("decode error")
    }
}

func TestAv2BvString (t *testing.T) {
    currectAvid := "av170001"
    currectBvid := "BV17x411w7KC"
    bvid, err := EncodeString(currectAvid)
    if err != nil {
        t.Error(err)
        return
    }
    if bvid != currectBvid {
        t.Error("bvid error")
        return
    }
    avid, err := DecodeString(bvid)
    if err != nil {
        t.Error(err)
        return
    }
    if currectAvid != avid {
        t.Errorf("decode error")
        return
    }
}

func TestBv2AvString (t *testing.T) {
    currectAvid := "av414437047"
    currectBvid := "BV15V411S7Fq"
    avid, err := DecodeString(currectBvid)
    if err != nil {
        t.Error(err)
        return
    }
    if currectAvid != avid {
        t.Errorf("decode error")
        return
    }
}

func BenchmarkAv2Bv (b *testing.B) {
    rand.Seed(time.Now().UnixNano())
    b.StopTimer()
    b.ResetTimer()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        avid := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
        _ = Encode(avid)
    }
}

func BenchmarkBv2Av (b *testing.B) {
    s = []int{ 9,8,1,6,2,4 }
    tableBytes = []byte{
        'f','Z','o','d','R','9','X','Q','D','S','U','m','2','1','y','C','k','r','6', 'z',
        'B','q','i','v','e','Y','a','h','8','b','t','4','x','s','W','p','H','n', 'J','E',
        '7','j','L','5','V','G','3','g','u','M','T','K','N','P','A','w','c', 'F',
    }
    rand.Seed(time.Now().UnixNano())
    b.StopTimer()
    b.ResetTimer()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        bvid := []byte {'1', 0, 0, '4', 0, '1' ,0 ,'7', 0, 0 }
        for j := 0; j < 6; j++ {
            bvid[s[j]] = tableBytes[rand.Intn(58)]
        }
        _ = Decode(bvid)
    }
}
