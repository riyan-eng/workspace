package config

import (
	"fmt"
	"os"

	"server/infrastructure"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := sqladapter.NewAdapter(infrastructure.SqlDB, "postgres", "permissions")
	if err != nil {
		fmt.Printf("casbin: failed to initialize adapter - %v \n", err)
		os.Exit(1)
	}
	enforce, err := casbin.NewEnforcer("./casbin.conf", adapter)
	if err != nil {
		fmt.Printf("casbin: failed to create enforcer - %v \n", err)
		os.Exit(1)
	}

	policies := [][]string{
		{"ADMIN", "/perangkat/*", "(GET)|(POST)|(PATCH)|(DELETE)"},
	}

	_, err = infrastructure.SqlxDB.Exec("delete from permissions where p_type = 'p' ")
	if err != nil {
		fmt.Printf("casbin: failed to delete policies - %v \n", err)
		os.Exit(1)
	}
	_, err = enforce.AddPoliciesEx(policies)
	if err != nil {
		fmt.Printf("casbin: failed to add policies - %v \n", err)
		os.Exit(1)
	}
	enforce.LoadPolicy()

	return enforce
}
