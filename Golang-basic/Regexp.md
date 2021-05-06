## regular expressions

https://zetcode.com/golang/regex/

https://yourbasic.org/golang/regexp-cheat-sheet/

https://www.golangprograms.com/find-dns-records-programmatically.html

https://geekflare.com/regular-expression-tester/

https://code.tutsplus.com/tutorials/8-regular-expressions-you-should-know--net-6149

https://code.tutsplus.com/tutorials/regular-expressions-with-go-part-1--cms-30403

https://stackoverflow.com/questions/9363145/regex-for-extracting-filename-from-path

https://stackoverflow.com/questions/47131996/go-how-to-check-if-a-string-contains-multiple-substrings/47133055

***Check Filename use this regular expression***
```
"^[a-zA-Z0-9_\\-\\.]*$"
```
Explanations:
```
^ : start of string
[ : beginning of character group
a-z : any lowercase letter
A-Z : any uppercase letter
0-9 : any digit
_ : underscore 
- : hyphen (double escape signal \)
. : dot  (double escape signal \)
] : end of character group
* : zero or more of the given characters
$ : end of string
```
Example:
```
 Check validate sponsor_adxgg1.js 
 Check validate iplay-mutex-ssp.js 
 Check validate muachungplaza_rc4.js 
 Check validate 131.js 
 Check validate admCoreVtvLive.js 
 Check validate C:\AppServ\www\project\storage\pr.vn.js 
 Filename C:\AppServ\www\project\storage\pr.vn.js is invalid --> (\ signal in path is not in matching regular)
```
