#IfWinActive, Path of Exile
; #NoEnv
SetBatchLines, -1
#Include ./lib/Websocket.ahk

client:=new Client("ws://127.0.0.1:9527")
return

f2::
if (!client.toggle){
	client.toggle:=1
	SoundPlay, ./on.mp3
}else{
	client.toggle:=0
	SoundPlay, ./off.mp3
}


class Client extends WebSocket
{
	static toggle:=1

	OnOpen(Event)
	{
		SoundPlay, ./on.mp3
		; MsgBox,  "Connected to local server port 9527"
		; InputBox, Data, WebSocket, Enter some text to send through the websocket.
		; this.Send(Data)
	}

	OnMessage(Event)
	{
		message:=Event.data
		if (this.toggle){
			SendInput, {Enter}
			SendInput, %message%
			SendInput, {Enter}
		}

		; MsgBox, % "Received Data: " message
		; this.Close()
	}

	OnClose(Event)
	{
		; MsgBox, Websocket Closed
		SoundPlay, ./error.mp3
		this.Disconnect()
	}

	OnError(Event)
	{
		SoundPlay, ./error.mp3
		sleep, 1000
		; MsgBox, Websocket Error
		ExitApp
	}

	__Delete()
	{
		SoundPlay, ./error.mp3
		sleep, 1000
		; MsgBox, Exiting
		ExitApp
	}
}