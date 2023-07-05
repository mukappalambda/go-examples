# Steps

To follow these intructions, you can refer to [this](https://support.google.com/accounts/answer/185833?hl=en)
for how to generate a Google app password, and paste it in the `.password` file.

```bash
cp .password{.example,}
echo '<YOUR_GOOGLE_APP_PASSWORD>' > .password

go run main.go \
-username <YOUR_GMAIL_ADDR> \
-from <YOUR_GMAIL_ADDR> \
-to <YOUR_GMAIL_ADDR>
```
