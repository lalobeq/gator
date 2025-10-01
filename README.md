# Gator
A Blog Aggregator project made with Go (for learning purposes)

## Requirements
This project requires postgreSQL and Go to run the program.
To install postgreSQL, you can run in the terminal (for macOS):
`brew install postgresql@15`
For Go, well, if you are here it is possible you have Go already installed, if not, you can go
to the official site and download it:
https://go.dev/doc/install.

Once the requirements are met up, you can install gator with
`go install github.com/lalobec/gator`

## Configuration
Then you need to configure some stuff. In your home directory, create json file
named ".gatorconfig.json", add the following:
```json
{"db_url":"postgres://user:@localhost:5432/gator?sslmode=disable","current_user_name":"yourname"}
```

## Usage
Now that things are working, here is how you can use the tool:
### Commands
    - gator register <username>
    - gator login <username>
        In order to login, the username should already exist in the db.
    - gator users 
        List the users in the db
    A user must be logged in to use the following commands
    - gator addfeed <url>
    - gator follow <url>
    - gator following
    - gator unfollow <url>
    - gator browse limit n
        limit is optional, default is set to 2

### Reset
    You can run
    - gator reset 
        to reset the db, this will erase all!
