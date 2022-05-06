package main

import( 
	"os"
	"fmt"
	"errors"
	"github.com/hashicorp/go-version"
)

func main() {
	//Validating No input case
	if len(os.Args) <=1 {
		fmt.Println(errors.New("Error: No Input found, please enter valid input"))
		os.Exit(2)
	}

	//Comparing the two input versions
	result, err := compareVersions(os.Args[1], os.Args[2]);
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(result)

}

//Function to Compare two version strings using Hasicorp version library
func compareVersions(v1, v2 string) (int, error){

	//Coverting string into version
	version1, err := version.NewVersion(v1)
	if err != nil {
		return 0, errors.New("Error: Invalid input version")
	}

	//Coverting string into version
	version2, err := version.NewVersion(v2)
	if err != nil {
		return 0, errors.New("Error: Invalid input version")
	}

	//Comparing the versions
	if version1.GreaterThan(version2){
		return 1, nil
	} else if version1.LessThan(version2){
		return -1, nil
	} else {
		return 0, nil
	}
}