#!/usr/local/bin/python# 
 
# John Adams
# jna@retina.net
# 8/30/2017
# 
# slackwc: make_bcrypt.go
# Generic bcrypt encoder
# 

import bcrypt

user = 'bob'
password = b'WhatAboutBob?!'

# Hash a password for the first time, with a randomly-generated salt
hashed = bcrypt.hashpw(password, bcrypt.gensalt())

print "%s:%s\n" % (user, hashed)

