package org.qiniu;

import com.googlecode.jsonrpc4j.JsonRpcBasicServer;
import io.grpc.Server;

import java.io.IOException;
import java.io.PrintStream;

public class Main {
    public static void main(String[] args) {
//        Properties props = new Properties();
//        /* 定义kakfa 服务的地址，不需要将所有broker指定上 */
//        props.put("bootstrap.servers", "localhost:9092");
//        /* 制定consumer group */
//        props.put("group.id", "test");
//        /* 是否自动确认offset */
//        props.put("enable.auto.commit", "true");
//        /* 自动确认offset的时间间隔 */
//        props.put("auto.commit.interval.ms", "1000");
//        props.put("session.timeout.ms", "30000");
//        /* key的序列化类 */
//        props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
//        /* value的序列化类 */
//        props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
//        /* 定义consumer */
//        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(props);
//        /* 消费者订阅的topic, 可同时订阅多个 */
//        consumer.subscribe(Arrays.asList("foo", "bar"));
//
//        /* 读取数据，读取超时时间为100ms */
//        while (true) {
//            ConsumerRecords<String, String> records = consumer.poll(100);
//            for (ConsumerRecord<String, String> record : records)
//                System.out.printf("offset = %d, key = %s, value = %s", record.offset(), record.key(), record.value());
//        }
        PrintStream pipe = System.out;
        System.setOut(System.err);

        System.out.println("client starting");
        JsonRpcBasicServer server = new JsonRpcBasicServer(new DemoServiceImply(), DemoService.class);
        while (true) {
            try {
                int ret = server.handleRequest(System.in, pipe);
                System.out.printf("handle return %d\n", ret);
            } catch (IOException e) {
                System.out.println(e.toString());
                break;
            }
        }

        System.out.println("client end");
    }
}
