package com.qiniu.smartelf.util;

public class Id {
    private int id;
    private static int idCounter = 0;

    public Id() {
        this.id = ++idCounter;
    }

    public int getId() {
        return id;
    }
}
