<p align="center">
<a href="#">
    <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# simhash

**simhash** is a lightweight Go package for generating Simhash tokens and calculating their similarity using the 
[Moses Charikar Simhash](https://en.wikipedia.org/wiki/SimHash) algorithm. It is ideal for applications 
like text deduplication, plagiarism detection, and near-duplicate content detection and fingerprinting.

For detailed usage, check [this](https://pkg.go.dev/github.com/erfanmomeniii/simhash).

---

# Documentation

## Install

To get started with simhash, install it using:

```bash
go get github.com/erfanmomeniii/simhash
```

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/simhash"
```

## Quick Start

The following example demonstrates how to generate Simhash tokens and calculate similarity:

```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/simhash"
)

func main() {
	// Create a new Simhash instance
	s := simhash.NewSimhash()

	// Add features with weights
	s.AddFeature("example", 2)
	s.AddFeature("test", 5)

	// Generate a Simhash token
	token1 := s.GenerateToken()

	// Create another Simhash instance with similar features
	s2 := simhash.NewSimhash()
	s2.AddFeature("example", 2)
	s2.AddFeature("testcase", 5)

	// Generate another token
	token2 := s2.GenerateToken()

	// Compute similarity between the two tokens
	similarity := simhash.ComputeSimilarity(token1, token2)

	fmt.Printf("Token1: %s\nToken2: %s\nSimilarity: %f\n", token1, token2, similarity)
}
```
Output:
```
Token1: F9E6E6EF197C2B25
Token2: FDA981914657B7D1
Similarity: 43.75
```

## Features

### Add Feature

Add features with their weights to the Simhash generator:

```go
s.AddFeature("example", 5)
s.AddFeature(12345, 10)
```
### Generate Token

Generate a 64-bit hexadecimal Simhash token based on the added features:

```go
token := s.GenerateToken()
```
### Compute Similarity

Calculate the similarity between two Simhash tokens as a percentage (normalized Hamming distance):

```go
similarity := simhash.ComputeSimilarity(token1, token2)
```
### Supported Feature Types

The `AddFeature` method accepts the following types:
- Strings: e.g., "example"
- Numbers: e.g., 123, float64, etc.
- Byte slices: e.g., []byte("example")
- Any other type: Converted using JSON serialization

---

## Contributing

Pull requests are welcome! For any changes, please open an issue first to discuss the proposed modification. Ensure tests are updated accordingly.
