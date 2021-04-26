package com.qiniu.smartelf.server;

import io.netty.channel.Channel;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class MsgChannelMap {
    private static Map<Integer, Channel> map = new ConcurrentHashMap<Integer, Channel>();

    public static void add(int id, Channel channel) {
        map.put(id, channel);
    }

    public static Channel get(int id) {
        return map.get(id);
    }

    public static void deleteByChannel(Channel channel) {
        for (Map.Entry entry : map.entrySet()) {
            if (entry.getValue() == channel) {
                map.remove(entry.getKey());
            }
        }
    }

    public static void deleteById(int id) {
        map.remove(id);
    }
}
