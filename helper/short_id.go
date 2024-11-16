package helper

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/sqids/sqids-go"
)

func GeneratePrefixedID(prefix string, minLength int) (string, error) {
	if minLength < len(prefix) {
		return "", fmt.Errorf("minLength (%d) must be greater than prefix length (%d)", minLength, len(prefix))
	}

	options := sqids.Options{
		Alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
	sqids, err := sqids.New(options)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Sqids: %w", err)
	}

	var seeds []uint64
	for {
		seed, err := generateSecureRandomSeed()
		if err != nil {
			return "", fmt.Errorf("failed to generate random seed: %w", err)
		}
		seeds = append(seeds, seed)

		encoded, err := sqids.Encode(seeds)
		if err != nil {
			return "", fmt.Errorf("failed to encode seeds: %w", err)
		}

		if len(prefix)+len(encoded) >= minLength {
			return fmt.Sprintf("%s%s", prefix, encoded), nil
		}
	}
}

func generateSecureRandomSeed() (uint64, error) {
	var seed uint64
	err := binary.Read(rand.Reader, binary.BigEndian, &seed)
	if err != nil {
		return 0, err
	}
	return seed, nil
}
