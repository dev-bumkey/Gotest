package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}
type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 20}
	vip := VIPUser{
		User{"화랑", "hwarang", 43},
		3,
		250,
	}
	fmt.Printf("유저: %s ID : %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저 :%s ID : %s 나이: %d 레벨: %d 가격: %d\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price)
}
