package apitest

import (
	"testing"
	"slackwc/api"
	"github.com/stretchr/testify/assert"
)

/* test the users in our current password file */

func TestCheckValidUser(t *testing.T) {
	res := api.CheckValidUser("foo","bar")
	assert.Equal(t, res, false, "Invalid logins should not work")
}

func TestEmptyLogin(t *testing.T) {
	res := api.CheckValidUser("","")
	assert.Equal(t, res, false, "Empty logins should fail")
}

func TestPartialLogin(t *testing.T) {
	res := api.CheckValidUser("","xxx")
	assert.Equal(t, res, false, "Partial logins (no user) should fail")
}

func TestPartialLogin2(t *testing.T) {
	res := api.CheckValidUser("yyy","")
	assert.Equal(t, res, false, "Partial logins (no pw) should fail")
}

func TestValidLogin(t *testing.T) {
	res := api.CheckValidUser("bob","WhatAboutBob?!")
	assert.Equal(t, res, true, "Valid logins should work")
}
