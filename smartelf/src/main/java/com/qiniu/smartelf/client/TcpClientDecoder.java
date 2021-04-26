package com.qiniu.smartelf.client;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.qiniu.smartelf.model.RequestBody;
import com.qiniu.smartelf.model.ResponseBody;
import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.ByteToMessageDecoder;

import java.util.List;

public class TcpClientDecoder extends ByteToMessageDecoder {
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
        RequestBody req = objectMapper.readValue(data, RequestBody.class);
        out.add(req);
    }
}
