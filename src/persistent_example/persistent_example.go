package main

import (
	"C"
	"gosol"
	"fmt"
	"time"
	"unsafe"
	"os"
)

func on_err(sess gosol.SESSION, err *gosol.ErrEvent) {
	fmt.Println("\nERROR EVENT:")
	fmt.Println("\tFNName: ",     err.FNName )
	fmt.Println("\tRetCcode: ",   err.RetCode )
	fmt.Println("\tRCString: ",   err.RCStr )
	fmt.Println("\tSubCode: ",    err.SubCode )
	fmt.Println("\tSCString: ",   err.SCStr )
	fmt.Println("\tRespCode: ",   err.RespCode )
	fmt.Println("\tErr String: ", err.ErrStr )
}


func dt2str(i int) string {
	switch i {
	case 1: return "TOPIC"
	case 2: return "QUEUE"
	}
	return "NONE"
}

func on_msg(sess gosol.SESSION, msg *gosol.MsgEvent) {
	fmt.Println("\nMESSAGE EVENT:")

	cstr    := (*C.char)(msg.Buffer)
	payload := C.GoStringN(cstr, C.int(msg.BufLen))
	
	fmt.Println("\tDestType: ", dt2str(msg.DestType) )
	fmt.Println("\tDestination: ", msg.Destination )
	fmt.Println("\tFlow: ", msg.Flow )
	fmt.Println("\tBuffer: ", payload )
	fmt.Println("\tBufLen: ", msg.BufLen )
	fmt.Println("\tMsgId: ", msg.MsgId )
	fmt.Println("\tReqId: ", msg.ReqId )
	fmt.Println("\tRedelivered: ", msg.Redelivered )
	fmt.Println("\tDiscard: ", msg.Discard )
}

func usage_exit() {
	fmt.Println("\tUSAGE:\n\t\t" , os.Args[0] , " <file.properties>\n")
	os.Exit(0)
}
func check_args() {
	if len(os.Args) < 2 {
		usage_exit()
	} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Println("\n\tProperties file ", os.Args[1], " not found.\n")
		usage_exit()
	}
}

func main() {

	check_args()

	payload  :=  "hello world"
	bytes    := []byte(payload)
	buffptr  :=  unsafe.Pointer(&bytes[0])
	buflen   :=  len(payload)


	sess := gosol.Init( gosol.MsgHandler(on_msg), gosol.ErrHandler(on_err), nil, nil )

	gosol.Connect( sess, os.Args[1] )

	gosol.BindQueue( sess, "q0",  gosol.STORE_FWD, gosol.AUTO_ACK )

	for i := 0; i < 3; i++ {
		fmt.Println("\nSending persistent msg: " , payload , " len:" , buflen)
		gosol.SendPersistent( sess, "q0", gosol.QUEUE, buffptr, buflen )
		gosol.SendPersistent( sess, "q0/mapped/1", gosol.TOPIC, buffptr, buflen )
	}

	time.Sleep( 2 * time.Second )

	gosol.UnbindQueue( sess, "q0" )

	gosol.Disconnect( sess )
}

