package service

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_service_SolveChallenge(t *testing.T) {
	tests := []struct {
		name string

		challengeString       string
		solvedChallengeString string
		publicKeyString       string
	}{
		{
			name:                  "1",
			challengeString:       "T2UCDtHvrK3uhoyJhzqZ9Fbr+Q/hGIPV/rmxaofm+SI=",
			solvedChallengeString: "VF05gLWIEFavQ9GYIFsjgb+fMmVRUThqICVbxyEqlt1zxO1ytpbyFL+eiTwpiHMBE3w0QfRj9Bk3r6+8UHMcItgTLP2ML8oGTyW4/7YG347LofqHyhCvFN/y04Y/NIkNSn8xD0AiRmzMi1VS9ZwUWHEWnKLwEIFWsBmIQ+YsxdLEKw6SrEKaCS2vDCwCZ4Yv+mTd2hvul5a7lVBVwPxH/KNgtrnrhFMsgkv97XWTCBjIX8aXE8UP2wfZvS8bk2SA1ZvRnqHgXmZUMt7ae5DkXCENBrQQrDAggsiB6Myh/1gJX8nTWMPODTiZ7RYqTLUbEpUqcy51CTFi3GGNf6OSIQ==",
			publicKeyString: `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmdB5OtML3hnrwVe9jJbdfekrUoF0dia9is6y4CH5fRcq/glxk9kTiVtogYIT49MaqTBPtRbkgzrhM2ViVkzSeGb6quJiRU/voLJTrMUa18PdniGurLcGOEiiNBUo2NMjfqQjY4XfFwahMozORX3DiJ/ayhFJbWPkCl0KyzruqG0VozxQV9eCX7lpdBE+6FvxZuG9by2ZQyTwXR2y8zDPR4Mo0X9n4WlFxdz0wMZ73PV0j3lMW4sfyzxi+p+cGUjT6/uqtFY9dE7BPVH8VLbRYSWCIqd5JH5wuyFk1dsQviVYN1jaEAXqFJYuS7TX8myBUvHCqrHpgDmBlXhHiq1HbQIDAQAB
-----END PUBLIC KEY-----
`,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKeyBlock, _ := pem.Decode([]byte(tt.publicKeyString))
			if publicKeyBlock == nil {
				log.Fatalf("Failed to parse PEM block containing the public key")
			}

			publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
			if err != nil {
				log.Fatalf("Failed to parse public key: %v", err)
			}

			rsaPubKey, ok := publicKey.(*rsa.PublicKey)
			if !ok {
				log.Fatalf("Could not convert to rsa.PublicKey")
			}

			fmt.Printf("\nSuccessfully parsed public key: %v", rsaPubKey)

			challenge, err := base64.StdEncoding.DecodeString(tt.challengeString)
			assert.NoError(t, err)

			solvedChallenge, err := base64.StdEncoding.DecodeString(tt.solvedChallengeString)

			hashed := sha256.Sum256(challenge)

			err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], solvedChallenge)

			fmt.Printf("\n\nSuccessfully verified\n\n")
		})
	}
}

func verify(challengeString, solvedChallengeString, publicKeyString string) error {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyString))
	if publicKeyBlock == nil {
		return errors.Errorf("Failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return errors.Wrap(err, "Failed to parse public key")
	}

	rsaPubKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.Errorf("Could not convert to rsa.PublicKey")
	}

	fmt.Println("Successfully parsed public key")

	challenge, err := base64.StdEncoding.DecodeString(challengeString)
	if err != nil {
		return errors.Wrap(err, "Failed to parse challenge")
	}

	solvedChallenge, err := base64.StdEncoding.DecodeString(solvedChallengeString)
	if err != nil {
		return errors.Wrap(err, "Failed to parse solved challenge")
	}

	hashed := sha256.Sum256(challenge)

	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], solvedChallenge)
	if err != nil {
		return errors.Wrap(err, "Failed to verify")
	}

	fmt.Println("Successfully verified")

	return nil
}

func Test_service_verify(t *testing.T) {
	type args struct {
		challengeString       string
		solvedChallengeString string
		publicKeyString       string
	}
	tests := []struct {
		name                  string
		challengeString       string
		solvedChallengeString string
		publicKeyString       string
		wantErr               bool
	}{
		{
			name:                  "1",
			challengeString:       "T2UCDtHvrK3uhoyJhzqZ9Fbr+Q/hGIPV/rmxaofm+SI=",
			solvedChallengeString: "VF05gLWIEFavQ9GYIFsjgb+fMmVRUThqICVbxyEqlt1zxO1ytpbyFL+eiTwpiHMBE3w0QfRj9Bk3r6+8UHMcItgTLP2ML8oGTyW4/7YG347LofqHyhCvFN/y04Y/NIkNSn8xD0AiRmzMi1VS9ZwUWHEWnKLwEIFWsBmIQ+YsxdLEKw6SrEKaCS2vDCwCZ4Yv+mTd2hvul5a7lVBVwPxH/KNgtrnrhFMsgkv97XWTCBjIX8aXE8UP2wfZvS8bk2SA1ZvRnqHgXmZUMt7ae5DkXCENBrQQrDAggsiB6Myh/1gJX8nTWMPODTiZ7RYqTLUbEpUqcy51CTFi3GGNf6OSIQ==",
			publicKeyString:       "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmdB5OtML3hnrwVe9jJbdfekrUoF0dia9is6y4CH5fRcq/glxk9kTiVtogYIT49MaqTBPtRbkgzrhM2ViVkzSeGb6quJiRU/voLJTrMUa18PdniGurLcGOEiiNBUo2NMjfqQjY4XfFwahMozORX3DiJ/ayhFJbWPkCl0KyzruqG0VozxQV9eCX7lpdBE+6FvxZuG9by2ZQyTwXR2y8zDPR4Mo0X9n4WlFxdz0wMZ73PV0j3lMW4sfyzxi+p+cGUjT6/uqtFY9dE7BPVH8VLbRYSWCIqd5JH5wuyFk1dsQviVYN1jaEAXqFJYuS7TX8myBUvHCqrHpgDmBlXhHiq1HbQIDAQAB",
			wantErr:               false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			err := s.verify(tt.challengeString, tt.solvedChallengeString, tt.publicKeyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
