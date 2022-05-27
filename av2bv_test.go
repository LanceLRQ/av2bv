package av2bv

import (
	"math/rand"
	"testing"
	"time"
)

func TestAv(t *testing.T) {
	bvid, err := EncodeString("170001")
	if err != nil {
		t.Fatal(err)
	}
	if bvid != "BV17x411w7KC" {
		t.Fatalf("encode 170001 error, bvid: %s", bvid)
	}
	bvid, err = EncodeString("AV455017605")
	if err != nil {
		t.Fatal(err)
	}
	if bvid != "BV1Q541167Qg" {
		t.Fatalf("encode 455017605 error, bvid: %s", bvid)
	}
	bvid, err = EncodeString("av882584971")
	if err != nil {
		t.Fatal(err)
	}
	if bvid != "BV1mK4y1C7Bz" {
		t.Fatalf("encode 882584971 error, bvid: %s", bvid)
	}
}

func TestBV(t *testing.T) {
	avid, err := DecodeString("17x411w7KC")
	if err != nil {
		t.Fatal(err)
	}
	if avid != "av170001" {
		t.Fatalf("decode 17x411w7KC error, avid: %s", avid)
	}
	avid, err = DecodeString("BV1Q541167Qg")
	if err != nil {
		t.Fatal(err)
	}
	if avid != "av455017605" {
		t.Fatalf("decode 1Q541167Qg error, avid: %s", avid)
	}
	avid, err = DecodeString("bv1mK4y1C7Bz")
	if err != nil {
		t.Fatal(err)
	}
	if avid != "av882584971" {
		t.Fatalf("decode 1mK4y1C7Bz error, avid: %s", avid)
	}
}

func TestAvBv(t *testing.T) {
	for i := 0; i < 1048576; i++ {
		avid := uint64(rand.Int63n(1<<30 - 1))
		bvid := Encode(avid)
		avidd := Decode(&bvid)
		if avid != avidd {
			t.Fatalf("round %d, avid to encode: %x, avid decoded: %x, bvid: %s", i, avid, avidd, bvid)
		}
	}
}

func BenchmarkAv2Bv(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		avid := uint64(rand.Int63n(1<<30 - 1))
		_ = Encode(avid)
	}
}

func BenchmarkBv2Av(b *testing.B) {
	s := []int{9, 8, 1, 6, 2, 4}
	tableBytes := []byte{
		'f', 'Z', 'o', 'd', 'R', '9', 'X', 'Q', 'D', 'S', 'U', 'm', '2', '1', 'y', 'C', 'k', 'r', '6', 'z',
		'B', 'q', 'i', 'v', 'e', 'Y', 'a', 'h', '8', 'b', 't', '4', 'x', 's', 'W', 'p', 'H', 'n', 'J', 'E',
		'7', 'j', 'L', '5', 'V', 'G', '3', 'g', 'u', 'M', 'T', 'K', 'N', 'P', 'A', 'w', 'c', 'F',
	}
	rand.Seed(time.Now().UnixNano())
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bvid := BVID{'1', 0, 0, '4', 0, '1', 0, '7', 0, 0}
		for j := 0; j < 6; j++ {
			bvid[s[j]] = tableBytes[rand.Intn(58)]
		}
		_ = Decode(&bvid)
	}
}
