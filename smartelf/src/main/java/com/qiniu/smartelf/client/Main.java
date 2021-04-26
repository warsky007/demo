package com.qiniu.smartelf.client;

import com.qiniu.smartelf.model.ResponseBody;
import io.netty.channel.Channel;

public class Main {
    public static void main(String[] args) throws Exception {
        TcpClient client = new TcpClient("127.0.0.1", 8001);
        client.start();
        Channel channel = client.getChannel();

        ResponseBody register = new ResponseBody();
        register.agentId = "123456";
        channel.writeAndFlush(register);
    }
}
