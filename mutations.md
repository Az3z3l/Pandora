
``` 
mutation register_user($in:register_input){
  register(input: $in)
}

{
   "in":{
      "email": "email@email.email",
      "username": "username",
      "password": "password"
        } 
}
```


```
mutation AddnewChallenge ($in : add_challenge_data){
  add_challenge(input:$in)
}

{
  "in":{
    "name": "XQLi",
    "category": ["Web"],
    "value": 111,
    "flags": "flag{le-platform}",
    "visibility": false,
    "description": "KZDFUS2SKZNHETKEJJLVEVS2JNLDCUSTK5DFUV3CIV4FIUSYKJUFM3LUNNJEKNKVKRWEMVCSNRNE\nQVKWKF3U6VSCKVGUK42LBI======"
  }
}
```

```
mutation userProfile($in : useredit){
  UpdateUser(input:$in)
}

{
  "in": {
    "ID": "5f59cc3b2a9ec7c8ec10068d",
    "username": "Az3z3l",
    "email": "email@star.com"
  }
}
```


```
mutation newpassword($in : resetpwd){
  reset_pwd(input: $in)
}

{
  "in": {
    "ID": "5f59cc3b2a9ec7c8ec100s",
    "oldpwd": "aaaaaaaaaaaaaaaa",
    "newpwda": "asta",
    "newpwdb": "asta"
  }
}
```

```
mutation UpdataChallenge ($in : edit_challenge_data){
  edit_challenge(input:$in)
}

{
  "in":{
    "ID": "5f57818e23104197b1306ba6",
    "name": "xqli II - revenge of recaptcha",
    "description":"lqs+ssx=ilqx",
    "category": ["Web", "Rev"],
    "value": 11134,
    "flags": ["flag{le-platform}"],
  }
}
```


```
mutation publicise( $in : public ){
  challvisibility(input:$in)
}

{
  "in": {
    "ID": "5f5780da23104197b1306ba5",
    "visibility": true
  }
}
```

```
mutation flaggy($in :flagInput){
  flag_submit(input:$in)
}


{
  "in": {
    "ID": "5f5bb5e18d9e7a44959d167a",
    "flag": "flag{1000-pt-chall}"
  }
}
```