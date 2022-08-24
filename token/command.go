package token

import (
	"fmt"
	"os"
)

const helpString = `user token manage tools

add         Add user token, if exists update.
del         Delete user token.
flush       Flush user token to cache.
help        Get user token manage tools help info.
exit        Exit token manage.

Use "help" for more information about a command.`

func StartTokenCmd() {
	fmt.Println("use [help] cat help info.")
	for {
		var cmdStr string
		fmt.Print(">>> ")
		_, err := fmt.Scanf("%s", &cmdStr)
		if err != nil {
			fmt.Println("The input is incorrect, please re-enter")
			continue
		}
		switch cmdStr {
		case "add":
			addCommand()
		case "del":
			delCommand()
		case "flush":
			flushCmd()
		case "exit", "quit":
			os.Exit(0)
		default:
			helpCommand()
		}
	}
}

func helpCommand() {
	fmt.Println(helpString)
}

func addCommand() {
	// var qq string
	// fmt.Print("input user qq number: ")
	// _, err := fmt.Scanf("%s", &qq)
	// if err != nil {
	// 	fmt.Println("user token create fail, please retry")
	// 	return
	// }
	// token := utils.CreateToken()
	// userToken := &modles.UserToken{
	// 	QQ:    qq,
	// 	Token: token,
	// }
	// err = modles.CreateToken(userToken)
	// if err != nil {
	// 	fmt.Println("user token create fail, please retry")
	// 	return
	// }
	// flushCmd()
	// fmt.Printf("token: %s\n", token)
	// fmt.Println("create user token info success !!!")
	// return
}

func delCommand() {
	// var token string
	// fmt.Print("input user token: ")
	// _, err := fmt.Scanf("%s", &token)
	// if err != nil {
	// 	fmt.Println("user token delete fail, please retry")
	// 	return
	// }
	// userToken := &modles.UserToken{
	// 	Token: token,
	// }
	// err = modles.DeleteToken(userToken)
	// if err != nil {
	// 	fmt.Println("user token delete fail, please retry")
	// 	return
	// }
	// flushCmd()
	// fmt.Println("delete user token info success !!!")
	// return
}

func flushCmd() {
	// tokens, err := modles.GetTokensFromDB()
	// if err != nil {
	// 	fmt.Println("flush user token to db fail")
	// 	return
	// }
	// var tokenSlice []string
	// for _, t := range tokens {
	// 	tokenSlice = append(tokenSlice, t.Token)
	// }
	// err = modles.WriteTokensToCache(tokenSlice)
	// if err != nil {
	// 	fmt.Println("flush user token to db fail")
	// 	return
	// }
	// fmt.Println("flush user tokens success !!!")
	// return
}
