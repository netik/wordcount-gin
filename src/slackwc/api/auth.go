/*
 * John Adams
 * jna@retina.net
 * 8/30/2017
 *
 * slackwc: auth.go
 * Authentication Routines
 *
 */

package api

import (
	"bufio"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

func CheckValidUser(user, pw string) bool {
	/* Validates a username and password against a local
	 * bcrypt flat-file database, returing true if valid, false if not.
	 */
	var err error

	if user == "" || pw == "" {
		return false
	}

	inFile,err := os.Open(AuthFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") { // Ignore comments
			result := strings.Split(scanner.Text(), ":")

			if len(result) == 2 {
				if user == result[0] {
					err = bcrypt.CompareHashAndPassword([]byte(result[1]), []byte(pw))
					if err == nil {
						return true
					}
				}
			}
		}
	}

	// fail closed
	return false
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		// Verify Basic Auth Token
		if len(token) == 0 {
			RespondWithError(401, "Authorization Required", c)
			return
		}

		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			RespondWithError(403, "Authorization Failed", c)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !CheckValidUser(pair[0], pair[1]) {
			RespondWithError(403, "Authorization Failed", c)
			return
		}

	}
}
