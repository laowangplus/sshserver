package main

// Forward from local port 9000 to remote port 9999
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
)

var (
	username         = "saijunchen"
	password         = "bgEFT961x1d/uq+vdVi0VhdpLAMrm43f"
	serverAddrString = "106.52.224.60:22"

	localAddrString    = "localhost:62975"
	remoteAddrString   = "bj-cdb-m3cl1jfs.sql.tencentcdb.com:62975"
	localAddrString2   = "localhost:61365"
	remoteAddrString2  = "bj-cdb-6blgboli.sql.tencentcdb.com:61365"
	localAddrString3   = "localhost:61207"
	remoteAddrString3  = "sh-cdb-ojdb9hns.sql.tencentcdb.com:61207"
	localAddrString5   = "localhost:61990"
	remoteAddrString5  = "mysql-log-for-outer-select.xeknow.com:61990"
	localAddrString6   = "localhost:62972"
	remoteAddrString6  = "bj-cdb-jxuuzz8q.sql.tencentcdb.com:62972"
	localAddrString7   = "localhost:62992"
	remoteAddrString7  = "bj-cdb-35o08k06.sql.tencentcdb.com:62992"
	localAddrString8   = "localhost:61989"
	remoteAddrString8  = "bj-cdb-eqx67j7k.sql.tencentcdb.com:61989"
	localAddrString9   = "localhost:3306"
	remoteAddrString9  = "10.66.163.189:3306"
	localAddrString10  = "localhost:61581"
	remoteAddrString10 = "bj-cdb-3gza49ze.sql.tencentcdb.com:61581"

	//用户灾备
	localAddrString11  = "localhost:61092"
	remoteAddrString11 = "bj-cdb-hidruum8.sql.tencentcdb.com:61092"
	localAddrString12  = "localhost:61105"
	remoteAddrString12 = "bj-cdb-mndes97k.sql.tencentcdb.com:61105"
	localAddrString13  = "localhost:61082"
	remoteAddrString13 = "bj-cdb-b689sf7s.sql.tencentcdb.com:61082"
	localAddrString14  = "localhost:61104"
	remoteAddrString14 = "bj-cdb-7rw961rc.sql.tencentcdb.com:61104"
	localAddrString15  = "localhost:61057"
	remoteAddrString15 = "bj-cdb-cgaqqqpw.sql.tencentcdb.com:61057"
	localAddrString16  = "localhost:61431"
	remoteAddrString16 = "bj-cdb-qhzbtd8s.sql.tencentcdb.com:61431"
)

func forward(localConn net.Conn, config *ssh.ClientConfig, serverAddr_agement, remoteAddr_agement string) {
	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", serverAddr_agement, config)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}
	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", remoteAddr_agement)
	// Copy localConn.Reader to sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
	// Copy sshConn.Reader to localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}

func execRoot() {
	fmt.Println("一分钟一次检测")
	testDb1()
	testDb2()
	testDb3()
	testDb4()
	testDb5()
	testDb6()
	testDb7()
	testDb8()
	testDb9()
	testDb10()
	testDb11()
	testDb12()
	testDb13()
	testDb14()
	fmt.Println("检测完成")
}

func main() {
	go sshForward(localAddrString, serverAddrString, remoteAddrString)
	go sshForward(localAddrString2, serverAddrString, remoteAddrString2)
	go sshForward(localAddrString3, serverAddrString, remoteAddrString3)
	go sshForward(localAddrString5, serverAddrString, remoteAddrString5)
	go sshForward(localAddrString6, serverAddrString, remoteAddrString6)
	go sshForward(localAddrString7, serverAddrString, remoteAddrString7)
	go sshForward(localAddrString8, serverAddrString, remoteAddrString8)
	go sshForward(localAddrString9, serverAddrString, remoteAddrString9)
	go sshForward(localAddrString10, serverAddrString, remoteAddrString10)
	go sshForward(localAddrString11, serverAddrString, remoteAddrString11)
	go sshForward(localAddrString12, serverAddrString, remoteAddrString12)
	go sshForward(localAddrString13, serverAddrString, remoteAddrString13)
	go sshForward(localAddrString14, serverAddrString, remoteAddrString14)
	go sshForward(localAddrString15, serverAddrString, remoteAddrString15)
	go sshForward(localAddrString16, serverAddrString, remoteAddrString16)
	execRoot()

	c := cron.New()
	spec := "10 * * * * ?"
	c.AddFunc(spec, func() {
		execRoot()
	})
	c.Start()
	select {} //阻塞主线程停止
}

func sshForward(localAddr_agement, serverAddr_agement, remoteAddr_agement string) {
	// Setup SSH config (type *ssh.ClientConfig)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	// Setup localListener (type net.Listener)
	localListener, err := net.Listen("tcp", localAddr_agement)
	if err != nil {
		log.Fatalf("net.Listen failed: %v", err)
	}
	for {
		// Setup localConn (type net.Conn)
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}
		go forward(localConn, config, serverAddr_agement, remoteAddr_agement)
	}
}

