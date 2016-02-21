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

### Signature String

The signature string contains all the information required to create a valid signature
that can be correctly be reproduced by Ellcrys. A signature string is signed using the private key
associated with an address (aka *signer* or *signatory address*). A valid signature string is composed of the following format: 

```text
RequestMethod + '\n'   				// request method (use `POST` for transfers)
URIEncode(*RequestURI*) + '\n'      // encoded request uri
AddressID + '\n'			        // address id or signer 
SeedIDs  + '\n'     		 		// comma separated list of seed ids
Timestamp + '\n' 					// unix time
```

The signature package provides the `GetSoleTransferSignatureString` method to help create a valid signature string.

```go
import "github.com/ellcrys/signature"
import "fmt"

sig := signature.GetSoleTransferSignatureString("42503020", []string{"46577,42654,599902"}, 1405882889)
fmt.Println(sig)    // POST\n%2Fv1%2Fseeds%2Ftransfer\n42503020\n46577,42654,599902\n1405882889
```



