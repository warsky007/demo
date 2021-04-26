package com.qiniu.smartelf.server;

public class Main {
    public static void main(String[] args) throws Exception {
        HttpServer httpServer = new HttpServer(8000);
        TcpServer tcpServer = new TcpServer(8001);
        Thread t1 = new Thread(httpServer);
        Thread t2 = new Thread(tcpServer);
        t1.start();
        t2.start();
        Thread.sleep(10000000);
    }
}
