package simhash

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strconv"
)

type Vector [64]int64

type Feature struct {
	value  []byte
	weight uint64
}

type Simhash struct {
	features []Feature
}

func NewSimhash() *Simhash {
	return &Simhash{
		features: make([]Feature, 0),
	}
}

func toBytes(value any) ([]byte, error) {
	switch v := value.(type) {
	case string:
		// If it's a string
		return []byte(v), nil
	case int:
		// Convert int to int64 to ensure a fixed size
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, int64(v))
		if err != nil {
			return nil, fmt.Errorf("failed to convert number to bytes: %v", err)
		}

		return buf.Bytes(), nil
	case int32, int64, uint, uint32, uint64, float32, float64:
		// If it's a number
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert number to bytes: %v", err)
		}

		return buf.Bytes(), nil
	case []byte:
		// If it's already a byte slice
		return v, nil
	default:
		// For any other type, attempt JSON serialization
		bytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to bytes: %v", err)
		}

		return bytes, nil
	}
}

func (s *Simhash) AddFeature(value any, weight uint64) error {
	byteValue, err := toBytes(value)
	if err != nil {
		return err
	}

	s.features = append(s.features, Feature{value: byteValue, weight: weight})

	return nil
}

func (f *Feature) vectorize() Vector {
	var v Vector

	h := fnv.New64a()
	h.Reset()
	h.Write(f.value)
	sum := h.Sum64()

	for i := uint8(0); i < 64; i++ {
		bit := (sum >> i) & 1
		if bit == 1 {
			v[i] += int64(f.weight)
		} else {
			v[i] -= int64(f.weight)
		}
	}

	return v
}

func uint64ToHex64(uintValue uint64) string {
	hexStr := fmt.Sprintf("%X", uintValue)
	return hexStr
}

func hexToBinaryUint64(hexStr string) uint64 {
	uintValue, _ := strconv.ParseUint(hexStr, 16, 64)
	return uintValue
}

func HammingDistance(a, b uint64) int {
	v := a ^ b
	var c int
	for v != 0 {
		v &= v - 1
		c++
	}
	return c
}

func (s *Simhash) GenerateToken() string {
	var v, total Vector
	for _, f := range s.features {
		v = f.vectorize()
		for i := 0; i < 64; i++ {
			total[i] += v[i]
		}
	}

	var token uint64
	for i := 0; i < 64; i++ {
		if total[i] >= 0 {
			token |= 1 << i
		}
	}

	return uint64ToHex64(token)
}

func ComputeSimilarity(token1, token2 string) float64 {
	// Convert hex tokens to uint64
	hash1 := hexToBinaryUint64(token1)
	hash2 := hexToBinaryUint64(token2)

	// Compute Hamming distance between the two tokens
	dist := HammingDistance(hash1, hash2)

	// Calculate similarity (1 - normalized Hamming distance)
	similarity := 1 - float64(dist)/64.0

	return float64(int64(similarity*10000)) / 100.0
}
