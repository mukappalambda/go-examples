# Rate limiting HTTP requests

Example output:

```bash
rate is 1; burst is 4
00-th request, 0s, response: data
01-th request, 201ms, response: data
02-th request, 402ms, response: data
03-th request, 602ms, response: data
04-th request, 803ms, response: rate limiting
05-th request, 1.004s, response: data
06-th request, 1.205s, response: rate limiting
07-th request, 1.407s, response: rate limiting
08-th request, 1.608s, response: rate limiting
09-th request, 1.809s, response: rate limiting
10-th request, 2.01s, response: data
11-th request, 2.211s, response: rate limiting
12-th request, 2.412s, response: rate limiting
13-th request, 2.614s, response: rate limiting
14-th request, 2.815s, response: rate limiting
15-th request, 3.016s, response: data
16-th request, 3.216s, response: rate limiting
17-th request, 3.417s, response: rate limiting
18-th request, 3.618s, response: rate limiting
19-th request, 3.819s, response: rate limiting
20-th request, 4.02s, response: data
```
