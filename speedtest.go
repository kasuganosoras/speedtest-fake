package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
	"os"
)

func mt_rand(min,max int64) int64{
    maxBigInt:=big.NewInt(max)
    i,_:=rand.Int(rand.Reader,maxBigInt)
    if i.Int64()<min{
        mt_rand(min,max)    
    }
    return i.Int64()
}

func cv(number int64) int {
	tempstr := strconv.FormatInt(number, 10);
	newint, err := strconv.Atoi(tempstr);
	if err != nil {
		//
	}
	return newint;
}

func vc(number string) int64 {
	newint, err := strconv.ParseInt(number, 10, 64);
	if err != nil {
		//
	}
	return newint;
}

func main() {
	if len(os.Args) < 11 {
		fmt.Println("Invalid arguments")
		fmt.Println("")
		fmt.Println("Usage: command <area> <ip> <isp> <server_area> <radius> <randomping_min> <randomping_max> <upload_min> <upload_max> <download_min> <download_max>")
		os.Exit(0)
	} else {
		area := os.Args[1]; 					                                                      // 地区
		ip := os.Args[2]; 								                                                  // IP
		isp := os.Args[3]; 				                                                          // 测速运营商
		server_area := os.Args[4]; 						                                              // 测速点位置
		radius := os.Args[5]; 									                                            // 测速点距离
		randomping := strconv.FormatInt(mt_rand(vc(os.Args[6]), vc(os.Args[7])), 10); 			// 测速点延迟，10 是最小值，100 是最大值
		randomping = randomping + "." + strconv.FormatInt(mt_rand(10, 999), 10); 	          // 测速点延迟随机小数点
		upload := strconv.FormatInt(mt_rand(vc(os.Args[8]), vc(os.Args[9])), 10); 					// 随机上传速度，1000 最小值，2000 最大值
		download := strconv.FormatInt(mt_rand(vc(os.Args[10]), vc(os.Args[11])), 10); 			// 随机下载速度，1000 最小值，2000 最大值
		upload = upload + "." + strconv.FormatInt(mt_rand(10, 99), 10); 			              // 随机小数点
		download = download + "." + strconv.FormatInt(mt_rand(10, 99), 10); 		            // 随机小数点
		fmt.Println("Retrieving speedtest.net configuration...");
		time.Sleep(3 * time.Second);
		fmt.Println("Testing from " + area + " (" + ip + ")...");
		time.Sleep(800000 * time.Microsecond);
		fmt.Println("Retrieving speedtest.net server list...");
		time.Sleep(2 * time.Second);
		fmt.Println("Selecting best server based on ping...");
		time.Sleep(8 * time.Second);
		fmt.Println("Hosted by " + isp + " (" + server_area + ") [" + radius + " km]: " + randomping + " ms");
		time.Sleep(50000 * time.Microsecond);
		fmt.Print("Testing download speed");
		for i := 0; i < 80; i++ {
			fmt.Print(".");
			time.Sleep(time.Duration(cv(mt_rand(5000, 50000))) * time.Microsecond);
		}
		fmt.Println("\nDownload: " + upload + " Mbit/s");
		time.Sleep(80000 * time.Microsecond);
		fmt.Print("Testing upload speed");
		for a := 0; a < 96; a++ {
			fmt.Print(".");
			time.Sleep(time.Duration(cv(mt_rand(5000, 50000))) * time.Microsecond);
		}
		fmt.Println("\nUpload: " + download + " Mbit/s");
	}
}
