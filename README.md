# jubilant-potato
A little tool to convert go json file to rust version 

## example
json file in go like
```
type UserInfo struct{
    User  string  `json:"user"`
    Addr  Address `json:"address"`
}

type Address struct{
    a string `json:"a"`
    somelist []string `json:"somelist"`
}
```

and it would be converted into rust file like
```
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
struct UserInfo{
    user:String,
    address:Address,
}

#[derive(Serialize, Deserialize)]
struct Address{
    a:String,
    somelist:Vec<String>,
}
```

## Tips
Currently it doesn't consider key words such as 'type'. You could correct it by add macro 
```
#[serde(rename=xxx)]
```

## how to use
```
cd scripts
chmod +x run.sh
```
Change the file path in that script and just run it