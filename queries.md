### PLIS NOTE 
* If the second part(if any) of the blocks is the input part. Use as `query variables`



``` 
query get_one_user{
  one_user(username:"uno"){
    s{
      username
      password
      
    }
  }
}
```

```
query get_all_users{
  user{
    username
  }
}
```

```
query one_user($in : String!){
  oneuser(username: $in){
    username
  }
}

{
  "in": "username"
}
```

```
query get_all_challenge{
  challenge{
    name
    description
    category
    value
    flags
  }
}
```

```
query get_one_challenge($in : String!){
  onechallenge(name:$in){
    name
    description
    category
    value
    flags
  }
}

{
  "in":"XQLi"
}
```

```
{
  scoreboard {
    username
    solved {
      ChallID
    }
    score
  }
}
```