package com.qiniu.smartelf.server;

import io.netty.channel.Channel;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class AgentChannelMap {
    private static Map<String, Channel> map = new ConcurrentHashMap<String, Channel>();

    public static void add(String id, Channel channel) {
        map.put(id, channel);
    }

    public static Channel get(String id) {
        return map.get(id);
    }

    public static void deleteByChannel(Channel channel) {
        for (Map.Entry entry : map.entrySet()) {
            if (entry.getValue() == channel) {
                map.remove(entry.getKey());
            }
        }
    }

    public static void deleteById(String id) {
        map.remove(id);
    }

    public static String getIdByChannel(Channel channel) {
        for (Map.Entry entry : map.entrySet()) {
            if (entry.getValue() == channel) {
                return entry.getKey().toString();
            }
        }
        return null;
    }
}
