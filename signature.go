// Package signature contains convenience method to help create various transaction 
// signatures to be used on the Ellcrys platform.
package signature

import (
	"strings"
	"fmt"
	"time"
	"errors"
	"net/url"
	"strconv"
 	"github.com/ellcrys/crypto"
)

const transferEndpoint = "/v1/seeds/transfer"

// Construct signature string for sole transfer.
// 	Format: RequestMethod + '\n' + CanonicalEndpointURI + '\n' + AddressID + '\n' + ShellIDs + '\n' + CurrentTime
// 	Where:
//   RequestMethod: POST
// 	 CanonicalEndpointURI: /v1/seed/transfer | Must be URI encoded
// 	 AddressID: Address id (aka signer id)
// 	 ShellIDs: Comma separated list of shell id
// 	 CurrentTime: Current unix time
func GetSoleTransferSignatureString(addressID string, shellIDs []string, currentUnixTime int64) string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%d", "POST", url.QueryEscape(transferEndpoint), addressID, strings.Join(shellIDs, ","), currentUnixTime)
}


// Creates a sole transfer signature. The signature contains the signature string and the signature 
// concantenated together.
func SignSoleTransfer(addressID, addressPrivKey string, shellIDs []string) (string, error) {
	var signatureString = GetSoleTransferSignatureString(addressID, shellIDs, time.Now().Unix())
	signer, err := crypto.ParsePrivateKey([]byte(addressPrivKey))
	if err != nil {
		return "", err
	}
	sig, err := signer.Sign([]byte(signatureString))
	if err != nil {
		return "", err
	}
	return (signatureString + "\n" + sig), nil
}


// Given a sole transfer signature, it attempts to parse it. 
// If everything goes well, a map containing the various signature data is returned.
// Rules:
// - Signature must have 6 parts
func ParseSoleTransferSignature(signature string) (map[string]interface{}, error) {
	
	var data = map[string]interface{}{}
	parts := strings.Split(strings.Trim(signature, "\n"), "\n")
	
	// expect 5 parts (request method, encoded url path, address id, shell ids, current time and signature)
	if len(parts) != 6 {
		return data, errors.New("invalid signature: 6 signature parts are required")
	}

	data["method"] = parts[0]
	data["uri"] = parts[1]
	data["address_id"] = parts[2]
	data["seed_ids"] = strings.Split(parts[3], ",")
	ts, _ := strconv.Atoi(parts[4])
	data["timestamp"] = ts 
	data["signature"] = parts[5]

	return data, nil 
}