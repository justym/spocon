# spocon
Dead simply spotify cli controller

# Why 
I just wanted to controll playback from cli

# Setup 
1. Please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
2. Please set redirect uri to http://127.0.0.1:14565/oauth/callback 
3. Please set these values like below.
```
export SPOTIFY_ID=your_client_id
export SPOTIFY_KEY=your_client_key
```
4. Please run the commands below

```
go get github.com/justym/spocon
cd $GOPATH/src/github.com/justym/spocon
go build
```

or 

Please run ```git clone [this repo]```. After that, you can use this app from ./bin/spocon

If you are in a trouble, please see caution part in this README

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
- This cli-app uses Spotify Web API. So, please check their [web page](https://developer.spotify.com/documentation/web-api/) carefully.
- This app uses that Web API through Authorization code flow (OAuth2.0).
- This cli-app uses this package [github.com/nmrshll/oauth2-noserver](https://github.com/nmrshll/oauth2-noserver) to do authentication rapidly. So please see this package carefully too.

# Finally
Welcome ...
- Code review :dizzy:
- Issue :dizzy:
- Pull Request:dizzy:


