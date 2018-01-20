pls-kvs
====================

"Please Key Value Store My Stuff"

A simple command-line key value store. Accepts commands as HTTP requests (GET, PUT, and DELETE) on port 54321, e.g.:
```
curl -X PUT http://localhost:54321/key/key_string -d "value_string"
```
