package com.qiniu.smartelf.client;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.RequestBody;
import com.qiniu.smartelf.model.ResponseBody;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

public class TcpClientHandler extends ChannelInboundHandlerAdapter {
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg)
            throws Exception {
        if (msg instanceof RequestBody) {
            RequestBody req = (RequestBody) msg;

            if (req.content.equals("ping")) {
                ResponseBody rsp = new ResponseBody();
                rsp.id = req.id;
                rsp.content = "pang";
                ctx.channel().writeAndFlush(rsp);
            }
        }
    }
}
