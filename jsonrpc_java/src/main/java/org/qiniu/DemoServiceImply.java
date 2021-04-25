package org.qiniu;

import org.qiniu.model.CountReply;
import org.qiniu.model.CountRequest;

import java.util.HashMap;

public class DemoServiceImply implements DemoService {
    public CountReply Test(CountRequest req) {
        Integer sum = req.a + req.b;
        CountReply reply = new CountReply();
        reply.message = "test message";
        reply.data = new HashMap<String,Object>();
        reply.data.put("value", sum);
        return reply;
    }
}
