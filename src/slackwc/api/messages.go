/*
 * John Adams
 * jna@retina.net
 * 8/30/2017
 *
 * slackwc: messages.go
 * Message passing types (json)
 *
 */

package api

type WordRequest struct {
	Input string `json:"input"`
}

type WordList struct {
	Count int            `json:"count"`
	Words map[string]int `json:"words"`
}
