package com.qiniu.smartelf.server;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.RequestBody;
import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.MessageToByteEncoder;


public class TcpMsgEncoder<H extends RequestBody> extends MessageToByteEncoder<Object> {

    protected void encode(ChannelHandlerContext ctx, Object msg, ByteBuf out) throws Exception {
        if (msg instanceof RequestBody) {
           byte[] bytes = new ObjectMapper().writeValueAsBytes(msg);
           out.writeInt(bytes.length);
           out.writeBytes(bytes);
        }
    }
}