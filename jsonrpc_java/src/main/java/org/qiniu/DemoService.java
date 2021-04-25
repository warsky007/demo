package org.qiniu;

import org.qiniu.model.CountReply;
import org.qiniu.model.CountRequest;

public interface DemoService {
    public CountReply Test(CountRequest req);
}
