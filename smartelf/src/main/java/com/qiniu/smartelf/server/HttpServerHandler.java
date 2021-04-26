package com.qiniu.smartelf.server;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.RequestBody;
import com.qiniu.smartelf.model.ResponseBody;
import com.qiniu.smartelf.util.Id;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpContent;

import static io.netty.handler.codec.http.HttpHeaderNames.CONTENT_LENGTH;
import static io.netty.handler.codec.http.HttpHeaderNames.CONTENT_TYPE;
import static io.netty.handler.codec.http.HttpResponseStatus.OK;
import static io.netty.handler.codec.http.HttpVersion.HTTP_1_1;

public class HttpServerHandler extends ChannelInboundHandlerAdapter {
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg)
            throws Exception {
        RequestBody body;

        if (msg instanceof HttpContent) {
            HttpContent httpContent = (HttpContent) msg;
            ByteBuf content = httpContent.content();
            byte[] bytes = new byte[content.readableBytes()];
            content.readBytes(bytes);
            content.release();

            ObjectMapper objectMapper = new ObjectMapper();
            body = objectMapper.readValue(bytes, RequestBody.class);
            Channel channel = AgentChannelMap.get(body.agentId);
            if (channel != null) {
                body.id = new Id().getId();
                channel.writeAndFlush(body);
                MsgChannelMap.add(body.id, ctx.channel());
            } else {
                ResponseBody rsp = new ResponseBody();
                rsp.agentId = body.agentId;
                rsp.content = new String("can't find agent id");
                FullHttpResponse response = new DefaultFullHttpResponse(
                        HTTP_1_1, OK, Unpooled.wrappedBuffer(new ObjectMapper().writeValueAsBytes(rsp)));
                response.headers().set(CONTENT_TYPE, "application/json");
                response.headers().set(CONTENT_LENGTH, response.content().readableBytes());
                ctx.channel().writeAndFlush(response);
            }
        }
    }

    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause)
            throws Exception {
        ctx.close();
    }
}