func testDb1() {
	Db, err := sqlx.Open("mysql", "outer_select:5agEebp5eVu5J2Kc@tcp("+localAddrString+")/db_ex_alive?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【现网直播业务库】失败", err)
	}
	fmt.Println("链接【现网直播业务库】成功", localAddrString, "outer_select", "5agEebp5eVu5J2Kc")
}

func testDb2() {
	Db, err := sqlx.Open("mysql", "outer_select:Xiaoe@select20200107@tcp("+localAddrString2+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【现网直播日志数据库】失败", err)
	}
	fmt.Println("链接【现网直播日志数据库】成功", localAddrString2, "outer_select", "Xiaoe@select20200107")
}

func testDb3() {
	Db, err := sqlx.Open("mysql", "outer_select:g5#wsB@gIytC05I6@tcp("+localAddrString3+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【sh-cdb-ojdb9hns.sql.tencentcdb.com】失败", err)
	}
	fmt.Println("链接【sh-cdb-ojdb9hns.sql.tencentcdb.com】成功", localAddrString3, "outer_select", "g5#wsB@gIytC05I6")
}

func testDb4() {
	Db, err := sqlx.Open("mysql", "outer_select:2J1k0N4Fj9@pNWbt@tcp("+localAddrString5+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【mysql-log-for-outer-select.xeknow.com】失败", err)
	}
	fmt.Println("链接【mysql-log-for-outer-select.xeknow.com】成功", localAddrString5, "outer_select", "2J1k0N4Fj9@pNWbt")
}

func testDb5() {
	Db, err := sqlx.Open("mysql", "outer_select:0RdHT0km40uOSd@tcp("+localAddrString6+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【现网核心业务库】失败", err)
	}
	fmt.Println("链接【现网核心业务库】成功", localAddrString6, "outer_select", "0RdHT0km40uOSd")
}

func testDb6() {
	Db, err := sqlx.Open("mysql", "outer_select:MvLdRY@A5d69vjDl@tcp("+localAddrString7+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【现网次级业务库】失败", err)
	}
	fmt.Println("链接【现网次级业务库】成功", localAddrString7, "outer_select", "MvLdRY@A5d69vjDl")
}

func testDb7() {
	Db, err := sqlx.Open("mysql", "outer_select:xiaoe@dk6K965d6rtwss@tcp("+localAddrString8+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【现网直播统计数据库】失败", err)
	}
	fmt.Println("链接【现网直播统计数据库】成功", localAddrString8, "outer_select", "xiaoe@dk6K965d6rtwss")
}

func testDb8() {
	Db, err := sqlx.Open("mysql", "outer_select:5o+Jlq4NEGCp1EPx@tcp("+localAddrString10+")/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【订单数据库】失败", err)
	}
	fmt.Println("链接【订单数据库】成功", localAddrString10, "outer_select", "5o+Jlq4NEGCp1EPx")
}

func testDb9() {
	Db, err := sqlx.Open("mysql", "outer_select:MmHtJ1k26UaQXpix@tcp("+localAddrString11+")/db_ex_business?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【用户灾备1库】失败", err)
	}
	fmt.Println("链接【用户灾备1库】成功", localAddrString11, "outer_select", "MmHtJ1k26UaQXpix")
}

func testDb10() {
	Db, err := sqlx.Open("mysql", "outer_select:4Q7NNs5R2CT5Ki8I@tcp("+localAddrString12+")/db_ex_business?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【用户灾备2库】失败", err)
	}
	fmt.Println("链接【用户灾备2库】成功", localAddrString12, "outer_select", "4Q7NNs5R2CT5Ki8I")
}

func testDb11() {
	Db, err := sqlx.Open("mysql", "outer_select:qLVMI32Z0EnQ3Kg8@tcp("+localAddrString13+")/db_ex_business?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【用户灾备3库】失败", err)
	}
	fmt.Println("链接【用户灾备3库】成功", localAddrString13, "outer_select", "qLVMI32Z0EnQ3Kg8")
}

func testDb12() {
	Db, err := sqlx.Open("mysql", "outer_select:ulzFN72eVWdMjKBl@tcp("+localAddrString14+")/db_ex_business?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【用户灾备4库】失败", err)
	}
	fmt.Println("链接【用户灾备4库】成功", localAddrString14, "outer_select", "ulzFN72eVWdMjKBl")
}

func testDb13() {
	Db, err := sqlx.Open("mysql", "outer_select:JWyaMUKb89CuToM9@tcp("+localAddrString15+")/db_ex_business?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【用户灾备5库】失败", err)
	}
	fmt.Println("链接【用户灾备5库】成功", localAddrString15, "outer_select", "JWyaMUKb89CuToM9")
}

func testDb14() {
	Db, err := sqlx.Open("mysql", "inner_select:uzv1iavGAjVa0Bn7@tcp("+localAddrString16+")/db_ex_big_class?charset=utf8mb4&parseTime=True&loc=Local")
	if err = Db.Ping(); err != nil {
		fmt.Println("链接【班课数据库】失败", err)
	}
	fmt.Println("链接【班课数据库】成功", localAddrString16, "inner_select", "uzv1iavGAjVa0Bn7")
}
