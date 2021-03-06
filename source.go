package main

import(

	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
)

func replaceAtIndex(in string, i int) string {
    out := []rune(in)
    
    for  ; i<len(in) ; i++{
    	out[i-1]=out[i]
    }
    out[i-1]='\n'
    return string(out)
}

func checkWithFullName(sectionName string, studentName string)bool{

	studentName=strings.ToUpper(studentName)

	file , err:=os.Open(sectionName)

	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	for scanner.Scan(){

		fileStudentName:=scanner.Text()

		fileStudentName=strings.ToUpper(fileStudentName)

		if fileStudentName==studentName{
			return true
			break
		}

	}
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	return false

}

func checkInResultMain(firstName string , lastName string , resultFile string)bool{

	firstName=strings.ToUpper(firstName)
	lastName=strings.ToUpper(lastName)

	file , err := os.Open(resultFile)

	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	for scanner.Scan(){
		redLin:=scanner.Text()
		redLin=strings.ToUpper(redLin)

		if ( strings.Contains(redLin,firstName) &&  strings.Contains(redLin,lastName) ){
			return true
		}
	}
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return false

}

func checkInResult(studentname string)int32{

	var count int32
	count=0
	studentnam:=strings.Split(studentname," ")
	for i:=1 ; i<=3 ; i++{
		if i==1{
			if checkInResultMain(studentnam[0],studentnam[1],"casis/results.txt")==true{
				count++
			}
		}
		if i==2{
			if checkInResultMain(studentnam[0],studentnam[1],"lme/results.txt")==true{
				count++
			}
		}
		if i==3{
			if checkInResultMain(studentnam[0],studentnam[1],"patsy/results.txt")==true{
				count++
			}
		}
	}


	return count

}

func checkinSection(fullNam string , whichTypeofName int32 , sectionPath string)bool{

	printSection := make([]string, 0)
	
	checkifExist:=0

	

		for i:=1 ; i<=14 ; i++{


			if i==1 {

				if checkWithFullName("sections/1st.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/1st.txt"{
					printSection = append(printSection,"1st")
					}

				}
			}
			if i==2 {

				if checkWithFullName("sections/2nd.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/2nd.txt"{
					printSection = append(printSection,"2nd")
					}
				}
			}
			if i==3 {

				if checkWithFullName("sections/3rd.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/3rd.txt"{
					printSection = append(printSection,"3rd")
					}
				}
			}
			if i==4 {

				if checkWithFullName("sections/4th.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/4th.txt"{
					printSection = append(printSection,"4th")
					}					
				}
			}

			if i==5 {

				if checkWithFullName("sections/5th.txt" , fullNam)==true{

					checkifExist=1
					if sectionPath!="sections/5th.txt"{
					printSection = append(printSection,"5th")
					}

				}
			}
			if i==6 {

				if checkWithFullName("sections/mon.txt" , fullNam)==true{

					checkifExist=1
					if sectionPath!="sections/mon.txt"{
					printSection = append(printSection,"mon")
					}
					
				}
			}
			if i==7 {

				if checkWithFullName("sections/tue.txt" , fullNam)==true{

					checkifExist=1
					if sectionPath!="sections/tue.txt"{
					printSection = append(printSection,"tue")
					}

				}
			}
			if i==8 {

				if checkWithFullName("sections/wed.txt" , fullNam)==true{
					
					checkifExist=1
					if sectionPath!="sections/wed.txt"{
					printSection = append(printSection,"wed")
					}
				}
			}

			if i==9 {

				if checkWithFullName("sections/thu.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/thu.txt"{
					printSection = append(printSection,"thu")
					}
				}
			}
			if i==10 {

				if checkWithFullName("sections/fri.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/fri.txt"{
					printSection = append(printSection,"fri")
					}
				}
			}
			if i==11 {

				if checkWithFullName("sections/u300.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/u300.txt"{
					printSection = append(printSection,"u300")
					}
				}
			}
			if i==12 {

				if checkWithFullName("sections/u600.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/u600.txt"{
					printSection = append(printSection,"u600")
					}
				}
			}

			if i==13 {

				if checkWithFullName("sections/u1600.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/u1600.txt"{
					printSection = append(printSection,"u1600")
					}
				}
			}
			if i==14 {

				if checkWithFullName("sections/k.txt" , fullNam)==true{
					checkifExist=1
					if sectionPath!="sections/k.txt"{
					printSection = append(printSection,"k")
					}
				}
			}
						
			
		}



		if (checkifExist==1 && whichTypeofName==3){

				//fmt.Println(fullNam,"found in : ",printSection)
				var h int

				for h=0 ; fullNam[h]!=' '; {
					h++
				}

				h++

				fullNam=replaceAtIndex(fullNam,h)

				//fmt.Println("yep : ",fullNam)

				

			}


		if (checkifExist==1) {

				fmt.Println(fullNam,"found in  : ",printSection)


				if checkInResult(fullNam)==3{
					fmt.Println(fullNam," won PLATINUM award !")
				}else{
					if checkInResult(fullNam)==2{
						fmt.Println(fullNam," won GOLD award !")
					}else{
						if checkInResult(fullNam)==1{
							fmt.Println(fullNam," won SILVER award !")
						}else{
							fmt.Println(fullNam," won BRONZE award !")
						}
					}

				}
				return true	

			}


	return false

}

