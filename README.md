# spocon
Dead simply spotify cli controller

# Why 
I just wanted to controll playback from cli

# Setup 
1. Please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
2. Please set redirect uri to http://127.0.0.1:14565/oauth/callback 
2. Please set the values into ``` .env ``` file
3. Please run the commands below

```
go get github.com/justym/spocon
cd $GOPATH/src/github.com/justym/spocon
go build
```
If you are in a trouble, please watch caution part

# Usage
```
spocon [sub cmds]
```

## Sub cmmands list

- start
  - start command can play/resume your spotify playback
  - :construction: Mainly, please use this to resume. Now: play function is something wrong :construction:
- pause
  - pause command can pasue your spotify playback
- next
  - next command can listen next song
- prev
  - prev command can listen previous song
- status 
  - status command can get your currently playback song name
- help
  - help command can see short help message 

# Thanksfull Package used 
- github.com/spf13/cobra
- github.com/nmrshll/oauth2-noserver

# Caution
- This cli-app use spotify web api. So, please check their web page carefully.
- Spotify Premium user can connect to Spotiy Web API through Authorization code flow (OAuth2.0).
- Spotify Web API restrict user access limit (60 tiems / per an hour) 
- This cli-app use this package [github.com/nmrshll/oauth2-noserver](https://github.com/nmrshll/oauth2-noserver) to do authentication rapidly. So please see this package carefully too.

# Finally
:dizzy: Pull Request welcome :dizzy:
Code review welcome :dizzy:
Issue welcome :dizzy:
