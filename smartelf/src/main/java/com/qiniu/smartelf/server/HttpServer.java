package com.qiniu.smartelf.server;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.codec.http.HttpRequestDecoder;
import io.netty.handler.codec.http.HttpResponseEncoder;

public class HttpServer implements Runnable {
    private int port;

    public HttpServer(int port) {
        this.port = port;
    }

    public void run() {
        ServerBootstrap b;
        ChannelFuture f = null;
        EventLoopGroup bossGroup = new NioEventLoopGroup(3);
        EventLoopGroup workerGroup = new NioEventLoopGroup(10);
        try {
            b = new ServerBootstrap();
            b.group(bossGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                    .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        public void initChannel(SocketChannel ch)
                                throws Exception {
                            ch.pipeline().addLast(
                                    new HttpResponseEncoder());
                            ch.pipeline().addLast(
                                    new HttpRequestDecoder());
                            ch.pipeline().addLast(
                                    new HttpServerHandler());
                        }
                    }).option(ChannelOption.SO_BACKLOG, 128)
                    .childOption(ChannelOption.SO_KEEPALIVE, true);
            f = b.bind(port).sync();
            f.channel().closeFuture().sync();
        } catch (InterruptedException e) {
            System.out.printf("server interrupted: %s", e.toString());
        } finally {
            workerGroup.shutdownGracefully();
            bossGroup.shutdownGracefully();
            if (f != null) {
                f.channel().close();
            }
        }
    }
}
