package com.qiniu.smartelf.server;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.ResponseBody;
import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.ByteToMessageDecoder;

import java.util.List;

public class TcpMsgDecoder extends ByteToMessageDecoder {
    protected void decode(ChannelHandlerContext ctx, ByteBuf buffer, List<Object> out) throws Exception {
        if (buffer.readableBytes() < 4) {
            return;
        }
        buffer.markReaderIndex();
        int dataLength = buffer.readInt();

        if (buffer.readableBytes() < dataLength) {
            buffer.resetReaderIndex();
            return;
        }
        byte[] data = new byte[dataLength];
        buffer.readBytes(data);

        ObjectMapper objectMapper = new ObjectMapper();
        ResponseBody rsp = objectMapper.readValue(data, ResponseBody.class);
        out.add(rsp);
    }
}
