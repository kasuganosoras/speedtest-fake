<?php
$area = "China Telecom Guangdong"; 					// 地区
$ip = "183.15.xxx.xx"; 								// IP
$isp = "W Professional Services Limited"; 			// 测速运营商
$server_area = "New Territories"; 					// 测速点位置
$radius = "18.14"; 									// 测速点距离
$randomping = mt_rand(10, 100); 					// 测速点延迟，10 是最小值，100 是最大值
$randomping = "{$randomping}." . mt_rand(10, 999); 	// 测速点延迟随机小数点
$upload = mt_rand(10000, 20000); 					// 随机上传速度，1000 最小值，2000 最大值
$download = mt_rand(10000, 20000); 					// 随机下载速度，1000 最小值，2000 最大值
$upload = "{$upload}." . mt_rand(10, 99); 			// 随机小数点
$download = "{$download}." . mt_rand(10, 99); 		// 随机小数点
echo "Retrieving speedtest.net configuration...\n";
sleep(3);
echo "Testing from {$area} ({$ip})...\n";
usleep(800000);
echo "Retrieving speedtest.net server list...\n";
sleep(2);
echo "Selecting best server based on ping...\n";
sleep(mt_rand(5, 10));
echo "Hosted by {$isp} ({$server_area}) [{$radius} km]: {$randomping} ms\n";
usleep(50000);
echo "Testing download speed";
for($i = 0;$i < 80;$i++) {
	echo ".";
	usleep(mt_rand(5000, 50000));
}
echo "\nDownload: {$upload} Mbit/s\n";
usleep(80000);
echo "Testing upload speed";
for($i = 0;$i < 96;$i++) {
	echo ".";
	usleep(mt_rand(5000, 50000));
}
echo "\nUpload: {$download} Mbit/s\n";
