#pragma once

#include <string>

// STATE:

class TCPState {
public:
    ~TCPState(){}
    virtual void open() = 0;
    virtual void close() = 0;
    virtual void transmit( const std::string & _msg ) = 0;
};

class TCPEstablished {
public:
    virtual void open() override {

    }
    virtual void close() override {

    }
    virtual void transmit( const std::string & _msg ){

    }
};

class TCPClosed {
public:
    virtual void open() override {

    }
    virtual void close() override {

    }
    virtual void transmit( const std::string & _msg ){

    }
};

class TCPListen {
public:
    virtual void open() override {

    }
    virtual void close() override {

    }
    virtual void transmit( const std::string & _msg ){

    }
};