func fullName(){
	
	var name1,name2,name3 string

	// fmt.Println("you are in fullName")

	fmt.Printf("Enter your full name : ")

	fmt.Scanf("%s %s %s",&name1,&name2,&name3)

	fmt.Println("Entered name is : "+name1+" "+name2+" "+name3)	

	// fmt.Println("length of name1 : ",len(name1),"and name2 : ",len(name2),"and name3 : ",len(name3))

	if len(name3)==0{
		name:=[]string{name1,name2}
		fullNam:=strings.Join(name," ")
		if checkinSection(fullNam,2,"ffwef")==false{
			fmt.Println("Name not found")
		}
	} else{
		name:=[]string{name1,name2,name3}
		fullNam:=strings.Join(name," ")
		if checkinSection(fullNam,3,"ffwef")==false{
			fmt.Println("Name not found")
		}
	}

	//fmt.Println("Full name is : "+fullNam)

}


func checkWithPartName(sectwithFile string , pname string)bool{

	file , err := os.Open(sectwithFile)

	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	foo:=strings.Split(sectwithFile,"/")

	ans:=strings.Split(foo[1],".")

	pname=strings.ToUpper(pname)
	checkforExist:=false
	scanner:=bufio.NewScanner(file)

	for scanner.Scan(){

		fileLine:=scanner.Text()
		fileLine=strings.ToUpper(fileLine)

		if strings.Contains(fileLine,pname)==true{
			fmt.Println(pname," found in : ",ans[0]," as ",fileLine)
			checkforExist=true

			if checkInResult(fileLine)==3{
					fmt.Println(fileLine," won PLATINUM award !")
				}else{
					if checkInResult(fileLine)==2{
						fmt.Println(fileLine," won GOLD award !")
					}else{
						if checkInResult(fileLine)==1{
							fmt.Println(fileLine," won SILVER award !")
						}else{
							fmt.Println(fileLine," won BRONZE award !")
						}
					}

				}


		}


	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	if checkforExist==true{
		return true
	}else{
		return false
	}


}


func checkpartialInSection(partname string)bool{

	var checkifExist int32

	for i:=1 ; i<=14 ; i++{


			if i==1{

				if checkWithPartName("sections/1st.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==2{

				if checkWithPartName("sections/2nd.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==3{

				if checkWithPartName("sections/3rd.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==4{

				if checkWithPartName("sections/4th.txt" , partname)==true{
					checkifExist=1
										
				}
			}

			if i==5{

				if checkWithPartName("sections/5th.txt" , partname)==true{

					checkifExist=1
					
				}
			}
			if i==6{

				if checkWithPartName("sections/mon.txt" , partname)==true{

					checkifExist=1
					
				}
			}
			if i==7{

				if checkWithPartName("sections/tue.txt" , partname)==true{

					checkifExist=1
					
				}
			}
			if i==8{

				if checkWithPartName("sections/wed.txt" , partname)==true{
					
					checkifExist=1
					
				}
			}

			if i==9{

				if checkWithPartName("sections/thu.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==10{

				if checkWithPartName("sections/fri.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==11{

				if checkWithPartName("sections/u300.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==12{

				if checkWithPartName("sections/u600.txt" , partname)==true{
					checkifExist=1
					
				}
			}

			if i==13{

				if checkWithPartName("sections/u1600.txt" , partname)==true{
					checkifExist=1
					
				}
			}
			if i==14{

				if checkWithPartName("sections/k.txt" , partname)==true{
					checkifExist=1
					
				}
			}
						
			
		}//end of for

		if checkifExist==0{
			return false
		}else{
			return true
		} 


}

func partialName(){

	var partName string

	fmt.Println("you are in partialName : ")

	fmt.Printf("Enter partial Name : ")
	fmt.Scanf("%s",&partName)

	if checkpartialInSection(partName)==false{
		fmt.Println("Name not found ! ")
	}


}

func checkandPrintsectionstudent(sectFileinfo string){

	file , err := os.Open(sectFileinfo)
	if err!=nil{
		log.Fatal(err)
	}

	defer file.Close()

	scanner:=bufio.NewScanner(file)

	for scanner.Scan(){
		readName:=scanner.Text()
		checkinSection(readName,2,sectFileinfo)
	}

}

func sectionName(){

	var choice int32
	var sectName string
	// fmt.Println("you are in sectionName")

	fmt.Printf("If you want see the list of sections press-> 1 otherwise press->0 : ")

	fmt.Scanf("%d",&choice)

	if choice==1{
		fmt.Println("1. 1st\n2. 2nd\n3. 3rd\n4. 4th\n5. 5th\n6. mon\n7. tue\n8. wed\n9. thu\n10. fri\n11. u300\n12. u600\n13. u1600\n14. k")
	}
	fmt.Printf("Enter Section name : ")

	fmt.Scanf("%s",&sectName)

	sectName=strings.ToLower(sectName)

	s:=[]string{"sections/",sectName,".txt"}

	wholeseectName:=strings.Join(s,"")


	fmt.Println("\n\nList of student of ",sectName," section...\n\n")


	checkandPrintsectionstudent(wholeseectName)


}

func main() {

	var choice int32

	choice=1

	for choice >=1 {

		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("!-------------------Menu----------------------------------------!")
		fmt.Println("!------------1.input full name----------------------------------!")
		fmt.Println("!------------2.input partial name-------------------------------!")
		fmt.Println("!------------3.input section name-------------------------------!")
		fmt.Println("!-------------------4.Exit--------------------------------------!")
		fmt.Println("-----------------------------------------------------------------")

		fmt.Printf("Enter your choice : ")
		fmt.Scanf("%d",&choice)

		switch choice {

			case 1 : fullName()

			case 2 : partialName()

			case 3 : sectionName()

			case 4 : os.Exit(0)

			default : fmt.Println("PLEASE ENTER CORRECT CHOICE ! ")

		}// end of switch
	}// end of for
	
}// end of main
