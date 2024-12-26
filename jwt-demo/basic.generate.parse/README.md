# Generating and parsing JWT

```console
$ # run help message
$ go run main.go -help
Usage of /tmp/go-build1073546655/b001/exe/main:
  -exp.iss string
        expected issuer (default "my.auth.server")
  -iss string
        issuer (default "my.auth.server")
  -require.exp
        requires exp claim. default to false
$ # run with the default setting
$ go run main.go
Signed first token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.BgjKCNps7Rx6IPmq_BzL1wwfiVOmpogyowiEi6iVVBk"
Signed second token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS5hdXRoLnNlcnZlciIsImtleSI6InZhbHVlIiwicm9sZXMiOiJhZG1pbiB1c2VyIiwic2NvcGUiOiJyZWFkOm1lc3NhZ2VzIHdyaXRlOm1lc3NhZ2VzIn0.ruVUfO76S2dcW-Y-1dhNHjuTsqgJ5sRbbFi5296ZLg4"
key: "iss"; value: "my.auth.server"
key: "key"; value: "value"
key: "roles"; value: "admin user"
key: "scope"; value: "read:messages write:messages"
$ # run with different issuer and expected issuer
$ go run main.go -iss test1 -exp.iss test2
Signed first token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.BgjKCNps7Rx6IPmq_BzL1wwfiVOmpogyowiEi6iVVBk"
Signed second token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0MSIsImtleSI6InZhbHVlIiwicm9sZXMiOiJhZG1pbiB1c2VyIiwic2NvcGUiOiJyZWFkOm1lc3NhZ2VzIHdyaXRlOm1lc3NhZ2VzIn0.9I5treDPfVLcrJX5KRWjYtmgQcenUfjEm7jzkfP-x3A"
2024/12/26 10:48:56 Error parsing the token string: token has invalid claims: token has invalid issuer
exit status 1
$ run with the same issuer and expected issuer
$ go run main.go -iss test1 -exp.iss test1
Signed first token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.BgjKCNps7Rx6IPmq_BzL1wwfiVOmpogyowiEi6iVVBk"
Signed second token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0MSIsImtleSI6InZhbHVlIiwicm9sZXMiOiJhZG1pbiB1c2VyIiwic2NvcGUiOiJyZWFkOm1lc3NhZ2VzIHdyaXRlOm1lc3NhZ2VzIn0.9I5treDPfVLcrJX5KRWjYtmgQcenUfjEm7jzkfP-x3A"
key: "iss"; value: "test1"
key: "key"; value: "value"
key: "roles"; value: "admin user"
key: "scope"; value: "read:messages write:messages"
$ # run with requiring the exp claim
$ go run main.go -iss test1 -exp.iss test1 -require.exp
Signed first token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.BgjKCNps7Rx6IPmq_BzL1wwfiVOmpogyowiEi6iVVBk"
Signed second token successfully, signed token string: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUxODQ5NjgsImlzcyI6InRlc3QxIiwia2V5IjoidmFsdWUiLCJyb2xlcyI6ImFkbWluIHVzZXIiLCJzY29wZSI6InJlYWQ6bWVzc2FnZXMgd3JpdGU6bWVzc2FnZXMifQ.4cx93Iv-Rtmt4oOSiMkuxBPTzE-OtaHDxQFNEBp1KTo"
key: "iss"; value: "test1"
key: "key"; value: "value"
key: "roles"; value: "admin user"
key: "scope"; value: "read:messages write:messages"
key: "exp"; value: %!q(float64=1.735184968e+09)
```
