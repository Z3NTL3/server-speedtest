package main

import (
	"context"
	test "go.jonnrb.io/speedtest/speedtestdotnet"
	"fmt"
	"go.jonnrb.io/speedtest/units"
	"time"
)

/* 
* Programmed by Z3NTL3
*/

const (
	timeout = 3
)

func DownloadSpeed(client *test.Client,ctx context.Context, done chan int) (){
	downloadSpeed := make(chan units.BytesPerSecond, 1000 * 4)

	_, err := client.Config(ctx)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	servers,err := client.LoadAllServers(ctx)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	bytes_second, err := servers[0].ProbeDownloadSpeed(ctx, client, downloadSpeed)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	download_speed := bytes_second * 8
	fmt.Println("Download Speed:" , download_speed)
	done <- 1
}

func UploadSpeed(client *test.Client,ctx context.Context, done chan int) (){
	uploadSpeed := make(chan units.BytesPerSecond, 1000 * 4)

	_, err := client.Config(ctx)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	servers,err := client.LoadAllServers(ctx)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	bytes_second, err := servers[0].ProbeUploadSpeed(ctx, client, uploadSpeed)
	if(err != nil){
		fmt.Println("Error:" , err)
		return
	}

	upload_speed := bytes_second * 8
	fmt.Println("Upload Speed:" , upload_speed)
	done <- 1
}


func main() {
	client := new(test.Client)
	done1 := make(chan int,1)
	done2 := make(chan int,1)

	ctx, cancel := context.WithDeadline(context.Background(),  time.Now().Add(time.Second * timeout))
	defer cancel()

	loop_i := []int{1,2}

	for i,_ := range loop_i {
		if i == 0 {
			go DownloadSpeed(client,ctx,done1)
		} else {
			go UploadSpeed(client,ctx,done2)
		}
	}
	<-done1
	<-done2
}
