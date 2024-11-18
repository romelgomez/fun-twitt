# fun-twitt
fun-twitt is api mock of twitter/X in golang


## models

db provider: `sqlite`

- [X] users  ~ a model ~ ids start with `user_`
- [X] tweets ~ a model ~ ids start with `tweet_` 
- [X] follow ~ a model ~ ids start with `fw_`
- [X] timeline ~ a query

## Setting Up Local Development

1. Run migrations to initialize the database:
   ```bash
   go run github.com/steebchen/prisma-client-go db push

## test

`go test ./... -v` 

## Postman Views

### User model

![post_user](images/post_user.png)
![patch_usert](images/patch_user.png)
![get_user](images/get_user.png)
![get_all_user](images/get_all_user.png)

### Tweet model

![post_tweet](images/post_tweet.png)
![get_tweet_by_user_id](images/get_tweet_by_user_id.png)

### Follow model

![get_follow_follower](images/get_follow_follower.png)

### Timeline query

![get_timeline](images/get_timeline.png)