#pragma once

// std
#include <string>
#include <vector>

// VISITOR: division Structre & Algorithm

class Visitor;

// ----------------------------------
// Element
// ----------------------------------
class Hardware {
public:
    Hardware(){}
    virtual ~Hardware(){}

    virtual void accept( Visitor * _visitor ) = 0;
    int price;
    std::string name;
};

class Floppy : public Hardware{
public:
    virtual void accept( Visitor * _visitor ) override;
};

class Processor : public Hardware{
public:
    virtual void accept( Visitor * _visitor ) override;
};

// ----------------------------------
// Visitor
// ----------------------------------
class Visitor{
public:

    Visitor(){}
    virtual ~Visitor(){}

    virtual void visit( Floppy * _floppy ) = 0;
    virtual void visit( Processor * _Cpu ) = 0;
};

class VisitorPrice : public Visitor {
public:

    VisitorPrice(){}
    ~VisitorPrice(){}

    virtual void visit( Floppy * _floppy ) override;
    virtual void visit( Processor * _processor ) override;
private:
    int total;
};

class VisitorInventory : public Visitor {
public:

    VisitorInventory(){}
    ~VisitorInventory(){}

    virtual void visit( Floppy * _floppy ) override;
    virtual void visit( Processor * _processor ) override;
private:
    std::vector<std::string> m_namesList;
};








