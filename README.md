# feishu bot talker

feishu bot talker is a cli tool for handle feishu api to communicate with feishu bots.

## Getting Started

```bash
# build
$ make go-build
```

### Show bot talker content schema

```bash
# checkout feishu api content schema to @ everybody
$ cd cmd/feishu-bot-talker
$ ./feishu-bot-talker template show @all

# checkout feishu api content schema to @ somebody
$ ./feishu-bot-talker template show @someone

# checkout feishu api content schema that contains some text
$ ./feishu-bot-talker template show complex

# checkout feishu api content schema that contains one line text
$ ./feishu-bot-talker template show simple
```

### Sending message

```bash
# create a template file that will @all to edit
# this will output a file in current dir with name feishu-template-output
$ ./feishu-bot-talker template new @all

# edit you file to say whatever you want
$ vi feishu-template-output

# send it
$ ./feishu-bot-talker --app-id your-bot-app-id --app-secret your-bot-app-secret --receive-id-type chat_id --receive-id-value your-chat-id --content-file feishu-template-output send
```

## milestone

* Feishu with post type content support
* Feishu bot answer features
* Feishu bot answer with llm