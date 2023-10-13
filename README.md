
### HackNight Details : 

#### Decscription : 

A modern-ish looking system monitor that uses what we would call the "next-gen" languages rust and go! It's an alternative to htop and bpytop.

#### Setup : 

- You would need golang and Rust toolchains installed before hand. And most of all, this project works only in linux.
- You can execute `build.sh` to install this application from local copy (this will build with all your changes as well)
- `sudo systemctl start sysperf.service`
- `sysperf`. It is possible the font of the terminal window it too large, hence try making the fonts lower to render the application (This is actually and issue I'd like you to solve :-)

#### Maintainer : 

- P K Navin Shrinivas (@NavinShrinivas)
- Mukund Deepak (@mukunddeepak)

# SysPerf 

A modern-ish looking system monitor that uses what we would call the "next-gen" languages rust and go!

## Features 
- Uses a client server  model 
- Uses GRPC to connect the both

## Downsides : 
- Uses unsafe rust :( 
- Over complicated for a simple task, but tbf this project was done only to follow the guidelines of SE as part of the course in our Univerity!

## Where could it be used:
- Due to it being built on rust and go compilation and setting up the system is done very quickly.
- Due to its ightweight nature it will not take over too many computer resources to obtain the necessary information about the system.
## How to use : 

> Note : Needs Rust and go toochains installed before hand


![image](https://user-images.githubusercontent.com/72858215/227838766-70c9314f-7060-418e-924c-f8d85c10f859.png)
