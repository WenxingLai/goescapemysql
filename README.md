# goescapemysql

Package goescapemysql provides a go function that escape a string as 
[mysql-server](https://github.com/mysql/mysql-server) 
[does](https://github.com/mysql/mysql-server/blob/c4f63caa8d9f30b2850672291e0ad0928dd89d0e/mysys/charset.c#L789). 

This function escapes strings by adding backslashes before special characters, such as making the null byte `\0` and 
the newline byte `\n`. 

The implementation also converts `%` to `\%` and `_` to `\_`, respectively. 

Also, `RealRealEscapeString` is provided, which does not escape `%` and `_`.

##### Acknowledgement

Package goescapemysql is for personal use only and does not guarantee anything.

##### License

See [WTFPL](http://www.wtfpl.net/txt/copying/).
