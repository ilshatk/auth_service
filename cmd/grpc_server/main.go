package main

import (
	"fmt"

	"github.com/fatih/color"
	desc "github.com/ilshatk/protos/gen/go/user_v1"
	desc2 "github.com/ilshatk/auth_service/tree/main/pkg/user_v1"
)

type Server struct {
	desc.UnimplementedUserServiceV1Server
}

func main() {

	fmt.Println(color.GreenString("Hello, world!"))
}
