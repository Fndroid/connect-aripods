package main

import (
	"encoding/base64"
	"log"
	"os"
	"os/exec"

	"github.com/caseymrm/go-pmset"
)

var script = "RmFzZFVBUyAxLjEwMS4xMA4AAAAED///AAEAAgADAf//AAANAAEAAWsAAAAAAAAABAIABAACAAUABg0ABQADeAAAAAAAC//+AAf//QH//gAADQAHAAI0AAAAAQAF//wACAr//AAECmZybWsNAAgAAW0AAAADAAQACQ4ACQABsQAKEQAKABYASQBPAEIAbAB1AGUAdABvAG8AdABoAv/9AAACAAYAAgALAAwNAAsAA3gAAAALABf/+wAN//oB//sAAA0ADQABMgABAA0AEP/5Cv/5AAQKb3NheAL/+gAAAgAMAAIADgAPDQAOAANsAAIAAAAA//j/9//2Af/4AAAB//cAAAH/9gAAAgAPAAIAEAARDQAQAANsAAIAAAAA//X/9P/zAf/1AAAB//QAAAH/8wAAAgARAAIAEgATDQASAANsAAIAAAAA//L/8f/wAf/yAAAB//EAAAH/8AAAAgATAAIAFAAVDQAUAAJpAAAAFwAaABYAFw0AFgADSQAAAAAAAP/vABj/7gv/7wAwMAAWZ2V0Zmlyc3RtYXRjaGluZ2RldmljZQAWZ2V0Rmlyc3RNYXRjaGluZ0RldmljZQIAGAACABn/7Q0AGQABbwAAAAAAAP/sC//sABgwAApkZXZpY2VuYW1lAApkZXZpY2VOYW1lAv/tAAAC/+4AAA0AFwADWAAAAAAALQAa/+sAGw0AGgAEWgABABgAKAAcAB3/6v/pDQAcAAJFAAAAGAAfAB4AHw0AHgADbAAFABgAHQAg/+j/5w0AIAACYwAAABgAHQAhACINACEAAm4AAwAYABsAIwAkDQAjAAFvAAAAGQAb/+YL/+YAHjAADW5hbWVvcmFkZHJlc3MADW5hbWVPckFkZHJlc3MNACQAAW8AAAAYABn/5Qv/5QAKMAAGZGV2aWNlAAANACIAAW0AAAAbABz/5Ar/5AAEClRFWFQB/+gAAAH/5wAADQAfAAFvAAAAHQAe/+ML/+MAGDAACmRldmljZW5hbWUACmRldmljZU5hbWUNAB0AAUwAAAAiACQAJQ0AJQABbwAAACIAI//iC//iAAowAAZkZXZpY2UAAAL/6gAAAf/pAAAL/+sACjAABmRldmljZQAADQAbAANsAAUAAwAMACb/4f/gDQAmAAJjAAAAAwAMACcAKA0AJwACbgADAAMACgApACoNACkAA0kAAAAGAAr/3//e/90L/98AHjAADXBhaXJlZGRldmljZXMADXBhaXJlZERldmljZXMC/94AAAL/3QAADQAqAAJuAAMAAwAGACsALA0AKwABbwAAAAQABv/cC//cACYwABFpb2JsdWV0b290aGRldmljZQARSU9CbHVldG9vdGhEZXZpY2UNACwAAW0AAAADAAT/2wr/2wAIC21pc2NjdXJhDQAoAAFtAAAACgAL/9oK/9oABApsaXN0Af/hAAAB/+AAAAIAFQACAC0ALg0ALQADbAACAAAAAP/Z/9j/1wH/2QAAAf/YAAAB/9cAAAIALgACAC8AMA0ALwACaQAAABsAHgAxADINADEAA0kAAAAAAAD/1gAz/9UL/9YAHDAADHRvZ2dsZWRldmljZQAMdG9nZ2xlRGV2aWNlAgAzAAIANP/UDQA0AAFvAAAAAAAA/9ML/9MACjAABmRldmljZQAAAv/UAAAC/9UAAA0AMgABawAAAAAALQA1AgA1AAIANgA3DQA2AAJyAAAAAAAJADgAOQ0AOAACbgAAAAAABwA6ADsNADoAATEAAAAFAAf/0gr/0gAECnN0cnENADsAA2wABQAAAAUAPP/R/9ANADwAAmMAAAAAAAUAPQA+DQA9AAJuAAMAAAADAD8AQA0APwABbwAAAAEAA//PC//PAB4wAA1uYW1lb3JhZGRyZXNzAA1uYW1lT3JBZGRyZXNzDQBAAAFvAAAAAAAB/84L/84ACjAABmRldmljZQAADQA+AAFtAAAAAwAE/80K/80ABApURVhUAf/RAAAB/9AAAA0AOQABbwAAAAAAAP/MC//MACQwABBxdW90ZWRkZXZpY2VuYW1lABBxdW90ZWREZXZpY2VOYW1lAgA3AAIAQQBCDQBBAANsAAIACgAK/8v/yv/JAf/LAAAB/8oAAAH/yQAAAgBCAAIAQwBEDQBDAARaAAAACgAcAEUARv/I/8cNAEUAAUgAAAAKABAARw0ARwADbAAFAAoADwBI/8b/xQ0ASAACYwAAAAoADwBJAEoNAEkAAm4AAwAKAA0ASwBMDQBLAAFvAAAACwAN/8QL/8QAGjAAC2lzY29ubmVjdGVkAAtpc0Nvbm5lY3RlZA0ATAABbwAAAAoAC//DC//DAAowAAZkZXZpY2UAAA0ASgABbQAAAA0ADv/CCv/CAAQKYm9vbAH/xgAAAf/FAAANAEYAAm4AAwATABgATQBODQBNAANJAAAAFAAY/8H/wP+/C//BACAwAA5vcGVuY29ubmVjdGlvbgAOb3BlbkNvbm5lY3Rpb24C/8AAAAL/vwAADQBOAAFvAAAAEwAU/74L/74ACjAABmRldmljZQAAAv/IAAAB/8cAAAIARAACAE8AUA0ATwADbAACAB0AHf+9/7z/uwH/vQAAAf+8AAAB/7sAAAIAUAACAFEAUg0AUQADSQACAB0AJP+6AFP/uQr/ugAYLnN5c29leGVjVEVYVP//gAAAAAAAVEVYVA0AUwACYgAAAB0AIABUAFUNAFQAAW0AAAAdAB4AVg4AVgABsQBXEQBXAEgALwB1AHMAcgAvAGwAbwBjAGEAbAAvAGIAaQBuAC8AUwB3AGkAdABjAGgAQQB1AGQAaQBvAFMAbwB1AHIAYwBlACAALQBzACANAFUAAW8AAAAeAB//uAv/uAAkMAAQcXVvdGVkZGV2aWNlbmFtZQAQcXVvdGVkRGV2aWNlTmFtZQL/uQAAAgBSAAIAWP+3DQBYAAFMAAAAJQAtAFkNAFkAAmIAAAAlACwAWgBbDQBaAAFtAAAAJQAmAFwOAFwAAbEAXREAXQAWAEMAbwBuAG4AZQBjAHQAaQBuAGcAIA0AWwADbAAFACYAKwBe/7b/tQ0AXgACYwAAACYAKwBfAGANAF8AAm4AAwAmACkAYQBiDQBhAAFvAAAAJwAp/7QL/7QAHjAADW5hbWVvcmFkZHJlc3MADW5hbWVPckFkZHJlc3MNAGIAAW8AAAAmACf/swv/swAKMAAGZGV2aWNlAAANAGAAAW0AAAApACr/sgr/sgAEClRFWFQB/7YAAAH/tQAAAv+3AAACADAAAgBjAGQNAGMAA2wAAgAAAAD/sf+w/68B/7EAAAH/sAAAAf+vAAACAGQAAgBlAGYNAGUAA2wAAgAAAAD/rv+t/6wB/64AAAH/rQAAAf+sAAACAGYAAgBnAGgNAGcAA2wAAgAAAAD/q/+q/6kB/6sAAAH/qgAAAf+pAAACAGgAAgBp/6gNAGkAAmkAAAAfACIAagBrDQBqAANJAAIAAAAA/6cAbP+mCv+nABguYWV2dG9hcHBudWxsAACAAAAAkAAqKioqDQBsAAFvAAAAAAAA/6UL/6UACDAABGFyZ3MAAAL/pgAADQBrAAFrAAAAAAATAG0CAG0AAgBuAG8NAG4AAnIAAAAAAAYAcABxDQBwAAJuAAAAAAAEAHIAcw0AcgACNAAAAAEABP+kAHQK/6QABApjb2JqDQB0AAFtAAAAAgAD/6MD/6MAAQ0AcwABbwAAAAAAAf+iC/+iAAgwAARhcmdzAAANAHEAAW8AAAAAAAD/oQv/oQAaMAALYWlycG9kc25hbWUAC0FpclBvZHNOYW1lAgBvAAIAdf+gDQB1AAFMAAAABwATAHYNAHYAA0kAAAAHABL/nwB3/54L/58AHDAADHRvZ2dsZWRldmljZQAMdG9nZ2xlRGV2aWNlAgB3AAIAeP+dDQB4AANJAAAACAAO/5wAef+bC/+cADAwABZnZXRmaXJzdG1hdGNoaW5nZGV2aWNlABZnZXRGaXJzdE1hdGNoaW5nRGV2aWNlAgB5AAIAev+aDQB6AAFvAAAACQAK/5kL/5kAGjAAC2FpcnBvZHNuYW1lAAtBaXJQb2RzTmFtZQL/mgAAAv+bAAAC/50AAAL/ngAAAv+gAAAC/6gAAA4AAgAADxAAAwAK/5gAewB8AH0AfgB/AID/l/+W/5UB/5gAABAAewAI/5T/k/+S/5H/kP+P/47/jQr/lAAECnBpbXIL/5MAMDAAFmdldGZpcnN0bWF0Y2hpbmdkZXZpY2UAFmdldEZpcnN0TWF0Y2hpbmdEZXZpY2UL/5IAHDAADHRvZ2dsZWRldmljZQAMdG9nZ2xlRGV2aWNlCv+RABguYWV2dG9hcHBudWxsAACAAAAAkAAqKioqC/+QABowAAthaXJwb2RzbmFtZQALQWlyUG9kc05hbWUB/48AAAH/jgAAAf+NAAAOAHwAAgT/jACBA/+MAAIOAIEAAgAAggCDBgCCAAP/iwCE/4oK/4sABApjb2JqDgCEAAEUAIUOAIUAAxgAAP+JAAkK/4kABApmcm1rBv+KAAAGAIMAA/+IAIb/hwr/iAAECmNvYmoOAIYAARQAhw4AhwACFgAA/4YK/4YABApvc2F4Bv+HAAAOAH0ABxD/hQAX/4T/gwCIAIn/ggv/hQAwMAAWZ2V0Zmlyc3RtYXRjaGluZ2RldmljZQAWZ2V0Rmlyc3RNYXRjaGluZ0RldmljZQ7/hAACBP+BAIoD/4EAAQ4AigABAP+AC/+AABgwAApkZXZpY2VuYW1lAApkZXZpY2VOYW1lAv+DAAAQAIgAAv9//34L/38AGDAACmRldmljZW5hbWUACmRldmljZU5hbWUL/34ACjAABmRldmljZQAAEACJAAn/ff98/3v/ev95/3j/d/92/3UK/30ACAttaXNjY3VyYQv/fAAmMAARaW9ibHVldG9vdGhkZXZpY2UAEUlPQmx1ZXRvb3RoRGV2aWNlC/97AB4wAA1wYWlyZWRkZXZpY2VzAA1wYWlyZWREZXZpY2VzCv96AAQKbGlzdAr/eQAECmtvY2wK/3gABApjb2JqCv93ABguY29yZWNudGUqKioqAAAAAAAAEAAqKioqC/92AB4wAA1uYW1lb3JhZGRyZXNzAA1uYW1lT3JBZGRyZXNzCv91AAQKVEVYVBH/ggAuFwAs4OEsaisAAuMmW+TlbAwABmtoGwABoecs6CagCB0AB6EPWQADaFtPWf/qDw4AfgAHEP90ADL/c/9yAIsAjP9xC/90ABwwAAx0b2dnbGVkZXZpY2UADHRvZ2dsZURldmljZQ7/cwACBP9wAI0D/3AAAQ4AjQABAP9vC/9vAAowAAZkZXZpY2UAAAL/cgAAEACLAAL/bv9tC/9uAAowAAZkZXZpY2UAAAv/bQAkMAAQcXVvdGVkZGV2aWNlbmFtZQAQcXVvdGVkRGV2aWNlTmFtZRAAjAAJ/2z/a/9q/2n/aP9nAFb/ZgBcC/9sAB4wAA1uYW1lb3JhZGRyZXNzAA1uYW1lT3JBZGRyZXNzCv9rAAQKVEVYVAr/agAECnN0cnEL/2kAGjAAC2lzY29ubmVjdGVkAAtpc0Nvbm5lY3RlZAr/aAAECmJvb2wL/2cAIDAADm9wZW5jb25uZWN0aW9uAA5vcGVuQ29ubmVjdGlvbgr/ZgAYLnN5c29leGVjVEVYVP//gAAAAAAAVEVYVBH/cQAuoOAs4SbiLEWxT6DjLOQmCx0ACqBqKwAFWQADaE/moSVqDAAHT+ig4CzhJiUPDw4AfwAHEP9lAGv/ZP9jAI4Aj/9iCv9lABguYWV2dG9hcHBudWxsAACAAAAAkAAqKioqC/9kAAgwAARhcmdzAAAC/2MAABAAjgAB/2EL/2EACDAABGFyZ3MAABAAjwAE/2D/X/9e/10K/2AABApjb2JqC/9fABowAAthaXJwb2RzbmFtZQALQWlyUG9kc05hbWUL/14AMDAAFmdldGZpcnN0bWF0Y2hpbmdkZXZpY2UAFmdldEZpcnN0TWF0Y2hpbmdEZXZpY2UL/10AHDAADHRvZ2dsZWRldmljZQAMdG9nZ2xlRGV2aWNlEf9iABSg4GsvRdFPKirBaysAAmsrAAMPDw4AgAABsQCQEQCQACoARgBuAGQAcgBvAGkAZAAnAHMAIABBAGkAcgBQAG8AZABzACAAUAByAG8B/5cAAAH/lgAAAf+VAABhc2NyAAEADPre3q0="

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Println("Starting to detect audio status...")
	args := os.Args[1:]
	if len(args) != 1 {
		panic("Usage: connect-airpods \"Fndroid's AriPods\"")
	}

	scriptBin, err := base64.StdEncoding.DecodeString(script)
	check(err)

	err = os.WriteFile("./script.scpt", scriptBin, 0644)
	check(err)

	channel := make(chan pmset.AssertionChange)
	go func() {
		for change := range channel {
			if change.Action == "Created" {
				if change.Pid.Name == "com.apple.Music.playback" || change.Pid.Name == "Playing audio" {
					log.Println("Audio is playing")
					cmd := exec.Command("osascript", "./script.scpt", args[0])
					cmd.Run()
				}
			}
		}
	}()
	pmset.SubscribeAssertionChangesAndRun(channel)
}
