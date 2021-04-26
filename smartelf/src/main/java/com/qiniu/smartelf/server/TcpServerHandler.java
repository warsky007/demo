package com.qiniu.smartelf.server;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.ResponseBody;
import io.netty.buffer.Unpooled;
import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpResponse;

import static io.netty.handler.codec.http.HttpHeaderNames.CONTENT_LENGTH;
import static io.netty.handler.codec.http.HttpHeaderNames.CONTENT_TYPE;
import static io.netty.handler.codec.http.HttpResponseStatus.OK;
import static io.netty.handler.codec.http.HttpVersion.HTTP_1_1;

public class TcpServerHandler extends ChannelInboundHandlerAdapter {
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg)
            throws Exception {
        if (msg instanceof ResponseBody) {
            ResponseBody rsp = (ResponseBody) msg;
            if (rsp.id == 0) {
                System.out.printf("agent %s register\n", rsp.agentId);
                AgentChannelMap.add(rsp.agentId, ctx.channel());
            } else {
                Channel httpChannel = MsgChannelMap.get(rsp.id);
                if (httpChannel != null) {
                    MsgChannelMap.deleteById(rsp.id);
                    FullHttpResponse response = new DefaultFullHttpResponse(
                            HTTP_1_1, OK, Unpooled.wrappedBuffer(new ObjectMapper().writeValueAsBytes(msg)));
                    response.headers().set(CONTENT_TYPE, "application/json");
                    response.headers().set(CONTENT_LENGTH,
                            response.content().readableBytes());
                    httpChannel.writeAndFlush(response);
                }
            }
        }
    }

    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
        ctx.flush();
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        System.out.printf("agent %s offline\n", AgentChannelMap.getIdByChannel(ctx.channel()));
        AgentChannelMap.deleteByChannel(ctx.channel());
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause)
            throws Exception {
        cause.printStackTrace();
        ctx.close();
    }
}
