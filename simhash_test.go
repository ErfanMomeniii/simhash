package simhash

import (
	"testing"
)

func TestSimhash_AddFeature(t *testing.T) {
	s := NewSimhash()

	err := s.AddFeature("test", 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(s.features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(s.features))
	}

	if string(s.features[0].value) != "test" {
		t.Fatalf("expected feature value 'test', got %s", string(s.features[0].value))
	}

	if s.features[0].weight != 1 {
		t.Fatalf("expected feature weight 1, got %d", s.features[0].weight)
	}
}

func TestSimhash_GenerateToken(t *testing.T) {
	s := NewSimhash()
	s.AddFeature("test", 1)
	s.AddFeature("example", 2)

	token := s.GenerateToken()

	if len(token) == 0 {
		t.Fatal("expected a non-empty token")
	}
}

func TestHammingDistance(t *testing.T) {
	a := uint64(0b101010)
	b := uint64(0b111000)

	dist := HammingDistance(a, b)

	expected := 2
	if dist != expected {
		t.Fatalf("expected Hamming distance %d, got %d", expected, dist)
	}
}

func TestComputeSimilarityWithLowDifference(t *testing.T) {
	s1 := NewSimhash()
	s1.AddFeature("hello", 5)
	s1.AddFeature("world", 1)

	s2 := NewSimhash()
	s2.AddFeature("hello", 5)
	s2.AddFeature("golang", 1)

	token1 := s1.GenerateToken()
	token2 := s2.GenerateToken()

	similarity := ComputeSimilarity(token1, token2)

	expectedSimilarity := 90.0

	if similarity < expectedSimilarity {
		t.Fatalf("expected a similarity score %f", similarity)
	}

	if similarity > 100.0 || similarity < 0.00 {
		t.Fatalf("expected similarity between 0 and 100, got %f", similarity)
	}
}

func TestComputeSimilarityWithHighDifference(t *testing.T) {
	s1 := NewSimhash()
	s1.AddFeature("hello", 1)
	s1.AddFeature("world", 5)

	s2 := NewSimhash()
	s2.AddFeature("hello", 1)
	s2.AddFeature("golang", 5)

	token1 := s1.GenerateToken()
	token2 := s2.GenerateToken()

	similarity := ComputeSimilarity(token1, token2)

	expectedSimilarity := 60.0

	if similarity > expectedSimilarity {
		t.Fatalf("expected a similarity score %f", similarity)
	}

	if similarity > 100.0 || similarity < 0.00 {
		t.Fatalf("expected similarity between 0 and 100, got %f", similarity)
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		input    any
		expected []byte
	}{
		{"test", []byte("test")},
		{42, []byte{42, 0, 0, 0, 0, 0, 0, 0}},
		{42.2, []byte{154, 153, 153, 153, 153, 25, 69, 64}},
	}

	for _, test := range tests {
		output, err := toBytes(test.input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if string(output) != string(test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
