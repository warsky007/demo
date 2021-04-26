package com.qiniu.smartelf.server;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

public class TcpServer implements Runnable {
    private int port;

    public TcpServer(int port) {
        this.port = port;
    }

    public void run() {
        ServerBootstrap b;
        ChannelFuture f = null;
        EventLoopGroup bossGroup = new NioEventLoopGroup(10);
        EventLoopGroup workerGroup = new NioEventLoopGroup(100);
        try {
            b = new ServerBootstrap();
            b.group(bossGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                    .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        public void initChannel(SocketChannel ch)
                                throws Exception {
                            ch.pipeline().addLast(
                                    new TcpMsgEncoder<>());
                            ch.pipeline().addLast(
                                    new TcpMsgDecoder());
                            ch.pipeline().addLast(
                                    new TcpServerHandler());
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
