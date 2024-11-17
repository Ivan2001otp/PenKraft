package utils

import ("fmt"
"time"
)

var TTL time.Duration = 1 * time.Hour

func Logger(obj any){
	
	//fmt.Println("Type is :",reflect.TypeOf(obj))

	switch v := obj.(type) {
		
	case int:
			fmt.Println("v is of type : ",v);
		break;

	case float64:
		fmt.Println("v is of type : ",v);
		break;

		case string:
			fmt.Println(v);
			break;
			
		default:
			fmt.Println("v is of default type : ",v);

	}
	fmt.Println(obj)
}