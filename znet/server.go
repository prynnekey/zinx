package znet

import (
	"fmt"
	"net"

	"github.com/prynnekey/study/zinx/ziface"
)

type Server struct {
	Name      string // 服务器的名称
	IPVersion string // //tcp4 or other
	Ip        string // 服务器绑定的ip
	Port      int    // 服务器绑定的端口
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at Ip : %s, Port : %d\n", s.Ip, s.Port)

	go func() {
		// 1. 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Printf("resolve tcp addr error: %v\n", err)
			return
		}

		// 2. 监听服务器地址
		listner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("listen %s error: %v\n", s.IPVersion, err)
			return
		}

		fmt.Printf("start Zinx server successfully %s, listening...\n", s.Name)

		// 3. 阻塞等待客户端连接,处理客户端连接业务(读写)
		for {
			// 如果有客户端连接过来,会停止阻塞
			conn, err := listner.AcceptTCP()
			if err != nil {
				fmt.Printf("accept err: %v\n", err)
				continue
			}

			// 已经与客户端建立了连接,做一些事情
			go func() {
				for {
					buf := make([]byte, 512)
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("read err: %v\n", err)
						continue
					}

					fmt.Printf("server receive buf: %s, n = %d\n", buf, count)

					// 回显功能
					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Printf("write back buf err: %v\n", err)
						continue
					}
				}
			}()
		}

	}()
}

// 停止服务器
func (s *Server) Stop() {
	//TODO: 将一些服务器的资源、状态或者一些已经开辟的连接信息进行停止或回收
}

// 开启业务服务器
func (s *Server) Serve() {
	s.Start()

	// TODO: 做一些其他的业务

	// 阻塞在这里
	select {}
}
