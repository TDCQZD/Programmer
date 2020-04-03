## crypto-js

CryptoJS (crypto.js) 为 JavaScript 提供了各种各样的加密算法。目前已支持的算法包括：
- MD5
- SHA-1
- SHA-256
- AES
- Rabbit
- MARC4
- HMAC
- HMAC-MD5
- HMAC-SHA1
- HMAC-SHA256
- PBKDF2

示例代码：
```
var digest = Crypto.MD5("Message");

var digestBytes = Crypto.MD5("Message", { asBytes: true });
var digestString = Crypto.MD5("Message", { asString: true });
```

API: https://github.com/brix/crypto-js