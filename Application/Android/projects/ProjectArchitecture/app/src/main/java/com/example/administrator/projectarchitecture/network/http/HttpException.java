
package com.example.administrator.projectarchitecture.network.http;


public class HttpException extends Exception {

    /**
     * serialVersionUID
     */
    private static final long serialVersionUID = 4010634120321127684L;

    public HttpException() {
        super();
    }

    public HttpException(String detailMessage, Throwable throwable) {
        super(detailMessage, throwable);
    }

    public HttpException(String detailMessage) {
        super(detailMessage);
    }

    public HttpException(Throwable throwable) {
        super(throwable);
    }


}
