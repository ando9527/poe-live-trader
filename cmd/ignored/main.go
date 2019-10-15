package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ando9527/poe-live-trader/pkg/db/ignored"
	"github.com/sirupsen/logrus"
)

func main(){
	c:=ignored.NewClient()
	c.Connect("sqlite.db")

	c.Migration()


	for{

		fmt.Println("1. Add user into ignored list.")
		fmt.Println("2. Remove user from ignored list.")
		fmt.Println("3. Display ignored list.")
		fmt.Println("4. Quit.")
		fmt.Print("What you choice?(1)")
		s, e := getInput()
		fmt.Println(s)
		if e != nil {
			logrus.Error(e)
			continue
		}
		if s=="2"{
			fmt.Print("Removing User Name?")
			n, e := getInput()
			if e != nil {
				logrus.Error( e)
				continue
			}
			e = c.Remove(n)
			if e != nil {
				logrus.Error("Failed to remove user from list",e)
				continue
			}
			fmt.Printf("Success to remove %s from list\n", n)
			continue
		}else if s =="3"{
			users, e := c.GetAll()
			if e != nil {
				logrus.Error("Failed to get list, ",e)
				continue
			}
			fmt.Println("==========================")
			fmt.Println("Ignored List: ")
			for _,v:=range users{
				fmt.Println(v.Name)
			}
			fmt.Println("==========================")
			continue
		}else if s =="4"{
			os.Exit(0)
		} else{
			fmt.Print("Add User Name?")
			name,e :=getInput()
			if e != nil {
				logrus.Error(e)
				continue
			}
			e = c.Add(name)
			if e != nil {
				logrus.Error("Failed to add user in to list", e)
				continue
			}
			fmt.Printf("Succes to add %s into list\n", name)
		}

	}

}


func getInput()(s string, e error){
	reader := bufio.NewReader(os.Stdin)
	text, e := reader.ReadString('\n')
	if e != nil {
		return "",e
	}
	s=strings.Replace(text,"\n","", -1)
	s=strings.Replace(s,"\r","", -1)
	s=strings.Replace(s," ","", -1)
	s=strings.Replace(s,"@","", -1)
	return s, e
}