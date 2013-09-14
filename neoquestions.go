package main

import (
	"flag"
	"fmt"
  "bytes"
  "strconv"
  "sync"
  "time"
  "encoding/binary"
  "encoding/json"
  "crypto/rand"
  "net/http"
)

var (
	consumer_key        = flag.String("consumer_key", "mTicOMGrYaaEZzj89waDg", "Consumer Key")
	consumer_secret     = flag.String("consumer_secret", "m9EwKCoGqCsNnA2EGoVPTxZN5gYU9UxwB1I8wYmE6g", "Consumer Secret")
  access_token        = flag.String("access_token", "892256485-ngmRRVv868YLWv78C8VGz8tTrRKPctMpJp5KwT9a", "Access Token")
  access_token_secret = flag.String("access_token_secret", "C2kjINNwjoJU19LvR7mwDPgSvlMjv06PRPlCPEBhGA", "Access Token Secret")
)

var (
  nonceLock    sync.Mutex
  nonceCounter uint64
)

// nonce returns a unique string.
func nonce() string {
  nonceLock.Lock()
  defer nonceLock.Unlock()
  if nonceCounter == 0 {
    binary.Read(rand.Reader, binary.BigEndian, &nonceCounter)
  }
  result := strconv.FormatUint(nonceCounter, 16)
  nonceCounter += 1
  return result
}

func main() {
	flag.Parse()

  msg := bytes.NewReader([]byte(`{"status":"Bob","testing new client"}`))

  client := &http.Client{}

  var oauth_nonce = nonce()
  var oauth

  var authStr = `OAuth oauth_consumer_key="` + *consumer_key +
    `", oauth_nonce="` + oauth_nonce +
    `", oauth_signature="`+" +
    `", oauth_signature_method="HMAC-SHA1"` +
    `, oauth_timestamp="` + strconv.FormatInt(time.Now().Unix(), 10) +
    `, oauth_token="` + *access_token +
    `", oauth_version="1.0"`

  fmt.Println(authStr)

  req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/statuses/update.json", msg)
  req.Header.Add("Authorization", authStr)
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }

	// Send final carriage return, just to be neat.
	fmt.Println(resp)
}
