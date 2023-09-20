package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("检测LNMP v2.0 被投毒检测\n")
	fmt.Println("Step 1:检查操作系统")
	cmd := exec.Command("cat", "/etc/os-release")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("无法获取操作系统信息，退出:")

	} else {
		osInfo := string(output)
		if strings.Contains(osInfo, "CentOS") {
			fmt.Println("当前系统是CentOS")
		} else {
			fmt.Println("当前系统不是CentOS,暂时安全")
			return
		}

	}

	fmt.Println("\nStep 2:检查定时任务")
	filename := "/usr/sbin/crond"
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("无法获取定时文件信息，说明没有定时任务，安全！如果还担心，可执行crontab -l 查看")
	} else {
		modifiedTime := fileInfo.ModTime()
		fmt.Printf("定时任务最后一次修改时间： %+v", modifiedTime)
	}

	// --------------------------------------------------------------------------------------
	//          安装包
	// --------------------------------------------------------------------------------------
	fmt.Println("\nStep 3:检查安装包md5")
	fmt.Println("正常文件MD5：（下载版）1236630dcea1c5a617eb7a2ed6c457ed （全量包,疑似）6811dcbdb8b689f869c51f6cc9a34247 ")
	filename = "/root/lnmp2.0.tar.gz"
	_, err = os.Stat(filename)
	if err == nil {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		defer file.Close()

		hash := md5.New()
		if _, err := io.Copy(hash, file); err != nil {
			fmt.Println("无法计算文件的MD5值，请自行计算比对。:", err)
		} else {
			md5Hash := hash.Sum(nil)
			md5String := hex.EncodeToString(md5Hash)
			fmt.Println("文件的MD5值:", md5String)
			if strings.Compare(md5String, "1236630dcea1c5a617eb7a2ed6c457ed") == 0 {
				fmt.Println("安装包MD5与官方下载版一致。安全。检查结束")
			} else {
				fmt.Println("安装包MD5与官方不一致！！！ 请结合定时任务情况做评估。")
			}
			if strings.Compare(md5String, "6811dcbdb8b689f869c51f6cc9a34247") == 0 {
				fmt.Println("安装包应该是全下载版，理论安全。")
			} else {
				fmt.Println("安装包MD5与官方不一致！！！ 请结合定时任务情况做评估。")
			}
		}
	} else if os.IsNotExist(err) {
		fmt.Println("没有发现安装包....")
		fmt.Println("如果安装包存放在别处，请执行以下命令后再重新检测。\n\n    mv 安装包路径 /root/lnmp2.0.tar.gz \n ")
	}

	// --------------------------------------------------------------------------------------
	//          检索lnmp.sh
	// --------------------------------------------------------------------------------------
	fmt.Println("\nStep 4:检查是否存在恶意文件lnmp.sh ")
	var checkApp string
	fmt.Print("是否全盘扫码恶意文件（耗时很久），如果是，输入数字1：")
	fmt.Scan(&checkApp)
	if strings.Compare("1", checkApp) == 0 {
		fmt.Println("您选中全盘扫描，请耐心等待！")
		cmd = exec.Command("find", "/", "-name", "lnmp.sh")
		output, err = cmd.Output()
		if err != nil {
			fmt.Println("执行命令时出错:", err)
			return
		}
		files := string(output)
		if len(files) > 0 {
			fmt.Println("中招了，恶意程序路径:", files)
		} else {
			fmt.Println("没找到恶意程序lnmp.sh,安全")
		}
	} else {
		fmt.Println("您没选全盘扫描")
	}

	fmt.Println("检查结束")
}
