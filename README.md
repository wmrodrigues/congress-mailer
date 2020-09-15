# congress-mailer
Simple mailer process to send a bunch of mail messages to addresses of a list from a csv file 

The motivation to build this small and simple project was based on appeal made by congressman [Kim Kataguiri](https://www.instagram.com/kimkataguiri), saying that we should ask to our fellow deputies congressmen for support to approve his anti privileges constitutional amendment. Well, this project served to its purposes but I think it can be used by anyone else that wants to help.

### Requirements:
* [Go lang 1.13](https://golang.org/doc/go1.13)

### Instructions:
After downloading this source code, go to the project directory:
```
username@machine:~$ cd congress-mailer
```

Assuming you already installed [Go](https://golang.org/), now you should:

1. Configure the csv source file with the mail addresses that you want to send messages to:
     ```
     username@machine:~$ cd configs
     username@machine:~$ cp file.csv.example file.csv
     ```
   Open the file.csv file using any text editor you want (I certainly prefer [Visual Studio Code](https://code.visualstudio.com/download), but sure, you can use nano if you insist, it's free, who cares?) and check it out:    
   
   The file must have the follow layout:

    > Name,email,treatment

    *e.g.:*
    ```
    Isaac Newton,isaac.newton@email.com,Mr
    Archimedes of Syracuse,syracuse.archimedes@email.com,Dr
    Leonhard Euler,leonhard.euler,Dr
    Euclid of Alexandria,alexandria.euclid@email.com,Mr
    Carl Friedrich Gauss,carl.gauss@email.com,Dr
    Albert Einstein,albert.einstein@email.com,Mr
    ```
    
    You can input how many address you want, really, it's up to you, be creative!

2. Configure the settings.json file, this is a very important part:
    ```
    username@machine:~$ cd configs
    username@machine:~$ cp settings.json.example settings.json
    ```
    Open the settings.json file using any text editor you want (I certainly prefer [Visual Studio Code](https://code.visualstudio.com/download), but sure, you can use nano if you insist, it's free, who cares?) and fill out the attributes:
    ```
    {
      "smtp": {
        "host": "smtp.gmail.com",
        "port": 587,
        "username": "username@gmail.com",
        "password": "yourTopSecretPassword",
        "from": "Your Awesome Nome"
      },
      "mail": {
        "subject": "My awesome subject"
      }
    }
    ```
    
    You should check those settings out with the email provider you want to use.

3. Prepare you message template, in here, you can write whatever you want and use the data you put on your csv file:
    ```
    username@machine:~$ cd configs
    username@machine:~$ cp message.template.example message.template.json
    ``` 

    You can write something like:
    
    ```
    Hello {{.Tratment}} {{.Name}}, my old friend...
    I've come to talk with you again...
    ```

**Done, you're all set!**  
You're ready to use the congress-mailer!

To build it, run the following command:
```
username@machine:~$ go build ./cmd/mailer
```

A binary file should be generated and then, you run the following command:
```
./mailer --file=/path/to/your/awesmome/file.csv
```

If everything went well, you'll see some logs like these:
```
2020/09/09 21:08:50 starting the congress-mailer...
2020/09/09 21:08:50 loading settings file
2020/09/09 21:08:50 loading csv file from /path/to/your/awesmome/file.csv
2020/09/09 21:08:50 now the fun part, sending mail
sending mail 0 to Mr. Isaac Newton  - isaac.newton@email.com
2020/09/09 21:08:55 mail succesfully sent to isaac.newton@email.com
sending mail 1 to Dr. Archimedes of Syracuse  - syracuse.archimedes@email.com
2020/09/09 21:08:55 mail succesfully sent to syracuse.archimedes@email.com
```

Now, if something different then those messages show up, well, then you certainly did something wrong, please go back to the first instruction and check all out one more time.

If you don't have go installed and you also don't want to, then you can use [Docker](https://docs.docker.com/get-docker/) to run the project (after also doing the steps 1 and 2):
``` 
username@machine:~$ docker run --rm -v /path/to/your/awesmome/file.csv:/usr/local/file.csv -v $(pwd):/app -w /app golang:1.13 bash -c "go build ./cmd/mailer && ./mailer --file=/usr/local/file.csv" 
```

This is fun, right?

Now, if you don't have go nor even Docker, seriously dude, what world are you living in?

Did you enjoy this? Send me your feedback, I'm a really nice person and accept your suggestions.  
You can talk to me at washington.moises@gmail.com  
Tks! 

Built with <span style="color:transparent; text-shadow: 0 0 0 yellow;">♥</span> by Wash