go-webperms
==========

`go-webperms` is a dead-simple Authorization package. It introduces `Group` and
`Permission` types which can be used to regulate access control of users in a
web application context.  Because web applications often have unique
requirements for the type which expresses a user's identity, go-webperms
expresses access control in terms of an `User` interface.
