Vim�UnDo� ����xH������\v�xq��/���ZH              	                       Y{]�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             Y{\\     �                 �              5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Y{\a     �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Y{\a     �                  5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                                             Y{\e     �                 
import fmt5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Y{\e     �                  5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �                 func main()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �             �                 func main(){}5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �               	5�_�      
           	          ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �               	messages := make()5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      	         	messages := make(chan string)5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   	      	5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   
      	go func ()5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   
      	go func (){}5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   
      	go func (){messages<-}5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   
      	go func (){messages <-}5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �      
   
      	go func (){messages <- ""}5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �         
      	go func (){messages <- "ping"}5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �   	            	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{\�     �   
            	msg:= <-messages5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]      �               	5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]    �               	fmt.Println()5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             Y{]/    �      
         	go func (){messages <- "ping"}5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             Y{]M     �               
import fmt5�_�                       
    ����                                                                                                                                                                                                                                                                                                                                                             Y{]O     �               
import fmt5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]Q     �               import fmt""5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]T     �               import ""fmt""5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]U     �               import fmt""5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             Y{]V     �               import ""fmt""5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Y{]W     �               import "fmt""5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Y{]\    �                  package main       import "fmt"       func main(){       	messages := make(chan string)       !	go func (){messages <- "ping"}()       	msg:= <-messages       	fmt.Println(msg)   }5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             Y{]�    �                  package main       import "fmt"       func main() {       	messages := make(chan string)       #	go func() { messages <- "ping" }()       	msg := <-messages       	fmt.Println(msg)   }5��