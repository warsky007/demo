package com.qiniu.smartelf.client;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.RequestBody;
import com.qiniu.smartelf.model.ResponseBody;
import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.MessageToByteEncoder;


public class TcpClientEncoder<H extends RequestBody> extends MessageToByteEncoder<Object> {

    protected void encode(ChannelHandlerContext ctx, Object msg, ByteBuf out) throws Exception {
        if (msg instanceof ResponseBody) {
           byte[] bytes = new ObjectMapper().writeValueAsBytes(msg);
           out.writeInt(bytes.length);
           out.writeBytes(bytes);
        }
    }
}