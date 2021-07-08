# chatServer

### Usage
1. `git clone https://github.com/bhankey/chatServer.git`
2. `cd chatServer`
3. `if needed change the server configuration file: 'backend/config.yaml'`
4. `docker-compose up`
5. `use api (default port: 9000)`

### Configuration

configuration of server parameters is carried out using the .yaml file. Examples of configuration file /backend/config/config.yaml. Also you can pass your file throughout flag 'config-path'

### Documentation
To see the documentation, start server and go through the "http://127.0.0.1:9000/swagger/index.html"

### Basic entities
You can see them by opening /database/schema.png

### Examples of usage Api methods

#### Create new user

##### Example request
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"username": "some_user"}' \
http://localhost:9000/users/add
```

##### Example response
```
{
    "user_id": 1
}
```

#### Crete new chat between users

##### Example request
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"name": "chat", "users": [1, 2]}' \
http://localhost:9000/chats/add
```

##### Example response
```
{
    "chat_id": 1
}
```

#### Send new message to chat

##### Example request
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"chat": 1, "author": 1, "text": "message"}' \
http://localhost:9000/messages/add
```

##### Example response
```
{
"message_id": 1
}
```

#### Get chats where user is added (sorted by latest message in chat)

##### Example request
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"user": "1}' \
http://localhost:9000/chats/get
```

##### Example response
```
[ 
    {
        "created_at": "2021-07-08T02:16:18.468708Z",
        "id": 1,
        "name": "string",
        "users_id": [1, 2]
    }
]
```

#### Get messages from chat (sorted by message creation time)

##### Example request
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"chat": 1}' \
http://localhost:9000/messages/get
```

##### Example response
```
[
    {
        "author": 1,
        "chat": 1,
        "created_at": "2021-07-08T02:16:18.468708Z",
        "id": 1,
        "text": "string"
    }
]
```

#### This is the task for trainee backend in Avito (original task: https://gist.github.com/FZambia/82b5b8d89c3005eff9ece3e01882c76f)
