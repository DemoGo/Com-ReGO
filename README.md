# Com-ReGO

SOME FEATURES

A program in Go that takes name of student and reads the result in result file and tells student win or not.
Takes the input as section name and show the status of every student.
Takes the input as partial name and suggest full name and status.
It also gives the information of student that also study in other section.


NOTE: You should download the sections,lme,patsy,casis folders and save the code and name it final.go in same directory 
      then you test
      
      link to download all files=https://drive.google.com/folderview?id=0BwKr2qfLihCUOFF1TVhhaVEycGs&usp=sharing


1. Creating an envrionment

	1.1	Install golang package according to system configuration. URL : https://golang.org/dl/
	1.2	Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go.
		 For example:
			tar -C /usr/local -xzf go1.6.2.linux-386.tar.gz
			Add /usr/local/go/bin to the PATH environment variable. 
			You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

			export PATH=$PATH:/usr/local/go/bin

2. How to run program

	2.1	In your go workspace save all file and directories:
						    a. mainProject.go
						    b. casis
						    c. lme
						    d. patsy
						    e. sections

	2.2 	In terminal first go to that workspace and type command : go run mainProject.go


3. Limitation of Program

	3.1 	Enter name with.
	
	3.2 	Enter sections name as mentioned.
	
	3.3 	It might show errors sometime because the program opens to much file in system but i tested
		at ubuntu 12.04LTS it is working smoothly.

4. Package involve in program
	
	4.1 bufio	Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
			creating another object (Reader or Writer) that also implements the interface but 
			provides buffering and some help for textual I/O. 

	4.2 fmt		Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 

	4.3 log		Package log implements a simple logging package.It defines a type, Logger, with methods 
			for formatting output.That logger writes to standard error and prints the date and time 
			of each logged message.

	4.4 os 		Package os provides a platform-independent interface to operating system functionality. 

	4.5 strings	Package strings implements simple functions to manipulate UTF-8 encoded strings. 

5. Testing environmet 

	I tested this on 32-bit , 2.3 GHz system (amd processor) and at operating system 12.04LTS
	
		

