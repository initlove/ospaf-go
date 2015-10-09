##Find Avaiable User by Composition
Find the avaiable user by the full char composition.
Two parameters: level(how many chars), charSet(the char set).

##How to use
- Add your own [accounts.json](#accounts.json)
- Modify the .go file (change 'level' and 'charSet' if you like)
- Make and Run

```
make
./find-avaiable-users-by-compose > output
``` 

##accounts.json
```
[
	{
		"Type": "Basic",
		"User": "golang001",
		"Password": "qwe123456"
	},
	{
		"Type": "Basic",
		"User": "golang002",
		"Password": "qwe123456"
	}
]
```
