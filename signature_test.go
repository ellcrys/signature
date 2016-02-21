package signature

import (
	"testing"
	"time"
	"strings"
	// "github.com/ellcrys/seed"
	"github.com/ellcrys/crypto"
	"github.com/stretchr/testify/assert"
)

var KEYS = []string{
	"-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----",
}

// TestGetSoleTransferSignatureString tests that an expected signature string is returned
func TestGetSoleTransferSignatureString(t *testing.T) {
	var expected = "POST\n%2Fv1%2Fseeds%2Ftransfer\naddr_123\nseed_id,seed_id2,seed_id3\n1460053234"
	var sigstring = GetSoleTransferSignatureString("addr_123", []string{ "seed_id", "seed_id2", "seed_id3" }, 1460053234)
	assert.Equal(t, expected, sigstring)
}

// TestSignSoleTransfer tests that an expected signature is returned 
func TestSignSoleTransfer(t *testing.T) {
	ts := time.Now().Unix()
	expected, err := SignSoleTransfer("addr_123", KEYS[0], []string{ "seed_id", "seed_id2", "seed_id3" })
	assert.Nil(t, err)

	signer, err := crypto.ParsePrivateKey([]byte(KEYS[0]))
	assert.Nil(t, err)
	var sigstring = GetSoleTransferSignatureString("addr_123", []string{ "seed_id", "seed_id2", "seed_id3" }, ts)
	sig, err := signer.Sign([]byte(sigstring))
	sigParts := strings.Split(sig, "\n")

	assert.Nil(t, err)
	assert.Equal(t, (sigstring + "\n" + sigParts[len(sigParts)-1]), expected)
}

// TestParseSoleTransferSignature tests that correct sole transfer signature will be parsed correctly
func TestParseSoleTransferSignature(t *testing.T) {
	signature := "POST\n%2Fv1%2Fseeds%2Ftransfer\naddr_123\nseed_id,seed_id2\n1456085144\nxxx_signature_xxx"
	sigData, err := ParseSoleTransferSignature(signature)
	assert.Nil(t, err)
	assert.Equal(t, sigData["method"], "POST")
	assert.Equal(t, sigData["uri"], "%2Fv1%2Fseeds%2Ftransfer")
	assert.Equal(t, sigData["address_id"], "addr_123")
	assert.Equal(t, len(sigData["seed_ids"].([]string)), 2)
	assert.Equal(t, sigData["seed_ids"].([]string)[0], "seed_id")
	assert.Equal(t, sigData["seed_ids"].([]string)[1], "seed_id2")
	assert.Equal(t, sigData["timestamp"], 1456085144)
	assert.Equal(t, sigData["signature"], "xxx_signature_xxx")
}

// TestParseSoleTransferSignatureIncompleteSignature tests that a parse will fail against a signature that is missing all required signature parts
func TestParseSoleTransferSignatureIncompleteSignature(t *testing.T) {
	signature := "POST\n%2Fv1%2Fseeds%2Ftransfer\naddr_123\nseed_id,seed_id2\nxxx_signature_xxx"
	_, err := ParseSoleTransferSignature(signature)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid signature: 6 signature parts are required")
}