#pragma once

#include <string>

using namespace std;

// TEMPLATE METHOD: "openDocument" - template method, specialization implemented in derived classes

class Document{
public:

};

class Reader {
public:

    bool openDocument( const std::string & _docName ){

        if( ! canOpen(_docName) ){
            return false;
        }

        Document * doc = read();

        const bool infoPresence = getInfo();
    }

    void closeDocument(){
        // ...
    }

    virtual bool canOpen( const std::string & _docName ) = 0;
    virtual Document * read() = 0;
    virtual bool getInfo() = 0;

private:


};

class PDFReader : public Reader {
public:

    virtual bool canOpen( const std::string & _docName ) override {
        // ...
    }

    virtual Document * read() override {
        // ...
    }

    virtual bool getInfo() override {
        // ...
    }

private:


};
