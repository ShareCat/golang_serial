package main

import (
	"time"
	"tools/snet/utils/conn"

	log "github.com/donnie4w/go-logger/logger"
)

func main() {
	conn, err := conn.ConnectServer("", 9000)
	if err != nil {
		log.Error(err)
		return
	}

	go func() {
		buffer := make([]byte, 2014)
		for {
			n, _ := conn.TcpConn.Read(buffer)
			if n > 0 {
				log.Debug(n, string(buffer[:n]))
			}
		}
	}()

	for {
		time.Sleep(time.Second)
		conn.TcpConn.Write([]byte("hello2,send from net client开展计算机科学基础课程两个星期了，我终于可以相信，最初的目标“零基础教育”是一个很好的出发点。不仅因为它能检验和改进我的学识和教学方法，而且因为零基础的人似乎更容易学会正确而干净的知识。这一期课程的成员来自中国大陆，香港，台湾，有各种各样的背景。有些同学没有理科背景，没上过大学，还有的完全是出于兴趣，还有的却已经博士毕业几年。他们短短两个星期以来的表现，思考问题的角度和深度，学习的态度和动力，让我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。我惊讶又欣慰。每当看到他们的进步，就觉得这一切的辛苦都是值得的。我正在改变一些人的人生。"))
	}

}
