package main

import(
  "fmt"
  "net"
  "os"
  "encoding/json"
  "bytes"
  "io"
)

//Defining json structs
type User struct{
  Name UserInfo
  Email []Email
}

type UserInfo struct {
  userName string
}

type Email struct {
  emailAddress string
}

//Function to create a user using User struct
func (u User) String() string {

  userString := u.UserInfo.userName

  for _, emailData := range u.Email {
    userString += "\n" + emailData.emailAddress
  }

  return userString
}

//Function main creates an instance of the user and prompts the user for input, checks for errors,
// encodes and decodes the user info, sends the user info and reads it back to the client
func main() {

  user := User{

    Name: UserInfo{userName: "TestUser"},
    Email: []Email{Email{emailAddress: "testuser@test.com"}

      }
  }

    if len(os.Args) != 2 {
      fmt.Println("Usage: ", os.Args[0], "host:port")
      os.Exit(1)
    }

    address: os.Args[1]

    conn, err := net.Dial("tcp", address)
    checkError(err)

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    encoder.Encode(user)
    var newUser User
    decoder.Decode(&newUser)
    fmt.Println(newUser.String())
    os.Exit(0)

  }

  //Function to check for errors and print them if any
  func checkError(err error) {

    if err != nil {
      fmt.Println("Error", err.Error())
      os.Exit(1)
    }

  }

/* NOTE: same script can be used to store gob serialised data

    1- In import change "encoding/json" to "encoding/gob"
    2- In the main function change user.json to user.gob also
    3- In the laod funciton change the encoder and decoder values to gob.NewEncoder and gob.NewDecoder
    
 */
