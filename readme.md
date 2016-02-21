# Signature Utility

This package contains convenience method to help create various
transaction signatures to be used on the Ellcrys platform.

## Getting started

First, install package

```go
go get github.com/ellcrys/signature
```

## Sole Transfer Signatures

A sole transfer signature is required by the Ellcrys platform to
authorize transfer request for seeds that have sole ownership. This signature
needs to be provided in the `X-Signature` header of the transfer request.

### Signature String Format

The signature string contains all the information required to create a valid signature
that can be correctly reproduced by Ellcrys. The signature string is signed using the RSA private key
associated with an address (aka *signer* or *signatory address*). A valid signature string is composed of the following format: 

```text
RequestMethod 	+ '\n'   				// request method (use `POST` for transfers)
RequestURI 		+ '\n'      			// URI encoded request uri (uri scheme and host not required)
AddressID 		+ '\n'			        // address id or signer id
SeedIDs  		+ '\n'     		 		// comma separated list of seed ids
Timestamp 		 						// unix time
```

### Full Signature Format

The is the format required to construct a full signature string that can be provided as the `X-Signature`
header value. 

```text
RequestMethod 	+ '\n'   				// request method (use `POST` for transfers)
RequestURI 		+ '\n'      			// URI encoded request uri (uri scheme and host not required)
AddressID 		+ '\n'			        // address id or signer 
SeedIDs  		+ '\n'     		 		// comma separated list of seed ids
Timestamp 		+ '\n' 					// unix time
Signature     							// signature create from signing the signature string with the signer's private key
```

The signature package provides the `GetSoleTransferSignatureString` method to help create a valid signature string.

```go
import (
   "github.com/ellcrys/signature"
   "fmt"
)   

sigStr := signature.GetSoleTransferSignatureString("42503020", []string{"46577","42654","599902"}, 1405882889)
fmt.Println(sigStr)    // POST\n%2Fv1%2Fseeds%2Ftransfer\n42503020\n46577,42654,599902\n1405882889
```

### Sign Sole Transfer

Use the `SignSoleTransfer` method to construct a full signature string to be used as the `X-Signature` string. This method
will create a valid signature string, sign the string and return a full signature.

```go
import (
   "github.com/ellcrys/signature"
   "fmt"
)   

privKey := "-----BEGIN RSA PRIVATE KEY-----\nMIICXg...A/A==\n-----END RSA PRIVATE KEY-----\n"
sig := signature.SignSoleTransfer("42503020", privKey, []string{"46577","42654","599902"})
fmt.Println(sig)    // POST\n%2Fv1%2Fseeds%2Ftransfer\n42503020\n46577,42654,599902\n1405882889\nxxx_signature_xxx
```


